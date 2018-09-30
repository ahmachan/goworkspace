package routers

import (
	"diormDm/services"
	//"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type MyRouter struct {
	*mux.Router
}

var service = new(services.MyServices)


func (r *MyRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//defer tools.CatchError(w, req)
	r.Router.ServeHTTP(w, req)
}

func NewAPIMux() http.Handler {
	r := &MyRouter{mux.NewRouter()}
	s := r.PathPrefix("/api").Subrouter()
	service.Init("./conf/conf.ini")

	initUserRouters(s)
	initRechargeRouters(s)
	
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("web/"))))
	return r
}
