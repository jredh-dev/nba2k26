// SPDX-License-Identifier: AGPL-3.0
// Copyright (C) 2025 NBA 2K26 Attribute System

package attributes

import (
	"fmt"
	"strings"
)

// LengthString represents a height or wingspan measurement in the format 6'7" or 6'7
// This type provides standardized parsing and conversion to inches.
type LengthString string

// NewLength creates a LengthString from a string, normalizing the format
// Accepts: "6'7\"", "6'7", "7'4\"", "7'4"
func NewLength(s string) (LengthString, error) {
	// Normalize: ensure it has the trailing quote
	normalized := s
	if !strings.HasSuffix(s, "\"") {
		normalized = s + "\""
	}

	// Validate by attempting to parse
	_, err := LengthToInches(normalized)
	if err != nil {
		return "", err
	}

	return LengthString(normalized), nil
}

// MustNewLength creates a LengthString, panicking on invalid input
// Use this in tests and other contexts where the input is known to be valid
func MustNewLength(s string) LengthString {
	l, err := NewLength(s)
	if err != nil {
		panic(err)
	}
	return l
}

// Inches converts the LengthString to total inches
// Returns the number of inches (e.g., "6'7\"" → 79, "7'4\"" → 88)
func (l LengthString) Inches() int {
	inches, err := LengthToInches(string(l))
	if err != nil {
		// Should never happen if LengthString was created via New/MustNew
		panic(fmt.Sprintf("invalid LengthString: %s: %v", l, err))
	}
	return inches
}

// String returns the standardized string representation with trailing quote
// Always returns format: "6'7\"", "7'4\"", etc.
func (l LengthString) String() string {
	return string(l)
}

// StringShort returns the string representation without trailing quote
// Returns format: "6'7", "7'4", etc.
func (l LengthString) StringShort() string {
	s := string(l)
	return strings.TrimSuffix(s, "\"")
}

// Feet returns the number of feet
func (l LengthString) Feet() int {
	return l.Inches() / 12
}

// RemainingInches returns the inches component after feet
// For "6'7\"", returns 7
func (l LengthString) RemainingInches() int {
	return l.Inches() % 12
}

// ParseLength is an alias for NewLength for backwards compatibility
func ParseLength(s string) (LengthString, error) {
	return NewLength(s)
}
