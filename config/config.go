package config

import (
	"os"

	"github.com/spf13/cast"
)

var (
	NewStatusId  = "2cb46139-7c5f-4e0d-88f6-50fe79986fa6"
	SentStatusId = "f9c31d41-7e91-42a1-859b-e0991357566c"
	ReadStatusId = "303d5102-7595-4ba4-9c67-0c5d969d6a3e"
)

const (
	COURIER             = "courier"
	BRANCH_USER         = "branch_user"
	CLIENT              = "client"
	ANDROID_PLATFORM_ID = "6bd7c2e3-d35e-47df-93ce-ed54ed53f95f"
	IOS_PLATFORM_ID     = "f6631db7-09d0-4cd9-a03a-b7a590abb8c1"
)

// Config ...
type Config struct {
	ServiceName string
	Environment string // develop, staging, production
	LogLevel    string

	PostgresHost          string
	PostgresPort          int
	PostgresUser          string
	PostgresPassword      string
	PostgresDatabase      string
	MaxPgxPoolConnections int

	RPCPort  string
	HTTPPort string

	CourierServiceHost string
	CourierServicePort string

	GorushHost string
	GorushPort string

	UserServiceHost string
	UserServicePort string
}

// Load loads environment vars and inflates Config
func Load() Config {

	c := Config{}

	c.ServiceName = cast.ToString(getEnvOrReturnDefault("SERVICE_NAME", "template_notification_service"))
	c.Environment = cast.ToString(getEnvOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(getEnvOrReturnDefault("LOG_LEVEL", "debug"))

	c.PostgresHost = cast.ToString(getEnvOrReturnDefault("POSTGRES_HOST", "0.0.0.0"))
	c.PostgresPort = cast.ToInt(getEnvOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresUser = cast.ToString(getEnvOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getEnvOrReturnDefault("POSTGRES_PASSWORD", "admin"))
	c.PostgresDatabase = cast.ToString(getEnvOrReturnDefault("POSTGRES_DATABASE", "template_notification_service"))
	c.MaxPgxPoolConnections = cast.ToInt(getEnvOrReturnDefault("MAX_PGX_POOL_CONNECTIONS", 5))

	c.RPCPort = cast.ToString(getEnvOrReturnDefault("RPC_PORT", ":8081"))
	c.HTTPPort = cast.ToString(getEnvOrReturnDefault("HTTP_PORT", ":8080"))

	c.UserServiceHost = cast.ToString(getEnvOrReturnDefault("USER_SERVICE_HOST", "user-service"))
	c.UserServicePort = cast.ToString(getEnvOrReturnDefault("USER_GRPC_PORT", "80"))

	c.CourierServiceHost = cast.ToString(getEnvOrReturnDefault("COURIER_SERVICE_HOST", "courier-service"))
	c.CourierServicePort = cast.ToString(getEnvOrReturnDefault("COURIER_GRPC_PORT", "80"))

	c.GorushHost = cast.ToString(getEnvOrReturnDefault("GORUSH_SERVICE_HOST", "gorush-service"))
	c.GorushPort = cast.ToString(getEnvOrReturnDefault("GORUSH_GRPC_PORT", "3000"))

	return c
}

func getEnvOrReturnDefault(key string, defaultValue interface{}) interface{} {
	if os.Getenv(key) == "" {
		return defaultValue
	}
	return os.Getenv(key)
}
