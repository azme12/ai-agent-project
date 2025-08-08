package agent

import (
	"fmt"
	"strconv"
	"strings"
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

	// Send daily reminder at configured time
	if s.shouldSendDailyReminder() {
		s.sendDailyReminder()
	}

	// Check for meetings starting soon
	s.checkUpcomingMeetings()

	// Process recurring tasks
	s.processRecurringTasks()
}

func (s *Scheduler) shouldSendDailyReminder() bool {
	now := time.Now()

	// Parse the configured reminder time
	reminderTime := s.config.DailyReminderTime
	if reminderTime == "" {
		reminderTime = "09:00" // Default to 9 AM
	}

	// Parse hour and minute from reminder time
	parts := strings.Split(reminderTime, ":")
	if len(parts) != 2 {
		return false
	}

	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		return false
	}

	minute, err := strconv.Atoi(parts[1])
	if err != nil {
		return false
	}

	// Check if it's the configured time
	return now.Hour() == hour && now.Minute() == minute
}

func (s *Scheduler) sendDailyReminder() {
	s.logger.Info("Sending daily reminder")

	// Get today's tasks
	tasks := s.getTodaysTasks()

	// Get today's meetings
	meetings := s.getTodaysMeetings()

	reminderBody := "Good morning! Here's the daily summary:\n\n"

	if len(tasks) > 0 {
		reminderBody += "Today's Tasks:\n"
		for i, task := range tasks {
			reminderBody += fmt.Sprintf("%d. %s\n", i+1, task)
		}
		reminderBody += "\n"
	}

	if len(meetings) > 0 {
		reminderBody += " Today's Meetings:\n"
		for _, meeting := range meetings {
			reminderBody += fmt.Sprintf("• %s at %s\n", meeting.Title, meeting.StartTime.Format("15:04"))
		}
		reminderBody += "\n"
	}

	reminderBody += "Have a productive day! "

	// Send reminder email
	if s.email != nil {
		err := s.email.SendEmail(
			s.config.UserEmail,
			"Daily Summary - AI Assistant",
			reminderBody,
		)
		if err != nil {
			s.logger.Error("Failed to send daily reminder", "error", err)
		}
	}
}

func (s *Scheduler) processRecurringTasks() {
	s.logger.Debug("Processing recurring tasks")

	// Check for weekly tasks on Monday
	now := time.Now()
	if now.Weekday() == time.Monday && now.Hour() == 9 && now.Minute() == 0 {
		s.sendWeeklySummary()
	}

	// Check for monthly tasks on the 1st
	if now.Day() == 1 && now.Hour() == 9 && now.Minute() == 0 {
		s.sendMonthlySummary()
	}
}

func (s *Scheduler) sendWeeklySummary() {
	s.logger.Info("Sending weekly summary")

	summaryBody := " Weekly Summary\n\n"
	summaryBody += "This week's accomplishments:\n"
	summaryBody += "• Completed project milestones\n"
	summaryBody += "• Scheduled team meetings\n"
	summaryBody += "• Responded to important emails\n\n"
	summaryBody += "Next week's priorities:\n"
	summaryBody += "• Review pending tasks\n"
	summaryBody += "• Plan upcoming meetings\n"
	summaryBody += "• Follow up on action items"

	if s.email != nil {
		err := s.email.SendEmail(
			s.config.UserEmail,
			"Weekly Summary - AI Assistant",
			summaryBody,
		)
		if err != nil {
			s.logger.Error("Failed to send weekly summary", "error", err)
		}
	}
}

func (s *Scheduler) sendMonthlySummary() {
	s.logger.Info("Sending monthly summary")

	summaryBody := " Monthly Summary\n\n"
	summaryBody += "This month's key achievements:\n"
	summaryBody += "• Completed major project phases\n"
	summaryBody += "• Attended important meetings\n"
	summaryBody += "• Maintained communication with stakeholders\n\n"
	summaryBody += "Next month's focus areas:\n"
	summaryBody += "• Strategic planning\n"
	summaryBody += "• Team coordination\n"
	summaryBody += "• Performance review"

	if s.email != nil {
		err := s.email.SendEmail(
			s.config.UserEmail,
			"Monthly Summary - AI Assistant",
			summaryBody,
		)
		if err != nil {
			s.logger.Error("Failed to send monthly summary", "error", err)
		}
	}
}

func (s *Scheduler) getTodaysTasks() []string {
	// TODO: Integrate with a real task management system
	return []string{
		"Review pending emails",
		"Check calendar for today's meetings",
		"Update project status",
		"Follow up on action items",
		"Prepare for tomorrow's meetings",
	}
}

func (s *Scheduler) getTodaysMeetings() []api.Event {
	if s.calendar != nil {
		events, err := s.calendar.GetUpcomingEvents()
		if err != nil {
			s.logger.Error("Failed to get today's meetings", "error", err)
			return []api.Event{}
		}

		// Filter for today's events
		var todaysEvents []api.Event
		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		tomorrow := today.AddDate(0, 0, 1)

		for _, event := range events {
			if event.StartTime.After(today) && event.StartTime.Before(tomorrow) {
				todaysEvents = append(todaysEvents, event)
			}
		}

		return todaysEvents
	}

	return []api.Event{}
}

func (s *Scheduler) checkUpcomingMeetings() {
	s.logger.Debug("Checking upcoming meetings")

	if s.calendar != nil {
		events, err := s.calendar.GetUpcomingEvents()
		if err != nil {
			s.logger.Error("Failed to get upcoming events", "error", err)
			return
		}

		now := time.Now()
		reminderMinutes := s.config.MeetingReminderMinutes
		if reminderMinutes == 0 {
			reminderMinutes = 15 // Default to 15 minutes
		}

		for _, event := range events {
			// Check if meeting starts within the reminder window
			timeUntilMeeting := event.StartTime.Sub(now)
			if timeUntilMeeting <= time.Duration(reminderMinutes)*time.Minute && timeUntilMeeting > 0 {
				s.sendMeetingReminder(event)
			}
		}
	}
}

func (s *Scheduler) sendMeetingReminder(event api.Event) {
	s.logger.Info("Sending meeting reminder", "meeting", event.Title)

	reminderBody := fmt.Sprintf(" Meeting Reminder\n\n")
	reminderBody += fmt.Sprintf("Meeting: %s\n", event.Title)
	reminderBody += fmt.Sprintf("Time: %s\n", event.StartTime.Format("Monday, January 2, 2006 at 15:04"))
	reminderBody += fmt.Sprintf("Duration: %s\n", event.EndTime.Sub(event.StartTime).String())

	if len(event.Attendees) > 0 {
		reminderBody += fmt.Sprintf("Attendees: %s\n", strings.Join(event.Attendees, ", "))
	}

	reminderBody += "\nPlease join on time! "

	if s.email != nil {
		err := s.email.SendEmail(
			s.config.UserEmail,
			"Meeting Reminder - AI Assistant",
			reminderBody,
		)
		if err != nil {
			s.logger.Error("Failed to send meeting reminder", "error", err)
		}
	}
}
