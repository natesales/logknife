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

	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

var version = "dev" // Set by build process

// Flags
var opts struct {
	IPs         bool   `short:"i" long:"ips" description:"Match IP addresses"`
	IPv4        bool   `short:"4" long:"ipv4" description:"Match IPv4 addresses"`
	IPv6        bool   `short:"6" long:"ipv6" description:"Match IPv6 addresses"`
	Action      string `short:"a" long:"action" description:"\"replace\" to replace with dummy values, \"remove\" to replace with asterisks"`
	ShowVersion bool   `short:"V" long:"version" description:"Show version and exit"`
}

var (
	regexIPv4 = `(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`
	regexIPv6 = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
)

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
		ipMap[ip] = newIp.String()[:len(newIp.String())-1]
	}

	return ipMap[ip]
}

func main() {
	// Parse cli flags
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		if !strings.Contains(err.Error(), "Usage") {
			log.Fatal(err)
		}
		os.Exit(1)
	}

	if opts.ShowVersion {
		log.Infof("logknife version %s https://github.com/natesales/logkife", version)
		os.Exit(0)
	}

	// Validate action flag
	if !(opts.Action == "replace" || opts.Action == "remove") {
		log.Fatal("--action must be \"replace\" or \"remove\"")
	}

	// Set filters to apply
	configuredFilters := map[string]bool{}
	if opts.IPv4 {
		configuredFilters[regexIPv4] = true
	}
	if opts.IPv6 {
		configuredFilters[regexIPv6] = true
	}
	if opts.IPs {
		configuredFilters[regexIPv4] = true
		configuredFilters[regexIPv6] = true
	}

	// Open buffered stdin reader
	reader := bufio.NewReader(os.Stdin)

	// Loop over stdin lines
	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}

		// Apply replacements for each defined filter
		modified := string(input)
		for filter := range configuredFilters {
			for _, entry := range regexp.MustCompile(filter).FindAllString(modified, -1) {
				var replacement string
				if opts.Action == "remove" {
					replacement = "******"
				} else if opts.Action == "replace" {
					replacement = getInternalIP(entry)
				} else {
					log.Fatal("Unknown action")
				}
				modified = strings.Replace(modified, entry, replacement, -1)
			}
		}

		fmt.Println(modified)
	}
}
