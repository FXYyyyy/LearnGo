package main

import (
	"firstGo/Gorm/DbInit"
	_const "firstGo/Gorm/const"
	"fmt"
)

func main() {
	DbInit.InitMysqlDB()

	/*u := &_const.Users{
		//ID: 11,
		Name: "kk3",
		Age: 3,
		BirthDay: "2019",
	}*/

	/*
	id := data.CreateUser(u)
	fmt.Println("id = ", id)

	f := []interface{}{
		"name",
	}
	id := data.CreateUserByFiled(f, u)
	fmt.Println("id = ", id)
	 */

	/*us := []_const.Users{
		{
			Name: "xixi1",
			Age: 4,
			BirthDay: "2019",
		},
		{
			Name: "xixi2",
			Age: 4,
			BirthDay: "2019",
		},
	}
	data.CreateUsers(us)*/

	/*
	m := []map[string]interface{}{
		{
			"Name": "juju",
		},
		{
			"Name": "juju2",
		},
	}
	data.CreateUserByMaps(m)
	 */

	var ret, ret2 _const.Users
	x := DbInit.MyDb.First(&ret)
	fmt.Println("first: ", ret)
	fmt.Println("affect rows = ", x.RowsAffected)

	DbInit.MyDb.Last(&ret2)
	fmt.Println("last: ", ret2)
}
