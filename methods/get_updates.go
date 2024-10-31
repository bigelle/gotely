package methods

import (
	"encoding/json"
	"fmt"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/types"
)

type GetUpdates struct {
	Offset         *int      `json:"offset,omitempty"`
	Limit          *int      `json:"limit,omitempty"`
	Timeout        *int      `json:"timeout,omitempty"`
	AllowedUpdates *[]string `json:"allowed_updates,omitempty"`
}

func (g *GetUpdates) SetOffset(o int) *GetUpdates {
	g.Offset = &o
	return g
}

func (g *GetUpdates) SetLimit(l int) *GetUpdates {
	g.Limit = &l
	return g
}

func (g *GetUpdates) SetTimeout(t int) *GetUpdates {
	g.Timeout = &t
	return g
}

func (g *GetUpdates) SetAllowedUpdates(s []string) *GetUpdates {
	g.AllowedUpdates = &s
	return g
}

func (g GetUpdates) Execute() ([]types.Update, error) {
	bot := telego.GetBot()

	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}

	b, err := bot.MakeGetRequest("getUpdates", data)
	if err != nil {
		return nil, err
	}

	var resp types.ApiResponse[[]types.Update]
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}
	if !resp.Ok{
		return nil, fmt.Errorf("%d: %s", resp.ErrorCode, *resp.Description)
	}
	return resp.Result, nil
}
