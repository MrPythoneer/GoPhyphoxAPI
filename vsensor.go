package phyphox

type VSensor struct {
	prefix  string
	phyphox *Phyphox
}

func (s *VSensor) Value() float64 {
	return s.phyphox.getBuffer()[s.prefix]["buffer"][0]
}
