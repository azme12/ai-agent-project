# AI Agent API Documentation

Complete API documentation for the AI Agent Executive Assistant automation service.

## Table of Contents

- [Overview](#overview)
- [Base URL](#base-url)
- [Authentication](#authentication)
- [Endpoints](#endpoints)
  - [Health & Status](#health--status)
  - [Meeting Scheduling](#meeting-scheduling)
  - [Email Automation](#email-automation)
  - [Natural Language Processing](#natural-language-processing)
- [Error Codes](#error-codes)
- [Request/Response Formats](#requestresponse-formats)
- [Examples](#examples)
- [Environment Variables](#environment-variables)
- [Rate Limiting](#rate-limiting)
- [Testing](#testing)

## Overview

The AI Agent API provides RESTful endpoints for automating executive assistant tasks including:

- **Meeting Scheduling**: Natural language meeting creation using Google Calendar
- **Email Automation**: Intelligent email sending using SendGrid
- **Natural Language Processing**: AI-powered command processing using Google Gemini
- **Health Monitoring**: Service status and health checks

## Base URL

```
Local Development: http://localhost:8080
Production: https://your-domain.com
```

## Authentication

Currently, the API does not require authentication. However, it relies on external API keys for:

- Google Calendar API
- SendGrid API  
- Google Gemini API

These keys should be configured via environment variables.

## Endpoints

### Health & Status

#### GET /health

Returns the health status of the AI Agent service.

**Request:**
```http
GET /health
```

**Response:**
```json
{
  "status": "healthy",
  "service": "ai-agent",
  "version": "1.0.0"
}
```

**Status Codes:**
- `200 OK`: Service is healthy

---

#### GET /status

Returns detailed service status including all available endpoints.

**Request:**
```http
GET /status
```

**Response:**
```json
{
  "status": "running",
  "service": "ai-agent",
  "endpoints": [
    "GET /health",
    "POST /schedule",
    "POST /email",
    "POST /nlp",
    "GET /status"
  ]
}
```

**Status Codes:**
- `200 OK`: Service is running

### Meeting Scheduling

#### POST /schedule

Schedules a meeting using natural language processing.

**Request:**
```http
POST /schedule
Content-Type: application/json

{
  "task": "Schedule a meeting with john@example.com tomorrow at 2 PM about project review"
}
```

**Request Body:**
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `task` | string | Yes | Natural language description of the meeting to schedule |

**Response:**
```json
{
  "status": "success",
  "message": "Task processed successfully",
  "task": "Schedule a meeting with john@example.com tomorrow at 2 PM about project review"
}
```

**Status Codes:**
- `200 OK`: Meeting scheduled successfully
- `400 Bad Request`: Invalid request body or missing task
- `500 Internal Server Error`: Service error or API integration failure

**Features:**
- Natural language parsing for meeting details
- Automatic attendee extraction from email addresses
- Time parsing (tomorrow, next week, specific dates)
- Topic extraction for meeting description
- Conflict detection and resolution
- Automatic calendar integration

**Example Tasks:**
```
"Schedule a meeting with john@example.com tomorrow at 2 PM about project review"
"Book a 1-hour meeting with the team next Monday at 10 AM"
"Set up a call with client@company.com on Friday at 3 PM to discuss proposal"
"Schedule a 30-minute meeting with sarah@example.com today at 4 PM"
```

### Email Automation

#### POST /email

Sends emails using natural language processing.

**Request:**
```http
POST /email
Content-Type: application/json

{
  "task": "Send email to client@example.com saying thank you for the meeting and we'll follow up next week"
}
```

**Request Body:**
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `task` | string | Yes | Natural language description of the email to send |

**Response:**
```json
{
  "status": "success",
  "message": "Email task processed successfully",
  "task": "Send email to client@example.com saying thank you for the meeting and we'll follow up next week"
}
```

**Status Codes:**
- `200 OK`: Email sent successfully
- `400 Bad Request`: Invalid request body or missing task
- `500 Internal Server Error`: Service error or SendGrid API failure

**Features:**
- Natural language parsing for email content
- Automatic recipient extraction from email addresses
- Content generation based on context
- Professional email formatting
- SendGrid integration for reliable delivery
- Follow-up email automation

**Example Tasks:**
```
"Send email to client@example.com saying thank you for the meeting and we'll follow up next week"
"Email john@example.com about the project status update"
"Send a follow-up email to sarah@company.com regarding yesterday's discussion"
"Email the team about the new project timeline"
```

### Natural Language Processing

#### POST /nlp

Processes natural language commands using Google Gemini AI.

**Request:**
```http
POST /nlp
Content-Type: application/json

{
  "command": "What meetings do I have today?"
}
```

**Request Body:**
| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `command` | string | Yes | Natural language command or query |

**Response:**
```json
{
  "status": "success",
  "response": "I'll check your calendar for today's meetings. You have 2 meetings scheduled: 1) Team Standup at 10:00 AM, 2) Client Review at 2:00 PM.",
  "command": "What meetings do I have today?"
}
```

**Status Codes:**
- `200 OK`: Command processed successfully
- `400 Bad Request`: Invalid request body or missing command
- `500 Internal Server Error`: Service error or Gemini API failure

**Capabilities:**
- Calendar queries (meetings, availability, schedule)
- Email management (drafts, sent emails, follow-ups)
- Task management and reminders
- General assistant queries
- Context-aware responses
- Multi-turn conversations

**Example Commands:**
```
"What meetings do I have today?"
"Check my availability for tomorrow"
"What emails did I send yesterday?"
"Remind me to call the client at 3 PM"
"What's on my schedule for next week?"
"Send a reminder to the team about the deadline"
```

## Error Codes

| Status Code | Description | Common Causes |
|-------------|-------------|---------------|
| `200` | Success | Request processed successfully |
| `400` | Bad Request | Missing required fields, invalid JSON, empty task/command |
| `405` | Method Not Allowed | Using wrong HTTP method (GET instead of POST, etc.) |
| `500` | Internal Server Error | Service error, API integration failure, configuration issues |

**Error Response Format:**
```json
{
  "error": "Error message description",
  "status": "error",
  "timestamp": "2024-01-01T00:00:00Z"
}
```

## Request/Response Formats

### Common Request Headers
```http
Content-Type: application/json
Accept: application/json
```

### Common Response Headers
```http
Content-Type: application/json
```

### Success Response Format
All successful responses follow this format:
```json
{
  "status": "success",
  "message": "Description of what was accomplished",
  "data": {} // Optional additional data
}
```

### Error Response Format
All error responses follow this format:
```json
{
  "status": "error",
  "error": "Description of the error",
  "code": "ERROR_CODE" // Optional error code
}
```

## Examples

### Complete Workflow Example

1. **Check Service Health**
```bash
curl -X GET http://localhost:8080/health
```

2. **Schedule a Meeting**
```bash
curl -X POST http://localhost:8080/schedule \
  -H "Content-Type: application/json" \
  -d '{
    "task": "Schedule a meeting with john@example.com tomorrow at 2 PM about project review"
  }'
```

3. **Send Follow-up Email**
```bash
curl -X POST http://localhost:8080/email \
  -H "Content-Type: application/json" \
  -d '{
    "task": "Send email to john@example.com thanking him for the meeting and confirming next steps"
  }'
```

4. **Query Calendar**
```bash
curl -X POST http://localhost:8080/nlp \
  -H "Content-Type: application/json" \
  -d '{
    "command": "What meetings do I have tomorrow?"
  }'
```

### JavaScript/Node.js Example

```javascript
const axios = require('axios');

const API_BASE_URL = 'http://localhost:8080';

// Health check
async function checkHealth() {
  try {
    const response = await axios.get(`${API_BASE_URL}/health`);
    console.log('Service status:', response.data);
  } catch (error) {
    console.error('Health check failed:', error.message);
  }
}

// Schedule meeting
async function scheduleMeeting(task) {
  try {
    const response = await axios.post(`${API_BASE_URL}/schedule`, {
      task: task
    });
    console.log('Meeting scheduled:', response.data);
  } catch (error) {
    console.error('Failed to schedule meeting:', error.message);
  }
}

// Send email
async function sendEmail(task) {
  try {
    const response = await axios.post(`${API_BASE_URL}/email`, {
      task: task
    });
    console.log('Email sent:', response.data);
  } catch (error) {
    console.error('Failed to send email:', error.message);
  }
}

// Process NLP command
async function processCommand(command) {
  try {
    const response = await axios.post(`${API_BASE_URL}/nlp`, {
      command: command
    });
    console.log('AI response:', response.data);
  } catch (error) {
    console.error('Failed to process command:', error.message);
  }
}

// Usage examples
checkHealth();
scheduleMeeting('Schedule a meeting with team@company.com tomorrow at 10 AM');
sendEmail('Send email to client@example.com about project update');
processCommand('What meetings do I have today?');
```

### Python Example

```python
import requests
import json

API_BASE_URL = 'http://localhost:8080'

def check_health():
    """Check service health"""
    try:
        response = requests.get(f'{API_BASE_URL}/health')
        response.raise_for_status()
        return response.json()
    except requests.exceptions.RequestException as e:
        print(f'Health check failed: {e}')
        return None

def schedule_meeting(task):
    """Schedule a meeting"""
    try:
        response = requests.post(
            f'{API_BASE_URL}/schedule',
            json={'task': task},
            headers={'Content-Type': 'application/json'}
        )
        response.raise_for_status()
        return response.json()
    except requests.exceptions.RequestException as e:
        print(f'Failed to schedule meeting: {e}')
        return None

def send_email(task):
    """Send an email"""
    try:
        response = requests.post(
            f'{API_BASE_URL}/email',
            json={'task': task},
            headers={'Content-Type': 'application/json'}
        )
        response.raise_for_status()
        return response.json()
    except requests.exceptions.RequestException as e:
        print(f'Failed to send email: {e}')
        return None

def process_command(command):
    """Process NLP command"""
    try:
        response = requests.post(
            f'{API_BASE_URL}/nlp',
            json={'command': command},
            headers={'Content-Type': 'application/json'}
        )
        response.raise_for_status()
        return response.json()
    except requests.exceptions.RequestException as e:
        print(f'Failed to process command: {e}')
        return None

# Usage examples
if __name__ == '__main__':
    # Check health
    health = check_health()
    print('Health status:', health)
    
    # Schedule meeting
    meeting_result = schedule_meeting('Schedule a meeting with team@company.com tomorrow at 10 AM')
    print('Meeting result:', meeting_result)
    
    # Send email
    email_result = send_email('Send email to client@example.com about project update')
    print('Email result:', email_result)
    
    # Process command
    command_result = process_command('What meetings do I have today?')
    print('Command result:', command_result)
```

## Environment Variables

The AI Agent service uses the following environment variables:

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `GOOGLE_CALENDAR_API_KEY` | Google Calendar API key | "" | Yes* |
| `SENDGRID_API_KEY` | SendGrid API key | "" | Yes* |
| `GEMINI_API_KEY` | Google Gemini API key | "" | Yes* |
| `SERVER_PORT` | HTTP server port | "8080" | No |
| `LOG_LEVEL` | Logging level | "info" | No |
| `GOOGLE_CALENDAR_URL` | Google Calendar API URL | "https://www.googleapis.com/calendar/v3" | No |
| `SENDGRID_URL` | SendGrid API URL | "https://api.sendgrid.com/v3" | No |
| `GEMINI_URL` | Gemini API URL | "https://generativelanguage.googleapis.com/v1beta" | No |
| `FROM_EMAIL` | Sender email address | "ai-assistant@yourdomain.com" | No |
| `FROM_NAME` | Sender name | "AI Assistant" | No |
| `CALENDAR_ID` | Google Calendar ID | "primary" | No |
| `TIMEZONE` | Timezone for events | "UTC" | No |
| `DAILY_REMINDER_TIME` | Daily reminder time | "09:00" | No |
| `MEETING_REMINDER_MINUTES` | Meeting reminder minutes | 15 | No |

*Required for full functionality. Without API keys, the service runs in mock mode.

## Rate Limiting

Currently, the API does not implement rate limiting. However, it's recommended to:

- Limit requests to reasonable frequencies
- Implement exponential backoff for retries
- Monitor API usage for external services (Google Calendar, SendGrid, Gemini)

## Testing

### Using cURL

```bash
# Health check
curl -X GET http://localhost:8080/health

# Schedule meeting
curl -X POST http://localhost:8080/schedule \
  -H "Content-Type: application/json" \
  -d '{"task": "Schedule meeting with test@example.com tomorrow at 2 PM"}'

# Send email
curl -X POST http://localhost:8080/email \
  -H "Content-Type: application/json" \
  -d '{"task": "Send email to test@example.com about test"}'

# NLP command
curl -X POST http://localhost:8080/nlp \
  -H "Content-Type: application/json" \
  -d '{"command": "What is the weather today?"}'
```

### Using Postman

1. Import the provided Postman collection: `AI_Agent_API_Postman_Collection.json`
2. Import the environment: `AI_Agent_API_Postman_Environment.json`
3. Update the environment variables with your API keys
4. Start testing the endpoints

### Automated Testing

Use the provided test script:

```bash
# Run the test suite
./test.sh
```

## Support

For issues, questions, or contributions:

1. Check the service logs for detailed error information
2. Verify all required environment variables are set
3. Ensure external API keys are valid and have proper permissions
4. Test with the health endpoint first to verify service status

---

**API Version:** 1.0.0  
**Last Updated:** January 2024  
**Service:** AI Agent Executive Assistant
