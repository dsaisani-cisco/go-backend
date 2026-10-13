package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"danielsaisani/cmd/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

type stubClock struct {
	time time.Time
}

func (s *stubClock) Now() time.Time {
	return s.time
}

func TestAPI_Now(t *testing.T) {
	gin.SetMode(gin.TestMode)

	date := time.Date(2026, 1, 1, 12, 30, 0, 0, time.UTC)
	api := New(service.New(&stubClock{time: date}))

	router := gin.New()
	api.RegisterRoutes(router)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/now", nil))

	assert.Equal(t, http.StatusOK, recorder.Code)

	var body nowResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	assert.Equal(t, date, body.Now)
}
