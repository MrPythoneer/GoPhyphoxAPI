package phyphox

import "testing"

func TestLightSensor(t *testing.T) {
	client, err := PhyphoxConnect("192.168.294.123:8080")
	if err != nil {
		t.Fatal(err)
	}

	lightSensor := client.RegisterSensor(LIGHT).(VSensor)

	client.Start()
	println(lightSensor.Value())
	client.Stop()
}
