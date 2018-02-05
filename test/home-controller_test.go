package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../controller"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	router := SetUpRouter()

	homeController := controller.HomeController{}
	homeController.InitRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
