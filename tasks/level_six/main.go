package main

import (
	"fmt"
	"project/tasks/level_six/calculator"
	"project/tasks/level_six/rate"
	"project/tasks/level_six/traffic"
	"project/tasks/level_six/weather"
	"time"
)

func main() {
	calculator := calculator.MyPriceCalculator{
		TrafficClient: &traffic.RealTrafficClient{}, // В продакшене используется настоящий клиент, а мы подключим структуру-заглушку для имитации его работы
	}

	price := calculator.CalculatePrice(
		rate.TripParameters{Distance: 8.5, Duration: 20},
		time.Now(),
		weather.WeatherData{Condition: weather.HeavyRain, WindSpeed: 10},
		55.751244, 37.618423,
	)

	fmt.Printf("Ваша цена: %.0f руб.\n", price)
}
