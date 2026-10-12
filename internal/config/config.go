package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	_ = godotenv.Load()
}

func GetServiceBusConnectionString() string {
	return os.Getenv("SERVICE_BUS_CONNECTION_STRING")
}

func GetQueueName() string {
	return os.Getenv("QUEUE_NAME")
}
