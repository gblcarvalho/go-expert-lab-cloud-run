package gateways

type CEPGatewayInterface interface {
	GetLocation(cep string) (CEPLocation, error)
}

type WeatherGatewayInterface interface {
	GetWeather(locality string) (WeatherTemp, error)
}

type CEPLocation struct {
	Locality string
}

type WeatherTemp struct {
	Celsius    float64
}
