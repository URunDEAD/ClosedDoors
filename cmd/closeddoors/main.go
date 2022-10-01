package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var (
	database *sql.DB
	router   *mux.Router
)

type JsonResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func init() {
	log.Println("Starting ClosedDoors service...")

	host := flag.String("host", "127.0.0.1", "Specifys the host used by the SQL database.")
	user := flag.String("user", "root", "Specifys the user of the SQL database.")
	port := flag.Int("port", 3306, "Specifys the port used by the SQL database.")
	passwd := flag.String("passwd", "", "Specify the password of the SQL database.")
	dbName := flag.String("db-name", "doors", "Specifys the name of a database to be used on image startup.")

	flag.Parse()

	initDatabse(*host, *user, *passwd, *dbName, *port)
	initRouter()

}

func main() {

	log.Fatal(http.ListenAndServe(":8080", router))
	database.Close()
}

func initRouter() {

	log.Println("Starting API...")

	router = mux.NewRouter()

	router.HandleFunc("/check", CheckKeyHandler).Methods("POST")

	router.HandleFunc("/register", RegisterKeyHandler).Methods("POST")
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

func RegisterKeyHandler(w http.ResponseWriter, r *http.Request) {

	key_sha := r.FormValue("hash")

	RegisterKey(key_sha)

	response := JsonResponse{Type: "success", Message: "Registered"}
	json.NewEncoder(w).Encode(response)

	log.Println("Added key hash: " + key_sha + " ")
}

func RegisterKey(key_sha string) {
	qry := "INSERT INTO doors (key_sha, expire_time) VALUES('" + key_sha + "', CURRENT_TIMESTAMP)"

	_, err := database.Exec(qry)

	if err != nil {
		panic(err.Error())
	}
}

func CheckKeyHandler(w http.ResponseWriter, r *http.Request) {

	key_sha := r.FormValue("hash")

	response := JsonResponse{Type: "success", Message: "Valid"}

	if !CheckKey(key_sha) {
		response.Message = "Invalid"
	}

	json.NewEncoder(w).Encode(response)

}

func CheckKey(key_sha string) bool {

	qry := "SELECT key_sha FROM doors WHERE key_sha='" + key_sha + "'"

	rows, err := database.Query(qry)

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	if rows.Next() {
		log.Println("key: " + key_sha + " status: VALID")

		return true
	}

	log.Println("key: " + key_sha + " status: INVALID ")

	return false

}
