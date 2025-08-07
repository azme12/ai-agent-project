#!/bin/bash

echo "🐳 Testing AI Agent Docker Container"
echo "===================================="

# Build the Docker image
echo "🔨 Building Docker image..."
docker build -t ai-agent .

if [ $? -eq 0 ]; then
    echo "✅ Docker image built successfully"
else
    echo "❌ Docker build failed"
    exit 1
fi

# Run the container
echo ""
echo "🚀 Starting AI Agent container..."
docker run -d --name ai-agent-test -p 8080:8080 ai-agent

# Wait for container to start
echo "⏳ Waiting for container to start..."
sleep 10

# Test health endpoint
echo ""
echo "📋 Testing health endpoint..."
curl -s http://localhost:8080/health

if [ $? -eq 0 ]; then
    echo ""
    echo "✅ Health check passed"
else
    echo ""
    echo "❌ Health check failed"
fi

# Test other endpoints
echo ""
echo "📅 Testing schedule endpoint..."
curl -X POST http://localhost:8080/schedule \
  -H "Content-Type: application/json" \
  -d '{"task": "Schedule meeting with team tomorrow at 10 AM"}' \
  -s

echo ""
echo "📧 Testing email endpoint..."
curl -X POST http://localhost:8080/email \
  -H "Content-Type: application/json" \
  -d '{"task": "Send email to john@example.com about project status"}' \
  -s

echo ""
echo "🧠 Testing NLP endpoint..."
curl -X POST http://localhost:8080/nlp \
  -H "Content-Type: application/json" \
  -d '{"command": "What is my schedule for today?"}' \
  -s

echo ""
echo "🛑 Stopping container..."
docker stop ai-agent-test
docker rm ai-agent-test

echo ""
echo "🎉 Docker test completed successfully!"
