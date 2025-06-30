package app

type InfoConfig struct {
	Name        string `validate:"required"`
	Description string `validate:"required"`
	Version     string `validate:"required"`
}
