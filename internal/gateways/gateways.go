package gateways

type CEPGatewayInterface interface {
	GetLocation(cep string) (CEPLocation, error)
}

type WeatherGatewayInterface interface {
	GetWeather(latitude string, longitude string) (WeatherTemp, error)
}

type CEPLocation struct {
	Latitude string
	Longitude string
}

type WeatherTemp struct {
	Celsius    string
	Fahrenheit string
}
