package base

import (
	"context"
	"database/sql"

	"github.com/alisson-paulino7/apigin/infrastructure/persistence/user/base"
)

type repository struct {
	ctx context.Context
	tx *sql.Tx
}

func NewRepository(ctx context.Context, tx *sql.Tx) BaseInterface {
	return &repository{
		ctx: ctx,
		tx: tx,
	}
}

func (r *repository) Insert(user *User) error {
	params := &base.InsertParams{}
	if user.Name != nil {
		params.Name = *user.Name
	}
	if user.Age != nil {
		params.Age = int32(*user.Age)
	}
	if user.Document != nil {
		params.Document = *user.Document
	}
	row, err := base.New(r.tx).Insert(r.ctx, params)
	if err != nil {
		return err
	}
	user.ID = &row.ID
	user.CreatedAt = &row.CreatedAt
	
	return nil
}