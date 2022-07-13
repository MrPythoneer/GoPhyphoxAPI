package phyphox

import (
	"fmt"
	"testing"
	"time"
)

func TestVSensor(t *testing.T) {
	client, err := PhyphoxConnect("192.168.193.215:8080")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Stop()

	lightSensor, ok := client.RegisterSensor(LIGHT).(VSensor)
	if !ok {
		t.Fatal("There is no such a sensor")
	}

	_, err = client.Start()
	if err != nil {
		t.Fatal(err)
	}

	val, ok := lightSensor.Value()
	if !ok {
		t.Fatal("Could not receive value")
	}

	fmt.Println("LIGHT: ", val)
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

	valX, ok := magSensor.GetX()
	if !ok {
		t.Fatal("Could not receive X")
	}
	fmt.Println("X: ", valX)

	_, ok = magSensor.GetY()
	if !ok {
		fmt.Println("Y cannot be received. Correct")
	}

	valZ, ok := magSensor.GetZ()
	if !ok {
		t.Fatal("Could not receive Z")
	}
	fmt.Println("Z: ", valZ)
}

func TestStartStop(t *testing.T) {
	client, err := PhyphoxConnect("192.168.193.215:8080")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		stopped, err := client.Stop()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("Stopped: ", stopped)
	}()

	started, err := client.Start()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Started: ", started)

	time.Sleep(time.Second * 5)
}
