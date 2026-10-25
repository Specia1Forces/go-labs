package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TaskStore struct {
	mutex  sync.RWMutex
	nextID int
	tasks  map[int]Task
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		nextID: 1,
		tasks:  make(map[int]Task),
	}
}

func (store *TaskStore) Add(title string) Task {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	task := Task{
		ID:    store.nextID,
		Title: title,
		Done:  false,
	}

	store.tasks[task.ID] = task
	store.nextID++

	return task
}

func (store *TaskStore) List() []Task {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	tasks := make([]Task, 0, len(store.tasks))
	for _, task := range store.tasks {
		tasks = append(tasks, task)
	}

	return tasks
}

type API struct {
	store *TaskStore
}

func NewAPI(store *TaskStore) *API {
	return &API{store: store}
}

func (api *API) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", api.handleTasks)

	return mux
}

func (api *API) handleTasks(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		api.listTasks(response)
	case http.MethodPost:
		api.createTask(response, request)
	default:
		writeJSON(response, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed",
		})
	}
}

func (api *API) listTasks(response http.ResponseWriter) {
	writeJSON(response, http.StatusOK, api.store.List())
}

func (api *API) createTask(response http.ResponseWriter, request *http.Request) {
	var input struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
		writeJSON(response, http.StatusBadRequest, map[string]string{
			"error": "invalid JSON",
		})
		return
	}

	input.Title = strings.TrimSpace(input.Title)
	if input.Title == "" {
		writeJSON(response, http.StatusBadRequest, map[string]string{
			"error": "title must not be empty",
		})
		return
	}

	task := api.store.Add(input.Title)
	writeJSON(response, http.StatusCreated, task)
}

func writeJSON(response http.ResponseWriter, status int, value any) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)

	if err := json.NewEncoder(response).Encode(value); err != nil {
		log.Printf("encode JSON response: %v", err)
	}
}
