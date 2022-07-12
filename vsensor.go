package phyphox

type VSensor struct {
	prefix  string
	phyphox *Phyphox
}

func (s *VSensor) Value() (float64, error) {
	valueb, ok := s.phyphox.buffer[s.prefix].(map[string]any)
	if !ok {
		return 0, ErrBufferParse
	}

	values, ok := valueb["buffer"].([]any)
	if !ok {
		return 0, ErrBufferParse
	}

	value, ok := values[0].(float64)
	if !ok {
		return 0, ErrBufferParse
	}

	return value, nil
}
