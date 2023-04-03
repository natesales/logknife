package logknife

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// Knife reads from a reader and prints to stdout
func Knife(r io.Reader, redact, ips, uuids bool, redactionPattern string) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	reader := bufio.NewReader(r)

	// Loop over stdin lines
	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}

		modified := string(input)

		if ips {
			for _, ip := range append(IPv4.FindAllString(modified, -1), IPv6.FindAllString(modified, -1)...) {
				log.Debugf("Found IP: %s", ip)
				if redact {
					modified = strings.ReplaceAll(modified, ip, redactionPattern)
				} else {
					modified = strings.ReplaceAll(modified, ip, ReplaceIP(ip, rng))
				}
			}
		}

		if uuids {
			for _, uuid := range UUID.FindAllString(modified, -1) {
				log.Debugf("Found UUID: %s", uuid)
				if redact {
					modified = strings.ReplaceAll(modified, uuid, redactionPattern)
				} else {
					modified = strings.ReplaceAll(modified, uuid, ReplaceUUID(uuid))
				}
			}
		}

		fmt.Println(modified)
	}
}
