package response

import (
	"time"
)

// created at 08-19-2022
type UserGeneral struct {
	UserID         int64     `json:"user_id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	ClientID       string    `json:"client_id"`
	Status         int64     `json:"status"`
	Locale         string    `json:"locale"`
	AdditionalInfo string    `json:"additional_info"`
	LastToken      time.Time `json:"last_token"`
	CreatedBy      int64     `json:"created_by"`
	CreatedClient  string    `json:"created_client"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedBy      int64     `json:"updated_by"`
	UpdatedClient  string    `json:"updated_client"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedBy      int64     `json:"deleted_by"`
	DeletedClient  string    `json:"deleted_client"`
	DeletedAt      time.Time `json:"deleted_at"`
}
