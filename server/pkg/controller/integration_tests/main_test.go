package integration_tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestHealthRoute(t *testing.T) {
	//req, err := http.NewRequest("GET", "http://127.0.0.1:8000/todo/healthz", nil)
	port := os.Getenv("API_HOST_PORT")
	req, err := http.NewRequest("GET", fmt.Sprintf("http://127.0.0.1:%s/todo/healthz", port), nil)
	if err != nil {
		t.Fatalf("Could not create req: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Could not make req: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", res.StatusCode)
	}
}

func TestGetTaskStatuses(t *testing.T) {
	port := os.Getenv("API_HOST_PORT")

	body := map[string]string{
		"status": "ACTIVIE",
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST",
		fmt.Sprintf("http://127.0.0.1:%s/todo/task-status", port),
		bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Error creating POST request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error executing POST request: %v", err)
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			t.Errorf("Error closing response body: %v", err)
		}
	}(resp.Body)

	fmt.Println("Status Code:", resp.StatusCode)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Could not read response body: %v", err)
	}
	bodyString := string(bodyBytes)
	t.Logf("Response body: %s", bodyString)

	//port := os.Getenv("API_HOST_PORT")
	//req, err := http.NewRequest("POST", fmt.Sprintf("http://127.0.0.1:%s/todo/task-status", port), nil)
	//if err != nil {
	//	t.Fatalf("Could not create req: %v", err)
	//}
	//
	//res, err := http.DefaultClient.Do(req)
	//if err != nil {
	//	t.Fatalf("Could not make req: %v", err)
	//}
	//
	//t.Logf("Response status: %d", res.StatusCode)
	//
	//if res.StatusCode != http.StatusOK {
	//	t.Errorf("Expected status 200, got %d", res.StatusCode)
	//}
	//
	//bodyBytes, err := io.ReadAll(res.Body)
	//if err != nil {
	//	t.Fatalf("Could not read response body: %v", err)
	//}
	//bodyString := string(bodyBytes)
	//t.Logf("Response body: %s", bodyString)
}
