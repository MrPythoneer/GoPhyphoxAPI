package phyphox

import (
	"errors"
	"fmt"
)

// JSON-formatted buffer cannot be parsed correctly
var ErrBufferParse = errors.New("cannot parse the buffer correctly")
var ErrBufferVarNotExist = errors.New("buffer does not contain this variable")

// The given sensor type does not match the type of the sensor returned by RegisterXYZSensor or RegisterVSensor
type ErrSensorWrongType struct {
	t        SensorType
	expected string
}

// The sensor type does not match any known
type ErrSensorUnknown struct {
	t SensorType
}

// The experiment does not use such a sensor
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
