package config

import "time"

type HttpConfig struct {
	Network      string
	Addr         string
	Timeout      time.Duration
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}

type GrpcConfig struct {
	Name    string
	Version string
	Addr    string
}

type EtcdConfig struct {
	Addrs   []string
	Timeout time.Duration
}

type DBConfig struct {
	Dialect         string
	Addr            string
	DSN             string        `yaml:"dsn"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifeTime time.Duration `yaml:"conn_max_life_time"`
}
