package model

import "database/sql"

// created at 06-20-2022
// updated at 07-04-2022
type ActivationAccount struct {
	ActivationID sql.NullInt64  `json:"activation_id"`
	Counter      sql.NullInt64  `json:"counter"`
	Code         sql.NullInt64  `json:"code"`
	EmailTo      sql.NullString `json:"email_to"`
	Expire       sql.NullInt64  `json:"expire"`
	Status       sql.NullInt64  `json:"status"`
	LinkValidate sql.NullString `json:"link_validate"`
	LinkResend   sql.NullString `json:"link_resend"`
}
