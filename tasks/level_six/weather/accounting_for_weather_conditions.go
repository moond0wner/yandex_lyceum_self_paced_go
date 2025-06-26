package weather

type WeatherCondition int

const (
	Clear WeatherCondition = iota // Ключевое слово iota присваивает каждой константе числовое значение по порядку (0, 1, 2, 3 и т.д.)
	Rain
	HeavyRain
	Snow
)

type WeatherData struct {
	Condition WeatherCondition
	WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64 {
	defaultRatio := 1.0

	switch weather.Condition {
	case Rain:
		defaultRatio += 0.125
	case HeavyRain:
		defaultRatio += 0.2
	case Snow:
		defaultRatio += 0.15
	}
	if weather.WindSpeed > 15 {
		defaultRatio += 0.1
	}
	return defaultRatio
}
