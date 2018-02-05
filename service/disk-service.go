package service

import (
	"github.com/shirou/gopsutil/disk"
)

// DiskService .
type DiskService struct {
}

// GetUsage .
func (diskService DiskService) GetUsage() (*disk.UsageStat, error) {
	return disk.Usage("/")
}
