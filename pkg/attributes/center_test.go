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
			weightLbs:      GetDefaultWeight("6'8\""),
			wingspanInches: MustLengthToInches(GetDefaultWingspan("6'8\"")),
			want:           99,
		},
		{
			name:           "6'9\" (always 98)",
			heightInches:   MustLengthToInches("6'9"),
			weightLbs:      GetDefaultWeight("6'9\""),
			wingspanInches: MustLengthToInches(GetDefaultWingspan("6'9\"")),
			want:           98,
		},
		{
			name:           "6'10\" (always 96)",
			heightInches:   MustLengthToInches("6'10"),
			weightLbs:      GetDefaultWeight("6'10\""),
			wingspanInches: MustLengthToInches(GetDefaultWingspan("6'10\"")),
			want:           96,
		},
		// 6'11" height
		{
			name:           "6'11\" at 215 lbs",
			heightInches:   MustLengthToInches("6'11"),
			weightLbs:      215,
			wingspanInches: MustLengthToInches(GetDefaultWingspan("6'11\"")),
			want:           94,
		},
		// 7'0" height
		{
			name:           "7'0\" at 250 lbs",
			heightInches:   MustLengthToInches("7'0"),
			weightLbs:      250,
			wingspanInches: MustLengthToInches(GetDefaultWingspan("7'0\"")),
			want:           91,
		},
		// 7'1" height
		{
			name:           "7'1\" at 257 lbs",
			heightInches:   MustLengthToInches("7'1"),
			weightLbs:      257,
			wingspanInches: MustLengthToInches(GetDefaultWingspan("7'1\"")),
			want:           82,
		},
		// 7'2" height
		{
			name:           "7'2\" at 244 lbs",
			heightInches:   MustLengthToInches("7'2"),
			weightLbs:      244,
			wingspanInches: MustLengthToInches(GetDefaultWingspan("7'2\"")),
			want:           80,
		},
		// 7'3" height
		{
			name:           "7'3\" at 250 lbs",
			heightInches:   MustLengthToInches("7'3"),
			weightLbs:      250,
			wingspanInches: MustLengthToInches(GetDefaultWingspan("7'3\"")),
			want:           75,
		},
		// 7'4" height
		{
			name:           "7'4\" at 260 lbs",
			heightInches:   MustLengthToInches(CENTER_MAX_HEIGHT),
			weightLbs:      260,
			wingspanInches: MustLengthToInches(GetDefaultWingspan(CENTER_MAX_HEIGHT)),
			want:           70,
		},
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
			weightLbs:      GetDefaultWeight("6'7"),
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

func TestDrivingDunk2(t *testing.T) {
	// Test that DrivingDunk2 uses the deficit model: 99 - heightDeficit - wingspanDeficit - weightDeficit
	// Should match the original DrivingDunk values (at baseline weight 270 lbs)
	tests := []struct {
		name           string
		heightInches   int
		weightLbs      int
		wingspanInches int
		wantCap        int
	}{
		// Sample tests across different heights - should match DrivingDunk exactly
		{
			name:           "6'7\" with 6'7\" wingspan",
			heightInches:   MustLengthToInches("6'7"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("6'7"),
			wantCap:        95,
		},
		{
			name:           "6'7\" with 6'9\" wingspan",
			heightInches:   MustLengthToInches("6'7"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("6'9"),
			wantCap:        98,
		},
		{
			name:           "6'7\" with 7'1\" wingspan (max)",
			heightInches:   MustLengthToInches("6'7"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'1"),
			wantCap:        99,
		},
		{
			name:           "7'0\" with 7'0\" wingspan",
			heightInches:   MustLengthToInches("7'0"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'0"),
			wantCap:        83,
		},
		{
			name:           "7'0\" with 7'3\" wingspan",
			heightInches:   MustLengthToInches("7'0"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'3"),
			wantCap:        86,
		},
		{
			name:           "7'0\" with 7'6\" wingspan (max)",
			heightInches:   MustLengthToInches("7'0"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'6"),
			wantCap:        89,
		},
		{
			name:           "7'4\" with 7'4\" wingspan",
			heightInches:   MustLengthToInches("7'4"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'4"),
			wantCap:        66,
		},
		{
			name:           "7'4\" with 7'10\" wingspan (max)",
			heightInches:   MustLengthToInches("7'4"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'10"),
			wantCap:        70,
		},
		// Edge cases
		{
			name:           "6'10\" with 6'10\" wingspan",
			heightInches:   MustLengthToInches("6'10"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("6'10"),
			wantCap:        90,
		},
		{
			name:           "7'2\" with 7'2\" wingspan",
			heightInches:   MustLengthToInches("7'2"),
			weightLbs:      270,
			wingspanInches: MustLengthToInches("7'2"),
			wantCap:        72,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCap := DrivingDunk2(tt.heightInches, tt.weightLbs, tt.wingspanInches)
			assert.Equal(t, tt.wantCap, gotCap, "DrivingDunk2(%d, %d, %d) = %d, want %d",
				tt.heightInches, tt.weightLbs, tt.wingspanInches, gotCap, tt.wantCap)

			// Verify that DrivingDunk2 matches original DrivingDunk (at baseline weight)
			originalCap := DrivingDunk(tt.heightInches, tt.weightLbs, tt.wingspanInches)
			assert.Equal(t, originalCap, gotCap, "DrivingDunk2 should match DrivingDunk at baseline weight")
		})
	}
}
