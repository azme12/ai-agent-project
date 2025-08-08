# Postman Setup Guide for AI Agent API

This guide will help you set up and use the Postman collection for testing the AI Agent API.

## Quick Setup

### 1. Import the Collection

1. Open Postman
2. Click **Import** button
3. Select **Upload Files**
4. Choose `AI_Agent_API_Postman_Collection.json`
5. Click **Import**

### 2. Import the Environment

1. Click **Import** again
2. Select **Upload Files**
3. Choose `AI_Agent_API_Postman_Environment.json`
4. Click **Import**

### 3. Configure Environment Variables

1. Click the **Environments** dropdown (top right)
2. Select **AI Agent API Environment**
3. Click the **eye icon** to view variables
4. Update the following variables:

| Variable | Description | Example Value |
|----------|-------------|---------------|
| `base_url` | API base URL | `http://localhost:8080` |
| `google_calendar_api_key` | Google Calendar API key | `your_actual_key_here` |
| `sendgrid_api_key` | SendGrid API key | `your_actual_key_here` |
| `gemini_api_key` | Gemini API key | `your_actual_key_here` |
| `from_email` | Sender email | `ai-assistant@yourdomain.com` |
| `user_email` | User email | `user@yourdomain.com` |

### 4. Start the AI Agent Service

```bash
# Navigate to your project directory
cd /path/to/ai-agent-project

# Start the service
go run cmd/main.go
```

### 5. Test the API

1. **Health Check**: Run the "Health Check" request to verify the service is running
2. **Service Status**: Run the "Service Status" request to see all available endpoints
3. **Schedule Meeting**: Test meeting scheduling with natural language
4. **Send Email**: Test email automation
5. **NLP Command**: Test natural language processing

## Collection Structure

The collection is organized into four main folders:

### Health & Status
- **Health Check**: Verify service is running
- **Service Status**: Get detailed service information

### Meeting Scheduling
- **Schedule Meeting**: Create meetings using natural language

### Email Automation
- **Send Email**: Send emails using natural language

### Natural Language Processing
- **Process NLP Command**: Query the AI assistant

## Example Requests

### Schedule a Meeting
```json
{
  "task": "Schedule a meeting with john@example.com tomorrow at 2 PM about project review"
}
```

### Send an Email
```json
{
  "task": "Send email to client@example.com saying thank you for the meeting and we'll follow up next week"
}
```

### Process NLP Command
```json
{
  "command": "What meetings do I have today?"
}
```

## Troubleshooting

### Common Issues

1. **Connection Refused**
   - Ensure the AI Agent service is running
   - Check the `base_url` variable is correct
   - Verify the port (default: 8080) is not blocked

2. **400 Bad Request**
   - Check request body format is valid JSON
   - Ensure required fields are not empty
   - Verify Content-Type header is set to `application/json`

3. **500 Internal Server Error**
   - Check service logs for detailed error information
   - Verify API keys are correctly configured
   - Ensure external services (Google Calendar, SendGrid, Gemini) are accessible

4. **API Key Issues**
   - Verify API keys are valid and have proper permissions
   - Check environment variables are set correctly
   - Ensure keys are not expired or rate-limited

### Debug Steps

1. **Check Service Health**
   ```bash
   curl -X GET http://localhost:8080/health
   ```

2. **View Service Logs**
   ```bash
   # If running with go run
   go run cmd/main.go
   
   # If running with Docker
   docker logs ai-agent
   ```

3. **Test Individual Endpoints**
   ```bash
   # Health check
   curl -X GET http://localhost:8080/health
   
   # Status check
   curl -X GET http://localhost:8080/status
   ```

## Environment Variables Setup

Create a `.env` file in your project root:

```bash
# Copy the example file
cp env.example .env

# Edit with your actual values
nano .env
```

Required variables:
```bash
GOOGLE_CALENDAR_API_KEY=your_google_calendar_api_key_here
SENDGRID_API_KEY=your_sendgrid_api_key_here
GEMINI_API_KEY=your_gemini_api_key_here
FROM_EMAIL=ai-assistant@yourdomain.com
USER_EMAIL=user@yourdomain.com
```

## Getting API Keys

### Google Calendar API
1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Create a new project or select existing
3. Enable Google Calendar API
4. Create credentials (API Key)
5. Set up OAuth 2.0 for calendar access

### SendGrid API
1. Go to [SendGrid Dashboard](https://app.sendgrid.com/)
2. Navigate to Settings > API Keys
3. Create a new API Key
4. Set appropriate permissions (Mail Send)

### Google Gemini API
1. Go to [Google AI Studio](https://ai.google.dev/)
2. Create a new API key
3. Enable Gemini API access

## Advanced Usage

### Using Pre-request Scripts

You can add pre-request scripts to automatically set headers or variables:

```javascript
// Set Content-Type header
pm.request.headers.add({
    key: 'Content-Type',
    value: 'application/json'
});
```

### Using Tests

Add tests to verify responses:

```javascript
// Test health endpoint
pm.test("Status code is 200", function () {
    pm.response.to.have.status(200);
});

pm.test("Response has required fields", function () {
    const jsonData = pm.response.json();
    pm.expect(jsonData).to.have.property('status');
    pm.expect(jsonData).to.have.property('service');
});
```

### Environment Switching

Create multiple environments for different stages:

- **Local Development**: `http://localhost:8080`
- **Staging**: `https://staging-api.yourdomain.com`
- **Production**: `https://api.yourdomain.com`

## Support

If you encounter issues:

1. Check the [API Documentation](API_Documentation.md)
2. Review the [README](README.md) for setup instructions
3. Check service logs for error details
4. Verify all environment variables are set correctly

---

**Happy Testing! 🚀**
