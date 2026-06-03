package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const LogURL = "http://4.224.186.213/evaluation-service/logs"

type LogRequest struct {
	Stack   string `json:"stack"`
	Level   string `json:"level"`
	Package string `json:"package"`
	Message string `json:"message"`
}

func Log(token, stack, level, packageName, message string) error {
	payload := LogRequest{
		Stack:   stack,
		Level:   level,
		Package: packageName,
		Message: message,
	}

	jsonData, _ := json.Marshal(payload)

	req, _ := http.NewRequest(
		"POST",
		LogURL,
		bytes.NewBuffer(jsonData),
	)

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Logging API Status:", resp.Status)

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response:", string(body))

	return nil

}
