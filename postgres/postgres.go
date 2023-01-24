package postgres

import (
	"bank/model"
	"database/sql"
)

type PG struct {
	db *sql.DB
}

func New(db *sql.DB) *PG {
	return &PG{
		db: db,
	}
}
func (pg PG) GetUser(name string) (*model.User, error) {
	var user model.User
	err := pg.db.QueryRow("SELECT name,age from data where name =$1", name).Scan(&user.Name, &user.Age)
	return &user, err
}
func (pg PG) InsertUser(user *model.User) error {
	_, err := pg.db.Exec("INSERT  into data (name,age) values ($1,$2)", user.Name, user.Age)
	return err
}
