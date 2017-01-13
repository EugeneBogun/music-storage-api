package main

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
)

func main() {
    router := httprouter.New()

    err := http.ListenAndServe(":8080", router)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }
}