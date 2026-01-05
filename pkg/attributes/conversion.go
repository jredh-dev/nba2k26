// SPDX-License-Identifier: AGPL-3.0
// Copyright (C) 2025 NBA 2K26 Attribute System

package attributes

import "fmt"

// LengthToInches converts a length string like "6'7"" to total inches (79)
// Works for both height and wingspan measurements
func LengthToInches(length string) (int, error) {
	var feet, inches int
	_, err := fmt.Sscanf(length, "%d'%d\"", &feet, &inches)
	if err != nil {
		return 0, fmt.Errorf("invalid length format: %s", length)
	}
	return feet*12 + inches, nil
}

// InchesToLength converts total inches (79) to length string "6'7""
// Works for both height and wingspan measurements
func InchesToLength(totalInches int) string {
	feet := totalInches / 12
	inches := totalInches % 12
	return fmt.Sprintf("%d'%d\"", feet, inches)
}

// WeightToInt converts a weight string like "215" to int
func WeightToInt(weight string) (int, error) {
	var w int
	_, err := fmt.Sscanf(weight, "%d", &w)
	if err != nil {
		return 0, fmt.Errorf("invalid weight format: %s", weight)
	}
	return w, nil
}
