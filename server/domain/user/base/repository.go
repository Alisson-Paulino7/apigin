package base

import (
	"context"
	"database/sql"

	"github.com/alisson-paulino7/apigin/infrastructure/persistence/user/base"
	"github.com/alisson-paulino7/apigin/utils"
	"github.com/google/uuid"
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

func (r *repository) FindByID(id uuid.UUID) (*User, error) {
	row, err := base.New(r.tx).FindByID(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:        &row.ID,
		CreatedAt: &row.CreatedAt,
		Name:      &row.Name,
		Age:       utils.PointerInt(int(row.Age)),
		Document:  &row.Document,
	},nil
}

func (r *repository) Update(user *User) error {
	params := &base.UpdateParams{}
	if user.ID != nil {
		params.ID = *user.ID
	}
	if user.Name != nil {
		params.Name = *user.Name
	}
	if user.Age != nil {
		params.Age = int32(*user.Age)
	}
	if user.Document != nil {
		params.Document = *user.Document
	}
	if err := base.New(r.tx).Update(r.ctx, params); err != nil {
		return err
	}
	return nil
}

func (r *repository) Delete(id uuid.UUID) error {	
	if err := base.New(r.tx).Delete(r.ctx, id); err != nil {
		return err
	}
	return nil
}

func (r *repository) FindAll() ([]*User, error) {
	rows, err := base.New(r.tx).FindAll(r.ctx)
	if err != nil {
		return nil, err
	}

	var users []*User
	for _, row := range rows {
		users = append(users, &User{
			ID:        &row.ID,
			CreatedAt: &row.CreatedAt,
			UpdatedAt: utils.ValidTime(row.UpdatedAt),
			Name:      &row.Name,
			Age:       utils.PointerInt(int(row.Age)),
			Document:  &row.Document,
		})
	}
	return users, nil
}