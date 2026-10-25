package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetEmptyTaskList(t *testing.T) {
	api := NewAPI(NewTaskStore())
	request := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	response := httptest.NewRecorder()

	api.Routes().ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
	}
	if response.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("Content-Type = %q, want application/json", response.Header().Get("Content-Type"))
	}

	var tasks []Task
	if err := json.NewDecoder(response.Body).Decode(&tasks); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(tasks) != 0 {
		t.Fatalf("len(tasks) = %d, want 0", len(tasks))
	}
}

func TestPostTask(t *testing.T) {
	api := NewAPI(NewTaskStore())
	request := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`{"title":"Learn net/http"}`))
	response := httptest.NewRecorder()

	api.Routes().ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusCreated)
	}

	var task Task
	if err := json.NewDecoder(response.Body).Decode(&task); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if task.ID != 1 || task.Title != "Learn net/http" || task.Done {
		t.Fatalf("created task = %+v", task)
	}
}

func TestPostTaskErrors(t *testing.T) {
	tests := []struct {
		name string
		body string
	}{
		{name: "empty title", body: `{"title":"   "}`},
		{name: "invalid JSON", body: `{"title":`},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			api := NewAPI(NewTaskStore())
			request := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(test.body))
			response := httptest.NewRecorder()

			api.Routes().ServeHTTP(response, request)

			if response.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want %d", response.Code, http.StatusBadRequest)
			}

			var errorResponse map[string]string
			if err := json.NewDecoder(response.Body).Decode(&errorResponse); err != nil {
				t.Fatalf("decode error response: %v", err)
			}
			if errorResponse["error"] == "" {
				t.Fatal("error response must contain an error message")
			}
		})
	}
}

func TestUnsupportedMethod(t *testing.T) {
	api := NewAPI(NewTaskStore())
	request := httptest.NewRequest(http.MethodPut, "/tasks", nil)
	response := httptest.NewRecorder()

	api.Routes().ServeHTTP(response, request)

	if response.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusMethodNotAllowed)
	}
}

func TestGetTasksAfterAddingTask(t *testing.T) {
	api := NewAPI(NewTaskStore())
	handler := api.Routes()

	postRequest := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`{"title":"Write tests"}`))
	postResponse := httptest.NewRecorder()
	handler.ServeHTTP(postResponse, postRequest)

	if postResponse.Code != http.StatusCreated {
		t.Fatalf("POST status = %d, want %d", postResponse.Code, http.StatusCreated)
	}

	getRequest := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	getResponse := httptest.NewRecorder()
	handler.ServeHTTP(getResponse, getRequest)

	if getResponse.Code != http.StatusOK {
		t.Fatalf("GET status = %d, want %d", getResponse.Code, http.StatusOK)
	}

	var tasks []Task
	if err := json.NewDecoder(getResponse.Body).Decode(&tasks); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(tasks) != 1 || tasks[0].Title != "Write tests" {
		t.Fatalf("tasks = %+v, want one task with title Write tests", tasks)
	}
}
