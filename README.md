# URL Redirect Service

## Project Structure
- cmd/url-redirect-service/main.go - Main application code
- Dockerfile - Multi-stage build configuration
- docker-compose.yml - Development environment setup
- go.mod - Go module definition

## Getting Started
1. Build and run the service:
   ```
   docker-compose up --build
   ```

2. Test the endpoints:
   - Health check: http://localhost:80/health
   - Demo redirect: http://localhost:80/redirect/demo
   - Service info: http://localhost:80/

## Development
The service includes:
- Health check endpoint
- Request logging
- Basic redirect functionality
- Docker development environment

## Next Steps
1. Add persistent storage for redirect mappings
2. Implement authentication/authorization
3. Add metrics and monitoring
4. Expand redirect configuration options
