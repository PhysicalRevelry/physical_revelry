# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a monorepo for the "physical_revelry" application containing:
- **Mobile**: Flutter app using Riverpod for state management and Hooks for widget lifecycle
- **Backend**: Go server with GraphQL API
- **Shared**: Common types and utilities

## Monorepo Structure

```
physical_revelry/
├── .github/workflows/    # CI/CD workflows
├── backend/             # Go backend with GraphQL
│   ├── cmd/server/      # Main server entry point
│   ├── internal/        # Private application code
│   ├── graph/           # GraphQL schema and resolvers
│   └── pkg/             # Public packages
├── mobile/              # Flutter mobile app
│   ├── lib/             # Dart source code
│   ├── test/            # Widget and unit tests
│   ├── android/         # Android-specific files
│   └── ios/             # iOS-specific files
├── shared/types/        # Shared type definitions
├── docker-compose.yml   # Local development orchestration
└── .env.example         # Environment variables template
```

## Development Commands

### Mobile App (Flutter)
All Flutter commands must be run from the `mobile/` directory:

```bash
cd mobile/
flutter pub get              # Install dependencies
flutter run                  # Run the app
flutter test                 # Run tests
flutter analyze              # Static analysis
flutter build apk           # Build Android APK
flutter build ios           # Build for iOS
```

### Backend (Go)
All Go commands must be run from the `backend/` directory:

```bash
cd backend/
go mod tidy                  # Clean up dependencies
go run cmd/server/main.go    # Run development server
go build -o server cmd/server/main.go  # Build binary
go test ./...                # Run tests
```

### Full Stack Development
From the root directory:

```bash
docker-compose up            # Start all services
docker-compose down          # Stop all services
```

## Mobile App Architecture

- **State Management**: Uses Riverpod with StateNotifier pattern
- **UI**: HookConsumerWidget instead of StatefulWidget for cleaner code
- **Dependencies**: flutter_riverpod, flutter_hooks, hooks_riverpod
- **Testing**: Widgets wrapped with ProviderScope for Riverpod integration

## Backend Architecture

- **Server**: Go with net/http
- **GraphQL**: Schema-first approach with gqlgen (when implemented)
- **Structure**: Clean architecture with internal packages for business logic

## Key Files

- `mobile/pubspec.yaml` - Flutter dependencies and configuration
- `backend/go.mod` - Go module dependencies
- `backend/graph/schema.graphqls` - GraphQL schema definition
- `docker-compose.yml` - Local development environment setup