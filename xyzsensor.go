package phyphox

// Represents a sensor with three variables:
// X, Y, and Z
type XYZSensor struct {
	prefix  string
	phyphox *Phyphox
}

// Returns the X value of the sensor from phyphox SensorsData
func (s *XYZSensor) GetX() (float64, bool) {
	return s.Get("X")
}

// Returns the Y value of the sensor from phyphox SensorsData
func (s *XYZSensor) GetY() (float64, bool) {
	return s.Get("Y")
}

// Returns the Z value of the sensor from phyphox SensorsData
func (s *XYZSensor) GetZ() (float64, bool) {
	return s.Get("Z")
}

// Returns the value of the sensor from phyphox SensorsData
func (s *XYZSensor) Get(axis string) (float64, bool) {
	value, ok := s.phyphox.SensorsData[s.prefix+axis]
	if !ok {
		return 0, false
	}

	return value, true
}

// Next Update() calls will fetch
// the X value form the sensor
func (s *XYZSensor) IncludeX() {
	s.phyphox.query += s.prefix + "X&"
}

// Next Update() calls will fetch
// the Y value form the sensor
func (s *XYZSensor) IncludeY() {
	s.phyphox.query += s.prefix + "Y&"
}

// Next Update() calls will fetch
// the Z value form the sensor
func (s *XYZSensor) IncludeZ() {
	s.phyphox.query += s.prefix + "Z&"
}

// Next Update() calls will fetch all the data
// from the sensor, except for time
func (s *XYZSensor) IncludeAll() {
	s.phyphox.query += s.prefix + "X&" +
		s.prefix + "Y&" +
		s.prefix + "Z&"
}

// Next Update() calls will fetch the sensor time
func (s *XYZSensor) IncludeTime() {
	s.phyphox.query += s.prefix + "_time&"
}
