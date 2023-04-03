package logknife

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/google/uuid"
)

var (
	ips   = map[string]string{}
	uuids = map[string]string{}
)

func randIPv4(rng *rand.Rand) string {
	return fmt.Sprintf("10.%d.%d.%d", rng.Intn(255), rng.Intn(255), rng.Intn(255))
}

func randHex(length int, rng *rand.Rand) string {
	var newHex string
	for i := 0; i < length; i++ {
		newHex += fmt.Sprintf("%x", rng.Intn(16))
	}
	return newHex
}

func randIPv6(rng *rand.Rand) string {
	return fmt.Sprintf("2001:db8:%s:%s::", randHex(4, rng), randHex(4, rng))
}

// ReplaceIP takes the true IP and returns a persistent mapping to an internal IP
func ReplaceIP(ip string, rng *rand.Rand) string {
	if _, ok := ips[ip]; !ok {
		if strings.Contains(ip, ".") { // If IPv4
			rIP := randIPv4(rng)
			for {
				if _, ok := ips[rIP]; !ok {
					ips[ip] = rIP
					break
				}
				rIP = randIPv4(rng)
			}
		} else { // If IPv6
			rIP := randIPv6(rng)
			for {
				if _, ok := ips[rIP]; !ok {
					ips[ip] = rIP
					break
				}
				rIP = randIPv6(rng)
			}
		}
	}
	return ips[ip]
}

// ReplaceUUID takes the true UUID and returns a persistent mapping to an internal UUID
func ReplaceUUID(u string) string {
	if _, ok := uuids[u]; !ok {
		uuids[u] = uuid.New().String()
	}
	return uuids[u]
}
