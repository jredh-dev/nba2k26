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
// - Pattern: Height is the primary factor, weight affects at maximum height
// - Wingspan doesn't appear to affect this attribute
func DrivingLayup(heightInches, weightLbs, wingspanInches int) int {
	// Height-based caps (discovered through testing)
	switch heightInches {
	case MustLengthToInches(CENTER_MIN_HEIGHT): // 79" (6'7")
		return 99
	case MustLengthToInches("6'8"): // 80" (6'8")
		return 99
	case MustLengthToInches("6'9"): // 81" (6'9")
		// Weight doesn't affect cap at this height (always 98)
		return 98
	case MustLengthToInches("6'10"): // 82" (6'10")
		// Weight doesn't affect cap at this height (always 96)
		return 96
	case MustLengthToInches("6'11"): // 83" (6'11")
		// Weight affects cap at this height (92-94 range)
		switch {
		case weightLbs <= 250:
			return 94
		case weightLbs <= 288:
			return 93
		default: // >= 289 lbs (max is 290)
			return 92
		}
	case MustLengthToInches("7'0"): // 84" (7'0")
		// Weight affects cap at this height (89-93 range)
		switch {
		case weightLbs <= 225:
			return 93
		case weightLbs <= 244:
			return 92
		case weightLbs <= 262:
			return 91
		case weightLbs <= 280:
			return 90
		default: // >= 281 lbs (max is 290)
			return 89
		}
	case MustLengthToInches("7'1"): // 85" (7'1")
		// Weight affects cap at this height (77-86 range)
		switch {
		case weightLbs <= 226:
			return 86
		case weightLbs <= 234:
			return 85
		case weightLbs <= 242:
			return 84
		case weightLbs <= 249:
			return 83
		case weightLbs <= 257:
			return 82
		case weightLbs <= 264:
			return 81
		case weightLbs <= 272:
			return 80
		case weightLbs <= 280:
			return 79
		case weightLbs <= 287:
			return 78
		default: // >= 288 lbs (max is 290)
			return 77
		}
	case MustLengthToInches("7'2"):
		switch {
		case weightLbs <= 223:
			return 84
		case weightLbs <= 228:
			return 83
		case weightLbs <= 233:
			return 82
		case weightLbs <= 239:
			return 81
		case weightLbs <= 244:
			return 80
		case weightLbs <= 249:
			return 79
		case weightLbs <= 254:
			return 78
		case weightLbs <= 259:
			return 77
		case weightLbs <= 264:
			return 76
		case weightLbs <= 269:
			return 75
		case weightLbs <= 275:
			return 74
		case weightLbs <= 280:
			return 73
		case weightLbs <= 285:
			return 72
		default: // >= 286 lbs (max is 290)
			return 71
		}
	case MustLengthToInches("7'3"): // 87" (7'3")
		// At 7'3", weight affects the cap (64-80 range)
		// Min weight (230 lbs) and 231 lbs both return 80
		switch {
		case weightLbs <= 231:
			return 80
		case weightLbs <= 234:
			return 79
		case weightLbs <= 238:
			return 78
		case weightLbs <= 242:
			return 77
		case weightLbs <= 246:
			return 76
		case weightLbs <= 250:
			return 75
		case weightLbs <= 254:
			return 74
		case weightLbs <= 258:
			return 73
		case weightLbs <= 262:
			return 72
		case weightLbs <= 266:
			return 71
		case weightLbs <= 270:
			return 70
		case weightLbs <= 274:
			return 69
		case weightLbs <= 277:
			return 68
		case weightLbs <= 281:
			return 67
		case weightLbs <= 285:
			return 66
		case weightLbs <= 289:
			return 65
		case weightLbs == 290:
			return 64
		default:
			return 0 // Invalid weight
		}
	case MustLengthToInches(CENTER_MAX_HEIGHT):
		switch {
		case weightLbs <= 231:
			return 77
		case weightLbs <= 235:
			return 76
		case weightLbs <= 239:
			return 75
		case weightLbs <= 243:
			return 74
		case weightLbs <= 247:
			return 73
		case weightLbs <= 251:
			return 72
		case weightLbs <= 256:
			return 71
		case weightLbs <= 260:
			return 70
		case weightLbs <= 264:
			return 69
		case weightLbs <= 268:
			return 68
		case weightLbs <= 272:
			return 67
		case weightLbs <= 276:
			return 66
		case weightLbs <= 280:
			return 65
		case weightLbs <= 285:
			return 64
		case weightLbs <= 289:
			return 63
		case weightLbs == 290:
			return 62
		default:
			return 0 // Invalid weight
		}
	default:
		return 0 // Not yet tested
	}
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
			return 86
		case MustLengthToInches("7'0"):
			return 87
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

// DrivingDunk2 calculates wingspan + weight penalties from max cap for Driving Dunk.
// Returns the total penalty from 99 based on height, wingspan, and weight.
//
// Current implementation: Only wingspan penalties at baseline weight (270 lbs).
// Weight penalties will be added after testing weight variations.
//
// Formula (future): 99 - DrivingDunk2(height, weight, wingspan) = Final Cap
func DrivingDunk2(heightInches, weightLbs, wingspanInches int) int {
	// For each height, calculate penalty from 99 based on wingspan
	// All data currently at baseline weight (270 lbs)
	// Weight modifier will be added later

	switch heightInches {
	case MustLengthToInches("6'7"): // 79"
		// Max cap at this height: 99 (with 7'1"+ wingspan)
		switch wingspanInches {
		case MustLengthToInches("6'7"):
			return 4 // Cap 95
		case MustLengthToInches("6'8"):
			return 2 // Cap 97
		case MustLengthToInches("6'9"):
			return 1 // Cap 98
		default: // 7'1" and above
			return 0 // Cap 99
		}
	case MustLengthToInches("6'8"): // 80"
		// Max cap at this height: 99 (with 7'0"+ wingspan)
		switch wingspanInches {
		case MustLengthToInches("6'8"):
			return 5 // Cap 94
		case MustLengthToInches("6'9"):
			return 4 // Cap 95
		case MustLengthToInches("6'10"):
			return 3 // Cap 96
		case MustLengthToInches("6'11"):
			return 1 // Cap 98
		default: // 7'0" and above
			return 0 // Cap 99
		}
	case MustLengthToInches("6'9"): // 81"
		// Max cap at this height: 99 (with 7'3" wingspan)
		switch wingspanInches {
		case MustLengthToInches("6'9"):
			return 7 // Cap 92
		case MustLengthToInches("6'10"):
			return 6 // Cap 93
		case MustLengthToInches("6'11"):
			return 5 // Cap 94
		case MustLengthToInches("7'0"):
			return 4 // Cap 95
		case MustLengthToInches("7'1"):
			return 3 // Cap 96
		case MustLengthToInches("7'2"):
			return 1 // Cap 98
		default: // 7'3" and above
			return 0 // Cap 99
		}
	case MustLengthToInches("6'10"): // 82"
		// Max cap at this height: 96 (with 7'4" wingspan)
		switch wingspanInches {
		case MustLengthToInches("6'10"):
			return 6 // Cap 90
		case MustLengthToInches("6'11"):
			return 5 // Cap 91
		case MustLengthToInches("7'0"):
			return 4 // Cap 92
		case MustLengthToInches("7'1"):
			return 3 // Cap 93
		case MustLengthToInches("7'2"):
			return 2 // Cap 94
		case MustLengthToInches("7'3"):
			return 1 // Cap 95
		default: // 7'4" and above
			return 0 // Cap 96
		}
	case MustLengthToInches("6'11"): // 83"
		// Max cap at this height: 92 (with 7'5" wingspan)
		switch wingspanInches {
		case MustLengthToInches("6'11"):
			return 6 // Cap 86
		case MustLengthToInches("7'0"):
			return 5 // Cap 87
		case MustLengthToInches("7'1"):
			return 4 // Cap 88
		case MustLengthToInches("7'2"):
			return 3 // Cap 89
		case MustLengthToInches("7'3"):
			return 2 // Cap 90
		case MustLengthToInches("7'4"):
			return 1 // Cap 91
		default: // 7'5" and above
			return 0 // Cap 92
		}
	case MustLengthToInches("7'0"): // 84"
		// Max cap at this height: 89 (with 7'6" wingspan)
		switch wingspanInches {
		case MustLengthToInches("7'0"):
			return 6 // Cap 83
		case MustLengthToInches("7'1"):
			return 5 // Cap 84
		case MustLengthToInches("7'2"):
			return 4 // Cap 85
		case MustLengthToInches("7'3"):
			return 3 // Cap 86
		case MustLengthToInches("7'4"):
			return 2 // Cap 87
		case MustLengthToInches("7'5"):
			return 1 // Cap 88
		default: // 7'6" and above
			return 0 // Cap 89
		}
	case MustLengthToInches("7'1"): // 85"
		// Max cap at this height: 82 (with 7'6"+ wingspan)
		switch wingspanInches {
		case MustLengthToInches("7'1"):
			return 5 // Cap 77
		case MustLengthToInches("7'2"):
			return 4 // Cap 78
		case MustLengthToInches("7'3"):
			return 3 // Cap 79
		case MustLengthToInches("7'4"):
			return 2 // Cap 80
		case MustLengthToInches("7'5"):
			return 1 // Cap 81
		default: // 7'6" and above
			return 0 // Cap 82
		}
	case MustLengthToInches("7'2"): // 86"
		// Max cap at this height: 77 (with 7'8" wingspan)
		switch wingspanInches {
		case MustLengthToInches("7'2"):
			return 5 // Cap 72
		case MustLengthToInches("7'3"):
			return 5 // Cap 72 (same as 7'2")
		case MustLengthToInches("7'4"):
			return 4 // Cap 73
		case MustLengthToInches("7'5"):
			return 3 // Cap 74
		case MustLengthToInches("7'6"):
			return 2 // Cap 75
		case MustLengthToInches("7'7"):
			return 1 // Cap 76
		default: // 7'8" and above
			return 0 // Cap 77
		}
	case MustLengthToInches("7'3"): // 87"
		// Max cap at this height: 72 (with 7'8"+ wingspan)
		switch wingspanInches {
		case MustLengthToInches("7'3"):
			return 4 // Cap 68
		case MustLengthToInches("7'4"):
			return 3 // Cap 69
		case MustLengthToInches("7'5"):
			return 3 // Cap 69 (same as 7'4")
		case MustLengthToInches("7'6"):
			return 2 // Cap 70
		case MustLengthToInches("7'7"):
			return 1 // Cap 71
		default: // 7'8" and above
			return 0 // Cap 72
		}
	case MustLengthToInches("7'4"): // 88"
		// Max cap at this height: 70 (with 7'9"+ wingspan)
		switch wingspanInches {
		case MustLengthToInches("7'4"):
			return 4 // Cap 66
		case MustLengthToInches("7'5"):
			return 3 // Cap 67
		case MustLengthToInches("7'6"):
			return 2 // Cap 68
		case MustLengthToInches("7'7"):
			return 2 // Cap 68 (same as 7'6")
		case MustLengthToInches("7'8"):
			return 1 // Cap 69
		default: // 7'9" and above
			return 0 // Cap 70
		}
	default:
		return 99 // Unknown height
	}
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
