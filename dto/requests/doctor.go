package requests

import "encoding/json"

type CreateDoctorRequest struct {
	Name         string          `json:"name"`
	Schedule     json.RawMessage `json:"schedule"`
	Email        string          `json:"email"`
	Password     string          `json:"password"`
	SpecialityID int64           `json:"speciality_id"`
}
