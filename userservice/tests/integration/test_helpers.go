package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

const (
	contentType = "application/json"
	urlReg      = "http://localhost:44044/registration"
	urlLog      = "http://localhost:44044/login"
)

func uniqueEmail() string {
	return "testgmail" + uuid.NewString() + "@gmail.com"
}

func registrateUser(t *testing.T) (string, string) {
	email := uniqueEmail()
	pass := "somePass"

	body := map[string]string{
		"first_name":  "Ivan",
		"middle_name": "Ivanovich",
		"last_name":   "Ivanov",
		"password":    pass,
		"email":       email,
	}

	b, err := json.Marshal(body)
	require.NoError(t, err)

	resp, err := http.Post(urlReg, contentType, bytes.NewReader(b))
	require.NoError(t, err)
	defer resp.Body.Close()

	expStatusCode := http.StatusOK

	var resBody struct {
		IsRegistered bool `json:"is_registered"`
	}

	require.NoError(t, json.NewDecoder(resp.Body).Decode(&resBody))
	require.True(t, resBody.IsRegistered)
	require.Equal(t, expStatusCode, resp.StatusCode)

	return email, pass
}

func loginUser(t *testing.T, email string, pass string) string {
	body := map[string]string{
		"email":    email,
		"password": pass,
	}

	b, err := json.Marshal(body)
	require.NoError(t, err)

	resp, err := http.Post(urlLog, contentType, bytes.NewReader(b))
	require.NoError(t, err)
	defer resp.Body.Close()

	cookies := resp.Cookies()
	var cookie *http.Cookie
	for _, c := range cookies {
		if c.Name == "sessionId" {
			cookie = c
		}
	}

	expBody := map[string]string{
		"first_name":  "Ivan",
		"middle_name": "Ivanovich",
		"last_name":   "Ivanov",
	}
	expStatusCode := http.StatusOK

	var bodyResp struct {
		User map[string]string `json:"user"`
	}

	require.NoError(t, json.NewDecoder(resp.Body).Decode(&bodyResp))
	require.Equal(t, expBody, bodyResp.User)
	require.Equal(t, expStatusCode, expStatusCode)
	require.NotNil(t, cookie)
	require.NotEmpty(t, cookie.Value)

	return cookie.Value
}
