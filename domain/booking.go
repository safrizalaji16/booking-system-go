package domain

import "github.com/uptrace/bun"

type Booking struct {
	bun.BaseModel
	ID   int64  `bun:"id,pk,autoincrement" json:"id"`
	Name string `bun:"name" json:"name"`
}
