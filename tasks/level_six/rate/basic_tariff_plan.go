package rate

const (
	pricePerKm     float64 = 10.0
	pricePerMinute float64 = 2.0
)

type TripParameters struct {
	Distance float64
	Duration float64
}

func CalculateBasePrice(p TripParameters) float64 {
	return p.Distance*pricePerKm + p.Duration*pricePerMinute
}
