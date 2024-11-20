package assertions

import (
	"fmt"
)

type ErrInvalidParam string

func (e ErrInvalidParam) Error() string {
	return string(e)
}

// currently used only for strings and []string
func ParamNotEmpty[T string | []string](val T, par string) error {
	isEmpty := false
	switch v := any(val).(type) {
	case string:
		isEmpty = (v == "")
	case []any:
		isEmpty = (len(v) == 0)
	case map[any]any:
		isEmpty = (len(v) == 0)
	}
	if isEmpty {
		return ErrInvalidParam(fmt.Sprintf("%s parameter can't be empty", par))
	}
	return nil
}
