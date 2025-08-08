package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/azme12/ai-agent-project/internal/agent"
	"github.com/azme12/ai-agent-project/internal/api"
	"github.com/azme12/ai-agent-project/internal/config"
	"github.com/azme12/ai-agent-project/pkg/logger"
)

var agentService *agent.Service

func main() {
	// Initialize logger and config
	logr := logger.New()
	cfg, err := config.Load()
	if err != nil {
		logr.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	// Initialize API clients
	calendar := api.NewGoogleCalendarService(cfg)
	email := api.NewSendGridEmailService(cfg)
	nlp := api.NewGeminiService(cfg)

	// Initialize agent
	agentService = agent.NewService(cfg, logr, calendar, email, nlp)

	// Start the agent
	if err := agentService.Start(); err != nil {
		logr.Error("Failed to start agent", "error", err)
		os.Exit(1)
	}

	// HTTP endpoints
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/schedule", scheduleHandler)
	http.HandleFunc("/email", emailHandler)
	http.HandleFunc("/nlp", nlpHandler)
	http.HandleFunc("/status", statusHandler)

	// Use configured port
	port := cfg.ServerPort
	if port == "" {
		port = "8080"
	}

	logr.Info("Starting HTTP server", "port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logr.Error("HTTP server failed", "error", err)
		os.Exit(1)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "healthy",
		"service": "ai-agent",
		"version": "1.0.0",
	})
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "running",
		"service": "ai-agent",
		"endpoints": []string{
			"GET /health",
			"POST /schedule",
			"POST /email",
			"POST /nlp",
			"GET /status",
		},
	})
}

func scheduleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req struct {
		Task string `json:"task"`
	}
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Task == "" {
		http.Error(w, "Task is required", http.StatusBadRequest)
		return
	}

	if err := agentService.ProcessTask(req.Task); err != nil {
		http.Error(w, fmt.Sprintf("Failed to process task: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Task processed successfully",
		"task":    req.Task,
	})
}

func emailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req struct {
		Task string `json:"task"`
	}
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Task == "" {
		http.Error(w, "Task is required", http.StatusBadRequest)
		return
	}

	if err := agentService.ProcessTask(req.Task); err != nil {
		http.Error(w, fmt.Sprintf("Failed to process email task: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Email task processed successfully",
		"task":    req.Task,
	})
}

func nlpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req struct {
		Command string `json:"command"`
	}
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Command == "" {
		http.Error(w, "Command is required", http.StatusBadRequest)
		return
	}

	response, err := agentService.ProcessNLPCommand(req.Command)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to process NLP command: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":   "success",
		"response": response,
		"command":  req.Command,
	})
}
