package test

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/URunDEAD/ClosedDoors/pkg/cmd/closeddoors"
)

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

func TestCheckRegisteredKey(t *testing.T) {
	type addTest struct {
		inputSha string
		expected bool
	}
	// test matrix
	var addTests = []addTest{
		{"497b94cde9b006bc414f19af515a5462266704316e6d44d5cadaf4194cdcf5fa", true},
		{"37273a264f262f861c3cd0fbcbd67a8878090e63dad4d37d4cad17b7ce0741a6", true},
		{"d0c3b98207936c0a0fae6757695814fdc2a6a3457469e7b669d00d68687a5c26", true},
		{"76ba2389a80bb0ee15d8107f676ea464d777647ed810051dfae38d9abd7350ee", true},
		{"ecbe9eb7cec2f0d826cb1dba9d0f7a1a691a898a4ccce4d87d604d09094a3150", true},
	}

	expire_time := time.Now().Add(time.Hour * 1).Format("2006-01-02 15:04:05")
	print(expire_time)

	//populate test db
	for _, test := range addTests {
		closeddoors.RegisterKey(test.inputSha, expire_time)
	}

	//start testing
	for _, test := range addTests {
		if closeddoors.CheckKey(test.inputSha) != test.expected {
			t.Errorf("Key %s does no show as found but should have been", test.inputSha)
		}
	}
}
