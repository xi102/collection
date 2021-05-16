package routers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitRouter(t *testing.T) {
	router := InitRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"message\":\"hello,xi102\"}", w.Body.String())
}
