package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/azme12/ai-agent-project/internal/config"
)

type EmailService struct {
	config *config.Config
	client *http.Client
}

func NewSendGridEmailService(cfg *config.Config) *EmailService {
	return &EmailService{
		config: cfg,
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (e *EmailService) SendEmail(to, subject, body string) error {
	// TODO: Implement actual SendGrid API integration

	if e.config.SendGridAPIKey == "" {
		fmt.Printf("SendGrid API key not configured. Using mock implementation.\n")
		fmt.Printf("Sending email to: %s\nSubject: %s\nBody: %s\n", to, subject, body)
		return nil
	}

	// Real SendGrid API implementation:
	url := "https://api.sendgrid.com/v3/mail/send"

	emailData := map[string]interface{}{
		"personalizations": []map[string]interface{}{
			{
				"to": []map[string]string{
					{"email": to},
				},
			},
		},
		"from": map[string]string{
			"email": "your-verified-sender@yourdomain.com",
		},
		"subject": subject,
		"content": []map[string]string{
			{
				"type":  "text/plain",
				"value": body,
			},
		},
	}

	jsonData, _ := json.Marshal(emailData)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer "+e.config.SendGridAPIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := e.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("sendgrid API error: %d", resp.StatusCode)
	}

	fmt.Printf("Sending email to: %s\nSubject: %s\nBody: %s\n", to, subject, body)

	return nil
}
