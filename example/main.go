package main

import "github.com/MrPythoneer/phyphox"

func main() {

	a, _ := phyphox.PhyphoxConnect("192.168.294.123:8080")
	lights := a.RegisterSensor(phyphox.LIGHT)

	client.start()
	lights.value()
	client.stop()
}
