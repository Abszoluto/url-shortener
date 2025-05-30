package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler).Methods("GET")
    r.HandleFunc("/", ShortenHandler).Methods("POST")
    r.HandleFunc("/{short}", RedirectHandler).Methods("GET")
    r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
    return r
}