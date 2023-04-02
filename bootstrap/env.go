package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	ServerHost      string 
	ServerPort      string 
	AppEnv         string
	ContextTimeout int
	DbDriver       string
	DbHost         string
	DbPort         string
	DbPass         string
	DbName         string
	DbUser         string
	ServerKey string
	ClientKey string
	ClientUrl string
	MidtransEnv int
	Expiry int
	GinMode        string
}

type EnvStruct struct {
	ServerHost      string `mapstructure:"SERVER_HOST"`
	ServerPort      string `mapstructure:"SERVER_PORT"`
	AppEnv         string `mapstructure:"APP_ENV"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	DbDriverdev    string `mapstructure:"DRIVER_DEV"`
	DbHostdev      string `mapstructure:"DB_HOST_DEV"`
	DbPortdev      string `mapstructure:"DB_PORT_DEV"`
	DbPassdev      string `mapstructure:"DB_PASS_DEV"`
	DbNamedev      string `mapstructure:"DB_NAME_DEV"`
	DbUserdev      string `mapstructure:"DB_USER_DEV"`
	DbDriver       string `mapstructure:"DRIVER"`
	DbHost         string `mapstructure:"DB_HOST"`
	DbPort         string `mapstructure:"DB_PORT"`
	DbPass         string `mapstructure:"DB_PASS"`
	DbName         string `mapstructure:"DB_NAME"`
	DbUser         string `mapstructure:"DB_USER"`
	ServerKey         string `mapstructure:"SERVER_KEY"`
	Expiry int `mapstructure:"EXPIRY"`
	ClientUrl         string `mapstructure:"CLIENT_URL"`
	ClientKey         string `mapstructure:"CLIENT_KEY"`
}

func NewEnv() *Env {
	env := EnvStruct{}
	config := Env{}

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "DEVELOPMENT" {
		config.ServerHost = env.ServerHost
		config.ServerPort = env.ServerPort
		config.ContextTimeout = env.ContextTimeout
		config.AppEnv = env.AppEnv
		config.DbDriver = env.DbDriverdev
		config.DbHost = env.DbHostdev
		config.DbPort = env.DbPortdev
		config.DbPass = env.DbPassdev
		config.DbName = env.DbNamedev
		config.DbUser = env.DbUserdev
		config.ServerKey = env.ServerKey
		config.ClientKey = env.ClientKey
		config.ClientUrl = env.ClientUrl
		config.MidtransEnv = 1
		config.Expiry = env.Expiry
		config.GinMode = "debug"
		log.Println("The App is running in development env")
	} else if env.AppEnv == "PRODUCTION" {
		config.ServerHost = env.ServerHost
		config.ServerPort = env.ServerPort
		config.ContextTimeout = env.ContextTimeout
		config.AppEnv = env.AppEnv
		config.DbDriver = env.DbDriver
		config.DbHost = env.DbHost
		config.DbPort = env.DbPort
		config.DbPass = env.DbPass
		config.DbName = env.DbName
		config.DbUser = env.DbUser
		config.MidtransEnv = 2
		config.Expiry = env.Expiry
		config.GinMode = "release"
		log.Println("The App is running in production env")
	} else {
		log.Fatal("Environment can't be loaded")
	}

	return &config
}
