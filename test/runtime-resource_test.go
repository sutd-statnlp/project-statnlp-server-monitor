package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../resource"
	"github.com/stretchr/testify/assert"
)

func TestGetGoOS(t *testing.T) {
	router := SetUpRouter()

	runtimeResource := resource.RuntimeResource{}
	runtimeResource.InitRoutes(router)

	// Get runtime go OS
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/runtime/goos", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.True(t, true, w.Body.Len() > 0)

}
