package config

import "github.com/spf13/viper"

func ENV() string {
	return viper.GetString("env")
}

func GetDbPort() string {
	return viper.GetString("port")
}

func GetDbHost() string {
	return viper.GetString("mysql.dbhost")
}

func GetDbName() string {
	return viper.GetString("mysql.dbname")
}

func GetDbUser() string {
	return viper.GetString("mysql.dbuser")
}

func GetDbPassword() string {
	return viper.GetString("mysql.dbpass")
}
