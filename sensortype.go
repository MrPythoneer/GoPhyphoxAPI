package phyphox

type SensorType string

const (
	ACCELEROMETER       SensorType = "acc"
	GYROSCOPE           SensorType = "gyr"
	LIGHT               SensorType = "light"
	PROXIMITY           SensorType = "prox"
	LINEAR_ACCELERATION SensorType = "lin"
	MAGNETIC_FIELD      SensorType = "mag"
)
