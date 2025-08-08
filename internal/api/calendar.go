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

type CalendarService struct {
	config *config.Config
	client *http.Client
}

type CalendarEvent struct {
	Summary            string                 `json:"summary"`
	Description        string                 `json:"description,omitempty"`
	Start              CalendarDateTime       `json:"start"`
	End                CalendarDateTime       `json:"end"`
	Attendees          []CalendarAttendee     `json:"attendees,omitempty"`
	Reminders          CalendarReminders      `json:"reminders,omitempty"`
	Location           string                 `json:"location,omitempty"`
	ExtendedProperties map[string]interface{} `json:"extendedProperties,omitempty"`
}

type CalendarDateTime struct {
	DateTime string `json:"dateTime"`
	TimeZone string `json:"timeZone"`
}

type CalendarAttendee struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

type CalendarReminders struct {
	UseDefault bool `json:"useDefault"`
}

type CalendarEventsResponse struct {
	Items []CalendarEvent `json:"items"`
}

func NewGoogleCalendarService(cfg *config.Config) *CalendarService {
	return &CalendarService{
		config: cfg,
		client: &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *CalendarService) ScheduleMeeting(attendees []string, startTime time.Time, duration time.Duration, title string) error {
	if c.config.GoogleCalendarAPIKey == "" {
		return c.mockScheduleMeeting(attendees, startTime, duration, title)
	}

	url := fmt.Sprintf("%s/calendars/%s/events", c.config.GoogleCalendarURL, c.config.CalendarID)

	// Convert attendees to proper format
	var calendarAttendees []CalendarAttendee
	for _, email := range attendees {
		calendarAttendees = append(calendarAttendees, CalendarAttendee{
			Email: email,
		})
	}

	event := CalendarEvent{
		Summary: title,
		Start: CalendarDateTime{
			DateTime: startTime.Format(time.RFC3339),
			TimeZone: c.config.TimeZone,
		},
		End: CalendarDateTime{
			DateTime: startTime.Add(duration).Format(time.RFC3339),
			TimeZone: c.config.TimeZone,
		},
		Attendees: calendarAttendees,
		Reminders: CalendarReminders{
			UseDefault: true,
		},
	}

	jsonData, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.config.GoogleCalendarAPIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to schedule meeting: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("calendar API error: %d - %s", resp.StatusCode, string(body))
	}

	fmt.Printf("Successfully scheduled meeting:\nTitle: %s\nAttendees: %v\nStart: %s\nDuration: %v\n",
		title, attendees, startTime.Format("2006-01-02 15:04:05"), duration)

	return nil
}

func (c *CalendarService) GetUpcomingEvents() ([]Event, error) {
	if c.config.GoogleCalendarAPIKey == "" {
		return c.mockGetUpcomingEvents()
	}

	url := fmt.Sprintf("%s/calendars/%s/events", c.config.GoogleCalendarURL, c.config.CalendarID)

	// Get events for the next 7 days
	now := time.Now()
	endTime := now.AddDate(0, 0, 7)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	q := req.URL.Query()
	q.Add("timeMin", now.Format(time.RFC3339))
	q.Add("timeMax", endTime.Format(time.RFC3339))
	q.Add("singleEvents", "true")
	q.Add("orderBy", "startTime")
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Authorization", "Bearer "+c.config.GoogleCalendarAPIKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get events: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("calendar API error: %d - %s", resp.StatusCode, string(body))
	}

	var response CalendarEventsResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	// Convert to internal Event format
	var events []Event
	for _, calEvent := range response.Items {
		startTime, _ := time.Parse(time.RFC3339, calEvent.Start.DateTime)
		endTime, _ := time.Parse(time.RFC3339, calEvent.End.DateTime)

		var attendees []string
		for _, attendee := range calEvent.Attendees {
			attendees = append(attendees, attendee.Email)
		}

		events = append(events, Event{
			Title:     calEvent.Summary,
			StartTime: startTime,
			EndTime:   endTime,
			Attendees: attendees,
		})
	}

	return events, nil
}

func (c *CalendarService) mockScheduleMeeting(attendees []string, startTime time.Time, duration time.Duration, title string) error {
	fmt.Printf("Google Calendar API key not configured. Using mock implementation.\n")
	fmt.Printf("Scheduling meeting:\nTitle: %s\nAttendees: %v\nStart: %s\nDuration: %v\n",
		title, attendees, startTime.Format("2006-01-02 15:04:05"), duration)
	return nil
}

func (c *CalendarService) mockGetUpcomingEvents() ([]Event, error) {
	fmt.Printf("Google Calendar API key not configured. Returning mock events.\n")

	// Return some mock events
	now := time.Now()
	return []Event{
		{
			Title:     "Team Standup",
			StartTime: now.Add(1 * time.Hour),
			EndTime:   now.Add(1*time.Hour + 30*time.Minute),
			Attendees: []string{"team@company.com"},
		},
		{
			Title:     "Client Meeting",
			StartTime: now.Add(3 * time.Hour),
			EndTime:   now.Add(3*time.Hour + 1*time.Hour),
			Attendees: []string{"client@company.com"},
		},
	}, nil
}

type Event struct {
	Title     string
	StartTime time.Time
	EndTime   time.Time
	Attendees []string
}
