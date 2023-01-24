package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func NewConnection(username, password, database string) (*sql.DB, error) {
	connection := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, database)
	log.Println(connection)
	return sql.Open("postgres", connection)
}
