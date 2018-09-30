package main

import (
	"fmt"
	"log"
	"net/http"
	//"time"
	//"github.com/widuu/gomysql"
	"diormDm/routers"
	//"./services"
)

func main() {
	//service := new(services.MyServices)
	//service.Init("./conf/conf.ini")
	/*
		service.GetUserInfo(7)
		service.GetUserDetails(6)
		service.UpdateUser(6,"wendy")
	*/
	/*
		service.GetUserListByGender(1)
		service.GetUserListByGender(2)

		var value = make(map[string]interface{})
		value["username"] = "hady"
		value["password"] = "123456"
		value["user_type"]=2
		value["gender"]=2
		newUserId,_ := service.InsertUser(value)
		fmt.Printf("\nnew user id :%d\n",newUserId)
	*/
	//if newUserId>0 {
	//   service.UpdateUser(newUserId,"Helen")
	//}
	//conf := goini.SetConfig("./conf/conf.ini")
	//charset := conf.GetValue("database", "charset")
	//fmt.Println(charset)

	//n := c.Query("INSERT INTO user (`username`,`password`,`user_type`) VALUES ('HeungLy','136600',1),('Poply','136600',2)")
	//fmt.Println(n)
	//d := c.Query("select * from user")
	//fmt.Println(d)
	//t:=c.SetTable("user")
	//fmt.Println(t.Fileds("user_id,username").FindAll())

	//n := c.Query("update user set username='ceshishenma' where user_id =11 ")
	//fmt.Println(n)
	/*
			ut:=c.SetTable("user as u")
		data := ut.Fileds("u.user_id", "d.keywords", "u.username", "u.password","u.user_type","d.age").Join("user_data as d", "u.user_id = d.user_id").FindAll()
		fmt.Println(data)
	*/

	//map[int]map[string]string
	/*
		   data := t.Fileds("user_id", "password", "username","user_type").Where("user_id=7").FindOne()
		   fmt.Println(data)
		   fmt.Println("---------------")
		   userInfo:=data[1]

		   fmt.Println("---------------")
		   fmt.Printf("userId:%s\n",userInfo["user_id"])
		   fmt.Printf("name:%s\n",userInfo["username"])
		   fmt.Printf("userType:%s\n",userInfo["user_type"])
		   fmt.Println("---------------")
		   for _,item := range data{
			   for key,val:=range item{
				   fmt.Printf("item[%s]=%s\t%s\n",key,val,item[key])
			   }
		   }
		   fmt.Printf("\n")
		   gomysql.Print(data)

		   details := t.Fileds("user_id", "password", "username","user_type").Where("user_id=7").Find()
		   fmt.Printf(details["username"])
	*/

	/*
		n = c.Query("update user set username='ceshishenma' where id =17 ")
		fmt.Println(n)

		n = c.Query("delete from user where id=16 ")
		fmt.Println(n)

		data = c.Query("select username,password from user")
		fmt.Println(data)

		var value = make(map[string]interface{})
		value["username"] = "widuu"
		value["password"] = "widuu"
		_, err = t.Insert(value)
		fmt.Println(err)

		n, err = c.SetTable("user").Delete("id = 16")
		fmt.Println(n, err)

		var sss = make(map[string]interface{})
		sss["username"] = "widuuweb"
		r, err := t.Where("username = 'widuu'").Update(sss)
		fmt.Println(r, err)

		data = c.SetTable("user").SetParam([]string{"*"}).Where("id>1").Limit(1, 5).OrderBy("id Desc").FindAll()
		for _, v := range data {
			for k, value := range v {
				fmt.Println(k, value)
			}
		}
	*/
    /*
	srv := &http.Server{
		Addr: "0.0.0.0:3001",
		//Addr: "0.0.0.0:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Handler:      routers.NewAPIMux(),
	}
	*/

	hostPort := fmt.Sprintf(":%s", "3002")
	log.Printf("\nStart to listening the incoming requests on http address: %s\n", hostPort)

	// Run our server in a goroutine so that it doesn't block.
	/*
	go func() {		
	    if err := srv.ListenAndServe(); err != nil {
		    log.Fatal("ListenAndServe: ", err)
		    fmt.Printf("\nListenAndServe:%v\n", err)
	    }
	}()
	*/
	
	if err := http.ListenAndServe(":3002", routers.NewAPIMux()); err != nil {
		log.Fatal("ListenAndServe: ", err)
		fmt.Printf("\nListenAndServe:%v\n",err)
	}
}

/*
func startHttp() {
    if err := http.ListenAndServe(":8080", routers.NewAPIMux()); err != nil {
        log.Fatal("ListenAndServe: ", err)
        fmt.Printf("\nListenAndServe:%v\n",err)
    }
}
*/
