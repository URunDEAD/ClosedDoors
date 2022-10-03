package closeddoors

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/URunDEAD/ClosedDoors/pkg/cmd/database"
	"github.com/gorilla/mux"
)

var (
	router *mux.Router
	db     *database.Database
)

func init() {
	log.Println("Starting\n",
		"   _____ _                    _ _____                             \n",
		"  / ____| |                  | |  __ \\                           \n",
		" | |    | | ___  ___  ___  __| | |  | | ___   ___  _ __ ___       \n",
		" | |    | |/ _ \\/ __|/ _ \\/ _` | |  | |/ _ \\ / _ \\| '__/ __|  \n",
		" | |____| | (_) \\__ \\  __/ (_| | |__| | (_) | (_) | |  \\__ \\  \n",
		"  \\_____|_|\\___/|___/\\___|\\__,_|_____/ \\___/ \\___/|_|  |___/")

}

type JsonResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func StartServer(mysqlHost, mysqlUser, mysqlPasswd, mysqlDbName string, mysqlPort int) {

	log.Println("Connecting to SQL database...")
	db = database.NewSQLConnection().
		SetHost(mysqlHost).
		SetUser(mysqlUser).
		SetPasswd(mysqlPasswd).
		SetDBName(mysqlDbName).
		SetPort(mysqlPort)
	db.StartConnection()
	log.Println("Connection Established!")

	log.Println("Creating doors table...")
	db.InitDatabase()
	log.Println("Done!")

	log.Println("Initiating API...")
	InitRouter()
	log.Println("Done!")
	log.Println("Waiting for requests...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func InitRouter() {
	log.Println("Starting API...")

	router = mux.NewRouter()

	router.HandleFunc("/check", CheckKeyHandler).Methods("POST")

	router.HandleFunc("/register", RegisterKeyHandler).Methods("POST")
}

func RegisterKeyHandler(w http.ResponseWriter, r *http.Request) {
	key_sha := r.FormValue("hash")
	expireTime := r.FormValue("expire-time")

	db.RegisterKey(key_sha, expireTime)

	response := JsonResponse{Type: "success", Message: "Registered"}
	json.NewEncoder(w).Encode(response)

	log.Println("Added key hash: " + key_sha + " ")
}

func CheckKeyHandler(w http.ResponseWriter, r *http.Request) {
	key_sha := r.FormValue("hash")

	response := JsonResponse{Type: "success", Message: "Valid"}

	if !db.CheckKey(key_sha) {
		response.Message = "Invalid"
	}
	log.Println("key: " + key_sha + " status: " + response.Message)
	json.NewEncoder(w).Encode(response)

}
