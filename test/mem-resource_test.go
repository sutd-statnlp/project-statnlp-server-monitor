package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../resource"
	"github.com/stretchr/testify/assert"
)

func TestGetVirtual(t *testing.T) {
	router := SetUpRouter()

	memResource := resource.MemResource{}
	memResource.InitRoutes(router)

	// Get mem info
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/mem/virtual", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.True(t, true, w.Body.Len() > 0)

}

func TestGetSwap(t *testing.T) {
	router := SetUpRouter()

	memResource := resource.MemResource{}
	memResource.InitRoutes(router)

	// Get mem swap
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/mem/swap", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.True(t, true, w.Body.Len() > 0)

}
