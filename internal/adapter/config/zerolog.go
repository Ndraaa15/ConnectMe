package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewZerolog() {
	var consoleLogger io.Writer = zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
		FormatLevel: func(i interface{}) string {
			return "[" + i.(string) + "]"
		},
	}

	timeString := time.Now().Format("2006-01-02")
	_, err := os.Open(fmt.Sprintf("log/connectme-%s.log", timeString))
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(fmt.Sprintf("log/connectme-%s.log", timeString))
			if err != nil {
				log.Fatal().Err(err).Msg("failed to create log file")
			}
		} else {
			log.Fatal().Err(err).Msg("failed to open log file")
		}
	}

	fileLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("log/connectme-%s.log", timeString),
		MaxSize:    5,
		MaxBackups: 10,
		MaxAge:     14,
		Compress:   true,
	}

	multiOutput := zerolog.MultiLevelWriter(consoleLogger, fileLogger)

	zerolog.ErrorFieldName = "err"
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}

	log.Logger = zerolog.New(multiOutput).With().Timestamp().Caller().Logger()
}
