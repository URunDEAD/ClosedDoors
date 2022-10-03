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
	closeddoors.InitDatabase(host, user, passwd, dbName, port)
}

func TestCheckRegisterKey(t *testing.T) {

	// test matrix
	var addTests = []addTest{
		{"497b94cde9b006bc414f19af515a5462266704316e6d44d5cadaf4194cdcf5fa", false},
		{"37273a264f262f861c3cd0fbcbd67a8878090e63dad4d37d4cad17b7ce0741a6", false},
		{"d0c3b98207936c0a0fae6757695814fdc2a6a3457469e7b669d00d68687a5c26", false},
		{"76ba2389a80bb0ee15d8107f676ea464d777647ed810051dfae38d9abd7350ee", true},
		{"ecbe9eb7cec2f0d826cb1dba9d0f7a1a691a898a4ccce4d87d604d09094a3150", true},
	}

	//populate test db
	for _, test := range addTests {
		if test.expected == true {
			closeddoors.RegisterKey(test.input)
		}
	}

	//start testing
	for _, test := range addTests {
		if closeddoors.CheckKey(test.input) != test.expected {
			if test.expected == false {
				t.Errorf("Key %s shows as found but should not have been", test.input)
			} else {
				t.Errorf("Key %s shows as not found but should have been", test.input)
			}
		}
	}
}
