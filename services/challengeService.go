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
	//"bytes"
	//	"encoding/json"
	//"fmt"
	//"log"
	"net/http"
	cm "onlineAccountCreator/common"
)

//ChallengeService service
type ChallengeService struct {
	Host     string
	ClientID string
	APIKey   string
}

//Challenge template
type Challenge struct {
	Question string `json:"question"`
	Key      string `json:"key"`
	Answer   string `json:"answer"`
}

//ChallengeResponse res
type ChallengeResponse struct {
	Success bool  `json:"success"`
	ID      int64 `json:"id"`
	Code    int   `json:"code"`
}

//SendChallenge challenge
func (c *ChallengeService) SendChallenge(chal *Challenge) *ChallengeResponse {
	var rtn = new(ChallengeResponse)
	var sURL = c.Host + "/rs/challenge"
	aJSON := cm.GetJSONEncode(chal)
	req, fail := cm.GetRequest(sURL, http.MethodPost, aJSON)
	if !fail {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("u-client-id", c.ClientID)
		req.Header.Set("u-api-key", c.APIKey)
		code := cm.ProcessServiceCall(req, &rtn)
		rtn.Code = code
	}
	return rtn
}

// GetChallenge get challenge
func (c *ChallengeService) GetChallenge(lan string) *Challenge {
	var rtn = new(Challenge)
	var gURL = c.Host + "/rs/challenge/" + lan
	//fmt.Println(gURL)
	req, fail := cm.GetRequest(gURL, http.MethodGet, nil)
	if !fail {
		req.Header.Set("u-client-id", c.ClientID)
		req.Header.Set("u-api-key", c.APIKey)
		cm.ProcessServiceCall(req, &rtn)

	}
	return rtn
}
