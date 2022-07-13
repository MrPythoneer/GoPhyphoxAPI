package phyphox

type XYZSensor struct {
	prefix  string
	phyphox *Phyphox
}

func (s *XYZSensor) GetX() (float64, bool) {
	return s.Get("X")
}

func (s *XYZSensor) GetY() (float64, bool) {
	return s.Get("Y")
}

func (s *XYZSensor) GetZ() (float64, bool) {
	return s.Get("Z")
}

func (s *XYZSensor) Get(axis string) (float64, bool) {
	value, ok := s.phyphox.SensorsData[s.prefix+axis]
	if !ok {
		return 0, false
	}

	return value, true
}

func (s *XYZSensor) IncludeX() {
	s.phyphox.query += s.prefix + "X&"
}

func (s *XYZSensor) IncludeY() {
	s.phyphox.query += s.prefix + "Y&"
}

func (s *XYZSensor) IncludeZ() {
	s.phyphox.query += s.prefix + "Z&"
}

func (s *XYZSensor) IncludeAll() {
	s.phyphox.query += s.prefix + "X&" +
		s.prefix + "Y&" +
		s.prefix + "Z&"
}

func (s *XYZSensor) IncludeTime() {
	s.phyphox.query += s.prefix + "_time&"
}
