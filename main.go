package main

import (
    "log"
    "net/http"
)

func main() {
    routes := SetupRouter()
    log.Println("Servidor rodando em http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", routes))
}