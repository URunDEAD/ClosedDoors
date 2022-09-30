package main

import (
	"crypto/sha256"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var (
	database *sql.DB
	router   *mux.Router
)

func init() {
	log.Println("Starting ClosedDoors service...")

	host := flag.String("host", "127.0.0.1", "Specifys the host used by the SQL database.")
	user := flag.String("user", "root", "Specifys the user of the SQL database.")
	port := flag.Int("port", 3306, "Specifys the port used by the SQL database.")
	passwd := flag.String("passwd", "", "Specify the password of the SQL database.")
	dbName := flag.String("db-name", "doors", "Specifys the name of a database to be used on image startup.")

	flag.Parse()

	initDatabse(*host, *user, *passwd, *dbName, *port)

}

func main() {

	// router := mux.NewRouter()
	// print(router)
}

func initRouter() {

	router = mux.NewRouter()

	// router.HandleFunc("/check/")
}

func initDatabse(host, user, passwd, dbName string, port int) *sql.DB {

	log.Println("Connecting to SQL database...")

	var (
		err error
	)

	connectionString := user + ":" + passwd + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + dbName

	for {

		database, err = sql.Open("mysql", connectionString)
		if err != nil {
			panic(err)
		}

		err = database.Ping()
		if err == nil {
			log.Println("Connected!")
			break
		}

		time.Sleep(1 * time.Second)
		continue

	}
	defer database.Close()

	qry := "CREATE TABLE IF NOT EXISTS doors (" +
		"key_sha CHAR(64) PRIMARY KEY," +
		"expire_time TIMESTAMP" +
		")"

	_, err = database.Exec(qry)

	if err != nil {
		panic(err.Error())
	}

	return database
}

func registerKey(key_sha string, database *sql.DB) {
	qry := "INSERT INTO doors (key_sha, expire_time) VALUES('" + key_sha + "', CURRENT_TIMESTAMP)"

	_, err := database.Exec(qry)

	if err != nil {
		panic(err.Error())
	}

	log.Println("Added key hash: " + key_sha + " ")
}

func checkKey(key string, database *sql.DB) bool {
	key_sha := sha256.Sum256([]byte(key))
	qry := "SELECT key_sha FROM doors WHERE key_sha='" + fmt.Sprintf("%x", key_sha) + "'"

	rows, err := database.Query(qry)

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	if rows.Next() {
		log.Println("key: " + key + " status: VALID")
		return true
	}

	log.Println("key: " + key + " status: INVALID ")
	return false

}
