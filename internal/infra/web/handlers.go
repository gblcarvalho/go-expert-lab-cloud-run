package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gblcarvalho/go-expert-lab-cloud-run/internal/gateways"
	"github.com/gblcarvalho/go-expert-lab-cloud-run/internal/usecases"
	"github.com/gblcarvalho/go-expert-lab-cloud-run/internal/utils"
	"github.com/go-chi/chi/v5"
)

type WeatherHandler struct {
	cepGateway     gateways.CEPGatewayInterface
	weatherGateway gateways.WeatherGatewayInterface
}

func NewWeatherHandler(
	cepGateway gateways.CEPGatewayInterface,
	weatherGateway gateways.WeatherGatewayInterface,
) *WeatherHandler {
	return &WeatherHandler{
		cepGateway:     cepGateway,
		weatherGateway: weatherGateway,
	}
}

func (h *WeatherHandler) Get(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	getWeatherUC := usecases.NewGetWeatherUseCase(h.cepGateway, h.weatherGateway)
	output, err := getWeatherUC.Execute(cep)
	if err != nil {
		if errors.Is(err, utils.ErrInvalidCEP) {
			http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		} else if errors.Is(err, utils.ErrCEPNotFound) {
			http.Error(w, "can not find zipcode", http.StatusNotFound)
		} else {
			http.Error(w, "can not find weather", http.StatusServiceUnavailable)
		}
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
