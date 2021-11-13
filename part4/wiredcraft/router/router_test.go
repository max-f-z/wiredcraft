package router

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"wiredcraft/store"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	// Setup in memory storage for unit test
	store.SetupInMemoryStore("polaris")

	router := SetupRouter()

	req, err := http.NewRequest("GET", "/welcome", nil)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, "hello polaris", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPut(t *testing.T) {
	// Setup in memory storage for unit test
	store.SetupInMemoryStore("polaris")

	router := SetupRouter()

	body := `{
			"name": "alex"
		}`

	req, err := http.NewRequest("PUT", "/welcome", bytes.NewReader([]byte(body)))
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	req, err = http.NewRequest("GET", "/welcome", nil)
	assert.Nil(t, err)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, "hello alex", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}
