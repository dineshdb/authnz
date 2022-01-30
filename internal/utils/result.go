package utils

type GenericResult struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
