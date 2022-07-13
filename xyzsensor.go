package phyphox

type XYZSensor struct {
	prefix  string
	phyphox *Phyphox
}

func (s *XYZSensor) GetX() (float64, error) {
	return s.Get("X")
}

func (s *XYZSensor) GetY() (float64, error) {
	return s.Get("Y")
}

func (s *XYZSensor) GetZ() (float64, error) {
	return s.Get("Z")
}

func (s *XYZSensor) Get(axis string) (float64, error) {
	value, ok := s.phyphox.sensorsData[s.prefix+axis]
	if !ok {
		return 0, ErrBufferVarNotExist
	}

	return value, nil
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

func (s *XYZSensor) IncludeTime() {
	s.phyphox.query += s.prefix + "_time&"
}
