package phyphox

import (
	"fmt"
	"testing"
)

func TestLightSensor(t *testing.T) {
	client, err := PhyphoxConnect("192.168.193.215:8080")
	if err != nil {
		t.Fatal(err)
	}

	lightSensor := client.RegisterSensor(LIGHT).(VSensor)

	_, err = client.Start()
	if err != nil {
		t.Fatal(err)
	}

	val, err := lightSensor.Value()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("LIGHT: ", val)

	_, err = client.Stop()
	if err != nil {
		t.Fatal(err)
	}
}
