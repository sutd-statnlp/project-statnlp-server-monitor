package service

import (
	"github.com/shirou/gopsutil/net"
)

// NetService .
type NetService struct {
}

// GetInterfaces .
func (netService NetService) GetInterfaces() ([]net.InterfaceStat, error) {
	return net.Interfaces()
}
