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
				Timeout: &Data_Redis_Timeout{
					Dial:               &durationpb.Duration{Seconds: 5},
					Read:               &durationpb.Duration{Seconds: 10},
					Write:              &durationpb.Duration{Seconds: 10},
					Idle:               &durationpb.Duration{Seconds: 30},
					IdleCheckFrequency: &durationpb.Duration{Nanos: 500000},
				},
			},
			Postgres: &Data_Postgres{
				Hostname: "postgres",
				Port:     5432,
				Database: "postgres",
				Schema:   "postgres",
				Username: "postgres",
				Password: "postgres",
				Ssl:      Data_Postgres_disable,
				Timeout: &Data_Postgres_Timeout{
					Dial:  &durationpb.Duration{Seconds: 5},
					Read:  &durationpb.Duration{Seconds: 10},
					Write: &durationpb.Duration{Seconds: 10},
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
