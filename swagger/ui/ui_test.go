package ui

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/myml/swag"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	assert := require.New(t)
	api := swag.New()
	h := Handler("/swagger_ui/", api)
	s := httptest.NewServer(h)
	resp, err := http.Get(s.URL + "/swagger_ui/")
	assert.NoError(err)
	defer resp.Body.Close()
	assert.Equal(resp.StatusCode, http.StatusOK)

	resp, err = http.Get(s.URL + "/swagger_ui/swagger.json")
	assert.NoError(err)
	defer resp.Body.Close()
	assert.Equal(resp.StatusCode, http.StatusOK)
}
