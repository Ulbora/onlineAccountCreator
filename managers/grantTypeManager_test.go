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
var addO2cGrant int64

func TestGatewayClientGrantService_getToken(t *testing.T) {
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

func TestGatewayClientGrantService_AddOauth2Client(t *testing.T) {
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
		addO2cGrant = res.ClientID
	}
}

func TestGatewayClientGrantService_AddClientGrantTypes(t *testing.T) {
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
	cc.ClientID = strconv.FormatInt(addO2cGrant, 10)
	res := g.AddClientGrantType(&cc)
	fmt.Print("role res: ")
	fmt.Println(res)
	if !res.Success {
		t.Fail()
	}
}

func TestGatewayClientGrantService_DeleteClientGrantTypes(t *testing.T) {
	var gt services.GrantTypeService
	gt.Token = testToken
	gt.Host = "http://localhost:3000"
	gt.ClientID = "403"
	gtypes := gt.GetGrantTypeList(strconv.FormatInt(addO2cGrant, 10))
	fmt.Print("gtypes: ")
	fmt.Println(gtypes)
	var failed bool
	for _, r := range *gtypes {
		res := gt.DeleteGrantType(strconv.FormatInt(r.ID, 10))
		if !res.Success {
			failed = true
		}
	}
	if failed {
		t.Fail()
	}
}

func TestGatewayClientGrantService_DeleteOauth2Client(t *testing.T) {
	var c services.ClientService
	c.Token = testToken
	c.Host = "http://localhost:3000"
	c.ClientID = "403"
	res := c.DeleteClient(strconv.FormatInt(addO2cGrant, 10))
	if !res.Success {
		t.Fail()
	}
}
