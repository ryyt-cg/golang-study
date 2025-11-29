package app

type ServerConfig struct {
	Host     string `yaml:"host" validate:"required"`
	HttpPort string `yaml:"httpPort" validate:"required"`
}
