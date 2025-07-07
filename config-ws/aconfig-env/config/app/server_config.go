package app

type ServerConfig struct {
	Host     string `validate:"required"`
	HttpPort string `validate:"required"`
}
