package model

import "database/sql"

// created at 08-16-2022
// updated at 08-18-2022
type ResourceGeneral struct {
	ResourceID sql.NullInt64  `json:"resource_id"`
	ClientID   sql.NullString `json:"client_id"`
	Surname    sql.NullString `json:"surname"`
	Nickname   sql.NullString `json:"nickname"`
	AccessTo   sql.NullString `json:"access_to"`
	CreatedAt  sql.NullTime   `json:"created_at"`
	UpdatedAt  sql.NullTime   `json:"updated_at"`
	DeletedAt  sql.NullTime   `json:"deleted_at"`
}
