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

func TestCreatePractice(t *testing.T) {
	app := SetupApp()

	Registration(app, `{"nip":"ARN-2402001","password":"123","name":"Luthfi"}`)
	Login(app, `{"nip":"ARN-2402001","password":"123"}`)

	resp, err := CreatePractice(app, `{
    "trackId": 2,
    "title": "Latihan 2",
    "body": "Latihan ini anda akan menguji kemampuan React anda",
    "questions": [
      {
        "question": "Apakah 1 + 1 = 2?",
        "answerChoices": [
          {
            "answer": "Benar",
            "isCorrect": true
          },
          {
            "answer": "Salah",
            "isCorrect": false
          }
        ]
      }
    ]
	}`)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var response fiber.Map

	json.NewDecoder(resp.Body).Decode(&response)

	assert.NotEmpty(t, response["data"].(map[string]any))
}

func CreatePractice(app *fiber.App, payload string) (*http.Response, error) {
	req := httptest.NewRequest(
		"POST",
		"/api/practices",
		strings.NewReader(payload),
	)

	req.Header.Set("Content-Type", "application/json")

	return app.Test(req)
}

func GetPractices(app *fiber.App) (*http.Response, error) {
	req := httptest.NewRequest(
		"GET",
		"/api/practices",
		nil,
	)

	req.Header.Set("Content-Type", "application/json")

	return app.Test(req)
}

func GetPractice(app *fiber.App) (*http.Response, error) {
	req := httptest.NewRequest(
		"GET",
		"/api/practices/1",
		nil,
	)

	return app.Test(req)
}
