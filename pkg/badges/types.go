package badges

import "github.com/jredh-dev/nba2k26/pkg/scraper"

// BadgeTier represents the tier level of a badge
type BadgeTier int

const (
	// BadgeTierNone indicates the badge is not available
	BadgeTierNone BadgeTier = iota
	// BadgeTierBronze is the lowest badge tier
	BadgeTierBronze
	// BadgeTierSilver is the second badge tier
	BadgeTierSilver
	// BadgeTierGold is the third badge tier
	BadgeTierGold
	// BadgeTierHallOfFame is the fourth badge tier
	BadgeTierHallOfFame
	// BadgeTierLegendary is the highest badge tier
	BadgeTierLegendary
)

// String returns the string representation of a BadgeTier
func (b BadgeTier) String() string {
	return [...]string{"None", "Bronze", "Silver", "Gold", "Hall of Fame", "Legendary"}[b]
}

// BadgeFunc is a function that calculates the maximum tier for a badge based on attribute caps
type BadgeFunc func(attrs *scraper.AttributeCaps) BadgeTier

// BadgeCategory represents the category of a badge
type BadgeCategory int

const (
	// BadgeCategoryFinishing includes inside scoring badges
	BadgeCategoryFinishing BadgeCategory = iota
	// BadgeCategoryShooting includes outside scoring badges
	BadgeCategoryShooting
	// BadgeCategoryPlaymaking includes playmaking badges
	BadgeCategoryPlaymaking
	// BadgeCategoryDefense includes defensive badges
	BadgeCategoryDefense
	// BadgeCategoryRebounding includes rebounding badges
	BadgeCategoryRebounding
	// BadgeCategoryPhysicals includes physical/general offense badges
	BadgeCategoryPhysicals
	// BadgeCategoryAllAround includes all-around badges
	BadgeCategoryAllAround
)

// String returns the string representation of a BadgeCategory
func (c BadgeCategory) String() string {
	return [...]string{
		"Finishing",
		"Shooting",
		"Playmaking",
		"Defense",
		"Rebounding",
		"Physicals",
		"All-Around",
	}[c]
}

// Badge represents a badge with its metadata and calculation function
type Badge struct {
	// Name is the display name of the badge
	Name string
	// Category is the badge category (Finishing, Shooting, etc.)
	Category BadgeCategory
	// Description is a short description of what the badge does
	Description string
	// Calc is the function that calculates the maximum tier for this badge
	Calc BadgeFunc
}

// AttributeRequirement represents a single attribute requirement for a badge tier
type AttributeRequirement struct {
	// Attribute is the name of the attribute (e.g., "Driving Dunk", "Vertical")
	Attribute string
	// Bronze is the minimum attribute value for Bronze tier (0 if not available)
	Bronze int
	// Silver is the minimum attribute value for Silver tier (0 if not available)
	Silver int
	// Gold is the minimum attribute value for Gold tier (0 if not available)
	Gold int
	// HallOfFame is the minimum attribute value for Hall of Fame tier (0 if not available)
	HallOfFame int
	// Legendary is the minimum attribute value for Legendary tier (0 if not available)
	Legendary int
	// MinHeight is the minimum height requirement in inches (0 if no restriction)
	MinHeight int
	// MaxHeight is the maximum height requirement in inches (0 if no restriction)
	MaxHeight int
}

// BadgeRequirements represents all requirements for a badge
type BadgeRequirements struct {
	// Name is the badge name
	Name string
	// Category is the badge category from NBA2KLab data
	Category string
	// Type indicates if requirements are "Primary" (all must be met) or "Secondary" (any can be met)
	Type string
	// Requirements is the list of attribute requirements
	Requirements []AttributeRequirement
}
