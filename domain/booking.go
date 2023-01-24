package domain

import (
	"time"

	"github.com/uptrace/bun"
)

type Booking struct {
	bun.BaseModel
	ID        int64     `bun:"id,pk,autoincrement" json:"id"`
	DoctorID  int64     `bun:"doctor_id,nullzero" json:"dotor_id"`
	PatientID int64     `bun:"patient_id,nullzero" json:"patent_id"`
	Doctor    *Doctor   `bun:"rel:belongs-to,join:doctor_id=id" json:"doctor"`
	Patient   *Patient  `bun:"rel:belongs-to,join:patient_id=id" json:"patient"`
	StartDate time.Time `bun:"start_date" json:"start_date"`
	EndDate   time.Time `bun:"end_date" json:"end_date"`
}
