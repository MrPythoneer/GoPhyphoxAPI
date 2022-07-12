package phyphox

import (
	"encoding/json"
	"io"
	"net"
)

type Phyphox struct {
	conn   net.Conn
	config map[string]any
	buffer map[string]any
	query  string
}

func PhyphoxConnect(address string) (*Phyphox, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	configRaw, err := io.ReadAll(conn)
	if err != nil {
		return nil, err
	}

	var config map[string]any
	err = json.Unmarshal(configRaw, &config)
	if err != nil {
		return nil, err
	}

	phyphox := new(Phyphox)
	phyphox.conn = conn
	phyphox.config = config
	return phyphox, nil
}

func (p *Phyphox) RegisterSensor(sensor SensorType) any {
	found := false
	for _, v := range p.config["inputs"].([]map[string]any) {
		if v["source"].(string) == string(sensor) {
			found = true
		}
	}

	if !found {
		return nil
	}

	prefix := string(sensor)
	switch sensor {
	case ACCELEROMETER, GYROSCOPE, LINEAR_ACCELERATION, MAGNETIC_FIELD:
		return XYZSensor{prefix: prefix, phyphox: p}
	case LIGHT, PROXIMITY:
		p.query = prefix + "&"
		return VSensor{prefix: prefix, phyphox: p}
	}

	return nil
}

func (p *Phyphox) getBuffer() map[string]map[string][]float64 {
	return p.buffer["buffer"].(map[string]map[string][]float64)
}
