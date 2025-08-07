package agent

import (
	"fmt"
	"time"

	"github.com/azme12/ai-agent-project/internal/api"
	"github.com/azme12/ai-agent-project/internal/config"
	"github.com/azme12/ai-agent-project/pkg/logger"
)

type Scheduler struct {
	config   *config.Config
	logger   *logger.Logger
	stopCh   chan struct{}
	calendar *api.CalendarService
	email    *api.EmailService
}

func NewScheduler(cfg *config.Config, log *logger.Logger) *Scheduler {
	return &Scheduler{
		config: cfg,
		logger: log,
		stopCh: make(chan struct{}),
	}
}

func (s *Scheduler) Start() error {
	s.logger.Info("Starting scheduler")

	go s.run()

	return nil
}

func (s *Scheduler) Stop() error {
	s.logger.Info("Stopping scheduler")
	close(s.stopCh)
	return nil
}

func (s *Scheduler) run() {
	// Check for scheduled tasks every minute
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	// Also run immediately on startup
	s.checkScheduledTasks()

	for {
		select {
		case <-ticker.C:
			s.checkScheduledTasks()

		case <-s.stopCh:
			s.logger.Info("Scheduler stopped")
			return
		}
	}
}

func (s *Scheduler) checkScheduledTasks() {
	s.logger.Debug("Checking scheduled tasks")

	// TODO: Implement actual scheduled task checking
	// This could include:
	// - Checking for upcoming meetings and sending reminders
	// - Sending daily task summaries
	// - Processing recurring tasks

	// Example: Send daily reminder at 9 AM
	now := time.Now()
	if now.Hour() == 9 && now.Minute() == 0 {
		s.sendDailyReminder()
	}

	// Example: Check for meetings starting in 15 minutes
	s.checkUpcomingMeetings()
}

func (s *Scheduler) sendDailyReminder() {
	s.logger.Info("Sending daily reminder")

	// TODO: Get actual tasks from a task management system
	tasks := []string{
		"Review pending emails",
		"Check calendar for today's meetings",
		"Update project status",
	}

	reminderBody := "Here are your tasks for today:\n"
	for i, task := range tasks {
		reminderBody += fmt.Sprintf("%d. %s\n", i+1, task)
	}

	// Send reminder email
	if s.email != nil {
		err := s.email.SendEmail(
			"user@example.com",
			"Daily Task Reminder",
			reminderBody,
		)
		if err != nil {
			s.logger.Error("Failed to send daily reminder", "error", err)
		}
	}
}

func (s *Scheduler) checkUpcomingMeetings() {
	s.logger.Debug("Checking upcoming meetings")

	// TODO: Get actual upcoming events from calendar
	if s.calendar != nil {
		events, err := s.calendar.GetUpcomingEvents()
		if err != nil {
			s.logger.Error("Failed to get upcoming events", "error", err)
			return
		}

		now := time.Now()
		for _, event := range events {
			// Check if meeting starts in 15 minutes
			if event.StartTime.Sub(now) <= 15*time.Minute && event.StartTime.Sub(now) > 0 {
				s.sendMeetingReminder(event)
			}
		}
	}
}

func (s *Scheduler) sendMeetingReminder(event api.Event) {
	s.logger.Info("Sending meeting reminder", "meeting", event.Title)

	reminderBody := fmt.Sprintf("Reminder: You have a meeting '%s' starting at %s",
		event.Title, event.StartTime.Format("15:04"))

	if s.email != nil {
		err := s.email.SendEmail(
			"user@example.com",
			"Meeting Reminder",
			reminderBody,
		)
		if err != nil {
			s.logger.Error("Failed to send meeting reminder", "error", err)
		}
	}
}
