package phyphox

type SensorType string

const (
	ACCELEROMETER       SensorType = "accelerometer"
	GYROSCOPE           SensorType = "gyroscope"
	LIGHT               SensorType = "light"
	PROXIMITY           SensorType = "proximity"
	LINEAR_ACCELERATION SensorType = "linear_acceleration"
	MAGNETIC_FIELD      SensorType = "magnetic_field"
)

func (st SensorType) Prefix() string {
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
