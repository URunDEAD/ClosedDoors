package database

import (
	"database/sql"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	database                   *sql.DB
	user, passwd, dbName, host string
	port                       int
}

func NewSQLConnection() *Database {
	var (
		db Database
	)

	return &db
}

func (x *Database) SetHost(value string) *Database {
	x.host = value
	return x
}

func (x *Database) SetUser(value string) *Database {
	x.user = value
	return x
}

func (x *Database) SetPasswd(value string) *Database {
	x.passwd = value
	return x
}

func (x *Database) SetDBName(value string) *Database {
	x.dbName = value
	return x
}

func (x *Database) SetPort(value int) *Database {
	x.port = value
	return x
}

func (x *Database) CheckKey(key_sha string) bool {
	qry := "SELECT key_sha FROM doors WHERE key_sha='" + key_sha + "' AND expire_time>=" + "CURRENT_TIMESTAMP" + ""

	rows, err := x.database.Query(qry)

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	return rows.Next()

}

func (x *Database) RegisterKey(key_sha, expireTime string) {
	qry := "INSERT INTO doors (key_sha, expire_time) VALUES('" + key_sha + "', '" + expireTime + "')"

	_, err := x.database.Exec(qry)

	if err != nil {
		panic(err.Error())
	}
}

func (x *Database) InitDatabase() {
	var (
		err error
	)

	qry := "CREATE TABLE IF NOT EXISTS doors (" +
		"key_sha CHAR(64) PRIMARY KEY," +
		"expire_time TIMESTAMP" +
		")"

	_, err = x.database.Exec(qry)

	if err != nil {
		panic(err.Error())
	}

}

func (x *Database) StartConnection() {
	var (
		err error
	)
	connectionString := x.user + ":" + x.passwd + "@tcp(" + x.host + ":" + strconv.Itoa(x.port) + ")/" + x.dbName

	for {

		x.database, err = sql.Open("mysql", connectionString)
		if err != nil {
			panic(err)
		}

		err = x.database.Ping()
		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
		continue

	}
}
