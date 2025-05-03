package main

import (
	"fmt"
	"github.com/Vidyuallatha/glofox/src/controllers"
	"log"
	"net/http"
)

func main() {
	port := ":9000"
	serverURL := "http://localhost" + port
	fmt.Println("==============================GLOFOX==============================")
	fmt.Println("Server listening on", serverURL)

	http.HandleFunc("/classes", controllers.HandleClasses)
	http.HandleFunc("/bookings", controllers.HandleBookings)

	log.Println("Server running at", serverURL)
	log.Fatal(http.ListenAndServe(port, nil))
}
