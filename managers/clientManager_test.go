package managers

import (
	services "ApiGatewayAdminPortal/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"testing"
)

//var addO2c int64
var addO2cmang int64

func TestGatewayClientService_getToken(t *testing.T) {
	if testToken == "" {
		req, _ := http.NewRequest("POST", tokenURL, nil)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("Client Add err: ")
			fmt.Println(cErr)
		} else {
			defer resp.Body.Close()
			var tres TokenResponse
			decoder := json.NewDecoder(resp.Body)
			error := decoder.Decode(&tres)
			if error != nil {
				log.Println(error.Error())
			} else {
				testToken = tres.Token
				//fmt.Print("token: ")
				//fmt.Println(testToken)
			}
		}
	}
}

func TestGatewayClientService_AddOauth2Client(t *testing.T) {
	var g GatewayAccountService
	g.Token = testToken
	g.Host = "http://localhost:3000"
	g.ClientID = "403"

	var cc GatewayAccount
	cc.Email = "testEmail"
	cc.FirstName = "bob"
	cc.LastName = "bob"
	cc.Name = "bob bob"
	cc.WebSite = "www.test.com"
	res := g.AddOauth2Client(&cc)
	if !res.Success || res.ClientID == 0 {
		t.Fail()
	} else {
		addO2cmang = res.ClientID
	}
}

func TestGatewayClientService_DeleteOauth2Client(t *testing.T) {
	var c services.ClientService
	c.Token = testToken
	c.Host = "http://localhost:3000"
	c.ClientID = "403"
	res := c.DeleteClient(strconv.FormatInt(addO2cmang, 10))
	if !res.Success {
		t.Fail()
	}
}
