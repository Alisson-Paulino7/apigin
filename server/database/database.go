package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var instance *sql.DB
  
const (
	host     = "localhost"
	port     = 5431
	user     = "postgres"
	password = "dev123456"
	dbname   = "manager"
)


func Init() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err = db.Ping();err != nil {
		panic(err)
	}

	instance = db

	return nil
}

func NewTransation(ctx context.Context) (*sql.Tx, error) {
	return instance.BeginTx(ctx, &sql.TxOptions{})
}


// func Get() error {
// 	return nil
// }