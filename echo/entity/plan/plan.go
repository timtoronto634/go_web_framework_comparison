package plan

import (
	"errors"
	"time"
)

type Plan struct {
	Candidate   *Candidate
	Reservation *Reservation
}

func NewPlan(hostUserID string, date time.Time) *Plan {
	return &Plan{
		Candidate: NewCandidate(hostUserID, date),
	}
}

func (p *Plan) Request(guestUserID string) error {
	if p.Candidate.Status != Open {
		return errors.New("Candidate is not open for reservation")
	}

	p.Reservation = NewReservation(p.Candidate.ID, guestUserID)
	p.Candidate.Status = Scheduling

	return nil
}

// ... その他のメソッドもここに実装 ...
