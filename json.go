package helper

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func MustJsonString(obj interface{}) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		logrus.WithError(err).Warn("method", "MustJsonString")
		return ""
	}

	jsonStr := string(bytes)

	return jsonStr
}

func JsonParse(jsonBytes []byte, receiver interface{}) error {
	return json.Unmarshal(jsonBytes, receiver)
}

func JsonString(obj interface{}) (*string, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	jsonStr := string(bytes)

	return &jsonStr, nil
}
