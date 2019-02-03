package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"personal-website-server/env"
	"personal-website-server/rest"
)

func main() {
	env := env.GetEnv()

	if env.Deployment == "docker" {
		r := rest.SetupRoutes(env)
		port := os.Getenv("PORT")
		fmt.Println("Server listening on port " + port)
		log.Fatal(http.ListenAndServe(":"+port, r))
	} else {
		//	Running on lambda
		handleLambdaRequest(env)
	}
}
