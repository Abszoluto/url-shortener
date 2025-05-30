package main

import (
    "html/template"
    "math/rand"
    "net/http"
    "time"
    "github.com/gorilla/mux"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    tpl.Execute(w, nil)
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
    url := r.FormValue("url")
    short := generateShortURL()
    SaveURL(short, url)

    tpl.Execute(w, map[string]string{"ShortURL": "http://localhost:8080/" + short})
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    short := vars["short"]

    if original, ok := GetURL(short); ok {
        http.Redirect(w, r, original, http.StatusFound)
    } else {
        http.NotFound(w, r)
    }
}

func generateShortURL() string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    rand.Seed(time.Now().UnixNano())
    b := make([]byte, 6)
    for i := range b {
        b[i] = charset[rand.Intn(len(charset))]
    }
    return string(b)
}
