package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../resource"
	"github.com/stretchr/testify/assert"
)

func TestGetCPUInfo(t *testing.T) {
	router := SetUpRouter()

	cpuResource := resource.CPUResource{}
	cpuResource.InitRoutes(router)

	// Get cpu info
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/cpu/info", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.True(t, true, w.Body.Len() > 0)

}

func TestGetPercent(t *testing.T) {
	router := SetUpRouter()

	cpuResource := resource.CPUResource{}
	cpuResource.InitRoutes(router)

	// Get cpu info
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/cpu/percent", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.True(t, true, w.Body.Len() > 0)

}
