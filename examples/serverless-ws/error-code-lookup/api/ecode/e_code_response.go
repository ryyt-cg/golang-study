package ecode

type Response struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Resolution  string `json:"resolution"`
}
