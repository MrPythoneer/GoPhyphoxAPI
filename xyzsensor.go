package phyphox

type XYZSensor struct {
	prefix  string
	phyphox *Phyphox
}

// func (s *XYZSensor) getX() float64 {
// 	return s.phyphox.getBuffer()[s.prefix+"X"]["buffer"][0]
// }

// func (s *XYZSensor) getY() float64 {
// 	return s.phyphox.getBuffer()[s.prefix+"Y"]["buffer"][0]
// }

// func (s *XYZSensor) getZ() float64 {
// 	return s.phyphox.getBuffer()[s.prefix+"Z"]["buffer"][0]
// }

func (s *XYZSensor) IncludeX() {
	s.phyphox.query += s.prefix + "X"
}

func (s *XYZSensor) IncludeY() {
	s.phyphox.query += s.prefix + "Y"
}

func (s *XYZSensor) IncludeZ() {
	s.phyphox.query += s.prefix + "Z"
}

func (s *XYZSensor) IncludeTime() {
	s.phyphox.query += s.prefix + "_time"
}
