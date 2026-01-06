# Attribute Modifier System Design

## Problem Statement

Some attributes (like Driving Dunk) are affected by multiple physical characteristics in an **additive** way:
- Height determines a base cap
- Wingspan adds/subtracts from that base
- Weight adds/subtracts from that total

Current lookup table approach doesn't scale when all three dimensions vary.

## Proposed Solution: Additive Modifier System

Instead of storing every combination, calculate the cap as:

```
Final Cap = Base(height) + WingspanModifier(height, wingspan) + WeightModifier(height, weight)
```

### Example: Driving Dunk

**Current data (270 lbs baseline):**
- 7'0"H, 270LBS, 7'0"WS → 83
- 7'0"H, 270LBS, 7'3"WS → 86 (+3 from wingspan)
- 7'0"H, 270LBS, 7'6"WS → 89 (+6 from wingspan)

**If weight also matters:**
- 7'0"H, 250LBS, 7'3"WS → 88 (+2 from lighter weight?)
- 7'0"H, 290LBS, 7'3"WS → 84 (-2 from heavier weight?)

### Implementation Approach

```go
func DrivingDunk(heightInches, weightLbs, wingspanInches int) int {
    // 1. Get base cap for this height (at min wingspan, baseline weight)
    base := getDrivingDunkBase(heightInches)
    
    // 2. Calculate wingspan modifier
    wingspanMod := getDrivingDunkWingspanModifier(heightInches, wingspanInches)
    
    // 3. Calculate weight modifier
    weightMod := getDrivingDunkWeightModifier(heightInches, weightLbs)
    
    // 4. Return total (clamped to 0-99)
    return clamp(base + wingspanMod + weightMod, 0, 99)
}

func getDrivingDunkBase(heightInches int) int {
    // Base cap at minimum wingspan and baseline weight (270 lbs)
    // This is what we currently have in our lookup tables
    switch heightInches {
    case MustLengthToInches("6'7"):
        return 95  // with 6'7" wingspan
    case MustLengthToInches("7'0"):
        return 83  // with 7'0" wingspan
    // ... etc
    }
}

func getDrivingDunkWingspanModifier(heightInches, wingspanInches int) int {
    // Calculate how much wingspan differs from minimum
    minWingspan := getMinWingspan(heightInches)
    wingspanDiff := wingspanInches - minWingspan
    
    // Roughly +1 cap per inch of wingspan (may vary by height)
    return wingspanDiff
}

func getDrivingDunkWeightModifier(heightInches, weightLbs int) int {
    // Calculate modifier based on weight difference from baseline (270 lbs)
    baselineWeight := 270
    weightDiff := weightLbs - baselineWeight
    
    // Heavier = lower cap (negative modifier)
    // Lighter = higher cap (positive modifier)
    // Need to test: exact ratio (e.g., -1 per 10 lbs?)
    return -(weightDiff / 10)  // Example: every 10 lbs = -1 cap
}
```

## Testing Strategy

To implement the modifier system, we need to test:

### 1. Weight Variations at One Height
Pick 7'0" height with 7'3" wingspan (middle of range), test different weights:
- 7'0"H, 220LBS, 7'3"WS → ? (lightest)
- 7'0"H, 240LBS, 7'3"WS → ?
- 7'0"H, 260LBS, 7'3"WS → ?
- 7'0"H, 270LBS, 7'3"WS → 86 (known baseline)
- 7'0"H, 280LBS, 7'3"WS → ?
- 7'0"H, 290LBS, 7'3"WS → ? (heaviest)

**Goal**: Determine weight modifier rate (e.g., -1 per X lbs)

### 2. Verify Additivity
Test if modifiers truly add:
- 7'0"H, 270LBS, 7'0"WS → 83 (base)
- 7'0"H, 270LBS, 7'3"WS → 86 (base + 3 from wingspan)
- 7'0"H, 250LBS, 7'3"WS → ? (should be 86 + weight modifier)

If the last value = 86 + (weight modifier from step 1), then system is additive!

### 3. Cross-Check Other Heights
Test same weight variations at 6'7" and 7'4" to verify modifier is consistent.

## Alternative: Non-Additive System

If testing shows modifiers are NOT additive (i.e., wingspan effect changes based on weight), we may need:
- Full lookup tables (current approach)
- Or more complex formulas with interaction terms

## Benefits of Modifier System

1. **Compact Storage**: Instead of 10 heights × 7 wingspans × 6 weights = 420 values, store ~30 base values + modifier logic
2. **Easier Testing**: Only need to test weight at 1-2 heights to determine modifier
3. **Interpolation**: Can handle any weight value, not just tested ones
4. **Pattern Discovery**: Easier to see underlying game mechanics

## Next Steps

1. **Test weight variations** (6 data points at 7'0" height)
2. **Analyze if additive** (compare expected vs actual)
3. **Implement modifier functions** if additive
4. **Update tests** to verify modifier system matches data
5. **Document findings** in center-findings.md

## Notes

- Driving Layup already uses a similar pattern (height base + weight modifier)
- This approach may apply to other attributes (Standing Dunk, Block, Speed, etc.)
- Keep current lookup tables as fallback if modifier system doesn't fit
