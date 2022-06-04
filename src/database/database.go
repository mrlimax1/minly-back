package database

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"log"
	"os"
)

func Connect() *pg.DB {
	opt, err := pg.ParseURL(fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("db_login"),
		os.Getenv("db_password"),
		"127.0.0.1",
		os.Getenv("db_port"),
		os.Getenv("db_db")))
	if err != nil {
		log.Fatal(err)
	}
	return pg.Connect(opt)

}

func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		(*Sites)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}
