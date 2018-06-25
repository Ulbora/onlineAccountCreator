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
	urts "onlineAccountCreator/routes"
	"strconv"
)

//AddGwRoutes AddGwRoutes
func (g *GatewayAccountService) AddGwRoutes(acct *GatewayAccount) *services.GatewayResponse {
	var rtn services.GatewayResponse
	rts := urts.GetActiveRoutes(acct.ActiveRoutes)
	//fmt.Print("rts in routes: ")
	//fmt.Println(rts)

	var r services.GatewayRouteService
	r.ClientID = g.ClientID
	r.Host = g.GwHost
	r.Token = g.Token

	var ru services.GatewayRouteURLService
	ru.ClientID = g.ClientID
	ru.Host = g.GwHost
	ru.Token = g.Token

	for _, rt := range *rts.RestRoutes {
		var rr services.GatewayRoute
		rr.ClientID, _ = strconv.ParseInt(acct.ClientID, 10, 64)
		rr.Route = rt.Route
		res := r.AddRoute(&rr)
		rtn.Success = res.Success
		//fmt.Print("res in route loop: ")
		//fmt.Println(res)
		for _, rtu := range *rt.RoutesURLs {
			var rru services.GatewayRouteURL
			rru.ClientID, _ = strconv.ParseInt(acct.ClientID, 10, 64)
			rru.RouteID = res.ID
			rru.Name = rtu.Name
			rru.URL = rtu.URL
			ru.AddRouteURL(&rru)
			//fmt.Print("resu in url loop: ")
			//fmt.Println(resu)
		}
	}
	return &rtn
}
