# Data Validation Summary

## Scraped Data Overview

**Total Builds**: 903 valid builds  
**Source**: NBA2KLab MyPlayer Builder API  
**Position**: Center  
**Heights**: 6'7" to 7'4" (79-88 inches)  
**Date Scraped**: 2026-01-05

### Builds Per Height

| Height | Builds | Weight Range | Wingspan Range |
|--------|--------|--------------|----------------|
| 6'7" (79") | 84 | 215-270 lbs | 79-85" |
| 6'8" (80") | 91 | 215-275 lbs | 80-86" |
| 6'9" (81") | 91 | 215-279 lbs | 81-87" |
| 6'10" (82") | 91 | 215-279 lbs | 82-88" |
| 6'11" (83") | 91 | 215-279 lbs | 83-89" |
| 7'0" (84") | 91 | 215-279 lbs | 84-90" |
| 7'1" (85") | 91 | 220-290 lbs | 85-91" |
| 7'2" (86") | 91 | 220-290 lbs | 86-92" |
| 7'3" (87") | 91 | 230-290 lbs | 87-93" |
| 7'4" (88") | 91 | 230-290 lbs | 88-94" |

## Validation Results

### ✅ Confirmed Correct

1. **CloseShot**: Always 99 (all builds validated)
2. **PassAccuracy**: Always 99 (all builds validated)
3. **DrivingDunk Data Inconsistency Resolution**:
   - 7'4" / 7'4" / 270 lbs = **64 cap** ✅
   - Confirms weight test data was correct
   - Wingspan test data was wrong (likely collected at different weight)

### 📊 DrivingLayup Pattern (Scraped Data)

**Weight-Independent Heights** (default wingspan):
- 6'7" (79"): Always 99
- 6'8" (80"): Always 99
- 6'9" (81"): Always 98
- 6'10" (82"): Always 96

**Weight-Dependent Heights**:
- 6'11" (83"): 94 (light) → 93 (255+ lbs)
- 7'0" (84"): 93 (215) → 90 (270+)
- 7'1" (85"): 86 (220) → 79 (275+)
- 7'2" (86"): 84 (220) → 73 (280+)
- 7'3" (87"): 80 (230) → 64 (290)
- 7'4" (88"): 77 (230) → 62 (290)

**Key Insight**: Weight penalties accelerate at 7'0" and above.

### ❌ Function Discrepancies

**Our Manual DrivingLayup Function:**

1. **7'3" (87") at 260 lbs**:
   - Our function: 72
   - Scraped data: 72 ✅
   - Original test (250 lbs): 75 ✅

2. **7'4" (88") at 260 lbs**:
   - Our function: 77
   - Scraped data: 70 ❌
   - **Error**: 7 points off

**Root Cause**: Our function uses linear interpolation but the actual game mechanics appear non-linear at extreme heights.

## Data Quality Issues

### Bounds Errors in scraper/main.go

The bounds used in `cmd/scraper/main.go` had inaccuracies:

**Incorrect:**
```go
83: {MaxWeight: 290, ...} // 6'11"
84: {MaxWeight: 290, ...} // 7'0"
```

**Should be** (based on scraped valid builds):
```go
83: {MaxWeight: 279, ...} // 6'11" - weight 280+ returned errors
84: {MaxWeight: 279, ...} // 7'0" - weight 280+ returned errors
```

### API Out-of-Bounds Behavior

**Expected**: ~1001 builds (10 heights × variable combinations)  
**Actual**: 903 valid builds

**Errors**: 98 builds (9.8%) returned "no results" from API
- Mostly at extreme weight/wingspan combinations
- Confirms in-game restrictions not fully documented

## Next Steps

1. ✅ **Update DrivingLayup function** with scraped data
2. ✅ **Fix bounds data** in `pkg/attributes/bounds.go`
3. ⏳ **Implement remaining attributes** (DrivingDunk with weight, StandingDunk, etc.)
4. ⏳ **Generate test fixtures** from scraped data
5. ⏳ **Scrape remaining positions** (PG, SG, SF, PF)

## Files

- **Scraped Data**: `data/Center_caps.json` (gitignored)
- **Validation Script**: `cmd/validate/main.go`
- **Scraper**: `pkg/scraper/nba2klab.go`
- **Functions**: `pkg/attributes/center.go`
