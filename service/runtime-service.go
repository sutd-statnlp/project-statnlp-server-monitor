package service

import (
	"runtime"
)

// RuntimeService .
type RuntimeService struct {
}

// GetGOOS .
func (runtimeService RuntimeService) GetGOOS() string {
	return runtime.GOOS
}
