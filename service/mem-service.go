package service

import (
	"github.com/shirou/gopsutil/mem"
)

// MemService .
type MemService struct {
}

// GetVirtualMemory .
func (memService MemService) GetVirtualMemory() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}
