package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"regexp"
	"strings"
)

var (
	regexIPv4 = `(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`
	regexIPv6 = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
)

// filters maps a string to a regex
var filters = map[string]string{
	"ipv4": regexIPv4,
	"ipv6": regexIPv6,
}

var ipMap = map[string]string{}

// getInternalIP takes the true IP and returns a persistent mapping to an internal address
func getInternalIP(ip string) string {
	if val, ok := ipMap[ip]; ok { // If the IP already has a internal mapped value
		return val
	}

	var newIp net.IP
	if strings.Contains(ip, ".") { // If IPv4
		// Set first octet to 10
		newIp = append(newIp, 10)

		// Append random data for the rest of the address
		for i := 0; i < 3; i++ { // 4 octets - 1 for 10 prefix
			number := uint8(rand.Intn(255))
			newIp = append(newIp, number)
		}
		ipMap[ip] = newIp.String()
	} else { // If IPv6
		// Append 2001:db8 documentation prefix
		newIp = append(newIp, 32)
		newIp = append(newIp, 1)
		newIp = append(newIp, 13)
		newIp = append(newIp, 184)

		// Append random data for the rest of the address
		for i := 0; i < 16-4; i++ { // 16 - 4 for 2001:db8: prefix
			number := uint8(rand.Intn(255))
			newIp = append(newIp, number)
		}
		ipMap[ip] = newIp.String()
	}

	return ipMap[ip]
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}

		// Apply replacements for each defined filter
		modified := string(input)
		for _, f := range []string{"ipv4", "ipv6"} {
			for _, entry := range regexp.MustCompile(filters[f]).FindAllString(modified, -1) {
				modified = strings.Replace(modified, entry, getInternalIP(entry), -1)
			}
		}

		fmt.Println(modified)
	}
}
