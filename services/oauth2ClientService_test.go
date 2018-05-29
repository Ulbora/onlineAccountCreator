/*
 Copyright (C) 2017 Ulbora Labs Inc. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2017 Ken Williamson
 All rights reserved.

 Certain inventions and disclosures in this file may be claimed within
 patents owned or patent applications filed by Ulbora Labs Inc., or third
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

package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"testing"
)

var CIDac int64

//var tokenac string

func TestClientService_getToken(t *testing.T) {
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

func TestClientService_AddClient(t *testing.T) {
	var c ClientService
	c.ClientID = "403"
	c.Host = "http://localhost:3000"
	c.Token = testToken
	var uri RedirectURI
	uri.URI = "http://googole.com"
	var uris []RedirectURI
	uris = append(uris, uri)
	var cc Client
	cc.Email = "ken@ken.com"
	cc.Enabled = true
	cc.Name = "A Big Company"
	cc.RedirectURIs = uris
	res := c.AddClient(&cc)
	fmt.Print("res: ")
	fmt.Println(res)
	CIDac = res.ClientID
	if res.Success != true {
		t.Fail()
	}
}

func TestClientService_AddClientUrl(t *testing.T) {
	var c ClientService
	c.ClientID = "403"
	c.Host = "://localhost:3000"
	c.Token = testToken
	var uri RedirectURI
	uri.URI = "http://googole.com"
	var uris []RedirectURI
	uris = append(uris, uri)
	var cc Client
	cc.Email = "ken@ken.com"
	cc.Enabled = true
	cc.Name = "A Big Company"
	cc.RedirectURIs = uris
	res := c.AddClient(&cc)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success == true {
		t.Fail()
	}
}

func TestClientService_AddClientReq(t *testing.T) {
	var c ClientService
	c.ClientID = "403"
	c.Host = "http://localhost:30001"
	c.Token = testToken
	var uri RedirectURI
	uri.URI = "http://googole.com"
	var uris []RedirectURI
	uris = append(uris, uri)
	var cc Client
	cc.Email = "ken@ken.com"
	cc.Enabled = true
	cc.Name = "A Big Company"
	cc.RedirectURIs = uris
	res := c.AddClient(&cc)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success == true {
		t.Fail()
	}
}

// func TestClientService_UpdateClient(t *testing.T) {
// 	var c ClientService
// 	c.ClientID = "403"
// 	c.Host = "http://localhost:3000"
// 	c.Token = tempToken
// 	var cc Client
// 	cc.Email = "ken@ken1.com"
// 	cc.Enabled = true
// 	cc.Name = "A Really Big Company"
// 	cc.WebSite = "http://www.ulbora.com"
// 	cc.ClientID = CID
// 	res := c.UpdateClient(&cc)
// 	fmt.Print("res: ")
// 	fmt.Println(res)
// 	if res.Success != true {
// 		t.Fail()
// 	}
// }

func TestClientService_GetClient(t *testing.T) {
	var c ClientService
	c.ClientID = "403"
	c.Host = "http://localhost:3000"
	c.Token = testToken
	fmt.Print("CID: ")
	fmt.Println(CIDac)
	res := c.GetClient(strconv.FormatInt(CIDac, 10))
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Enabled != true {
		t.Fail()
	}
}

func TestClientService_GetClientUrl(t *testing.T) {
	var c ClientService
	c.ClientID = "403"
	c.Host = "://localhost:3000"
	c.Token = testToken
	fmt.Print("CID: ")
	fmt.Println(CIDac)
	res := c.GetClient(strconv.FormatInt(CIDac, 10))
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Enabled == true {
		t.Fail()
	}
}

func TestClientService_GetClientReq(t *testing.T) {
	var c ClientService
	c.ClientID = "403"
	c.Host = "http://localhost:30001"
	c.Token = testToken
	fmt.Print("CID: ")
	fmt.Println(CIDac)
	res := c.GetClient(strconv.FormatInt(CIDac, 10))
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Enabled == true {
		t.Fail()
	}
}

// func TestClientService_SearchClient(t *testing.T) {
// 	var c ClientService
// 	c.ClientID = "403"
// 	c.Host = "http://localhost:3000"
// 	c.Token = tempToken
// 	var cc Client
// 	cc.Name = "Big"
// 	res := c.SearchClient(&cc)
// 	fmt.Print("searched res: ")
// 	fmt.Println(res)
// 	if res == nil || len(*res) == 0 {
// 		t.Fail()
// 	}
// }

// func TestClientService_GetClientList(t *testing.T) {
// 	var c ClientService
// 	c.ClientID = "403"
// 	c.Host = "http://localhost:3000"
// 	c.Token = tempToken
// 	res := c.GetClientList()
// 	fmt.Print("res list: ")
// 	fmt.Println(res)
// 	fmt.Print("len: ")
// 	fmt.Println(len(*res))
// 	if res == nil || len(*res) == 0 {
// 		t.Fail()
// 	}
// }

func TestClientService_DeleteClient(t *testing.T) {
	var c ClientService
	c.ClientID = "403"
	c.Host = "http://localhost:3000"
	c.Token = testToken
	res := c.DeleteClient(strconv.FormatInt(CIDac, 10))
	fmt.Print("res deleted: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestClientService_DeleteClientReq(t *testing.T) {
	var c ClientService
	c.ClientID = "403"
	c.Host = "http://localhost:30001"
	c.Token = testToken
	res := c.DeleteClient(strconv.FormatInt(CIDac, 10))
	fmt.Print("res deleted: ")
	fmt.Println(res)
	if res.Success == true {
		t.Fail()
	}
}
