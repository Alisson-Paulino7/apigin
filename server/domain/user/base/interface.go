package base

type BaseInterface interface {
	Insert(user *User) error
	// Update(user *User) error
	// Delete(id string) error
	// FindAll() ([]*User, error)
	// FindByID(id string) (*User, error)
}