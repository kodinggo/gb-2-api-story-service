package config

import "github.com/spf13/viper"

func GetDbHost() string {
	return viper.GetString("database.host")
}

func GetDbName() string {
	return viper.GetString("database.dbname")
}

func GetDbUser() string {
	return viper.GetString("database.user")
}

func GetDbPassword() string {
	return viper.GetString("database.password")
}

func GetDbPort() string {
	return viper.GetString("database.port")
}
