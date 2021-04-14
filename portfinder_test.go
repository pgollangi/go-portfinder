package portfinder

import (
	"net"
	"testing"
)

func TestGetPort(t *testing.T) {
	l, err := net.Listen("tcp", ":3000")
	defer l.Close()
	if err != nil {
		t.Fatal(err)
	}
	port, err := GetPort(PortFinderOptions{
		StartPort: 3000,
		StopPort:  3005,
	})
	if port != 3001 {
		t.Fail()
	}

}
