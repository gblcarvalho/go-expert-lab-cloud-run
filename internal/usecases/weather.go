package usecases

import (
	"github.com/gblcarvalho/go-expert-lab-cloud-run/internal/gateways"
	"github.com/gblcarvalho/go-expert-lab-cloud-run/internal/utils"
	"github.com/gblcarvalho/go-expert-lab-cloud-run/internal/valueobjects"
)

type GetWeatherUseCase struct {
	cepGateway     gateways.CEPGatewayInterface
	weatherGateway gateways.WeatherGatewayInterface
}

type GetWeatherUseCaseOutput struct {
	Celsius    float64 `json:"temp_C,omitempty"`
	Fahrenheit float64 `json:"temp_F,omitempty"`
	Kelvin     float64 `json:"temp_K,omitempty"`
}

func NewGetWeatherUseCase(
	cepGateway gateways.CEPGatewayInterface,
	weatherGateway gateways.WeatherGatewayInterface,
) *GetWeatherUseCase {
	return &GetWeatherUseCase{
		cepGateway:     cepGateway,
		weatherGateway: weatherGateway,
	}
}

func (uc *GetWeatherUseCase) Execute(cepStr string) (GetWeatherUseCaseOutput, error) {
	cep, err := valueobjects.NewCEP(cepStr)
	if err != nil {
		return GetWeatherUseCaseOutput{}, err
	}

	location, err := uc.cepGateway.GetLocation(cep.Value())
	if err != nil {
		return GetWeatherUseCaseOutput{}, utils.ErrCEPNotFound
	}
	weather, err := uc.weatherGateway.GetWeather(location.Locality)
	if err != nil {
		return GetWeatherUseCaseOutput{}, utils.ErrWeather
	}

	return GetWeatherUseCaseOutput{
		Celsius: weather.Celsius,
		Fahrenheit: CelsiusToFahrenheit(weather.Celsius),
		Kelvin: CelsiusToKelvin(weather.Celsius),
	}, nil
}

func CelsiusToFahrenheit(tempC float64) float64 {
	return tempC * 1.8 + 32
}

func CelsiusToKelvin(tempC float64) float64 {
	return tempC + 273
}
