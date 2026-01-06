# Quick Testing Reference

## 🎯 CURRENT PRIORITY: Driving Dunk Weight Modifier Testing

✅ **Driving Layup is COMPLETE** (all 10 heights, 6'7" - 7'4")
✅ **Driving Dunk wingspan data is COMPLETE** (all 10 heights, baseline weight)

### Goal: Determine Weight Modifier Rate

We need to test how weight affects Driving Dunk caps to implement an additive modifier system.

**Hypothesis**: `Final Cap = Base(height) + WingspanModifier + WeightModifier`

### Test at 7'0" Height, 7'3" Wingspan

**Known baseline**: 7'0"H, 270LBS, 7'3"WS → **86**

**Test these weights:**

| Test # | Height | Weight | Wingspan | Expected Baseline | Actual Cap? | Weight Effect |
|--------|--------|--------|----------|-------------------|-------------|---------------|
| 1      | 7'0"   | 220    | 7'3"     | 86                | ❓          | ? |
| 2      | 7'0"   | 240    | 7'3"     | 86                | ❓          | ? |
| 3      | 7'0"   | 260    | 7'3"     | 86                | ❓          | ? |
| **4**  | **7'0"**   | **270**    | **7'3"**     | **86**                | **86 ✅**      | **0 (baseline)** |
| 5      | 7'0"   | 280    | 7'3"     | 86                | ❓          | ? |
| 6      | 7'0"   | 290    | 7'3"     | 86                | ❓          | ? |

**Analysis**: 
- If 220 lbs → 90 cap, then weight modifier = +4 for -50 lbs (baseline is 270)
- This gives us the rate: **+4 / 50 lbs = +1 per 12.5 lbs lighter**
- Or: **-1 per 12.5 lbs heavier**

### Verification Tests (if modifier system works)

**Test at 6'7" height to verify consistency:**

| Height | Weight | Wingspan | Current Known | With Weight Modifier? |
|--------|--------|----------|---------------|------------------------|
| 6'7"   | 220    | 6'9"     | ?             | 98 + modifier          |
| 6'7"   | 270    | 6'9"     | 98 ✅         | 98 (baseline)          |
| 6'7"   | 290    | 6'9"     | ?             | 98 + modifier          |

If modifiers match, system is confirmed additive!

---

## Testing Steps

### In Game:
1. Go to MyCareer → Create New Build
2. Select **Center** position
3. Set exact height/weight/wingspan from table above
4. Navigate to attributes screen
5. **Record the MAX cap** for Driving Layup (or Driving Dunk)
6. Screenshot if helpful

### After Testing:
Report back with format:
```
6'8"H 220LBS 6'9"WS → Driving Layup cap = XX
```

I'll update the code/tests with your findings.

---

## Completed Attributes ✅

### Driving Layup (COMPLETE - All 10 Heights)
- **Pattern**: Height-based with weight impact at 6'11"+
- **6'7" - 6'10"**: Fixed caps (99, 99, 98, 96) - weight doesn't matter
- **6'11"+**: Weight affects cap, impact increases with height
- **Full range**: 62-99 across all builds
- **Wingspan**: Does NOT affect this attribute

### If Linear (Most Likely for Driving Layup):
- Each inch of height changes cap by same amount
- Example: 99, 96, 93, 90... (decreases by 3)
- Easy to implement as formula

### If Stepped:
- Groups of heights share same cap
- Example: 6'7-6'9 = 99, 6'10-7'0 = 90, etc.
- Implement as ranges

### If Non-Linear:
- Changes vary by height
- Example: 99, 97, 94, 90, 85... (decreases vary)
- Implement as lookup table

---

## Quick Notes Space

Use this space to jot down findings as you test:

```
6'8" → 
6'9" → 
6'10" → 
...
```

Once you have 3-4 data points, we can start seeing the pattern!
