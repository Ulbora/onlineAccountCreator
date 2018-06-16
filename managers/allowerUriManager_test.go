/*
 Copyright (C) 2017 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2017 Ken Williamson
 All rights reserved.

 Certain inventions and disclosures in this file may be claimed within
 patents owned or patent applications filed by Ulbora Labs LLC., or third
 parties.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as published
 by the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

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

var addO2curl int64

func TestGatewayClientUrlService_getToken(t *testing.T) {
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

func TestGatewayClientUrlService_AddOauth2Client(t *testing.T) {
	var g GatewayAccountService
	g.Token = testToken
	g.Host = "http://localhost:3000"
	g.ClientID = "403"

	var cc GatewayAccount
	cc.Email = "testEmail"
	cc.FirstName = "boburl"
	cc.LastName = "boburl"
	cc.Name = "boburl boburl"
	cc.WebSite = "www.test.com"
	res := g.AddOauth2Client(&cc)
	if !res.Success || res.ClientID == 0 {
		t.Fail()
	} else {
		addO2curl = res.ClientID
	}
}

func TestGatewayClientUrlService_AddClientRoles(t *testing.T) {
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
	cc.ClientID = strconv.FormatInt(addO2curl, 10)
	res := g.AddClientRole(&cc)
	fmt.Print("role res: ")
	fmt.Println(res)
	if !res.Success {
		t.Fail()
	}
}

func TestGatewayClientUrlService_insertAllowedURI(t *testing.T) {
	var au services.AllowedURIService
	au.ClientID = "403"
	au.Host = "http://localhost:3000"
	au.Token = testToken

	suc, _ := insertAllowedURI("0", "ggg", au)
	fmt.Println(suc)
	if !suc {
		t.Fail()
	}
}

func TestGatewayClientUrlService_insertRoleURI(t *testing.T) {
	var g GatewayAccountService
	g.Token = testToken
	g.Host = "http://localhost:3000"
	g.ClientID = "403"
	suc := g.insertRoleURI(false, 0, 0)
	if suc {
		t.Fail()
	}
}

func TestGatewayClientUrlService_insertRoleURI2(t *testing.T) {
	var g GatewayAccountService
	g.Token = testToken
	g.Host = "http://localhost:3000"
	g.ClientID = "403"
	suc := g.insertRoleURI(true, 0, 0)
	if suc {
		t.Fail()
	}
}

func TestGatewayClientUrlService_AddAllowedUris(t *testing.T) {
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
	cc.ClientID = strconv.FormatInt(addO2curl, 10)
	var sel ulboraUris.UlboraSelection
	sel.Oauth2 = true
	sel.APIGateway = true
	cc.UlboraSelected = &sel
	//---------------modify GatewayAccount to include UlboraSelection attr--------------------------------------

	res := g.AddAllowedUris(&cc)
	fmt.Print("role uri res: ")
	fmt.Println(res)
	if !res.Success {
		t.Fail()
	}
}

func TestGatewayClientUrlService_DeleteRoleUrl(t *testing.T) {
	var cr services.RoleURIService
	cr.ClientID = "403"
	cr.Host = "http://localhost:3000"
	cr.Token = testToken

	var rs services.ClientRoleService
	rs.Token = testToken
	rs.Host = "http://localhost:3000"
	rs.ClientID = "403"
	roles := rs.GetClientRoleList(strconv.FormatInt(addO2curl, 10))
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

func TestGatewayClientUrlService_DeleteClientRoles(t *testing.T) {
	var rs services.ClientRoleService
	rs.Token = testToken
	rs.Host = "http://localhost:3000"
	rs.ClientID = "403"
	roles := rs.GetClientRoleList(strconv.FormatInt(addO2curl, 10))
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

func TestGatewayClientUrlService_DeleteClientAllowedUris(t *testing.T) {

	var au services.AllowedURIService
	au.ClientID = "403"
	au.Host = "http://localhost:3000"
	au.Token = testToken
	uris := au.GetAllowedURIList(strconv.FormatInt(addO2curl, 10))
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

func TestGatewayClientUrlService_DeleteOauth2Client(t *testing.T) {
	var c services.ClientService
	c.Token = testToken
	c.Host = "http://localhost:3000"
	c.ClientID = "403"
	res := c.DeleteClient(strconv.FormatInt(addO2curl, 10))
	if !res.Success {
		t.Fail()
	}
}
