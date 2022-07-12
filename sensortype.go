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
