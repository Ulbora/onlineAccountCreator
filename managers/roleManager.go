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
	adminRole = "admin"
	userRole  = "user"
)

//AddClientRole AddClientRole
func (g *GatewayAccountService) AddClientRole(acct *GatewayAccount) *services.ClientRoleResponse {
	var rs services.ClientRoleService
	rs.Token = g.Token
	rs.Host = g.Host
	rs.ClientID = g.ClientID

	var rr services.ClientRole
	rr.ClientID, _ = strconv.ParseInt(acct.ClientID, 10, 0)
	rr.Role = adminRole
	res := rs.AddClientRole(&rr)
	fmt.Print("res in add client admin role: ")
	fmt.Println(res)
	rr.Role = userRole
	res2 := rs.AddClientRole(&rr)
	fmt.Print("res in add client user role: ")
	fmt.Println(res2)
	return res2
}
