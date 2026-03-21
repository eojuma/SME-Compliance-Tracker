package main

import (
	"fmt"
	"log"
	"net/http"
	"SME-compliance-tracker/routes"
)

func main() {
	router := routes.SetupRoutes()
	fmt.Println("Server running on http://localhoost:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Server failed: ", err)
	}
}
