package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportRoute(t *testing.T) {
	server := httptest.NewServer(RouterFactory())
	defer server.Close()

	req, err := http.NewRequest("GET", server.URL+"/report", nil)
	assert.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	assert.Equal(t, res.StatusCode, http.StatusOK)
}