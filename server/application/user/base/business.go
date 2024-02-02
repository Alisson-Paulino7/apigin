package base

import (
	"context"

	"github.com/alisson-paulino7/apigin/database"
	"github.com/alisson-paulino7/apigin/domain/user/base"
	"github.com/google/uuid"
)

// Context é como o significado momentaneo
// Cada requisição tem um contexto específico

// Find all Users
func FindAll(ctx context.Context) (out *FindAllResponse, err error) {
	
	return nil, nil
}

// Find an User by ID
func FindByID(ctx context.Context, in *uuid.UUID) (out *FindByIDResponse, err error) {
	out = &FindByIDResponse{
		Name:     out.Name,
		Age:      out.Age,
		Document: out.Document,
	}

	return out, nil
}

// Create an User
func Create(ctx context.Context, in *CreateRequest) (out *CreateResponse, err error) {

	tx, err := database.NewTransation(ctx)
	if err != nil {
		return nil, err
	}

	repository := base.NewRepository(ctx, tx)

	usr := &base.User{
		Name: 		in.Name,
		Age: 		in.Age,
		Document: 	in.Document,
	}
	if err := repository.Insert(usr); err != nil {
		return nil, err
	} 
	defer tx.Rollback()	

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &CreateResponse{
		ID: 	  	usr.ID,
		CreatedAt: 	usr.CreatedAt,
		Name:     	usr.Name,
		Age:      	usr.Age,
		Document: 	usr.Document,
	},nil

}

// Update an User
func Update(ctx context.Context, in *UpdateRequest) error {
	return nil
}

// Delete an User
func Delete(ctx context.Context, in *uuid.UUID) error {
	return nil
}
