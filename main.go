package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v76"

	"github.com/euler-b/stripe-demo-simple/handlers"
)

func main() {
	// Configurar la clave de API de Stripe
	stripe.Key = os.Getenv("stripeKey")
	
	// Configurar enrutador usando Gorilla Mux
	router := mux.NewRouter()

	// Rutas
	router.HandleFunc("/payment", handlers.MakePayment).Methods("POST")
	router.HandleFunc("/", handlers.Home)

	// Iniciar el servidor
	log.Println("Servidor funcionando en el puerto :8000")
	err := http.ListenAndServe("localhost:8000", router)
	log.Fatal(err)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}
