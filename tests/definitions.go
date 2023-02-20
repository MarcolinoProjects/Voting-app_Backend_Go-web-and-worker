package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"votingMicroservicesApp/pkg/models"
)

// assertVotingSessionFetched asserts that the voting session was fetched successfully.
func assertVotingSessionFetched(t *testing.T, w *httptest.ResponseRecorder) {
	assert.Equal(t, 200, w.Code)

	// Parse the response body into a Voting struct.
	var votingSessionFetch models.Voting
	err := json.Unmarshal(w.Body.Bytes(), &votingSessionFetch)
	if err != nil {
		t.Fatal("Failed to unmarshal response body")
	}
}

// assertVotingSessionDeleted asserts that the voting session was deleted successfully.
func assertVotingSessionDeleted(t *testing.T, w *httptest.ResponseRecorder) {
	assert.Equal(t, 200, w.Code)
}

// whenFetchVotingSessionRequest sends a GET request to fetch a voting session.
func whenFetchVotingSessionRequest(t *testing.T, votingSession models.Voting, router *gin.Engine) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", "/api/"+votingSession.UUID, nil)
	if err != nil {
		t.Fatal("Failed to create GET request")
	}
	req.Header.Add("Content-Type", "application/json")

	// Send the request and record the response.
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

// whenDeleteVotingSessionRequest sends a DELETE request to delete a voting session.
func whenDeleteVotingSessionRequest(t *testing.T, votingSession models.Voting, router *gin.Engine) *httptest.ResponseRecorder {
	req, err := http.NewRequest("DELETE", "/api/"+votingSession.UUID, nil)
	if err != nil {
		t.Fatal("Failed to create DELETE request")
	}
	req.Header.Add("Content-Type", "application/json")

	// Send the request and record the response.
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

// assertVotingSessionWasCreated asserts that the voting session was created successfully.
func assertVotingSessionWasCreated(t *testing.T, w *httptest.ResponseRecorder) (models.Voting, error) {
	assert.Equal(t, 201, w.Code)

	// Parse the response body into a Voting struct.
	var votingSession models.Voting
	err := json.Unmarshal(w.Body.Bytes(), &votingSession)
	if err != nil {
		t.Fatal("Failed to unmarshal response body")
	}
	return votingSession, err
}

// whenCreateRequest sends a POST request to create a voting session.
func whenCreateRequest(payload *strings.Reader, router *gin.Engine, w *httptest.ResponseRecorder) *http.Request {
	req, _ := http.NewRequest("POST", "/api/create", payload)
	req.Header.Add("Content-Type", "application/json")

	// Send the request and record the response.
	router.ServeHTTP(w, req)
	return req
}

// givenVotingSession creates a voting session payload for testing.
func givenVotingSession() *strings.Reader {
	return strings.NewReader(`  {
    "uuid": null,
    "name": "voting",
    "candidates": [
        {
          "uuid": "null",
          "name": "Roberts",
          "votes": 5
        },
        {
          "uuid": "null",
          "name": "Park",
          "votes": 0
        }
    ]
  }`)
}
