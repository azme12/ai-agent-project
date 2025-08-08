package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/azme12/ai-agent-project/internal/config"
)

type EmailService struct {
	config *config.Config
	client *http.Client
}

type SendGridEmail struct {
	Personalizations []SendGridPersonalization `json:"personalizations"`
	From             SendGridFrom              `json:"from"`
	Subject          string                    `json:"subject"`
	Content          []SendGridContent         `json:"content"`
}

type SendGridPersonalization struct {
	To []SendGridTo `json:"to"`
}

type SendGridTo struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

type SendGridFrom struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

type SendGridContent struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func NewSendGridEmailService(cfg *config.Config) *EmailService {
	return &EmailService{
		config: cfg,
		client: &http.Client{Timeout: 30 * time.Second},
	}
}

func (e *EmailService) SendEmail(to, subject, body string) error {
	if e.config.SendGridAPIKey == "" {
		return e.mockSendEmail(to, subject, body)
	}

	url := fmt.Sprintf("%s/mail/send", e.config.SendGridURL)

	emailData := SendGridEmail{
		Personalizations: []SendGridPersonalization{
			{
				To: []SendGridTo{
					{
						Email: to,
					},
				},
			},
		},
		From: SendGridFrom{
			Email: e.config.FromEmail,
			Name:  e.config.FromName,
		},
		Subject: subject,
		Content: []SendGridContent{
			{
				Type:  "text/plain",
				Value: body,
			},
		},
	}

	jsonData, err := json.Marshal(emailData)
	if err != nil {
		return fmt.Errorf("failed to marshal email data: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+e.config.SendGridAPIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := e.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("sendgrid API error: %d - %s", resp.StatusCode, string(body))
	}

	fmt.Printf("Successfully sent email to: %s\nSubject: %s\nBody: %s\n", to, subject, body)
	return nil
}

func (e *EmailService) mockSendEmail(to, subject, body string) error {
	fmt.Printf("SendGrid API key not configured. Using mock implementation.\n")
	fmt.Printf("Sending email to: %s\nSubject: %s\nBody: %s\n", to, subject, body)
	return nil
}
