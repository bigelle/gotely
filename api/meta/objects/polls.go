package objects

type Poll struct {
	Id                   string
	Question             string
	Options              []PollOption
	TotalVoterCount      int
	IsClosed             bool
	IsAnonymous          bool
	Type                 string
	AllowMultipleAnswers bool
	CorrectOptionId      int
	OpenPeriod           int
	CloseDate            int
	Explanation          string
	ExplanationEntities  []MessageEntity
	QuestionEntities     []MessageEntity
}

type PollOption struct {
	Text         string
	VoterCount   int
	TextEntities []MessageEntity
}

//TODO: PollAnswer, InputPollOption
