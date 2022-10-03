package closeddoors

import (
	"database/sql"
	"encoding/json"
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

func StartServer(mysqlHost, mysqlUser, mysqlPasswd, mysqlDbName string, mysqlPort int) {
	log.Println("Starting ClosedDoors service...")
	log.Println("Connecting to SQL database...")
	InitDatabase(mysqlHost, mysqlUser, mysqlPasswd, mysqlDbName, mysqlPort)
	log.Println("Connected!")
	InitRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
	database.Close()
}
func InitRouter() {
	log.Println("Starting API...")

	router = mux.NewRouter()

	router.HandleFunc("/check", CheckKeyHandler).Methods("POST")

	router.HandleFunc("/register", RegisterKeyHandler).Methods("POST")
}

func InitDatabase(host, user, passwd, dbName string, port int) {
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

}

func RegisterKeyHandler(w http.ResponseWriter, r *http.Request) {
	key_sha := r.FormValue("hash")
	expireTime := r.FormValue("expire-time")

	RegisterKey(key_sha, expireTime)

	response := JsonResponse{Type: "success", Message: "Registered"}
	json.NewEncoder(w).Encode(response)

	log.Println("Added key hash: " + key_sha + " ")
}

func RegisterKey(key_sha, expireTime string) {
	qry := "INSERT INTO doors (key_sha, expire_time) VALUES('" + key_sha + "', '" + expireTime + "')"

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
	log.Println("key: " + key_sha + " status: " + response.Message)
	json.NewEncoder(w).Encode(response)

}

func CheckKey(key_sha string) bool {
	qry := "SELECT key_sha FROM doors WHERE key_sha='" + key_sha + "' AND expire_time>=" + "CURRENT_TIMESTAMP" + ""

	rows, err := database.Query(qry)

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	return rows.Next()

}
