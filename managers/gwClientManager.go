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
	//"fmt"

	"strconv"
)

//AddGwClient AddGwClient
func (g *GatewayAccountService) AddGwClient(acct *GatewayAccount) *services.GatewayResponse {
	var rtn services.GatewayResponse

	var c services.GatewayClientService
	c.ClientID = g.ClientID
	c.Host = g.GwHost
	c.Token = g.Token
	c.APIKey = g.APIKey
	var cc services.GatewayClient
	cc.ClientID, _ = strconv.ParseInt(acct.ClientID, 10, 64)
	cc.APIKey = generateAPIKey()
	cc.Enabled = false
	cc.Level = "small"
	//fmt.Print("gw client add req: ")
	//fmt.Println(cc)
	res := c.AddClient(&cc)
	//fmt.Print("gw client add res in manager: ")
	//fmt.Println(res)
	rtn.Success = res.Success
	rtn.ID = cc.ClientID
	rtn.Code = res.Code

	return &rtn
}
