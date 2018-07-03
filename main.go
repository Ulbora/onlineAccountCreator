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

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	hand "onlineAccountCreator/handlers"
	mgn "onlineAccountCreator/managers"
	"os"
)

var templates *template.Template
var h hand.Handler

func main() {
	var captchaSecret string
	var credSecret string
	if len(os.Args) == 3 {
		captchaSecret = os.Args[1]
		credSecret = os.Args[2]
	}
	h.GetCaptchaSecret(captchaSecret)
	h.GetCredentialsSecret(credSecret)

	h.Templates = template.Must(template.ParseFiles("./static/index.html", "./static/header.html",
		"./static/navbar.html", "./static/status.html"))
	var ac mgn.GatewayAccountService
	ac.Host = h.GetOauth2Host()
	ac.GwHost = h.GetGwHost()
	ac.UserHost = h.GetUserHost()
	ac.ClientID = h.GetClientID()
	ac.APIKey = h.GetAPIKey()
	h.AcctService = ac
	router := mux.NewRouter()

	router.HandleFunc("/", h.HandleIndex)
	router.HandleFunc("/addAccount", h.HandleAddAccount)
	router.HandleFunc("/status", h.HandleStatus)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Online Account Creator!")
	fmt.Println("Listening on :8050...")
	http.ListenAndServe(":8050", router)

}
