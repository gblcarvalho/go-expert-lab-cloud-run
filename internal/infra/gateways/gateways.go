package gateways

import "github.com/gblcarvalho/go-expert-lab-cloud-run/internal/gateways"

type ViaCEPGateway struct {}

func NewViaCEPGateway() *ViaCEPGateway {
	return &ViaCEPGateway{}
}

func (g *ViaCEPGateway) GetLocation(cep string) (gateways.CEPLocation, error) {
	return gateways.CEPLocation{
		Latitude: "",
		Longitude: "",
	}, nil
}

type WeatherAPIGateway struct {
	APIKey string
}

func NewWeatherAPIGateway(apiKey string) *WeatherAPIGateway {
	return &WeatherAPIGateway{APIKey: apiKey}
}

func (w *WeatherAPIGateway)	GetWeather(latitude string, longitude string) (gateways.WeatherTemp, error) {
	return gateways.WeatherTemp{
		Celsius: "0",
		Fahrenheit: "0",
	}, nil
}
