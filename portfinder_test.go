package portfinder

import (
	"net"
	"testing"
)

func TestValidateStartPort(t *testing.T) {
	_, err := GetPort(&PortFinderOptions{})
	if err == nil {
		t.Fail()
	}
	_, err = GetPort(&PortFinderOptions{StartPort: -1})
	if err == nil {
		t.Fail()
	}
}

func TestValidateStopPort(t *testing.T) {
	_, err := GetPort(&PortFinderOptions{StartPort: 10, StopPort: 9})
	if err == nil {
		t.Fail()
	}
}

func TestGetPort(t *testing.T) {
	l, err := net.Listen("tcp", ":3000")
	defer l.Close()
	if err != nil {
		t.Fatal(err)
	}
	port, err := GetPort(&PortFinderOptions{
		StartPort: 3000,
		StopPort:  3005,
	})
	if port != 3001 {
		t.Fail()
	}
}

func TestIsOpen(t *testing.T) {
	isOpen, err := IsPortOpen(9999)
	if err != nil {
		t.Fatal(err)
	}
	if isOpen != true {
		t.Fail()
	}
	l, err := net.Listen("tcp", ":9999")
	defer l.Close()
	if err != nil {
		t.Fatal(err)
	}
	isOpen, err = IsPortOpen(9999)
	if err != nil {
		t.Fatal(err)
	}
	if isOpen != false {
		t.Fail()
	}
}
