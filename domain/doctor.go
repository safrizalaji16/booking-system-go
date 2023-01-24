package domain

import (
	"encoding/json"

	"github.com/uptrace/bun"
)

type Doctor struct {
	bun.BaseModel
	ID           int64           `bun:"id,pk,autoincrement" json:"id"`
	Name         string          `bun:"name" json:"name"`
	Schedule     json.RawMessage `bun:"schedule" json:"schedule"`
	Email        string          `bun:"email" json:"email"`
	Password     string          `bun:"password" json:"password"`
	SpecialityID int64           `bun:"speciality_id,nullzero" json:"speciality_id"`
	Speciality   *Speciality     `bun:"rel:belongs-to,join:speciality_id=id" json:"speciality"`
	Bookings     []*Booking      `bun:"rel:has-many,join:id=doctor_id"`
}
