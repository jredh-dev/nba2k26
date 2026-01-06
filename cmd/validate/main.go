// SPDX-License-Identifier: AGPL-3.0
// Copyright (C) 2025 NBA 2K26 Attribute System

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jredh-dev/nba2k26/pkg/attributes"
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

	fmt.Printf("Loaded %d builds from scraped data\n\n", len(caps))

	// Create lookup map
	buildMap := make(map[string]*scraper.AttributeCaps)
	for i := range caps {
		key := fmt.Sprintf("%d-%d-%d", caps[i].Height, caps[i].Wingspan, caps[i].Weight)
		buildMap[key] = &caps[i]
	}

	// Test cases from our manual testing
	tests := []struct {
		name   string
		height int
		ws     int
		weight int
		checks map[string]int // attribute -> expected value
	}{
		{
			name:   "6'7\" default build",
			height: 79,
			ws:     82,
			weight: 245, // Closest to default 243 (step 5)
			checks: map[string]int{
				"close_shot":    99,
				"driving_layup": 99,
				"pass_accuracy": 99,
			},
		},
		{
			name:   "7'4\" min wingspan (data inconsistency test)",
			height: 88,
			ws:     88,
			weight: 270,
			checks: map[string]int{
				"driving_dunk": 64, // Confirmed: not 66!
			},
		},
		{
			name:   "6'7\" min weight min wingspan",
			height: 79,
			ws:     79,
			weight: 215,
			checks: map[string]int{
				"close_shot":    99,
				"driving_layup": 99,
				"pass_accuracy": 99,
			},
		},
		{
			name:   "7'3\" default build",
			height: 87,
			ws:     91,
			weight: 260,
			checks: map[string]int{
				"close_shot":        99,
				"driving_layup":     75,
				"pass_accuracy":     99,
				"standing_dunk":     99,
				"block":             99,
				"offensive_rebound": 99,
				"defensive_rebound": 99,
			},
		},
	}

	passed := 0
	failed := 0

	for _, tt := range tests {
		fmt.Printf("Testing: %s (H=%d\" WS=%d\" W=%dlbs)\n", tt.name, tt.height, tt.ws, tt.weight)

		key := fmt.Sprintf("%d-%d-%d", tt.height, tt.ws, tt.weight)
		build, exists := buildMap[key]

		if !exists {
			fmt.Printf("  ❌ Build not found in scraped data\n")
			failed++
			continue
		}

		allPassed := true
		for attr, expected := range tt.checks {
			actual := getAttributeValue(build, attr)
			if actual == expected {
				fmt.Printf("  ✅ %-20s = %d\n", attr, actual)
			} else {
				fmt.Printf("  ❌ %-20s = %d (expected %d)\n", attr, actual, expected)
				allPassed = false
			}
		}

		if allPassed {
			passed++
		} else {
			failed++
		}
		fmt.Println()
	}

	// Test our functions against scraped data
	fmt.Println("=== Testing Our Functions vs Scraped Data ===")
	fmt.Println()

	funcTests := []struct {
		name      string
		height    int
		ws        int
		weight    int
		attrName  string
		attrFunc  func(int, int, int) int
		tolerance int // Allow some difference due to rounding
	}{
		{
			name:     "CloseShot - 6'7\" default",
			height:   79,
			ws:       82,
			weight:   245,
			attrName: "close_shot",
			attrFunc: attributes.CloseShot,
		},
		{
			name:     "DrivingLayup - 6'7\" min weight",
			height:   79,
			ws:       79,
			weight:   215,
			attrName: "driving_layup",
			attrFunc: attributes.DrivingLayup,
		},
		{
			name:     "DrivingLayup - 7'4\" default",
			height:   88,
			ws:       91,
			weight:   260,
			attrName: "driving_layup",
			attrFunc: attributes.DrivingLayup,
		},
		{
			name:     "PassAccuracy - any",
			height:   84,
			ws:       87,
			weight:   250,
			attrName: "pass_accuracy",
			attrFunc: attributes.PassAccuracy,
		},
	}

	for _, tt := range funcTests {
		fmt.Printf("Testing: %s\n", tt.name)

		// Get our function's result
		ourValue := tt.attrFunc(tt.height, tt.ws, tt.weight)

		// Lookup in scraped data
		key := fmt.Sprintf("%d-%d-%d", tt.height, tt.ws, tt.weight)

		build, exists := buildMap[key]
		if !exists {
			fmt.Printf("  ⚠️  Build not in scraped data (H=%d WS=%d W=%d)\n", tt.height, tt.ws, tt.weight)
			continue
		}

		scrapedValue := getAttributeValue(build, tt.attrName)
		diff := abs(ourValue - scrapedValue)

		if diff <= tt.tolerance {
			fmt.Printf("  ✅ Our func: %d, Scraped: %d (diff: %d)\n", ourValue, scrapedValue, diff)
			passed++
		} else {
			fmt.Printf("  ❌ Our func: %d, Scraped: %d (diff: %d, tolerance: %d)\n",
				ourValue, scrapedValue, diff, tt.tolerance)
			failed++
		}
		fmt.Println()
	}

	// Summary
	fmt.Println("==========================================")
	fmt.Printf("Results: %d passed, %d failed\n", passed, failed)
	if failed > 0 {
		os.Exit(1)
	}
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
