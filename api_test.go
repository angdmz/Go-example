package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Pong struct {
	Sarasa  string `json:"sarasa"`
	Poronga int    `json:"poronga"`
}

func TestPingRoute(t *testing.T) {
	router := ServeRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var pong Pong
	err := json.Unmarshal(w.Body.Bytes(), &pong)
	assert.Nil(t, err)
	assert.Equal(t, pong.Sarasa, "jajs dice poronga")
	assert.Equal(t, pong.Poronga, 123)
}

func TestTrue(t *testing.T) {
	assert.True(t, true)
}
