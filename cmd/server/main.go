package main

import (
	"net/http"

	"github.com/gblcarvalho/go-expert-lab-cloud-run/internal/infra/gateways"
	"github.com/gblcarvalho/go-expert-lab-cloud-run/internal/infra/web"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	cepGateway := gateways.NewViaCEPGateway()
	weatherGateway := gateways.NewWeatherAPIGateway("")

	handler := web.NewWeatherHandler(cepGateway, weatherGateway)

	r.Get("/weather/{cep}", handler.Get)
	http.ListenAndServe(":8080", r)
}
