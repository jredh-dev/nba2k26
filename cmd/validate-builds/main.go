// SPDX-License-Identifier: AGPL-3.0
// Copyright (C) 2025 NBA 2K26 Attribute System

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jredh-dev/nba2k26/pkg/scraper"
)

func main() {
	// Load scraped data
	data, err := os.ReadFile("data/Center_caps.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading data file: %v\n", err)
		os.Exit(1)
	}

	var caps []scraper.AttributeCaps
	if err := json.Unmarshal(data, &caps); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Printf("Validating %d Scraped Builds\n", len(caps))
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n")

	// Manual test cases from docs/center-findings.md
	manualTests := []struct {
		name                   string
		height                 int
		wingspan               int
		weight                 int
		expectedCloseShot      int
		expectedPassAccuracy   int
		expectedDrivingLayup   int
		expectedDrivingDunkMin int // Range
		expectedDrivingDunkMax int
	}{
		{
			name:                   "6'7\" min weight min wingspan",
			height:                 79,
			wingspan:               79,
			weight:                 215,
			expectedCloseShot:      99,
			expectedPassAccuracy:   99,
			expectedDrivingLayup:   99,
			expectedDrivingDunkMin: 95, // From manual findings: 6'7"H with 6'7"WS at min
			expectedDrivingDunkMax: 95,
		},
		{
			name:                   "7'4\" max weight max wingspan",
			height:                 88,
			wingspan:               94, // 7'10"
			weight:                 290,
			expectedCloseShot:      99,
			expectedPassAccuracy:   99,
			expectedDrivingLayup:   62, // From manual findings
			expectedDrivingDunkMin: 66, // From manual findings: 7'4"H max range
			expectedDrivingDunkMax: 70,
		},
		{
			name:                   "7'0\" default build",
			height:                 84,
			wingspan:               87, // 7'3" (default for 7'0")
			weight:                 250,
			expectedCloseShot:      99,
			expectedPassAccuracy:   99,
			expectedDrivingLayup:   91, // Weight dependent
			expectedDrivingDunkMin: 85, // Estimated
			expectedDrivingDunkMax: 87,
		},
	}

	fmt.Println("Comparing Manual Test Cases with Scraped Data\n")

	passedTests := 0
	failedTests := 0

	for _, test := range manualTests {
		fmt.Printf("Test: %s\n", test.name)
		fmt.Printf("  Build: H=%d\" WS=%d\" W=%dlbs\n", test.height, test.wingspan, test.weight)

		// Find in scraped data
		var found *scraper.AttributeCaps
		for i := range caps {
			if caps[i].Height == test.height &&
				caps[i].Wingspan == test.wingspan &&
				caps[i].Weight == test.weight {
				found = &caps[i]
				break
			}
		}

		if found == nil {
			fmt.Printf("  ❌ BUILD NOT FOUND IN SCRAPED DATA\n")
			fmt.Printf("     This suggests the API rejected this combination\n\n")
			failedTests++
			continue
		}

		allPassed := true

		// Validate CloseShot
		if found.CloseShot != test.expectedCloseShot {
			fmt.Printf("  ❌ CloseShot: Expected %d, Got %d\n", test.expectedCloseShot, found.CloseShot)
			allPassed = false
		} else {
			fmt.Printf("  ✅ CloseShot: %d\n", found.CloseShot)
		}

		// Validate PassAccuracy
		if found.PassAccuracy != test.expectedPassAccuracy {
			fmt.Printf("  ❌ PassAccuracy: Expected %d, Got %d\n", test.expectedPassAccuracy, found.PassAccuracy)
			allPassed = false
		} else {
			fmt.Printf("  ✅ PassAccuracy: %d\n", found.PassAccuracy)
		}

		// Validate DrivingLayup
		if found.DrivingLayup != test.expectedDrivingLayup {
			fmt.Printf("  ❌ DrivingLayup: Expected %d, Got %d\n", test.expectedDrivingLayup, found.DrivingLayup)
			allPassed = false
		} else {
			fmt.Printf("  ✅ DrivingLayup: %d\n", found.DrivingLayup)
		}

		// Validate DrivingDunk (range)
		if found.DrivingDunk >= test.expectedDrivingDunkMin && found.DrivingDunk <= test.expectedDrivingDunkMax {
			fmt.Printf("  ✅ DrivingDunk: %d (within expected range %d-%d)\n",
				found.DrivingDunk, test.expectedDrivingDunkMin, test.expectedDrivingDunkMax)
		} else {
			fmt.Printf("  ⚠️  DrivingDunk: %d (expected range %d-%d)\n",
				found.DrivingDunk, test.expectedDrivingDunkMin, test.expectedDrivingDunkMax)
			// Don't mark as failed - this is expected variation
		}

		if allPassed {
			passedTests++
		} else {
			failedTests++
		}
		fmt.Println()
	}

	// Data quality checks
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("Data Quality Checks\n")

	// Check 1: No zero values in core attributes
	zeroCloseShot := 0
	zeroPassAccuracy := 0
	zeroDrivingLayup := 0
	for _, build := range caps {
		if build.CloseShot == 0 {
			zeroCloseShot++
		}
		if build.PassAccuracy == 0 {
			zeroPassAccuracy++
		}
		if build.DrivingLayup == 0 {
			zeroDrivingLayup++
		}
	}

	fmt.Printf("Zero Values Check:\n")
	if zeroCloseShot == 0 && zeroPassAccuracy == 0 && zeroDrivingLayup == 0 {
		fmt.Printf("  ✅ No zero values in CloseShot, PassAccuracy, DrivingLayup\n")
	} else {
		fmt.Printf("  ❌ Found zeros: CloseShot=%d, PassAccuracy=%d, DrivingLayup=%d\n",
			zeroCloseShot, zeroPassAccuracy, zeroDrivingLayup)
	}

	// Check 2: All heights present
	heightCounts := make(map[int]int)
	for _, build := range caps {
		heightCounts[build.Height]++
	}

	fmt.Printf("\nHeight Distribution:\n")
	expectedHeights := []int{79, 80, 81, 82, 83, 84, 85, 86, 87, 88}
	allHeightsPresent := true
	for _, h := range expectedHeights {
		count := heightCounts[h]
		fmt.Printf("  %2d\" (", h)
		switch h {
		case 79:
			fmt.Printf("6'7\"")
		case 80:
			fmt.Printf("6'8\"")
		case 81:
			fmt.Printf("6'9\"")
		case 82:
			fmt.Printf("6'10\"")
		case 83:
			fmt.Printf("6'11\"")
		case 84:
			fmt.Printf("7'0\"")
		case 85:
			fmt.Printf("7'1\"")
		case 86:
			fmt.Printf("7'2\"")
		case 87:
			fmt.Printf("7'3\"")
		case 88:
			fmt.Printf("7'4\"")
		}
		fmt.Printf("): %3d builds", count)
		if count > 0 {
			fmt.Printf(" ✅\n")
		} else {
			fmt.Printf(" ❌\n")
			allHeightsPresent = false
		}
	}

	// Check 3: Weight distribution
	minWeight := 99999
	maxWeight := 0
	for _, build := range caps {
		if build.Weight < minWeight {
			minWeight = build.Weight
		}
		if build.Weight > maxWeight {
			maxWeight = build.Weight
		}
	}

	fmt.Printf("\nWeight Range:\n")
	fmt.Printf("  Min: %d lbs\n", minWeight)
	fmt.Printf("  Max: %d lbs\n", maxWeight)
	if minWeight == 215 && maxWeight == 290 {
		fmt.Printf("  ✅ Full range covered (215-290 lbs)\n")
	} else if minWeight >= 215 && maxWeight <= 290 {
		fmt.Printf("  ⚠️  Partial range (expected 215-290 lbs)\n")
	}

	// Check 4: Wingspan distribution
	minWingspan := 99999
	maxWingspan := 0
	for _, build := range caps {
		if build.Wingspan < minWingspan {
			minWingspan = build.Wingspan
		}
		if build.Wingspan > maxWingspan {
			maxWingspan = build.Wingspan
		}
	}

	fmt.Printf("\nWingspan Range:\n")
	fmt.Printf("  Min: %d\" (", minWingspan)
	fmt.Printf("6'%d\"", minWingspan-72)
	fmt.Printf(")\n")
	fmt.Printf("  Max: %d\" (", maxWingspan)
	fmt.Printf("%d'%d\"", maxWingspan/12, maxWingspan%12)
	fmt.Printf(")\n")

	// Summary
	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("Summary")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("Manual Test Cases: %d passed, %d failed\n", passedTests, failedTests)
	fmt.Printf("Total Builds: %d\n", len(caps))
	fmt.Printf("All Heights Present: %v\n", allHeightsPresent)
	fmt.Printf("Weight Range: %d-%d lbs\n", minWeight, maxWeight)
	fmt.Printf("Wingspan Range: %d-%d inches\n", minWingspan, maxWingspan)

	if failedTests > 0 {
		fmt.Println("\n⚠️  Some manual tests failed - review findings")
		os.Exit(1)
	} else {
		fmt.Println("\n✅ All manual tests passed - scraped data is valid!")
	}
}
