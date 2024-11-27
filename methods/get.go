package methods

import (
	"encoding/json"
	"strings"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/errors"
	"github.com/bigelle/tele.go/internal"
	"github.com/bigelle/tele.go/types"
)

type GetUserProfilePhotos struct {
	UserId int
	Offset *int
	Limit  *int
}

func (g GetUserProfilePhotos) Validate() error {
	if g.UserId < 1 {
		return errors.ErrInvalidParam("user_id parameter can't be empty")
	}
	if g.Limit != nil {
		if *g.Limit < 1 || *g.Limit > 100 {
			return errors.ErrInvalidParam("limit parameter must be between 1 and 100")
		}
	}
	return nil
}

func (g GetUserProfilePhotos) MarshalJSON() ([]byte, error) {
	type alias GetUserProfilePhotos
	return json.Marshal(alias(g))
}

func (g GetUserProfilePhotos) Execute() (*types.UserProfilePhotos, error) {
	return internal.MakeGetRequest[types.UserProfilePhotos](telego.GetToken(), "getUserProfilePhotos", g)
}

type GetFile struct {
	FileId string
}

func (g GetFile) Validate() error {
	if strings.TrimSpace(g.FileId) == "" {
		return errors.ErrInvalidParam("file_id parameter can't be empty")
	}
	return nil
}

func (g GetFile) MarshalJSON() ([]byte, error) {
	type alias GetFile
	return json.Marshal(alias(g))
}
