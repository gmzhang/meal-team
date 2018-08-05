package config

import (
	"os"
)

var (
	// DatabaseHost ...
	DatabaseHost string
	// DatabaseUser ...
	DatabaseUser string
	// DatabasePass ...
	DatabasePass string
	// DatabaseName ...
	DatabaseName string
	// DatabasePort ...
	DatabasePort string
)

func init() {
	DatabaseHost = GetEnvWithDefault("MYSQL_HOST", "")
	DatabasePort = GetEnvWithDefault("MYSQL_PORT", "")
	DatabaseUser = GetEnvWithDefault("MYSQL_USER", "")
	DatabasePass = GetEnvWithDefault("MYSQL_PASSWORD", "")
	DatabaseName = GetEnvWithDefault("MYSQL_DBNAME", "")

}


func GetEnvWithDefault(key, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return val
}
