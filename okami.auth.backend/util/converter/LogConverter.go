package util

import (
	"okami.auth.backend/constanta"
	"okami.auth.backend/model"
	"time"
)

func ToLogActivity(activityTime time.Time, requestType string, api string, resourceClientID string) (result model.LogActivity) {
	result.ActivityTime.Time = activityTime
	result.ActivityDetail.String = requestType + constanta.PrefixLog + api
	result.ResourceClientID.String = resourceClientID
	return
}

func ToLogAuditSystem(auditTime time.Time, actionType string, tableName string, resourceClientID string, dataOld string, dataNew string, actionID int16) (result model.LogAuditSystem) {
	result.AuditTime.Time = auditTime
	result.AuditDetail.String = actionType + constanta.PrefixLog + tableName
	result.ResourceClientID.String = resourceClientID
	result.DataOld.String = dataOld
	result.DataNew.String = dataNew
	result.Action.Int16 = actionID
	return
}
