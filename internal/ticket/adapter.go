package ticket

import (
	"context"
	"database/sql"
	"fmt"
	q "github.com/core-go/sql"
	"reflect"
)

func NewTicketRepository(db *sql.DB) *TicketAdapter {
	modelType := reflect.TypeOf(Ticket{})
	jsonColumnMap := q.MakeJsonColumnMap(modelType)
	keys, _ := q.FindPrimaryKeys(modelType)
	schema := q.CreateSchema(modelType)
	return &TicketAdapter{DB: db, ModelType: modelType, Keys: keys, Schema: schema, JsonColumnMap: jsonColumnMap}
}

type TicketAdapter struct {
	DB            *sql.DB
	ModelType     reflect.Type
	Keys          []string
	Schema        *q.Schema
	JsonColumnMap map[string]string
}

func (r *TicketAdapter) Load(ctx context.Context, id string) (*Ticket, error) {
	var tickets []Ticket
	query := fmt.Sprintf("select * from tickets where id = %s limit 1", q.BuildParam(1))
	err := q.Query(ctx, r.DB, nil, &tickets, query, id)
	if err != nil {
		return nil, err
	}
	if len(tickets) > 0 {
		return &tickets[0], nil
	}
	return nil, nil
}

func (r *TicketAdapter) Create(ctx context.Context, ticket *Ticket) (int64, error) {
	query, args := q.BuildToInsert("tickets", ticket, q.BuildParam, r.Schema)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, nil
	}
	return res.RowsAffected()
}

func (r *TicketAdapter) Update(ctx context.Context, ticket *Ticket) (int64, error) {
	query, args := q.BuildToUpdate("tickets", ticket, q.BuildParam, r.Schema)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, nil
	}
	return res.RowsAffected()
}

func (r *TicketAdapter) Patch(ctx context.Context, ticket map[string]interface{}) (int64, error) {
	colMap := q.JSONToColumns(ticket, r.JsonColumnMap)
	query, args := q.BuildToPatch("tickets", colMap, r.Keys, q.BuildParam, r.Schema.Fields)
	res, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (r *TicketAdapter) Delete(ctx context.Context, id string) (int64, error) {
	query := "delete from tickets where id = ?"
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