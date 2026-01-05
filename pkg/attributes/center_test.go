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
		name           string
		heightInches   int
		weightLbs      int
		wingspanInches int
		want           int
	}{
		{
			name:           "minimum size center",
			heightInches:   MustLengthToInches("6'7"),
			weightLbs:      215,
			wingspanInches: MustLengthToInches("6'7"),
			want:           99,
		},
		{
			name:           "maximum size center",
			heightInches:   MustLengthToInches("7'4"),
			weightLbs:      290,
			wingspanInches: MustLengthToInches("7'10"),
			want:           99,
		},
		{
			name:           "medium build",
			heightInches:   MustLengthToInches("7'0"),
			weightLbs:      250,
			wingspanInches: MustLengthToInches("7'4"),
			want:           99,
		},
		{
			name:           "short wingspan",
			heightInches:   MustLengthToInches("6'10"),
			weightLbs:      240,
			wingspanInches: MustLengthToInches("6'10"),
			want:           99,
		},
		{
			name:           "long wingspan",
			heightInches:   MustLengthToInches("6'10"),
			weightLbs:      240,
			wingspanInches: MustLengthToInches("7'6"),
			want:           99,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CloseShot(tt.heightInches, tt.weightLbs, tt.wingspanInches)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestPassAccuracy verifies Pass Accuracy is always 99 regardless of physical characteristics
func TestPassAccuracy(t *testing.T) {
	tests := []struct {
		name           string
		heightInches   int
		weightLbs      int
		wingspanInches int
		want           int
	}{
		{
			name:           "minimum size center",
			heightInches:   MustLengthToInches("6'7"),
			weightLbs:      215,
			wingspanInches: MustLengthToInches("6'7"),
			want:           99,
		},
		{
			name:           "maximum size center",
			heightInches:   MustLengthToInches("7'4"),
			weightLbs:      290,
			wingspanInches: MustLengthToInches("7'10"),
			want:           99,
		},
		{
			name:           "medium build",
			heightInches:   MustLengthToInches("7'0"),
			weightLbs:      250,
			wingspanInches: MustLengthToInches("7'4"),
			want:           99,
		},
		{
			name:           "light weight",
			heightInches:   MustLengthToInches("7'2"),
			weightLbs:      220,
			wingspanInches: MustLengthToInches("7'5"),
			want:           99,
		},
		{
			name:           "heavy weight",
			heightInches:   MustLengthToInches("7'2"),
			weightLbs:      285,
			wingspanInches: MustLengthToInches("7'5"),
			want:           99,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PassAccuracy(tt.heightInches, tt.weightLbs, tt.wingspanInches)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestDrivingLayup verifies Driving Layup caps based on height
// Pattern discovered: Height is primary factor (taller = lower cap)
func TestDrivingLayup(t *testing.T) {
	tests := []struct {
		name           string
		heightInches   int
		weightLbs      int
		wingspanInches int
		want           int
	}{
		{
			name:           "minimum height (6'7\") at minimum build",
			heightInches:   MustLengthToInches(CENTER_MIN_HEIGHT),
			weightLbs:      GetBounds(CENTER_MIN_HEIGHT).MinWeight,
			wingspanInches: MustLengthToInches(GetBounds(CENTER_MIN_HEIGHT).MinWingspan),
			want:           99,
		},
		{
			name:           "maximum height (7'4\") at maximum build",
			heightInches:   MustLengthToInches(CENTER_MAX_HEIGHT),
			weightLbs:      GetBounds(CENTER_MAX_HEIGHT).MaxWeight,
			wingspanInches: MustLengthToInches(GetBounds(CENTER_MAX_HEIGHT).MaxWingspan),
			want:           62,
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
			got := DrivingLayup(tt.heightInches, tt.weightLbs, tt.wingspanInches)
			assert.Equal(t, tt.want, got, "DrivingLayup(%d, %d, %d) = %d, want %d",
				tt.heightInches, tt.weightLbs, tt.wingspanInches, got, tt.want)
		})
	}
}
