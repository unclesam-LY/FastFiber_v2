package config

import "fmt"

type System struct {
	Mode         string   `yaml:"mode"`
	Host         string   `yaml:"host"`
	Port         int      `yaml:"port"`
	Env          string   `yaml:"env"`
	Version      string   `yaml:"version"`
	AllowOrigins []string `yaml:"allow_origins"`
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
