// SPDX-License-Identifier: AGPL-3.0
// Copyright (C) 2025 NBA 2K26 Attribute System

package attributes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeightToInches(t *testing.T) {
	tests := []struct {
		height string
		want   int
	}{
		{"6'7\"", 79},
		{"7'4\"", 88},
		{"5'9\"", 69},
		{"7'10\"", 94},
	}

	for _, tt := range tests {
		t.Run(tt.height, func(t *testing.T) {
			got, err := HeightToInches(tt.height)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInchesToHeight(t *testing.T) {
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
			got := InchesToHeight(tt.inches)
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
