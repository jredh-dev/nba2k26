# NBA 2K26 Badge System Implementation Plan

**Created:** 2026-01-08  
**Status:** Planning Phase

## Overview

Badges in NBA 2K26 are special abilities/animations that become available based on a player's attribute values. Each badge has multiple tiers, and the maximum tier available depends on meeting attribute thresholds.

## Badge Tier System

### Tier Levels
```go
type BadgeTier int

const (
    BadgeTierNone BadgeTier = iota
    BadgeTierBronze
    BadgeTierSilver
    BadgeTierGold
    BadgeTierHallOfFame
    BadgeTierLegendary
)

func (b BadgeTier) String() string {
    return [...]string{"None", "Bronze", "Silver", "Gold", "Hall of Fame", "Legendary"}[b]
}
```

## Architecture

### Badge Function Signature

Each badge is a **pure function** that takes attribute values and returns the highest available tier:

```go
// BadgeFunc calculates the maximum tier for a badge based on attribute caps
type BadgeFunc func(attrs *AttributeCaps) BadgeTier

// Example: Posterizer badge (requires DrivingDunk + Vertical)
func Posterizer(attrs *AttributeCaps) BadgeTier {
    if attrs.DrivingDunk >= 90 && attrs.Vertical >= 85 {
        return BadgeTierHallOfFame
    }
    if attrs.DrivingDunk >= 85 && attrs.Vertical >= 80 {
        return BadgeTierGold
    }
    if attrs.DrivingDunk >= 80 && attrs.Vertical >= 75 {
        return BadgeTierSilver
    }
    if attrs.DrivingDunk >= 75 && attrs.Vertical >= 70 {
        return BadgeTierBronze
    }
    return BadgeTierNone
}
```

### Registry System

All badges registered in a central map for easy lookup:

```go
// Badge represents a badge with its metadata and calculation function
type Badge struct {
    Name        string
    Category    BadgeCategory
    Description string
    Calc        BadgeFunc
}

type BadgeCategory int

const (
    BadgeCategoryFinishing BadgeCategory = iota
    BadgeCategoryShooting
    BadgeCategoryPlaymaking
    BadgeCategoryDefense
    BadgeCategoryRebounding
    BadgeCategoryPhysicals
)

// Global badge registry
var BadgeRegistry = map[string]*Badge{
    "Posterizer": {
        Name:        "Posterizer",
        Category:    BadgeCategoryFinishing,
        Description: "Increases the chances of successfully dunking on defenders",
        Calc:        Posterizer,
    },
    // ... more badges
}
```

### Usage

```go
// Get all available badges for a build
func GetAvailableBadges(attrs *AttributeCaps) map[string]BadgeTier {
    result := make(map[string]BadgeTier)
    
    for name, badge := range BadgeRegistry {
        tier := badge.Calc(attrs)
        if tier > BadgeTierNone {
            result[name] = tier
        }
    }
    
    return result
}

// Get specific badge tier
func GetBadgeTier(badgeName string, attrs *AttributeCaps) (BadgeTier, error) {
    badge, ok := BadgeRegistry[badgeName]
    if !ok {
        return BadgeTierNone, fmt.Errorf("badge %q not found", badgeName)
    }
    return badge.Calc(attrs), nil
}
```

## Data Collection Strategy

### Phase 1: Manual Testing (Immediate)

Since NBA2KLab API doesn't provide badge data, we need **in-game testing**:

1. **Test Known Builds**: Use scraped builds with known attribute combinations
2. **Check Badge Availability**: In character creator, see which badges light up
3. **Record Maximum Tiers**: Note the highest tier available for each badge
4. **Document Requirements**: Record which attributes are required

**Testing Template:**
```
Build: Center 7'0" / 255 lbs / 85" wingspan
Attributes:
  - DrivingDunk: 82
  - Vertical: 88
  - Strength: 95

Finishing Badges:
  ✓ Posterizer: Gold (req: DrivingDunk 80+, Vertical 85+)
  ✓ Slithery Finisher: Silver (req: DrivingLayup 75+)
  ✗ Limitless Takeoff: None (req: DrivingDunk 90+, Vertical 90+)
```

### Phase 2: Pattern Recognition (After 20-30 Tests)

Look for patterns:
- Single attribute requirements (e.g., "Post Spin Technician" → PostControl only)
- Multi-attribute requirements (e.g., "Posterizer" → DrivingDunk + Vertical)
- Position-specific badges (e.g., "Brick Wall" → Centers/PFs only)
- Height/weight/wingspan requirements (e.g., tall centers get "Intimidator" easier)

### Phase 3: Validation (After Implementation)

Create test matrix:
```go
func TestBadgeAccuracy(t *testing.T) {
    tests := []struct {
        build    *AttributeCaps
        badge    string
        expected BadgeTier
    }{
        {
            build:    &AttributeCaps{DrivingDunk: 90, Vertical: 85, ...},
            badge:    "Posterizer",
            expected: BadgeTierHallOfFame,
        },
        // ... more test cases from manual testing
    }
    
    for _, tt := range tests {
        actual, err := GetBadgeTier(tt.badge, tt.build)
        require.NoError(t, err)
        assert.Equal(t, tt.expected, actual, 
            "Badge %s with build %+v", tt.badge, tt.build)
    }
}
```

## Badge Categories (NBA 2K26)

### Finishing Badges (Centers)
- **Posterizer** - Dunk on defenders (DrivingDunk + Vertical)
- **Slithery Finisher** - Avoid contact on layups (DrivingLayup + Agility)
- **Limitless Takeoff** - Dunk from farther out (DrivingDunk + Vertical)
- **Backdown Punisher** - Post moves (PostControl + Strength)
- **Drop Stepper** - Post hop shots (PostControl + Vertical)
- **Fast Twitch** - Faster layup/dunk animations (DrivingLayup/Dunk)
- **Fearless Finisher** - Contact finishing (DrivingLayup/Dunk + Strength)
- **Giant Slayer** - Finish vs tall defenders (DrivingLayup + Height inverse)
- **Masher** - Stronger standing dunks (StandingDunk + Strength)
- **Post Spin Technician** - Post spins (PostControl)
- **Pro Touch** - Better timing windows (CloseShot + any finishing)
- **Putback Boss** - Offensive rebound putbacks (OffensiveRebound + StandingDunk)
- **Rise Up** - Standing dunk success (StandingDunk + Vertical)

### Shooting Badges (Centers)
- **Agent 3** - Corner threes (ThreePointShot)
- **Catch & Shoot** - Spot-up shooting (MidRange/ThreePoint)
- **Claymore** - Set shot timing (MidRange/ThreePoint)
- **Corner Specialist** - Corner shooting (ThreePointShot)
- **Deadeye** - Contested shots (MidRange/ThreePoint)
- **Middy Magician** - Mid-range shots (MidRangeShot)
- **Space Creator** - Off-dribble shots (MidRange/ThreePoint + BallHandle)

### Playmaking Badges (Centers)
- **Anchor** - Team defense boost (DefensiveAttributes)
- **Break Starter** - Outlet passes (PassAccuracy)
- **Bail Out** - Difficult passes (PassAccuracy + BallHandle)
- **Dimer** - Assist bonus (PassAccuracy)
- **Handles for Days** - Stamina on dribbles (BallHandle + Stamina)
- **Needle Threader** - Pass through traffic (PassAccuracy)
- **Post Playmaker** - Passes from post (PostControl + PassAccuracy)
- **Quick First Step** - Faster drives (Speed + Acceleration)
- **Special Delivery** - Alley-oop passes (PassAccuracy + Height)
- **Unpluckable** - Ball security (BallHandle)

### Defense/Rebounding Badges (Centers)
- **Anchor** - Defensive anchor (InteriorDefense + Block + Height)
- **Boxout Beast** - Boxing out (DefensiveRebound + Strength)
- **Brick Wall** - Screen effectiveness (Strength + Height)
- **Chase Down Artist** - Chase blocks (Speed + Block)
- **Challenger** - Contest shots (PerimeterDefense + Wingspan)
- **Clamps** - Perimeter defense (PerimeterDefense + Speed)
- **Interceptor** - Steal passes (Steal + PassPerception)
- **Intimidator** - Opponent shot penalty (InteriorDefense + Height)
- **Lightning Reflexes** - Faster contests (PerimeterDefense)
- **Off-Ball Pest** - Off-ball defense (PerimeterDefense + Speed)
- **Pogo Stick** - Multiple jump attempts (Vertical + Block)
- **Post Lockdown** - Post defense (InteriorDefense + Strength)
- **Rebound Chaser** - Rebound pursuit (DefensiveRebound + Speed)
- **Work Horse** - Defensive stamina (InteriorDefense + Stamina)

### Physical Badges (Centers)
- **Brick Wall** - Screen effectiveness (Strength + Weight)
- **Bulldozer** - Physical screens (Strength)
- **Immovable Enforcer** - Resist contact (Strength)
- **Physical Finisher** - Contact finishing (Strength + any finishing)

## Implementation Phases

### Phase 1: Core Infrastructure ✅ (Next)
**Files:** `pkg/badges/types.go`, `pkg/badges/registry.go`

- [ ] Define `BadgeTier` enum
- [ ] Define `BadgeFunc` signature
- [ ] Define `Badge` struct with metadata
- [ ] Create `BadgeRegistry` map
- [ ] Implement `GetAvailableBadges()` function
- [ ] Implement `GetBadgeTier()` function
- [ ] Unit tests for infrastructure

**Deliverable:** Badge system foundation with no badges implemented yet

### Phase 2: High-Priority Badges (5-10 badges)
**Focus:** Most common/important Center badges

Start with badges that have **single attribute requirements** (easiest to test):

1. **Post Spin Technician** - PostControl only
2. **Putback Boss** - OffensiveRebound + StandingDunk
3. **Rise Up** - StandingDunk + Vertical
4. **Pogo Stick** - Vertical + Block
5. **Boxout Beast** - DefensiveRebound + Strength
6. **Brick Wall** - Strength + Height (or Weight?)
7. **Anchor** - InteriorDefense + Block + Height
8. **Posterizer** - DrivingDunk + Vertical
9. **Intimidator** - InteriorDefense + Height
10. **Post Lockdown** - InteriorDefense + Strength

**Testing Process:**
1. Pick 3-5 scraped builds with varying attributes
2. Test each build in-game
3. Record available badges and max tiers
4. Implement badge functions based on findings
5. Add test cases for each badge

### Phase 3: Medium Priority (10-15 badges)
**Focus:** Common shooting, playmaking, and defense badges

- Shooting badges (if any apply to centers)
- Playmaking badges (passing/post playmaking)
- Remaining defense badges
- Finishing badges (layups, dunks)

### Phase 4: Low Priority (Remaining badges)
**Focus:** Rare/situational badges

- Perimeter-focused badges (less relevant for centers)
- Advanced playmaking badges
- Niche finishing badges

### Phase 5: Other Positions
**Scope:** Repeat for PG, SG, SF, PF

Each position has different badge priorities and requirements

## Testing Infrastructure

### Test Data Structure

Store manual test results in JSON:

```json
{
  "test_date": "2026-01-08",
  "position": "Center",
  "builds": [
    {
      "height": 84,
      "wingspan": 88,
      "weight": 260,
      "attributes": {
        "driving_dunk": 82,
        "vertical": 88,
        "strength": 95,
        "standing_dunk": 92
      },
      "badges": {
        "Posterizer": "Gold",
        "Rise Up": "Hall of Fame",
        "Brick Wall": "Silver"
      }
    }
  ]
}
```

**File:** `data/badge_tests_center.json` (git-tracked)

### Validation Tool

Create `cmd/badge-validator/main.go`:

```go
func main() {
    // Load test data
    tests := loadBadgeTests("data/badge_tests_center.json")
    
    // Validate each test case
    for _, test := range tests {
        attrs := &AttributeCaps{
            Height:       test.Height,
            Wingspan:     test.Wingspan,
            Weight:       test.Weight,
            DrivingDunk:  test.Attributes.DrivingDunk,
            // ... all attributes
        }
        
        for badgeName, expectedTier := range test.Badges {
            actualTier, err := GetBadgeTier(badgeName, attrs)
            if err != nil {
                log.Printf("❌ Badge %q not implemented", badgeName)
                continue
            }
            
            if actualTier.String() != expectedTier {
                log.Printf("⚠️  MISMATCH: %q - Expected: %s, Got: %s",
                    badgeName, expectedTier, actualTier)
            } else {
                log.Printf("✅ %q - %s", badgeName, actualTier)
            }
        }
    }
}
```

## CLI Tool Enhancement

Add badge queries to main calculator:

```bash
# Get all badges for a build
go run cmd/calculator/main.go \
  --position Center \
  --height 7-0 \
  --wingspan 7-3 \
  --weight 260 \
  --show-badges

# Output:
# Attributes: [... attribute caps ...]
#
# Available Badges (23):
#
# Finishing (6):
#   🥇 Posterizer (Hall of Fame)
#   🥇 Rise Up (Hall of Fame)
#   🥈 Slithery Finisher (Silver)
#   🥉 Fast Twitch (Bronze)
#   🥉 Putback Boss (Bronze)
#   🥉 Masher (Bronze)
#
# Defense (8):
#   🥇 Anchor (Hall of Fame)
#   🥇 Intimidator (Hall of Fame)
#   🥇 Pogo Stick (Hall of Fame)
#   🥈 Boxout Beast (Silver)
#   🥈 Post Lockdown (Silver)
#   🥈 Rebound Chaser (Silver)
#   🥉 Chase Down Artist (Bronze)
#   🥉 Interceptor (Bronze)
#
# [... more categories ...]

# Filter by category
go run cmd/calculator/main.go \
  --position Center --height 7-0 --wingspan 7-3 --weight 260 \
  --show-badges --category Finishing

# Show specific badge
go run cmd/calculator/main.go \
  --position Center --height 7-0 --wingspan 7-3 --weight 260 \
  --check-badge Posterizer
  
# Output:
# Posterizer: Hall of Fame ✅
# Requirements:
#   - Driving Dunk: 82 (requires 80+) ✅
#   - Vertical: 88 (requires 85+) ✅
```

## Documentation

### Badge Reference Doc

Create `docs/BADGE-REFERENCE.md`:

```markdown
# NBA 2K26 Badge Requirements (Centers)

## Finishing Badges

### Posterizer
**Description:** Increases success rate of dunking on defenders

**Tiers:**
- Bronze: DrivingDunk 75+, Vertical 70+
- Silver: DrivingDunk 80+, Vertical 75+
- Gold: DrivingDunk 85+, Vertical 80+
- Hall of Fame: DrivingDunk 90+, Vertical 85+

**Testing Notes:**
- Tested on 7'0" Center (260 lbs, 88" wingspan)
- DrivingDunk 82, Vertical 88 → Gold tier confirmed
- DrivingDunk 90, Vertical 87 → HOF confirmed

[... more badges ...]
```

## Success Criteria

### Phase 1 Complete When:
- [ ] Badge types and registry implemented
- [ ] Can query badges for any build
- [ ] Unit tests pass
- [ ] Documentation complete

### Phase 2 Complete When:
- [ ] 10+ high-priority badges implemented
- [ ] Manual testing data for 5+ builds
- [ ] Validation tool confirms accuracy
- [ ] Badge requirements documented

### Phase 3+ Complete When:
- [ ] 30+ Center badges implemented
- [ ] Validation tool shows 95%+ accuracy
- [ ] CLI tool supports badge queries
- [ ] Ready to expand to other positions

## Open Questions

1. **Position-specific badges**: Do badge requirements change by position?
2. **Height/weight/wingspan**: Do physical measurements affect thresholds directly?
3. **Legendary tier**: Is there a Legendary tier in 2K26? (Some years have it)
4. **Badge interactions**: Do some badges require other badges?
5. **Cap system**: Can you equip all available badges, or is there a limit?

## Next Steps

1. **Spawn agent** for Phase 1 infrastructure:
   ```bash
   agentctl agent spawn "implement badge system infrastructure" --repo jredh-dev/nba2k26
   ```

2. **Create badge package structure:**
   ```
   pkg/badges/
     types.go      # BadgeTier, BadgeFunc, Badge structs
     registry.go   # BadgeRegistry map, GetBadgeTier/GetAvailableBadges
     finishing.go  # Finishing badge functions
     shooting.go   # Shooting badge functions
     playmaking.go # Playmaking badge functions
     defense.go    # Defense/rebounding badge functions
     physical.go   # Physical badge functions
   ```

3. **Manual testing session**: Test 3-5 builds in-game to gather initial badge data

4. **Implement first 5 badges** with known single-attribute requirements

5. **Create validation infrastructure** to track accuracy

---

**Contributors:** AI Agent System  
**License:** AGPL-3.0
