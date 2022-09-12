package model

import "database/sql"

// created at 06-17-2022
// updated at 06-24-2022
type UserGeneral struct {
	UserID   sql.NullInt64  `json:"user_id"`
	Username sql.NullString `json:"username"`
	Email    sql.NullString `json:"email"`
	Password sql.NullString `json:"password"`
	ClientID sql.NullString `json:"client_id"`
	Status   sql.NullInt64  `json:"status"`
	//ActivationID   sql.NullInt64  `json:"activation_id"`
	Locale         sql.NullString `json:"locale"`
	AdditionalInfo sql.NullString `json:"additional_info"`
	LastToken      sql.NullTime   `json:"last_token"`
	CreatedBy      sql.NullInt64  `json:"created_by"`
	CreatedClient  sql.NullString `json:"created_client"`
	CreatedAt      sql.NullTime   `json:"created_at"`
	UpdatedBy      sql.NullInt64  `json:"updated_by"`
	UpdatedClient  sql.NullString `json:"updated_client"`
	UpdatedAt      sql.NullTime   `json:"updated_at"`
	DeletedBy      sql.NullInt64  `json:"deleted_by"`
	DeletedClient  sql.NullString `json:"deleted_client"`
	DeletedAt      sql.NullTime   `json:"deleted_at"`
}

// created at 06-20-2022
// updated at 06-21-2022
//type UserSimple struct {
//	UserID         int64  `json:"user_id"`
//	Username       string `json:"username"`
//	Email          string `json:"email"`
//	ClientID       string `json:"client_id"`
//	Status         int64  `json:"status"`
//	Locale         string `json:"locale"`
//	AdditionalInfo string `json:"additional_info"`
//}
