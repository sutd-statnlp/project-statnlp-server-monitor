package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../resource"
	"github.com/stretchr/testify/assert"
)

func TestGetNetInterfaces(t *testing.T) {
	router := SetUpRouter()

	netResource := resource.NetResource{}
	netResource.InitRoutes(router)

	// Get net interfaces
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/net/interface", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.True(t, true, w.Body.Len() > 0)

}

func TestGetNetConnection(t *testing.T) {
	router := SetUpRouter()

	netResource := resource.NetResource{}
	netResource.InitRoutes(router)

	// Get net connection
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/net/connection", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.True(t, true, w.Body.Len() > 0)

}
