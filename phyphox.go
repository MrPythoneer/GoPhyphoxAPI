package phyphox

import (
	"encoding/json"
	"io"
	"net/http"
)

type Phyphox struct {
	address     string
	query       string
	Sensors     []string
	SensorsData map[string]float64
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

	sensors := make([]string, 0)
	for _, v := range config["inputs"].([]any) {
		sensor := v.(map[string]any)["source"]
		sensors = append(sensors, sensor.(string))
	}

	phyphox := new(Phyphox)
	phyphox.address = address
	phyphox.Sensors = sensors
	return phyphox, nil
}

func (p *Phyphox) RegisterSensor(sensor SensorType) (any, bool) {
	if !p.HasSensor(sensor) {
		return nil, false
	}

	prefix := sensor.Prefix()
	switch sensor {
	case ACCELEROMETER, GYROSCOPE, LINEAR_ACCELERATION, MAGNETIC_FIELD:
		return XYZSensor{prefix: prefix, phyphox: p}, true
	case LIGHT, PROXIMITY:
		p.query += prefix + "&"
		return VSensor{prefix: prefix, phyphox: p}, true
	}

	return nil, false
}

func (p *Phyphox) HasSensor(sensor SensorType) bool {
	for _, v := range p.Sensors {
		if v == string(sensor) {
			return true
		}
	}

	return false
}

func (p *Phyphox) Update() error {
	res, err := p.execute("/get?" + p.query)

	buffer, ok := res["buffer"].(map[string]any)
	if !ok {
		return ErrBufferParse
	}

	data := make(map[string]float64, 0)
	for k, v := range buffer {
		variable, ok := v.(map[string]any)
		if !ok {
			return ErrBufferParse
		}

		varbuff, ok := variable["buffer"].([]any)
		if !ok {
			return ErrBufferParse
		}

		value, ok := varbuff[0].(float64)
		if !ok {
			return ErrBufferParse
		}

		data[k] = value
	}

	p.SensorsData = data

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
