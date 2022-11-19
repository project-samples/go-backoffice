package test

import (
	"context"
	"database/sql"
	"fmt"
	q "github.com/core-go/sql"
	"reflect"
)

func NewTestRepository(db *sql.DB) *TestAdapter {
	modelType := reflect.TypeOf(Test{})
	jsonColumnMap := q.MakeJsonColumnMap(modelType)
	keys, _ := q.FindPrimaryKeys(modelType)
	schema := q.CreateSchema(modelType)
	return &TestAdapter{DB: db, ModelType: modelType, Keys: keys, Schema: schema, JsonColumnMap: jsonColumnMap}
}

type TestAdapter struct {
	DB            *sql.DB
	ModelType     reflect.Type
	Keys          []string
	Schema        *q.Schema
	JsonColumnMap map[string]string
}

func (r *TestAdapter) Load(ctx context.Context, id string) (*Test, error) {
	var tests []Test
	query := fmt.Sprintf("select * from tests where testid = %s limit 1", q.BuildParam(1))
	err := q.Query(ctx, r.DB, nil, &tests, query, id)
	if err != nil {
		return nil, err
	}
	if len(tests) > 0 {
		return &tests[0], nil
	}
	return nil, nil
}

func (r *TestAdapter) Create(ctx context.Context, test *Test) (int64, error) {
	query, args := q.BuildToInsert("tests", test, q.BuildParam, r.Schema)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, nil
	}
	return res.RowsAffected()
}

func (r *TestAdapter) Update(ctx context.Context, test *Test) (int64, error) {
	query, args := q.BuildToUpdate("tests", test, q.BuildParam, r.Schema)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, nil
	}
	return res.RowsAffected()
}

func (r *TestAdapter) Patch(ctx context.Context, test map[string]interface{}) (int64, error) {
	colMap := q.JSONToColumns(test, r.JsonColumnMap)
	query, args := q.BuildToPatch("tests", colMap, r.Keys, q.BuildParam, r.Schema.Fields)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (r *TestAdapter) Delete(ctx context.Context, id string) (int64, error) {
	query := "delete from tests where testid = ?"
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
