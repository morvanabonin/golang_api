package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/morvanabonin/golang_api/app"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Hello World API")

	apiRoute := mux.NewRouter()
	apiRoute.HandleFunc("/pokemons/", app.ListHandler).
		Methods("GET")

	srv := &http.Server{
		Handler:      apiRoute,
		Addr:         "127.0.0.1:5000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		return 
	}
}
