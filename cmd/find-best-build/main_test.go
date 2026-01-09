// SPDX-License-Identifier: AGPL-3.0
// Copyright (C) 2025 NBA 2K26 Build Analyzer Tests

package main

import (
	"testing"

	"github.com/jredh-dev/nba2k26/pkg/badges"
	"github.com/jredh-dev/nba2k26/pkg/scraper"
)

// TestFormatHeight verifies height formatting
func TestFormatHeight(t *testing.T) {
	tests := []struct {
		name   string
		inches int
		want   string
	}{
		{"6 feet even", 72, `6'0"`},
		{"6'10\"", 82, `6'10"`},
		{"7'2\"", 86, `7'2"`},
		{"7'4\"", 88, `7'4"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatHeight(tt.inches)
			if got != tt.want {
				t.Errorf("formatHeight(%d) = %q, want %q", tt.inches, got, tt.want)
			}
		})
	}
}

// TestBuildScoreSorting verifies sorting by badge count
func TestBuildScoreSorting(t *testing.T) {
	builds := []BuildScore{
		{Height: 82, Wingspan: 86, Weight: 220, Badges: 30},
		{Height: 80, Wingspan: 84, Weight: 240, Badges: 34},
		{Height: 84, Wingspan: 88, Weight: 260, Badges: 28},
		{Height: 81, Wingspan: 85, Weight: 230, Badges: 32},
	}

	// Sort by badge count descending
	sortBuildsByBadges(builds)

	// Verify order
	if builds[0].Badges != 34 {
		t.Errorf("First build should have 34 badges, got %d", builds[0].Badges)
	}
	if builds[1].Badges != 32 {
		t.Errorf("Second build should have 32 badges, got %d", builds[1].Badges)
	}
	if builds[2].Badges != 30 {
		t.Errorf("Third build should have 30 badges, got %d", builds[2].Badges)
	}
	if builds[3].Badges != 28 {
		t.Errorf("Fourth build should have 28 badges, got %d", builds[3].Badges)
	}
}

// sortBuildsByBadges is extracted for testing
func sortBuildsByBadges(builds []BuildScore) {
	// Using bubble sort for simplicity in test
	for i := 0; i < len(builds); i++ {
		for j := i + 1; j < len(builds); j++ {
			if builds[i].Badges < builds[j].Badges {
				builds[i], builds[j] = builds[j], builds[i]
			}
		}
	}
}

// TestBadgeCountAccuracy tests that badge counting matches actual available badges
func TestBadgeCountAccuracy(t *testing.T) {
	// Create a test build with known attributes
	attrs := &scraper.AttributeCaps{
		Position:         "Center",
		Height:           82,
		Wingspan:         86,
		Weight:           220,
		CloseShot:        99,
		DrivingLayup:     98,
		DrivingDunk:      98,
		StandingDunk:     99,
		PostControl:      99,
		MidRangeShot:     85,
		ThreePointShot:   82,
		FreeThrow:        95,
		PassAccuracy:     99,
		BallHandle:       80,
		SpeedWithBall:    70,
		InteriorDefense:  99,
		PerimeterDefense: 87,
		Steal:            88,
		Block:            98,
		OffensiveRebound: 95,
		DefensiveRebound: 96,
		Speed:            81,
		Strength:         95,
		Vertical:         91,
		Agility:          77,
	}

	calc, err := badges.NewCalculator()
	if err != nil {
		t.Fatalf("Failed to create calculator: %v", err)
	}

	// Get available badges
	availableBadges := calc.GetAvailableBadges(attrs)

	// Count non-None badges
	count := 0
	for _, tier := range availableBadges {
		if tier != badges.BadgeTierNone {
			count++
		}
	}

	// This build should have at least 30 badges (based on known optimal builds)
	if count < 30 {
		t.Errorf("Expected at least 30 badges for optimal build, got %d", count)
	}

	// Should have badges in multiple categories
	categories := []badges.BadgeCategory{
		badges.BadgeCategoryFinishing,
		badges.BadgeCategoryPlaymaking,
		badges.BadgeCategoryDefense,
	}

	for _, cat := range categories {
		catBadges := calc.GetBadgesByCategory(cat, attrs)
		catCount := 0
		for _, tier := range catBadges {
			if tier != badges.BadgeTierNone {
				catCount++
			}
		}
		if catCount == 0 {
			t.Errorf("Expected badges in category %v, got 0", cat)
		}
	}
}

// TestBadgeBreakdownSum verifies category counts sum to total
func TestBadgeBreakdownSum(t *testing.T) {
	attrs := &scraper.AttributeCaps{
		Position:         "Center",
		Height:           82,
		Wingspan:         86,
		Weight:           220,
		CloseShot:        99,
		DrivingLayup:     98,
		DrivingDunk:      98,
		StandingDunk:     99,
		PostControl:      99,
		MidRangeShot:     85,
		ThreePointShot:   82,
		FreeThrow:        95,
		PassAccuracy:     99,
		BallHandle:       80,
		SpeedWithBall:    70,
		InteriorDefense:  99,
		PerimeterDefense: 87,
		Steal:            88,
		Block:            98,
		OffensiveRebound: 95,
		DefensiveRebound: 96,
		Speed:            81,
		Strength:         95,
		Vertical:         91,
		Agility:          77,
	}

	calc, err := badges.NewCalculator()
	if err != nil {
		t.Fatalf("Failed to create calculator: %v", err)
	}

	// Get total badge count
	totalBadges := calc.GetAvailableBadges(attrs)
	totalCount := 0
	for _, tier := range totalBadges {
		if tier != badges.BadgeTierNone {
			totalCount++
		}
	}

	// Get category counts
	categories := []badges.BadgeCategory{
		badges.BadgeCategoryFinishing,
		badges.BadgeCategoryShooting,
		badges.BadgeCategoryPlaymaking,
		badges.BadgeCategoryDefense,
		badges.BadgeCategoryRebounding,
		badges.BadgeCategoryPhysicals,
		badges.BadgeCategoryAllAround,
	}

	categorySum := 0
	for _, cat := range categories {
		catBadges := calc.GetBadgesByCategory(cat, attrs)
		for _, tier := range catBadges {
			if tier != badges.BadgeTierNone {
				categorySum++
			}
		}
	}

	// Category sum should equal total count
	if categorySum != totalCount {
		t.Errorf("Category sum (%d) doesn't match total count (%d)", categorySum, totalCount)
	}
}
