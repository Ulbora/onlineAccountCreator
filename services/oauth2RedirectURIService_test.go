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
	"fmt"
	"strconv"
	"testing"
)

var CID2rdu int64
var rdIDrdu int64
var tokenrdu string

func TestRedirectURIService_getToken(t *testing.T) {

}

func TestRedirectURIService_AddClient(t *testing.T) {
	var c ClientService
	c.ClientID = "403"
	c.Host = "http://localhost:3000"
	c.Token = tokenrdu
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
	fmt.Print("add client res: ")
	fmt.Println(res)
	CID2rdu = res.ClientID
	if res.Success != true {
		t.Fail()
	}
}

func TestRedirectURIService_AddRedirectURI(t *testing.T) {
	var c RedirectURIService
	c.ClientID = "403"
	c.Host = "http://localhost:3000"
	c.Token = tokenrdu
	var uri RedirectURI
	uri.URI = "http://yahoooo.com"
	uri.ClientID = CID2rdu
	res := c.AddRedirectURI(&uri)

	fmt.Print("add uri res: ")
	fmt.Println(res)
	rdIDrdu = res.ID
	if res.Success != true {
		t.Fail()
	}
}

func TestRedirectURIService_GetRedirectURIList(t *testing.T) {
	var c RedirectURIService
	c.ClientID = "403"
	c.Host = "http://localhost:3000"
	c.Token = tokenrdu
	res := c.GetRedirectURIList(strconv.FormatInt(CID2rdu, 10))
	fmt.Print("uri res list: ")
	fmt.Println(res)
	fmt.Print("len: ")
	fmt.Println(len(*res))
	if res == nil || len(*res) != 2 {
		t.Fail()
	}
}

// func TestRedirectURIService_DeleteRedirectURI(t *testing.T) {
// 	var c RedirectURIService
// 	c.ClientID = "403"
// 	c.Host = "http://localhost:3000"
// 	c.Token = tempToken
// 	res := c.DeleteRedirectURI(strconv.FormatInt(rdID, 10))
// 	fmt.Print("res deleted uri: ")
// 	fmt.Println(res)
// 	if res.Success != true {
// 		t.Fail()
// 	}
// }

// func TestRedirectURIService_DeleteClient(t *testing.T) {
// 	var c ClientService
// 	c.ClientID = "403"
// 	c.Host = "http://localhost:3000"
// 	c.Token = tempToken
// 	res := c.DeleteClient(strconv.FormatInt(CID2, 10))
// 	fmt.Print("res deleted client: ")
// 	fmt.Println(res)
// 	if res.Success != true {
// 		t.Fail()
// 	}
// }
