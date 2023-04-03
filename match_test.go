package logknife

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexIPv4(t *testing.T) {
	tcs := []struct {
		in    string
		match bool
	}{
		{"192", false},
		{"192.1", false},
		{"192.0.2.1", true},
		{"2001:db8:85a3:0:0:8a2e:370:7334", false},
	}

	for _, tc := range tcs {
		assert.Equalf(t, tc.match, IPv4.MatchString(tc.in), "IPv4.MatchString(%s)", tc.in)
	}
}

func TestRegexIPv6(t *testing.T) {
	tcs := []struct {
		in    string
		match bool
	}{
		{"192", false},
		{"192.1", false},
		{"192.0.2.1", false},
		{"2001:db8:85a3:0:0:8a2e:370:7334", true},
		{"2001::", true},
		{"2001::1", true},
	}

	for _, tc := range tcs {
		assert.Equalf(t, tc.match, IPv6.MatchString(tc.in), "IPv6.MatchString(%s)", tc.in)
	}
}

func TestRegexUUID(t *testing.T) {
	tcs := []struct {
		in    string
		match bool
	}{
		{"192", false},
		{"foo", false},
		{"foo-foo-foo", false},
		{"d08a1ec8-6b1e-4cb9-b9b7-7e3fe946f781", true},
	}

	for _, tc := range tcs {
		assert.Equalf(t, tc.match, UUID.MatchString(tc.in), "UUID.MatchString(%s)", tc.in)
	}
}
