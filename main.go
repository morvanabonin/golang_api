package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello World API")

	route := mux.NewRouter()
	route.HandleFunc("/app", PokemonHandler)
	http.Handle("/", route)
}
