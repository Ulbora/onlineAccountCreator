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
	//"fmt"
	//"bytes"
	//"encoding/json"
	//"fmt"
	//"log"
	"net/http"
	cm "onlineAccountCreator/common"
)

//MailServerService service
type MailServerService struct {
	Token    string
	ClientID string
	APIKey   string
	UserID   string
	Hashed   string
	Host     string
}

//MailMessage to send
type MailMessage struct {
	ToEmail     string `json:"toEmail"`
	FromEmail   string `json:"fromEmail"`
	Subject     string `json:"subject"`
	TextMessage string `json:"text"`
	HTMLMessage string `json:"html"`
}

//MailResponse res
type MailResponse struct {
	Success bool   `json:"success"`
	ID      int64  `json:"id"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

//SendMail send mail
func (m *MailServerService) SendMail(mailMessage *MailMessage) *MailResponse {
	var rtn = new(MailResponse)
	var sendURL = m.Host + "/rs/mail/send"
	aJSON := cm.GetJSONEncode(mailMessage)
	req, fail := cm.GetRequest(sendURL, http.MethodPost, aJSON)
	if !fail {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+m.Token)
		req.Header.Set("u-client-id", m.ClientID)
		req.Header.Set("clientId", m.ClientID)
		req.Header.Set("userId", m.UserID)
		req.Header.Set("hashed", m.Hashed)
		req.Header.Set("u-api-key", m.APIKey)
		code := cm.ProcessServiceCall(req, &rtn)
		rtn.Code = code
		//fmt.Print("rtn: ")
		//fmt.Println(rtn)
	}
	return rtn
}
