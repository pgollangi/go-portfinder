package portfinder

import "fmt"

type PortFinderOptions struct {
	//Host to find available port on.
	Host string
	// Minimum port
	StartPort int
	//Maximum port
	StopPort int
}

func GetPort(options PortFinderOptions) (int, error) {
	if options.StartPort < 0 {
		return -1, fmt.Errorf("Provided options.startPort(%d) is less than 0, which are cannot be bound.", options.StartPort)
	}
	if options.StopPort < options.StartPort {
		return -1, fmt.Errorf("Provided options.stopPort(%d) is less than options.startPort (%d)", options.StopPort, options.StartPort)
	}
	return -1, nil
}
