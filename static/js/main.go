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

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	//services "UlboraCmsV3/services"

	usession "github.com/Ulbora/go-better-sessions"
	oauth2 "github.com/Ulbora/go-oauth2-client"
	"github.com/gorilla/mux"
)

var s usession.Session

//var token *oauth2.Token
var tokenMap map[string]*oauth2.Token
var credentialToken *oauth2.Token

var templates = template.Must(template.ParseFiles("./static/index.html", "./static/header.html", "./static/headerChart.html",
	"./static/footer.html", "./static/navbar.html", "./static/clients.html", "./static/addClient.html",
	"./static/editClient.html", "./static/oauth2.html", "./static/redirectUrls.html", "./static/grantTypes.html",
	"./static/roles.html", "./static/allowedUris.html", "./static/secSideMenu.html", "./static/ulboraUris.html",
	"./static/users.html", "./static/editUser.html", "./static/gwSideMenu.html", "./static/gateway.html",
	"./static/gatewayClient.html", "./static/gatewayRoutes.html", "./static/editGatewayRoute.html",
	"./static/gatewayRouteUrls.html", "./static/gatewayRouteUrlsByRoute.html", "./static/gatewayRouteUrl.html",
	"./static/gatewayRouteUrlsStatus.html", "./static/gatewayRouteUrlsPerformance.html",
	"./static/gatewayRouteUrlPerformanceByDate.html", "./static/gatewayRouteUrlsErrors.html",
	"./static/gatewayRouteUrlViewErrors.html"))

//var username string

func main() {
	//gob.Register(oauth2.Token)
	s.MaxAge = sessingTimeToLive
	s.Name = userSession
	if os.Getenv("SESSION_SECRET_KEY") != "" {
		s.SessionKey = os.Getenv("SESSION_SECRET_KEY")
	} else {
		s.SessionKey = "115722gggg14ddfg4567"
	}

	tokenMap = make(map[string]*oauth2.Token)

	router := mux.NewRouter()

	//securety routes
	router.HandleFunc("/", handleIndex)
	router.HandleFunc("/clients", handleClients)
	router.HandleFunc("/addClient", handleAddClient)
	router.HandleFunc("/editClient/{clientId}", handleEditClient)
	router.HandleFunc("/newClient", handleNewClient)
	router.HandleFunc("/updateClient", handleUpdateClient)

	router.HandleFunc("/users/{clientId}", handleUsers)
	router.HandleFunc("/newUser", handleNewUser)
	router.HandleFunc("/editUser/{username}/{clientId}", handleEditUser)
	router.HandleFunc("/updateUserInfo", handleUpdateUserInfo)
	router.HandleFunc("/updateUserEnable", handleUpdateUserEnable)
	router.HandleFunc("/updateUserPw", handleUpdateUserPw)

	router.HandleFunc("/oauth2/{clientId}", handleOauth2)

	router.HandleFunc("/clientRedirectUrls/{clientId}", handleRedirectURLs)
	router.HandleFunc("/addRedirectUrl", handleRedirectURLAdd)
	router.HandleFunc("/deleteRedirectUri/{id}/{clientId}", handleRedirectURLDelete)

	router.HandleFunc("/clientGrantTypes/{clientId}", handleGrantType)
	router.HandleFunc("/addGrantType", handleGrantTypeAdd)
	router.HandleFunc("/deleteGrantType/{id}/{clientId}", handleGrantTypeDelete)

	router.HandleFunc("/clientRoles/{clientId}", handleRoles)
	router.HandleFunc("/addClientRole", handleRoleAdd)
	router.HandleFunc("/deleteClientRoles/{id}/{clientId}", handleRoleDelete)

	router.HandleFunc("/clientAllowedUris/{clientId}", handleAllowedUris)
	router.HandleFunc("/addAllowedUri", handleAllowedUrisAdd)
	router.HandleFunc("/editAllowedUri", handleAllowedUrisUpdate)
	router.HandleFunc("/deleteAllowedUri/{id}/{roleId}/{clientId}", handleAllowedUrisDelete)

	router.HandleFunc("/ulboraUris/{clientId}", handleUlboraUris)
	router.HandleFunc("/ulboraUrisAdd", handleUlboraUrisAdd)

	// gateway client
	router.HandleFunc("/gateway/{clientId}", handleGateway)
	router.HandleFunc("/gatewayClient/{clientId}", handleGatewayClient)
	router.HandleFunc("/addGatewayClient", handleAddGatewayClient)
	router.HandleFunc("/updateGatewayClient", handleUpdateGatewayClient)

	//gateway routes
	router.HandleFunc("/gatewayRoutes/{clientId}", handleRoutes)
	router.HandleFunc("/addGatewayRoute", handleRoutesAdd)
	router.HandleFunc("/deleteGatewayRoute/{id}/{clientId}", handleRoutesDelete)
	router.HandleFunc("/editGatewayRoute/{id}/{clientId}", handleRouteEdit)
	router.HandleFunc("/updateGatewayRoute", handleRouteUpdate)

	//gateway route uris
	router.HandleFunc("/gatewayRouteUrls/{clientId}", handleRouteURLs)
	router.HandleFunc("/gatewayRouteUrlsByRoute/{id}/{clientId}", handleRouteURLsByRoute)
	router.HandleFunc("/addGatewayRouteUrl", handleRouteURLAdd)
	router.HandleFunc("/activateGatewayRouteUrl/{id}/{routeId}/{clientId}", handleRouteURLActivate)
	router.HandleFunc("/updateGatewayRouteUrl", handleRouteURLUpdate)
	router.HandleFunc("/editGatewayRouteUrl/{id}/{routeId}/{clientId}", handleRouteURLEdit)
	router.HandleFunc("/deleteGatewayRouteUrl/{id}/{routeId}/{clientId}", handleRouteURLDelete)

	//gateway status uris
	router.HandleFunc("/gatewayRouteUrlsStatus/{routeId}/{clientId}", handleRouteURLsStatus)
	router.HandleFunc("/resetBreaker/{urlId}/{routeId}/{clientId}", handleResetBreaker)

	//gateway performance
	router.HandleFunc("/gatewayRouteUrlsPermformance/{routeId}/{clientId}", handleRouteURLsPerformance)
	router.HandleFunc("/viewPerformanceByDate/{urlId}/{routeId}/{clientId}", handleRouteURLPerformanceByDate)

	//gateway errors
	router.HandleFunc("/gatewayRouteUrlsErrors/{routeId}/{clientId}", handleRouteURLsErrors)
	router.HandleFunc("/viewErrors/{urlId}/{routeId}/{clientId}", handleRouteURLError)
	router.HandleFunc("/viewErrors/{urlId}/{routeId}/{clientId}/{page}", handleRouteURLError)

	router.HandleFunc("/tokenHandler", handleToken)
	router.HandleFunc("/login", handleLogin)
	router.HandleFunc("/logout", handleLogout)

	// admin resources
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	//http.Handle("/js", fs)

	fmt.Println("API Gateway Admin Portal running!")
	log.Println("Listening on :8091...")
	http.ListenAndServe(":8091", router)
}
