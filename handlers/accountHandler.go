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
	//"bytes"
	"fmt"
	//"net"
	//"strings"
	//"github.com/gorilla/mux"
	"net/http"
)

//HandleAddAccount HandleAddAccount
func (h *Handler) HandleAddAccount(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	email := r.FormValue("email")
	companyName := r.FormValue("companyName")
	website := r.FormValue("website")
	recaptchaResp := r.FormValue("g-recaptcha-response")

	fmt.Print("firstName: ")
	fmt.Println(firstName)

	fmt.Print("lastName: ")
	fmt.Println(lastName)

	fmt.Print("email: ")
	fmt.Println(email)

	fmt.Print("companyName: ")
	fmt.Println(companyName)

	fmt.Print("website: ")
	fmt.Println(website)

	fmt.Print("recaptchaResp: ")
	fmt.Println(recaptchaResp)

	h.sendCaptcha(recaptchaResp)

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
