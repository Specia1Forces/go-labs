package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewTaskStore()
	api := NewAPI(store)

	const address = ":8080"
	log.Printf("server started on http://localhost%s/tasks", address)

	if err := http.ListenAndServe(address, api.Routes()); err != nil {
		log.Fatal(err)
	}
}
