package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitRouter(t *testing.T) {
	router := InitRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/files/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
