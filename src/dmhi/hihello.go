package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"dmhi/router"

	"dm01"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kylelemons/go-gypsy/yaml"
)

const (
	PORT = "8086"
)

func main() {
	fmt.Println("hi golang")
	sum := dm01.Sum(3, 4.2)
	fmt.Println("the sum is:", sum)

	config, err := yaml.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Get("path"))
	fmt.Println(config.GetBool("enabled"))

	fdName := "badBoy"
	//fdSize:=256.32
	spec := "special clothes"
	token := "gT5mKcp5yPgoE38b8fzNheYZaBLW1zkjM20wm1"
	fileType := dm01.FileTypeDmMismatch{
		Full:     spec,
		Spec:     spec,
		Token:    token,
		Expected: "yaml.Scalar",
		Name:     fdName,
		//Size:     fdSize
	}

	log.Print(fileType.ShowInfo())
	log.Print("\n******\n")
	fileType.ShowToken()
	log.Print("\n******\n")
	fileType.FunPtrA1()
	log.Print("\n******\n")
	fileType.FunPtrA2()
	log.Print("\n******\n")

	// Create the Gin engine.
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	// Routes.
	router.Load(
		// Cores.
		g,
		// Middlwares.
		middlewares...,
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	hostPort := fmt.Sprintf(":%s", PORT)
	log.Printf("Start to listening the incoming requests on http address: %s", hostPort)
	log.Printf(http.ListenAndServe(hostPort, g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	baseHost := fmt.Sprintf("http://127.0.0.1:%s", PORT)
	for i := 0; i < 3; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(baseHost + "/users/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Print(resp)
		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
