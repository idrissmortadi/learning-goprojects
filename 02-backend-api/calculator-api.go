package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type RequestOperands struct {
	Number1 float64 `json:"number1"`
	Number2 float64 `json:"number2"`
}

var logger *slog.Logger

func init() {
	// Initialize the logger
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received request at home endpoint", "method", r.Method, "url", r.URL)
	fmt.Fprintf(w, "Hello, World!\n")
}

func createOperation(operation string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Request received", "endpoint", operation, "method", r.Method, "url", r.URL)

		if r.Method != http.MethodPost {
			logger.Error("Invalid method", "method", r.Method)
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}

		var data RequestOperands
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			logger.Error("Error decoding JSON", "error", err)
			http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
			return
		}

		logger.Info("Payload decoded", "number1", data.Number1, "number2", data.Number2)

		var result float64
		switch operation {
		case "add":
			result = data.Number1 + data.Number2
		case "subtract":
			result = data.Number1 - data.Number2
		case "multiply":
			result = data.Number1 * data.Number2
		case "divide":
			if data.Number2 == 0 {
				logger.Error("Cannot divide by zero")
				http.Error(w, "Cannot divide by zero", http.StatusBadRequest)
				return
			}
			result = data.Number1 / data.Number2
		default:
			logger.Error("Invalid operation", "operation", operation)
			http.Error(w, "Invalid operation", http.StatusBadRequest)
			return
		}

		logger.Info("Operation successful", "operation", operation, "result", result)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]float64{"result": result})
	}
}

func main() {
	logger.Info("Starting server on :3000")

	http.HandleFunc("/", home)

	http.HandleFunc("/add", createOperation("add"))
	http.HandleFunc("/subtract", createOperation("subtract"))
	http.HandleFunc("/multiply", createOperation("multiply"))
	http.HandleFunc("/divide", createOperation("divide"))

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		logger.Error("Server failed to start", "error", err)
	}
}
