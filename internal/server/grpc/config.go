package grpc

import (
	"net"
	"strconv"
)

type Config struct {
	Host string `json:"host" yaml:"host" env:"HOST" validate:"notEmpty"`
	Port int    `json:"port" yaml:"port" env:"PORT" validate:"notEmpty"`
}

func (c *Config) Addr() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}
