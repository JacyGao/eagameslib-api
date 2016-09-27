/* Start the service process from here */
package main

import (
	"github.com/eagameslib-api/config"
	"github.com/eagameslib-api/router"
	"log"
	"net/http"
)

func main() {
	initRouter := router.NewRouter()
	log.Printf("Listenning to %s:%s", config.HOST, config.PORT)
	log.Fatal(http.ListenAndServe(":"+config.PORT, initRouter))
}
