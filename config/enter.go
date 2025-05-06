package config

type Config struct {
	DB     DB     `yaml:"db"`
	Redis  Redis  `yaml:"redis"`
	Log    Log    `yaml:"log"`
	System System `yaml:"system"`
	Jwt    Jwt    `yaml:"jwt"`
}
