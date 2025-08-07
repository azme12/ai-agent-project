# AI Agent Project - Executive Assistant Automation

An intelligent AI agent that automates executive assistant tasks including scheduling meetings, sending emails, and processing natural language commands.

## 🎯 Project Overview

This AI agent serves as an **Executive Assistant in a Tech Startup**, automating high-impact, repetitive tasks to increase productivity and efficiency.

### Core Features

- **📅 Meeting Scheduling**: Automatically schedule meetings using Google Calendar API
- **📧 Email Automation**: Send emails and follow-ups using SendGrid API
- **🧠 Natural Language Processing**: Process commands using Google Gemini API
- **⏰ Proactive Reminders**: Automated daily task reminders and meeting notifications
- **🌐 REST API**: HTTP endpoints for triggering actions programmatically

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
└─────────────────┘    └─────────────────┘    │ • Gemini        │
                                              └─────────────────┘
```

## 🚀 Quick Start

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
   export GOOGLE_CALENDAR_API_KEY="your_calendar_api_key"
   export SENDGRID_API_KEY="your_sendgrid_api_key"
   export GEMINI_API_KEY="your_gemini_api_key"
   export SERVER_PORT="8080"
   export LOG_LEVEL="info"
   ```

3. **Run the application**
   ```bash
   go run cmd/main.go
   ```

4. **Test the health endpoint**
   ```bash
   curl http://localhost:8080/health
   ```

## 📋 API Endpoints

### Health Check
```bash
GET /health
```
Returns service status.

### Schedule Meeting
```bash
POST /schedule
Content-Type: application/json

{
  "task": "Schedule a meeting with John tomorrow at 2 PM"
}
```

### Send Email
```bash
POST /email
Content-Type: application/json

{
  "task": "Send email to client@example.com about project update"
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

## 🔧 Configuration

The agent uses environment variables for configuration:

| Variable | Description | Default |
|----------|-------------|---------|
| `GOOGLE_CALENDAR_API_KEY` | Google Calendar API key | "" |
| `SENDGRID_API_KEY` | SendGrid API key | "" |
| `GEMINI_API_KEY` | Google Gemini API key | "" |
| `SERVER_PORT` | HTTP server port | "8080" |
| `LOG_LEVEL` | Logging level | "info" |

## 🎯 Use Cases

### Executive Assistant Tasks Automated

1. **Meeting Management**
   - Schedule meetings based on natural language commands
   - Send meeting reminders 15 minutes before start
   - Handle meeting conflicts and rescheduling

2. **Email Automation**
   - Send follow-up emails after meetings
   - Send daily task reminders
   - Process email requests via natural language

3. **Task Management**
   - Daily task summaries
   - Proactive reminders for deadlines
   - Natural language task creation

4. **Calendar Integration**
   - Sync with Google Calendar
   - Check availability
   - Manage recurring meetings

## 🔄 Proactive Actions

The agent runs proactive tasks in the background:

- **Daily Reminders**: Sends task summaries at 9 AM
- **Meeting Alerts**: Notifies 15 minutes before meetings
- **Periodic Checks**: Monitors calendar and tasks every minute

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
└── go.mod                   # Go module file
```

### Adding New Features

1. **New API Integration**: Add service in `internal/api/`
2. **New Task Type**: Update handler in `internal/agent/handler.go`
3. **New Endpoint**: Add handler in `cmd/main.go`

## 🧪 Testing

### Manual Testing
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

## 🔒 Security

- API keys are loaded from environment variables
- No hardcoded credentials in source code
- HTTP endpoints validate input data
- Error handling prevents information leakage

## 📈 Monitoring

- Health check endpoint for monitoring
- Structured logging with different levels
- Error tracking and reporting
- Performance metrics via HTTP endpoints

## 🚀 Deployment

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

#### Docker Testing
```bash
# Run automated Docker tests
./docker-test.sh
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

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.

## 🆘 Support

For issues and questions:
- Create an issue in the GitHub repository
- Check the logs for debugging information
- Verify API keys are correctly configured

---

**Built with ❤️ using Go, Google APIs, and modern automation practices.**
