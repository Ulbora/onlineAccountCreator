package mysqldb

/*
 Copyright (C) 2016 Ulbora Labs Inc. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2016 Ken Williamson
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

import (
	crud "github.com/Ulbora/go-crud-mysql"
)

//ConnectAuthDb connect to db
func ConnectAuthDb(host, user, pw, dbName string) bool {
	res := crud.InitializeMysql(host, user, pw, dbName)
	return res
}

//ConnectionAuthTest get a row. Passing in tx allows for transactions
func ConnectionAuthTest() *crud.DbRow {
	var as []interface{}
	rowPtr := crud.Get(ConnectionAuthTestQuery, as...)
	return rowPtr
}
