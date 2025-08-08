# AI Agent Project - Executive Assistant Automation

An intelligent AI agent that automates executive assistant tasks including scheduling meetings, sending emails, and processing natural language commands.

##  Project Overview

This AI agent serves as an **Executive Assistant in a Tech Startup**, automating high-impact, repetitive tasks to increase productivity and efficiency. The agent uses advanced natural language processing to understand user commands and takes proactive actions to manage your schedule, communications, and tasks.

###  Core Features

- **📅 Smart Meeting Scheduling**: Automatically schedule meetings using Google Calendar API with natural language parsing
- **📧 Email Automation**: Send emails and follow-ups using SendGrid API with intelligent content generation
- **🧠 Natural Language Processing**: Process commands using Google Gemini API for human-like understanding
- **⏰ Proactive Reminders**: Automated daily task reminders, meeting notifications, and weekly/monthly summaries
- **🌐 REST API**: HTTP endpoints for triggering actions programmatically
- **⚙️ Configurable**: All settings via environment variables for easy deployment

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   HTTP Server   │    │   Agent Service │    │   API Clients   │
│   (Port 8080)   │◄──►│   (Main Logic)  │◄──►│  (Calendar,     │
│                 │    │                 │    │   Email, NLP)   │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         │                       │                       │
         ▼                       ▼                       ▼
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Endpoints:    │    │   Components:   │    │   External      │
│ • /health       │    │ • Handler       │    │   APIs:         │
│ • /schedule     │    │ • Scheduler     │    │ • Google        │
│ • /email        │    │ • Logger        │    │   Calendar      │
│ • /nlp          │    │ • Config        │    │ • SendGrid      │
│ • /status       │    │ • Parser        │    │ • Gemini        │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

##  Quick Start

### Prerequisites

- Go 1.21 or higher
- API keys for:
  - Google Calendar API
  - SendGrid API
  - Google Gemini API

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/azme12/ai-agent-project.git
   cd ai-agent-project
   ```

2. **Set up environment variables**
   ```bash
   # Copy the example environment file
   cp env.example .env
   
   # Edit the .env file with your API keys
   nano .env
   ```

3. **Install dependencies**
   ```bash
   go mod download
   ```

4. **Run the application**
   ```bash
   go run cmd/main.go
   ```

5. **Test the health endpoint**
   ```bash
   curl http://localhost:8080/health
   ```

## 📋 API Documentation

### Health Check
```bash
GET /health
```
Returns service status and version information.

**Response:**
```json
{
  "status": "healthy",
  "service": "ai-agent",
  "version": "1.0.0"
}
```

### Service Status
```bash
GET /status
```
Returns service status and available endpoints.

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

### Schedule Meeting
```bash
POST /schedule
Content-Type: application/json

{
  "task": "Schedule a meeting with john@example.com tomorrow at 2 PM about project review"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Task processed successfully",
  "task": "Schedule a meeting with john@example.com tomorrow at 2 PM about project review"
}
```

### Send Email
```bash
POST /email
Content-Type: application/json

{
  "task": "Send email to client@example.com saying thank you for the meeting"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Email task processed successfully",
  "task": "Send email to client@example.com saying thank you for the meeting"
}
```

### Process NLP Command
```bash
POST /nlp
Content-Type: application/json

{
  "command": "What meetings do I have today?"
}
```

**Response:**
```json
{
  "status": "success",
  "response": "I'll check your calendar for today's meetings...",
  "command": "What meetings do I have today?"
}
```

## 🔧 Configuration

The agent uses environment variables for all configuration. Copy `env.example` to `.env` and customize:

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

##  Use Cases

### Executive Assistant Tasks Automated

1. **Meeting Management**
   - Schedule meetings based on natural language commands
   - Send meeting reminders 15 minutes before start
   - Handle meeting conflicts and rescheduling
   - Extract attendees, times, and topics from natural language

2. **Email Automation**
   - Send follow-up emails after meetings
   - Send daily task reminders
   - Process email requests via natural language
   - Generate email content based on context

3. **Task Management**
   - Daily task summaries at 9 AM
   - Weekly summaries on Mondays
   - Monthly summaries on the 1st
   - Proactive reminders for deadlines

4. **Calendar Integration**
   - Sync with Google Calendar
   - Check availability
   - Manage recurring meetings
   - Query calendar information

## 🔄 Proactive Actions

The agent runs proactive tasks in the background:

- **Daily Reminders**: Sends task summaries at configured time (default: 9 AM)
- **Meeting Alerts**: Notifies before meetings (default: 15 minutes)
- **Weekly Summaries**: Sends weekly summaries on Mondays
- **Monthly Summaries**: Sends monthly summaries on the 1st
- **Periodic Checks**: Monitors calendar and tasks every minute

## 🧪 Testing

### Run the Test Suite
```bash
# Make sure the service is running first
go run cmd/main.go

# In another terminal, run tests
./test.sh
```

### Manual Testing Examples
```bash
# Test scheduling
curl -X POST http://localhost:8080/schedule \
  -H "Content-Type: application/json" \
  -d '{"task": "Schedule meeting with team tomorrow at 10 AM"}'

# Test email
curl -X POST http://localhost:8080/email \
  -H "Content-Type: application/json" \
  -d '{"task": "Send email to john@example.com about project status"}'

# Test NLP
curl -X POST http://localhost:8080/nlp \
  -H "Content-Type: application/json" \
  -d '{"command": "What is my schedule for today?"}'
```

## 🛠️ Development

### Project Structure
```
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── agent/
│   │   ├── handler.go       # Task processing logic
│   │   ├── scheduler.go     # Proactive scheduling
│   │   └── service.go       # Main agent service
│   ├── api/
│   │   ├── calendar.go      # Google Calendar integration
│   │   ├── email.go         # SendGrid integration
│   │   └── gemini.go        # Gemini NLP integration
│   └── config/
│       └── config.go        # Configuration management
├── pkg/
│   └── logger/
│       └── logger.go        # Logging utilities
├── env.example              # Environment variables template
├── docker-compose.yml       # Docker Compose configuration
├── Dockerfile               # Docker build configuration
├── test.sh                  # Test script
└── go.mod                   # Go module file
```

### Adding New Features

1. **New API Integration**: Add service in `internal/api/`
2. **New Task Type**: Update handler in `internal/agent/handler.go`
3. **New Endpoint**: Add handler in `cmd/main.go`
4. **New Configuration**: Add to `internal/config/config.go`

### Development Workflow
```bash
# Start development
go run cmd/main.go

# Run tests
./test.sh

# Build for production
go build -o ai-agent cmd/main.go

# Run production build
./ai-agent
```

##  Deployment

### Local Development
```bash
go run cmd/main.go
```

### Production Build
```bash
go build -o ai-agent cmd/main.go
./ai-agent
```

### Docker Deployment

#### Build and Run with Docker
```bash
# Build the Docker image
docker build -t ai-agent .

# Run the container
docker run -d --name ai-agent -p 8080:8080 ai-agent

# Test the container
curl http://localhost:8080/health
```

#### Using Docker Compose
```bash
# Start with docker-compose
docker-compose up -d

# View logs
docker-compose logs -f

# Stop the service
docker-compose down
```

#### Environment Variables with Docker
```bash
# Run with environment variables
docker run -d --name ai-agent \
  -p 8080:8080 \
  -e GOOGLE_CALENDAR_API_KEY="your_key" \
  -e SENDGRID_API_KEY="your_key" \
  -e GEMINI_API_KEY="your_key" \
  ai-agent
```

### Cloud Deployment

#### Google Cloud Run (Recommended)
```bash
# Build and deploy to Cloud Run
gcloud run deploy ai-agent \
  --source . \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --set-env-vars GOOGLE_CALENDAR_API_KEY=your_key,SENDGRID_API_KEY=your_key,GEMINI_API_KEY=your_key
```

#### AWS ECS
```bash
# Build and push to ECR
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin your-account.dkr.ecr.us-east-1.amazonaws.com
docker build -t ai-agent .
docker tag ai-agent:latest your-account.dkr.ecr.us-east-1.amazonaws.com/ai-agent:latest
docker push your-account.dkr.ecr.us-east-1.amazonaws.com/ai-agent:latest
```

## 🔒 Security

- API keys are loaded from environment variables
- No hardcoded credentials in source code
- HTTP endpoints validate input data
- Error handling prevents information leakage
- Docker runs as non-root user
- Health checks for monitoring

## 📈 Monitoring

- Health check endpoint for monitoring
- Structured logging with different levels
- Error tracking and reporting
- Performance metrics via HTTP endpoints
- Docker health checks
- Graceful shutdown handling




---

**Built with ❤️ using Go, Google APIs, and modern automation practices.**
