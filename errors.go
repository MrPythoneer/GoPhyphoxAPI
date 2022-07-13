package phyphox

import (
	"errors"
	"fmt"
)

var ErrBufferParse = errors.New("cannot parse the buffer correctly")
var ErrBufferVarNotExist = errors.New("buffer does not contain this variable")

type ErrSensorWrongType struct {
	t        SensorType
	expected string
}

type ErrSensorUnknown struct {
	t SensorType
}

type ErrSensorNotUsed struct {
	t SensorType
}

func (e *ErrSensorWrongType) Error() string {
	return fmt.Sprintf("wrong type of sensors %s: expected %s", e.t, e.expected)
}

func (e *ErrSensorUnknown) Error() string {
	return fmt.Sprintf("unknown type of sensor (%s)", e.t)
}

func (e *ErrSensorNotUsed) Error() string {
	return fmt.Sprintf("the experiment does not use sensor %s", e.t)
}
