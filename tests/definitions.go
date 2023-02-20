package tests

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"votingMicroservicesApp/pkg/models"
)

func assertVotingSessionFetched(t *testing.T, w *httptest.ResponseRecorder) {
	assert.Equal(t, 200, w.Code)
	var votingSessionFetch models.Voting
	err := json.Unmarshal(w.Body.Bytes(), &votingSessionFetch)
	if err != nil {
		t.Fatal("Failure on Unmarshal")
	}
}
func assertVotingSessionDeleted(t *testing.T, w *httptest.ResponseRecorder) {
	assert.Equal(t, 200, w.Code)
}

func whenFetchVotingSessionRequest(t *testing.T, req *http.Request, err error, votingSession models.Voting, w *httptest.ResponseRecorder, router *gin.Engine) *httptest.ResponseRecorder {
	req, err = http.NewRequest("GET", "/api/"+votingSession.UUID, nil)
	if err != nil {
		t.Fatal("Failure on Request")
	}
	req.Header.Add("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
func whenDeleteVotingSessionRequest(t *testing.T, req *http.Request, err error, votingSession models.Voting, w *httptest.ResponseRecorder, router *gin.Engine) *httptest.ResponseRecorder {
	req, err = http.NewRequest("DELETE", "/api/"+votingSession.UUID, nil)
	if err != nil {
		t.Fatal("Failure on Request")
	}
	req.Header.Add("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func assertVotingSessionWasCreated(t *testing.T, w *httptest.ResponseRecorder) (models.Voting, error) {
	assert.Equal(t, 201, w.Code)
	var votingSession models.Voting
	err := json.Unmarshal(w.Body.Bytes(), &votingSession)
	if err != nil {
		t.Fatal("Failure on Unmarshal")
	}
	return votingSession, err
}

func whenCreateRequest(payload *strings.Reader, router *gin.Engine, w *httptest.ResponseRecorder) *http.Request {
	req, _ := http.NewRequest("POST", "/api/create", payload)
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	return req
}

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
