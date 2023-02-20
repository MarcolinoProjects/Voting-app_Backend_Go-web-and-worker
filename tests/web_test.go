package tests

import (
	"log"
	"net/http/httptest"
	"os"
	"testing"
	"votingMicroservicesApp/pkg/config"
	"votingMicroservicesApp/pkg/handlers"
)

func TestMain(m *testing.M) {
	config.InitializeAppConfig(false)
	code := m.Run()
	config.ShutDown()
	log.Println("SHUT")
	os.Exit(code)
}

func TestCreateVotingSession(t *testing.T) {
	router := handlers.SetupRouter()
	w := httptest.NewRecorder()

	payload := givenVotingSession()
	whenCreateRequest(payload, router, w)
	_, _ = assertVotingSessionWasCreated(t, w)
}

func TestRetrieveInfoAboutVoteSession(t *testing.T) {
	router := handlers.SetupRouter()
	w := httptest.NewRecorder()

	payload := givenVotingSession()
	req := whenCreateRequest(payload, router, w)
	votingSession, err := assertVotingSessionWasCreated(t, w)

	w = whenFetchVotingSessionRequest(t, req, err, votingSession, w, router)
	assertVotingSessionFetched(t, w)
}
func TestDeleteVoteSession(t *testing.T) {
	router := handlers.SetupRouter()
	w := httptest.NewRecorder()

	payload := givenVotingSession()
	req := whenCreateRequest(payload, router, w)
	votingSession, err := assertVotingSessionWasCreated(t, w)

	w = whenDeleteVotingSessionRequest(t, req, err, votingSession, w, router)
	assertVotingSessionDeleted(t, w)
}
