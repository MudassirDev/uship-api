package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MudassirDev/uship-api/internel/web"
	"github.com/joho/godotenv"
)

var (
	PORT    string
	HANDLER http.Handler
)

func init() {
	godotenv.Load()

	log.Println("loading envs")

	port := os.Getenv("PORT")
	validateEnv(port, "PORT")
	PORT = port

	apiKey := os.Getenv("API_KEY")
	validateEnv(apiKey, "API_KEY")

	log.Println("envs loaded")

	handler := web.CreateMux(apiKey)
	HANDLER = handler
}

func main() {
	srv := http.Server{
		Addr:    ":" + PORT,
		Handler: HANDLER,
	}

	log.Println("server is listening at port", PORT)
	log.Fatal(srv.ListenAndServe())
}

func validateEnv(env, name string) {
	if env == "" {
		log.Fatalf("env variable not initialized: %v", name)
	}
}
