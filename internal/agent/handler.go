package agent

import (
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

	// Route based on command type
	if strings.Contains(strings.ToLower(task), "schedule") || strings.Contains(strings.ToLower(task), "meeting") {
		return h.handleScheduleTask(task)
	} else if strings.Contains(strings.ToLower(task), "email") || strings.Contains(strings.ToLower(task), "send") {
		return h.handleEmailTask(task)
	} else {
		h.logger.Info("Unknown task type, using NLP response", "task", task)
	}

	return nil
}

func (h *Handler) handleScheduleTask(task string) error {
	h.logger.Info("Handling schedule task", "task", task)

	// TODO: Parse meeting details from task
	// For now, use placeholder values
	attendees := []string{"example@email.com"}
	startTime := time.Now().Add(1 * time.Hour)
	duration := 30 * time.Minute
	title := "Meeting from AI Agent"

	return h.calendar.ScheduleMeeting(attendees, startTime, duration, title)
}

func (h *Handler) handleEmailTask(task string) error {
	h.logger.Info("Handling email task", "task", task)

	// TODO: Parse email details from task
	// For now, use placeholder values
	to := "recipient@email.com"
	subject := "Email from AI Agent"
	body := "This is an automated email from the AI Agent."

	return h.email.SendEmail(to, subject, body)
}
