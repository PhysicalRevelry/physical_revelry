# Physical Revelry

A full-stack application with Flutter mobile frontend and Go backend.

## Project Structure

This is a monorepo containing:

- **Mobile**: Flutter app with Riverpod state management
- **Backend**: Go server with GraphQL API
- **Shared**: Common type definitions

## Quick Start

### Prerequisites

- Flutter SDK
- Go 1.21+
- Docker & Docker Compose (optional)

### Development

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd physical_revelry
   ```

2. **Start the mobile app**
   ```bash
   cd mobile/
   flutter pub get
   flutter run
   ```

3. **Start the backend** (in a new terminal)
   ```bash
   cd backend/
   go run cmd/server/main.go
   ```

4. **Or use Docker Compose** (from root directory)
   ```bash
   docker-compose up
   ```

## Development Commands

See `CLAUDE.md` for detailed development commands and architecture information.
