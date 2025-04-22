package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestGetTaskStatuses(t *testing.T) {
	p := os.Getenv("PORT")

	body := map[string]string{
		"status": "ACTIVIE",
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Could not marshal body: %v", err)
	}

	req, err := http.NewRequest("POST",
		fmt.Sprintf("http://127.0.0.1:%s/todo/task-status", p),
		bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Creating POST request error: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Executing POST request error: %v", err)
	}

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			t.Errorf("Closing response body error: %v", err)
		}
	}(resp.Body)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Reading response body error: %v", err)
	}
	bodyString := string(bodyBytes)
	t.Logf("Response body: %s", bodyString)
}
