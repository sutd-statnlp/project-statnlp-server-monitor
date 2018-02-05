package service

import (
	"github.com/shirou/gopsutil/host"
)

// HostService .
type HostService struct {
}

// GetInfo .
func (hostService HostService) GetInfo() (*host.InfoStat, error) {
	return host.Info()
}

// GetTemperature .
func (hostService HostService) GetTemperature() ([]host.TemperatureStat, error) {
	return host.SensorsTemperatures()
}
