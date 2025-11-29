package info

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AppInfo struct {
	AppName     string `json:"appName"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Env         string `json:"env"`
	Ip          string `json:"ip"`
	GitCommit   string `json:"gitCommit"`
}

func (ai AppInfo) Validate() error {
	return validation.ValidateStruct(&ai,
		validation.Field(&ai.AppName, validation.Required),
		validation.Field(&ai.Description, validation.Required),
		validation.Field(&ai.Version, validation.Required),
	)
}
