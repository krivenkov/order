package http

type Config struct {
	Host string `json:"host" yaml:"host" env:"HOST"`
	Port int    `json:"port" yaml:"port" env:"PORT"`
}
