package main

import (
    "fmt"
    "dbDmApp/controllers"
)

func main(){
    fmt.Println("show dbDmApp")
    result:=dmgz.Sqrt(8)
    fmt.Println("the result is:",result)

    userRes:=users.getAllUsers()
}
