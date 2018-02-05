package service

import (
	"github.com/shirou/gopsutil/load"
)

// LoadService .
type LoadService struct {
}

// GetAverage .
func (loadService LoadService) GetAverage() (*load.AvgStat, error) {
	return load.Avg()
}

// GetMisc .
func (loadService LoadService) GetMisc() (*load.MiscStat, error) {
	return load.Misc()
}
