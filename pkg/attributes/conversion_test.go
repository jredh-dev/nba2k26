// SPDX-License-Identifier: AGPL-3.0
// Copyright (C) 2025 NBA 2K26 Attribute System

package attributes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLengthToInches(t *testing.T) {
	tests := []struct {
		length string
		want   int
	}{
		{"6'7\"", 79},
		{"6'7", 79}, // Without trailing quote
		{"7'4\"", 88},
		{"7'4", 88}, // Without trailing quote
		{"5'9\"", 69},
		{"7'10\"", 94},
		{"7'10", 94}, // Without trailing quote
	}

	for _, tt := range tests {
		t.Run(tt.length, func(t *testing.T) {
			got, err := LengthToInches(tt.length)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMustLengthToInches(t *testing.T) {
	// Valid inputs should not panic
	assert.Equal(t, 79, MustLengthToInches("6'7\""))
	assert.Equal(t, 79, MustLengthToInches("6'7"))
	assert.Equal(t, 88, MustLengthToInches("7'4"))

	// Invalid input should panic
	assert.Panics(t, func() {
		MustLengthToInches("invalid")
	})
}

func TestInchesToLength(t *testing.T) {
	tests := []struct {
		inches int
		want   string
	}{
		{79, "6'7\""},
		{88, "7'4\""},
		{69, "5'9\""},
		{94, "7'10\""},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := InchesToLength(tt.inches)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestWeightToInt(t *testing.T) {
	tests := []struct {
		weight string
		want   int
	}{
		{"215", 215},
		{"290", 290},
		{"250", 250},
	}

	for _, tt := range tests {
		t.Run(tt.weight, func(t *testing.T) {
			got, err := WeightToInt(tt.weight)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMustWeightToInt(t *testing.T) {
	// Valid inputs should not panic
	assert.Equal(t, 215, MustWeightToInt("215"))
	assert.Equal(t, 290, MustWeightToInt("290"))

	// Invalid input should panic
	assert.Panics(t, func() {
		MustWeightToInt("invalid")
	})
}
