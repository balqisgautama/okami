package model

import "database/sql"

type LogAuditSystem struct {
	AuditID          sql.NullInt64  `json:"audit_id"`
	AuditTime        sql.NullTime   `json:"audit_time"`
	AuditDetail      sql.NullString `json:"audit_detail"`
	ResourceClientID sql.NullString `json:"resource_client_id"`
	DataOld          sql.NullString `json:"data_old"`
	DataNew          sql.NullString `json:"data_new"`
	Action           sql.NullInt16  `json:"action"`
}
