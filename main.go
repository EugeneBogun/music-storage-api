package main

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "bitbucket.org/ebogun/music-storage/controllers"
)

func main() {
    router := httprouter.New()

    router.POST("/auth/registration", controllers.HandleRegistration)
    err := http.ListenAndServe(":8080", router)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }
}