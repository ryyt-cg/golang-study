package ecode

import (
	"error-code-lookup/pkg/model"
)

type Ecode struct {
	model.Base
	ID          string `gorm:"primaryKey" json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Resolution  string `json:"resolution"`
}

func (ec *Ecode) ToEcodeResponse(ecode *Ecode) *Response {
	return &Response{
		ID:          ecode.ID,
		Code:        ecode.Code,
		Description: ecode.Description,
		Resolution:  ecode.Resolution,
	}
}

func (ec *Ecode) ToEcodeResponses(ecodes []Ecode) []Response {
	ecodeResponses := make([]Response, len(ecodes))

	for i, v := range ecodes {
		ecodeResponses[i] = *ec.ToEcodeResponse(&v)
	}

	return ecodeResponses
}
