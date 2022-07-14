package main

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

func randIPv4() string {
	return fmt.Sprintf("10.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255))
}

func randHex(length int) string {
	var newHex string
	for i := 0; i < length; i++ {
		newHex += fmt.Sprintf("%x", rand.Intn(16))
	}
	return newHex
}

func randIPv6() string {
	return fmt.Sprintf("2001:db8:%s:%s::", randHex(4), randHex(4))
}

// ReplaceIP takes the true IP and returns a persistent mapping to an internal IP
func ReplaceIP(ip string) string {
	if _, ok := ips[ip]; !ok {
		if strings.Contains(ip, ".") { // If IPv4
			rIP := randIPv4()
			for {
				if _, ok := ips[rIP]; !ok {
					ips[ip] = rIP
					break
				}
				rIP = randIPv4()
			}
		} else { // If IPv6
			rIP := randIPv6()
			for {
				if _, ok := ips[rIP]; !ok {
					ips[ip] = rIP
					break
				}
				rIP = randIPv6()
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
