package tools

import (
	"fmt"
	"math"
	"net"
	"net/url"
	"strconv"
)

// MessageHex returns Hexadecimal string of a byte slice passed as a parameter
func MessageHex(b []byte) string {
	var s string
	s += "[ "
	for i := 0; i < len(b); i++ {
		s += fmt.Sprintf("%02x", b[i])
		if i < len(b)-1 {
			s += " "
		}
	}
	s += " ]"

	return s
}

// ProtocolIDString returns string with protocol deacription based on the id
func ProtocolIDString(id uint8) string {
	switch id {
	case 1:
		return "IS-IS Level 1"
	case 2:
		return "IS-IS Level 2"
	case 3:
		return "OSPFv2"
	case 4:
		return "Direct"
	case 5:
		return "Static configuration"
	case 6:
		return "OSPFv3"
	case 7:
		return "BGP"
	default:
		return "Unknown"
	}
}

// AddLevel adds a number of \t defined in level to allign and return the string
func AddLevel(level ...int) string {
	var s string
	if level != nil {
		for i := 0; i < level[0]; i++ {
			s += "\t"
		}
	}
	return s
}

// RawBytesToJSON converts a slice of bytes into a comma separated JSON representation
func RawBytesToJSON(rb []byte) []byte {
	b := []byte{}
	b = append(b, '[')
	for i := 0; i < len(rb); i++ {
		b = append(b, fmt.Sprintf("%d", rb[i])...)
		if i < len(rb)-1 {
			b = append(b, ',')
		}
	}
	b = append(b, ']')

	return b
}

// HostAddrValidator parser host address passed as a string, and make sure it follows X.X.X.X:YYZZ format
func HostAddrValidator(addr string) error {
	host, port, _ := net.SplitHostPort(addr)
	if host == "" || port == "" {
		return fmt.Errorf("host or port cannot be ''")
	}
	// Try to resolve if the hostname was used in the address
	if ip, err := net.LookupIP(host); err != nil || ip == nil {
		// Check if IP address was used in address instead of a host name
		if net.ParseIP(host) == nil {
			return fmt.Errorf("fail to parse host part of address")
		}
	}
	np, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("fail to parse port with error: %w", err)
	}
	if np == 0 || np > math.MaxUint16 {
		return fmt.Errorf("the value of port is invalid")
	}
	return nil
}

// URLAddrValidation validates that passed addrress has valid URL formating.
func URLAddrValidation(addr string) error {
	endpoint, err := url.Parse(addr)
	if err != nil {
		return err
	}
	host, port, _ := net.SplitHostPort(endpoint.Host)
	if host == "" || port == "" {
		return fmt.Errorf("host or port cannot be ''")
	}
	// Try to resolve if the hostname was used in the address
	if ip, err := net.LookupIP(host); err != nil || ip == nil {
		// Check if IP address was used in address instead of a host name
		if net.ParseIP(host) == nil {
			return fmt.Errorf("fail to parse host part of address")
		}
	}
	np, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("fail to parse port with error: %w", err)
	}
	if np == 0 || np > math.MaxUint16 {
		return fmt.Errorf("the value of port is invalid")
	}
	return nil
}