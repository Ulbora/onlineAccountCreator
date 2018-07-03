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
	//"fmt"
	//"fmt"
	oauth2 "github.com/Ulbora/go-oauth2-client"
	"html/template"
	"net"
	manager "onlineAccountCreator/managers"
	sr "onlineAccountCreator/services"
	"os"
)

//Handler Handler
type Handler struct {
	AcctService      manager.GatewayAccountService
	Templates        *template.Template
	CaptchaSecret    string
	ClientCredSecret string
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
	//fmt.Print("captcha secret: ")
	//fmt.Println(h.CaptchaSecret)
}

//GetCaptchaHost GetCaptchaHost
func (h *Handler) GetCaptchaHost() string {
	var rtn string
	if os.Getenv("CAPTCHA_HOST") != "" {
		rtn = os.Getenv("CAPTCHA_HOST")
	} else {
		rtn = "https://www.google.com/recaptcha/api/siteverify"
	}
	//fmt.Print("captcha host: ")
	//fmt.Println(rtn)
	return rtn
}

//GetCredentialsSecret GetCredentialsSecret
func (h *Handler) GetCredentialsSecret(cs string) {
	if os.Getenv("OAUTH2_CREDENTIALS_SECRET") != "" {
		h.ClientCredSecret = os.Getenv("OAUTH2_CREDENTIALS_SECRET")
	} else {
		h.ClientCredSecret = cs
	}
}

//GetCredentialsToken GetCredentialsToken
func (h *Handler) GetCredentialsToken() string {
	var rtn string
	//fmt.Println("getting Client Credentials token")
	var tn oauth2.ClientCredentialsToken
	tn.OauthHost = h.GetOauth2Host()
	tn.ClientID = h.GetClientID()
	tn.Secret = h.ClientCredSecret
	//fmt.Print("ClientCredentialsToken tn: ")
	//fmt.Println(tn)
	resp := tn.ClientCredentialsToken()
	//fmt.Print("credentils token response: ")
	//fmt.Println(resp)
	if resp != nil {
		rtn = resp.AccessToken
		//fmt.Print("new credentials token: ")
		//fmt.Println(resp.AccessToken)
	}
	return rtn
}

func (h *Handler) sendCaptcha(recaptchaResp string) *sr.CaptchaResponse {

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
	// fmt.Print("captcha secret: ")
	// fmt.Println(h.CaptchaSecret)

	// fmt.Print("client ip address: ")
	// fmt.Println(ipAddr)

	// fmt.Print("recaptchaResp: ")
	// fmt.Println(recaptchaResp)

	var s sr.CaptchaService
	s.Host = h.GetCaptchaHost()
	var cap sr.Captcha
	cap.Remoteip = ipAddr
	cap.Secret = h.CaptchaSecret
	cap.Response = recaptchaResp
	res := s.SendCaptchaCall(cap)

	return res
}

// add method to send email
