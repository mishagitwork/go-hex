package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type SpecificLevelWriter struct {
	io.Writer
	Levels []zerolog.Level
}

func (w SpecificLevelWriter) WriteLevel(level zerolog.Level, p []byte) (int, error) {
	for _, l := range w.Levels {
		if l == level {
			return w.Write(p)
		}
	}
	return len(p), nil
}

func createLoggerWriter() zerolog.LevelWriter {
	runLogFile, _ := os.OpenFile(
		fmt.Sprintf("logs/%s.log", time.Now().Format(time.RFC3339)),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)

	return zerolog.MultiLevelWriter(
		SpecificLevelWriter{
			Writer: zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05.999999999"},
			Levels: []zerolog.Level{
				zerolog.DebugLevel, zerolog.InfoLevel, zerolog.WarnLevel,
			},
		},
		SpecificLevelWriter{
			Writer: zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05.999999999"},
			Levels: []zerolog.Level{
				zerolog.ErrorLevel, zerolog.FatalLevel, zerolog.PanicLevel,
			},
		},
		zerolog.ConsoleWriter{Out: runLogFile, NoColor: true, TimeFormat: "15:04:05.999999999"},
	)
}
