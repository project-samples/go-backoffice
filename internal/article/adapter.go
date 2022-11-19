package article

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	q "github.com/core-go/sql"
	"reflect"
)

func NewArticleRepository(db *sql.DB, toArray func(interface{}) interface {
	driver.Valuer
	sql.Scanner
}) *ArticleAdapter {
	modelType := reflect.TypeOf(Article{})
	jsonColumnMap := q.MakeJsonColumnMap(modelType)
	keys, _ := q.FindPrimaryKeys(modelType)
	schema := q.CreateSchema(modelType)
	buildParam := q.GetBuild(db)
	return &ArticleAdapter{DB: db, ModelType: modelType, Keys: keys, Schema: schema, JsonColumnMap: jsonColumnMap,
		BuildParam: buildParam, toArray: toArray}
}

type ArticleAdapter struct {
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

func (r *ArticleAdapter) Load(ctx context.Context, id string) (*Article, error) {
	var articles []Article
	query := fmt.Sprintf("select * from articles where id = %s limit 1", r.BuildParam(1))
	err := q.QueryWithArray(ctx, r.DB, nil, &articles, r.toArray, query, id)
	if err != nil {
		return nil, err
	}
	if len(articles) > 0 {
		return &articles[0], nil
	}
	return nil, nil
}

func (r *ArticleAdapter) Create(ctx context.Context, article *Article) (int64, error) {
	query, args := q.BuildToInsertWithArray("articles", article, r.BuildParam, true, r.toArray, r.Schema)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (r *ArticleAdapter) Update(ctx context.Context, article *Article) (int64, error) {
	query, args := q.BuildToUpdateWithArray("articles", article, r.BuildParam, false, r.toArray, r.Schema)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (r *ArticleAdapter) Patch(ctx context.Context, article map[string]interface{}) (int64, error) {
	colMap := q.JSONToColumns(article, r.JsonColumnMap)
	query, args := q.BuildToPatchWithArray("articles", colMap, r.Keys, r.BuildParam, r.toArray, r.Schema.Fields)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (r *ArticleAdapter) Delete(ctx context.Context, id string) (int64, error) {
	query := "delete from articles where id = $1"
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
