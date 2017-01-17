package main

import (
	"fmt"
	"github.com/eugenebogun/music-storage/controllers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()

	router.POST("/auth/registration", controllers.HandleRegistration)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
	fmt.Println("Server is started.")
}
