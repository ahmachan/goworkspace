package routers

import (
	//"diormDm/services"
	//"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

//var service = new(services.MyServices)

func initRechargeRouters(r *mux.Router) {

	//@api/user/login
	s := r.PathPrefix("/recharge").Subrouter()
	s.HandleFunc("/sign", Signhandler)
}

func Signhandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("recharge sign handler"))
}