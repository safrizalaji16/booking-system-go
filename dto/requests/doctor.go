package requests

import "encoding/json"

type CreateDoctorRequest struct {
	Name     string          `json:"name"`
	Schdule  json.RawMessage `json:"schdule"`
	Email    string          `json:"email"`
	Password string          `json:"password"`
}
