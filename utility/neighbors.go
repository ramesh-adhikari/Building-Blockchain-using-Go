package utility

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func IsFoundHost(guessHost string, port uint16) bool {
	_, err := http.Get("http://" + fmt.Sprintf("%s:%d", guessHost, port))
	if err != nil {
		print(err.Error())
		return false
	}
	return true
}

var PATTERN = regexp.MustCompile(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?\.){3})(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`)

func FindNeighbors(myHost string, myPort uint16, startIp uint8, endIp uint8, startPort uint16, endPort uint16) []string {
	address := fmt.Sprintf("%s:%d", myHost, myPort)

	m := PATTERN.FindStringSubmatch(myHost)
	if m == nil {
		return nil
	}
	prefixHost := m[1]
	lastIp, _ := strconv.Atoi(m[len(m)-1])
	neighbors := make([]string, 0)

	for port := startPort; port < endPort; port++ {
		for ip := startIp; ip < endIp; ip++ {
			guessHost := fmt.Sprintf("%s%d", prefixHost, lastIp+int(ip))
			guessTarget := fmt.Sprintf("%s:%d", guessHost, port)
			if guessTarget != address && IsFoundHost(guessHost, port) {
				neighbors = append(neighbors, guessTarget)
			}
		}

	}
	return neighbors
}

func GetHost() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "127.0.0.1"
	}
	fmt.Println(hostname)
	address, err := net.LookupHost(hostname)

	if err != nil {
		return "127.0.0.1"
	}
	// fmt.Println(address)
	return address[0]
}
