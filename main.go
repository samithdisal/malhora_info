package main

import (
	"flag"
	"net/http"

	"log"

	"fmt"

	"github.com/gorilla/mux"
	"github.com/samithdisal/malhora_info/api"
)

const apiRoutesPrefix = "/api/v1"

// APIRouter handles API reqursts
var APIRouter = mux.NewRouter()

func main() {

	port := flag.Int("port", 8090, "Port to serve http content")

	api.Configure(APIRouter)

	http.Handle(apiRoutesPrefix, APIRouter)

	http.Handle("/", http.FileServer(http.Dir("public")))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))

}
