package usersdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersPort     = "mysql_users_port"
	mysqlUsersSchema   = "mysql_users_schema"
)

var (
	// Client is the connection to our users databasae
	Client *sql.DB

	username string
	password string
	host     string
	port     string
	schema   string
)

func init() {
	// FIXME: move to a more globally file
	godotenv.Load()

	username = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	host = os.Getenv(mysqlUsersHost)
	port = os.Getenv(mysqlUsersPort)
	schema = os.Getenv(mysqlUsersSchema)

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		username,
		password,
		host,
		port,
		schema,
	)
	var err error

	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	// mysql.SetLogger()
	log.Println("database successfully configured")
}
