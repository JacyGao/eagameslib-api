/* Start the service process from here */
package main

import (
    "log"
    "net/http"
    "github.com/eagames/router"
    "github.com/eagames/config"
)

func main() {
    initRouter := router.NewRouter()
    log.Printf("Listenning to %s:%s", config.HOST, config.PORT)
    log.Fatal(http.ListenAndServe(":"+config.PORT, initRouter))
}