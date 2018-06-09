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
	"fmt"
	"math/rand"
	"time"
)

const (
	prodFreeUserPortalRedirectURL = "http://www.userportal.myapigateway.com/tokenImplicitHandler"
	localUserPortalRedirectURL    = "http://localhost:8092/tokenImplicitHandler"
	localUlboraCmsRedirectURL     = "http://localhost:8090/admin/token"
)

//GatewayAccountService GatewayAccountService
type GatewayAccountService struct {
	Token    string
	Host     string
	UserHost string
	ClientID string
}

//GatewayAccount GatewayAccount
type GatewayAccount struct {
	Name      string
	Username  string
	FirstName string
	LastName  string
	WebSite   string
	Email     string
	ClientID  string
}

//AddGatewayAccount AddGatewayAccount
func (g *GatewayAccountService) AddGatewayAccount(acct *GatewayAccount) *services.ClientResponse {
	var rtn services.ClientResponse
	var cid int64
	var c services.ClientService
	c.Token = g.Token
	c.Host = g.Host
	c.ClientID = g.ClientID
	res := g.AddOauth2Client(acct)
	if res.Success {
		cid = res.ClientID
		rtn.ClientID = cid
		fmt.Print("cid: ")
		fmt.Println(cid)
		rtn.Success = true
	}
	return &rtn
	//fmt.Println(c)
}

func generateTempPassword() string {
	return randStringRunes(20)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//RandStringRunes RandStringRunes
func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
