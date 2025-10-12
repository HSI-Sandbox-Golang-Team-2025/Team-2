package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/HSI-Sandbox-Golang-Team-2025/Team-2/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRegistrationSuccess(t *testing.T) {
	app := SetupApp()

	resp, err := Registration(app, `{"nip":"ARN-2402001","password":"123","name":"Luthfi"}`)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var response fiber.Map

	json.NewDecoder(resp.Body).Decode(&response)

	responseData := response["data"].(map[string]any)

	assert.NotEmpty(t, responseData["token"].(string))
}

func TestRegistrationFailed(t *testing.T) {
	app := SetupApp()

	Registration(app, `{"nip":"ARN-2402001","password":"123","name":"Luthfi"}`)

	resp, err := Registration(app, `{"nip":"ARN-2402001","password":"123","name":"Luthfi"}`)

	assert.NoError(t, err)
	assert.Equal(t, 500, resp.StatusCode)

	var response fiber.Map

	json.NewDecoder(resp.Body).Decode(&response)

	assert.NotEmpty(t, response["message"].(string))
}

func TestLoginSuccess(t *testing.T) {
	app := SetupApp()

	Registration(app, `{"nip":"ARN-2402001","password":"123","name":"Luthfi"}`)

	resp, err := Login(app, `{"nip":"ARN-2402001","password":"123"}`)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var response fiber.Map

	json.NewDecoder(resp.Body).Decode(&response)

	responseData := response["data"].(map[string]any)

	assert.NotEmpty(t, responseData["token"].(string))
}

func TestLoginFailed(t *testing.T) {
	app := SetupApp()

	resp, err := Login(app, `{"nip":"ARN-2402001","password":"123"}`)

	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)

	var response fiber.Map

	json.NewDecoder(resp.Body).Decode(&response)

	assert.Equal(t, "Invalid credentials!", response["message"].(string))
}

func Registration(app *fiber.App, payload string) (*http.Response, error) {
	req := httptest.NewRequest(
		"POST",
		"/api/auth/register",
		strings.NewReader(payload),
	)

	req.Header.Set("Content-Type", "application/json")

	return app.Test(req)
}

func Login(app *fiber.App, payload string) (*http.Response, error) {
	req := httptest.NewRequest(
		"POST",
		"/api/auth/login",
		strings.NewReader(payload),
	)

	req.Header.Set("Content-Type", "application/json")

	return app.Test(req)
}
