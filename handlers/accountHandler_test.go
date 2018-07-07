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

	//"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_HandleAddAccount(t *testing.T) {
	testCap = true
	var h Handler
	w := httptest.NewRecorder()

	r, _ := http.NewRequest("POST", "/test?firstName=bob&lastName=bobby&companyName=bob and co&email=bob@bobco.com&website=bobco.com&g-recaptcha-response=55444", nil)

	h.HandleAddAccount(w, r)
}

func TestHandler_HandleStatus(t *testing.T) {
	testCap = true
	var h Handler
	h.Templates = template.Must(template.ParseFiles("status.html"))
	w := httptest.NewRecorder()

	r, _ := http.NewRequest("GET", "/test?success=true", nil)

	h.HandleStatus(w, r)
}

func TestHandler_HandleActivation(t *testing.T) {
	testCap = true
	var h Handler
	h.Templates = template.Must(template.ParseFiles("status.html"))
	w := httptest.NewRecorder()

	r, _ := http.NewRequest("GET", "/test?clientId=123&email=bob@bob.com", nil)

	h.HandleActivation(w, r)
}

func TestHandler_HandleActivation2(t *testing.T) {
	testCap = false
	var h Handler
	h.Templates = template.Must(template.ParseFiles("status.html"))
	w := httptest.NewRecorder()

	r, _ := http.NewRequest("GET", "/test?clientId=123&email=bob@bob.com", nil)

	h.HandleActivation(w, r)
}
