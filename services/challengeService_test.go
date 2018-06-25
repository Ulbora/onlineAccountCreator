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
	"fmt"
	//"reflect"
	"testing"
)

var key string

func TestChallengeService_GetChallenge(t *testing.T) {
	var c ChallengeService
	c.Host = "http://localhost:3003"
	res := c.GetChallenge("en_us")
	fmt.Print("res: ")
	fmt.Println(res)
	key = res.Key
	if res.Question == "" {
		t.Fail()
	}
}

func TestChallengeService_SendChallenge(t *testing.T) {
	var c ChallengeService
	c.Host = "http://localhost:3003"
	var ch Challenge
	ch.Answer = "some answer"
	ch.Key = key
	res := c.SendChallenge(&ch)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success == true {
		t.Fail()
	}
}
