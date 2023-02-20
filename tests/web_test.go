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
	// Initialize the app configuration and shut it down when the tests are finished.
	config.InitializeAppConfig(false)
	code := m.Run()
	config.ShutDown()
	log.Println("App configuration shut down.")
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
	whenCreateRequest(payload, router, w)
	votingSession, _ := assertVotingSessionWasCreated(t, w)

	w = whenFetchVotingSessionRequest(t, votingSession, router)
	assertVotingSessionFetched(t, w)
}

func TestDeleteVoteSession(t *testing.T) {
	router := handlers.SetupRouter()
	w := httptest.NewRecorder()

	payload := givenVotingSession()
	whenCreateRequest(payload, router, w)
	votingSession, _ := assertVotingSessionWasCreated(t, w)

	w = whenDeleteVotingSessionRequest(t, votingSession, router)
	assertVotingSessionDeleted(t, w)
}
