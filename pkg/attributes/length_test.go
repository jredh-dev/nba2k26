// SPDX-License-Identifier: AGPL-3.0
// Copyright (C) 2025 NBA 2K26 Attribute System

package attributes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewLength(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantString  string
		wantInches  int
		shouldError bool
	}{
		{
			name:       "with trailing quote",
			input:      "6'7\"",
			wantString: "6'7\"",
			wantInches: 79,
		},
		{
			name:       "without trailing quote",
			input:      "6'7",
			wantString: "6'7\"",
			wantInches: 79,
		},
		{
			name:       "tall height with quote",
			input:      "7'4\"",
			wantString: "7'4\"",
			wantInches: 88,
		},
		{
			name:       "tall height without quote",
			input:      "7'4",
			wantString: "7'4\"",
			wantInches: 88,
		},
		{
			name:       "max center wingspan",
			input:      "7'10\"",
			wantString: "7'10\"",
			wantInches: 94,
		},
		{
			name:       "max center wingspan no quote",
			input:      "7'10",
			wantString: "7'10\"",
			wantInches: 94,
		},
		{
			name:        "invalid format",
			input:       "6-7",
			shouldError: true,
		},
		{
			name:        "empty string",
			input:       "",
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLength(tt.input)

			if tt.shouldError {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantString, got.String())
			assert.Equal(t, tt.wantInches, got.Inches())
		})
	}
}

func TestMustNewLength(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		l := MustNewLength("6'7")
		assert.Equal(t, "6'7\"", l.String())
		assert.Equal(t, 79, l.Inches())
	})

	t.Run("invalid input panics", func(t *testing.T) {
		assert.Panics(t, func() {
			MustNewLength("invalid")
		})
	})
}

func TestLengthString_String(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "standard format",
			input: "6'7",
			want:  "6'7\"",
		},
		{
			name:  "already has quote",
			input: "7'4\"",
			want:  "7'4\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := MustNewLength(tt.input)
			assert.Equal(t, tt.want, l.String())
		})
	}
}

func TestLengthString_StringShort(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "removes trailing quote",
			input: "6'7\"",
			want:  "6'7",
		},
		{
			name:  "no quote to remove",
			input: "7'4",
			want:  "7'4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := MustNewLength(tt.input)
			assert.Equal(t, tt.want, l.StringShort())
		})
	}
}

func TestLengthString_Inches(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantInches int
	}{
		{
			name:       "6'7\"",
			input:      "6'7",
			wantInches: 79,
		},
		{
			name:       "7'4\"",
			input:      "7'4",
			wantInches: 88,
		},
		{
			name:       "5'9\"",
			input:      "5'9",
			wantInches: 69,
		},
		{
			name:       "7'10\"",
			input:      "7'10",
			wantInches: 94,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := MustNewLength(tt.input)
			assert.Equal(t, tt.wantInches, l.Inches())
		})
	}
}

func TestLengthString_Feet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantFeet int
	}{
		{
			name:     "6'7\"",
			input:    "6'7",
			wantFeet: 6,
		},
		{
			name:     "7'4\"",
			input:    "7'4",
			wantFeet: 7,
		},
		{
			name:     "5'9\"",
			input:    "5'9",
			wantFeet: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := MustNewLength(tt.input)
			assert.Equal(t, tt.wantFeet, l.Feet())
		})
	}
}

func TestLengthString_RemainingInches(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		wantRemaining int
	}{
		{
			name:          "6'7\"",
			input:         "6'7",
			wantRemaining: 7,
		},
		{
			name:          "7'4\"",
			input:         "7'4",
			wantRemaining: 4,
		},
		{
			name:          "5'9\"",
			input:         "5'9",
			wantRemaining: 9,
		},
		{
			name:          "7'10\"",
			input:         "7'10",
			wantRemaining: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := MustNewLength(tt.input)
			assert.Equal(t, tt.wantRemaining, l.RemainingInches())
		})
	}
}

func TestParseLength(t *testing.T) {
	// Test that ParseLength is an alias for NewLength
	l1, err1 := NewLength("6'7")
	l2, err2 := ParseLength("6'7")

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.Equal(t, l1, l2)
}
