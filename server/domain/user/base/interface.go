package base

import "github.com/google/uuid"

type BaseInterface interface {
	Insert(user *User) error
	Update(user *User) error
	Delete(id uuid.UUID) error
	FindAll() ([]*User, error)
	FindByID(id uuid.UUID) (*User, error)
}