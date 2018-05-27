package mysqldb

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var connectedmygw bool
var clientIDmygw int64
var insertIDmygw int64
var routeIDmygw int64
var routeURLIDmygw int64
var routeURLID2mygw int64

func TestConnectGwDb(t *testing.T) {
	time.Sleep(time.Second * 10)
	connectedmygw = ConnectGwDb("localhost:3306", "admin", "admin", "ulbora_api_gateway")
	if connectedmygw != true {
		fmt.Println("database init failed")
		t.Fail()
	} else {
		fmt.Println("database opened in mysqldb package")
	}
}

func TestConnectionGwTest(t *testing.T) {
	rowPtr := ConnectionGwTest()
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
