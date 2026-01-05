// SPDX-License-Identifier: AGPL-3.0
// Copyright (C) 2025 NBA 2K26 Attribute System

package attributes

import "fmt"

// PhysicalBounds represents the valid weight and wingspan ranges for a given height
type PhysicalBounds struct {
	Height      string
	MinWeight   int
	MaxWeight   int
	MinWingspan string
	MaxWingspan string
}

// CenterHeightBounds maps each valid center height to its weight/wingspan constraints
// This data is discovered through in-game testing
var CenterHeightBounds = map[string]PhysicalBounds{
	"6'7\"": {
		Height:      "6'7\"",
		MinWeight:   215,
		MaxWeight:   270,
		MinWingspan: "6'7\"",
		MaxWingspan: "7'1\"",
	},
	// TODO: Add 6'8", 6'9", 6'10", 6'11", 7'0", 7'1", 7'2", 7'3"
	"7'4\"": {
		Height:      "7'4\"",
		MinWeight:   215, // TODO: Verify this
		MaxWeight:   290,
		MinWingspan: "6'7\"", // TODO: Verify this
		MaxWingspan: "7'10\"",
	},
}

// GetBounds returns the physical bounds for a given height, or nil if height is invalid
func GetBounds(height string) *PhysicalBounds {
	if bounds, ok := CenterHeightBounds[height]; ok {
		return &bounds
	}
	return nil
}

// ValidateCenter checks if a height/weight/wingspan combination is valid for a Center
func ValidateCenter(height, weight, wingspan string) bool {
	bounds := GetBounds(height)
	if bounds == nil {
		return false
	}

	// Parse weight as int
	w := 0
	_, err := fmt.Sscanf(weight, "%d", &w)
	if err != nil {
		return false
	}

	// Check weight range
	if w < bounds.MinWeight || w > bounds.MaxWeight {
		return false
	}

	// TODO: Add wingspan validation
	// Need to convert wingspan strings to comparable format

	return true
}
