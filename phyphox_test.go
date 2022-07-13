package phyphox

import (
	"fmt"
	"testing"
	"time"
)

const HostAddr string = "192.168.193.215:8080"

func TestVSensor(t *testing.T) {
	experiment, err := PhyphoxConnect(HostAddr)
	if err != nil {
		t.Fatal(err)
	}

	lightSensor, err := experiment.RegisterVSensor(LIGHT)
	if err != nil {
		t.Error(err)
	}

	_, err = experiment.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer experiment.Stop()

	val, ok := lightSensor.Value()
	if !ok {
		t.Fatal("Could not receive value")
	}

	fmt.Println("LIGHT: ", val)
}

func TestXYZSensor(t *testing.T) {
	experiment, err := PhyphoxConnect(HostAddr)
	if err != nil {
		t.Fatal(err)
	}

	magSensor, err := experiment.RegisterXYZSensor(MAGNETIC_FIELD)
	if err != nil {
		t.Error(err)
	}

	magSensor.IncludeX()
	magSensor.IncludeZ()

	_, err = experiment.Start()
	if err != nil {
		t.Fatal(err)
	}

	defer experiment.Stop()

	valX, ok := magSensor.X()
	if !ok {
		t.Fatal("Could not receive X")
	}
	fmt.Println("X: ", valX)

	_, ok = magSensor.Y()
	if !ok {
		fmt.Println("Y cannot be received. Correct")
	}

	valZ, ok := magSensor.Z()
	if !ok {
		t.Fatal("Could not receive Z")
	}
	fmt.Println("Z: ", valZ)
}

func TestStartStop(t *testing.T) {
	experiment, err := PhyphoxConnect(HostAddr)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		stopped, err := experiment.Stop()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("Stopped: ", stopped)
	}()

	started, err := experiment.Start()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Started: ", started)

	time.Sleep(time.Second * 5)
}
