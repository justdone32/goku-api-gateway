package goku_log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

const (
	PanicLevel = logrus.PanicLevel
	FatalLevel = logrus.FatalLevel
	ErrorLevel = logrus.ErrorLevel
	WarnLevel  = logrus.WarnLevel
	InfoLevel  = logrus.InfoLevel
	DebugLevel = logrus.DebugLevel
	TraceLevel = logrus.TraceLevel
)

var (
	writer *FileWriterByPeriod = nil
	logger                     = logrus.New()

	logEnable = true
)

type Level = logrus.Level
type Fields = logrus.Fields
type Entry = logrus.Entry

func ParseLevel(lvl string) (Level, error) {
	return logrus.ParseLevel(lvl)
}

func init() {
	logger.SetLevel(logrus.WarnLevel)
	logger.SetFormatter(&LineFormatter{
		//ForceColors: true,
		//FullTimestamp:             true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	writer = NewFileWriteBytePeriod()
	logrus.RegisterExitHandler(func() {
		Close()
	})
}
func GetLogger() *logrus.Logger {
	return logger
}
func SetLevel(level Level) {

	logger.SetLevel(level)

}
func SetOutPut(enable bool, dir, file string, period LogPeriod, expire int) {
	logEnable = enable
	logger.SetOutput(writer)
	if enable {
		writer.Set(dir, file, period, time.Duration(expire)*time.Hour*24)
		writer.Open()
	} else {
		writer.Close()
	}
}

func Close() {
	writer.Close()
}
func WithFields(fields Fields) *Entry {

	return logger.WithFields(fields)
}

// Trace logs a message at level Trace on the standard logger.
func Trace(args ...interface{}) {
	if !logEnable {
		return
	}
	logger.Trace(args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {

	if !logEnable {
		return
	}
	logger.Debug(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {

	if !logEnable {
		return
	}
	logger.Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {

	if !logEnable {
		return
	}
	logger.Warn(args...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {

	if !logEnable {
		return
	}
	logger.Warning(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {

	if !logEnable {
		return
	}
	logger.Error(args...)
}
func panicout(args ...interface{}) string {
	defer func() {
		if e := recover(); e != nil {
			Close()
		}
	}()
	s, _ := encode(PanicLevel, args...)
	_, _ = os.Stderr.Write(s)
	logger.Panic(args...)
	return string(s)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	panic(panicout(args...))
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	s, e := encode(FatalLevel, args...)
	if e != nil {
		return
	}
	_, _ = os.Stderr.Write(s)
	logger.Fatal(args...)
}

// Tracef logs a message at level Trace on the standard logger.
func Tracef(format string, args ...interface{}) {

	if !logEnable {
		return
	}
	logger.Tracef(format, args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {

	if !logEnable {
		return
	}
	logger.Debugf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {

	if !logEnable {
		return
	}
	logger.Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {

	if !logEnable {
		return
	}
	logger.Warnf(format, args...)
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(format string, args ...interface{}) {

	if !logEnable {
		return
	}
	logger.Warningf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {

	if !logEnable {
		return
	}
	logger.Errorf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalf(format string, args ...interface{}) {

	s, e := encode(FatalLevel, fmt.Sprintf(format, args...))
	if e != nil {
		return
	}
	_, _ = os.Stdout.Write(s)
	logger.Fatalf(format, args...)

}

// Traceln logs a message at level Trace on the standard logger.
func Traceln(args ...interface{}) {

	if !logEnable {
		return
	}
	logger.Traceln(args...)
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {

	if !logEnable {
		return
	}
	logger.Debugln(args...)
}

// 无视lever输出一段trace 信息
func Print(args ...interface{}) {

	s, e := encode(TraceLevel, args...)
	if e != nil {
		return
	}
	if len(s) == 0 {
		return
	}

	if logEnable {
		writer.Write(s)
	} else {
		os.Stdout.Write(s)
	}
}

func encode(level Level, args ...interface{}) ([]byte, error) {
	entry := logrus.NewEntry(logger)
	entry.Message = fmt.Sprintln(args...)
	entry.Level = level
	entry.Time = time.Now()
	s, e := logger.Formatter.Format(entry)
	if e != nil {
		fmt.Println(e)
		fmt.Println(args...)
	}
	return s, e
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(args ...interface{}) {

	if !logEnable {
		return
	}
	logger.Infoln(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(args ...interface{}) {

	if !logEnable {
		return
	}
	logger.Warnln(args...)
}

// Warningln logs a message at level Warn on the standard logger.
func Warningln(args ...interface{}) {

	if !logEnable {
		return
	}
	logger.Warningln(args...)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...interface{}) {

	if !logEnable {
		return
	}
	logger.Errorln(args...)
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalln(args ...interface{}) {

	s, e := encode(FatalLevel, args...)
	if e != nil {
		return
	}
	_, _ = os.Stderr.Write(s)
	logger.Fatalln(args...)
}
