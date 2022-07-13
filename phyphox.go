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

// Connects to the remote experiment at the given address
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

// Returns VSensor representing a sensor in the experiment.
// Since the sensor has only one variable, it's automatically
// going to be fetched with Update()
func (p *Phyphox) RegisterVSensor(sensor SensorType) (*VSensor, error) {
	if !p.HasSensor(sensor) {
		return nil, &ErrSensorNotUsed{sensor}
	}

	switch sensor {
	case ACCELEROMETER, GYROSCOPE, LINEAR_ACCELERATION, MAGNETIC_FIELD:
		return nil, &ErrSensorWrongType{sensor, "VSensor"}
	case LIGHT, PROXIMITY:
		prefix := sensor.prefix()
		p.query += prefix + "&"
		return &VSensor{prefix: prefix, phyphox: p}, nil
	}

	return nil, &ErrSensorUnknown{sensor}
}

// Returns VSensor representing a sensor in the experiment.
// Since the sensor has several variables, none will be fetched
// with Update(). In order to fetch data from the sensor,
// IncludeX, IncludeY or IncludeZ should be called.
func (p *Phyphox) RegisterXYZSensor(sensor SensorType) (*XYZSensor, error) {
	if !p.HasSensor(sensor) {
		return nil, &ErrSensorNotUsed{sensor}
	}

	switch sensor {
	case ACCELEROMETER, GYROSCOPE, LINEAR_ACCELERATION, MAGNETIC_FIELD:
		return &XYZSensor{
			prefix:  sensor.prefix(),
			phyphox: p,
		}, nil
	case LIGHT, PROXIMITY:
		return nil, &ErrSensorWrongType{sensor, "XYZSensor"}
	}

	return nil, &ErrSensorUnknown{sensor}
}

// Returns XYZSensor or VSensor representing a sensor in the experiment.
//
// Since VSensor has only one variable, it's automatically going to be fetched
// with Update()
//
// Since XYZSensor has several variables, none will be fetched with Update().
// In order to fetch data from the sensor, IncludeX, IncludeY or IncludeZ should
// be called.
//
func (p *Phyphox) RegisterSensor(sensor SensorType) (any, error) {
	if !p.HasSensor(sensor) {
		return nil, &ErrSensorNotUsed{sensor}
	}

	prefix := sensor.prefix()
	switch sensor {
	case ACCELEROMETER, GYROSCOPE, LINEAR_ACCELERATION, MAGNETIC_FIELD:
		return &XYZSensor{prefix: prefix, phyphox: p}, nil
	case LIGHT, PROXIMITY:
		p.query += prefix + "&"
		return &VSensor{prefix: prefix, phyphox: p}, nil
	}

	return nil, &ErrSensorUnknown{sensor}
}

// Returns true if the experiment uses the given
// sensor type, otheriwse, false will be returned
func (p *Phyphox) HasSensor(sensor SensorType) bool {
	for _, v := range p.Sensors {
		if v == string(sensor) {
			return true
		}
	}

	return false
}

// Requests the remote host for the latest data.
// The data will be saved to the SensorsData field
func (p *Phyphox) Fetch() error {
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

// Starts measuring.
//
// By default, Fetch() is called automatically
func (p *Phyphox) Start() (bool, error) {
	res, err := p.execute("/control?cmd=start")
	p.Fetch()
	return res["result"].(bool), err
}

// Stops measuring
func (p *Phyphox) Stop() (bool, error) {
	res, err := p.execute("/control?cmd=stop")
	return res["result"].(bool), err
}

// Clears the experiment buffer
func (p *Phyphox) Clear() (bool, error) {
	res, err := p.execute("/control?cmd=clear")
	return res["result"].(bool), err
}

// Executes some remote command on the host.
// Returns the JSON-like result of the command
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
