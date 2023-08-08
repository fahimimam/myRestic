package logging

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"runtime"
	"strings"
	"time"
)

var (
	logger *logrus.Logger
)

func NewLogger() *logrus.Logger {
	logger = logrus.New()
	logger.SetReportCaller(true)
	logger.Formatter = getCustomLogFormatter()

	level := viper.GetString("log.level")
	if level == "" {
		level = "error"
	}
	l, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.Error("Failed to parse log level: ", err)
	}
	logger.Level = l
	return logger
}

func Get() logrus.FieldLogger {
	if logger != nil {
		return logger
	}
	return NewLogger()
}

func getCustomLogFormatter() *nested.Formatter {
	return &nested.Formatter{
		TimestampFormat: time.RFC3339Nano,
		HideKeys:        true,
		NoColors:        false,
		NoFieldsColors:  false,
		ShowFullLevel:   true,
		CallerFirst:     true,
		CustomCallerFormatter: func(f *runtime.Frame) string {
			return fmt.Sprintf(" File:%s:%d", formatFilePath(f.File), f.Line)
		},
	}
}

// formatFilePath - Shortens File Name.
func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
