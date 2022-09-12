package model

import "database/sql"

// created at 08-22-2022
type ActivationUser struct {
	ActivationID      sql.NullInt64  `json:"activation_id"`
	Counter           sql.NullInt64  `json:"counter"`
	Code              sql.NullString `json:"code"`
	ExpiredAt         sql.NullTime   `json:"expired_at"`
	Status            sql.NullInt64  `json:"status"`
	UserID            sql.NullInt64  `json:"user_id"`
	EmailTo           sql.NullString `json:"email_to"`
	EmailLinkValidate sql.NullString `json:"email_link_validate"`
	EmailLinkResend   sql.NullString `json:"email_link_resend"`
}
