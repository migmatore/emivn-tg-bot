package logging

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func GetLogger(ctx context.Context) *Logger {
	return loggerFromContext(ctx)
}

func (l *Logger) SetLoggingLevel(level string) {
	logrusLevel, err := logrus.ParseLevel(level)
	if err != nil {
		log.Fatalln(err)
	}

	l.SetLevel(logrusLevel)
}

func NewLogger() *Logger {
	l := logrus.New()

	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			filename := path.Base(f.File)
			// TODO fix only logging.go filename with interface
			return fmt.Sprintf("%s:%d", filename, f.Line), fmt.Sprintf("%s()", f.Function)
		},
		DisableColors: false,
		FullTimestamp: true,
	}
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.InfoLevel)

	return &Logger{Logger: l}
}
