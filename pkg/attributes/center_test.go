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

// TestDrivingLayup verifies Driving Layup caps based on height and weight
// Pattern discovered:
// - Height is primary factor (taller = lower cap)
// - At 7'4", weight also matters (heavier = lower cap)
func TestDrivingLayup(t *testing.T) {
	tests := []struct {
		name           string
		heightInches   int
		weightLbs      int
		wingspanInches int
		want           int
	}{
		{
			name:           "minimum height (6'7\") at minimum weight",
			heightInches:   MustLengthToInches(CENTER_MIN_HEIGHT),
			weightLbs:      GetBounds(CENTER_MIN_HEIGHT).MinWeight,
			wingspanInches: MustLengthToInches(GetBounds(CENTER_MIN_HEIGHT).MinWingspan),
			want:           99,
		},
		// 7'2" height - weight variations (CONFIRMED)
		{
			name:           "7'2\" at 223 lbs (lightest tested)",
			heightInches:   MustLengthToInches("7'2"),
			weightLbs:      223,
			wingspanInches: MustLengthToInches("7'2"),
			want:           84,
		},
		{
			name:           "7'2\" at 244 lbs",
			heightInches:   MustLengthToInches("7'2"),
			weightLbs:      244,
			wingspanInches: MustLengthToInches("7'2"),
			want:           80,
		},
		{
			name:           "7'2\" at 269 lbs",
			heightInches:   MustLengthToInches("7'2"),
			weightLbs:      269,
			wingspanInches: MustLengthToInches("7'2"),
			want:           75,
		},
		{
			name:           "7'2\" at 290 lbs (maximum weight)",
			heightInches:   MustLengthToInches("7'2"),
			weightLbs:      290,
			wingspanInches: MustLengthToInches("7'2"),
			want:           71,
		},
		// 7'3" height - weight variations
		{
			name:           "7'3\" at 230 lbs (minimum weight)",
			heightInches:   MustLengthToInches("7'3"),
			weightLbs:      230,
			wingspanInches: MustLengthToInches("7'3"),
			want:           80,
		},
		{
			name:           "7'3\" at 250 lbs",
			heightInches:   MustLengthToInches("7'3"),
			weightLbs:      250,
			wingspanInches: MustLengthToInches("7'3"),
			want:           75,
		},
		{
			name:           "7'3\" at 270 lbs",
			heightInches:   MustLengthToInches("7'3"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'3"),
			want:           70,
		},
		{
			name:           "7'3\" at 290 lbs (maximum weight)",
			heightInches:   MustLengthToInches("7'3"),
			weightLbs:      290,
			wingspanInches: MustLengthToInches("7'3"),
			want:           64,
		},
		// 7'4" height - weight variations
		{
			name:           "7'4\" at 230 lbs (lightest tested)",
			heightInches:   MustLengthToInches(CENTER_MAX_HEIGHT),
			weightLbs:      230,
			wingspanInches: MustLengthToInches("7'5"),
			want:           77,
		},
		{
			name:           "7'4\" at 232 lbs",
			heightInches:   MustLengthToInches(CENTER_MAX_HEIGHT),
			weightLbs:      232,
			wingspanInches: MustLengthToInches("7'5"),
			want:           76,
		},
		{
			name:           "7'4\" at 240 lbs",
			heightInches:   MustLengthToInches(CENTER_MAX_HEIGHT),
			weightLbs:      240,
			wingspanInches: MustLengthToInches("7'5"),
			want:           74,
		},
		{
			name:           "7'4\" at 265 lbs",
			heightInches:   MustLengthToInches(CENTER_MAX_HEIGHT),
			weightLbs:      265,
			wingspanInches: MustLengthToInches("7'5"),
			want:           68,
		},
		{
			name:           "7'4\" at 287 lbs",
			heightInches:   MustLengthToInches(CENTER_MAX_HEIGHT),
			weightLbs:      287,
			wingspanInches: MustLengthToInches("7'5"),
			want:           63,
		},
		{
			name:           "7'4\" at 290 lbs (maximum weight)",
			heightInches:   MustLengthToInches(CENTER_MAX_HEIGHT),
			weightLbs:      GetBounds(CENTER_MAX_HEIGHT).MaxWeight,
			wingspanInches: MustLengthToInches(GetBounds(CENTER_MAX_HEIGHT).MaxWingspan),
			want:           62,
		},
		// TODO: Add intermediate heights as you test them
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DrivingLayup(tt.heightInches, tt.weightLbs, tt.wingspanInches)
			assert.Equal(t, tt.want, got, "DrivingLayup(%d, %d, %d) = %d, want %d",
				tt.heightInches, tt.weightLbs, tt.wingspanInches, got, tt.want)
		})
	}
}
