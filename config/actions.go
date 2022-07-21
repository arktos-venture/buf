package config

import (
	"os"

	"github.com/hashicorp/go-hclog"
)

// LogFlagParse level logs
func LogFlagParse(name string, c *Logger) hclog.Logger {
	var level string

	if levels.MatchString(c.Format) {
		level = c.Format
	} else {
		level = "info"
	}

	return hclog.New(&hclog.LoggerOptions{
		Name:       name,
		Level:      hclog.LevelFromString(level),
		JSONFormat: format.MatchString(c.Format),
		Output:     os.Stdout,
	})
}

func (cfg *Data) CheckConfig(logger hclog.Logger, data []Service, services, envs []string) {
	var log hclog.Logger = logger.Named("CheckConfig")

	for _, s := range data {
		switch s {
		case Service_PostgreSQL:
			if cfg.GetPostgres() == nil {
				log.Error("failed to read postgres config")
				os.Exit(1)
			}
		case Service_Redis:
			if cfg.GetRedis() == nil {
				log.Error("failed to read redis config")
				os.Exit(1)
			}
		case Service_Nats:
			if cfg.GetNats() == nil {
				log.Error("failed to read nats config")
				os.Exit(1)
			}
		case Service_Meilisearch:
			if cfg.GetMeilisearch() == nil {
				log.Error("failed to read meilisearch config")
				os.Exit(1)
			}
		case Service_Keycloak:
			if cfg.GetKeycloak() == nil {
				log.Error("failed to read keycloak config")
				os.Exit(1)
			}
		}
	}

	// Check Mandatory services
	if services != nil {
		for _, c := range services {
			if _, ok := cfg.GetServices()[c]; !ok {
				log.With("config", c).Error("missing config to start, exit")
				os.Exit(1)
			}
		}
	}

	// Check Mandatory envs
	if envs != nil {
		for _, e := range envs {
			if env := os.Getenv(e); env == "" {
				log.With("env", e).Error("missing env to start, exit")
				os.Exit(1)
			}
		}
	}
}
