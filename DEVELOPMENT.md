# Development Guide

This document provides instructions for developers contributing to the awake project.

## Setup

1. Clone the repository
   ```
   git clone https://github.com/yourusername/awake.git
   cd awake
   ```

2. Install Task (if not already installed)
   ```
   go install github.com/go-task/task/v3/cmd/task@latest
   ```

3. Install golangci-lint (if not already installed)
   ```
   # macOS
   brew install golangci-lint

   # Other platforms
   # See https://golangci-lint.run/usage/install/
   ```

4. Set up git hooks
   ```
   task hooks
   ```

## Development Workflow

### Building

```
task build
```

This will create the `awake` binary in the current directory.

### Running Tests

```
task test
```

### Running the Linter

```
task lint
```

### Running Pre-commit Checks Manually

You can manually run the same checks that the pre-commit hook runs:

```
task precommit
```

This will run both the linter and tests.

## Git Hooks

### Pre-commit Hook

The pre-commit hook runs automatically before each commit and:

1. Checks if Task is installed
2. Checks if golangci-lint is installed
3. Uses `task precommit` to run the linter and tests

If any of these checks fail, the commit will be aborted.

## Project Structure

- `main.go` - Main application code
- `main_test.go` - Tests for the application
- `.githooks/` - Git hooks
- `scripts/` - Development and automation scripts
