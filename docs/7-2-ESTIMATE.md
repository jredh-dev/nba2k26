# 7'2" Height Estimate - NEEDS TESTING

## Prediction Based on Pattern

**Build to test:**
- Height: 7'2"
- Weight: 220 lbs (minimum for this height)
- Wingspan: 7'2" (or any valid wingspan)

**Expected Driving Layup Cap: 85**

## Reasoning

**Known data points:**
- 7'3" at 230 lbs (min weight for 7'3") → 80
- 7'4" at 230 lbs (min weight for 7'4") → 77

**Challenge:** 7'2" has different minimum weight (220 lbs vs 230 lbs)

**Pattern observed:**
- Each inch taller ≈ 3 point decrease
- Each 4 lbs heavier ≈ 1 point decrease

**Calculation:**
- 7'3" at 230 lbs = 80
- 7'2" is 1 inch shorter (+3 points) AND 10 lbs lighter (+2.5 points)
- **Estimated: 80 + 3 + 2.5 = 85.5 ≈ 85**

## Full Weight Range Estimate

If the ~4 lbs per point pattern holds:

| Weight | Estimated Cap |
|--------|---------------|
| 220    | 85 |
| 230    | 82-83 |
| 240    | 80 |
| 250    | 78 |
| 260    | 75 |
| 270    | 73 |
| 280    | 70 |
| 290    | 67 |

## How to Test

1. Create Center build: 7'2" / 220 lbs / 7'2" wingspan
2. Check Driving Layup cap
3. Report result

**If cap = 85:** Pattern confirmed! ✅  
**If cap = 83-87:** Close enough, pattern holds  
**If cap is way off:** We need to reconsider the model

---

**Testing priority:** Just test 220 lbs first to verify the base cap prediction.
