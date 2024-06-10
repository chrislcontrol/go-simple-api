package main

import (
	"fmt"
	"github/chrislcontrol/go-simple-api/routers"
	"log"
	"net/http"
)


func main() {
	routerV1 := routers.RouterV1()
	url := "127.0.0.1:8000"
	
	fmt.Println("Application running on: " + url)

	log.Fatal(http.ListenAndServe(url, routerV1))

}
