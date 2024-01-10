package config

type Config struct {
	System System `yaml:"system"`
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	ES     Es     `yaml:"es"`
	Jwt    Jwt    `yaml:"jwt"`
	Site   Site   `yaml:"site"`
}
