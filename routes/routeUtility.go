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

package routes

const (
	contentURLBlue   = "http://content:3008"
	templateURLBlue  = "http://templates:3009"
	mailURLBlue      = "http://mail:8080"
	imagesURLBlue    = "http://images:8080"
	challengeURLBlue = "http://challenge:8080"
	productURLBlue   = "http://product:8080"
	orderURLBlue     = "http://order:8080"
	customerURLBlue  = "http://customer:8080"

	routeNameBlue = "blue"

	contentRoute   = "content"
	templateRoute  = "template"
	mailRoute      = "mail"
	imageRoute     = "image"
	challengeRoute = "challenge"
	productRoute   = "product"
	orderRoute     = "order"
	customerRoute  = "customer"
)

//GatewayRoute GatewayRoute
type GatewayRoute struct {
	ID         int64
	ClientID   int64
	Route      string
	RoutesURLs *[]GatewayRouteURL
}

//GatewayRouteURL GatewayRouteURL
type GatewayRouteURL struct {
	Name string
	URL  string
}

//GatewayRoutes GatewayRoutes
type GatewayRoutes struct {
	RestRoutes *[]GatewayRoute
}

//ActiveRoutes ActiveRoutes
type ActiveRoutes struct {
	Content   bool
	Templates bool
	Mail      bool
	Images    bool
	Challenge bool
	Product   bool
	Order     bool
	Customer  bool
}

//GetActiveRoutes GetActiveRoutes
func GetActiveRoutes(s *ActiveRoutes) *GatewayRoutes {
	var rtn GatewayRoutes
	var grs = make([]GatewayRoute, 0)

	if s != nil && s.Content {
		var ctru GatewayRouteURL
		ctru.Name = routeNameBlue
		ctru.URL = contentURLBlue

		var ctgrus = make([]GatewayRouteURL, 0)
		ctgrus = append(ctgrus, ctru)

		var ctgr GatewayRoute
		ctgr.Route = contentRoute
		ctgr.RoutesURLs = &ctgrus

		grs = append(grs, ctgr)
	}

	if s != nil && s.Templates {
		var tru GatewayRouteURL
		tru.Name = routeNameBlue
		tru.URL = templateURLBlue

		var tgrus = make([]GatewayRouteURL, 0)
		tgrus = append(tgrus, tru)

		var tgr GatewayRoute
		tgr.Route = templateRoute
		tgr.RoutesURLs = &tgrus

		grs = append(grs, tgr)
	}

	if s != nil && s.Mail {
		var mru GatewayRouteURL
		mru.Name = routeNameBlue
		mru.URL = mailURLBlue

		var mgrus = make([]GatewayRouteURL, 0)
		mgrus = append(mgrus, mru)

		var mgr GatewayRoute
		mgr.Route = mailRoute
		mgr.RoutesURLs = &mgrus

		grs = append(grs, mgr)
	}

	if s != nil && s.Images {
		var curu GatewayRouteURL
		curu.Name = routeNameBlue
		curu.URL = imagesURLBlue

		var cugrus = make([]GatewayRouteURL, 0)
		cugrus = append(cugrus, curu)

		var cugr GatewayRoute
		cugr.Route = imageRoute
		cugr.RoutesURLs = &cugrus

		grs = append(grs, cugr)
	}

	if s != nil && s.Challenge {
		var cru GatewayRouteURL
		cru.Name = routeNameBlue
		cru.URL = challengeURLBlue

		var cgrus = make([]GatewayRouteURL, 0)
		cgrus = append(cgrus, cru)

		var cgr GatewayRoute
		cgr.Route = challengeRoute
		cgr.RoutesURLs = &cgrus

		grs = append(grs, cgr)
	}

	if s != nil && s.Product {
		var pru GatewayRouteURL
		pru.Name = routeNameBlue
		pru.URL = productURLBlue

		var pgrus = make([]GatewayRouteURL, 0)
		pgrus = append(pgrus, pru)

		var pgr GatewayRoute
		pgr.Route = productRoute
		pgr.RoutesURLs = &pgrus

		grs = append(grs, pgr)
	}

	if s != nil && s.Order {
		var oru GatewayRouteURL
		oru.Name = routeNameBlue
		oru.URL = orderURLBlue

		var ogrus = make([]GatewayRouteURL, 0)
		ogrus = append(ogrus, oru)

		var ogr GatewayRoute
		ogr.Route = orderRoute
		ogr.RoutesURLs = &ogrus

		grs = append(grs, ogr)
	}

	if s != nil && s.Customer {
		var curu GatewayRouteURL
		curu.Name = routeNameBlue
		curu.URL = customerURLBlue

		var cugrus = make([]GatewayRouteURL, 0)
		cugrus = append(cugrus, curu)

		var cugr GatewayRoute
		cugr.Route = customerRoute
		cugr.RoutesURLs = &cugrus

		grs = append(grs, cugr)
	}
	rtn.RestRoutes = &grs
	return &rtn
}
