package getip

import (
	"errors"
	"fmt"
	"net"
)

type IPProtocol int

const (
	IPv4 IPProtocol = iota
	IPv6
	Unknown
)

type IPGetter interface {
	GetIP(hostname string, protocol IPProtocol) (string, error)
}

type getterImpl struct{}

func New() IPGetter {
	return &getterImpl{}
}

func (g *getterImpl) GetIP(hostname string, protocol IPProtocol) (string, error) {
	ips, err := net.LookupIP(hostname)
	if err != nil {
		return "", fmt.Errorf("failed to lookup ip; %s; %w", hostname, err)
	}
	switch protocol {
	case IPv4:
		return getIPv4(ips)
	case IPv6:
		return getIPv6(ips)
	}
	return "", errors.New("no address")
}

func getIPv4(ips []net.IP) (string, error) {
	for _, ip := range ips {
		if ip.To4() != nil {
			return ip.String(), nil
		}
	}
	return "", fmt.Errorf("no IPv4 address")
}

func getIPv6(ips []net.IP) (string, error) {
	for _, ip := range ips {
		if ip.To16() != nil {
			return ip.String(), nil
		}
	}
	return "", fmt.Errorf("no IPv6 address")
}
