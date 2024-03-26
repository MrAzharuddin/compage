package utils

import "github.com/sirupsen/logrus"

type Log struct {
	log *logrus.Logger
}

func NewLog() Log {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return Log{
		log: log,
	}
}

func (l Log) GetLogger() *logrus.Logger {
	return l.log
}
