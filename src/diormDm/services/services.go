package services

import (
	"diormDm/models"
	"fmt"
)

type MyServices struct{}

var bingo = new(models.MyModel)

func (s *MyServices) Init(filepath string) {
	bingo.Init(filepath)
}

func (s *MyServices) InsertUser(userParams map[string]interface{}) (num int, err error) {
	res, err := bingo.InsertUser(userParams)
	fmt.Println(res)
	fmt.Println("---------------")
	return res, err
}

func (s *MyServices) GetUserInfo(userId int) {
	data := bingo.GetUserInfo(userId)
	fmt.Println(data)
	fmt.Println("---------------")
}

func (s *MyServices) GetUserDetails(userId int) {
	var data = bingo.GetUserDetails(userId)
	fmt.Println(data)
	fmt.Println("---------------")
}

func (s *MyServices) UpdateUser(userId int, username string) {
	var res, _ = bingo.UpdateUser(userId, username)
	fmt.Println(res)
	fmt.Println("---------------")
}

func (s *MyServices) ModifyUser(userId int, username string) {
	var res = bingo.ModifyUser(userId, username)
	fmt.Println(res)
	fmt.Println("---------------")
}

func (s *MyServices) GetUserListByGender(gender int) {
	var res = bingo.GetUserListByGender(gender)
	fmt.Println(res)
	fmt.Println("---------------")
}
