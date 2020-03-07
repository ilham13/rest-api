package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// var DB *sql.DB

// func Init() {
// 	var err error
// 	DB, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/shopee")
// 	CheckErr(err)

// 	fmt.Println("You connected to your database.")
// }

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/shopee")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
