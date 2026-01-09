// SPDX-License-Identifier: AGPL-3.0
// Copyright (C) 2025 NBA 2K26 Badge System

package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/jredh-dev/nba2k26/pkg/attributes"
	"github.com/jredh-dev/nba2k26/pkg/badges"
	"github.com/jredh-dev/nba2k26/pkg/scraper"
)

func main() {
	// Command-line flags
	position := flag.String("position", "Center", "Position (Center, PG, SG, SF, PF)")
	heightStr := flag.String("height", "", "Height in format 7-0 or 84 (inches)")
	wingspanStr := flag.String("wingspan", "", "Wingspan in format 7-3 or 87 (inches)")
	weight := flag.Int("weight", 0, "Weight in pounds")
	category := flag.String("category", "", "Filter by category (Finishing, Shooting, Playmaking, Defense, Rebounding, Physicals, AllAround)")
	badge := flag.String("badge", "", "Check specific badge only")
	minTier := flag.String("min-tier", "Bronze", "Minimum tier to display (Bronze, Silver, Gold, HoF, Legendary)")
	showAll := flag.Bool("all", false, "Show all badges including unavailable (None tier)")
	showAttrs := flag.Bool("show-attributes", false, "Show calculated attribute values")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: badge-checker [OPTIONS]\n\n")
		fmt.Fprintf(os.Stderr, "Check badge availability for NBA 2K26 builds.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  # Check all badges for a 7'0\" Center\n")
		fmt.Fprintf(os.Stderr, "  badge-checker --position Center --height 7-0 --wingspan 7-3 --weight 260\n\n")
		fmt.Fprintf(os.Stderr, "  # Check only finishing badges\n")
		fmt.Fprintf(os.Stderr, "  badge-checker --height 84 --wingspan 87 --weight 260 --category Finishing\n\n")
		fmt.Fprintf(os.Stderr, "  # Check specific badge\n")
		fmt.Fprintf(os.Stderr, "  badge-checker --height 7-0 --wingspan 7-3 --weight 260 --badge Posterizer\n\n")
		fmt.Fprintf(os.Stderr, "  # Show all badges including unavailable\n")
		fmt.Fprintf(os.Stderr, "  badge-checker --height 7-0 --wingspan 7-3 --weight 260 --all\n\n")
	}

	flag.Parse()

	// Validate required flags
	if *heightStr == "" || *wingspanStr == "" || *weight == 0 {
		fmt.Fprintf(os.Stderr, "Error: --height, --wingspan, and --weight are required\n\n")
		flag.Usage()
		os.Exit(1)
	}

	// Parse height and wingspan
	height, err := parseHeight(*heightStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing height: %v\n", err)
		os.Exit(1)
	}

	wingspan, err := parseHeight(*wingspanStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing wingspan: %v\n", err)
		os.Exit(1)
	}

	// Calculate attribute caps using attribute system
	attrs := calculateAttributeCaps(height, wingspan, *weight)

	// Print build summary
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Printf("Build: %s\n", *position)
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Printf("Height:   %d\" (%s)\n", height, formatHeight(height))
	fmt.Printf("Wingspan: %d\" (%s)\n", wingspan, formatHeight(wingspan))
	fmt.Printf("Weight:   %d lbs\n\n", *weight)

	// Show attributes if requested
	if *showAttrs {
		printAttributes(attrs)
		fmt.Println()
	}

	// Initialize badge calculator
	calc, err := badges.NewCalculator()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing badge calculator: %v\n", err)
		os.Exit(1)
	}

	// Parse minimum tier
	minTierValue := parseTier(*minTier)

	// Handle specific badge query
	if *badge != "" {
		tier, err := calc.GetBadgeTier(*badge, attrs)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		printBadgeDetails(*badge, tier, attrs, calc)
		return
	}

	// Handle category filter
	var badgeTiers map[string]badges.BadgeTier
	if *category != "" {
		cat, err := parseCategory(*category)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		badgeTiers = calc.GetBadgesByCategory(cat, attrs)
	} else {
		badgeTiers = calc.GetAvailableBadges(attrs)
	}

	// Filter by minimum tier
	if !*showAll {
		filtered := make(map[string]badges.BadgeTier)
		for name, tier := range badgeTiers {
			if tier >= minTierValue {
				filtered[name] = tier
			}
		}
		badgeTiers = filtered
	}

	// Group badges by category
	grouped := groupByCategory(badgeTiers, calc)

	// Print results
	if len(badgeTiers) == 0 {
		fmt.Printf("No badges available at %s tier or higher.\n", *minTier)
		return
	}

	fmt.Printf("Available Badges (%d):\n\n", len(badgeTiers))

	categories := []badges.BadgeCategory{
		badges.BadgeCategoryFinishing,
		badges.BadgeCategoryShooting,
		badges.BadgeCategoryPlaymaking,
		badges.BadgeCategoryDefense,
		badges.BadgeCategoryRebounding,
		badges.BadgeCategoryPhysicals,
		badges.BadgeCategoryAllAround,
	}

	for _, cat := range categories {
		badgesInCat := grouped[cat]
		if len(badgesInCat) == 0 {
			continue
		}

		fmt.Printf("%s (%d):\n", categoryName(cat), len(badgesInCat))

		// Sort badges by tier (highest first), then by name
		sort.Slice(badgesInCat, func(i, j int) bool {
			if badgesInCat[i].tier != badgesInCat[j].tier {
				return badgesInCat[i].tier > badgesInCat[j].tier
			}
			return badgesInCat[i].name < badgesInCat[j].name
		})

		for _, b := range badgesInCat {
			fmt.Printf("  %s %s (%s)\n", tierEmoji(b.tier), b.name, b.tier)
		}
		fmt.Println()
	}
}

// badgeInfo holds badge name and tier for sorting
type badgeInfo struct {
	name string
	tier badges.BadgeTier
}

// calculateAttributeCaps fetches real attribute caps from NBA2KLab API with fallback to local calculations
func calculateAttributeCaps(height, wingspan, weight int) *scraper.AttributeCaps {
	// Try to fetch real data from NBA2KLab first
	client := scraper.NewClient()
	attrs, err := client.GetAttributeCaps("Center", height, wingspan, weight)

	if err == nil && attrs != nil {
		return attrs
	}

	// Fallback to local attribute functions if API fails
	fmt.Fprintf(os.Stderr, "Warning: Failed to fetch from NBA2KLab API, using local calculations: %v\n", err)
	return &scraper.AttributeCaps{
		Height:           height,
		Wingspan:         wingspan,
		Weight:           weight,
		CloseShot:        attributes.CloseShot(height, weight, wingspan),
		DrivingLayup:     attributes.DrivingLayup(height, weight, wingspan),
		DrivingDunk:      attributes.DrivingDunk(height, weight, wingspan),
		StandingDunk:     attributes.StandingDunk(height, weight, wingspan),
		PostControl:      attributes.PostControl(height, weight, wingspan),
		MidRangeShot:     attributes.MidRangeShot(height, weight, wingspan),
		ThreePointShot:   attributes.ThreePointShot(height, weight, wingspan),
		FreeThrow:        attributes.FreeThrow(height, weight, wingspan),
		PassAccuracy:     attributes.PassAccuracy(height, weight, wingspan),
		BallHandle:       attributes.BallHandle(height, weight, wingspan),
		SpeedWithBall:    attributes.SpeedWithBall(height, weight, wingspan),
		InteriorDefense:  attributes.InteriorDefense(height, weight, wingspan),
		PerimeterDefense: attributes.PerimeterDefense(height, weight, wingspan),
		Steal:            attributes.Steal(height, weight, wingspan),
		Block:            attributes.Block(height, weight, wingspan),
		OffensiveRebound: attributes.OffensiveRebound(height, weight, wingspan),
		DefensiveRebound: attributes.DefensiveRebound(height, weight, wingspan),
		Speed:            attributes.Speed(height, weight, wingspan),
		Agility:          attributes.Agility(height, weight, wingspan),
		Strength:         attributes.Strength(height, weight, wingspan),
		Vertical:         attributes.Vertical(height, weight, wingspan),
	}
}

// printBadgeDetails prints detailed information about a specific badge
func printBadgeDetails(name string, tier badges.BadgeTier, attrs *scraper.AttributeCaps, calc *badges.Calculator) {
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Printf("%s\n", name)
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Printf("Tier: %s %s\n\n", tierEmoji(tier), tier)

	if tier == badges.BadgeTierNone {
		fmt.Printf("This badge is not available for this build.\n")
		fmt.Printf("Check badge requirements and adjust your build's height, weight, or wingspan.\n")
	} else {
		fmt.Printf("✅ This badge is available for this build.\n")
	}
}

// groupByCategory groups badges by their category
func groupByCategory(badgeTiers map[string]badges.BadgeTier, calc *badges.Calculator) map[badges.BadgeCategory][]badgeInfo {
	grouped := make(map[badges.BadgeCategory][]badgeInfo)

	// Get all badge requirements to access category information
	for name, tier := range badgeTiers {
		// Get badge tier to access requirement data (we ignore tier since we already have it)
		// This is a workaround - ideally Calculator would expose category lookup
		cat := inferCategoryFromName(name)
		grouped[cat] = append(grouped[cat], badgeInfo{name, tier})
	}

	return grouped
}

// inferCategoryFromName infers badge category from scraped data
// This is a temporary solution until Calculator exposes category metadata
func inferCategoryFromName(badgeName string) badges.BadgeCategory {
	// Inside Scoring badges
	insideScoring := []string{"Aerial Wizard", "Backdown Punisher", "Dream Shake", "Drop Stepper",
		"Fast Twitch", "Fearless Finisher", "Giant Slayer", "Masher", "Post Spin Technician",
		"Posterizer", "Pro Touch", "Putback Boss", "Rise Up", "Slithery Finisher",
		"Limitless Takeoff", "Acrobat", "Bully", "Tear Dropper", "Layup Mixmaster"}

	// Outside Scoring badges
	outsideScoring := []string{"Agent 3", "Blinders", "Catch & Shoot", "Claymore", "Corner Specialist",
		"Deadeye", "Guard Up", "Middy Magician", "Space Creator"}

	// Playmaking badges
	playmaking := []string{"Anchor", "Bail Out", "Break Starter", "Dimer", "Handles For Days",
		"Mismatch Expert", "Needle Threader", "Post Playmaker", "Quick First Step",
		"Special Delivery", "Unpluckable"}

	// Defense badges
	defense := []string{"Ankle Braces", "Challenger", "Clamps", "Glove", "Interceptor",
		"Lightning Reflexes", "Off-Ball Pest", "Pick Dodger", "Post Lockdown", "Work Horse",
		"Chase Down Artist", "Pogo Stick", "Intimidator"}

	// Rebounding badges
	rebounding := []string{"Boxout Beast", "Rebound Chaser"}

	// Physical/General Offense badges
	physicals := []string{"Brick Wall", "Bulldozer", "Physical Finisher"}

	// All-Around badges
	allAround := []string{"Versatility"}

	for _, b := range insideScoring {
		if b == badgeName {
			return badges.BadgeCategoryFinishing
		}
	}
	for _, b := range outsideScoring {
		if b == badgeName {
			return badges.BadgeCategoryShooting
		}
	}
	for _, b := range playmaking {
		if b == badgeName {
			return badges.BadgeCategoryPlaymaking
		}
	}
	for _, b := range defense {
		if b == badgeName {
			return badges.BadgeCategoryDefense
		}
	}
	for _, b := range rebounding {
		if b == badgeName {
			return badges.BadgeCategoryRebounding
		}
	}
	for _, b := range physicals {
		if b == badgeName {
			return badges.BadgeCategoryPhysicals
		}
	}
	for _, b := range allAround {
		if b == badgeName {
			return badges.BadgeCategoryAllAround
		}
	}

	// Default to finishing if unknown
	return badges.BadgeCategoryFinishing
}

// parseHeight parses height from "7-0" or "84" format
func parseHeight(s string) (int, error) {
	// Try feet-inches format first (e.g., "7-0")
	var feet, in int
	n, err := fmt.Sscanf(s, "%d-%d", &feet, &in)
	if err == nil && n == 2 {
		return feet*12 + in, nil
	}

	// Try simple inches format (e.g., "84")
	var inches int
	n, err = fmt.Sscanf(s, "%d", &inches)
	if err == nil && n == 1 {
		return inches, nil
	}

	return 0, fmt.Errorf("invalid height format %q (use 7-0 or 84)", s)
}

// formatHeight formats inches as feet-inches
func formatHeight(inches int) string {
	feet := inches / 12
	in := inches % 12
	return fmt.Sprintf("%d'%d\"", feet, in)
}

// parseTier converts string to BadgeTier
func parseTier(s string) badges.BadgeTier {
	switch strings.ToLower(s) {
	case "bronze":
		return badges.BadgeTierBronze
	case "silver":
		return badges.BadgeTierSilver
	case "gold":
		return badges.BadgeTierGold
	case "hof", "halloffame":
		return badges.BadgeTierHallOfFame
	case "legendary", "legend":
		return badges.BadgeTierLegendary
	default:
		return badges.BadgeTierBronze
	}
}

// parseCategory converts string to BadgeCategory
func parseCategory(s string) (badges.BadgeCategory, error) {
	switch strings.ToLower(s) {
	case "finishing", "inside":
		return badges.BadgeCategoryFinishing, nil
	case "shooting", "outside":
		return badges.BadgeCategoryShooting, nil
	case "playmaking", "passing":
		return badges.BadgeCategoryPlaymaking, nil
	case "defense", "defensive":
		return badges.BadgeCategoryDefense, nil
	case "rebounding", "rebounds":
		return badges.BadgeCategoryRebounding, nil
	case "physicals", "physical":
		return badges.BadgeCategoryPhysicals, nil
	case "allaround", "all":
		return badges.BadgeCategoryAllAround, nil
	default:
		return 0, fmt.Errorf("unknown category %q", s)
	}
}

// categoryName returns human-readable category name
func categoryName(cat badges.BadgeCategory) string {
	switch cat {
	case badges.BadgeCategoryFinishing:
		return "Finishing"
	case badges.BadgeCategoryShooting:
		return "Shooting"
	case badges.BadgeCategoryPlaymaking:
		return "Playmaking"
	case badges.BadgeCategoryDefense:
		return "Defense"
	case badges.BadgeCategoryRebounding:
		return "Rebounding"
	case badges.BadgeCategoryPhysicals:
		return "Physicals"
	case badges.BadgeCategoryAllAround:
		return "All-Around"
	default:
		return "Unknown"
	}
}

// tierEmoji returns an emoji for the tier
func tierEmoji(tier badges.BadgeTier) string {
	switch tier {
	case badges.BadgeTierLegendary:
		return "💎"
	case badges.BadgeTierHallOfFame:
		return "🥇"
	case badges.BadgeTierGold:
		return "🥈"
	case badges.BadgeTierSilver:
		return "🥉"
	case badges.BadgeTierBronze:
		return "🔶"
	default:
		return "❌"
	}
}

// printAttributes prints all calculated attribute values
func printAttributes(attrs *scraper.AttributeCaps) {
	fmt.Println("Calculated Attribute Caps:")
	fmt.Printf("  Finishing:\n")
	fmt.Printf("    Close Shot:        %2d\n", attrs.CloseShot)
	fmt.Printf("    Driving Layup:     %2d\n", attrs.DrivingLayup)
	fmt.Printf("    Driving Dunk:      %2d\n", attrs.DrivingDunk)
	fmt.Printf("    Standing Dunk:     %2d\n", attrs.StandingDunk)
	fmt.Printf("    Post Control:      %2d\n", attrs.PostControl)
	fmt.Printf("  Shooting:\n")
	fmt.Printf("    Mid-Range Shot:    %2d\n", attrs.MidRangeShot)
	fmt.Printf("    Three-Point Shot:  %2d\n", attrs.ThreePointShot)
	fmt.Printf("    Free Throw:        %2d\n", attrs.FreeThrow)
	fmt.Printf("  Playmaking:\n")
	fmt.Printf("    Pass Accuracy:     %2d\n", attrs.PassAccuracy)
	fmt.Printf("    Ball Handle:       %2d\n", attrs.BallHandle)
	fmt.Printf("    Speed With Ball:   %2d\n", attrs.SpeedWithBall)
	fmt.Printf("  Defense:\n")
	fmt.Printf("    Interior Defense:  %2d\n", attrs.InteriorDefense)
	fmt.Printf("    Perimeter Defense: %2d\n", attrs.PerimeterDefense)
	fmt.Printf("    Steal:             %2d\n", attrs.Steal)
	fmt.Printf("    Block:             %2d\n", attrs.Block)
	fmt.Printf("  Rebounding:\n")
	fmt.Printf("    Offensive Rebound: %2d\n", attrs.OffensiveRebound)
	fmt.Printf("    Defensive Rebound: %2d\n", attrs.DefensiveRebound)
	fmt.Printf("  Physicals:\n")
	fmt.Printf("    Speed:             %2d\n", attrs.Speed)
	fmt.Printf("    Agility:           %2d\n", attrs.Agility)
	fmt.Printf("    Strength:          %2d\n", attrs.Strength)
	fmt.Printf("    Vertical:          %2d\n", attrs.Vertical)
}
