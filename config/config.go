package config

import (
	"git.mudu.tv/dev-public/utils"
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
	DatabaseHost = utils.GetEnvWithDefault("MYSQL_HOST", "")
	DatabasePort = utils.GetEnvWithDefault("MYSQL_PORT", "")
	DatabaseUser = utils.GetEnvWithDefault("MYSQL_USER", "")
	DatabasePass = utils.GetEnvWithDefault("MYSQL_PASSWORD", "")
	DatabaseName = utils.GetEnvWithDefault("MYSQL_DBNAME", "")

}
