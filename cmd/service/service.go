package service

import "time"

type Clock interface {
	Now() time.Time
}

type TimeService struct {
	clock Clock
}

func New(clock Clock) *TimeService {
	return &TimeService{clock: clock}
}

// Now returns..... the current time, what else did you expect?
func (s *TimeService) Now() time.Time {
	return s.clock.Now()
}
