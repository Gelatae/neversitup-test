package config

import (
	"fmt"
	"log"
	"os"
)

type OracleConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type AppConfig struct {
	Port string
}

func requireEnv(s string) string {
	result := os.Getenv(s)
	if result == "" {
		log.Panic(fmt.Errorf(`error unable to get env: %s`, s))
	}
	return result
}

func GetOracleConfig() OracleConfig {
	return OracleConfig{
		Host:     requireEnv("ORACLEDB_HOST"),
		Port:     requireEnv("ORACLEDB_PORT"),
		User:     requireEnv("ORACLEDB_USER"),
		Password: requireEnv("ORACLEDB_PASSWORD"),
		DBName:   requireEnv("ORACLEDB_NAME"),
	}
}

func GetAppConfig() AppConfig {
	conf := AppConfig{
		Port: requireEnv("APP_PORT"),
	}
	return conf
}
