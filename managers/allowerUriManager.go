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
	ulboraUris "ApiGatewayAdminPortal/ulborauris"
	"fmt"
	"strconv"
	"sync"
)

// const (
// 	codeGrant     = "code"
// 	implicitGrant = "implicit"
// 	clientGrant   = "client_credentials"
// )

//AddAllowedUris AddAllowedUris
func (g *GatewayAccountService) AddAllowedUris(acct *GatewayAccount, usel *ulboraUris.UlboraSelection) *services.RoleURIResponse {
	var rtn services.RoleURIResponse
	uuris := ulboraUris.GetUlboraURIs(usel)
	fmt.Println(uuris)

	var rs services.ClientRoleService
	rs.ClientID = g.ClientID
	rs.Host = g.Host
	rs.Token = g.Token
	rr := rs.GetClientRoleList(acct.ClientID)

	rMap := make(map[string]int64)

	for _, rrr := range *rr {
		rMap[rrr.Role] = rrr.ID
	}
	fmt.Println(rMap)

	var au services.AllowedURIService
	au.ClientID = g.ClientID
	au.Host = g.Host
	au.Token = g.Token
	var insCnt = 0
	var tmpUris = make([]ulboraUris.UlborURIs, 0)
	for _, uuri := range *uuris {
		tmpUris = append(tmpUris, uuri)
	}
	var failed = false
	var wg sync.WaitGroup
	for i := range tmpUris {
		wg.Add(1)
		//fmt.Println(tmpUris[i].URI)
		//var tu = tmpUris[i].URI
		go func(val ulboraUris.UlborURIs) {
			//fmt.Print("in thread: ")
			defer wg.Done()
			//fmt.Println(val.URI)
			insCnt++
			if rMap[val.Role] != 0 {
				var auu services.AllowedURI
				auu.ClientID, _ = strconv.ParseInt(acct.ClientID, 10, 64)
				auu.URI = val.URI
				aures := au.AddAllowedURI(&auu)
				if aures.Success != true {
					fmt.Print("error inserting record:")
					fmt.Println(val.URI)
					fmt.Println(aures)
					failed = true
				} else {
					var crr services.RoleURI
					crr.ClientRoleID = rMap[val.Role]
					crr.ClientAllowedURIID = aures.ID
					var cr services.RoleURIService
					cr.ClientID = g.ClientID
					cr.Host = g.Host
					cr.Token = g.Token

					crres := cr.AddRoleURI(&crr)
					if crres.Success != true {
						fmt.Println(crres)
					}
				}
			}
		}(tmpUris[i])
	}
	wg.Wait()
	if !failed {
		rtn.Success = true
	}
	fmt.Print("role url cnt: ")
	fmt.Println(insCnt)

	return &rtn
}
