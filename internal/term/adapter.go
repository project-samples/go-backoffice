package term

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	q "github.com/core-go/sql"
	"reflect"
)

func NewTermRepository(db *sql.DB, toArray func(interface{}) interface {
	driver.Valuer
	sql.Scanner
}) (*TermAdapter, error) {
	modelType := reflect.TypeOf(Term{})
	fieldsIndex, err := q.GetColumnIndexes(modelType)
	if err != nil {
		return nil, err
	}
	jsonColumnMap := q.MakeJsonColumnMap(modelType)
	keys, _ := q.FindPrimaryKeys(modelType)
	schema := q.CreateSchema(modelType)
	buildParam := q.GetBuild(db)
	return &TermAdapter{DB: db, ModelType: modelType, FieldsIndex: fieldsIndex, Keys: keys, Schema: schema, JsonColumnMap: jsonColumnMap,
		BuildParam: buildParam, toArray: toArray}, nil
}

type TermAdapter struct {
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

func (r *TermAdapter) Load(ctx context.Context, id string) (*Term, error) {
	var terms []Term
	query := fmt.Sprintf("select * from terms where id = %s limit 1", r.BuildParam(1))
	err := q.QueryWithArray(ctx, r.DB, r.FieldsIndex, &terms, r.toArray, query, id)
	if err != nil {
		return nil, err
	}
	if len(terms) > 0 {
		return &terms[0], nil
	}
	return nil, nil
}

func (r *TermAdapter) Create(ctx context.Context, term *Term) (int64, error) {
	query, args := q.BuildToInsertWithArray("terms", term, r.BuildParam, true, r.toArray, r.Schema)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (r *TermAdapter) Update(ctx context.Context, term *Term) (int64, error) {
	query, args := q.BuildToUpdateWithArray("terms", term, r.BuildParam, false, r.toArray, r.Schema)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (r *TermAdapter) Patch(ctx context.Context, term map[string]interface{}) (int64, error) {
	colMap := q.JSONToColumns(term, r.JsonColumnMap)
	query, args := q.BuildToPatchWithArray("terms", colMap, r.Keys, r.BuildParam, r.toArray, r.Schema.Fields)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (r *TermAdapter) Delete(ctx context.Context, id string) (int64, error) {
	query := "delete from terms where id = $1"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return -1, nil
	}
	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}
