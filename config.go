package main

import (
	"os"
	"strconv"
	"strings"
)

type ConfigStruct struct {
	TimeZone       string
	Port           int
	ProxyPort      int
	Domain         string
	PublicFolder   string
	Authentication Authentication
}

type Authentication struct {
	Username string
	PassHash string
	Type     string // none, basic, forms
}

var Config *ConfigStruct

func InitConfig() {
	Config = &ConfigStruct{
		Domain:       getEnv("DOMAIN", "localhost"),
		Port:         getEnvAsInt("PORT", 3000),
		ProxyPort:    getEnvAsInt("PROXY_PORT", 3001),
		TimeZone:     getEnv("TZ", "UTC"),
		PublicFolder: getEnv("PUBLIC_FOLDER", "/public/"),
		Authentication: Authentication{
			Username: getEnv("AUTH_USER", "admin"),
			PassHash: getEnv("AUTH_PASS", SHA256Hash(GenerateKey(64))), // Generate a random password as default in case user forgot to set flag
			Type:     strings.ToLower(getEnv("AUTH_TYPE", "none")),
		},
	}
}
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valueStr := getEnv(name, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}

	return defaultVal
}
