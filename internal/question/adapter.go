package question

import (
	"context"
	"database/sql"
	"fmt"
	q "github.com/core-go/sql"
	"reflect"
)

func NewQuestionRepository(db *sql.DB) *QuestionAdapter {
	modelType := reflect.TypeOf(Question{})
	jsonColumnMap := q.MakeJsonColumnMap(modelType)
	keys, _ := q.FindPrimaryKeys(modelType)
	schema := q.CreateSchema(modelType)
	return &QuestionAdapter{DB: db, ModelType: modelType, Keys: keys, Schema: schema, JsonColumnMap: jsonColumnMap}
}

type QuestionAdapter struct {
	DB            *sql.DB
	ModelType     reflect.Type
	Keys          []string
	Schema        *q.Schema
	JsonColumnMap map[string]string
}

func (r *QuestionAdapter) Load(ctx context.Context, id string) (*Question, error) {
	var questions []Question
	query := fmt.Sprintf("select * from questions where id = %s limit 1", q.BuildParam(1))
	err := q.Query(ctx, r.DB, nil, &questions, query, id)
	if err != nil {
		return nil, err
	}
	if len(questions) > 0 {
		return &questions[0], nil
	}
	return nil, nil
}

func (r *QuestionAdapter) Create(ctx context.Context, question *Question) (int64, error) {
	query, args := q.BuildToInsert("questions", question, q.BuildParam, r.Schema)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, nil
	}
	return res.RowsAffected()
}

func (r *QuestionAdapter) Update(ctx context.Context, question *Question) (int64, error) {
	query, args := q.BuildToUpdate("questions", question, q.BuildParam, r.Schema)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, nil
	}
	return res.RowsAffected()
}

func (r *QuestionAdapter) Patch(ctx context.Context, question map[string]interface{}) (int64, error) {
	colMap := q.JSONToColumns(question, r.JsonColumnMap)
	query, args := q.BuildToPatch("questions", colMap, r.Keys, q.BuildParam, r.Schema.Fields)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (r *QuestionAdapter) Delete(ctx context.Context, id string) (int64, error) {
	query := "delete from questions where id = ?"
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
