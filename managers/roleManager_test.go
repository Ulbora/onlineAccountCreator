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
var addO2crole int64

func TestGatewayClientRoleService_getToken(t *testing.T) {
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

func TestGatewayClientRoleService_AddOauth2Client(t *testing.T) {
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
		addO2crole = res.ClientID
	}
}

func TestGatewayClientRoleService_AddClientRoles(t *testing.T) {
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
	cc.ClientID = strconv.FormatInt(addO2crole, 10)
	res := g.AddClientRole(&cc)
	fmt.Print("role res: ")
	fmt.Println(res)
	if !res.Success {
		t.Fail()
	}
}

func TestGatewayClientRoleService_DeleteClientRoles(t *testing.T) {
	var rs services.ClientRoleService
	rs.Token = testToken
	rs.Host = "http://localhost:3000"
	rs.ClientID = "403"
	roles := rs.GetClientRoleList(strconv.FormatInt(addO2crole, 10))
	fmt.Print("roles: ")
	fmt.Println(roles)
	var failed bool
	for _, r := range *roles {
		res := rs.DeleteClientRole(strconv.FormatInt(r.ID, 10))
		if !res.Success {
			failed = true
		}
	}
	if failed {
		t.Fail()
	}
}

func TestGatewayClientRoleService_DeleteOauth2Client(t *testing.T) {
	var c services.ClientService
	c.Token = testToken
	c.Host = "http://localhost:3000"
	c.ClientID = "403"
	res := c.DeleteClient(strconv.FormatInt(addO2crole, 10))
	if !res.Success {
		t.Fail()
	}
}
