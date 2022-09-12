package response

import "okami.auth.backend/util"

type ErrorResponse struct {
	ErrorCode             string      `json:"error_code"`
	ErrorMessage          string      `json:"error_message"`
	AdditionalInformation interface{} `json:"additional_information"`
}

type SuccessResponse struct {
	Status                string      `json:"status"`
	Message               string      `json:"message"`
	AdditionalInformation interface{} `json:"additional_information"`
}

type APIResponse struct {
	Okami OkamiMessage `json:"okami_auth"`
}

type OkamiMessage struct {
	Header  Header  `json:"header"`
	Payload Payload `json:"payload"`
}

type Payload struct {
	Status StatusResponse `json:"status"`
	Data   interface{}    `json:"data"`
	Other  interface{}    `json:"other"`
}

type Header struct {
	RequestID string `json:"request_id"`
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
}

type StatusResponse struct {
	Success        bool     `json:"success"`
	Code           string   `json:"code"`
	Message        string   `json:"message"`
	Detail         string   `json:"detail"`
	AdditionalInfo []string `json:"additional_info"`
}

func (input APIResponse) String() string {
	return util.StructToJSON(input)
}
