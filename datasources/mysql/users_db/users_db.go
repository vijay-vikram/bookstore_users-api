package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysqlUsersUsername = "mysqlUsersUsername"
	mysqlUsersPassword = "mysqlUsersPassword"
	mysqlUsersSchema   = "mysqlUsersSchema"
)

var (
	Client   *sql.DB
	username = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	schema   = os.Getenv(mysqlUsersSchema)
)

func init() {
	var dataSource = fmt.Sprintf("%s:%s@/%s", username, password, schema)

	var err error
	Client, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	err = Client.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Database successfully configured")
}
