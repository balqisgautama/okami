package response

// created at 06-24-2022
type UserSimple struct {
	UserID         int64  `json:"user_id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	ClientID       string `json:"client_id"`
	Status         int64  `json:"status"`
	Locale         string `json:"locale"`
	AdditionalInfo string `json:"additional_info"`
}

// created at 07-05-2022
type UserSimpleWithToken struct {
	UserID         int64  `json:"user_id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	ClientID       string `json:"client_id"`
	Status         int64  `json:"status"`
	Locale         string `json:"locale"`
	AdditionalInfo string `json:"additional_info"`
	UserToken      string `json:"user_token"`
}
