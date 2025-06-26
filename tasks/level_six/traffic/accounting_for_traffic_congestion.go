package traffic

type TrafficClient interface {
	GetTrafficLevel(lat, lng float64) int // 1–5
}

func GetTrafficMultiplier(trafficLevel int) float64 {
	return 1.0 + float64(trafficLevel-1)*0.1
}

type PriceCalculator struct {
	TrafficClient TrafficClient
}

type RealTrafficClient struct{}

func (c *RealTrafficClient) GetTrafficLevel(lat, lng float64) int {
	return 3 // Константное значение в нашем примере, в реальности оно будет вычисляться сервисом Яндекс Карты
}
