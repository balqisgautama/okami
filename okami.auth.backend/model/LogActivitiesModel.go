package model

import "database/sql"

type LogActivity struct {
	ActivityID       sql.NullInt64  `json:"activity_id"`
	ActivityTime     sql.NullTime   `json:"activity_time"`
	ActivityDetail   sql.NullString `json:"activity_detail"`
	ResourceClientID sql.NullString `json:"resource_client_id"`
}
