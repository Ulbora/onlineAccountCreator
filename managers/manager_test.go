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

var addO2clnt int64

func TestGatewayAccountService_getToken(t *testing.T) {
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

func TestGatewayAccountService_AddGatewayAccount(t *testing.T) {
	var g GatewayAccountService
	g.Token = testToken
	g.Host = "http://localhost:3000"
	g.ClientID = "403"
	var cc GatewayAccount
	cc.Email = "testEmail"
	cc.FirstName = "bob1"
	cc.LastName = "bob1"
	cc.Name = "bob1 bob1"
	cc.WebSite = "www.test1.com"
	res := g.AddGatewayAccount(&cc)
	if !res.Success || res.ClientID == 0 {
		t.Fail()
	} else {
		addO2clnt = res.ClientID
	}
}

func TestGatewayAccountService_DeleteOauthClient(t *testing.T) {
	var c services.ClientService
	c.Token = testToken
	c.Host = "http://localhost:3000"
	c.ClientID = "403"
	res := c.DeleteClient(strconv.FormatInt(addO2clnt, 10))
	if !res.Success {
		t.Fail()
	}
}

func TestGatewayAccountService_generateTempPassword(t *testing.T) {
	pw := generateTempPassword()
	fmt.Print("password: ")
	fmt.Println(pw)
	if len(pw) != 20 {
		t.Fail()
	}
}
