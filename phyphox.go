package phyphox

import (
	"encoding/json"
	"io"
	"net/http"
)

type Phyphox struct {
	address string
	config  map[string]any
	buffer  map[string]any
	query   string
}

func PhyphoxConnect(address string) (*Phyphox, error) {
	address = "http://" + address

	resp, err := http.Get(address)
	if err != nil {
		return nil, err
	}

	configRaw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var config map[string]any
	err = json.Unmarshal(configRaw, &config)
	if err != nil {
		return nil, err
	}

	phyphox := new(Phyphox)
	phyphox.address = address
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

func (p *Phyphox) Update() (bool, error) {
	return p.execute("/?get" + p.query)
}

func (p *Phyphox) Start() (bool, error) {
	return p.execute("/control?cmd=start")
}

func (p *Phyphox) Stop() (bool, error) {
	return p.execute("/control?cmd=stop")
}

func (p *Phyphox) Clear() (bool, error) {
	return p.execute("clear")
}

func (p *Phyphox) execute(command string) (bool, error) {
	resp, err := http.Get(p.address + command)
	if err != nil {
		return false, err
	}

	respRaw, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result map[string]bool
	json.Unmarshal(respRaw, &result)

	return result["result"], nil
}

func (p *Phyphox) getBuffer() map[string]map[string][]float64 {
	return p.buffer["buffer"].(map[string]map[string][]float64)
}
