// SPDX-License-Identifier: AGPL-3.0
// Copyright (C) 2025 NBA 2K26 Attribute System

package attributes

import "fmt"

// PhysicalBounds represents the valid weight and wingspan ranges for a given height
type PhysicalBounds struct {
	MinWeight       int
	MaxWeight       int
	DefaultWeight   int // The default/baseline weight for this height in-game
	MinWingspan     string
	MaxWingspan     string
	DefaultWingspan string // The default wingspan for this height in-game
}

// CenterBounds maps each valid center height to its weight/wingspan constraints
// This data is discovered through in-game testing
var CenterBounds = map[string]PhysicalBounds{
	"6'7\"": {
		MinWeight:       215,
		MaxWeight:       270,
		DefaultWeight:   243,
		MinWingspan:     "6'7\"",
		MaxWingspan:     "7'1\"",
		DefaultWingspan: "6'10\"",
	},
	"6'8\"": {
		MinWeight:       215,
		MaxWeight:       275,
		DefaultWeight:   245,
		MinWingspan:     "6'8\"",
		MaxWingspan:     "7'2\"",
		DefaultWingspan: "6'11\"",
	},
	"6'9\"": {
		MinWeight:       215,
		MaxWeight:       285,
		DefaultWeight:   250,
		MinWingspan:     "6'9\"",
		MaxWingspan:     "7'3\"",
		DefaultWingspan: "7'0\"",
	},
	"6'10\"": {
		MinWeight:       215,
		MaxWeight:       285,
		DefaultWeight:   250,
		MinWingspan:     "6'10\"",
		MaxWingspan:     "7'4\"",
		DefaultWingspan: "7'1\"",
	},
	"6'11\"": {
		MinWeight:       215,
		MaxWeight:       290,
		DefaultWeight:   253,
		MinWingspan:     "6'11\"",
		MaxWingspan:     "7'5\"",
		DefaultWingspan: "7'2\"",
	},
	"7'0\"": {
		MinWeight:       215,
		MaxWeight:       290,
		DefaultWeight:   253,
		MinWingspan:     "7'0\"",
		MaxWingspan:     "7'6\"",
		DefaultWingspan: "7'3\"",
	},
	"7'1\"": {
		MinWeight:       220,
		MaxWeight:       290,
		DefaultWeight:   255,
		MinWingspan:     "7'1\"",
		MaxWingspan:     "7'7\"",
		DefaultWingspan: "7'4\"",
	},
	"7'2\"": {
		MinWeight:       220,
		MaxWeight:       290,
		DefaultWeight:   255,
		MinWingspan:     "7'2\"",
		MaxWingspan:     "7'8\"",
		DefaultWingspan: "7'5\"",
	},
	"7'3\"": {
		MinWeight:       230,
		MaxWeight:       290,
		DefaultWeight:   260,
		MinWingspan:     "7'3\"",
		MaxWingspan:     "7'9\"",
		DefaultWingspan: "7'6\"",
	},
	"7'4\"": {
		MinWeight:       230,
		MaxWeight:       290,
		DefaultWeight:   260,
		MinWingspan:     "7'4\"",
		MaxWingspan:     "7'10\"",
		DefaultWingspan: "7'7\"",
	},
}

// GetBounds returns the physical bounds for a given height, or nil if height is invalid
func GetBounds(height string) *PhysicalBounds {
	if bounds, ok := CenterBounds[height]; ok {
		return &bounds
	}
	return nil
}

// GetDefaultWeight returns the default/baseline weight for a given height
// Returns -1 if height is invalid
func GetDefaultWeight(height string) int {
	bounds := GetBounds(height)
	if bounds == nil {
		return -1
	}
	return bounds.DefaultWeight
}

// GetDefaultWeightForInches returns the default weight for a height in inches
func GetDefaultWeightForInches(heightInches int) int {
	heightStr := InchesToLength(heightInches)
	return GetDefaultWeight(heightStr)
}

// GetDefaultWingspan returns the default wingspan for a given height
// Returns empty string if height is invalid
func GetDefaultWingspan(height string) string {
	bounds := GetBounds(height)
	if bounds == nil {
		return ""
	}
	return bounds.DefaultWingspan
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
