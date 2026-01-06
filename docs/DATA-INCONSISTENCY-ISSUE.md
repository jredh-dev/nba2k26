# Data Inconsistency Issue - DrivingDunk Weight Testing

## Problem Summary

We've discovered a significant discrepancy between our existing wingspan test data and new weight test data for the DrivingDunk attribute at 7'4" height.

## The Discrepancy

### Existing Data (Wingspan Tests)
All wingspan tests were conducted at **270 lbs**:
- **7'4"H / 7'4"WS / 270 lbs → 66 cap** (from TestDrivingDunk in center_test.go:544)

### New Data (Weight Tests)
Recent testing at 7'4" height with 7'4" wingspan:
- **7'4"H / 7'4"WS / 260 lbs → 64 cap** (baseline weight)
- **7'4"H / 7'4"WS / 290 lbs → 59 cap**

## Analysis

### Expected Behavior
If the weight formula is linear at ~1 cap per 6 lbs (from 260→290 data):
- 260 lbs → 64 cap ✅ (confirmed)
- 266 lbs → 63 cap (expected)
- 272 lbs → 62 cap (expected)
- 278 lbs → 61 cap (expected)
- 284 lbs → 60 cap (expected)
- 290 lbs → 59 cap ✅ (confirmed)

Therefore, **270 lbs should give ~62 cap**, not 66!

### The Problem
Our wingspan test data expects **66 cap at 270 lbs**, but weight testing suggests it should be **~62 cap**.

**This is a 4-point discrepancy!**

## Possible Explanations

### 1. **Wingspan Data Was Collected at Wrong Weight** ❓
Maybe the wingspan tests were NOT actually at 270 lbs? Possibilities:
- Tests were at 230 lbs (minimum weight)?
  - 230→260 = 30 lbs difference
  - At -1 per 6 lbs: 30÷6 = 5 points
  - 64 + 5 = 69 cap (still doesn't match 66)
- Tests were at a different weight entirely?

### 2. **Weight Formula Is Non-Linear** ❓
Maybe the weight penalty changes at different weight ranges:
- Lighter weights (230-260): Different rate
- Medium weights (260-270): Different rate
- Heavier weights (270-290): Different rate

### 3. **Rounding or Hidden Variables** ❓
Maybe there are:
- Integer rounding at different stages
- Hidden variables we haven't discovered
- Thresholds that change the formula

### 4. **Original Data Collection Error** ❌
The original wingspan test at 270 lbs might have been recorded incorrectly.

## Impact

This affects:
1. **All DrivingDunk wingspan test data** - Currently 31 test cases assume 270 lbs baseline
2. **DrivingDunk2 implementation** - Weight deficit calculation needs accurate baseline
3. **Future testing** - We need to standardize on a testing weight

## Recommendations

### Immediate Actions

1. **Re-test 7'4" at 270 lbs** 🎯
   - Build: 7'4"H / 7'4"WS / 270 lbs
   - Record the ACTUAL cap
   - This will confirm which dataset is correct

2. **Test additional weight points** 🎯
   - 7'4"H / 7'4"WS / 240 lbs
   - 7'4"H / 7'4"WS / 250 lbs
   - Map out the full curve from 230→290 lbs

3. **Verify weight formula linearity** 🎯
   - Test at 5-10 lb increments
   - Determine if rate is constant or variable

### Long-term Solutions

**Option A: Use Baseline Weight for All Tests** ✅ Recommended
- Re-collect ALL wingspan data at each height's baseline weight
- Pros: Clean separation of wingspan and weight effects
- Cons: Requires re-testing all 10 heights

**Option B: Use Standardized Test Weight** ⚠️ Not Recommended
- Keep 270 lbs as "standard testing weight"
- Document weight offset from baseline
- Pros: No re-testing needed
- Cons: Confusing, weight effects baked into wingspan data

**Option C: Three-Dimensional Lookup Tables** ⚠️ Complex
- Create height × wingspan × weight matrices
- Pros: Most accurate
- Cons: Requires extensive testing, huge data tables

## Current State

- ✅ DefaultWeight/DefaultWingspan fields added to bounds.go
- ✅ 7'4" baseline confirmed: 260 lbs, 7'7" wingspan
- ⚠️ Wingspan test data still uses 270 lbs
- ⚠️ Weight deficit not yet implemented in DrivingDunk2
- ❌ Data inconsistency unresolved

## Next Steps

**Blocking**: Re-test 7'4"H / 7'4"WS at 270 lbs to confirm which dataset is correct.

Until then, we cannot:
- Implement weight deficit in DrivingDunk2
- Trust the wingspan test data
- Proceed with testing other heights

## Testing Checklist

```
[ ] 7'4"H / 7'4"WS / 270 lbs → ? cap (CRITICAL - resolves inconsistency)
[ ] 7'4"H / 7'4"WS / 250 lbs → ? cap
[ ] 7'4"H / 7'4"WS / 240 lbs → ? cap
[ ] 7'4"H / 7'4"WS / 230 lbs → ? cap (minimum weight)
[ ] 7'4"H / 7'4"WS / 280 lbs → ? cap
[ ] Determine if weight rate is constant or variable
```

## References

- `pkg/attributes/center_test.go:544` - Wingspan test expecting 66 at 270 lbs
- `docs/center-findings.md:51` - Note about "270 lbs baseline"
- User-provided data: 260 lbs → 64 cap, 290 lbs → 59 cap
