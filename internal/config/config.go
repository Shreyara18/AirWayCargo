package config

import (
	"os"
	"strconv"
)

var (
	AppName      string
	Port         string
	ServiceToken string

	DatabaseURL      string
	MaxDBConnections int

	LogLevel                 string
	LogFilePathSize          string
	WorkFlowTasksNotifyQueue string
	NotificationSenderBot    string
)

func init() {
	AppName = os.Getenv("AppName")
	Port = os.Getenv("Port")
	ServiceToken = os.Getenv("ServiceToken")

	// All database configurations are stored here
	DatabaseURL = os.Getenv("DatabaseURL")
	MaxDBConnections, _ = strconv.Atoi(os.Getenv("MaxDBConnections"))

	// Log related configuration
	LogLevel = os.Getenv("LogLevel")
	LogFilePathSize = os.Getenv("LogFilePathSize")

	// Setup global logger
	config := utils.NewLoggerConfig(AppName)
	config.SetLevel(LogLevel)

	//Logger = utils.NewLogger(config)

	WorkFlowTasksNotifyQueue = os.Getenv("WorkFlowTasksNotifyQueue")
	if WorkFlowTasksNotifyQueue == "" {
		WorkFlowTasksNotifyQueue = "https://sqs.ap-south-1.amazonaws.com/816374819873/workflow-tasks-notify-dev"
	}
	NotificationSenderBot = os.Getenv("NotificationSenderBot")

}
