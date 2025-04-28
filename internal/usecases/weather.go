package usecases

import "github.com/gblcarvalho/go-expert-lab-cloud-run/internal/gateways"

type GetWeatherUseCase struct {
	cepGateway     gateways.CEPGatewayInterface
	weatherGateway gateways.WeatherGatewayInterface
}

type GetWeatherUseCaseOutput struct {
	Celsius    string `json:"temp_C,omitempty"`
	Fahrenheit string `json:"temp_F,omitempty"`
	Kelvin     string `json:"temp_K,omitempty"`
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

func (uc *GetWeatherUseCase) Execute(cep string) (GetWeatherUseCaseOutput, error) {
	location, err := uc.cepGateway.GetLocation(cep)
	if err != nil {
		return GetWeatherUseCaseOutput{}, err
	}
	weather, err := uc.weatherGateway.GetWeather(location.Latitude, location.Longitude)
	if err != nil {
		return GetWeatherUseCaseOutput{}, err
	}

	return GetWeatherUseCaseOutput{
		Celsius: weather.Celsius,
		Fahrenheit: weather.Fahrenheit,
		Kelvin: FahrenheitToKelvin(weather.Fahrenheit),
	}, nil
}

func FahrenheitToKelvin(tempF string) string {
	return "0"
}
