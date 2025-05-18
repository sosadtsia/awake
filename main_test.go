package main

import (
	"flag"
	"os"
	"testing"
	"time"
)

func TestParseFlags(t *testing.T) {
	// Save original args and flags
	oldArgs := os.Args
	oldFlagCommandLine := flag.CommandLine

	// Restore them after the test
	defer func() {
		os.Args = oldArgs
		flag.CommandLine = oldFlagCommandLine
	}()

	testCases := []struct {
		name          string
		args          []string
		expectedQuiet bool
		expectedDebug bool
		expectedHelp  bool
		expectedVer   bool
		expectedDur   time.Duration
		expectedBg    bool
	}{
		{
			name:          "default",
			args:          []string{"awake"},
			expectedQuiet: false,
			expectedDebug: false,
			expectedHelp:  false,
			expectedVer:   false,
			expectedDur:   0,
			expectedBg:    false,
		},
		{
			name:          "quiet mode",
			args:          []string{"awake", "-q"},
			expectedQuiet: true,
			expectedDebug: false,
			expectedHelp:  false,
			expectedVer:   false,
			expectedDur:   0,
			expectedBg:    false,
		},
		{
			name:          "debug mode",
			args:          []string{"awake", "-d"},
			expectedQuiet: false,
			expectedDebug: true,
			expectedHelp:  false,
			expectedVer:   false,
			expectedDur:   0,
			expectedBg:    false,
		},
		{
			name:          "with duration",
			args:          []string{"awake", "-t", "2h"},
			expectedQuiet: false,
			expectedDebug: false,
			expectedHelp:  false,
			expectedVer:   false,
			expectedDur:   2 * time.Hour,
			expectedBg:    false,
		},
		{
			name:          "help",
			args:          []string{"awake", "-h"},
			expectedQuiet: false,
			expectedDebug: false,
			expectedHelp:  true,
			expectedVer:   false,
			expectedDur:   0,
			expectedBg:    false,
		},
		{
			name:          "version",
			args:          []string{"awake", "-v"},
			expectedQuiet: false,
			expectedDebug: false,
			expectedHelp:  false,
			expectedVer:   true,
			expectedDur:   0,
			expectedBg:    false,
		},
		{
			name:          "background mode",
			args:          []string{"awake", "-b"},
			expectedQuiet: false,
			expectedDebug: false,
			expectedHelp:  false,
			expectedVer:   false,
			expectedDur:   0,
			expectedBg:    true,
		},
		{
			name:          "background shorthand",
			args:          []string{"awake", "-background"},
			expectedQuiet: false,
			expectedDebug: false,
			expectedHelp:  false,
			expectedVer:   false,
			expectedDur:   0,
			expectedBg:    true,
		},
		{
			name:          "combined",
			args:          []string{"awake", "-q", "-d", "-t", "30m"},
			expectedQuiet: true,
			expectedDebug: true,
			expectedHelp:  false,
			expectedVer:   false,
			expectedDur:   30 * time.Minute,
			expectedBg:    false,
		},
		{
			name:          "background with duration",
			args:          []string{"awake", "-b", "-t", "1h"},
			expectedQuiet: false,
			expectedDebug: false,
			expectedHelp:  false,
			expectedVer:   false,
			expectedDur:   1 * time.Hour,
			expectedBg:    true,
		},
		{
			name:          "all options",
			args:          []string{"awake", "-q", "-d", "-t", "30m", "-b"},
			expectedQuiet: true,
			expectedDebug: true,
			expectedHelp:  false,
			expectedVer:   false,
			expectedDur:   30 * time.Minute,
			expectedBg:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Reset flags for each test
			flag.CommandLine = flag.NewFlagSet(tc.args[0], flag.ExitOnError)

			// Set args for this test
			os.Args = tc.args

			// Parse flags
			opts := parseFlags()

			// Check results
			if opts.quiet != tc.expectedQuiet {
				t.Errorf("quiet: expected %v, got %v", tc.expectedQuiet, opts.quiet)
			}
			if opts.debug != tc.expectedDebug {
				t.Errorf("debug: expected %v, got %v", tc.expectedDebug, opts.debug)
			}
			if opts.help != tc.expectedHelp {
				t.Errorf("help: expected %v, got %v", tc.expectedHelp, opts.help)
			}
			if opts.version != tc.expectedVer {
				t.Errorf("version: expected %v, got %v", tc.expectedVer, opts.version)
			}
			if opts.duration != tc.expectedDur {
				t.Errorf("duration: expected %v, got %v", tc.expectedDur, opts.duration)
			}
			if opts.background != tc.expectedBg {
				t.Errorf("background: expected %v, got %v", tc.expectedBg, opts.background)
			}
		})
	}
}

func TestLogger(t *testing.T) {
	// These tests are simple and don't test actual output
	// but ensure the logger initialization works

	// Test quiet logger
	quietLogger := newLogger(true, false)
	if !quietLogger.quiet {
		t.Error("Expected quiet logger to have quiet=true")
	}
	if quietLogger.debug {
		t.Error("Expected quiet logger to have debug=false")
	}

	// Test debug logger
	debugLogger := newLogger(false, true)
	if debugLogger.quiet {
		t.Error("Expected debug logger to have quiet=false")
	}
	if !debugLogger.debug {
		t.Error("Expected debug logger to have debug=true")
	}

	// Test default logger
	defaultLogger := newLogger(false, false)
	if defaultLogger.quiet {
		t.Error("Expected default logger to have quiet=false")
	}
	if defaultLogger.debug {
		t.Error("Expected default logger to have debug=false")
	}
}
