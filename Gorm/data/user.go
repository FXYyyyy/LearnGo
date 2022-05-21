package data

import (
	. "firstGo/Gorm/DbInit"
	_const "firstGo/Gorm/const"
)

// CreateUser
/* @Description: 指定结构体创建
*  @param u
*  @return int
*/
func CreateUser(u *_const.Users) int {
	MyDb.Create(u)
	return u.ID
}

// CreateUserByFiled
/* @Description: 选定部分字段进行创建
*  @param filed
*  @param u
*  @return int
*/
func CreateUserByFiled(filed []interface{}, u *_const.Users) int {
	MyDb.Select(filed[0], filed[1:]...).Create(u)
	return u.ID
}

// CreateUsers
/* @Description: 批量创建，底层还是CreateInBatches分批创建
*  @param us
*/
func CreateUsers(us []_const.Users) {
	MyDb.Debug().Create(us)
}

// CreateUserByMaps
/* @Description: 使用map进行创建
*  @param m
*/
func CreateUserByMaps(m []map[string]interface{})  {
	MyDb.Model(&_const.Users{}).Create(m)
}