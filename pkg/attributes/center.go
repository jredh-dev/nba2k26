// SPDX-License-Identifier: AGPL-3.0
// Copyright (C) 2025 NBA 2K26 Attribute System

// Package attributes provides attribute cap calculation functions for NBA 2K26 character creation.
// Each function takes height, weight, and wingspan and returns the attribute cap (0-99).
package attributes

// Center position physical characteristic bounds
// Note: Weight and wingspan ranges are HEIGHT-DEPENDENT
// These constants represent the absolute extremes across all heights
const (
	CENTER_MIN_HEIGHT = "6'7\""
	CENTER_MAX_HEIGHT = "7'4\""
)

// Height-specific bounds (examples - add more as you discover them)
// Format: At height X, weight range is Y-Z, wingspan range is A-B
//
// 6'7" Center:
//   - Weight: 215 - 270 lbs
//   - Wingspan: 6'7" - 7'1"
//
// 7'4" Center:
//   - Weight: 215 - 290 lbs (TODO: verify min weight)
//   - Wingspan: 6'7" - 7'10" (TODO: verify min wingspan)
//
// TODO: Document bounds for all intermediate heights (6'8", 6'9", etc.)

// Center position attribute calculators
// All functions take measurements as integers:
//   - heightInches: height in inches (e.g., 79 for 6'7")
//   - weightLbs: weight in pounds (e.g., 215)
//   - wingspanInches: wingspan in inches (e.g., 79 for 6'7")

// CloseShot calculates the Close Shot attribute cap for a Center.
// This attribute is always 99 regardless of physical characteristics.
func CloseShot(heightInches, weightLbs, wingspanInches int) int {
	return 99
}

// PassAccuracy calculates the Pass Accuracy attribute cap for a Center.
// This attribute is always 99 regardless of physical characteristics.
func PassAccuracy(heightInches, weightLbs, wingspanInches int) int {
	return 99
}

// DrivingLayup calculates the Driving Layup attribute cap for a Center.
// Testing notes:
// - At minimum height (79" / 6'7"): cap is 99 (weight doesn't matter)
// - At maximum height (88" / 7'4"): cap is 62-77 (weight DOES matter: heavier = lower)
// - Pattern: Height is primary factor, weight creates penalties at 6'11"+
// - Wingspan does not affect this attribute
// - Data-driven implementation based on NBA2KLab API scraped data (903 builds)
func DrivingLayup(heightInches, weightLbs, _ int) int {
	// Lookup table generated from scraped data
	// Heights 6'7"-6'10" are weight-independent
	// Heights 6'11"+ have weight-dependent penalties

	type weightThreshold struct {
		maxWeight int
		value     int
	}

	// Weight-dependent heights use threshold tables
	// Format: if weight <= maxWeight, return value (check in order)
	lookupTable := map[int][]weightThreshold{
		79: {{99999, 99}}, // 6'7": always 99
		80: {{99999, 99}}, // 6'8": always 99
		81: {{99999, 98}}, // 6'9": always 98
		82: {{99999, 96}}, // 6'10": always 96
		83: { // 6'11": 93-94 range
			{250, 94},
			{99999, 93},
		},
		84: { // 7'0": 90-93 range
			{225, 93},
			{240, 92},
			{260, 91},
			{99999, 90},
		},
		85: { // 7'1": 79-86 range
			{225, 86},   // 220-225 = 86
			{230, 85},   // 226-230 = 85
			{240, 84},   // 231-240 = 84
			{245, 83},   // 241-245 = 83
			{260, 82},   // 246-260 = 82
			{270, 80},   // 261-270 = 80-81 (265=80, 260=81 in data, use 80)
			{99999, 79}, // 271+ = 79
		},
		86: { // 7'2": 73-84 range
			{220, 84},
			{225, 83},
			{230, 82},
			{235, 81},
			{245, 80}, // 236-245 = 80
			{250, 78},
			{255, 77},
			{260, 76},
			{265, 75},
			{275, 74}, // 266-275 = 74
			{99999, 73},
		},
		87: { // 7'3": 64-80 range
			{230, 80},
			{235, 78},
			{240, 77},
			{245, 76},
			{250, 75},
			{255, 73},
			{260, 72},
			{265, 71},
			{270, 70},
			{275, 68},
			{280, 67},
			{285, 66},
			{99999, 64},
		},
		88: { // 7'4": 62-77 range
			{230, 77},
			{235, 76},
			{240, 74},
			{245, 73},
			{250, 72},
			{255, 71},
			{260, 70},
			{265, 68},
			{270, 67},
			{275, 66},
			{280, 65},
			{285, 64},
			{99999, 62},
		},
	}

	thresholds, exists := lookupTable[heightInches]
	if !exists {
		return 0 // Invalid height
	}

	// Find first threshold where weight <= maxWeight
	for _, t := range thresholds {
		if weightLbs <= t.maxWeight {
			return t.value
		}
	}

	return 0 // Should never reach here if table is correct
}

// DrivingDunk calculates the Driving Dunk attribute cap for a Center.
// NOTE: Weight also affects this attribute - current implementation uses baseline weight.
// TODO: Implement weight modifiers (additive system: height_base + wingspan_modifier + weight_modifier)
func DrivingDunk(heightInches, weightLbs, wingspanInches int) int {
	// Current implementation: wingspan variations at baseline weight for each height
	switch heightInches {
	case MustLengthToInches("6'7"): // 79"
		switch wingspanInches {
		case MustLengthToInches("6'7"):
			return 95
		case MustLengthToInches("6'8"):
			return 97
		case MustLengthToInches("6'9"):
			return 98
		case MustLengthToInches("7'1"):
			return 99
		}
	case MustLengthToInches("6'8"): // 80"
		switch wingspanInches {
		case MustLengthToInches("6'8"):
			return 94
		case MustLengthToInches("6'9"):
			return 95
		case MustLengthToInches("6'10"):
			return 96
		case MustLengthToInches("6'11"):
			return 98
		case MustLengthToInches("7'0"):
			return 99
		case MustLengthToInches("7'2"):
			return 99
		}
	case MustLengthToInches("6'9"): // 81"
		switch wingspanInches {
		case MustLengthToInches("6'9"):
			return 92
		case MustLengthToInches("6'10"):
			return 93
		case MustLengthToInches("6'11"):
			return 94
		case MustLengthToInches("7'0"):
			return 95
		case MustLengthToInches("7'1"):
			return 96
		case MustLengthToInches("7'2"):
			return 98
		case MustLengthToInches("7'3"):
			return 99
		}
	case MustLengthToInches("6'10"): // 82"
		switch wingspanInches {
		case MustLengthToInches("6'10"):
			return 90
		case MustLengthToInches("6'11"):
			return 91
		case MustLengthToInches("7'0"):
			return 92
		case MustLengthToInches("7'1"):
			return 93
		case MustLengthToInches("7'2"):
			return 94
		case MustLengthToInches("7'3"):
			return 95
		case MustLengthToInches("7'4"):
			return 96
		}
	case MustLengthToInches("6'11"): // 83"
		switch wingspanInches {
		case MustLengthToInches("6'11"):
			switch {
			case weightLbs <= 290:
				return 85
			case weightLbs <= 268:
				return 86
			case weightLbs <= 225:
				return 87
			}
		case MustLengthToInches("7'0"):
			switch {
			case weightLbs <= 290:
				return 86
			case weightLbs <= 271:
				return 87
			case weightLbs <= 229:
				return 88
			}
		case MustLengthToInches("7'1"):
			return 88
		case MustLengthToInches("7'2"):
			return 89
		case MustLengthToInches("7'3"):
			return 90
		case MustLengthToInches("7'4"):
			return 91
		case MustLengthToInches("7'5"):
			return 92
		}
	case MustLengthToInches("7'0"): // 84"
		switch wingspanInches {
		case MustLengthToInches("7'0"):
			return 83
		case MustLengthToInches("7'1"):
			return 84
		case MustLengthToInches("7'2"):
			return 85
		case MustLengthToInches("7'3"):
			return 86
		case MustLengthToInches("7'4"):
			return 87
		case MustLengthToInches("7'5"):
			return 88
		case MustLengthToInches("7'6"):
			return 89
		}
	case MustLengthToInches("7'1"): // 85"
		switch wingspanInches {
		case MustLengthToInches("7'1"):
			return 77
		case MustLengthToInches("7'2"):
			return 78
		case MustLengthToInches("7'3"):
			return 79
		case MustLengthToInches("7'4"):
			return 80
		case MustLengthToInches("7'5"):
			return 81
		case MustLengthToInches("7'6"):
			return 82
		case MustLengthToInches("7'7"):
			return 82
		}
	case MustLengthToInches("7'2"): // 86"
		switch wingspanInches {
		case MustLengthToInches("7'2"):
			return 72
		case MustLengthToInches("7'3"):
			return 72
		case MustLengthToInches("7'4"):
			return 73
		case MustLengthToInches("7'5"):
			return 74
		case MustLengthToInches("7'6"):
			return 75
		case MustLengthToInches("7'7"):
			return 76
		case MustLengthToInches("7'8"):
			return 77
		}
	case MustLengthToInches("7'3"): // 87"
		switch wingspanInches {
		case MustLengthToInches("7'3"):
			return 68
		case MustLengthToInches("7'4"):
			return 69
		case MustLengthToInches("7'5"):
			return 69
		case MustLengthToInches("7'6"):
			return 70
		case MustLengthToInches("7'7"):
			return 71
		case MustLengthToInches("7'8"):
			return 72
		case MustLengthToInches("7'9"):
			return 72
		}
	case MustLengthToInches("7'4"): // 88"
		switch wingspanInches {
		case MustLengthToInches("7'4"):
			return 66
		case MustLengthToInches("7'5"):
			return 67
		case MustLengthToInches("7'6"):
			return 68
		case MustLengthToInches("7'7"):
			return 68
		case MustLengthToInches("7'8"):
			return 69
		case MustLengthToInches("7'9"):
			return 70
		case MustLengthToInches("7'10"):
			return 70
		}
	}
	return 0
}

// DrivingDunk2 calculates Driving Dunk using an additive deficit model.
// Formula: 99 - heightDeficit - wingspanDeficit - weightDeficit = Final Cap
//
// Current implementation: height and wingspan deficits are complete.
// Weight deficit will be added after testing weight variations.
func DrivingDunk2(heightInches, weightLbs, wingspanInches int) int {
	var heightDeficit int
	var wingspanDeficit int
	var weightDeficit int

	// Step 1: Calculate height deficit (how much height lowers the cap from 99)
	switch heightInches {
	case MustLengthToInches("6'7"): // 79"
		heightDeficit = 0 // Can reach 99
	case MustLengthToInches("6'8"): // 80"
		heightDeficit = 0 // Can reach 99
	case MustLengthToInches("6'9"): // 81"
		heightDeficit = 0 // Can reach 99
	case MustLengthToInches("6'10"): // 82"
		heightDeficit = 3 // Max 96
	case MustLengthToInches("6'11"): // 83"
		heightDeficit = 7 // Max 92
	case MustLengthToInches("7'0"): // 84"
		heightDeficit = 10 // Max 89
	case MustLengthToInches("7'1"): // 85"
		heightDeficit = 17 // Max 82
	case MustLengthToInches("7'2"): // 86"
		heightDeficit = 22 // Max 77
	case MustLengthToInches("7'3"): // 87"
		heightDeficit = 27 // Max 72
	case MustLengthToInches("7'4"): // 88"
		heightDeficit = 29 // Max 70
	default:
		heightDeficit = 99 // Unknown height
	}

	// Step 2: Calculate wingspan deficit (shorter wingspan = larger deficit)
	// Based on minimum wingspan for each height
	switch heightInches {
	case MustLengthToInches("6'7"): // 79"
		minWingspan := MustLengthToInches("6'7")
		switch wingspanInches {
		case minWingspan:
			wingspanDeficit = 4
		case minWingspan + 1:
			wingspanDeficit = 2
		case minWingspan + 2:
			wingspanDeficit = 1
		default: // >= minWingspan + 4 (7'1")
			wingspanDeficit = 0
		}
	case MustLengthToInches("6'8"): // 80"
		minWingspan := MustLengthToInches("6'8")
		switch wingspanInches {
		case minWingspan:
			wingspanDeficit = 5
		case minWingspan + 1:
			wingspanDeficit = 4
		case minWingspan + 2:
			wingspanDeficit = 3
		case minWingspan + 3:
			wingspanDeficit = 1
		default: // >= minWingspan + 4 (7'0")
			wingspanDeficit = 0
		}
	case MustLengthToInches("6'9"): // 81"
		minWingspan := MustLengthToInches("6'9")
		switch wingspanInches {
		case minWingspan:
			wingspanDeficit = 7
		case minWingspan + 1:
			wingspanDeficit = 6
		case minWingspan + 2:
			wingspanDeficit = 5
		case minWingspan + 3:
			wingspanDeficit = 4
		case minWingspan + 4:
			wingspanDeficit = 3
		case minWingspan + 5:
			wingspanDeficit = 1
		default: // >= minWingspan + 6 (7'3")
			wingspanDeficit = 0
		}
	case MustLengthToInches("6'10"): // 82"
		minWingspan := MustLengthToInches("6'10")
		switch wingspanInches {
		case minWingspan:
			wingspanDeficit = 6
		case minWingspan + 1:
			wingspanDeficit = 5
		case minWingspan + 2:
			wingspanDeficit = 4
		case minWingspan + 3:
			wingspanDeficit = 3
		case minWingspan + 4:
			wingspanDeficit = 2
		case minWingspan + 5:
			wingspanDeficit = 1
		default: // >= minWingspan + 6 (7'4")
			wingspanDeficit = 0
		}
	case MustLengthToInches("6'11"): // 83"
		minWingspan := MustLengthToInches("6'11")
		switch wingspanInches {
		case minWingspan:
			wingspanDeficit = 6
		case minWingspan + 1:
			wingspanDeficit = 5
		case minWingspan + 2:
			wingspanDeficit = 4
		case minWingspan + 3:
			wingspanDeficit = 3
		case minWingspan + 4:
			wingspanDeficit = 2
		case minWingspan + 5:
			wingspanDeficit = 1
		default: // >= minWingspan + 6 (7'5")
			wingspanDeficit = 0
		}
	case MustLengthToInches("7'0"): // 84"
		minWingspan := MustLengthToInches("7'0")
		switch wingspanInches {
		case minWingspan:
			wingspanDeficit = 6
		case minWingspan + 1:
			wingspanDeficit = 5
		case minWingspan + 2:
			wingspanDeficit = 4
		case minWingspan + 3:
			wingspanDeficit = 3
		case minWingspan + 4:
			wingspanDeficit = 2
		case minWingspan + 5:
			wingspanDeficit = 1
		default: // >= minWingspan + 6 (7'6")
			wingspanDeficit = 0
		}
	case MustLengthToInches("7'1"): // 85"
		minWingspan := MustLengthToInches("7'1")
		switch wingspanInches {
		case minWingspan:
			wingspanDeficit = 5
		case minWingspan + 1:
			wingspanDeficit = 4
		case minWingspan + 2:
			wingspanDeficit = 3
		case minWingspan + 3:
			wingspanDeficit = 2
		case minWingspan + 4:
			wingspanDeficit = 1
		default: // >= minWingspan + 5 (7'6")
			wingspanDeficit = 0
		}
	case MustLengthToInches("7'2"): // 86"
		minWingspan := MustLengthToInches("7'2")
		switch wingspanInches {
		case minWingspan:
			wingspanDeficit = 5
		case minWingspan + 1:
			wingspanDeficit = 5 // Same as min
		case minWingspan + 2:
			wingspanDeficit = 4
		case minWingspan + 3:
			wingspanDeficit = 3
		case minWingspan + 4:
			wingspanDeficit = 2
		case minWingspan + 5:
			wingspanDeficit = 1
		default: // >= minWingspan + 6 (7'8")
			wingspanDeficit = 0
		}
	case MustLengthToInches("7'3"): // 87"
		minWingspan := MustLengthToInches("7'3")
		switch wingspanInches {
		case minWingspan:
			wingspanDeficit = 4
		case minWingspan + 1:
			wingspanDeficit = 3
		case minWingspan + 2:
			wingspanDeficit = 3 // Same as +1
		case minWingspan + 3:
			wingspanDeficit = 2
		case minWingspan + 4:
			wingspanDeficit = 1
		default: // >= minWingspan + 5 (7'8")
			wingspanDeficit = 0
		}
	case MustLengthToInches("7'4"): // 88"
		minWingspan := MustLengthToInches("7'4")
		switch wingspanInches {
		case minWingspan:
			wingspanDeficit = 4
		case minWingspan + 1:
			wingspanDeficit = 3
		case minWingspan + 2:
			wingspanDeficit = 2
		case minWingspan + 3:
			wingspanDeficit = 2 // Same as +2
		case minWingspan + 4:
			wingspanDeficit = 1
		default: // >= minWingspan + 5 (7'9")
			wingspanDeficit = 0
		}
	default:
		wingspanDeficit = 0
	}

	// Step 3: Calculate weight deficit (heavier = larger deficit)
	// TODO: Implement after weight testing is complete
	// Known data points for 7'4": 260 lbs → 64 cap, 290 lbs → 59 cap (rate: -1 per 6 lbs)
	// See docs/DATA-INCONSISTENCY-ISSUE.md for blocking issue with baseline weight
	weightDeficit = 0 // Placeholder until testing resolves data inconsistency

	// Step 4: Calculate final cap
	finalCap := 99 - heightDeficit - wingspanDeficit - weightDeficit

	// Clamp to valid range [0, 99]
	if finalCap < 0 {
		return 0
	}
	if finalCap > 99 {
		return 99
	}

	return finalCap
}

// StandingDunk calculates the Standing Dunk attribute cap for a Center.
// TODO: Implement based on testing data
func StandingDunk(heightInches, weightLbs, wingspanInches int) int {
	// Stub: returns 0 until pattern is discovered
	return 0
}

// PostControl calculates the Post Control attribute cap for a Center.
// TODO: Implement based on testing data
func PostControl(heightInches, weightLbs, wingspanInches int) int {
	// Stub: returns 0 until pattern is discovered
	return 0
}

// MidRangeShot calculates the Mid-Range Shot attribute cap for a Center.
// TODO: Implement based on testing data
func MidRangeShot(heightInches, weightLbs, wingspanInches int) int {
	// Stub: returns 0 until pattern is discovered
	return 0
}

// ThreePointShot calculates the Three-Point Shot attribute cap for a Center.
// TODO: Implement based on testing data
func ThreePointShot(heightInches, weightLbs, wingspanInches int) int {
	// Stub: returns 0 until pattern is discovered
	return 0
}

// FreeThrow calculates the Free Throw attribute cap for a Center.
// TODO: Implement based on testing data
func FreeThrow(heightInches, weightLbs, wingspanInches int) int {
	// Stub: returns 0 until pattern is discovered
	return 0
}

// BallHandle calculates the Ball Handle attribute cap for a Center.
// TODO: Implement based on testing data
func BallHandle(heightInches, weightLbs, wingspanInches int) int {
	// Stub: returns 0 until pattern is discovered
	return 0
}

// SpeedWithBall calculates the Speed With Ball attribute cap for a Center.
// TODO: Implement based on testing data
func SpeedWithBall(heightInches, weightLbs, wingspanInches int) int {
	// Stub: returns 0 until pattern is discovered
	return 0
}

// InteriorDefense calculates the Interior Defense attribute cap for a Center.
// TODO: Implement based on testing data
func InteriorDefense(heightInches, weightLbs, wingspanInches int) int {
	// Stub: returns 0 until pattern is discovered
	return 0
}

// PerimeterDefense calculates the Perimeter Defense attribute cap for a Center.
// TODO: Implement based on testing data
func PerimeterDefense(heightInches, weightLbs, wingspanInches int) int {
	// Stub: returns 0 until pattern is discovered
	return 0
}

// Steal calculates the Steal attribute cap for a Center.
// TODO: Implement based on testing data
func Steal(heightInches, weightLbs, wingspanInches int) int {
	// Stub: returns 0 until pattern is discovered
	return 0
}

// Block calculates the Block attribute cap for a Center.
// TODO: Implement based on testing data
func Block(heightInches, weightLbs, wingspanInches int) int {
	// Stub: returns 0 until pattern is discovered
	return 0
}

// OffensiveRebound calculates the Offensive Rebound attribute cap for a Center.
// TODO: Implement based on testing data
func OffensiveRebound(heightInches, weightLbs, wingspanInches int) int {
	// Stub: returns 0 until pattern is discovered
	return 0
}

// DefensiveRebound calculates the Defensive Rebound attribute cap for a Center.
// TODO: Implement based on testing data
func DefensiveRebound(heightInches, weightLbs, wingspanInches int) int {
	// Stub: returns 0 until pattern is discovered
	return 0
}

// Speed calculates the Speed attribute cap for a Center.
// TODO: Implement based on testing data
func Speed(heightInches, weightLbs, wingspanInches int) int {
	// Stub: returns 0 until pattern is discovered
	return 0
}

// Agility calculates the Agility attribute cap for a Center.
// TODO: Implement based on testing data
func Agility(heightInches, weightLbs, wingspanInches int) int {
	// Stub: returns 0 until pattern is discovered
	return 0
}

// Strength calculates the Strength attribute cap for a Center.
// TODO: Implement based on testing data
func Strength(heightInches, weightLbs, wingspanInches int) int {
	// Stub: returns 0 until pattern is discovered
	return 0
}

// Vertical calculates the Vertical attribute cap for a Center.
// TODO: Implement based on testing data
func Vertical(heightInches, weightLbs, wingspanInches int) int {
	// Stub: returns 0 until pattern is discovered
	return 0
}
