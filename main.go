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

	host := os.Getenv("HOST")
	validateEnv(host, "HOST")

	env := os.Getenv("ENV")
	validateEnv(env, "ENV")

	pickupCountry := os.Getenv("PICKUP_COUNTRY")
	validateEnv(pickupCountry, "PICKUP_COUNTRY")

	pickupZip := os.Getenv("PICKUP_ZIP")
	validateEnv(pickupZip, "PICKUP_ZIP")

	pickupType := os.Getenv("PICKUP_TYPE")
	validateEnv(pickupType, "PICKUP_TYPE")

	log.Println("envs loaded")

	handler := web.CreateMux(apiKey, host, env, pickupCountry, pickupZip, pickupType)
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
