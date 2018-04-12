package service

import (
	"os/exec"

	"../util"
)

// GPUService .
type GPUService struct {
}

// GetInfo .
func (gpuService GPUService) GetInfo() (string, error) {
	output, err := exec.Command("nvidia-smi -q -x").Output()
	return util.ConvertXMLToJSON(string(output)), err
}
