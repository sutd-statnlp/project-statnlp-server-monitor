package service

import (
	"os/exec"

	"../util"
)

// GPUService .
type GPUService struct {
}

// GetInfo .
func (gpuService GPUService) GetInfo() string {
	output, err := exec.Command("nvidia-smi -q -x").Output()
	if err != nil {
		return "{}"
	}
	return util.ConvertXMLToJSON(string(output))
}
