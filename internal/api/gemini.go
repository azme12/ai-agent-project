package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/azme12/ai-agent-project/internal/config"
)

type GeminiService struct {
	config *config.Config
	client *http.Client
}

func NewGeminiService(cfg *config.Config) *GeminiService {
	return &GeminiService{
		config: cfg,
		client: &http.Client{Timeout: 30 * time.Second},
	}
}

func (g *GeminiService) ProcessCommand(command string) (string, error) {
	// TODO: Implement actual Gemini API integration

	if g.config.GeminiAPIKey == "" {
		fmt.Printf("Gemini API key not configured. Using mock implementation.\n")
		fmt.Printf("Processing command with Gemini: %s\n", command)

		// Mock response based on command type
		if len(command) > 0 {
			return fmt.Sprintf("Processed command: %s", command), nil
		}
		return "No command provided", nil
	}

	// Real Gemini API implementation:
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent"

	requestData := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{"text": fmt.Sprintf("You are an AI assistant. Process this command: %s", command)},
				},
			},
		},
	}

	jsonData, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Add API key as query parameter
	q := req.URL.Query()
	q.Add("key", g.config.GeminiAPIKey)
	req.URL.RawQuery = q.Encode()

	resp, err := g.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to process command: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("gemini API error: %d", resp.StatusCode)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	// Extract text from response
	if candidates, ok := response["candidates"].([]interface{}); ok && len(candidates) > 0 {
		if candidate, ok := candidates[0].(map[string]interface{}); ok {
			if content, ok := candidate["content"].(map[string]interface{}); ok {
				if parts, ok := content["parts"].([]interface{}); ok && len(parts) > 0 {
					if part, ok := parts[0].(map[string]interface{}); ok {
						if text, ok := part["text"].(string); ok {
							return text, nil
						}
					}
				}
			}
		}
	}

	fmt.Printf("Processing command with Gemini: %s\n", command)

	// Mock response based on command type
	if len(command) > 0 {
		return fmt.Sprintf("Processed command: %s", command), nil
	}

	return "No command provided", nil
}
