# Quick Testing Reference

## Next Tests Needed

### 🎯 Priority 1: Driving Layup (Complete Height Pattern)

Test these builds and record the Driving Layup cap:

| Height | Weight | Wingspan | Cap? |
|--------|--------|----------|------|
| 6'7"   | 215    | 6'7"     | ✅ 99 |
| **6'8"**   | **220**    | **6'9"**     | **❓** |
| **6'9"**   | **225**    | **6'10"**    | **❓** |
| **6'10"**  | **230**    | **6'11"**    | **❓** |
| **6'11"**  | **235**    | **7'0"**     | **❓** |
| **7'0"**   | **240**    | **7'1"**     | **❓** |
| **7'1"**   | **245**    | **7'2"**     | **❓** |
| **7'2"**   | **250**    | **7'3"**     | **❓** |
| **7'3"**   | **255**    | **7'4"**     | **❓** |
| 7'4"   | 260    | 7'5"     | ✅ 62 |

### 🎯 Priority 2: Driving Dunk (Wingspan Test)

Once Driving Layup pattern is found, test Driving Dunk with **one height** but varying wingspans:

**All tests at 7'0" height, 240 lbs:**

| Wingspan | Driving Dunk Cap? |
|----------|-------------------|
| 7'0"     | ❓ |
| 7'1"     | ❓ |
| 7'2"     | ❓ |
| 7'3"     | ❓ |
| 7'4"     | ❓ |
| 7'5"     | ❓ |
| 7'6"     | ❓ |

**Purpose**: Determine if wingspan affects dunk caps (likely yes)

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

## Pattern Clues

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
