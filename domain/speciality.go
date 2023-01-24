package domain

import "github.com/uptrace/bun"

type Speciality struct {
	bun.BaseModel
	ID      int64     `bun:"id,pk,autoincrement" json:"id"`
	Name    string    `bun:"name" json:"name"`
	Doctors []*Doctor `bun:"rel:has-many,join:id=speciality_id"`
}
