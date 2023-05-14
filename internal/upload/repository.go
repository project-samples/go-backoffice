package uploads

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type StorageRepository interface {
	Load(ctx context.Context, id string) (*Attachment, error)
	Update(ctx context.Context, id string, attachments Attachment) (int64, error)
}

func NewRepository(DB *sql.DB,
	Table string,
	columns FieldColumn, toArray func(interface{}) interface {
		driver.Valuer
		sql.Scanner
	}) *SqlRepository {
	return &SqlRepository{DB: DB, Table: Table, Columns: &columns, toArray: toArray}
}

type FieldColumn struct {
	Id   string
	File string
}

type SqlRepository struct {
	DB      *sql.DB
	Table   string
	Columns *FieldColumn
	toArray func(interface{}) interface {
		driver.Valuer
		sql.Scanner
	}
}

func (s *SqlRepository) Load(ctx context.Context, id string) (*Attachment, error) {
	var attachments = make([]Attachment, 0)
	query := fmt.Sprintf("select %s from %s where %s= $1", s.Columns.File, s.Table, s.Columns.Id)
	row := s.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(s.toArray(&attachments))
	if err != nil {
		return nil, err
	}
	if len(attachments) > 0 {
		return &attachments[0], nil
	}
	return nil, err
}

func (s *SqlRepository) Update(ctx context.Context, id string, attachments Attachment) (int64, error) {
	query := fmt.Sprintf("update %s set %s = $1 where %s =$2", s.Table, s.Columns.File, s.Columns.Id)
	stmt, er0 := s.DB.Prepare(query)
	if er0 != nil {
		return -1, er0
	}
	res, err := stmt.ExecContext(ctx, attachments, id)
	if err != nil {
		return -1, err
	}
	row, er2 := res.RowsAffected()

	if row < 0 {
		return -1, er2
	}
	return row, er2
}
