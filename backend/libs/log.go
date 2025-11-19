package libs

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type LogConf struct {
	Path     string        `json:"path"`
	Level    string        `json:"level"`
	KeepDays time.Duration `json:"keep_days"`
}

func (c *LogConf) InitLogrus() {
	logrus.SetReportCaller(true)
	if c.Path != "" {
		writer := c.BuildLogWriter()
		logrus.SetOutput(writer)
	}
	logrus.SetLevel(c.ParseLevel())
	logrus.SetFormatter(GetFormatter())
}

func (c *LogConf) BuildLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	if c.Path != "" {
		writer := c.BuildLogWriter()
		logger.SetOutput(writer)
	}
	logger.SetLevel(c.ParseLevel())
	logger.SetFormatter(GetFormatter())
	return logger
}

func (c *LogConf) ParseLevel() logrus.Level {
	level, err := logrus.ParseLevel(c.Level)
	if c.Level == "" || err != nil {
		return logrus.InfoLevel
	}
	return level
}

func (c *LogConf) BuildLogWriter() io.Writer {
	if c.KeepDays == 0 {
		c.KeepDays = 3
	}
	writer, _ := rotatelogs.New(
		c.Path+".%Y%m%d",
		rotatelogs.WithLinkName(c.Path),
		rotatelogs.WithMaxAge(c.KeepDays*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	return writer
}

func GetFormatter() logrus.Formatter {
	return &formatter{}
}

type formatter struct{}

func (f *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	callerInfo := ""
	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		callerInfo = fmt.Sprintf("[%s:%d] ", fName, entry.Caller.Line)
	}
	timestamp := entry.Time.Format("2006-01-02 15:04:05.000")
	level := strings.ToUpper(entry.Level.String())
	newLog := fmt.Sprintf("%s [%s] %s%s\n", timestamp, level, callerInfo, entry.Message)

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	b.WriteString(newLog)
	return b.Bytes(), nil
}
