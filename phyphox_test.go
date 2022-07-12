package phyphox

import (
	"fmt"
	"testing"
)

func TestVSensor(t *testing.T) {
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

func TestXYZSensor(t *testing.T) {
	client, err := PhyphoxConnect("192.168.193.215:8080")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Stop()

	magSensor, ok := client.RegisterSensor(MAGNETIC_FIELD).(XYZSensor)
	if !ok {
		t.Fatalf("The sensor is not a XYZSensor.")
	}
	magSensor.IncludeX()
	magSensor.IncludeZ()

	_, err = client.Start()
	if err != nil {
		t.Fatal(err)
	}

	valX, err := magSensor.GetX()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("X: ", valX)

	_, err = magSensor.GetY()
	if err != nil {
		fmt.Println("Y cannot be received. Correct")
	}

	valZ, err := magSensor.GetZ()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Z: ", valZ)
}
