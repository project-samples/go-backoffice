package product

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	q "github.com/core-go/sql"
	"reflect"
)

func NewProductRepository(db *sql.DB, toArray func(interface{}) interface {
	driver.Valuer
	sql.Scanner
}) *ProductAdapter {
	modelType := reflect.TypeOf(Product{})
	jsonColumnMap := q.MakeJsonColumnMap(modelType)
	keys, _ := q.FindPrimaryKeys(modelType)
	schema := q.CreateSchema(modelType)
	buildParam := q.GetBuild(db)
	return &ProductAdapter{DB: db, ModelType: modelType, Keys: keys, Schema: schema, JsonColumnMap: jsonColumnMap,
		BuildParam: buildParam, toArray: toArray}
}

type ProductAdapter struct {
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

func (r *ProductAdapter) Load(ctx context.Context, id string) (*Product, error) {
	var products []Product
	query := fmt.Sprintf("select * from products where id = %s limit 1", r.BuildParam(1))
	err := q.QueryWithArray(ctx, r.DB, nil, &products, r.toArray, query, id)
	if err != nil {
		return nil, err
	}
	if len(products) > 0 {
		return &products[0], nil
	}
	return nil, nil
}

func (r *ProductAdapter) Create(ctx context.Context, product *Product) (int64, error) {
	query, args := q.BuildToInsertWithArray("products", product, r.BuildParam, true, r.toArray, r.Schema)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, nil
	}
	return res.RowsAffected()
}

func (r *ProductAdapter) Update(ctx context.Context, product *Product) (int64, error) {
	query, args := q.BuildToUpdateWithArray("products", product, r.BuildParam, false, r.toArray, r.Schema)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, nil
	}
	return res.RowsAffected()
}

func (r *ProductAdapter) Patch(ctx context.Context, product map[string]interface{}) (int64, error) {
	colMap := q.JSONToColumns(product, r.JsonColumnMap)
	query, args := q.BuildToPatchWithArray("products", colMap, r.Keys, r.BuildParam, r.toArray, r.Schema.Fields)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (r *ProductAdapter) Delete(ctx context.Context, id string) (int64, error) {
	query := "delete from products where id = $1"
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
