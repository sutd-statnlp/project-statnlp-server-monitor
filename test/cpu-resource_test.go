package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../resource"
	"github.com/stretchr/testify/assert"
)

func TestGetCPUSumPercent(t *testing.T) {
	router := SetUpRouter()

	cpuResource := resource.CPUResource{}
	cpuResource.InitRoutes(router)

	// Get cpu summary percent
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/cpu/sum/percent", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.True(t, true, w.Body.Len() > 0)
}

func TestGetCPUSumTime(t *testing.T) {
	router := SetUpRouter()

	cpuResource := resource.CPUResource{}
	cpuResource.InitRoutes(router)

	// Get cpu summary time
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/cpu/sum/time", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.True(t, true, w.Body.Len() > 0)
}

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

	// Get cpu percent
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/cpu/percent", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.True(t, true, w.Body.Len() > 0)

}
func TestGetTime(t *testing.T) {
	router := SetUpRouter()

	cpuResource := resource.CPUResource{}
	cpuResource.InitRoutes(router)

	// Get cpu time
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/cpu/time", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.True(t, true, w.Body.Len() > 0)

}
