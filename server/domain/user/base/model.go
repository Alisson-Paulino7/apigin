package base

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID 			*uuid.UUID 
	CreatedAt 	*time.Time 
	UpdatedAt 	*time.Time
	DeletedAt 	*time.Time
	Name 		*string
	Age 		*int
	Document 	*string 
}