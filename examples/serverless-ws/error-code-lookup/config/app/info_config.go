package app

import validation "github.com/go-ozzo/ozzo-validation/v4"

type AppInfoConfig struct {
	Name        string
	Description string
	Env         string
	Version     string
	GitVersion  string
}

func (aic AppInfoConfig) Validate() error {
	return validation.ValidateStruct(&aic,
		validation.Field(&aic.Name, validation.Required),
		validation.Field(&aic.Description, validation.Required),
		validation.Field(&aic.Version, validation.Required),
	)
}
