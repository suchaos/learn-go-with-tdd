package command_line

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}

	server := &PlayerServer{&store}

	t.Run("return pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("return floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("return 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusNotFound)
	})
}

func assertResponseStatus(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("response code is wrong, got %d want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	return req
}
