package assertions

import (
	"fmt"
	"strings"
)

func IsStringEmpty(str string) bool {
	return str == "" || strings.TrimSpace(str) == ""
}

func IsSliceEmpty[T any](sl []T) bool {
	return len(sl) == 0
}

type ErrorEmptyParam struct {
	Param string
}

func (e ErrorEmptyParam) Error() string {
	return fmt.Sprintf("%s parameter can't be empty", e.Param)
}

func ParamNotEmpty(val, par string) error {
	if IsStringEmpty(val) {
		return ErrorEmptyParam{Param: par}
	}
	return nil
}
