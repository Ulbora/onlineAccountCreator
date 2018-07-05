/*
 Copyright (C) 2018 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2018 Ken Williamson
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
	//"fmt"
	"math/rand"
	urts "onlineAccountCreator/routes"
	"strconv"
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
	GwHost   string
	ClientID string
	APIKey   string
}

//GatewayAccount GatewayAccount
type GatewayAccount struct {
	Name           string
	Username       string
	FirstName      string
	LastName       string
	WebSite        string
	Email          string
	ClientID       string
	UlboraSelected *ulboraUris.UlboraSelection
	ActiveRoutes   *urts.ActiveRoutes
}

//AddGatewayAccount AddGatewayAccount
func (g *GatewayAccountService) AddGatewayAccount(acct *GatewayAccount) (*services.ClientResponse, string) {
	var rtn services.ClientResponse
	var cid int64
	var passwd string
	// var c services.ClientService
	// c.Token = g.Token
	// c.Host = g.Host
	// c.ClientID = g.ClientID
	res := g.AddOauth2Client(acct)
	if res.Success {
		cid = res.ClientID
		rtn.ClientID = cid
		//fmt.Print("cid: ")
		//fmt.Println(cid)
		acct.ClientID = strconv.FormatInt(res.ClientID, 10)
		//rtn.Success = true
		resu, pw := g.AddOauth2User(acct)
		//fmt.Print("add user in add account: ")
		//fmt.Println(resu)
		if resu.Success {
			passwd = pw
			resr := g.AddClientRole(acct)
			//fmt.Print("add roles in add account: ")
			//fmt.Println(resr)
			if resr.Success {
				resgt := g.AddClientGrantType(acct)
				//fmt.Print("add grant type in add account: ")
				//fmt.Println(resgt)
				if resgt.Success {
					resUris := g.AddAllowedUris(acct)
					//fmt.Print("add uris in add account: ")
					//fmt.Println(resUris)
					if resUris.Success {
						resgw := g.AddGwClient(acct)
						//fmt.Print("add gw client account: ")
						//fmt.Println(resgw)
						if resgw.Success {
							rtn.Success = true
						}
					}

				}
			}
		}
	}
	return &rtn, passwd
	//fmt.Println(c)
}

func generateTempPassword() string {
	return randStringRunes(20)
}

func generateAPIKey() string {
	return randStringRunes(35)
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
