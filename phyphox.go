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

	resp, err := http.Get(address + "/config?input")
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
	for _, v := range p.config["inputs"].([]any) {
		source := v.(map[string]any)["source"]
		if source.(string) == string(sensor) {
			found = true
		}
	}

	if !found {
		return nil
	}

	switch sensor {
	case ACCELEROMETER:
		return XYZSensor{prefix: "acc", phyphox: p}

	case GYROSCOPE:
		return XYZSensor{prefix: "gyr", phyphox: p}

	case LINEAR_ACCELERATION:
		return XYZSensor{prefix: "lin", phyphox: p}

	case MAGNETIC_FIELD:
		return XYZSensor{prefix: "mag", phyphox: p}

	case LIGHT:
		p.query += "illum&"
		return VSensor{prefix: "illum", phyphox: p}

	case PROXIMITY:
		p.query += "prox&"
		return VSensor{prefix: "prox", phyphox: p}
	}

	return nil
}

func (p *Phyphox) Update() error {
	res, err := p.execute("/get?" + p.query)

	buffer, ok := res["buffer"].(map[string]any)
	if !ok {
		return ErrBufferParse
	}

	p.buffer = buffer

	return err
}

func (p *Phyphox) Start() (bool, error) {
	res, err := p.execute("/control?cmd=start")
	p.Update()
	return res["result"].(bool), err
}

func (p *Phyphox) Stop() (bool, error) {
	res, err := p.execute("/control?cmd=stop")
	return res["result"].(bool), err
}

func (p *Phyphox) Clear() (bool, error) {
	res, err := p.execute("/control?cmd=clear")
	return res["result"].(bool), err
}

func (p *Phyphox) execute(command string) (map[string]any, error) {
	resp, err := http.Get(p.address + command)
	if err != nil {
		return nil, err
	}

	respRaw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	err = json.Unmarshal(respRaw, &result)

	return result, err
}
