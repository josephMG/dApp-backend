package config

import (
	"log"

	"github.com/spf13/viper"
)

// Env has environment stored
type Env struct {
	ServerPort         string `mapstructure:"SERVER_PORT"`
	Environment        string `mapstructure:"ENV"`
	LogOutput          string `mapstructure:"LOG_OUTPUT"`
	DBUsername         string `mapstructure:"DB_USER"`
	DBPassword         string `mapstructure:"DB_PASS"`
	DBHost             string `mapstructure:"DB_HOST"`
	DBPort             string `mapstructure:"DB_PORT"`
	DBName             string `mapstructure:"DB_NAME"`
	JWTSecret          string `mapstructure:"JWT_SECRET"`
	MaxMultipartMemory int64  `mapstructure:"MAX_MULTIPART_MEMORY"`
	TimeZone           string `mapstructure:"TIMEZONE"`
	EthereumURL        string `mapstructure:"ETHEREUM_URL"`
	ContractAddr       string `mapstructure:"CONTRACT_ADDRESS"`
	AccountPrivateKey  string `mapstructure:"ACCOUNT_PRIVATE_KEY"`
}

var globalEnv = Env{
	MaxMultipartMemory: 10 << 20, // 10 MB
}

func GetEnv() Env {
	return globalEnv
}

func NewEnv() *Env {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot read cofiguration", err)
	}

	viper.SetDefault("TIMEZONE", "UTC")

	err = viper.Unmarshal(&globalEnv)
	if err != nil {
		log.Fatal("environment cant be loaded: ", err)
	}

	return &globalEnv
}
