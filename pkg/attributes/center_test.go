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
		{
			name:           "6'8\" (always 99)",
			heightInches:   MustLengthToInches("6'8"),
			weightLbs:      240,
			wingspanInches: MustLengthToInches("6'8"),
			want:           99,
		},
		{
			name:           "6'9\" (always 98)",
			heightInches:   MustLengthToInches("6'9"),
			weightLbs:      250,
			wingspanInches: MustLengthToInches("6'9"),
			want:           98,
		},
		{
			name:           "6'10\" (always 96)",
			heightInches:   MustLengthToInches("6'10"),
			weightLbs:      250,
			wingspanInches: MustLengthToInches("6'10"),
			want:           96,
		},
		// 6'11" height - weight variations
		{
			name:           "6'11\" at 215 lbs (minimum weight)",
			heightInches:   MustLengthToInches("6'11"),
			weightLbs:      215,
			wingspanInches: MustLengthToInches("6'11"),
			want:           94,
		},
		{
			name:           "6'11\" at 290 lbs (maximum weight)",
			heightInches:   MustLengthToInches("6'11"),
			weightLbs:      290,
			wingspanInches: MustLengthToInches("6'11"),
			want:           92,
		},
		// 7'0" height - weight variations
		{
			name:           "7'0\" at 215 lbs (minimum weight)",
			heightInches:   MustLengthToInches("7'0"),
			weightLbs:      215,
			wingspanInches: MustLengthToInches("7'0"),
			want:           93,
		},
		{
			name:           "7'0\" at 290 lbs (maximum weight)",
			heightInches:   MustLengthToInches("7'0"),
			weightLbs:      290,
			wingspanInches: MustLengthToInches("7'0"),
			want:           89,
		},
		// 7'1" height - weight variations
		{
			name:           "7'1\" at 220 lbs (minimum weight)",
			heightInches:   MustLengthToInches("7'1"),
			weightLbs:      220,
			wingspanInches: MustLengthToInches("7'1"),
			want:           86,
		},
		{
			name:           "7'1\" at 257 lbs",
			heightInches:   MustLengthToInches("7'1"),
			weightLbs:      257,
			wingspanInches: MustLengthToInches("7'1"),
			want:           82,
		},
		{
			name:           "7'1\" at 290 lbs (maximum weight)",
			heightInches:   MustLengthToInches("7'1"),
			weightLbs:      290,
			wingspanInches: MustLengthToInches("7'1"),
			want:           77,
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

func TestDrivingDunk(t *testing.T) {
	tests := []struct {
		name           string
		heightInches   int
		weightLbs      int
		wingspanInches int
		want           int
	}{
		// 6'7" height - wingspan variations
		{
			name:           "6'7\" with 6'7\" wingspan",
			heightInches:   MustLengthToInches("6'7"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("6'7"),
			want:           95,
		},
		{
			name:           "6'7\" with 6'8\" wingspan",
			heightInches:   MustLengthToInches("6'7"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("6'8"),
			want:           97,
		},
		{
			name:           "6'7\" with 6'9\" wingspan",
			heightInches:   MustLengthToInches("6'7"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("6'9"),
			want:           98,
		},
		{
			name:           "6'7\" with 7'1\" wingspan",
			heightInches:   MustLengthToInches("6'7"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'1"),
			want:           99,
		},
		// 6'8" height - wingspan variations
		{
			name:           "6'8\" with 6'8\" wingspan",
			heightInches:   MustLengthToInches("6'8"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("6'8"),
			want:           94,
		},
		{
			name:           "6'8\" with 6'10\" wingspan",
			heightInches:   MustLengthToInches("6'8"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("6'10"),
			want:           96,
		},
		{
			name:           "6'8\" with 7'0\" wingspan",
			heightInches:   MustLengthToInches("6'8"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'0"),
			want:           99,
		},
		// 6'9" height - wingspan variations
		{
			name:           "6'9\" with 6'9\" wingspan",
			heightInches:   MustLengthToInches("6'9"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("6'9"),
			want:           92,
		},
		{
			name:           "6'9\" with 7'0\" wingspan",
			heightInches:   MustLengthToInches("6'9"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'0"),
			want:           95,
		},
		{
			name:           "6'9\" with 7'3\" wingspan",
			heightInches:   MustLengthToInches("6'9"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'3"),
			want:           99,
		},
		// 6'10" height - wingspan variations
		{
			name:           "6'10\" with 6'10\" wingspan",
			heightInches:   MustLengthToInches("6'10"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("6'10"),
			want:           90,
		},
		{
			name:           "6'10\" with 7'1\" wingspan",
			heightInches:   MustLengthToInches("6'10"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'1"),
			want:           93,
		},
		{
			name:           "6'10\" with 7'4\" wingspan",
			heightInches:   MustLengthToInches("6'10"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'4"),
			want:           96,
		},
		// 6'11" height - wingspan variations
		{
			name:           "6'11\" with 6'11\" wingspan",
			heightInches:   MustLengthToInches("6'11"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("6'11"),
			want:           86,
		},
		{
			name:           "6'11\" with 7'2\" wingspan",
			heightInches:   MustLengthToInches("6'11"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'2"),
			want:           89,
		},
		{
			name:           "6'11\" with 7'5\" wingspan",
			heightInches:   MustLengthToInches("6'11"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'5"),
			want:           92,
		},
		// 7'0" height - wingspan variations
		{
			name:           "7'0\" with 7'0\" wingspan",
			heightInches:   MustLengthToInches("7'0"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'0"),
			want:           83,
		},
		{
			name:           "7'0\" with 7'3\" wingspan",
			heightInches:   MustLengthToInches("7'0"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'3"),
			want:           86,
		},
		{
			name:           "7'0\" with 7'6\" wingspan",
			heightInches:   MustLengthToInches("7'0"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'6"),
			want:           89,
		},
		// 7'1" height - wingspan variations
		{
			name:           "7'1\" with 7'1\" wingspan",
			heightInches:   MustLengthToInches("7'1"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'1"),
			want:           77,
		},
		{
			name:           "7'1\" with 7'4\" wingspan",
			heightInches:   MustLengthToInches("7'1"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'4"),
			want:           80,
		},
		{
			name:           "7'1\" with 7'7\" wingspan",
			heightInches:   MustLengthToInches("7'1"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'7"),
			want:           82,
		},
		// 7'2" height - wingspan variations
		{
			name:           "7'2\" with 7'2\" wingspan",
			heightInches:   MustLengthToInches("7'2"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'2"),
			want:           72,
		},
		{
			name:           "7'2\" with 7'5\" wingspan",
			heightInches:   MustLengthToInches("7'2"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'5"),
			want:           74,
		},
		{
			name:           "7'2\" with 7'8\" wingspan",
			heightInches:   MustLengthToInches("7'2"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'8"),
			want:           77,
		},
		// 7'3" height - wingspan variations
		{
			name:           "7'3\" with 7'3\" wingspan",
			heightInches:   MustLengthToInches("7'3"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'3"),
			want:           68,
		},
		{
			name:           "7'3\" with 7'6\" wingspan",
			heightInches:   MustLengthToInches("7'3"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'6"),
			want:           70,
		},
		{
			name:           "7'3\" with 7'8\" wingspan",
			heightInches:   MustLengthToInches("7'3"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'8"),
			want:           72,
		},
		// 7'4" height - wingspan variations
		{
			name:           "7'4\" with 7'4\" wingspan",
			heightInches:   MustLengthToInches("7'4"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'4"),
			want:           66,
		},
		{
			name:           "7'4\" with 7'7\" wingspan",
			heightInches:   MustLengthToInches("7'4"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'7"),
			want:           68,
		},
		{
			name:           "7'4\" with 7'10\" wingspan",
			heightInches:   MustLengthToInches("7'4"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'10"),
			want:           70,
		},
		// NOTE: These tests use baseline weight (270). Weight also affects this attribute.
		// TODO: Add weight variation tests once modifier system is implemented
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DrivingDunk(tt.heightInches, tt.weightLbs, tt.wingspanInches)
			assert.Equal(t, tt.want, got, "DrivingDunk(%d, %d, %d) = %d, want %d",
				tt.heightInches, tt.weightLbs, tt.wingspanInches, got, tt.want)
		})
	}
}
