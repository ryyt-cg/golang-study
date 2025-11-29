package app

type DatabaseConfig struct {
	Driver         string               `yaml:"driver" validate:"required"`
	Host           string               `yaml:"host" validate:"required"`
	Port           int                  `yaml:"port"`
	Name           string               `yaml:"name" validate:"required"` // Database name
	Username       string               `yaml:"username"`
	Password       string               `yaml:"password"`
	ConnectionPool ConnectionPoolConfig `yaml:"connectionPool"`
}

type ConnectionPoolConfig struct {
	MaxIdleConnection int `default:"10" validate:"required"`
	MaxOpenConnection int `default:"100" validate:"required"`
	MaxIdleTime       int `yaml:"maxIdleTime" default:"50" validate:"required"` // in seconds
}
