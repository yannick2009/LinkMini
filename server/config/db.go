package config

import (
	"errors"
	"os"

	"github.com/surrealdb/surrealdb.go"
)

const DBConnectLink = "ws://localhost:8000/rpc"
const dbUser = "DB_USER"
const dbPassword = "DB_PASSWORD"
const test = "test"

var (
	ErrDBconnection = errors.New("DB connection failed")
	ErrDBTest       = errors.New("DB test failed")
	ErrDBAuth       = errors.New("DB authentication failed")
)

func ConnectDB() (*surrealdb.DB, error) {
	db, err := surrealdb.New(DBConnectLink)
	if err != nil {
		return nil, errors.Join(err, ErrDBconnection)
	}

	if _, err = db.Signin(map[string]interface{}{
		"user": os.Getenv(dbUser),
		"pass": os.Getenv(dbPassword),
	}); err != nil {
		return nil, errors.Join(err, ErrDBAuth)
	}

	if _, err = db.Use(test, test); err != nil {
		return nil, errors.Join(err, ErrDBTest)
	}

	return db, nil
}
