package mysqldb

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var connectedmyauth bool
var clientIDmyauth int64
var insertIDmyauth int64
var routeIDmyauth int64
var routeURLIDmyauth int64
var routeURLID2myauth int64

func TestConnectAuthDb(t *testing.T) {
	time.Sleep(time.Second * 10)
	connectedmyauth = ConnectAuthDb("localhost:3306", "admin", "admin", "ulbora_api_gateway")
	if connectedmyauth != true {
		fmt.Println("database init failed")
		t.Fail()
	} else {
		fmt.Println("database opened in mysqldb package")
	}
}

func TestConnectionAuthTest(t *testing.T) {
	rowPtr := ConnectionAuthTest()
	if rowPtr != nil {
		foundRow := rowPtr.Row
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		fmt.Print("Records found during test ")
		fmt.Println(int64Val)
		if err2 != nil {
			fmt.Print(err2)
		}
		if int64Val >= 0 {
			fmt.Println("database connection ok")
		} else {
			fmt.Println("database connection failed")
			t.Fail()
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
}
