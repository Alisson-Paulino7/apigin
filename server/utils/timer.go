package utils

import (
	"database/sql"
	"time"
)
func ValidTime(r sql.NullTime) *time.Time {
	if !r.Valid {
		return nil
	}
	return &r.Time
}