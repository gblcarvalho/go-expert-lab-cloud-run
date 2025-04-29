package main

import (
	"net/http"

	"github.com/gblcarvalho/go-expert-lab-cloud-run/configs"
	"github.com/gblcarvalho/go-expert-lab-cloud-run/internal/infra/gateways"
	"github.com/gblcarvalho/go-expert-lab-cloud-run/internal/infra/web"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	cepGateway := gateways.NewViaCEPGateway()
	weatherGateway := gateways.NewWeatherAPIGateway(configs.WeatherApiKey)

	handler := web.NewWeatherHandler(cepGateway, weatherGateway)

	r.Get("/weather/{cep}", handler.Get)
	http.ListenAndServe(":8080", r)
}
