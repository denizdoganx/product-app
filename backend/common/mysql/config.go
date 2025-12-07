package mysql

import "time"

type Config struct {
	Host                  string
	Port                  string
	Username              string
	Password              string
	DbName                string
	MaxConnections        int
	MaxIdleConnections    int
	MaxConnectionIdleTime time.Duration
	Timeout               string
	ReadTimeout           string
	WriteTimeout          string
}
