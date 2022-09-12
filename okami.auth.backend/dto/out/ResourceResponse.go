package response

import "time"

// created at 08-18-2022
type ResourceGeneral struct {
	ResourceID int64     `json:"resource_id"`
	ClientID   string    `json:"client_id"`
	Surname    string    `json:"surname"`
	Nickname   string    `json:"nickname"`
	AccessTo   string    `json:"access_to"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}
