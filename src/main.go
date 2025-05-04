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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Welcome to Glofox!"))
		if err != nil {
			return
		}
	})

	http.HandleFunc("/classes", controllers.HandleClasses)
	http.HandleFunc("/bookings", controllers.HandleBookings)

	log.Println("Server running at", serverURL)
	log.Fatal(http.ListenAndServe(port, nil))
}
