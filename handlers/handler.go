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
	"fmt"
	"html/template"
	"net"
	manager "onlineAccountCreator/managers"
	"os"
)

//Handler Handler
type Handler struct {
	AcctService   manager.GatewayAccountService
	Templates     *template.Template
	CaptchaSecret string
}

//GetOauth2Host GetOauth2Host
func (h *Handler) GetOauth2Host() string {
	var rtn string
	if os.Getenv("OAUTH2_HOST") != "" {
		rtn = os.Getenv("OAUTH2_HOST")
	} else {
		rtn = "http://localhost:3000"
	}
	return rtn
}

//GetUserHost GetUserHost
func (h *Handler) GetUserHost() string {
	var rtn string
	if os.Getenv("USER_HOST") != "" {
		rtn = os.Getenv("USER_HOST")
	} else {
		rtn = "http://localhost:3001"
	}
	return rtn
}

//GetGwHost GetGwHost
func (h *Handler) GetGwHost() string {
	var rtn string
	if os.Getenv("GATEWAY_HOST") != "" {
		rtn = os.Getenv("GATEWAY_HOST")
	} else {
		rtn = "http://localhost:3011"
	}
	return rtn
}

//GetClientID GetClientID
func (h *Handler) GetClientID() string {
	var rtn string
	if os.Getenv("CLIENT_ID") != "" {
		rtn = os.Getenv("CLIENT_ID")
	} else {
		rtn = "403"
	}
	return rtn
}

//GetAPIKey GetAPIKey
func (h *Handler) GetAPIKey() string {
	var rtn string
	if os.Getenv("API_KEY") != "" {
		rtn = os.Getenv("API_KEY")
	} else {
		rtn = "403"
	}
	return rtn
}

//GetCaptchaSecret GetCaptchaSecret
func (h *Handler) GetCaptchaSecret(ps string) {
	if os.Getenv("CAPTCHA_SECRET") != "" {
		h.CaptchaSecret = os.Getenv("CAPTCHA_SECRET")
	} else {
		h.CaptchaSecret = ps
	}
	fmt.Print("captcha secret: ")
	fmt.Println(h.CaptchaSecret)
}

func (h *Handler) sendCaptcha(recaptchaResp string) bool {
	var rtn = false

	var ipAddr string

	addrs, _ := net.InterfaceAddrs()
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipAddr = ipnet.IP.String()
				break
			}
		}
	}

	fmt.Print("client ip address: ")
	fmt.Println(ipAddr)
	return rtn
}
