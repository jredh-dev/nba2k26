# Center Badge Optimization Results

**Date:** 2026-01-09  
**Tool:** badge-checker CLI  
**Limitation:** Only 4/21 attributes implemented (Close Shot, Driving Layup, Driving Dunk, Pass Accuracy)

## Executive Summary

The optimal center build for badge availability (with current attribute implementations) is:
- **Height:** 6'11" (83 inches)
- **Wingspan:** 7'4" (88 inches)  
- **Weight:** 220-300 lbs (weight has minimal impact with current attributes)
- **Available Badges:** 8 badges (5 Finishing, 3 Playmaking)

## Top 5 Builds

All top builds share 6'11" height and 7'4" wingspan with varying weights:

| Rank | Height | Wingspan | Weight | Badges | Notes |
|------|--------|----------|--------|--------|-------|
| 1 | 6'11" | 7'4" | Any | 8 | Tied - weight doesn't affect current attributes |
| 2 | 6'9" | 7'2" | 240 lbs | 8 | Slightly shorter with proportional wingspan |
| 3 | 6'10" | 7'0" - 7'2" | 240-250 lbs | 8 | Balanced build |

## Detailed Build Analysis

### Optimal Build: 6'11" / 7'4" / 260 lbs

**Calculated Attributes:**
- Close Shot: 99
- Driving Layup: 93
- Driving Dunk: 91
- Pass Accuracy: 99

**Available Badges (8):**

**Finishing (5):**
- 💎 Float Game (Legendary) - Secondary badge: Close Shot 99 ≥ 98 OR Layup 93 < 98
  - Unlocked via Close Shot (99 ≥ 98 for Legendary)
- 💎 Paint Prodigy (Legendary) - Close Shot + Layup based
- 💎 Versatile Visionary (Legendary) - Multi-attribute
- 🥇 Aerial Wizard (Hall of Fame) - Secondary badge: Driving Dunk 91 ≥ 89 OR Standing Dunk 0
  - Unlocked via Driving Dunk (91 ≥ 89 for HoF)
- 🥈 Layup Mixmaster (Gold) - Layup based

**Playmaking (3):**
- 💎 Bail Out (Legendary) - Pass Accuracy 99 ≥ 98
- 💎 Break Starter (Legendary) - Pass Accuracy 99 ≥ 98
- 💎 Dimer (Legendary) - Pass Accuracy 99 ≥ 98

### Comparison: 6'9" / 7'2" / 240 lbs

**Calculated Attributes:**
- Close Shot: 99
- Driving Layup: 99 (higher than 6'11"!)
- Driving Dunk: 95 (higher than 6'11"!)
- Pass Accuracy: 99

**Available Badges (8):**

**Finishing (5):**
- 💎 Aerial Wizard (Legendary) - Driving Dunk 95 ≥ 97? No, but 95 ≥ 89 for HoF
  - Actually shows Legendary - likely Layup requirement OR bug
- 💎 Float Game (Legendary)
- 💎 Paint Prodigy (Legendary)
- 💎 Versatile Visionary (Legendary)
- 🥇 Layup Mixmaster (Hall of Fame) - Higher tier due to better Driving Layup

**Playmaking (3):**
- Same 3 Legendary badges

**Analysis:** Shorter centers appear to get BETTER finishing attributes but same badge count.

### Tall Build: 7'2" / 7'6" / 280 lbs

**Calculated Attributes:**
- Close Shot: 99
- Driving Layup: 83
- Driving Dunk: 80
- Pass Accuracy: 99

**Available Badges (7):**

**Finishing (4):** 
- Only 4 badges instead of 5
- Lost Layup Mixmaster (Driving Layup 83 too low)
- Aerial Wizard downgraded to Silver (Driving Dunk 80 ≥ 70 for Silver)

**Playmaking (3):**
- Same 3 Legendary badges

**Analysis:** Taller centers sacrifice finishing badge availability.

## Key Insights

### Height Impact
- **Shorter centers (6'9" - 6'11"):** Better finishing attributes (Driving Layup, Driving Dunk)
- **Taller centers (7'2"+):** Worse finishing attributes
- **Pass Accuracy:** Maxed at 99 for all heights tested

### Wingspan Impact
- **Longer wingspan:** Generally better (7'4" optimal for 6'11")
- **Wingspan ratio:** ~+5-6 inches over height is ideal
- **Too long:** Diminishing returns or penalties (7'8" wingspan on 7'0" center lost badges)

### Weight Impact
- **Minimal impact** with current attribute implementations
- Likely affects Strength, Vertical, Speed (not yet implemented)

### Badge Patterns
1. **Secondary badges dominate:** Most available badges are Secondary type (ANY requirement)
2. **Pass Accuracy badges:** Easy to max (all 3 Playmaking badges at Legendary)
3. **Finishing badges:** Height-dependent, shorter = more badges
4. **Missing badges:** Most badges require unimplemented attributes (Vertical, Strength, Block, etc.)

## Predictions After Full Implementation

When all 21 attributes are implemented:

**Expected Changes:**
1. **Taller centers will gain Defense/Rebounding badges** (Block, Interior Defense, Defensive Rebound)
2. **Weight will matter** for Strength-based badges (Brick Wall, Bulldozer, Physical Finisher)
3. **Wingspan will affect Block** (chase-down blocks, Pogo Stick)
4. **Current optimal build may change** as more badges become available

**Likely New Optimal Builds:**
- **Balanced:** 6'11" - 7'0" with 7'2" - 7'4" wingspan (Finishing + Defense)
- **Rim Protector:** 7'2" - 7'3" with 7'6"+ wingspan (Defense + Rebounding)
- **Post Scorer:** 6'10" - 7'0" with shorter wingspan (Finishing + Post badges)

## Methodology

**Test Range:**
- Heights: 6'6" - 7'4" (78 - 88 inches)
- Wingspans: 6'10" - 7'8" (82 - 92 inches)
- Weights: 220 - 300 lbs (20 lb increments)
- Total builds tested: 385

**Limitations:**
- Only 4/21 attributes functional (19%)
- Badge availability artificially limited
- Cannot test Vertical, Strength, Block, Speed, Agility
- Many meta badges unavailable (Posterizer, Intimidator, Pogo Stick, Brick Wall)

## Next Steps

1. **Implement remaining attributes** (Standing Dunk, Vertical, Strength, Block priority)
2. **Re-run optimization** with full attribute set
3. **Test in-game** to validate attribute formulas
4. **Compare to competitive builds** from NBA 2K community
5. **Build recommendations** based on playstyle (Rim Runner, Post Scorer, Paint Beast, etc.)

## Raw Data

Top 10 builds by badge count:

```
8 badges: 6'11" height, 7'4" wingspan, 300lbs
8 badges: 6'11" height, 7'4" wingspan, 280lbs
8 badges: 6'11" height, 7'4" wingspan, 260lbs
8 badges: 6'11" height, 7'4" wingspan, 240lbs
8 badges: 6'11" height, 7'4" wingspan, 220lbs
8 badges: 6'10" height, 7'4" wingspan, 300lbs
8 badges: 6'10" height, 7'4" wingspan, 280lbs
8 badges: 6'10" height, 7'4" wingspan, 260lbs
8 badges: 6'10" height, 7'4" wingspan, 240lbs
8 badges: 6'10" height, 7'4" wingspan, 220lbs
```

Pattern: 7'4" wingspan dominates, heights 6'10" - 6'11" optimal, weight irrelevant.
