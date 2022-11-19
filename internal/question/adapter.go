package question

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	q "github.com/core-go/sql"
	"reflect"
)

func NewQuestionRepository(db *sql.DB, toArray func(interface{}) interface {
	driver.Valuer
	sql.Scanner
}) *QuestionAdapter {
	modelType := reflect.TypeOf(Question{})
	jsonColumnMap := q.MakeJsonColumnMap(modelType)
	keys, _ := q.FindPrimaryKeys(modelType)
	schema := q.CreateSchema(modelType)
	buildParam := q.GetBuild(db)
	return &QuestionAdapter{DB: db, ModelType: modelType, Keys: keys, Schema: schema, JsonColumnMap: jsonColumnMap,
		BuildParam: buildParam, toArray: toArray}
}

type QuestionAdapter struct {
	DB            *sql.DB
	ModelType     reflect.Type
	Keys          []string
	Schema        *q.Schema
	JsonColumnMap map[string]string
	BuildParam    func(i int) string
	toArray       func(interface{}) interface {
		driver.Valuer
		sql.Scanner
	}
}

func (r *QuestionAdapter) Load(ctx context.Context, id string) (*Question, error) {
	var questions []Question
	query := fmt.Sprintf("select * from questions where id = %s limit 1", r.BuildParam(1))
	err := q.QueryWithArray(ctx, r.DB, nil, &questions, r.toArray, query, id)
	if err != nil {
		return nil, err
	}
	if len(questions) > 0 {
		return &questions[0], nil
	}
	return nil, nil
}

func (r *QuestionAdapter) Create(ctx context.Context, question *Question) (int64, error) {
	query, args := q.BuildToInsertWithArray("questions", question, r.BuildParam, true, r.toArray, r.Schema)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, nil
	}
	return res.RowsAffected()
}

func (r *QuestionAdapter) Update(ctx context.Context, question *Question) (int64, error) {
	query, args := q.BuildToUpdateWithArray("questions", question, r.BuildParam, false, r.toArray, r.Schema)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, nil
	}
	return res.RowsAffected()
}

func (r *QuestionAdapter) Patch(ctx context.Context, question map[string]interface{}) (int64, error) {
	colMap := q.JSONToColumns(question, r.JsonColumnMap)
	query, args := q.BuildToPatchWithArray("questions", colMap, r.Keys, r.BuildParam, r.toArray, r.Schema.Fields)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (r *QuestionAdapter) Delete(ctx context.Context, id string) (int64, error) {
	query := "delete from questions where id = $1"
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
