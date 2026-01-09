// SPDX-License-Identifier: AGPL-3.0
// Copyright (C) 2025 NBA 2K26 Build Analyzer

package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/jredh-dev/nba2k26/pkg/badges"
	"github.com/jredh-dev/nba2k26/pkg/scraper"
)

// BuildScore represents a build with its badge count and attribute quality
type BuildScore struct {
	Height   int
	Wingspan int
	Weight   int
	Badges   int
	Attrs    *scraper.AttributeCaps
	AvgCap   float64 // Average attribute cap
}

func main() {
	// Command-line flags
	minHeight := flag.Int("min-height", 78, "Minimum height in inches (6'6\" = 78)")
	maxHeight := flag.Int("max-height", 88, "Maximum height in inches (7'4\" = 88)")
	heightStep := flag.Int("height-step", 1, "Height increment in inches")

	minWingspan := flag.Int("min-wingspan", 78, "Minimum wingspan in inches")
	maxWingspan := flag.Int("max-wingspan", 92, "Maximum wingspan in inches")
	wingspanStep := flag.Int("wingspan-step", 2, "Wingspan increment in inches")

	minWeight := flag.Int("min-weight", 220, "Minimum weight in pounds")
	maxWeight := flag.Int("max-weight", 300, "Maximum weight in pounds")
	weightStep := flag.Int("weight-step", 20, "Weight increment in pounds")

	topN := flag.Int("top", 10, "Number of top builds to display")
	verbose := flag.Bool("verbose", false, "Show progress for each build")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: find-best-build [OPTIONS]\n\n")
		fmt.Fprintf(os.Stderr, "Find center builds with the most badges from NBA2KLab data.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  # Quick search\n")
		fmt.Fprintf(os.Stderr, "  find-best-build --top 5\n\n")
		fmt.Fprintf(os.Stderr, "  # Comprehensive search\n")
		fmt.Fprintf(os.Stderr, "  find-best-build --min-height 78 --max-height 88 --height-step 1 --wingspan-step 2 --weight-step 10\n\n")
	}

	flag.Parse()

	// Calculate total iterations
	heightCount := (*maxHeight-*minHeight) / *heightStep + 1
	wingspanCount := (*maxWingspan-*minWingspan) / *wingspanStep + 1
	weightCount := (*maxWeight-*minWeight) / *weightStep + 1
	total := heightCount * wingspanCount * weightCount

	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Printf("NBA 2K26 Center Build Analyzer\n")
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Printf("Height range:   %d\" - %d\" (step %d)\n", *minHeight, *maxHeight, *heightStep)
	fmt.Printf("Wingspan range: %d\" - %d\" (step %d)\n", *minWingspan, *maxWingspan, *wingspanStep)
	fmt.Printf("Weight range:   %d - %d lbs (step %d)\n", *minWeight, *maxWeight, *weightStep)
	fmt.Printf("Total builds:   %d\n", total)
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n")

	// Initialize scraper and badge calculator
	client := scraper.NewClient()
	calc, err := badges.NewCalculator()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing badge calculator: %v\n", err)
		os.Exit(1)
	}

	// Collect all builds
	var builds []BuildScore
	current := 0

	for height := *minHeight; height <= *maxHeight; height += *heightStep {
		for wingspan := *minWingspan; wingspan <= *maxWingspan; wingspan += *wingspanStep {
			for weight := *minWeight; weight <= *maxWeight; weight += *weightStep {
				current++

				// Fetch attributes from NBA2KLab
				attrs, err := client.GetAttributeCaps("Center", height, wingspan, weight)
				if err != nil {
					if *verbose {
						fmt.Printf("[%d/%d] ⚠️  H=%d\" WS=%d\" W=%dlbs - Error: %v\n",
							current, total, height, wingspan, weight, err)
					}
					continue
				}

				// Calculate badge count
				availableBadges := calc.GetAvailableBadges(attrs)

				badgeCount := len(availableBadges)

				// Calculate average attribute cap
				avgCap := calculateAverageAttributeCap(attrs)

				if *verbose {
					fmt.Printf("[%d/%d] ✓ H=%d\" WS=%d\" W=%dlbs - %d badges, Avg cap: %.1f\n",
						current, total, height, wingspan, weight, badgeCount, avgCap)
				} else if current%10 == 0 {
					fmt.Printf("Progress: %d/%d (%.1f%%)\r", current, total, float64(current)/float64(total)*100)
				}

				builds = append(builds, BuildScore{
					Height:   height,
					Wingspan: wingspan,
					Weight:   weight,
					Badges:   badgeCount,
					Attrs:    attrs,
					AvgCap:   avgCap,
				})
			}
		}
	}

	if !*verbose {
		fmt.Printf("\nCompleted: %d/%d builds analyzed\n\n", len(builds), total)
	}

	// Sort by badge count (descending)
	sort.Slice(builds, func(i, j int) bool {
		return builds[i].Badges > builds[j].Badges
	})

	// Display top N builds
	displayCount := *topN
	if displayCount > len(builds) {
		displayCount = len(builds)
	}

	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Printf("Top %d Builds by Badge Count\n", displayCount)
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n")

	for i := 0; i < displayCount; i++ {
		build := builds[i]
		fmt.Printf("%2d. %d badges (Avg cap: %.1f) - Height: %d\" (%s) | Wingspan: %d\" (%s) | Weight: %d lbs\n",
			i+1,
			build.Badges,
			build.AvgCap,
			build.Height,
			formatHeight(build.Height),
			build.Wingspan,
			formatHeight(build.Wingspan),
			build.Weight,
		)
	}

	// Show detailed stats for #1 build
	if len(builds) > 0 {
		best := builds[0]
		fmt.Printf("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
		fmt.Printf("Best Build: %s / %s / %d lbs (%d badges)\n",
			formatHeight(best.Height),
			formatHeight(best.Wingspan),
			best.Weight,
			best.Badges,
		)
		fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n")

		fmt.Println("⚠️  NOTE: These are ATTRIBUTE CAPS (max potential), not starting values!")
		fmt.Println("    Starting values are typically 60-80 points lower.")
		fmt.Println("    Badge availability is determined by attribute CAPS.\n")

		printAttributes(best.Attrs)
		fmt.Println()

		// Show badge breakdown by category
		fmt.Println("Badge Breakdown:")
		categories := []struct {
			name string
			cat  badges.BadgeCategory
		}{
			{"Finishing", badges.BadgeCategoryFinishing},
			{"Shooting", badges.BadgeCategoryShooting},
			{"Playmaking", badges.BadgeCategoryPlaymaking},
			{"Defense", badges.BadgeCategoryDefense},
			{"Rebounding", badges.BadgeCategoryRebounding},
			{"Physicals", badges.BadgeCategoryPhysicals},
			{"AllAround", badges.BadgeCategoryAllAround},
		}

		for _, cat := range categories {
			badgeMap := calc.GetBadgesByCategory(cat.cat, best.Attrs)
			count := 0
			for _, tier := range badgeMap {
				if tier != badges.BadgeTierNone {
					count++
				}
			}
			if count > 0 {
				fmt.Printf("  %-12s: %d badges\n", cat.name, count)
			}
		}
	}
}

// formatHeight converts inches to feet-inches format
func formatHeight(inches int) string {
	feet := inches / 12
	remainingInches := inches % 12
	return fmt.Sprintf("%d'%d\"", feet, remainingInches)
}

// calculateAverageAttributeCap returns the average of all attribute caps
func calculateAverageAttributeCap(attrs *scraper.AttributeCaps) float64 {
	total := attrs.CloseShot + attrs.DrivingLayup + attrs.DrivingDunk +
		attrs.StandingDunk + attrs.PostControl + attrs.MidRangeShot +
		attrs.ThreePointShot + attrs.FreeThrow + attrs.PassAccuracy +
		attrs.BallHandle + attrs.SpeedWithBall + attrs.InteriorDefense +
		attrs.PerimeterDefense + attrs.Steal + attrs.Block +
		attrs.OffensiveRebound + attrs.DefensiveRebound + attrs.Speed +
		attrs.Strength + attrs.Vertical + attrs.Agility

	return float64(total) / 21.0 // 21 attributes
}

// printAttributes displays attribute values in a readable format
func printAttributes(attrs *scraper.AttributeCaps) {
	fmt.Println("Attribute Caps (Maximum Potential):")
	fmt.Printf("  Finishing:\n")
	fmt.Printf("    Close Shot:      %d\n", attrs.CloseShot)
	fmt.Printf("    Driving Layup:   %d\n", attrs.DrivingLayup)
	fmt.Printf("    Driving Dunk:    %d\n", attrs.DrivingDunk)
	fmt.Printf("    Standing Dunk:   %d\n", attrs.StandingDunk)
	fmt.Printf("    Post Control:    %d\n", attrs.PostControl)

	fmt.Printf("  Shooting:\n")
	fmt.Printf("    Mid-Range:       %d\n", attrs.MidRangeShot)
	fmt.Printf("    Three-Point:     %d\n", attrs.ThreePointShot)
	fmt.Printf("    Free Throw:      %d\n", attrs.FreeThrow)

	fmt.Printf("  Playmaking:\n")
	fmt.Printf("    Pass Accuracy:   %d\n", attrs.PassAccuracy)
	fmt.Printf("    Ball Handle:     %d\n", attrs.BallHandle)
	fmt.Printf("    Speed w/ Ball:   %d\n", attrs.SpeedWithBall)

	fmt.Printf("  Defense:\n")
	fmt.Printf("    Interior:        %d\n", attrs.InteriorDefense)
	fmt.Printf("    Perimeter:       %d\n", attrs.PerimeterDefense)
	fmt.Printf("    Steal:           %d\n", attrs.Steal)
	fmt.Printf("    Block:           %d\n", attrs.Block)

	fmt.Printf("  Rebounding:\n")
	fmt.Printf("    Offensive:       %d\n", attrs.OffensiveRebound)
	fmt.Printf("    Defensive:       %d\n", attrs.DefensiveRebound)

	fmt.Printf("  Physicals:\n")
	fmt.Printf("    Speed:           %d\n", attrs.Speed)
	fmt.Printf("    Agility:         %d\n", attrs.Agility)
	fmt.Printf("    Strength:        %d\n", attrs.Strength)
	fmt.Printf("    Vertical:        %d\n", attrs.Vertical)
}
