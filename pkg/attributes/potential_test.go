// SPDX-License-Identifier: AGPL-3.0
// Copyright (C) 2025 NBA 2K26 Attribute System

package attributes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCloseShot verifies Close Shot is always 99 regardless of physical characteristics
func TestCloseShot(t *testing.T) {
	tests := []struct {
		name     string
		height   string
		weight   string
		wingspan string
		want     int
	}{
		{
			name:     "minimum size center",
			height:   "6'7\"",
			weight:   "215",
			wingspan: "6'7\"",
			want:     99,
		},
		{
			name:     "maximum size center",
			height:   "7'4\"",
			weight:   "290",
			wingspan: "7'10\"",
			want:     99,
		},
		{
			name:     "medium build",
			height:   "7'0\"",
			weight:   "250",
			wingspan: "7'4\"",
			want:     99,
		},
		{
			name:     "short wingspan",
			height:   "6'10\"",
			weight:   "240",
			wingspan: "6'10\"",
			want:     99,
		},
		{
			name:     "long wingspan",
			height:   "6'10\"",
			weight:   "240",
			wingspan: "7'6\"",
			want:     99,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CloseShot(tt.height, tt.weight, tt.wingspan)
			assert.Equal(t, tt.want, got, "CloseShot(%s, %s, %s) = %d, want %d",
				tt.height, tt.weight, tt.wingspan, got, tt.want)
		})
	}
}

// TestPassAccuracy verifies Pass Accuracy is always 99 regardless of physical characteristics
func TestPassAccuracy(t *testing.T) {
	tests := []struct {
		name     string
		height   string
		weight   string
		wingspan string
		want     int
	}{
		{
			name:     "minimum size center",
			height:   "6'7\"",
			weight:   "215",
			wingspan: "6'7\"",
			want:     99,
		},
		{
			name:     "maximum size center",
			height:   "7'4\"",
			weight:   "290",
			wingspan: "7'10\"",
			want:     99,
		},
		{
			name:     "medium build",
			height:   "7'0\"",
			weight:   "250",
			wingspan: "7'4\"",
			want:     99,
		},
		{
			name:     "light weight",
			height:   "7'2\"",
			weight:   "220",
			wingspan: "7'5\"",
			want:     99,
		},
		{
			name:     "heavy weight",
			height:   "7'2\"",
			weight:   "285",
			wingspan: "7'5\"",
			want:     99,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PassAccuracy(tt.height, tt.weight, tt.wingspan)
			assert.Equal(t, tt.want, got, "PassAccuracy(%s, %s, %s) = %d, want %d",
				tt.height, tt.weight, tt.wingspan, got, tt.want)
		})
	}
}

// TestDrivingLayup verifies Driving Layup caps based on height
// Pattern discovered: Height is primary factor (taller = lower cap)
func TestDrivingLayup(t *testing.T) {
	tests := []struct {
		name     string
		height   string
		weight   string
		wingspan string
		want     int
	}{
		{
			name:     "minimum height (6'7\")",
			height:   CENTER_MIN_HEIGHT,
			weight:   "215",
			wingspan: "6'7\"",
			want:     99,
		},
		{
			name:     "maximum height (7'4\")",
			height:   CENTER_MAX_HEIGHT,
			weight:   "290",
			wingspan: "7'10\"",
			want:     62,
		},
		// TODO: Add intermediate heights as you test them
		// {
		//     name:     "6'8\"",
		//     height:   "6'8\"",
		//     weight:   "220",
		//     wingspan: "6'9\"",
		//     want:     ??,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DrivingLayup(tt.height, tt.weight, tt.wingspan)
			assert.Equal(t, tt.want, got, "DrivingLayup(%s, %s, %s) = %d, want %d",
				tt.height, tt.weight, tt.wingspan, got, tt.want)
		})
	}
}
