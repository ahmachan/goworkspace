package models

import (
	"fmt"
	"github.com/widuu/gomysql"
)

type MyModel struct {
	connection *gomysql.Model
}

var conn *gomysql.Model

func (m *MyModel) Init(filepath string) {
	conn, err := gomysql.SetConfig(filepath)
	if err != nil {
		fmt.Println(err)
	}
	m.connection = conn // 获取数据库连接
}

func (m *MyModel) InsertUser(userRes map[string]interface{}) (num int, err error) {
	t := m.connection.SetTable("user").SetPk("user_id")
	nint, err := t.Insert(userRes)

	return nint, err
}

func (m *MyModel) GetUserInfo(userId int) map[int]map[string]string {
	t := m.connection.SetTable("user")
	//map[string]string
	where := fmt.Sprintf("user_id=%d", userId)
	data := t.Fileds("user_id", "password", "username", "user_type").Where(where).FindOne()
	//fmt.Println(data)
	//fmt.Println("---------------")
	return data
}

func (m *MyModel) GetUserDetails(userId int) map[string]string {
	t := m.connection.SetTable("user")

	where := fmt.Sprintf("user_id=%d", userId)
	data := t.Fileds("user_id", "password", "username", "user_type").Where(where).Find()

	return data
}

func (m *MyModel) UpdateUser(userId int, username string) (num int, err error) {
	t := m.connection.SetTable("user")
	where := fmt.Sprintf("user_id=%d", userId)
	var value = make(map[string]interface{})
	value["username"] = username
	nint, err := t.Where(where).Update(value)

	return nint, err
}

func (m *MyModel) ModifyUser(userId int, username string) interface{} {
	sqlStr := fmt.Sprintf("update user set username='%s' where user_id =%d", username, userId)
	n := m.connection.Query(sqlStr)

	return n
}

func (m *MyModel) GetUserListByGender(gender int) map[int]map[string]string {
	t := m.connection.SetTable("user")
	where := fmt.Sprintf("gender=%d", gender)
	data := t.Fileds("user_id", "password", "username", "user_type", "gender").Where(where).OrderBy("user_id desc,user_type asc").FindAll()

	return data
}
