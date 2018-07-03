package managers

import (
	services "ApiGatewayAdminPortal/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"testing"
)

var o2u1 = "bobbyb1"

func TestGatewayUserService_getToken(t *testing.T) {
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

func TestGatewayUserService_AddOauth2User(t *testing.T) {
	var g GatewayAccountService
	g.Token = testToken
	g.UserHost = "http://localhost:3001"
	g.ClientID = "403"

	var cc GatewayAccount
	cc.Email = "testEmail"
	cc.FirstName = "bob1"
	cc.LastName = "bob1"
	cc.Name = "bob1 bob1"
	cc.WebSite = "www.test1.com"
	cc.Username = o2u1
	cc.ClientID = "403"
	res, pw := g.AddOauth2User(&cc)
	if !res.Success {
		fmt.Print("password: ")
		fmt.Println(pw)
		t.Fail()
	}
}

func TestGatewayUserService_DeleteOauth2User(t *testing.T) {
	var u services.UserService
	u.Token = testToken
	u.Host = "http://localhost:3001"
	u.ClientID = "403"
	res := u.DeleteUser(o2u1, u.ClientID)
	if !res.Success {
		t.Fail()
	}
}
