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

package handlers

import (
	"strconv"
	"time"
	//"time"
	//"bytes"
	"fmt"
	//"net"
	//"strings"
	//"github.com/gorilla/mux"
	ulboraUris "ApiGatewayAdminPortal/ulborauris"
	"net/http"
	mgr "onlineAccountCreator/managers"
)

var testCap bool

//Page Page
type Page struct {
	Success bool
}

//HandleAddAccount HandleAddAccount
func (h *Handler) HandleAddAccount(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	email := r.FormValue("email")
	companyName := r.FormValue("companyName")
	website := r.FormValue("website")
	recaptchaResp := r.FormValue("g-recaptcha-response")

	// fmt.Print("firstName: ")
	// fmt.Println(firstName)

	// fmt.Print("lastName: ")
	// fmt.Println(lastName)

	// fmt.Print("email: ")
	// fmt.Println(email)

	// fmt.Print("companyName: ")
	// fmt.Println(companyName)

	// fmt.Print("website: ")
	// fmt.Println(website)

	// fmt.Print("recaptchaResp: ")
	// fmt.Println(recaptchaResp)

	res := h.sendCaptcha(recaptchaResp)
	fmt.Print("captcha res: ")
	fmt.Println(res)
	if res.Success || testCap {
		fmt.Print("captcha res when true: ")
		fmt.Println(res)
		var gws mgr.GatewayAccountService
		gws.Host = h.GetOauth2Host()
		gws.GwHost = h.GetGwHost()
		gws.UserHost = h.GetUserHost()
		gws.ClientID = h.GetClientID()
		gws.APIKey = h.GetAPIKey()
		gws.Token = h.GetCredentialsToken()
		var acct mgr.GatewayAccount
		acct.FirstName = firstName
		acct.LastName = lastName
		acct.Name = companyName
		acct.Email = email
		acct.WebSite = website
		acct.Username = email
		var sel ulboraUris.UlboraSelection
		sel.Oauth2 = true
		sel.APIGateway = true
		acct.UlboraSelected = &sel
		time.Sleep(5000 * time.Millisecond)
		var success = "false"
		//h.Templates.ExecuteTemplate(w, "index.html", nil)
		resAcct, pw := gws.AddGatewayAccount(&acct)
		if resAcct.Success || testCap {
			fmt.Print("Add gw acct res: ")
			fmt.Println(resAcct)

			fmt.Print("Add gw acct password: ")
			fmt.Println(pw)
			// send email to user-----------

		}

		http.Redirect(w, r, "/status?success="+success, http.StatusFound)

	}

	// var ipAddr string

	// addrs, _ := net.InterfaceAddrs()
	// for _, a := range addrs {
	// 	if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	// 		if ipnet.IP.To4() != nil {
	// 			ipAddr = ipnet.IP.String()
	// 			break
	// 		}
	// 	}
	// }

	// fmt.Print("client ip address: ")
	// fmt.Println(ipAddr)

	//h.Templates.ExecuteTemplate(w, "index.html", nil)
}

//HandleStatus HandleStatus
func (h *Handler) HandleStatus(w http.ResponseWriter, r *http.Request) {
	var successStr = r.URL.Query().Get("success")
	var success, _ = strconv.ParseBool(successStr)
	var p Page
	p.Success = success
	// vars := mux.Vars(r)
	// page := vars["content"]
	// if page == "" {
	// 	page = "home"
	// }
	fmt.Print("success: ")
	fmt.Println(success)

	//if page != "favicon.ico" {
	// var c services.ContentPageService
	// c.ClientID = getAuthCodeClient()
	// c.APIKey = getGatewayAPIKey()
	// c.Host = getContentHost()
	// c.PageSize = 100
	// h, res := c.GetPage(page)
	// var pg = new(pageContent)
	// pg.Cont = res
	// pg.MetaAuthor = h.MetaAuthor
	// pg.MetaKeyWords = h.MetaKeyWords
	// pg.MetaDesc = h.MetaDesc
	// pg.Title = h.Title

	h.Templates.ExecuteTemplate(w, "status.html", &p)

	//}
}
