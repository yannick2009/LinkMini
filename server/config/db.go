package config

import (
	"errors"
	"log"
	"os"

	"github.com/surrealdb/surrealdb.go"
)

const dbConnectLink = "ws://localhost:8000/rpc"
const dbUser = "DB_USER"
const dbPassword = "DB_PASSWORD"
const namespace = "linkmini"
const database = "linkmini"

var (
	ErrDBconnection = errors.New("DB connection failed")
	ErrDBAuth       = errors.New("DB authentication failed")
)

var DB *surrealdb.DB

func ConnectDB() error {
	db, err := surrealdb.New(dbConnectLink)
	if err != nil {
		return errors.Join(err, ErrDBconnection)
	}

	if _, err = db.Signin(map[string]interface{}{
		"user": os.Getenv(dbUser),
		"pass": os.Getenv(dbPassword),
	}); err != nil {
		log.Print("fail to auth to SurrealDB ❌")
		return errors.Join(err, ErrDBAuth)
	}

	if _, err = db.Use(namespace, database); err != nil {
		log.Print("fail to access to the database ❌")
		panic(err)
	}

	log.Print("Connected to DB ✅")
	DB = db
	return nil
}
