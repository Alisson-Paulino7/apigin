package base

import (
	"context"
	"database/sql"
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
	return nil
}