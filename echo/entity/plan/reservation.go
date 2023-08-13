package plan

import (
	"github.com/google/uuid"
)

type ReservationID uuid.UUID

type ReservationStatus string

const (
	Requesting ReservationStatus = "requesting"
	Accepted   ReservationStatus = "accepted"
	Declined   ReservationStatus = "declined"
)

type Reservation struct {
	ID          ReservationID
	CandidateID CandidateID
	GuestUserID string
	Status      ReservationStatus
}

func NewReservation(candidateID CandidateID, guestUserID string) *Reservation {
	return &Reservation{
		ID:          ReservationID(uuid.New()), // 予約ID生成の実装に依存
		CandidateID: candidateID,
		GuestUserID: guestUserID,
		Status:      Requesting,
	}
}

// func (r *Reservation) ChangeStatus(newStatus ReservationStatus, candidateStatus CandidateStatus) error {
// 	if !isValidReservationTransition(r.Status, newStatus, candidateStatus) {
// 		return errors.New("invalid reservation status transition")
// 	}
// 	r.Status = newStatus
// 	return nil
// }

// func isValidReservationTransition(current, newStatus ReservationStatus, candidateStatus CandidateStatus) bool {
// 	switch current {
// 	case Requesting:
// 		if candidateStatus != Open {
// 			return false
// 		}
// 		return newStatus == Accepted || newStatus == Declined
// 	case Accepted:
// 		return newStatus == Declined && candidateStatus != Scheduled
// 	case Declined:
// 		return false
// 	default:
// 		return false
// 	}
// }
