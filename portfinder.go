package portfinder

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

// PortFinderOptions are the options passed to GetPort
type PortFinderOptions struct {
	//Host to find available port on.
	Host string
	// Minimum port
	StartPort int
	//Maximum port
	StopPort int
	Timeout  time.Duration
}

// IsPortInUse checks if the give port on the host name is inUse.
// Return TRUE if port is in use otherwise FALSE.
func IsPortInUse(host string, port int, timeout time.Duration) bool {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", net.JoinHostPort(host, strconv.Itoa(port)))
	if err != nil {
		return false
	}
	conn, err := net.DialTimeout("tcp", tcpAddr.String(), timeout)
	if err != nil {
		return false
	}

	defer conn.Close()
	return true
}

// GetPort search for open port in the range provided
func GetPort(options PortFinderOptions) (int, error) {
	if options.StartPort <= 0 {
		return -1, fmt.Errorf("Provided options.startPort(%d) is less than 0, which are cannot be bound.", options.StartPort)
	}
	if options.StopPort < options.StartPort {
		return -1, fmt.Errorf("Provided options.stopPort(%d) is less than options.startPort (%d)", options.StopPort, options.StartPort)
	}
	allAddrs, err := localAddresses()
	if err != nil {
		return -1, err
	}
	var openPort = -1
SEARCHPORT:
	for openPort = options.StartPort; openPort <= options.StopPort; openPort++ {
		var wg sync.WaitGroup
		results := make(chan bool, len(allAddrs))
		for _, ip := range allAddrs {
			addr := ip.String()
			wg.Add(1)
			func() {
				defer wg.Done()
				inUse := IsPortInUse(addr, openPort, 3*time.Second)
				results <- inUse
			}()
		}
		wg.Wait()
		close(results)
		for inUse := range results {
			if inUse {
				continue SEARCHPORT
			}
		}
		break
	}
	return openPort, nil
}

func localAddresses() ([]net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("localAddresses: %+v\n", err.Error())
	}
	var addresses []net.IP
	for _, i := range ifaces {
		if i.Flags&net.FlagUp == 0 {
			// Interface is not active
			continue
		}
		addrs, err := i.Addrs()
		if err != nil {
			return nil, fmt.Errorf("localAddresses: %+v\n", err.Error())
		}
		for _, a := range addrs {
			switch v := a.(type) {
			case *net.IPNet:
				if v.IP.To4() != nil {
					addresses = append(addresses, v.IP)
				}
			case *net.IPAddr:
				addresses = append(addresses, v.IP)
			}

		}
	}
	return addresses, nil
}
