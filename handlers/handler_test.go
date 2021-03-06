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
	sr "onlineAccountCreator/services"
	"os"
	"testing"
)

func TestHandler_GetOauth2Host(t *testing.T) {
	var h Handler
	var oa2host = h.GetOauth2Host()
	if oa2host != "http://localhost:3000" {
		t.Fail()
	}
}

func TestHandler_GetOauth2Host2(t *testing.T) {
	var h Handler
	os.Setenv("OAUTH2_HOST", "oauthTestHost")
	var oa2host = h.GetOauth2Host()
	if oa2host != "oauthTestHost" {
		t.Fail()
	}
}

func TestHandler_GetUserHost(t *testing.T) {
	var h Handler
	var uhost = h.GetUserHost()
	if uhost != "http://localhost:3001" {
		t.Fail()
	}
}

func TestHandler_GetUserHost2(t *testing.T) {
	os.Setenv("USER_HOST", "userTestHost")
	var h Handler
	var uhost = h.GetUserHost()
	if uhost != "userTestHost" {
		t.Fail()
	}
}

func TestHandler_GetGwHost(t *testing.T) {
	var h Handler
	var gwhost = h.GetGwHost()
	if gwhost != "http://localhost:3011" {
		t.Fail()
	}
}

func TestHandler_GetGwHost2(t *testing.T) {
	os.Setenv("GATEWAY_HOST", "gwTestHost")
	var h Handler
	var gwhost = h.GetGwHost()
	if gwhost != "gwTestHost" {
		t.Fail()
	}
}

func TestHandler_GetClientID(t *testing.T) {
	var h Handler
	var cl = h.GetClientID()
	if cl != "403" {
		t.Fail()
	}
}

func TestHandler_GetClientID2(t *testing.T) {
	os.Setenv("CLIENT_ID", "555")
	var h Handler
	var cl = h.GetClientID()
	if cl != "555" {
		t.Fail()
	}
}

func TestHandler_GetAPIKey(t *testing.T) {
	var h Handler
	var k = h.GetAPIKey()
	if k != "403" {
		t.Fail()
	}
}

func TestHandler_GetAPIKey2(t *testing.T) {
	os.Setenv("API_KEY", "555")
	var h Handler
	var k = h.GetAPIKey()
	if k != "555" {
		t.Fail()
	}
}

func TestHandler_GetCaptchaSecret(t *testing.T) {
	var h Handler
	h.GetCaptchaSecret("125")
	if h.CaptchaSecret != "125" {
		t.Fail()
	}
}

func TestHandler_GetCaptchaSecret2(t *testing.T) {
	os.Setenv("CAPTCHA_SECRET", "555")
	var h Handler
	h.GetCaptchaSecret("")
	if h.CaptchaSecret != "555" {
		t.Fail()
	}
}

func TestHandler_GetCaptchaHost(t *testing.T) {
	var h Handler
	hst := h.GetCaptchaHost()
	if hst != "https://www.google.com/recaptcha/api/siteverify" {
		t.Fail()
	}
}

func TestHandler_GetCaptchaHost2(t *testing.T) {
	os.Setenv("CAPTCHA_HOST", "555")
	var h Handler
	hst := h.GetCaptchaHost()
	if hst != "555" {
		t.Fail()
	}
}

func TestHandler_GetCredentialsSecret(t *testing.T) {
	var h Handler
	h.GetCredentialsSecret("125")
	if h.ClientCredSecret != "125" {
		t.Fail()
	}
}

func TestHandler_GetCredentialsSecret2(t *testing.T) {
	os.Setenv("OAUTH2_CREDENTIALS_SECRET", "555444")
	var h Handler
	h.GetCredentialsSecret("")
	if h.ClientCredSecret != "555444" {
		t.Fail()
	}
}

func TestHandler_GetCredentialsToken(t *testing.T) {
	var h Handler
	os.Setenv("OAUTH2_HOST", "")
	os.Setenv("CLIENT_ID", "")
	h.ClientCredSecret = "badsecret"
	res := h.GetCredentialsToken()
	fmt.Print("credentils token: ")
	fmt.Println(res)
	if res != "" {
		t.Fail()
	}
}

func TestHandler_sendCaptcha(t *testing.T) {
	os.Setenv("CAPTCHA_HOST", "")
	var h Handler
	res := h.sendCaptcha("test")
	if res.Success != false {
		t.Fail()
	}
}

func TestHandler_GetFromEmailAddress(t *testing.T) {
	var h Handler
	h.GetFromEmailAddress("125")
	if h.FromEmailAddress != "125" {
		t.Fail()
	}
}

func TestHandler_GetFromEmailAddress2(t *testing.T) {
	os.Setenv("FROM_EMAIL_ADDRESS", "555444")
	var h Handler
	h.GetFromEmailAddress("")
	if h.FromEmailAddress != "555444" {
		t.Fail()
	}
}

func TestHandler_GetMailHost(t *testing.T) {
	var h Handler
	m := h.GetMailHost()
	if m != "http://localhost:3002" {
		t.Fail()
	}
}

func TestHandler_GetMailHost2(t *testing.T) {
	os.Setenv("MAIL_HOST", "555444")
	var h Handler
	m := h.GetMailHost()
	if m != "555444" {
		t.Fail()
	}
}

func TestHandler_sendEmail(t *testing.T) {
	os.Setenv("FROM_EMAIL_ADDRESS", "")
	os.Setenv("MAIL_HOST", "")
	var h Handler
	var mm sr.MailMessage
	mm.ToEmail = "some@some.com"
	mm.Subject = "Welcome to MyApiGateway.com"
	mm.HTMLMessage = "test message"
	res := h.sendMail(&mm)
	if res.Success != false {
		t.Fail()
	}
}
