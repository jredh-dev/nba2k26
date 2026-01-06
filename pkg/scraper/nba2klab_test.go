// SPDX-License-Identifier: AGPL-3.0
// Copyright (C) 2025 NBA 2K26 Attribute System

package scraper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAttributeCaps(t *testing.T) {
	client := NewClient()

	tests := []struct {
		name     string
		position string
		height   int
		wingspan int
		weight   int
		want     AttributeCaps
	}{
		{
			name:     "6'7\" Center default build",
			position: "Center",
			height:   79,
			wingspan: 82,
			weight:   243,
			want: AttributeCaps{
				Position:     "Center",
				Height:       79,
				Wingspan:     82,
				Weight:       243,
				CloseShot:    99,
				DrivingLayup: 99,
				PassAccuracy: 99,
			},
		},
		{
			name:     "7'3\" Center default build",
			position: "Center",
			height:   87,
			wingspan: 91,
			weight:   250,
			want: AttributeCaps{
				Position:         "Center",
				Height:           87,
				Wingspan:         91,
				Weight:           250,
				CloseShot:        99,
				DrivingLayup:     75,
				DrivingDunk:      73,
				StandingDunk:     99,
				PostControl:      99,
				PassAccuracy:     99,
				Block:            99,
				OffensiveRebound: 99,
				DefensiveRebound: 99,
			},
		},
		{
			name:     "7'4\" min wingspan min weight (data inconsistency test)",
			position: "Center",
			height:   88,
			wingspan: 88,
			weight:   270,
			want: AttributeCaps{
				Position:    "Center",
				Height:      88,
				Wingspan:    88,
				Weight:      270,
				DrivingDunk: 64, // Confirms weight data was correct, not wingspan data
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetAttributeCaps(tt.position, tt.height, tt.wingspan, tt.weight)
			require.NoError(t, err)
			require.NotNil(t, got)

			// Check all specified fields match
			assert.Equal(t, tt.want.Position, got.Position)
			assert.Equal(t, tt.want.Height, got.Height)
			assert.Equal(t, tt.want.Wingspan, got.Wingspan)
			assert.Equal(t, tt.want.Weight, got.Weight)

			if tt.want.CloseShot != 0 {
				assert.Equal(t, tt.want.CloseShot, got.CloseShot, "CloseShot mismatch")
			}
			if tt.want.DrivingLayup != 0 {
				assert.Equal(t, tt.want.DrivingLayup, got.DrivingLayup, "DrivingLayup mismatch")
			}
			if tt.want.DrivingDunk != 0 {
				assert.Equal(t, tt.want.DrivingDunk, got.DrivingDunk, "DrivingDunk mismatch")
			}
			if tt.want.StandingDunk != 0 {
				assert.Equal(t, tt.want.StandingDunk, got.StandingDunk, "StandingDunk mismatch")
			}
			if tt.want.PostControl != 0 {
				assert.Equal(t, tt.want.PostControl, got.PostControl, "PostControl mismatch")
			}
			if tt.want.PassAccuracy != 0 {
				assert.Equal(t, tt.want.PassAccuracy, got.PassAccuracy, "PassAccuracy mismatch")
			}
			if tt.want.Block != 0 {
				assert.Equal(t, tt.want.Block, got.Block, "Block mismatch")
			}
			if tt.want.OffensiveRebound != 0 {
				assert.Equal(t, tt.want.OffensiveRebound, got.OffensiveRebound, "OffensiveRebound mismatch")
			}
			if tt.want.DefensiveRebound != 0 {
				assert.Equal(t, tt.want.DefensiveRebound, got.DefensiveRebound, "DefensiveRebound mismatch")
			}
		})
	}
}

func TestGetAttributeCaps_InvalidBuild(t *testing.T) {
	client := NewClient()

	// Test with invalid height (below minimum for Centers)
	got, err := client.GetAttributeCaps("Center", 70, 75, 200)
	assert.Error(t, err)
	assert.Nil(t, got)
	assert.Contains(t, err.Error(), "no results returned")
}

func TestScrapeRange_SmallSample(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping API scrape test in short mode")
	}

	client := NewClient()

	// Scrape a small sample: 6'7" Center, wingspan 6'7"-6'9", weight 215-225 (step 5)
	results, err := client.ScrapeRange(
		"Center",
		[2]int{79, 79},      // Height: 6'7" only
		[2]int{79, 81},      // Wingspan: 6'7" to 6'9"
		[3]int{215, 225, 5}, // Weight: 215, 220, 225
	)

	require.NoError(t, err)
	require.NotEmpty(t, results)

	// Should have 3 wingspans × 3 weights = 9 results
	assert.Equal(t, 9, len(results))

	// Verify all results are for 6'7" Centers
	for _, r := range results {
		assert.Equal(t, "Center", r.Position)
		assert.Equal(t, 79, r.Height)
		assert.GreaterOrEqual(t, r.Wingspan, 79)
		assert.LessOrEqual(t, r.Wingspan, 81)
		assert.GreaterOrEqual(t, r.Weight, 215)
		assert.LessOrEqual(t, r.Weight, 225)
	}
}
