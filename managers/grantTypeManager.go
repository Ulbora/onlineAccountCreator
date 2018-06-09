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
	"strconv"
)

const (
	codeGrant     = "code"
	implicitGrant = "implicit"
	clientGrant   = "client_credentials"
)

//AddClientGrantType AddClientGrantType
func (g *GatewayAccountService) AddClientGrantType(acct *GatewayAccount) *services.GrantTypeResponse {
	var gt services.GrantTypeService
	gt.Token = g.Token
	gt.Host = g.Host
	gt.ClientID = g.ClientID

	var gg services.GrantType
	gg.ClientID, _ = strconv.ParseInt(acct.ClientID, 10, 0)
	gg.GrantType = codeGrant
	gt.AddGrantType(&gg)

	gg.GrantType = implicitGrant
	gt.AddGrantType(&gg)

	gg.GrantType = clientGrant
	res := gt.AddGrantType(&gg)

	fmt.Print("res in add client user role: ")
	fmt.Println(res)
	// if !res.Success || !res2.Success || !res3.Success {
	// 	res3.Success = false
	// }
	return res
}
