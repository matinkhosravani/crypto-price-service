package app

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	AppEnv              string `mapstructure:"APP_ENV"`
	PriceExpirationTime int    `mapstructure:"PRICE_EXPIRATION_TIME_MINUTE"`
	JSONListenAddr      string `mapstructure:"JSON_LISTEN_ADDR"`
	GRPCListenAddr      string `mapstructure:"GRPC_LISTEN_ADDR"`
	RedisHost           string `mapstructure:"REDIS_HOST"`
	RedisPort           string `mapstructure:"REDIS_PORT"`
	RedisPassword       string `mapstructure:"REDIS_PASSWORD"`
	RedisDatabase       int    `mapstructure:"REDIS_DATABASE"`
	CacheType           string `mapstructure:"CACHE_TYPE"`
	PriceAPI            string `mapstructure:"PRICE_API"`
}

func NewEnv() *Env {
	env := Env{}

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
