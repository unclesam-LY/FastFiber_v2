package config

type Config struct {
	DB     DB     `yaml:"db"`
	Log    Log    `yaml:"log"`
	System System `yaml:"system"`
}
