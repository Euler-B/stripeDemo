package handlers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Hola soy un servidor de prueba de concepto, escrito Golang, y estoy usando la API de Stripe 🫡 😛 🤟")
}
