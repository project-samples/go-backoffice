package company

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	q "github.com/core-go/sql"
)

type companyRole struct {
	CompanyId string `json:"companyId,omitempty" gorm:"column:companyId;primary_key" bson:"_id,omitempty" validate:"required,max=20,code"`
	UserId    string `json:"userId,omitempty" gorm:"column:userId;primary_key" bson:"_id,omitempty" dynamodbav:"userId,omitempty" firestore:"userId,omitempty" validate:"max=40"`
}

type CompanyAdapter struct {
	db                *sql.DB
	driver            string
	BuildParam        func(int) string
	CheckDelete       string
	Map               map[string]int
	FieldIndex        map[string]int
	modelType         reflect.Type
	companySchema     *q.Schema
	companyRoleSchema *q.Schema
}

func NewCompanyRepository(db *sql.DB) (*CompanyAdapter, error) {
	modelType := reflect.TypeOf(Company{})
	buildParam := q.GetBuild(db)
	var r companyRole
	subType := reflect.TypeOf(r)
	m, err := q.GetColumnIndexes(subType)
	if err != nil {
		return nil, err
	}
	m2, err := q.GetColumnIndexes(modelType)
	if err != nil {
		return nil, err
	}
	companySchema := q.CreateSchema(modelType)
	companyRoleSchema := q.CreateSchema(subType)
	driver := q.GetDriver(db)
	return &CompanyAdapter{db: db, driver: driver, BuildParam: buildParam, modelType: modelType, Map: m, FieldIndex: m2, companySchema: companySchema, companyRoleSchema: companyRoleSchema}, nil
}

func (s *CompanyAdapter) Load(ctx context.Context, id string) (*Company, error) {
	var companies []Company
	sql := fmt.Sprintf("select * from companies where companyId = %s", s.BuildParam(1))
	er1 := q.Query(ctx, s.db, s.FieldIndex, &companies, sql, id)
	if er1 != nil {
		return nil, er1
	}
	if len(companies) == 0 {
		return nil, nil
	}
	company := companies[0]
	users, er3 := getUsers(ctx, s.db, id, s.BuildParam, s.Map)
	if er3 != nil {
		return nil, er3
	}
	if len(users) > 0 {
		company.Users = users
	}
	return &company, nil
}

func (s *CompanyAdapter) Create(ctx context.Context, company *Company) (int64, error) {
	sts, err := buildInsertCompanyStatements(company, s.driver, s.BuildParam, s.companySchema, s.companyRoleSchema)
	if err != nil {
		return 0, err
	}
	return sts.Exec(ctx, s.db)
}

func (s *CompanyAdapter) Update(ctx context.Context, company *Company) (int64, error) {
	sts, err := buildUpdateCompanyStatements(company, s.driver, s.BuildParam, s.companySchema, s.companyRoleSchema)
	if err != nil {
		return 0, err
	}
	return sts.Exec(ctx, s.db)
}

func (s *CompanyAdapter) Patch(ctx context.Context, company map[string]interface{}) (int64, error) {
	sts, err := buildPatchCompanyStatements(company, s.BuildParam, s.modelType)
	if err != nil {
		return 0, err
	}
	return sts.Exec(ctx, s.db)
}

func (s *CompanyAdapter) Delete(ctx context.Context, id string) (int64, error) {
	if len(s.CheckDelete) > 0 {
		exist, er0 := checkExist(s.db, s.CheckDelete, id)
		if exist || er0 != nil {
			return -1, er0
		}
	}
	sts, er1 := buildDeleteCompanyStatements(id, s.BuildParam)
	if er1 != nil {
		return 0, er1
	}
	return sts.Exec(ctx, s.db)
}

func getUsers(ctx context.Context, db *sql.DB, companyId string, buildParam func(int) string, m map[string]int) ([]string, error) {
	var companyUsers []companyRole
	users := make([]string, 0)
	query := fmt.Sprintf(`select userId from companyusers where companyId = %s`, buildParam(1))
	err := q.Query(ctx, db, m, &companyUsers, query, companyId)
	if err != nil {
		return nil, err
	}
	for _, e := range companyUsers {
		users = append(users, e.UserId)
	}
	return users, nil
}

func buildInsertCompanyStatements(company *Company, driver string, buildParam func(int) string, companySchema *q.Schema, companyRoleSchema *q.Schema) (q.Statements, error) {
	modules, er1 := buildCompanyModules(company.CompanyId, company.Users)
	if er1 != nil {
		return nil, er1
	}
	sts := q.NewStatements(true)
	sts.Add(q.BuildToInsert("companies", company, buildParam, companySchema))
	if modules != nil {
		query, args, er2 := q.BuildToInsertBatch("companyusers", modules, driver, companyRoleSchema)
		if er2 != nil {
			return sts, er2
		}
		sts.Add(query, args)
	}
	return sts, nil
}

func buildCompanyModules(companyID string, users []string) ([]companyRole, error) {
	if users == nil || len(users) <= 0 {
		return nil, nil
	}
	modules := make([]companyRole, 0)
	for _, p := range users {
		m := toCompanyModules(companyID, p)
		m.CompanyId = companyID
		m.UserId = users[0]
		modules = append(modules, m)
	}
	return modules, nil
}

func toCompanyModules(CompanyID string, menu string) companyRole {
	s := strings.Split(menu, " ")
	p := companyRole{CompanyId: CompanyID, UserId: s[0]}
	return p
}

func buildUpdateCompanyStatements(company *Company, driver string, buildParam func(int) string, companySchema *q.Schema, companyRoleSchema *q.Schema) (q.Statements, error) {
	modules, er1 := buildCompanyModules(company.CompanyId, company.Users)
	if er1 != nil {
		return nil, er1
	}
	sts := q.NewStatements(true)
	sts.Add(q.BuildToUpdate("companies", company, buildParam, companySchema))

	deleteModules := fmt.Sprintf("delete from companyusers where companyId = %s", buildParam(1))
	sts.Add(deleteModules, []interface{}{company.CompanyId})

	if modules != nil {
		query, args, er2 := q.BuildToInsertBatch("companyusers", modules, driver, companyRoleSchema)
		if er2 != nil {
			return sts, er2
		}
		sts.Add(query, args)
	}
	return sts, nil
}

func checkExist(db *sql.DB, sql string, args ...interface{}) (bool, error) {
	rows, err := db.Query(sql, args...)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		return true, nil
	}
	return false, nil
}

func buildDeleteCompanyStatements(id string, buildParam func(int) string) (q.Statements, error) {
	sts := q.NewStatements(false)

	deleteModules := fmt.Sprintf("delete from companyusers where companyId = %s", buildParam(1))
	sts.Add(deleteModules, []interface{}{id})

	deleteRole := fmt.Sprintf("delete from companies where companyId = %s", buildParam(1))
	sts.Add(deleteRole, []interface{}{id})

	return sts, nil
}

func buildPatchCompanyStatements(json map[string]interface{}, buildParam func(int) string, modelType reflect.Type) (q.Statements, error) {
	sts := q.NewStatements(true)
	primaryKeyColumns, _ := q.FindPrimaryKeys(modelType)
	jsonColumnMap := q.MakeJsonColumnMap(modelType)
	columnMap := q.JSONToColumns(json, jsonColumnMap)
	sts.Add(q.BuildToPatch("companies", columnMap, primaryKeyColumns, buildParam))
	if json["users"] != nil {
		deleteModules := fmt.Sprintf("delete from companyusers where companyId = '%s';", buildParam(1))
		sts.Add(deleteModules, []interface{}{json["companyId"]})
		a := json["users"]
		t, ok := a.([]string)
		if ok {
			for i := 0; i < len(t); i++ {
				insertModules := fmt.Sprintf("insert into companyusers values ('%s','%s');", buildParam(1), buildParam(2))
				sts.Add(insertModules, []interface{}{json["companyId"], t[i]})
			}
		}
	}
	return sts, nil
}
