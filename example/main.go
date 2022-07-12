package main

func main() {
	client := Phyphox.connect("192.168.294.123", "8080")
	lights := client.registerSensor(LIGHT)

	client.start()
	lights.value()
	client.stop()
}
