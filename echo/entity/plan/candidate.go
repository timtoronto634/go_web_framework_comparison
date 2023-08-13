package plan

import (
	"time"

	"github.com/google/uuid"
)

type CandidateID uuid.UUID

type CandidateStatus string

const (
	Open       CandidateStatus = "open"
	Scheduling CandidateStatus = "scheduling"
	Scheduled  CandidateStatus = "scheduled"
	Deleted    CandidateStatus = "deleted"
)

type Candidate struct {
	ID         CandidateID
	HostUserID string
	Date       time.Time
	Status     CandidateStatus
}

func NewCandidate(hostUserID string, date time.Time) *Candidate {
	return &Candidate{
		ID:         CandidateID(uuid.New()),
		HostUserID: hostUserID,
		Date:       date,
		Status:     Open,
	}
}

// func (c *Candidate) ChangeStatus(newStatus CandidateStatus) error {
// 	if !isValidCandidateTransition(c.Status, newStatus) {
// 		return errors.New("invalid status transition")
// 	}
// 	c.Status = newStatus
// 	return nil
// }

// func isValidCandidateTransition(current, newStatus CandidateStatus) bool {
// 	switch current {
// 	case Open:
// 		return newStatus == Scheduled || newStatus == Deleted
// 	case Scheduled:
// 		return newStatus == Completed || newStatus == Deleted
// 	case Completed:
// 		return false
// 	case Deleted:
// 		return false
// 	default:
// 		return false
// 	}
// }
