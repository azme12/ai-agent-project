package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/azme12/ai-agent-project/internal/config"
)

type GeminiService struct {
	config *config.Config
	client *http.Client
}

type GeminiRequest struct {
	Contents []GeminiContent `json:"contents"`
}

type GeminiContent struct {
	Parts []GeminiPart `json:"parts"`
}

type GeminiPart struct {
	Text string `json:"text"`
}

type GeminiResponse struct {
	Candidates []GeminiCandidate `json:"candidates"`
}

type GeminiCandidate struct {
	Content GeminiContent `json:"content"`
}

func NewGeminiService(cfg *config.Config) *GeminiService {
	return &GeminiService{
		config: cfg,
		client: &http.Client{Timeout: 30 * time.Second},
	}
}

func (g *GeminiService) ProcessCommand(command string) (string, error) {
	if g.config.GeminiAPIKey == "" {
		return g.mockProcessCommand(command)
	}

	url := fmt.Sprintf("%s/models/gemini-pro:generateContent", g.config.GeminiURL)

	request := GeminiRequest{
		Contents: []GeminiContent{
			{
				Parts: []GeminiPart{
					{
						Text: fmt.Sprintf(`You are an AI executive assistant. Process this command and respond with a clear, actionable response: %s

Please respond in a helpful, professional manner. If the command involves scheduling, emailing, or task management, provide specific details about what actions should be taken.`, command),
					},
				},
			},
		},
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

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
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("gemini API error: %d - %s", resp.StatusCode, string(body))
	}

	var response GeminiResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	if len(response.Candidates) == 0 {
		return "", fmt.Errorf("no response from Gemini API")
	}

	if len(response.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("empty response from Gemini API")
	}

	return response.Candidates[0].Content.Parts[0].Text, nil
}

func (g *GeminiService) mockProcessCommand(command string) (string, error) {
	fmt.Printf("Gemini API key not configured. Using mock implementation.\n")
	fmt.Printf("Processing command with Gemini: %s\n", command)

	// Mock responses based on command type
	lowerCommand := strings.ToLower(command)

	if strings.Contains(lowerCommand, "schedule") || strings.Contains(lowerCommand, "meeting") {
		return "I'll help schedule a meeting. Please provide the attendees, date, time, and meeting title.", nil
	} else if strings.Contains(lowerCommand, "email") || strings.Contains(lowerCommand, "send") {
		return "I'll help send an email. Please provide the recipient, subject, and message content.", nil
	} else if strings.Contains(lowerCommand, "remind") || strings.Contains(lowerCommand, "task") {
		return "I'll set a reminder. Please provide the task details and deadline.", nil
	} else if strings.Contains(lowerCommand, "calendar") || strings.Contains(lowerCommand, "schedule") {
		return "I'll check the calendar. What specific information would you like to know about the schedule?", nil
	}

	return fmt.Sprintf("I understand you want me to: %s. How can I help with this?", command), nil
}
