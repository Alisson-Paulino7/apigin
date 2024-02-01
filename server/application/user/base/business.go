package base

import (
	"context"

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

	out = &CreateResponse{
		Name:     in.Name,
		Age:      in.Age,
		Document: in.Document,
	}

	return out, nil

}

// Update an User
func Update(ctx context.Context, in *UpdateRequest) error {
	return nil
}

// Delete an User
func Delete(ctx context.Context, in *uuid.UUID) error {
	return nil
}
