package phyphox

type VSensor struct {
	prefix  string
	phyphox *Phyphox
}

func (s *VSensor) Value() (float64, error) {
	value, ok := s.phyphox.SensorsData[s.prefix]
	if !ok {
		return 0, ErrBufferVarNotExist
	}
	return value, nil
}

func (s *VSensor) IncludeTime() {
	s.phyphox.query += s.prefix + "_time&"
}
