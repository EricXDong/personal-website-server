package main

import (
	"fmt"
	"log"
	"net/http"
	"personal-website-server/env"
	"personal-website-server/rest"
)

func main() {
	env := env.GetEnv()

	r := rest.SetupRoutes(env)
	fmt.Println("Server listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
