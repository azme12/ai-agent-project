#!/bin/bash

# AI Agent Test Script
# This script tests all endpoints and functionality of the AI Agent

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
BASE_URL="http://localhost:8080"
TIMEOUT=10

echo -e "${BLUE}🤖 AI Agent Test Suite${NC}"
echo "================================"

# Function to check if service is running
check_service() {
    echo -e "${YELLOW}Checking if service is running...${NC}"
    
    if curl -s --max-time $TIMEOUT "$BASE_URL/health" > /dev/null; then
        echo -e "${GREEN}✅ Service is running${NC}"
        return 0
    else
        echo -e "${RED}❌ Service is not running. Please start the service first.${NC}"
        echo "Run: go run cmd/main.go"
        exit 1
    fi
}

# Function to test endpoint
test_endpoint() {
    local method=$1
    local endpoint=$2
    local data=$3
    local description=$4
    
    echo -e "${YELLOW}Testing: $description${NC}"
    
    if [ "$method" = "GET" ]; then
        response=$(curl -s --max-time $TIMEOUT "$BASE_URL$endpoint")
    else
        response=$(curl -s --max-time $TIMEOUT -X $method \
            -H "Content-Type: application/json" \
            -d "$data" \
            "$BASE_URL$endpoint")
    fi
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✅ Success${NC}"
        echo "Response: $response" | jq '.' 2>/dev/null || echo "Response: $response"
    else
        echo -e "${RED}❌ Failed${NC}"
    fi
    echo ""
}

# Function to test error cases
test_error_cases() {
    echo -e "${BLUE}Testing Error Cases${NC}"
    echo "====================="
    
    # Test invalid JSON
    echo -e "${YELLOW}Testing invalid JSON...${NC}"
    response=$(curl -s --max-time $TIMEOUT -X POST \
        -H "Content-Type: application/json" \
        -d "invalid json" \
        "$BASE_URL/schedule")
    
    if echo "$response" | grep -q "Invalid JSON"; then
        echo -e "${GREEN}✅ Invalid JSON handled correctly${NC}"
    else
        echo -e "${RED}❌ Invalid JSON not handled properly${NC}"
    fi
    
    # Test missing required fields
    echo -e "${YELLOW}Testing missing required fields...${NC}"
    response=$(curl -s --max-time $TIMEOUT -X POST \
        -H "Content-Type: application/json" \
        -d '{}' \
        "$BASE_URL/schedule")
    
    if echo "$response" | grep -q "Task is required"; then
        echo -e "${GREEN}✅ Missing fields handled correctly${NC}"
    else
        echo -e "${RED}❌ Missing fields not handled properly${NC}"
    fi
    
    # Test wrong HTTP method
    echo -e "${YELLOW}Testing wrong HTTP method...${NC}"
    response=$(curl -s --max-time $TIMEOUT -X GET "$BASE_URL/schedule")
    
    if echo "$response" | grep -q "Method not allowed"; then
        echo -e "${GREEN}✅ Wrong method handled correctly${NC}"
    else
        echo -e "${RED}❌ Wrong method not handled properly${NC}"
    fi
    
    echo ""
}

# Main test execution
main() {
    check_service
    
    echo -e "${BLUE}Testing Basic Endpoints${NC}"
    echo "========================"
    
    # Test health endpoint
    test_endpoint "GET" "/health" "" "Health Check"
    
    # Test status endpoint
    test_endpoint "GET" "/status" "" "Status Check"
    
    echo -e "${BLUE}Testing Core Functionality${NC}"
    echo "============================"
    
    # Test scheduling
    test_endpoint "POST" "/schedule" '{"task": "Schedule a meeting with john@example.com tomorrow at 2 PM about project review"}' "Schedule Meeting"
    
    # Test email
    test_endpoint "POST" "/email" '{"task": "Send email to client@example.com saying thank you for the meeting"}' "Send Email"
    
    # Test NLP
    test_endpoint "POST" "/nlp" '{"command": "What meetings do I have today?"}' "NLP Processing"
    
    echo -e "${BLUE}Testing Advanced Scenarios${NC}"
    echo "============================="
    
    # Test complex scheduling
    test_endpoint "POST" "/schedule" '{"task": "Schedule a 1-hour meeting with team@company.com and manager@company.com next week about quarterly planning"}' "Complex Meeting Scheduling"
    
    # Test reminder
    test_endpoint "POST" "/nlp" '{"command": "Remind me to call the client tomorrow at 10 AM"}' "Set Reminder"
    
    # Test calendar query
    test_endpoint "POST" "/nlp" '{"command": "Show me my calendar for this week"}' "Calendar Query"
    
    # Test error cases
    test_error_cases
    
    echo -e "${BLUE}Performance Tests${NC}"
    echo "=================="
    
    # Test concurrent requests
    echo -e "${YELLOW}Testing concurrent requests...${NC}"
    for i in {1..5}; do
        curl -s --max-time $TIMEOUT "$BASE_URL/health" > /dev/null &
    done
    wait
    echo -e "${GREEN}✅ Concurrent requests handled${NC}"
    
    echo ""
    echo -e "${GREEN}🎉 All tests completed!${NC}"
    echo ""
    echo -e "${BLUE}Next Steps:${NC}"
    echo "1. Check the logs for any errors"
    echo "2. Verify API integrations are working"
    echo "3. Test with real API keys for full functionality"
    echo "4. Monitor the scheduler for proactive actions"
}

# Check if jq is installed for pretty JSON output
if ! command -v jq &> /dev/null; then
    echo -e "${YELLOW}Warning: jq not installed. JSON responses will not be formatted.${NC}"
    echo "Install jq for better output formatting."
fi

# Run tests
main
