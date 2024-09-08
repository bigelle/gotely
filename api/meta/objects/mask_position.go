package objects

import (
	"fmt"
	"reflect"
)

type MaskPosition struct {
	point  string
	xShift float32
	yShift float32
	scale  float32
}

func (m MaskPosition) Validate() error {
	if reflect.DeepEqual(m, MaskPosition{}) {
		return fmt.Errorf("all fields must be non-empty'")
	}
	return nil
}
