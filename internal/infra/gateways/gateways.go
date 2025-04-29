package gateways

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gblcarvalho/go-expert-lab-cloud-run/internal/gateways"
)

type ViaCEPGateway struct {}

const ViaCEPURL = "https://viacep.com.br/ws/%s/json/"
const WeatherAPIURL = "https://api.weatherapi.com/v1/current.json?q=%s&key=%s"

type ViaCEPResp struct {
	Localidade string `json:"localidade,omitempty"`
	Erro       string `json:"erro,omitempty"`
}

type WeatherAPIResp struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func NewViaCEPGateway() *ViaCEPGateway {
	return &ViaCEPGateway{}
}

func (g *ViaCEPGateway) GetLocation(cep string) (gateways.CEPLocation, error) {
	url := fmt.Sprintf(ViaCEPURL, cep)
	resp, err := http.Get(url)
	if err != nil {
		return gateways.CEPLocation{}, fmt.Errorf("erro ao fazer requisição: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return gateways.CEPLocation{}, fmt.Errorf("requisição falhou com status: %s", resp.Status)
	}

	var result ViaCEPResp
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return gateways.CEPLocation{}, fmt.Errorf("erro ao decodificar resposta: %w", err)
	}

	if result.Erro != "" {
		return gateways.CEPLocation{}, fmt.Errorf("requisição falhou")
	}

	return gateways.CEPLocation{
		Locality: result.Localidade,
	}, nil
}

type WeatherAPIGateway struct {
	apiKey string
}

func NewWeatherAPIGateway(apiKey string) *WeatherAPIGateway {
	return &WeatherAPIGateway{apiKey: apiKey}
}

func (w *WeatherAPIGateway)	GetWeather(locality string) (gateways.WeatherTemp, error) {
	escapedLocality := url.QueryEscape(locality)
	url := fmt.Sprintf(WeatherAPIURL, escapedLocality, w.apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return gateways.WeatherTemp{}, fmt.Errorf("erro ao fazer requisição: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return gateways.WeatherTemp{}, fmt.Errorf("requisição falhou com status: %s", resp.Status)
	}

	var result WeatherAPIResp
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return gateways.WeatherTemp{}, fmt.Errorf("erro ao decodificar resposta: %w", err)
	}

	return gateways.WeatherTemp{
		Celsius: result.Current.TempC,
	}, nil
}
