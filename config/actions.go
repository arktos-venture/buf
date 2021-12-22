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
