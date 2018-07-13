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
	//"github.com/gorilla/mux"
	"net/http"
)

//PageParams PageParams
type PageParams struct {
	UlboraCms bool
}

//HandleIndex HandleIndex
func (h *Handler) HandleIndex(w http.ResponseWriter, r *http.Request) {
	var ulboracms = r.URL.Query().Get("ulboraCms")
	// vars := mux.Vars(r)
	// page := vars["content"]
	// if page == "" {
	// 	page = "home"
	// }
	//fmt.Print("page in handler: ")
	//fmt.Println(page)

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
	var p PageParams
	if ulboracms == "true" {
		p.UlboraCms = true
	}

	h.Templates.ExecuteTemplate(w, "index.html", &p)

	//}
}
