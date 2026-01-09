package badges

import (
	"embed"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data/badge_requirements.json
var badgeDataFS embed.FS

// rawBadgeRequirement represents the JSON structure from NBA2KLab
type rawBadgeRequirement struct {
	Category  string `json:"Category"`
	Badge     string `json:"Badge"`
	Type      string `json:"Type"`
	Attribute string `json:"Attribute"`
	Bronze    int    `json:"Bronze"`
	Silver    int    `json:"Silver"`
	Gold      int    `json:"Gold"`
	HoF       any    `json:"HoF"`        // Can be int or empty string
	Legend    any    `json:"Legend"`     // Can be int or empty string
	MinHeight string `json:"Min_Height"` // Format: "6'3"
	MaxHeight string `json:"Max_Height"` // Format: "7'4"
	ID        string `json:"id"`
}

// parseHeight converts height string like "6'3" to inches (75)
func parseHeight(heightStr string) int {
	if heightStr == "" {
		return 0
	}

	parts := strings.Split(heightStr, "'")
	if len(parts) != 2 {
		return 0
	}

	feet, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0
	}

	inches, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0
	}

	return feet*12 + inches
}

// parseIntOrEmpty converts interface{} that can be int or empty string to int
func parseIntOrEmpty(val any) int {
	switch v := val.(type) {
	case float64:
		return int(v)
	case int:
		return v
	case string:
		if v == "" {
			return 0
		}
		result, _ := strconv.Atoi(v)
		return result
	default:
		return 0
	}
}

// LoadBadgeRequirements loads and parses badge requirements from embedded JSON
func LoadBadgeRequirements() (map[string]*BadgeRequirements, error) {
	data, err := badgeDataFS.ReadFile("data/badge_requirements.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read badge requirements: %w", err)
	}

	var rawReqs []rawBadgeRequirement
	if err := json.Unmarshal(data, &rawReqs); err != nil {
		return nil, fmt.Errorf("failed to parse badge requirements: %w", err)
	}

	// Group requirements by badge ID
	badgeMap := make(map[string]*BadgeRequirements)

	for _, raw := range rawReqs {
		badge, exists := badgeMap[raw.ID]
		if !exists {
			badge = &BadgeRequirements{
				Name:         raw.Badge,
				Category:     raw.Category,
				Type:         raw.Type,
				Requirements: []AttributeRequirement{},
			}
			badgeMap[raw.ID] = badge
		}

		// Add attribute requirement
		req := AttributeRequirement{
			Attribute:  raw.Attribute,
			Bronze:     raw.Bronze,
			Silver:     raw.Silver,
			Gold:       raw.Gold,
			HallOfFame: parseIntOrEmpty(raw.HoF),
			Legendary:  parseIntOrEmpty(raw.Legend),
			MinHeight:  parseHeight(raw.MinHeight),
			MaxHeight:  parseHeight(raw.MaxHeight),
		}

		badge.Requirements = append(badge.Requirements, req)
	}

	return badgeMap, nil
}
