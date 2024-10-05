package common

import (
	"os"

	"github.com/beardfriend/ddalggak/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

var logFile *os.File

func SetLog(logDirname string, logFilename string) {
	err := os.MkdirAll(logDirname, os.ModePerm)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(logDirname+"/"+logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	logFile = file

	multi := zerolog.MultiLevelWriter(os.Stdout, logFile)
	logger := zerolog.New(multi).With().Timestamp().Logger()

	logger = logger.Level(config.LogLevel)

	log.Logger = logger

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

}

func LogFileCloser() error {
	return logFile.Close()
}

func ErrorHttpLog(err error, message string, code int) {
	log.Error().Stack().Err(err).Int("status", code).Msg(message)
}

func ErrorHttpLogV2(err error, origin string, message string, code int) {
	log.Error().Stack().Err(err).Int("status", code).Str("origin", origin).Msg(message)
}

func ErrorLog(err error, message string) {
	log.Error().Stack().Err(err).Msg(message)
}

func PanicLog(err error, message string) {
	log.Panic().Stack().Err(err).Msg(message)
}

func WarnLog(message string) {
	log.Warn().Stack().Msg(message)
}

func InfoLog(message string) {
	log.Info().Stack().Msg(message)
}

func DebugLog(message string) {
	log.Debug().Msg(message)
}
