package objects

// placeholder for event
type VideoChatStarted struct {
}

type VideoChatEnded struct {
	Duration int
}

type VideoChatParticipantsInvited struct {
	Users []User
}

type VideoChatScheduled struct {
	StartDate int
}
