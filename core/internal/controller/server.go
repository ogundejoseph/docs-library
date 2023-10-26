package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/ogundejoseph/docs-library/core/internal/controller/docs"
)

var router *mux.Router

func initHandlers() {
	router.HandleFunc("/api/docs", controller.GetAllDocs).Methods("GET")
	router.HandleFunc("/api/docs/{id}", controller.GetDoc).Methods("GET")
	router.HandleFunc("/api/docs/new", controller.CreateDoc).Methods("POST")
	router.HandleFunc("/api/docs/update", controller.UpdateDoc).Methods("PUT")
	router.HandleFunc("/api/docs/delete/{id}", controller.DeleteDoc).Methods("DELETE")
}

func Start() {
	router = mux.NewRouter()

	initHandlers()
	fmt.Printf("router initialize and listening on 3020\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}