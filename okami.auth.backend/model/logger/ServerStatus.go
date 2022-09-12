package model

import "okami.auth.backend/util"

type ServerStatus struct {
	Status string `json:"status"`
}

func (object ServerStatus) String() string {
	return util.StructToJSON(object)
}
