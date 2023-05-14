package test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	q "github.com/core-go/sql"
	"reflect"
)

func NewTestRepository(db *sql.DB,
	toArray func(interface{}) interface {
		driver.Valuer
		sql.Scanner
	},
) (*TestAdapter, error) {
	modelType := reflect.TypeOf(Test{})
	fieldsIndex, err := q.GetColumnIndexes(modelType)
	if err != nil {
		return nil, err
	}
	jsonColumnMap := q.MakeJsonColumnMap(modelType)
	keys, _ := q.FindPrimaryKeys(modelType)
	schema := q.CreateSchema(modelType)
	buildParam := q.GetBuild(db)
	return &TestAdapter{DB: db, ModelType: modelType, FieldsIndex: fieldsIndex, Keys: keys, Schema: schema, JsonColumnMap: jsonColumnMap, BuildParam: buildParam, toArray: toArray}, nil
}

type TestAdapter struct {
	DB            *sql.DB
	ModelType     reflect.Type
	FieldsIndex   map[string]int
	Keys          []string
	Schema        *q.Schema
	JsonColumnMap map[string]string
	BuildParam    func(i int) string
	toArray       func(interface{}) interface {
		driver.Valuer
		sql.Scanner
	}
}

func (r *TestAdapter) Load(ctx context.Context, id string) (*Test, error) {
	var tests []Test
	query := fmt.Sprintf("select * from tests where testid = %s limit 1", r.BuildParam(1))
	err := q.QueryWithArray(ctx, r.DB, r.FieldsIndex, &tests, r.toArray, query, id)
	if err != nil {
		return nil, err
	}
	if len(tests) > 0 {
		return &tests[0], nil
	}
	return nil, nil
}

func (r *TestAdapter) Create(ctx context.Context, test *Test) (int64, error) {
	query, args := q.BuildToInsertWithArray("tests", test, r.BuildParam, false, r.toArray, r.Schema)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (r *TestAdapter) Update(ctx context.Context, test *Test) (int64, error) {
	query, args := q.BuildToUpdateWithArray("tests", test, r.BuildParam, false, r.toArray, r.Schema)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (r *TestAdapter) Patch(ctx context.Context, test map[string]interface{}) (int64, error) {
	colMap := q.JSONToColumns(test, r.JsonColumnMap)
	query, args := q.BuildToPatchWithArray("tests", colMap, r.Keys, r.BuildParam, r.toArray, r.Schema.Fields)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (r *TestAdapter) Delete(ctx context.Context, id string) (int64, error) {
	query := "delete from tests where testid = " + r.BuildParam(1)
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return -1, err
	}
	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}
