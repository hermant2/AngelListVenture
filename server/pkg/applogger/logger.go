package applogger

import (
	"github.com/sirupsen/logrus"
)

const (
	keyMessage = "message"
	keyError = "error"
)

type AppLogger struct {
	Client *logrus.Logger
}

var loggerInstance *AppLogger = nil
func Instance() *AppLogger{
	if loggerInstance == nil {
		logger := logrus.New()
		loggerInstance = &AppLogger{logger}
	}

	return loggerInstance
}

func (logger *AppLogger) Error(message string, err error) {
	params := map[string]interface{}{keyMessage: message, keyError: err}
	logger.Client.WithFields(params).Error()
}

func (logger *AppLogger) ErrorWithParams(message string, err error, params map[string]interface{}) {
	params[keyMessage] = message
	params[keyError] = err
	logger.Client.WithFields(params).Error()
}

func (logger *AppLogger) Warn(message string) {
	logger.Client.WithField(keyMessage, message).Warn()
}

func (logger *AppLogger) WarnWithParams(message string, params map[string]interface{}) {
	params[keyMessage] = message
	logger.Client.WithFields(params).Warn()
}

func (logger *AppLogger) Info(message string) {
	logger.Client.WithField(keyMessage, message).Info()
}

func (logger *AppLogger) InfoWithParams(message string, params map[string]interface{}) {
	params[keyMessage] = message
	logger.Client.WithFields(params).Info()
}

func (logger *AppLogger) Panic(message string, err error) {
	params := map[string]interface{}{keyMessage: message, keyError: err}
	logger.Client.WithFields(params).Panic()
}

func (logger *AppLogger) PanicWithParams(message string, err error, params map[string]interface{}) {
	params[keyMessage] = message
	params[keyError] = err
	logger.Client.WithFields(params).Panic()
}

func (logger *AppLogger) Fatal(message string, err error) {
	params := map[string]interface{}{keyMessage: message, keyError: err}
	logger.Client.WithFields(params).Fatal()
}

func (logger *AppLogger) FatalWithParams(message string, err error, params map[string]interface{}) {
	params[keyMessage] = message
	params[keyError] = err
	logger.Client.WithFields(params).Fatal()
}
