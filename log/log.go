package log

import (
	"fmt"
	"io"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/mattn/go-colorable"
)

type LogFormatter struct {
}

func init() {
	logrus.SetFormatter(&LogFormatter{})
	logrus.SetOutput(colorable.NewColorableStdout())
}

func (l *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level.String() {
	case "debug":
		levelColor = 36 // cyan
	case "info":
		levelColor = 34 // blue
	case "warning":
		levelColor = 33 // yellow
	case "error", "fatal":
		levelColor = 31 // red
	default:
		levelColor = 37 // white
	}

	return []byte(fmt.Sprintf("\x1b[%dm%s\x1b[0m: %s\n", levelColor, strings.ToUpper(entry.Level.String()), entry.Message)), nil
}

func SetLevel(level string) error {
	l, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	logrus.SetLevel(l)
	return nil
}

func Debug(msg string, out io.Writer) {
	logrus.SetOutput(out)
	logrus.Debug(msg)
}

func Info(msg string, out io.Writer) {
	logrus.SetOutput(out)
	logrus.Info(msg)
}

func Warning(msg string, out io.Writer) {
	logrus.SetOutput(out)
	logrus.Warning(msg)
}

func Error(msg string, out io.Writer) {
	logrus.SetOutput(out)
	logrus.Error(msg)
}

func Fatal(msg string, out io.Writer) {
	logrus.SetOutput(out)
	logrus.Fatal(msg)
}
