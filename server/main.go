package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Maysa87/Goland_to_do/router"
)

func main() {
	r := router.Router()
	fmt.Print("starting the server on port 9000...")
	log.Fatal(http.ListenAndServe(":9000", r))
}
