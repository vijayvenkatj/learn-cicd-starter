package auth

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {

	headers := http.Header{}
	headers.Add("Authorization", "ApiKey Testing123")

	apiKey, err := GetAPIKey(headers)
	require.NoError(t, err)
	assert.Equal(t, apiKey, "Testing123")

	headers.Del("Authorization")
	_, err = GetAPIKey(headers)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrNoAuthHeaderIncluded)
}
