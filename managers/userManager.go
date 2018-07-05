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
)

//AddOauth2User AddOauth2User
func (g *GatewayAccountService) AddOauth2User(acct *GatewayAccount) (*services.UserResponse, string) {
	var roleID int64
	var u services.UserService
	u.ClientID = g.ClientID
	u.Host = g.UserHost
	u.Token = g.Token
	resRoles := u.GetRoleList()
	for _, r := range *resRoles {
		if r.Role == "admin" {
			roleID = r.ID
			break
		}
	}
	var uu services.User
	uu.ClientID = acct.ClientID
	uu.Username = acct.Username
	uu.RoleID = roleID
	uu.FirstName = acct.FirstName
	uu.LastName = acct.LastName
	uu.EmailAddress = acct.Email
	uu.Enabled = true
	uu.Password = generateTempPassword()
	res := u.AddUser(&uu)
	//fmt.Print("res in add o2 user: ")
	//fmt.Println(res)
	return res, uu.Password
}
