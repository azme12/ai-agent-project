package main

import (
	"encoding/json"
	"io"
	"log"
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

	log.Printf("Starting HTTP server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logr.Error("HTTP server failed", "error", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
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

	var req struct {
		Task string `json:"task"`
	}
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := agentService.ProcessTask(req.Task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
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

	var req struct {
		Task string `json:"task"`
	}
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := agentService.ProcessTask(req.Task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
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

	var req struct {
		Command string `json:"command"`
	}
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	response, err := agentService.ProcessNLPCommand(req.Command)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"response": response})
}
