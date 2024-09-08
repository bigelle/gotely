package objects

type ProximityAlertTriggered struct {
	Traveler User
	Watcher  User
	Distance int
}

type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int
}

