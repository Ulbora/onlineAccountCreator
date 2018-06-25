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
	"html/template"
	manager "onlineAccountCreator/managers"
	"os"
)

//Handler Handler
type Handler struct {
	AcctService manager.GatewayAccountService
	Templates   *template.Template
}

//GetOauth2Host GetOauth2Host
func (h *Handler) GetOauth2Host() string {
	var rtn = ""
	if os.Getenv("OAUTH2_HOST") != "" {
		rtn = os.Getenv("OAUTH2_HOST")
	} else {
		rtn = "http://localhost:3000"
	}
	return rtn
}

//GetUserHost GetUserHost
func (h *Handler) GetUserHost() string {
	var rtn = ""
	if os.Getenv("USER_HOST") != "" {
		rtn = os.Getenv("USER_HOST")
	} else {
		rtn = "http://localhost:3001"
	}
	return rtn
}

//GetGwHost GetGwHost
func (h *Handler) GetGwHost() string {
	var rtn = ""
	if os.Getenv("GATEWAY_HOST") != "" {
		rtn = os.Getenv("GATEWAY_HOST")
	} else {
		rtn = "http://localhost:3011"
	}
	return rtn
}

//GetClientID GetClientID
func (h *Handler) GetClientID() string {
	var rtn = ""
	if os.Getenv("CLIENT_ID") != "" {
		rtn = os.Getenv("CLIENT_ID")
	} else {
		rtn = "403"
	}
	return rtn
}

//GetAPIKey GetAPIKey
func (h *Handler) GetAPIKey() string {
	var rtn = ""
	if os.Getenv("API_KEY") != "" {
		rtn = os.Getenv("API_KEY")
	} else {
		rtn = "403"
	}
	return rtn
}
