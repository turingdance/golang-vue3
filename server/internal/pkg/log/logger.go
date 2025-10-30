package log

import (
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

type LogConf struct {
	Level int
	File  string
}

func Initialize(appconf LogConf) {
	// logrus.SetFormatter(&logrus.TextFormatter{
	// 	TimestampFormat: "2006-01-02 15:04:05",
	// })
	fdir := filepath.Dir(appconf.File)
	os.MkdirAll(fdir, 755)
	logrus.SetFormatter(NewLogFormatter())
	// logrus.SetOutput(os.Stdout)
	//ile, err := os.OpenFile(appconf.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	rolefie := NewRolefile(appconf.File)
	writers := []io.Writer{rolefie, os.Stdout}
	//同时写文件和屏幕
	fileAndStdoutWriter := io.MultiWriter(writers...)
	logrus.SetLevel(logrus.Level(appconf.Level))
	logrus.SetOutput(fileAndStdoutWriter)
}
