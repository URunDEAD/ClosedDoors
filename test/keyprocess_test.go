package test

import (
	"os"
	"strconv"
	"testing"

	"github.com/URunDEAD/ClosedDoors/pkg/cmd/closeddoors"
)

type addTest struct {
	input    string
	expected bool
}

var addTests = []addTest{
	addTest{"497b94cde9b006bc414f19af515a5462266704316e6d44d5cadaf4194cdcf5fa", false},
}

var (
	host, user, passwd, dbName string
	port                       int
)

func init() {
	host = os.Getenv("MYSQL_HOST")
	user = os.Getenv("MYSQL_USER")
	port, _ = strconv.Atoi(os.Getenv("MYSQL_PORT"))
	passwd = os.Getenv("MYSQL_PASSWD")
	dbName = os.Getenv("MYSQL_DBNAME")

	// flag.Parse()

	closeddoors.InitDatabase(host, user, passwd, dbName, port)
}

func TestProcessKey(t *testing.T) {
	for _, test := range addTests {
		if closeddoors.CheckKey(test.input) != test.expected {
			t.Errorf("Key %s shows as found but should not have been", test.input)
		}
	}
}
