package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

type mockClock struct {
	time time.Time
}

// Moves the stub clock forward by the given number of milliseconds.
func (m *mockClock) MoveForwardBy(milliseconds int64) {
	duration, err := time.ParseDuration(fmt.Sprintf("%dms", milliseconds))
	if err != nil {
		panic(err)
	}
	m.time = m.time.Add(duration)
}

func (m *mockClock) Now() time.Time {
	return m.time
}

func TestService_Now(t *testing.T) {
	date := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	mockClock := &mockClock{time: date}
	s := New(mockClock)

	now := s.Now()
	assert.Equal(t, date, now)

	mockClock.MoveForwardBy(1000)

	now = s.Now()
	assert.Equal(t, date.Add(1*time.Second), now)
}
