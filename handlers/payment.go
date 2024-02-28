package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/euler-b/stripe-demo-simple/utils"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/charge"
	"github.com/stripe/stripe-go/v76/customer"
)

type PaymentRequest struct {
	Amount        int    `json:"amount"`
	Currency      string `json:"currency"`
	Description   string `json:"description"`
	Email         string `json:"email"`
	Name          string `jason:"name"`
	PaymentMethod string `json:"payment-method"`
}

func MakePayment(w http.ResponseWriter, r *http.Request) {
	// Decodificar el cuerpo de la solicitud
	var paymentReq PaymentRequest
	err := json.NewDecoder(r.Body).Decode(&paymentReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Crear un cliente en Stripe
	customerParams := &stripe.CustomerParams{
		Email:         stripe.String(paymentReq.Email),
		Name:          stripe.String(paymentReq.Name),
		PaymentMethod: stripe.String(paymentReq.PaymentMethod),
	}
	customer, err := customer.New(customerParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Crear un cargo en Stripe
	chargeParams := &stripe.ChargeParams{
		Amount:      stripe.Int64(int64(paymentReq.Amount)),
		Currency:    stripe.String(paymentReq.Currency),
		Customer:    stripe.String(customer.ID),
		Description: stripe.String(paymentReq.Description),
	}
	_, err = charge.New(chargeParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respuesta exitosa
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Pago exitoso"})
}
