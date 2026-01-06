# Quick Testing Reference

## 🎯 CURRENT PRIORITY: Driving Dunk (Wingspan Test)

✅ **Driving Layup is COMPLETE** (all 10 heights, 6'7" - 7'4")

### Driving Dunk - Wingspan Impact Test

**Question**: Does wingspan affect Driving Dunk caps?

**Test at 7'0" height, 250 lbs, varying wingspan:**

| Test # | Height | Weight | Wingspan | Driving Dunk Cap? |
|--------|--------|--------|----------|-------------------|
| 1      | 7'0"   | 250    | 7'0"     | ❓ (minimum wingspan) |
| 2      | 7'0"   | 250    | 7'2"     | ❓ |
| 3      | 7'0"   | 250    | 7'4"     | ❓ |
| 4      | 7'0"   | 250    | 7'6"     | ❓ (maximum wingspan) |

**What we're looking for:**
- ✅ If all 4 values are **the same** → Wingspan doesn't matter (like Driving Layup)
- ⚠️ If values are **different** → Wingspan DOES matter, need systematic testing

### If Wingspan Matters (Values Different):

Test minimum/maximum builds to establish range:

| Build Type | Height | Weight | Wingspan | Driving Dunk Cap? |
|------------|--------|--------|----------|-------------------|
| Min build  | 6'7"   | 215    | 6'7"     | ❓ |
| Max build  | 7'4"   | 290    | 7'10"    | ❓ |

Then test systematically like we did with Driving Layup.

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
