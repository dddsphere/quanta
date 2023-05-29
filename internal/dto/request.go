package dto

type (
	CreateList struct {
		ReqID  string
		UserID string
		Name   string `json:"name"`
	}
)
