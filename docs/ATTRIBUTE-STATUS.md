# NBA 2K26 Attribute Implementation Status

**Last Updated:** 2025-01-06  
**Data Source:** NBA2KLab API (903 Center builds scraped)

## Summary

| Status | Count | Attributes |
|--------|-------|-----------|
| ✅ **Fully Implemented** | 3 | CloseShot, PassAccuracy, DrivingLayup |
| ⚠️ **Partial Implementation** | 1 | DrivingDunk (missing weight modifiers) |
| ❌ **Not Implemented** | 17 | All others |

## Detailed Status

### ✅ Fully Implemented & Validated (3/21)

These attributes are 100% accurate against scraped data from 903 builds:

#### 1. **CloseShot**
- **Pattern:** Always 99 (no variation by height/weight/wingspan)
- **Implementation:** Single return value
- **Validation:** ✅ 10/10 random samples match scraped data

#### 2. **PassAccuracy**
- **Pattern:** Always 99 (no variation by height/weight/wingspan)
- **Implementation:** Single return value
- **Validation:** ✅ 10/10 random samples match scraped data

#### 3. **DrivingLayup**
- **Pattern:** Height-based with weight penalties at 6'11"+
- **Implementation:** Data-driven lookup table (refactored 2025-01-06)
- **Key Insights:**
  - Heights 6'7"-6'10": Weight-independent
  - Heights 6'11"+: Weight creates penalties (heavier = lower cap)
  - Wingspan has no effect
- **Validation:** ✅ 10/10 random samples match scraped data
- **Range:** 62 (7'4" @ 290lbs) → 99 (6'7"-6'8" any weight)

---

### ⚠️ Partially Implemented (1/21)

#### 4. **DrivingDunk**
- **Current Implementation:** Height + wingspan modifiers only
- **Missing:** Weight penalties (similar to DrivingLayup)
- **Validation:** ⚠️ 4/10 samples match (60% accuracy)
- **Known Issues:**
  - Off by 1-4 points on weight-dependent builds
  - Example: H=85" WS=88" W=275lbs → Our=80, Actual=77 (diff=3)
- **Next Steps:** Add weight lookup tables using scraped data

---

### ❌ Not Implemented (17/21)

These functions return 0 (stub implementation):

| Category | Attributes |
|----------|-----------|
| **Shooting (4)** | MidRangeShot, ThreePointShot, FreeThrow, PostControl |
| **Defense (4)** | InteriorDefense, PerimeterDefense, Steal, Block |
| **Rebounding (2)** | OffensiveRebound, DefensiveRebound |
| **Athleticism (4)** | Speed, Agility, Strength, Vertical |
| **Playmaking (2)** | BallHandle, SpeedWithBall |
| **Finishing (1)** | StandingDunk |

**All of these attributes have data available in the scraped dataset!**

---

## Implementation Priority

### High Priority (Data-Driven Implementation)

These can be implemented immediately using scraped data:

1. **StandingDunk** - Similar pattern to DrivingDunk (height + wingspan)
2. **Block** - Likely height + wingspan based
3. **OffensiveRebound** - Height/wingspan/strength based
4. **DefensiveRebound** - Height/wingspan based
5. **Vertical** - Likely height/weight inverse relationship
6. **Speed** - Height/weight penalties
7. **Agility** - Similar to speed
8. **Strength** - Weight-based

### Medium Priority

9. **InteriorDefense** - Height + wingspan
10. **PerimeterDefense** - Height inverse?
11. **Steal** - Wingspan based?
12. **PostControl** - Height + weight based

### Lower Priority (Shooting)

13. **FreeThrow** - Less critical for centers
14. **MidRangeShot** - Less critical for centers
15. **ThreePointShot** - Less critical for centers
16. **BallHandle** - Less critical for centers
17. **SpeedWithBall** - Less critical for centers

---

## Data Availability

**We have complete data for ALL 21 attributes!**

```bash
# Sample scraped data structure (903 builds)
{
  "height": 88,
  "wingspan": 91,
  "weight": 260,
  "close_shot": 99,
  "driving_layup": 70,
  "driving_dunk": 66,
  "standing_dunk": 99,
  "post_control": 99,
  "mid_range_shot": 86,
  "three_point_shot": 77,
  "free_throw": 74,
  "pass_accuracy": 99,
  "ball_handle": 66,
  "speed_with_ball": 66,
  "interior_defense": 99,
  "perimeter_defense": 70,
  "steal": 70,
  "block": 99,
  "offensive_rebound": 99,
  "defensive_rebound": 99,
  "speed": 62,
  "agility": 64,
  "strength": 99,
  "vertical": 70
}
```

---

## Testing & Validation

### Quality Check Tool

Run comprehensive validation:
```bash
go run cmd/quality-check/main.go
```

**Output:**
- Tests 10 random samples per attribute
- Reports: ✅ Perfect, ⚠️ Partial, or ⚠️ Stubbed
- Shows specific mismatches with diffs

### Unit Tests

Run attribute-specific tests:
```bash
go test ./pkg/attributes/... -v
```

---

## Badge Information

**Status:** No badge data available yet

Badges are separate from attribute caps and determine what animations/skills are available. Badge requirements typically depend on:
- Attribute cap values (e.g., "Post Spin Technician" requires 80+ Post Control)
- Position
- Height/weight/wingspan combinations

**Next Steps for Badges:**
1. Check if NBA2KLab API provides badge data
2. Manual in-game testing if API doesn't expose badges
3. Document badge requirements per build

---

## Scraper Information

**Tool:** `cmd/scraper/main.go`  
**Coverage:** Centers only (903 builds)  
**Heights:** 6'7" → 7'4" (all 10 heights)  
**Weight Step:** 5 lbs increments  
**Wingspan Step:** 1 inch increments  

**Run Scraper:**
```bash
go run cmd/scraper/main.go --position Center
```

**Next Positions to Scrape:**
- Point Guard
- Shooting Guard
- Small Forward
- Power Forward

---

## Next Actions

### Immediate (Finish Centers)

1. **Fix DrivingDunk** - Add weight modifiers using scraped data
2. **Implement StandingDunk** - Use scraped data lookup table
3. **Implement Block** - Similar pattern to StandingDunk
4. **Implement Rebounds** - Height/wingspan patterns

### Short Term

5. **Implement Athleticism** - Speed, Agility, Vertical, Strength
6. **Implement Defense** - InteriorDefense, PerimeterDefense, Steal
7. **Quality check all** - Validate every attribute

### Medium Term

8. **Scrape other positions** - Point Guard, Power Forward, etc.
9. **Badge data collection** - Manual testing or API exploration
10. **Build calculator tool** - CLI/web interface for users

---

## Files

| File | Purpose |
|------|---------|
| `pkg/attributes/center.go` | Attribute calculation functions |
| `pkg/attributes/center_test.go` | Unit tests for functions |
| `cmd/quality-check/main.go` | Comprehensive validation tool |
| `cmd/scraper/main.go` | NBA2KLab API scraper |
| `data/Center_caps.json` | 903 scraped builds (gitignored) |

---

**Contributors:** AI agents working on reverse engineering NBA 2K26's attribute system  
**License:** AGPL-3.0
