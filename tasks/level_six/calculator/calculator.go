package calculator

import (
	"math"
	"project/tasks/level_six/price"
	"project/tasks/level_six/rate"
	"project/tasks/level_six/time_coef"
	"project/tasks/level_six/traffic"
	"project/tasks/level_six/weather"
	"time"
)

type MyPriceCalculator traffic.PriceCalculator

func (c *MyPriceCalculator) CalculatePrice(trip rate.TripParameters, now time.Time, currentWeather weather.WeatherData, lat, lng float64) float64 {
	base := rate.CalculateBasePrice(trip)
	timeMult := time_coef.GetTimeMultiplier(now)
	weatherMult := weather.GetWeatherMultiplier(currentWeather)
	trafficMult := traffic.GetTrafficMultiplier(c.TrafficClient.GetTrafficLevel(lat, lng))

	finalPrice := base * timeMult * weatherMult * trafficMult

	return price.ApplyPriceLimits(math.Round(finalPrice))
}
