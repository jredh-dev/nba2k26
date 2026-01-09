package badges_test

import (
	"testing"

	"github.com/jredh-dev/nba2k26/pkg/badges"
	"github.com/jredh-dev/nba2k26/pkg/scraper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestNewCalculator tests calculator creation
func TestNewCalculator(t *testing.T) {
	calc, err := badges.NewCalculator()
	require.NoError(t, err)
	require.NotNil(t, calc)

	// Should have loaded badge requirements
	allBadges := calc.ListAllBadges()
	assert.Greater(t, len(allBadges), 0, "Should have loaded badge requirements")
}

// TestBadgeTierString tests BadgeTier string representation
func TestBadgeTierString(t *testing.T) {
	tests := []struct {
		tier     badges.BadgeTier
		expected string
	}{
		{badges.BadgeTierNone, "None"},
		{badges.BadgeTierBronze, "Bronze"},
		{badges.BadgeTierSilver, "Silver"},
		{badges.BadgeTierGold, "Gold"},
		{badges.BadgeTierHallOfFame, "Hall of Fame"},
		{badges.BadgeTierLegendary, "Legendary"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.tier.String())
		})
	}
}

// TestPosterizer tests the Posterizer badge calculation
// Posterizer requirements: Driving Dunk + Vertical (Primary - both must be met)
// Bronze: DD 73+, Vert 65+
// Silver: DD 87+, Vert 75+
// Gold: DD 93+, Vert 80+
// HoF: DD 96+, Vert 85+
// Legend: DD 99+, Vert 90+
func TestPosterizer(t *testing.T) {
	calc, err := badges.NewCalculator()
	require.NoError(t, err)

	tests := []struct {
		name         string
		drivingDunk  int
		vertical     int
		expectedTier badges.BadgeTier
	}{
		{
			name:         "Below Bronze threshold",
			drivingDunk:  70,
			vertical:     60,
			expectedTier: badges.BadgeTierNone,
		},
		{
			name:         "Bronze tier (both at minimum)",
			drivingDunk:  73,
			vertical:     65,
			expectedTier: badges.BadgeTierBronze,
		},
		{
			name:         "Bronze tier (DD high, Vert low)",
			drivingDunk:  90,
			vertical:     65,
			expectedTier: badges.BadgeTierBronze, // Limited by Vertical
		},
		{
			name:         "Silver tier",
			drivingDunk:  87,
			vertical:     75,
			expectedTier: badges.BadgeTierSilver,
		},
		{
			name:         "Gold tier",
			drivingDunk:  93,
			vertical:     80,
			expectedTier: badges.BadgeTierGold,
		},
		{
			name:         "Hall of Fame tier",
			drivingDunk:  96,
			vertical:     85,
			expectedTier: badges.BadgeTierHallOfFame,
		},
		{
			name:         "Legendary tier",
			drivingDunk:  99,
			vertical:     90,
			expectedTier: badges.BadgeTierLegendary,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attrs := &scraper.AttributeCaps{
				Position:    "Center",
				Height:      84, // 7'0"
				Wingspan:    88,
				Weight:      260,
				DrivingDunk: tt.drivingDunk,
				Vertical:    tt.vertical,
			}

			tier, err := calc.GetBadgeTier("Posterizer", attrs)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedTier, tier, "Tier mismatch for DD=%d, Vert=%d", tt.drivingDunk, tt.vertical)
		})
	}
}

// TestRiseUp tests the Rise Up badge calculation
// Rise Up requirements: Standing Dunk + Vertical (Primary - both must be met)
// Height restriction: 6'6" (78") minimum
func TestRiseUp(t *testing.T) {
	calc, err := badges.NewCalculator()
	require.NoError(t, err)

	tests := []struct {
		name         string
		height       int
		standingDunk int
		vertical     int
		expectedTier badges.BadgeTier
	}{
		{
			name:         "Too short for badge",
			height:       75, // 6'3"
			standingDunk: 99,
			vertical:     99,
			expectedTier: badges.BadgeTierNone,
		},
		{
			name:         "Minimum height with Bronze stats",
			height:       78, // 6'6"
			standingDunk: 72,
			vertical:     60,
			expectedTier: badges.BadgeTierBronze,
		},
		{
			name:         "7-foot with Hall of Fame stats",
			height:       84, // 7'0"
			standingDunk: 95,
			vertical:     69,
			expectedTier: badges.BadgeTierHallOfFame,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attrs := &scraper.AttributeCaps{
				Position:     "Center",
				Height:       tt.height,
				Wingspan:     tt.height + 4,
				Weight:       260,
				StandingDunk: tt.standingDunk,
				Vertical:     tt.vertical,
			}

			tier, err := calc.GetBadgeTier("Rise Up", attrs)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedTier, tier)
		})
	}
}

// TestDeadeye tests the Deadeye badge calculation (Secondary badge)
// Deadeye requirements: Mid-Range Shot OR Three-Point Shot (Secondary - either works)
func TestDeadeye(t *testing.T) {
	calc, err := badges.NewCalculator()
	require.NoError(t, err)

	tests := []struct {
		name         string
		midRange     int
		threePoint   int
		expectedTier badges.BadgeTier
	}{
		{
			name:         "Both below threshold",
			midRange:     70,
			threePoint:   70,
			expectedTier: badges.BadgeTierNone,
		},
		{
			name:         "Mid-range at Gold, three-point low",
			midRange:     92,
			threePoint:   50,
			expectedTier: badges.BadgeTierGold, // Only mid-range matters
		},
		{
			name:         "Three-point at Silver, mid-range low",
			midRange:     50,
			threePoint:   85,
			expectedTier: badges.BadgeTierSilver, // Only three-point matters
		},
		{
			name:         "Both at different tiers (takes highest)",
			midRange:     73,                   // Bronze
			threePoint:   92,                   // Gold
			expectedTier: badges.BadgeTierGold, // Takes higher tier
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attrs := &scraper.AttributeCaps{
				Position:       "Center",
				Height:         84,
				Wingspan:       88,
				Weight:         260,
				MidRangeShot:   tt.midRange,
				ThreePointShot: tt.threePoint,
			}

			tier, err := calc.GetBadgeTier("Deadeye", attrs)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedTier, tier)
		})
	}
}

// TestGetAvailableBadges tests listing all available badges for a build
func TestGetAvailableBadges(t *testing.T) {
	calc, err := badges.NewCalculator()
	require.NoError(t, err)

	// High-attribute build (should have many badges)
	attrs := &scraper.AttributeCaps{
		Position:         "Center",
		Height:           84, // 7'0"
		Wingspan:         88,
		Weight:           260,
		CloseShot:        99,
		DrivingLayup:     90,
		DrivingDunk:      85,
		StandingDunk:     92,
		PostControl:      96,
		MidRangeShot:     80,
		ThreePointShot:   75,
		FreeThrow:        85,
		PassAccuracy:     99,
		BallHandle:       70,
		SpeedWithBall:    70,
		InteriorDefense:  95,
		PerimeterDefense: 80,
		Steal:            75,
		Block:            90,
		OffensiveRebound: 85,
		DefensiveRebound: 90,
		Speed:            75,
		Strength:         95,
		Vertical:         85,
		Agility:          75,
	}

	available := calc.GetAvailableBadges(attrs)
	assert.Greater(t, len(available), 10, "High-attribute build should have many badges")

	// Check some specific badges we expect
	_, hasPosterizer := available["Posterizer"]
	assert.True(t, hasPosterizer, "Should have Posterizer with DD 85+ and Vert 85+")

	_, hasRiseUp := available["Rise Up"]
	assert.True(t, hasRiseUp, "Should have Rise Up with SD 92+ and Vert 85+")
}

// TestGetBadgesByCategory tests filtering badges by category
func TestGetBadgesByCategory(t *testing.T) {
	calc, err := badges.NewCalculator()
	require.NoError(t, err)

	// Build with good finishing attributes
	attrs := &scraper.AttributeCaps{
		Position:     "Center",
		Height:       84,
		Wingspan:     88,
		Weight:       260,
		DrivingLayup: 90,
		DrivingDunk:  90,
		StandingDunk: 90,
		PostControl:  90,
		CloseShot:    90,
		Vertical:     85,
		Strength:     85,
	}

	finishing := calc.GetBadgesByCategory(badges.BadgeCategoryFinishing, attrs)
	assert.Greater(t, len(finishing), 0, "Should have finishing badges")

	// Should not have many shooting badges with low shooting stats
	attrs.MidRangeShot = 50
	attrs.ThreePointShot = 50
	shooting := calc.GetBadgesByCategory(badges.BadgeCategoryShooting, attrs)
	assert.Equal(t, 0, len(shooting), "Should have no shooting badges with low shooting stats")
}

// TestBadgeNotFound tests error handling for unknown badge
func TestBadgeNotFound(t *testing.T) {
	calc, err := badges.NewCalculator()
	require.NoError(t, err)

	attrs := &scraper.AttributeCaps{
		Position: "Center",
		Height:   84,
	}

	_, err = calc.GetBadgeTier("NonexistentBadge", attrs)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}
