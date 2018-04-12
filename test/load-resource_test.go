package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../resource"
	"github.com/stretchr/testify/assert"
)

func TestGetAverage(t *testing.T) {
	router := SetUpRouter()

	resource.InitLoadRoutes(router)

	// Get load average
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/load/average", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.True(t, true, w.Body.Len() > 0)

}

func TestGetMisc(t *testing.T) {
	router := SetUpRouter()

	resource.InitLoadRoutes(router)

	// Get load misc
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/load/misc", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.True(t, true, w.Body.Len() > 0)

}
