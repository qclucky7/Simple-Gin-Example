package components

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *logrus.Logger

func SetUp() {
	rootPath, _ := os.Getwd()
	SetupLogger(filepath.Join(rootPath, "logs", "info.log"), filepath.Join(rootPath, "logs", "error.log"))
}

// 工作目录
func SetupLogger(infoPath string, errorPath string) {
	Logger = logrus.New()
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		DisableQuote:    true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return "", fmt.Sprintf("%s:%d", filepath.Base(frame.File), frame.Line)
		},
	})
	Logger.SetOutput(os.Stdout)

	fmt.Printf("INFO Path: %s Error Path: %s", infoPath, errorPath)

	infoFileHook := NewFileHook(logrus.InfoLevel, &lumberjack.Logger{
		Filename:   infoPath,
		MaxSize:    3,
		MaxBackups: 3,
		Compress:   true,
	})

	errorFileHook := NewFileHook(logrus.ErrorLevel, &lumberjack.Logger{
		Filename:   errorPath,
		MaxSize:    3,
		MaxBackups: 3,
		Compress:   true,
	})

	Logger.AddHook(infoFileHook)
	Logger.AddHook(errorFileHook)
}

type FileHook struct {
	Writer    *lumberjack.Logger
	LogLevels []logrus.Level
}

func NewFileHook(level logrus.Level, writer *lumberjack.Logger) *FileHook {
	return &FileHook{
		Writer: writer,
		LogLevels: []logrus.Level{
			level,
		},
	}
}

func (hook *FileHook) Levels() []logrus.Level {
	return hook.LogLevels
}

func (hook *FileHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write([]byte(line))
	return err
}
