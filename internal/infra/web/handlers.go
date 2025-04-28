package web

import (
	"encoding/json"
	"net/http"

	"github.com/gblcarvalho/go-expert-lab-cloud-run/internal/gateways"
	"github.com/gblcarvalho/go-expert-lab-cloud-run/internal/usecases"
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

/*
Em caso de sucesso:
	Código HTTP: 200
	Response Body: { "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }
Em caso de falha, caso o CEP não seja válido (com formato correto):
	Código HTTP: 422
	Mensagem: invalid zipcode
Em caso de falha, caso o CEP não seja encontrado:
	Código HTTP: 404
	Mensagem: can not find zipcode
*/
func (h *WeatherHandler) Get(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	getWeatherUC := usecases.NewGetWeatherUseCase(h.cepGateway, h.weatherGateway)
	output, err := getWeatherUC.Execute(cep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
