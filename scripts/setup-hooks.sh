#!/bin/sh
#
# Script to set up git hooks for the project

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

# Configure git to use our hooks directory
printf "${YELLOW}Configuring git to use hooks from .githooks directory...${NC}\n"
git config core.hooksPath .githooks

# Make all hooks executable
printf "${YELLOW}Making hooks executable...${NC}\n"
chmod +x .githooks/*

printf "${GREEN}Git hooks setup complete!${NC}\n"
printf "${YELLOW}Note: Pre-commit hook will check for golangci-lint and run tests before each commit.${NC}\n"
