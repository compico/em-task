package config

import "fmt"

type Database interface {
	GetUsername() string
	GetPassword() string
	GetHost() string
	GetPort() int
	GetDatabase() string
}

type database struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

func (d *database) GetUsername() string {
	return d.Username
}

func (d *database) GetPassword() string {
	return d.Password
}

func (d *database) GetHost() string {
	return d.Host
}

func (d *database) GetPort() int {
	return d.Port
}

func (d *database) GetDatabase() string {
	return d.Database
}

func (d *database) GetDsn() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.Database,
	)
}
