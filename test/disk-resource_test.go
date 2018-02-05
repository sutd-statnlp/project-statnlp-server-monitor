package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../resource"
	"github.com/stretchr/testify/assert"
)

func TestGetUsage(t *testing.T) {
	router := SetUpRouter()

	diskResource := resource.DiskResource{}
	diskResource.InitRoutes(router)

	// Get disk info
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/disk/usage", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.True(t, true, w.Body.Len() > 0)

}
