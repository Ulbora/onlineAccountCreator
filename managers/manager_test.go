package managers

import (
	services "ApiGatewayAdminPortal/services"
	ulboraUris "ApiGatewayAdminPortal/ulborauris"
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
	g.UserHost = "http://localhost:3001"
	g.ClientID = "403"
	var cc GatewayAccount
	cc.Email = "testEmail"
	cc.FirstName = "bob1"
	cc.LastName = "bob1"
	cc.Name = "bob1 bob1"
	cc.WebSite = "www.test1.com"
	cc.Username = "bobbbbb1"

	var sel ulboraUris.UlboraSelection
	sel.Oauth2 = true
	sel.APIGateway = true
	cc.UlboraSelected = &sel
	res := g.AddGatewayAccount(&cc)
	if !res.Success || res.ClientID == 0 {
		t.Fail()
	} else {
		addO2clnt = res.ClientID
	}
}

func TestGatewayAccountService_DeleteOauth2User(t *testing.T) {
	var u services.UserService
	u.Token = testToken
	u.Host = "http://localhost:3001"
	u.ClientID = "403"
	res := u.DeleteUser("bobbbbb1", strconv.FormatInt(addO2clnt, 10))
	if !res.Success {
		t.Fail()
	}
}

func TestGatewayAccountService_DeleteClientGrantTypes(t *testing.T) {
	var gt services.GrantTypeService
	gt.Token = testToken
	gt.Host = "http://localhost:3000"
	gt.ClientID = "403"
	gtypes := gt.GetGrantTypeList(strconv.FormatInt(addO2clnt, 10))
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

func TestGatewayAccountService_DeleteRoleUrl(t *testing.T) {
	var cr services.RoleURIService
	cr.ClientID = "403"
	cr.Host = "http://localhost:3000"
	cr.Token = testToken

	var rs services.ClientRoleService
	rs.Token = testToken
	rs.Host = "http://localhost:3000"
	rs.ClientID = "403"
	roles := rs.GetClientRoleList(strconv.FormatInt(addO2clnt, 10))
	var failed bool
	for _, r := range *roles {
		rurls := cr.GetRoleURIList(strconv.FormatInt(r.ID, 10))
		fmt.Print("role uris found for id: ")
		fmt.Println(rurls)
		fmt.Println(r)
		for _, ruri := range *rurls {
			res := cr.DeleteRoleURI(&ruri)
			if !res.Success {
				failed = true
			}
		}
	}

	// rurls := cr.GetRoleURIList(strconv.FormatInt(addO2curl, 10))
	// fmt.Print("role uris found: ")
	// fmt.Println(rurls)

	if failed {
		t.Fail()
	}
}

func TestGatewayAccountService_DeleteClientRoles(t *testing.T) {
	var rs services.ClientRoleService
	rs.Token = testToken
	rs.Host = "http://localhost:3000"
	rs.ClientID = "403"
	roles := rs.GetClientRoleList(strconv.FormatInt(addO2clnt, 10))
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

func TestGatewayAccountService_DeleteClientAllowedUris(t *testing.T) {

	var au services.AllowedURIService
	au.ClientID = "403"
	au.Host = "http://localhost:3000"
	au.Token = testToken
	uris := au.GetAllowedURIList(strconv.FormatInt(addO2clnt, 10))
	fmt.Print("uris: ")
	fmt.Println(uris)
	var failed bool
	for _, ur := range *uris {
		res := au.DeleteAllowedURI(strconv.FormatInt(ur.ID, 10))
		if !res.Success {
			failed = true
		}
	}
	if failed {
		t.Fail()
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
