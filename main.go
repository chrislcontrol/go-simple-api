package main

import (
	"github/chrislcontrol/go-simple-api/routers"
	"log"
	"net/http"
)

func main() {
	routerV1 := routers.RouterV1()
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", routerV1))
}
