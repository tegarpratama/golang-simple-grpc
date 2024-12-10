package message

type (
	MessageReq struct {
		Message string `json:"message"`
	}
)

type (
	MessageRes struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
)
