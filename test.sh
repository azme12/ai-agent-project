#!/bin/bash

echo "🧪 Testing AI Agent API Endpoints"
echo "=================================="

# Start the server in background
echo "🚀 Starting AI Agent server..."
go run cmd/main.go &
SERVER_PID=$!

# Wait for server to start
sleep 3

echo ""
echo "📋 Testing Health Endpoint"
echo "-------------------------"
curl -s http://localhost:8080/health
echo ""

echo ""
echo "📅 Testing Schedule Endpoint"
echo "---------------------------"
curl -X POST http://localhost:8080/schedule \
  -H "Content-Type: application/json" \
  -d '{"task": "Schedule meeting with team tomorrow at 10 AM"}' \
  -s
echo ""

echo ""
echo "📧 Testing Email Endpoint"
echo "------------------------"
curl -X POST http://localhost:8080/email \
  -H "Content-Type: application/json" \
  -d '{"task": "Send email to john@example.com about project status"}' \
  -s
echo ""

echo ""
echo "🧠 Testing NLP Endpoint"
echo "----------------------"
curl -X POST http://localhost:8080/nlp \
  -H "Content-Type: application/json" \
  -d '{"command": "What is my schedule for today?"}' \
  -s
echo ""

echo ""
echo "✅ All tests completed!"
echo ""

# Stop the server
echo "🛑 Stopping server..."
kill $SERVER_PID
wait $SERVER_PID 2>/dev/null

echo "🎉 Test script completed successfully!"
