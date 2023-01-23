package domain

import (
	"encoding/json"

	"github.com/uptrace/bun"
)

type Doctor struct {
	bun.BaseModel
	ID       int64           `bun:"id,pk,autoincrement" json:"id"`
	Name     string          `bun:"name" json:"name"`
	Schdule  json.RawMessage `bun:"schdule" json:"schdule"`
	Email    string          `bun:"email" json:"email"`
	Password string          `bun:"password" json:"password"`
}
