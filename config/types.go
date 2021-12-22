package config

import (
	"regexp"

	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

var (
	levels *regexp.Regexp = regexp.MustCompile("^(trace|debug|info|warn|error|fatal)$")
	format *regexp.Regexp = regexp.MustCompile("^(json)$")

	// DefaultConfig values for CLI
	DefaultConfig = &Bootstrap{
		ConfigFile: "config/default.yaml",
		LogOptions: &Logger{
			Level:  "info",
			Format: "logfmt",
		},
		Server: &Server{
			Http: &Server_HTTP{
				Hostname: "0.0.0.0:8000",
				Timeout: &durationpb.Duration{
					Seconds: 3,
				},
			},
			Grpc: &Server_GRPC{
				Hostname: "0.0.0.0:11000",
				Timeout: &durationpb.Duration{
					Seconds: 3,
				},
			},
		},
		Data: &Data{
			Redis: &Data_Redis{
				Hostname: "redis",
				Port:     6379,
				Database: 1,
				Username: "",
				Password: "",
				Timeout: &durationpb.Duration{
					Seconds: 3,
				},
			},
			Mongo: &Data_Mongo{
				Hostnames: []string{
					"mongo:27017",
				},
				Database: "mongo",
				Username: "",
				Password: "",
				Timeout: &durationpb.Duration{
					Seconds: 3,
				},
			},
		},
	}
)

// LogOptions struct parameters for logging level and format
type LogOptions struct {
	Level  string `yaml:"level,omitempty"`
	Format string `yaml:"format,omitempty"`
}
