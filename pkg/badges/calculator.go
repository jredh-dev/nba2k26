package badges

import (
	"fmt"
	"strings"

	"github.com/jredh-dev/nba2k26/pkg/scraper"
)

// Calculator calculates badge tiers based on attribute caps
type Calculator struct {
	requirements map[string]*BadgeRequirements
}

// NewCalculator creates a new badge calculator
func NewCalculator() (*Calculator, error) {
	reqs, err := LoadBadgeRequirements()
	if err != nil {
		return nil, err
	}

	return &Calculator{
		requirements: reqs,
	}, nil
}

// GetBadgeTier calculates the maximum tier available for a specific badge
func (c *Calculator) GetBadgeTier(badgeName string, attrs *scraper.AttributeCaps) (BadgeTier, error) {
	// Convert badge name to ID format (e.g., "Posterizer" -> "Posterizer", "Ankle Assassin" -> "AnkleAssassin")
	badgeID := strings.ReplaceAll(badgeName, " ", "")
	badgeID = strings.ReplaceAll(badgeID, "-", "")

	reqs, exists := c.requirements[badgeID]
	if !exists {
		return BadgeTierNone, fmt.Errorf("badge %q not found", badgeName)
	}

	// Check height restrictions (if any requirement has height restrictions, they apply)
	for _, req := range reqs.Requirements {
		if req.MinHeight > 0 && attrs.Height < req.MinHeight {
			return BadgeTierNone, nil
		}
		if req.MaxHeight > 0 && attrs.Height > req.MaxHeight {
			return BadgeTierNone, nil
		}
	}

	// Calculate tier based on requirements
	// Primary badges: ALL requirements must be met
	// Secondary badges: ANY requirement can be met
	if reqs.Type == "Primary" {
		return c.calculatePrimaryBadgeTier(reqs, attrs), nil
	}

	return c.calculateSecondaryBadgeTier(reqs, attrs), nil
}

// calculatePrimaryBadgeTier calculates tier when ALL requirements must be met
func (c *Calculator) calculatePrimaryBadgeTier(reqs *BadgeRequirements, attrs *scraper.AttributeCaps) BadgeTier {
	// Get the minimum tier across all requirements (bottleneck)
	minTier := BadgeTierLegendary

	for _, req := range reqs.Requirements {
		attrValue := c.getAttributeValue(req.Attribute, attrs)
		tier := c.getTierForRequirement(req, attrValue)

		if tier < minTier {
			minTier = tier
		}
	}

	return minTier
}

// calculateSecondaryBadgeTier calculates tier when ANY requirement can be met
func (c *Calculator) calculateSecondaryBadgeTier(reqs *BadgeRequirements, attrs *scraper.AttributeCaps) BadgeTier {
	// Get the maximum tier across any requirement
	maxTier := BadgeTierNone

	for _, req := range reqs.Requirements {
		attrValue := c.getAttributeValue(req.Attribute, attrs)
		tier := c.getTierForRequirement(req, attrValue)

		if tier > maxTier {
			maxTier = tier
		}
	}

	return maxTier
}

// getTierForRequirement determines the tier based on a single attribute requirement
func (c *Calculator) getTierForRequirement(req AttributeRequirement, attrValue int) BadgeTier {
	// Check from highest to lowest tier
	if req.Legendary > 0 && attrValue >= req.Legendary {
		return BadgeTierLegendary
	}
	if req.HallOfFame > 0 && attrValue >= req.HallOfFame {
		return BadgeTierHallOfFame
	}
	if req.Gold > 0 && attrValue >= req.Gold {
		return BadgeTierGold
	}
	if req.Silver > 0 && attrValue >= req.Silver {
		return BadgeTierSilver
	}
	if req.Bronze > 0 && attrValue >= req.Bronze {
		return BadgeTierBronze
	}

	return BadgeTierNone
}

// getAttributeValue extracts the attribute value from AttributeCaps
func (c *Calculator) getAttributeValue(attributeName string, attrs *scraper.AttributeCaps) int {
	// Map attribute names from NBA2KLab to our struct fields
	switch attributeName {
	case "Close Shot":
		return attrs.CloseShot
	case "Driving Layup", "Layup":
		return attrs.DrivingLayup
	case "Driving Dunk":
		return attrs.DrivingDunk
	case "Standing Dunk":
		return attrs.StandingDunk
	case "Post Control":
		return attrs.PostControl
	case "Mid-Range Shot":
		return attrs.MidRangeShot
	case "Three-Point Shot":
		return attrs.ThreePointShot
	case "Free Throw":
		return attrs.FreeThrow
	case "Pass Accuracy":
		return attrs.PassAccuracy
	case "Ball Handle":
		return attrs.BallHandle
	case "Speed With Ball":
		return attrs.SpeedWithBall
	case "Interior Defense":
		return attrs.InteriorDefense
	case "Perimeter Defense":
		return attrs.PerimeterDefense
	case "Steal":
		return attrs.Steal
	case "Block":
		return attrs.Block
	case "Offensive Rebound":
		return attrs.OffensiveRebound
	case "Defensive Rebound":
		return attrs.DefensiveRebound
	case "Speed":
		return attrs.Speed
	case "Strength":
		return attrs.Strength
	case "Vertical":
		return attrs.Vertical
	case "Agility":
		return attrs.Agility
	default:
		return 0
	}
}

// GetAvailableBadges returns all badges available for a build (tier > None)
func (c *Calculator) GetAvailableBadges(attrs *scraper.AttributeCaps) map[string]BadgeTier {
	result := make(map[string]BadgeTier)

	for _, reqs := range c.requirements {
		tier, err := c.GetBadgeTier(reqs.Name, attrs)
		if err != nil {
			continue
		}

		if tier > BadgeTierNone {
			result[reqs.Name] = tier
		}
	}

	return result
}

// GetBadgesByCategory returns all badges in a specific category
func (c *Calculator) GetBadgesByCategory(category BadgeCategory, attrs *scraper.AttributeCaps) map[string]BadgeTier {
	result := make(map[string]BadgeTier)
	categoryStr := c.mapCategoryToString(category)

	for _, reqs := range c.requirements {
		if !c.matchesCategory(reqs.Category, categoryStr) {
			continue
		}

		tier, err := c.GetBadgeTier(reqs.Name, attrs)
		if err != nil {
			continue
		}

		if tier > BadgeTierNone {
			result[reqs.Name] = tier
		}
	}

	return result
}

// mapCategoryToString maps BadgeCategory enum to NBA2KLab category strings
func (c *Calculator) mapCategoryToString(category BadgeCategory) string {
	switch category {
	case BadgeCategoryFinishing:
		return "Inside Scoring"
	case BadgeCategoryShooting:
		return "Outside Scoring"
	case BadgeCategoryPlaymaking:
		return "Playmaking"
	case BadgeCategoryDefense:
		return "Defense"
	case BadgeCategoryRebounding:
		return "Rebounding"
	case BadgeCategoryPhysicals:
		return "General Offense"
	case BadgeCategoryAllAround:
		return "All Around"
	default:
		return ""
	}
}

// matchesCategory checks if a badge category matches the filter
func (c *Calculator) matchesCategory(badgeCat, filterCat string) bool {
	return badgeCat == filterCat
}

// ListAllBadges returns all badge names
func (c *Calculator) ListAllBadges() []string {
	badges := make([]string, 0, len(c.requirements))
	for _, reqs := range c.requirements {
		badges = append(badges, reqs.Name)
	}
	return badges
}
