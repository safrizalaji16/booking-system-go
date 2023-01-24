package requests

import "time"

type CreateBookingRequest struct {
	DoctorID  int64     `json:"doctor_id"`
	PatientID int64     `json:"patient_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
