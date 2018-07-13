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
	//"time"
	//"time"
	//"bytes"
	//"fmt"
	//"net"
	//"strings"
	//"github.com/gorilla/mux"
	services "ApiGatewayAdminPortal/services"
	ulboraUris "ApiGatewayAdminPortal/ulborauris"
	"net/http"
	mgr "onlineAccountCreator/managers"
	sr "onlineAccountCreator/services"
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
	//fmt.Print("captcha res: ")
	//fmt.Println(res)
	if res.Success || testCap {
		//fmt.Print("captcha res when true: ")
		//fmt.Println(res)
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
		///----------------------add code to turn on ulboracms
		acct.UlboraSelected = &sel
		//time.Sleep(5000 * time.Millisecond)
		var success = "false"
		//h.Templates.ExecuteTemplate(w, "index.html", nil)
		resAcct, pw := gws.AddGatewayAccount(&acct)
		if resAcct.Success || testCap {
			//fmt.Print("Add gw acct res: ")
			//fmt.Println(resAcct)
			var cidStr = strconv.FormatInt(resAcct.ClientID, 10)
			var actURL = "http://" + r.Host + "/activate?clientId=" + cidStr + "&email=" + email + ""
			var htmlMessage = "<!DOCTYPE html><html><head> <title>Free MyApiGateway.com Account</title> <meta charset='UTF-8'> <meta name='viewport' content='width=device-width, initial-scale=1, shrink-to-fit=no'></head><body> <div style='background: rgb(19, 73, 128); width: 100%; color: white; padding: 1% 0 1% 2%; margin: 0 0 1% 0; font-weight: bold; font-size: 12pt;'> MyApiGateway.com </div><div style='text-align: center'> <div style='margin: 0 0 1% 0;'>Thank you for using MyApiGateway.com</div><div style='margin: 0 0 1% 0;'>Activate your account by <a href='" + actURL + "'>clicking here.</a></div><div style='margin: 0 0 2% 0;'>Your client id is: " + cidStr + "</div><div style='margin: 0 0 2% 0;'>Your username is: " + email + "</div><div style='margin: 0 0 1% 0;'>Your new password is: " + pw + "</div></div></body></html>"
			//fmt.Print("htmlMessage")
			//fmt.Println(htmlMessage)
			var mm sr.MailMessage
			mm.ToEmail = email
			mm.Subject = "Welcome to MyApiGateway.com"
			mm.HTMLMessage = htmlMessage
			mres := h.sendMail(&mm)
			//fmt.Print("sendEmail res: ")
			//fmt.Println(mres)
			if mres.Success || testCap {
				success = "true"
			}
			//fmt.Print("Add gw acct password: ")
			//fmt.Println(pw)

		}

		http.Redirect(w, r, "/status?success="+success, http.StatusFound)

	}
}

//HandleStatus HandleStatus
func (h *Handler) HandleStatus(w http.ResponseWriter, r *http.Request) {
	var successStr = r.URL.Query().Get("success")
	var success, _ = strconv.ParseBool(successStr)
	var p Page
	p.Success = success
	//fmt.Print("success: ")
	//fmt.Println(success)

	h.Templates.ExecuteTemplate(w, "status.html", &p)

}

//HandleActivation HandleActivation
func (h *Handler) HandleActivation(w http.ResponseWriter, r *http.Request) {
	var success = false
	var clientID = r.URL.Query().Get("clientId")
	var email = r.URL.Query().Get("email")
	var c services.ClientService
	c.Token = h.GetCredentialsToken()
	c.Host = h.GetOauth2Host()
	c.ClientID = h.GetClientID()
	cl := c.GetClient(clientID)
	//fmt.Print("res c1: ")
	//fmt.Println(cl)
	if (cl.ClientID > 0 && cl.Email == email) || testCap {
		cl.Enabled = true
		c.UpdateClient(cl)
		//fmt.Print("res client act: ")
		//fmt.Println(res)
		var cg services.GatewayClientService
		cg.ClientID = h.GetClientID()
		cg.Host = h.GetGwHost()
		cg.Token = h.GetCredentialsToken()
		cg.APIKey = h.GetAPIKey()
		gwc := cg.GetClient(clientID)
		fndCid, _ := strconv.ParseInt(clientID, 10, 64)
		if (gwc.ClientID > 0 && gwc.ClientID == fndCid) || testCap {
			gwc.Enabled = true
			cg.UpdateClient(gwc)
			//fmt.Print("gwres client act: ")
			//fmt.Println(gwres)
			success = true
		}
	}
	//fmt.Print("url:")
	//fmt.Println(r.URL.Scheme + r.Host + r.URL.Path)
	//fmt.Println(r.Host, r.URL.Path)
	//success = true
	if success {
		h.Templates.ExecuteTemplate(w, "activated.html", nil)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}

}
