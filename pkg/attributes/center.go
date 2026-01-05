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
	case MustLengthToInches("7'2"): // 86" (7'2")
		// At 7'2", weight affects the cap (71-84 range)
		// Pattern: ~5 lbs per point (wider intervals than 7'3"/7'4")
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
	// TODO: Add intermediate heights as you discover them
	// case MustLengthToInches("6'8"):
	//     return ??
	// case MustLengthToInches("6'9"):
	//     return ??
	default:
		return 0 // Not yet tested
	}
}

// DrivingDunk calculates the Driving Dunk attribute cap for a Center.
// TODO: Implement based on testing data
func DrivingDunk(heightInches, weightLbs, wingspanInches int) int {
	// Stub: returns 0 until pattern is discovered
	return 0
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
