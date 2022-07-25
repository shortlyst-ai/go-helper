package helper

import (
	"encoding/json"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type messageWrapper struct {
	Headers map[string]string `json:"headers"`
}

func getRequestIdFromMessage(messageBytes []byte, defaultValue func() string) string {

	msgObj := messageWrapper{}
	err := json.Unmarshal(messageBytes, &msgObj)
	if err != nil {
		logrus.WithField("method", "GetRequestIdFromMessage").WithError(err).Warning("error when unmarshalling message")
		if defaultValue != nil {
			return defaultValue()
		}
		return ""
	}

	reqId := msgObj.Headers["requestId"]
	if reqId == "" && defaultValue != nil {
		return defaultValue()
	}

	return reqId
}

func GetRequestIdFromMessage(message []byte) string {
	return getRequestIdFromMessage(message, UuidV4Fn())
}

func UuidV4Fn() func() string {
	return func() string {
		return uuid.NewV4().String()
	}
}
