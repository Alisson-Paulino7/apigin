// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package base

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	Delete(ctx context.Context, id uuid.UUID) error
	FindAll(ctx context.Context) ([]*UserTUser, error)
	FindByID(ctx context.Context, id uuid.UUID) (*UserTUser, error)
	Insert(ctx context.Context, arg *InsertParams) (*InsertRow, error)
	Update(ctx context.Context, arg *UpdateParams) error
}

var _ Querier = (*Queries)(nil)
