package phyphox

func main() {
	client, _ := PhyphoxConnect("192.168.294.123:8080")
	lights := client.RegisterSensor(LIGHT)

	client.start()
	lights.value()
	client.stop()
}
