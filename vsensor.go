package phyphox

type VSensor struct {
	prefix  string
	phyphox *Phyphox
}

func (s *VSensor) Value() (float64, bool) {
	value, ok := s.phyphox.SensorsData[s.prefix]
	if !ok {
		return 0, false
	}
	return value, true
}

func (s *VSensor) IncludeTime() {
	s.phyphox.query += s.prefix + "_time&"
}
