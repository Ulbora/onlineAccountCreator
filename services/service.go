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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//ProcessRespose ProcessRespose
func ProcessResponse(resp *http.Response, obj interface{}) bool {
	var rtn bool
	//fmt.Print("resp in processResponse: ")
	//fmt.Println(resp)
	if resp != nil {
		//fmt.Print("resp body: ")
		//fmt.Println(resp.Body)
		decoder := json.NewDecoder(resp.Body)
		var err error
		if obj != nil {
			err = decoder.Decode(obj)
			//fmt.Print("response obj: ")
			//fmt.Println(obj)
		}
		if err != nil {
			fmt.Print("response err: ")
			fmt.Println(err)
		} else {
			rtn = true
		}
	} else {
		log.Println("response = nil in processResponse")
	}
	return rtn
}

//GetJSONEncode GetJSONEncode
func GetJSONEncode(obj interface{}) *[]byte {
	//fmt.Print("obj in json: ")
	//fmt.Println(obj)
	aJSON, _ := json.Marshal(obj)
	return &aJSON
}
