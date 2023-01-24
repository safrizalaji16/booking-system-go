package domain

import "github.com/uptrace/bun"

type Patient struct {
	bun.BaseModel
	ID       int64      `bun:"id,pk,autoincrement" json:"id"`
	Name     string     `bun:"name" json:"name"`
	Email    string     `bun:"email" json:"email"`
	Password string     `bun:"password" json:"password"`
	Bookings []*Booking `bun:"rel:has-many,join:id=patient_id"`
}
