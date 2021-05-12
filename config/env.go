package config

import (
	"os"
)

func getStringFromEnv(key string, prefix string) string {
	if v, ok := os.LookupEnv(prefix + key); ok {
		return v
	}
	return ""
}

type GlobalConfig struct {
	AlphaVantageApiKey string
	EncrypterApi string
}

func GetConfig() GlobalConfig {
	config := GlobalConfig{
		AlphaVantageApiKey: getStringFromEnv("ALPHA_VANTAGE_API_KEY", ""),
		EncrypterApi: getStringFromEnv("ENCRYPTER_API", ""),
	}

	return config
}
