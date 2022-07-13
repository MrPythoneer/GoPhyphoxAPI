package phyphox

// Represents a single-value sensor
type VSensor struct {
	prefix  string
	phyphox *Phyphox
}

// Returns the value of the sensor from phyphox SensorsData
func (s *VSensor) Value() (float64, bool) {
	value, ok := s.phyphox.SensorsData[s.prefix]
	if !ok {
		return 0, false
	}
	return value, true
}

// Returns the Time value of the sensor from phyphox SensorsData
func (s *VSensor) Time() (float64, bool) {
	value, ok := s.phyphox.SensorsData[s.prefix+"_time"]
	if !ok {
		return 0, false
	}
	return value, true
}

// Next Update() calls will fetch the sensor time
func (s *VSensor) IncludeTime() {
	s.phyphox.query += s.prefix + "_time&"
}
