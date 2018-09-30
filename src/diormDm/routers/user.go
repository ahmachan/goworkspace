package routers

import (
	//"diormDm/services"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func initUserRouters(r *mux.Router) {
	//service.Init("./conf/conf.ini")
	/*
		if this.service == nil {
		   this.service = new(services.MyServices)
		}
		this.service.Init("./conf/conf.ini")
	*/
	//@api/user/login
	s := r.PathPrefix("/user").Subrouter()
	s.HandleFunc("/list", UserListHandler)
	s.HandleFunc("/add", UserAddHandler)
	s.HandleFunc("/login", UserLoginhandler)
}

func UserLoginhandler(w http.ResponseWriter, r *http.Request) {
	/*
		user := &model.User{}
		tools.UnMarshalJson(r, user)
		resp := &model.Resp{Code: "1001", Msg: "账号或者密码错误"}
		if user.Username == "sweetop" && user.Password == "123456" {
			resp = &model.Resp{Code: "0", Msg: "success"}
		}
		tools.MarshalJson(w, resp)
	*/
	w.Write([]byte("user login handler"))
}

func UserListHandler(w http.ResponseWriter, r *http.Request) {
	/*
		resp := &model.Resp{Code: "0", Msg: "success"}
		users := make([]model.User, 0)
		resp.Data = append(users, model.User{Username: "sweetop"})
		tools.MarshalJson(w, resp)
	*/
	service.GetUserListByGender(1)
}

func UserListAction() {
	/*
		resp := &model.Resp{Code: "0", Msg: "success"}
		users := make([]model.User, 0)
		resp.Data = append(users, model.User{Username: "sweetop"})
		tools.MarshalJson(w, resp)
	*/
	service.GetUserListByGender(1)
}

func UserAddHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("user add"))
	var value = make(map[string]interface{})
	value["username"] = "hady"
	value["password"] = "123456"
	value["user_type"] = 2
	value["gender"] = 2
	newUserId, _ := service.InsertUser(value)
	fmt.Printf("\nnew user id :%d\n", newUserId)
}
