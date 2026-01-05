// SPDX-License-Identifier: AGPL-3.0
// Copyright (C) 2025 NBA 2K26 Attribute System

package attributes

import "fmt"

// PhysicalBounds represents the valid weight and wingspan ranges for a given height
type PhysicalBounds struct {
	MinWeight   int
	MaxWeight   int
	MinWingspan string
	MaxWingspan string
}

// CenterHeightBounds maps each valid center height to its weight/wingspan constraints
// This data is discovered through in-game testing
var CenterHeightBounds = map[string]PhysicalBounds{
	"6'7\"": {
		MinWeight:   215,
		MaxWeight:   270,
		MinWingspan: "6'7\"",
		MaxWingspan: "7'1\"",
	},
	"6'8\"": {
		MinWeight:   215,
		MaxWeight:   275,
		MinWingspan: "6'8\"",
		MaxWingspan: "7'2\"",
	},
	"6'9\"": {
		MinWeight:   215,
		MaxWeight:   285,
		MinWingspan: "6'9\"",
		MaxWingspan: "7'3\"",
	},
	"6'10\"": {
		MinWeight:   215,
		MaxWeight:   285,
		MinWingspan: "6'10\"",
		MaxWingspan: "7'4\"",
	},
	"6'11\"": {
		MinWeight:   215,
		MaxWeight:   290,
		MinWingspan: "6'11\"",
		MaxWingspan: "7'5\"",
	},
	"7'0\"": {
		MinWeight:   215,
		MaxWeight:   290,
		MinWingspan: "7'0\"",
		MaxWingspan: "7'6\"",
	},
	"7'1\"": {
		MinWeight:   220,
		MaxWeight:   290,
		MinWingspan: "7'1\"",
		MaxWingspan: "7'7\"",
	},
	"7'2\"": {
		MinWeight:   220,
		MaxWeight:   290,
		MinWingspan: "7'2\"",
		MaxWingspan: "7'8\"",
	},
	"7'3\"": {
		MinWeight:   230,
		MaxWeight:   290,
		MinWingspan: "7'3\"",
		MaxWingspan: "7'9\"",
	},
	"7'4\"": {
		MinWeight:   230,
		MaxWeight:   290,
		MinWingspan: "7'4\"",
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
