package main

import (
	"flag"

	"github.com/URunDEAD/ClosedDoors/pkg/cmd/closeddoors"
)

func main() {

	host := flag.String("MYSQL-Host", "127.0.0.1", "Specifys the host used by the SQL database.")
	user := flag.String("MYSQL-User", "root", "Specifys the user of the SQL database.")
	port := flag.Int("MYSQL-Port", 3306, "Specifys the port used by the SQL database.")
	passwd := flag.String("MYSQL-Passwd", "", "Specify the password of the SQL database.")
	dbName := flag.String("MYSQL-DB-name", "doors", "Specifys the name of a database to be used on image startup.")

	flag.Parse()

	closeddoors.StartServer(*host, *user, *passwd, *dbName, *port)
}
