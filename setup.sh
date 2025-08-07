#!/bin/bash

echo "🚀 AI Agent Project Setup"
echo "========================"

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21 or higher."
    exit 1
fi

echo "✅ Go is installed: $(go version)"

# Check if the project compiles
echo ""
echo "🔨 Building project..."
if go build ./cmd/main.go; then
    echo "✅ Project builds successfully"
else
    echo "❌ Build failed. Please check the code."
    exit 1
fi

echo ""
echo "📋 Environment Variables Setup"
echo "============================="
echo ""
echo "Please set the following environment variables:"
echo ""
echo "export GOOGLE_CALENDAR_API_KEY=\"your_calendar_api_key\""
echo "export SENDGRID_API_KEY=\"your_sendgrid_api_key\""
echo "export GEMINI_API_KEY=\"your_gemini_api_key\""
echo "export SERVER_PORT=\"8080\""
echo "export LOG_LEVEL=\"info\""
echo ""
echo "💡 Note: If you don't have API keys yet, the agent will run in mock mode."
echo ""

# Create .env.example file
cat > .env.example << EOF
# AI Agent Environment Variables
# Copy this file to .env and fill in your API keys

# Google Calendar API Key (Optional - will use mock mode if not set)
GOOGLE_CALENDAR_API_KEY=

# SendGrid API Key (Optional - will use mock mode if not set)
SENDGRID_API_KEY=

# Google Gemini API Key (Optional - will use mock mode if not set)
GEMINI_API_KEY=

# Server Configuration
SERVER_PORT=8080
LOG_LEVEL=info
EOF

echo "📄 Created .env.example file"
echo ""
echo "🎯 Next Steps:"
echo "1. Copy .env.example to .env"
echo "2. Fill in your API keys (optional for testing)"
echo "3. Run: go run cmd/main.go"
echo "4. Test with: ./test.sh"
echo ""
echo "🎉 Setup completed successfully!"
