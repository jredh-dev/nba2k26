// SPDX-License-Identifier: AGPL-3.0
// Copyright (C) 2025 NBA 2K26 Attribute System

package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	"github.com/jredh-dev/nba2k26/pkg/attributes"
	"github.com/jredh-dev/nba2k26/pkg/scraper"
)

// attributeTest defines a test for one attribute function
type attributeTest struct {
	name     string
	attrName string
	attrFunc func(int, int, int) int
}

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

	fmt.Printf("Loaded %d builds from scraped data\n\n", len(caps))

	// All attribute functions to test
	tests := []attributeTest{
		{"CloseShot", "close_shot", attributes.CloseShot},
		{"PassAccuracy", "pass_accuracy", attributes.PassAccuracy},
		{"DrivingLayup", "driving_layup", attributes.DrivingLayup},
		{"DrivingDunk", "driving_dunk", attributes.DrivingDunk},
		{"StandingDunk", "standing_dunk", attributes.StandingDunk},
		{"PostControl", "post_control", attributes.PostControl},
		{"MidRangeShot", "mid_range_shot", attributes.MidRangeShot},
		{"ThreePointShot", "three_point_shot", attributes.ThreePointShot},
		{"FreeThrow", "free_throw", attributes.FreeThrow},
		{"BallHandle", "ball_handle", attributes.BallHandle},
		{"SpeedWithBall", "speed_with_ball", attributes.SpeedWithBall},
		{"InteriorDefense", "interior_defense", attributes.InteriorDefense},
		{"PerimeterDefense", "perimeter_defense", attributes.PerimeterDefense},
		{"Steal", "steal", attributes.Steal},
		{"Block", "block", attributes.Block},
		{"OffensiveRebound", "offensive_rebound", attributes.OffensiveRebound},
		{"DefensiveRebound", "defensive_rebound", attributes.DefensiveRebound},
		{"Speed", "speed", attributes.Speed},
		{"Agility", "agility", attributes.Agility},
		{"Strength", "strength", attributes.Strength},
		{"Vertical", "vertical", attributes.Vertical},
	}

	totalTests := 0
	totalPassed := 0
	totalFailed := 0
	totalStubs := 0

	for _, test := range tests {
		fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
		fmt.Printf("Testing: %s\n", test.name)
		fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")

		// Sample random builds for testing (10 samples)
		samples := []int{}
		for len(samples) < 10 && len(samples) < len(caps) {
			idx := rand.Intn(len(caps))
			samples = append(samples, idx)
		}

		passed := 0
		failed := 0
		stubbed := 0

		for _, idx := range samples {
			build := caps[idx]
			ourValue := test.attrFunc(build.Height, build.Weight, build.Wingspan)
			scrapedValue := getAttributeValue(&build, test.attrName)

			if ourValue == 0 {
				// Likely a stub
				stubbed++
			} else if ourValue == scrapedValue {
				passed++
			} else {
				failed++
				fmt.Printf("  ❌ H=%d\" WS=%d\" W=%dlbs: Our=%d, Scraped=%d (diff=%d)\n",
					build.Height, build.Wingspan, build.Weight, ourValue, scrapedValue, abs(ourValue-scrapedValue))
			}

			totalTests++
		}

		// Summary for this attribute
		if stubbed == len(samples) {
			fmt.Printf("  ⚠️  STUBBED - Function returns 0 (not implemented)\n")
			totalStubs++
		} else if failed == 0 {
			fmt.Printf("  ✅ PERFECT - All %d samples match scraped data!\n", passed)
			totalPassed++
		} else {
			fmt.Printf("  ⚠️  PARTIAL - %d passed, %d failed out of %d samples\n", passed, failed, len(samples))
			totalFailed++
		}
		fmt.Println()
	}

	// Overall summary
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Printf("OVERALL SUMMARY\n")
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Printf("✅ Fully Implemented & Accurate: %d\n", totalPassed)
	fmt.Printf("⚠️  Partially Implemented:       %d\n", totalFailed)
	fmt.Printf("⚠️  Stubbed (Not Implemented):  %d\n", totalStubs)
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Printf("Total Attributes: %d\n", len(tests))
	fmt.Printf("Total Test Samples: %d\n", totalTests)
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
}

func getAttributeValue(build *scraper.AttributeCaps, name string) int {
	switch name {
	case "close_shot":
		return build.CloseShot
	case "driving_layup":
		return build.DrivingLayup
	case "driving_dunk":
		return build.DrivingDunk
	case "standing_dunk":
		return build.StandingDunk
	case "post_control":
		return build.PostControl
	case "mid_range_shot":
		return build.MidRangeShot
	case "three_point_shot":
		return build.ThreePointShot
	case "free_throw":
		return build.FreeThrow
	case "pass_accuracy":
		return build.PassAccuracy
	case "ball_handle":
		return build.BallHandle
	case "speed_with_ball":
		return build.SpeedWithBall
	case "interior_defense":
		return build.InteriorDefense
	case "perimeter_defense":
		return build.PerimeterDefense
	case "steal":
		return build.Steal
	case "block":
		return build.Block
	case "offensive_rebound":
		return build.OffensiveRebound
	case "defensive_rebound":
		return build.DefensiveRebound
	case "speed":
		return build.Speed
	case "strength":
		return build.Strength
	case "vertical":
		return build.Vertical
	case "agility":
		return build.Agility
	default:
		return -1
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
