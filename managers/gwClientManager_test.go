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
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"testing"
)

//var GwCid int64
var GwCidStr string

func TestGwAccountService_getToken(t *testing.T) {
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

func TestGwAccountService_AddGwClient(t *testing.T) {
	var g GatewayAccountService
	g.Token = testToken
	g.GwHost = "http://localhost:3011"
	g.ClientID = "403"
	g.APIKey = "403"

	var cc GatewayAccount
	cc.Email = "testEmail"
	cc.FirstName = "bob"
	cc.LastName = "bob"
	cc.Name = "bob bob"
	cc.WebSite = "www.test.com"

	//GwCid = 55445588844444444
	GwCidStr = "55445588844444444"

	cc.ClientID = GwCidStr
	res := g.AddGwClient(&cc)

	fmt.Print("res in gw client add: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGwAccountService_DeleteGwClient(t *testing.T) {
	var c services.GatewayClientService
	c.ClientID = "403"
	c.Host = "http://localhost:3011"
	c.Token = testToken

	res := c.DeleteClient(GwCidStr)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}
