package phyphox

// Type of the sensor used in the experiment
type SensorType string

const (
	ACCELEROMETER       SensorType = "accelerometer"
	GYROSCOPE           SensorType = "gyroscope"
	LIGHT               SensorType = "light"
	PROXIMITY           SensorType = "proximity"
	LINEAR_ACCELERATION SensorType = "linear_acceleration"
	MAGNETIC_FIELD      SensorType = "magnetic_field"
)

// Variable that denotes name of the sensor in the JSON buffer
func (st SensorType) prefix() string {
	switch st {
	case ACCELEROMETER:
		return "acc"
	case GYROSCOPE:
		return "gyr"
	case LIGHT:
		return "illum"
	case PROXIMITY:
		return "prox"
	case LINEAR_ACCELERATION:
		return "lin"
	case MAGNETIC_FIELD:
		return "mag"
	}

	return ""
}
