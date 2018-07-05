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

package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"reflect"
	"testing"
)

func TestGatewayClientUrlService_getToken(t *testing.T) {
	if testToken == "" {
		req, _ := http.NewRequest("POST", tokenURL, nil)
		client := &http.Client{}
		resp, cErr := client.Do(req)
		if cErr != nil {
			fmt.Print("Client Add err: ")
			fmt.Println(cErr)
		} else {
			defer resp.Body.Close()
			var tres TokenResponse
			decoder := json.NewDecoder(resp.Body)
			error := decoder.Decode(&tres)
			if error != nil {
				log.Println(error.Error())
			} else {
				testToken = tres.Token
				//fmt.Print("token: ")
				//fmt.Println(testToken)
			}
		}
	}
}

func TestMailServerService_SendMail(t *testing.T) {
	var fromEmail = "some.some1.com"
	var toEmail = "another.some1.com"
	var m MailServerService
	m.ClientID = "403"
	m.APIKey = "403"
	m.Token = testToken
	m.Host = "http://localhost:3002"
	var mm MailMessage
	mm.FromEmail = fromEmail
	mm.ToEmail = toEmail
	mm.Subject = "Ulbora CMS V3 message"
	mm.TextMessage = "test message"
	mres := m.SendMail(&mm)
	fmt.Print("Send Mail Res: ")
	fmt.Println(mres)
	fmt.Println("Sending mail failed from " + fromEmail)
	if mres.Success == true {
		t.Fail()
	}
}
