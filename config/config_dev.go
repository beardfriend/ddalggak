//go:build dev

package config

import "github.com/rs/zerolog"

var (
	LogLevel = zerolog.DebugLevel
)
