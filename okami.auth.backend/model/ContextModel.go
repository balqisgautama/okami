package model

import logger "okami.auth.backend/model/logger"

type ContextModel struct {
	LoggerModel    logger.LoggerModel
	PermissionHave string
	//IsSignatureCheck     bool
	IsInternal           bool
	LimitedByCreatedBy   int64
	AuthAccessTokenModel AuthAccessTokenModel
	DBSchema             string
}

type AuthAccessTokenModel struct {
	RedisAuthAccessTokenModel
	ClientID                   string `json:"cid"`
	AuthenticationServerUserID int64  `json:"aid"`
	Locale                     string `json:"lang"`
}

type RedisAuthAccessTokenModel struct {
	ResourceUserID int64  `json:"rid"`
	Authentication string `json:"auth"`
	IPWhiteList    string `json:"ipl"`
	SignatureKey   string `json:"sign"`
}

func (input AuthAccessTokenModel) ConvertToRedisModel() RedisAuthAccessTokenModel {
	return RedisAuthAccessTokenModel{
		ResourceUserID: input.ResourceUserID,
		Authentication: input.Authentication,
		IPWhiteList:    input.IPWhiteList,
		SignatureKey:   input.SignatureKey,
	}
}
