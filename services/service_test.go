/*
 Copyright (C) 2017 Ulbora Labs Inc. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2017 Ken Williamson
 All rights reserved.

 Certain inventions and disclosures in this file may be claimed within
 patents owned or patent applications filed by Ulbora Labs Inc., or third
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
	"fmt"
	"net/http"
	"testing"
)

type GatewayClusterRouteURL struct {
	RouteID           int64  `json:"routeId"`
	Route             string `json:"route"`
	URLID             int64  `json:"urlId"`
	Name              string `json:"name"`
	URL               string `json:"url"`
	Active            bool   `json:"active"`
	CircuitOpen       bool   `json:"circuitOpen"`
	OpenFailCode      int    `json:"openFailCode"`
	FailoverRouteName string `json:"failoverRouteName"`
}

func TestProcessResponse(t *testing.T) {
	//var p GatewayClusterRouteURL
	//var r = new(http.Response)
	//r := httptest.NewRecorder()
	var rtn = make([]GatewayClusterRouteURL, 0)
	client := &http.Client{}
	req, reqperr := http.NewRequest("GET", "http://localhost:3011/rs/cluster/routes/challenge", nil)
	req.Header.Set("u-client-id", "403")
	req.Header.Set("u-api-key", "403")
	fmt.Print("reqperr: ")
	fmt.Println(reqperr)
	fmt.Print("req: ")
	fmt.Println(req)
	resp, reserr := client.Do(req)
	fmt.Print("reserr: ")
	fmt.Println(reserr)
	fmt.Print("resp: ")
	fmt.Println(resp)
	defer resp.Body.Close()
	// decoder := json.NewDecoder(resp.Body)
	// error := decoder.Decode(&rtn)
	// if error != nil {
	// 	log.Println(error.Error())
	// }
	// fmt.Print("response: ")
	// fmt.Println(rtn)

	suc := ProcessResponse(resp, &rtn)
	if suc != true {
		t.Fail()
	} else {
		fmt.Print("response: ")
		fmt.Println(rtn)
	}
}

func TestProcessResponseNil(t *testing.T) {
	var p GatewayClusterRouteURL
	suc := ProcessResponse(nil, &p)
	if suc != false {
		t.Fail()
	}
}

func TestProcessResponseMethod(t *testing.T) {
	var p GatewayClusterRouteURL
	//var r = new(http.Response)
	//r := httptest.NewRecorder()
	client := &http.Client{}
	req, reqperr := http.NewRequest("POST", "http://google.com", nil)
	fmt.Print("reqperr: ")
	fmt.Println(reqperr)
	fmt.Print("req: ")
	fmt.Println(req)
	resp, reserr := client.Do(req)
	fmt.Print("reserr: ")
	fmt.Println(reserr)
	fmt.Print("resp: ")
	fmt.Println(resp)
	suc := ProcessResponse(resp, &p)
	if suc != false {
		t.Fail()
	}
}

func TestProcessResponseArrayErr(t *testing.T) {
	var p GatewayClusterRouteURL
	//var r = new(http.Response)
	//r := httptest.NewRecorder()
	client := &http.Client{}
	req, reqperr := http.NewRequest("GET", "http://localhost:3011/rs/cluster/routes/challenge", nil)
	req.Header.Set("u-client-id", "403")
	req.Header.Set("u-api-key", "403")
	fmt.Print("reqperr: ")
	fmt.Println(reqperr)
	fmt.Print("req: ")
	fmt.Println(req)
	resp, reserr := client.Do(req)
	fmt.Print("reserr: ")
	fmt.Println(reserr)
	fmt.Print("resp: ")
	fmt.Println(resp)
	suc := ProcessResponse(resp, &p)
	if suc != false {
		t.Fail()
	}
}

func Test_getJSONEncode(t *testing.T) {
	var p GatewayClusterRouteURL
	j := GetJSONEncode(&p)
	fmt.Print("j: ")
	fmt.Println(j)
	if j == nil {
		t.Fail()
	}
}
