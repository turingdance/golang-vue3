package log

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// 配置日志切割
// LogFileCut 日志文件切割
func NewRolefile(fileName string) *rotatelogs.RotateLogs {
	logier, err := rotatelogs.New(
		// 切割后日志文件名称
		fileName+".%Y%m%d.log",
		rotatelogs.WithLinkName(fileName),        // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(30*24*time.Hour),   // 文件最大保存时间
		rotatelogs.WithRotationSize(1024*1024*5), // 日志切割时间间隔
	)

	if err != nil {
		panic(err)
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  logier,
		logrus.FatalLevel: logier,
		logrus.DebugLevel: logier,
		logrus.WarnLevel:  logier,
		logrus.ErrorLevel: logier,
		logrus.PanicLevel: logier,
	},
		// 设置分割日志样式
		&LogFormatter{})
	logrus.AddHook(lfHook)
	return logier
}
