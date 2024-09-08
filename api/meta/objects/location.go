package objects

type Location struct {
	longitude            float64
	latitude             float64
	horizontalAccuracy   float64
	livePeriod           int
	heading              int
	proximityAlertRadius int
}
