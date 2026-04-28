package cmd

import (
	"os"
)

func (server *ApiServer) loadEnv() []string {
	var errors []string

	server.apiRunType = os.Getenv("API_RUN_TYPE")
	if server.apiRunType == "" {
		errors = append(errors, "API_RUN_TYPE not set")
	}

	server.runtimeEnv = os.Getenv("RUNTIME_ENV")
	if server.runtimeEnv == "" {
		errors = append(errors, "RUNTIME_ENV not set")
	}

	// Instance: singleton pattern (phien ban ket noi duy nhat) DNS: Data Source Name (chuoi ket noi)
	server.dbInstance = os.Getenv("DB_DNS")
	if server.dbInstance == "" {
		errors = append(errors, "DB_DNS not set")
	}

	// server.logLevel = os.Getenv("LOG_LEVEL")
	// switch server.logLevel {
	// case "DEBUG":
	// 	log.SetOutputLevel(log.LOG_LEVEL_DEBUG)
	// case "INFO":
	// 	log.SetOutputLevel(log.LOG_LEVEL_INFO)
	// case "WARNING":
	// 	log.SetOutputLevel(log.LOG_LEVEL_WARNING)
	// case "ERROR":
	// 	log.SetOutputLevel(log.LOG_LEVEL_ERROR)
	// case "CRITICAL":
	// 	log.SetOutputLevel(log.LOG_LEVEL_CRITICAL)
	// default:
	// 	log.SetOutputLevel(log.LOG_LEVEL_DEBUG)
	// }

	return errors
}
