package service

import (
	"github.com/shirou/gopsutil/cpu"
)

// CPUService .
type CPUService struct {
}

// GetSummaryPercent .
func (cpuService CPUService) GetSummaryPercent() ([]float64, error) {
	return cpu.Percent(0, false)
}

// GetCount .
func (cpuService CPUService) GetCount() (int, error) {
	return cpu.Counts(true)
}

// GetSummaryTime .
func (cpuService CPUService) GetSummaryTime() ([]cpu.TimesStat, error) {
	return cpu.Times(false)
}

// GetInfo .
func (cpuService CPUService) GetInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

// GetPercent .
func (cpuService CPUService) GetPercent() ([]float64, error) {
	return cpu.Percent(0, true)
}

// GetTime .
func (cpuService CPUService) GetTime() ([]cpu.TimesStat, error) {
	return cpu.Times(true)
}
