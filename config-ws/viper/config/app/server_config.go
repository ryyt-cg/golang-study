package app

type ServerConfig struct {
	Host string `validate:"required"`
	Port int    `validate:"required,min=1,max=65535"`
}
