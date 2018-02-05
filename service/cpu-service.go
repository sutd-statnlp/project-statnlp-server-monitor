package service

import (
	"github.com/shirou/gopsutil/cpu"
)

// CPUService .
type CPUService struct {
}

// GetInfo .
func (cpuService CPUService) GetInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

// GetPercent .
func (cpuService CPUService) GetPercent() ([]float64, error) {
	return cpu.Percent(0, true)
}
