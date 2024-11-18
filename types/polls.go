package types

import (
	"fmt"
)

type Poll struct {
	Id                   string           `json:"id"`
	Question             string           `json:"question"`
	Options              []PollOption     `json:"options"`
	TotalVoterCount      int              `json:"total_voter_count"`
	IsClosed             bool             `json:"is_closed"`
	IsAnonymous          bool             `json:"is_anonymous"`
	Type                 string           `json:"type"`
	AllowMultipleAnswers bool             `json:"allow_multiple_answers"`
	QuestionEntities     []MessageEntity  `json:"question_entities"`
	CorrectOptionId      *int             `json:"correct_option_id,omitempty"`
	OpenPeriod           *int             `json:"open_period,omitempty"`
	CloseDate            *int             `json:"close_date,omitempty"`
	Explanation          *string          `json:"explanation,omitempty"`
	ExplanationEntities  *[]MessageEntity `json:"explanation_entities,omitempty"`
}

type PollOption struct {
	Text         string           `json:"text"`
	VoterCount   int              `json:"voter_count"`
	TextEntities *[]MessageEntity `json:"text_entities,omitempty"`
}

type PollAnswer struct {
	PollId    string `json:"poll_id"`
	User      *User  `json:"user"`
	OptionIds []int  `json:"option_ids"`
	VoterChat *Chat  `json:"voter_chat"`
}

type InputPollOption struct {
	Text          string           `json:"text"`
	TextParseMode *string          `json:"text_parse_mode,omitempty"`
	TextEntities  *[]MessageEntity `json:"text_entities,omitempty"`
}

func (i InputPollOption) Validate() error {
	if len(i.Text) < 1 || len(i.Text) > 100 {
		return fmt.Errorf("text must be between 1 and 100 characters")
	}
	if len(*i.TextEntities) != 0 && len(*i.TextParseMode) != 0 {
		return fmt.Errorf("parse mode and entities can't be used together")
	}
	return nil
}
