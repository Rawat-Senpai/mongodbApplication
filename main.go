package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Rawat-Senapi/mongoapi/router"
)

func main() {
	fmt.Println("connecting mongo db here")
	r := router.Router()
	fmt.Println("Server is getting connected ")
	log.Fatal(http.ListenAndServe(":4010", r))
	fmt.Println("Listening at port 4000 ...")
}
