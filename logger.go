// Copyright 2018 The goftp Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"fmt"
	"log"
)

// Logger is the interface that provides write logs.
type Logger interface {
	Print(sessionID string, message interface{})
	Printf(sessionID string, format string, v ...interface{})
	PrintCommand(sessionID string, command string, params string)
	PrintResponse(sessionID string, code int, message string)
}

// StdLogger implements Logger for write logs to standard output.
type StdLogger struct{}

// Print writes sessionID prefixed message to standard output.
func (logger *StdLogger) Print(sessionID string, message interface{}) {
	log.Printf("%s  %s", sessionID, message)
}

// Printf formats according to a format specifier and writes sessionID prefixed message to standard output.
func (logger *StdLogger) Printf(sessionID string, format string, v ...interface{}) {
	logger.Print(sessionID, fmt.Sprintf(format, v...))
}

// PrintCommand writes FTP command request logs to standard output.
func (logger *StdLogger) PrintCommand(sessionID string, command string, params string) {
	if command == "PASS" {
		log.Printf("%s > PASS ****", sessionID)
	} else {
		log.Printf("%s > %s %s", sessionID, command, params)
	}
}

// PrintResponse writes FTP command response logs to standard output.
func (logger *StdLogger) PrintResponse(sessionID string, code int, message string) {
	log.Printf("%s < %d %s", sessionID, code, message)
}

// DiscardLogger produces no output.
type DiscardLogger struct{}

// Print is nothing to do.
func (logger *DiscardLogger) Print(sessionID string, message interface{}) {}

// Printf is nothing to do.
func (logger *DiscardLogger) Printf(sessionID string, format string, v ...interface{}) {}

// PrintCommand is nothing to do.
func (logger *DiscardLogger) PrintCommand(sessionID string, command string, params string) {}

// PrintResponse is nothing to do.
func (logger *DiscardLogger) PrintResponse(sessionID string, code int, message string) {}
