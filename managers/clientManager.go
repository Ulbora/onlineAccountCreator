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

//AddOauth2Client AddOauth2Client
func (g *GatewayAccountService) AddOauth2Client(acct *GatewayAccount) *services.ClientResponse {
	var c services.ClientService
	c.Token = g.Token
	c.Host = g.Host
	c.ClientID = g.ClientID
	var cc services.Client
	cc.Name = acct.Name
	cc.Email = acct.Email
	cc.Enabled = false
	cc.WebSite = acct.WebSite

	var uris []services.RedirectURI
	var redirectURLs = []string{prodFreeUserPortalRedirectURL, localUserPortalRedirectURL, localUlboraCmsRedirectURL}
	for i := range redirectURLs {
		var uri services.RedirectURI
		uri.URI = redirectURLs[i]
		uris = append(uris, uri)
	}
	cc.RedirectURIs = uris
	//fmt.Print("cc: ")
	//fmt.Println(cc)
	res := c.AddClient(&cc)
	//fmt.Print("res in add o2 client: ")
	//fmt.Println(res)
	return res
}
