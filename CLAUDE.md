# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Flutter mobile application project named "physical_revelry" that targets both Android and iOS platforms. It's currently a basic Flutter starter app with a counter demo but appears to be set up for future development of a more substantial application.

## Development Commands

### Dependencies and Setup
- `flutter pub get` - Install dependencies
- `flutter pub upgrade` - Upgrade dependencies to latest versions
- `flutter pub outdated` - Check for outdated dependencies

### Running the App
- `flutter run` - Run the app (will prompt to select target device)
- `flutter run -d chrome` - Run in Chrome browser (for web development)
- `flutter run -d android` - Run on Android device/emulator
- `flutter run -d ios` - Run on iOS simulator (macOS only)

### Testing and Quality
- `flutter test` - Run all unit and widget tests
- `flutter test test/widget_test.dart` - Run specific test file
- `flutter analyze` - Run static analysis and linting
- `flutter doctor` - Check Flutter installation and dependencies

### Building
- `flutter build apk` - Build Android APK
- `flutter build ios` - Build for iOS (requires Xcode on macOS)
- `flutter build web` - Build for web deployment

## Project Structure

```
lib/
  main.dart           # Main entry point, contains MyApp and MyHomePage widgets
test/
  widget_test.dart    # Widget tests for the counter functionality
android/              # Android-specific configuration and build files
ios/                  # iOS-specific configuration and build files
```

## Key Configuration Files

- `pubspec.yaml` - Flutter project configuration, dependencies, assets
- `analysis_options.yaml` - Dart analyzer configuration with flutter_lints rules
- `android/app/build.gradle.kts` - Android build configuration
- `ios/Runner.xcodeproj/` - iOS project configuration

## Development Environment

- Dart SDK: ^3.8.1
- Uses flutter_lints for code quality enforcement
- Material Design with Cupertino Icons for cross-platform UI consistency
- Hot reload is available during development (`r` in terminal or IDE hot reload)