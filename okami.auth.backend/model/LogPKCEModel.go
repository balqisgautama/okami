package model

import "database/sql"

// created at 08-27-2022
type LogPKCE struct {
	LogPKCEID      sql.NullInt64  `json:"log_pkce_id"`
	Step1          sql.NullTime   `json:"step1"`
	Step2          sql.NullTime   `json:"step2"`
	Step3          sql.NullTime   `json:"step3"`
	UserClientID   sql.NullString `json:"user_client_id"`
	SecretCode     sql.NullString `json:"secret_code"`
	CodeChallenger sql.NullString `json:"code_challenger"`
}
