package models

type Config struct {
	Port        int
	Environment string
	Database    DatabaseConfig
	Cache       CacheConfig
}

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
	SslMode  string
	LogLevel int
}

type CacheConfig struct {
	Addr string
	Db   int
}
