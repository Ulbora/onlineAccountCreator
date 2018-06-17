/*
 Copyright (C) 2017 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2017 Ken Williamson
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

package routes

import (
	"fmt"
	"testing"
)

func TestGetActiveRoutes(t *testing.T) {
	var s ActiveRoutes
	s.Challenge = true
	s.Content = true
	s.Customer = true
	s.Images = true
	s.Mail = true
	s.Order = true
	s.Product = true
	s.Templates = true
	rts := GetActiveRoutes(&s)
	fmt.Print("res active routes: ")
	fmt.Println(rts.RestRoutes)
	for _, r := range *rts.RestRoutes {
		fmt.Print("active route: ")
		fmt.Println(r)
		fmt.Print("active route urls: ")
		fmt.Println(r.RoutesURLs)
	}
	//fmt.Println((*rts.RestRoutes)[5].Route)
	if len(*rts.RestRoutes) != 8 || (*rts.RestRoutes)[4].Route != "challenge" {
		t.Fail()
	}
}
