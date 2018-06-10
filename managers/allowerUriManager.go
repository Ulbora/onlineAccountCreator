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
				failed, aures := insertAllowedURI(acct.ClientID, val.URI, au)
				g.insertRoleURI(failed, rMap[val.Role], aures.ID)
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

func insertAllowedURI(cid string, uri string, au services.AllowedURIService) (bool, *services.AllowedURIResponse) {
	var failed = false
	var auu services.AllowedURI
	auu.ClientID, _ = strconv.ParseInt(cid, 10, 64)
	auu.URI = uri
	aures := au.AddAllowedURI(&auu)
	if !aures.Success {
		fmt.Print("error inserting record:")
		fmt.Println(uri)
		fmt.Println(aures)
		failed = true
	}
	return failed, aures
}

func (g *GatewayAccountService) insertRoleURI(failed bool, roleID int64, auID int64) bool {
	//fmt.Print("failed in insert role uri:")
	//fmt.Println(failed)
	var rtn = false
	if !failed {
		var crr services.RoleURI
		crr.ClientRoleID = roleID
		crr.ClientAllowedURIID = auID
		var cr services.RoleURIService
		cr.ClientID = g.ClientID
		cr.Host = g.Host
		cr.Token = g.Token

		crres := cr.AddRoleURI(&crr)
		//fmt.Print("inserting url:")
		//fmt.Println(crres)
		if !crres.Success {
			fmt.Print("inserting url error:")
			fmt.Println(crres)
		} else {
			//fmt.Print("inserting url record:")
			//fmt.Println(crres)
			rtn = true
		}
	}
	return rtn
}
