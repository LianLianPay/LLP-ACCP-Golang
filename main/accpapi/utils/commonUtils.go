package utils

import (
	"encoding/json"
	"fmt"
)

func ObjectToString(obj interface{}) (string, error) {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return "", fmt.Errorf("转换JSON字符串失败: %s", err)
	}

	jsonString := string(jsonBytes)
	return jsonString, nil
}

func StringToObject(jsonString string, obj interface{}) error {
	err := json.Unmarshal([]byte(jsonString), obj)
	if err != nil {
		return fmt.Errorf("反序列化失败: %s", err)
	}

	return nil
}
