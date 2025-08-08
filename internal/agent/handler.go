package agent

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/azme12/ai-agent-project/internal/api"
	"github.com/azme12/ai-agent-project/internal/config"
	"github.com/azme12/ai-agent-project/pkg/logger"
)

type Handler struct {
	config   *config.Config
	logger   *logger.Logger
	calendar *api.CalendarService
	email    *api.EmailService
	nlp      *api.GeminiService
}

type TaskRequest struct {
	Type        string    `json:"type"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Attendees   []string  `json:"attendees"`
	StartTime   time.Time `json:"start_time"`
	Duration    int       `json:"duration_minutes"`
	To          string    `json:"to"`
	Subject     string    `json:"subject"`
	Body        string    `json:"body"`
}

func NewHandler(cfg *config.Config, log *logger.Logger, cal *api.CalendarService, em *api.EmailService, nlp *api.GeminiService) *Handler {
	return &Handler{
		config:   cfg,
		logger:   log,
		calendar: cal,
		email:    em,
		nlp:      nlp,
	}
}

func (h *Handler) ProcessTask(task string) error {
	h.logger.Info("Processing task", "task", task)

	// Use NLP to understand the command
	response, err := h.nlp.ProcessCommand(task)
	if err != nil {
		h.logger.Error("Failed to process command with NLP", "error", err)
		return err
	}

	h.logger.Info("NLP response", "response", response)

	// Parse the task to determine action type
	taskRequest, err := h.parseTask(task)
	if err != nil {
		h.logger.Error("Failed to parse task", "error", err)
		return err
	}

	// Route based on task type
	switch taskRequest.Type {
	case "schedule":
		return h.handleScheduleTask(taskRequest)
	case "email":
		return h.handleEmailTask(taskRequest)
	case "reminder":
		return h.handleReminderTask(taskRequest)
	default:
		h.logger.Info("Unknown task type, using NLP response", "task", task)
		return nil
	}
}

func (h *Handler) parseTask(task string) (*TaskRequest, error) {
	lowerTask := strings.ToLower(task)

	// Initialize task request
	taskRequest := &TaskRequest{
		Duration: 30, // Default 30 minutes
	}

	// Determine task type
	if strings.Contains(lowerTask, "schedule") || strings.Contains(lowerTask, "meeting") {
		taskRequest.Type = "schedule"
		return h.parseScheduleTask(task, taskRequest)
	} else if strings.Contains(lowerTask, "email") || strings.Contains(lowerTask, "send") {
		taskRequest.Type = "email"
		return h.parseEmailTask(task, taskRequest)
	} else if strings.Contains(lowerTask, "remind") || strings.Contains(lowerTask, "reminder") {
		taskRequest.Type = "reminder"
		return h.parseReminderTask(task, taskRequest)
	}

	// Default to general task
	taskRequest.Type = "general"
	taskRequest.Title = task
	return taskRequest, nil
}

func (h *Handler) parseScheduleTask(task string, req *TaskRequest) (*TaskRequest, error) {
	// Extract attendees (email addresses)
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	emails := emailRegex.FindAllString(task, -1)
	req.Attendees = emails

	// Extract time information
	req.StartTime = h.extractTime(task)

	// Extract duration
	if strings.Contains(task, "hour") || strings.Contains(task, "hr") {
		req.Duration = 60
	} else if strings.Contains(task, "30") {
		req.Duration = 30
	} else if strings.Contains(task, "15") {
		req.Duration = 15
	}

	// Extract title
	req.Title = h.extractTitle(task)
	if req.Title == "" {
		req.Title = "Meeting scheduled by AI Assistant"
	}

	return req, nil
}

func (h *Handler) parseEmailTask(task string, req *TaskRequest) (*TaskRequest, error) {
	// Extract recipient
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	emails := emailRegex.FindAllString(task, -1)
	if len(emails) > 0 {
		req.To = emails[0]
	}

	// Extract subject
	req.Subject = h.extractSubject(task)
	if req.Subject == "" {
		req.Subject = "Message from AI Assistant"
	}

	// Extract body
	req.Body = h.extractBody(task)
	if req.Body == "" {
		req.Body = "This is an automated message from the AI Assistant."
	}

	return req, nil
}

func (h *Handler) parseReminderTask(task string, req *TaskRequest) (*TaskRequest, error) {
	req.Title = h.extractTitle(task)
	if req.Title == "" {
		req.Title = "Reminder from AI Assistant"
	}

	req.StartTime = h.extractTime(task)
	return req, nil
}

func (h *Handler) extractTime(task string) time.Time {
	now := time.Now()

	// Look for time patterns
	if strings.Contains(task, "tomorrow") {
		return now.AddDate(0, 0, 1)
	} else if strings.Contains(task, "next week") {
		return now.AddDate(0, 0, 7)
	} else if strings.Contains(task, "today") {
		return now
	}

	// Look for specific times
	timeRegex := regexp.MustCompile(`(\d{1,2}):?(\d{2})?\s*(am|pm)?`)
	matches := timeRegex.FindStringSubmatch(task)
	if len(matches) > 0 {
		hour := 9 // Default to 9 AM
		if len(matches) > 1 {
			// Parse hour
			if h, err := fmt.Sscanf(matches[1], "%d", &hour); err == nil && h > 0 {
				// Handle AM/PM
				if len(matches) > 3 && strings.ToLower(matches[3]) == "pm" && hour != 12 {
					hour += 12
				} else if len(matches) > 3 && strings.ToLower(matches[3]) == "am" && hour == 12 {
					hour = 0
				}
			}
		}

		// Set time for today
		return time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, now.Location())
	}

	// Default to 1 hour from now
	return now.Add(1 * time.Hour)
}

func (h *Handler) extractTitle(task string) string {
	// Simple title extraction - look for quoted text or key phrases
	quoteRegex := regexp.MustCompile(`"([^"]+)"`)
	matches := quoteRegex.FindStringSubmatch(task)
	if len(matches) > 1 {
		return matches[1]
	}

	// Look for "about" or "regarding" phrases
	if strings.Contains(task, "about") {
		parts := strings.Split(task, "about")
		if len(parts) > 1 {
			return strings.TrimSpace(parts[1])
		}
	}

	return ""
}

func (h *Handler) extractSubject(task string) string {
	// Similar to title extraction
	return h.extractTitle(task)
}

func (h *Handler) extractBody(task string) string {
	// Extract body content after key phrases
	if strings.Contains(task, "saying") {
		parts := strings.Split(task, "saying")
		if len(parts) > 1 {
			return strings.TrimSpace(parts[1])
		}
	}

	if strings.Contains(task, "message") {
		parts := strings.Split(task, "message")
		if len(parts) > 1 {
			return strings.TrimSpace(parts[1])
		}
	}

	return ""
}

func (h *Handler) handleScheduleTask(req *TaskRequest) error {
	h.logger.Info("Handling schedule task", "title", req.Title, "attendees", req.Attendees, "startTime", req.StartTime)

	if len(req.Attendees) == 0 {
		req.Attendees = []string{h.config.UserEmail}
	}

	duration := time.Duration(req.Duration) * time.Minute
	return h.calendar.ScheduleMeeting(req.Attendees, req.StartTime, duration, req.Title)
}

func (h *Handler) handleEmailTask(req *TaskRequest) error {
	h.logger.Info("Handling email task", "to", req.To, "subject", req.Subject)

	if req.To == "" {
		req.To = h.config.UserEmail
	}

	return h.email.SendEmail(req.To, req.Subject, req.Body)
}

func (h *Handler) handleReminderTask(req *TaskRequest) error {
	h.logger.Info("Handling reminder task", "title", req.Title, "time", req.StartTime)

	// Send reminder email
	subject := "Reminder: " + req.Title
	body := fmt.Sprintf("This is a reminder for: %s\nScheduled for: %s", req.Title, req.StartTime.Format("2006-01-02 15:04:05"))

	return h.email.SendEmail(h.config.UserEmail, subject, body)
}
