package assertions

import (
	"fmt"
)

type CanEmpty interface {
	string | map[string]string | []string | []struct{}
}

type ErrorEmptyParam string

func (e ErrorEmptyParam) Error() string {
	return fmt.Sprintf("%s parameter can't be empty", string(e))
}

func ParamNotEmpty[T CanEmpty](val T, par string) error {
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
		return ErrorEmptyParam(par)
	}
	return nil
}
