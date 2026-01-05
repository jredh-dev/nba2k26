# Testing Workflow

## Goal
Systematically discover attribute cap formulas for NBA 2K26 Center position by testing specific character builds.

## Current Priority: Driving Layup (All Heights)

We know:
- 6'7" → 99 cap
- 7'4" → 62 cap

**Next: Test all intermediate heights to find the pattern**

### Test Queue (Priority Order)

1. **Driving Layup - All Heights**
   - [ ] 6'8" / 220 lbs / 6'9" wingspan → ?
   - [ ] 6'9" / 225 lbs / 6'10" wingspan → ?
   - [ ] 6'10" / 230 lbs / 6'11" wingspan → ?
   - [ ] 6'11" / 235 lbs / 7'0" wingspan → ?
   - [ ] 7'0" / 240 lbs / 7'1" wingspan → ?
   - [ ] 7'1" / 245 lbs / 7'2" wingspan → ?
   - [ ] 7'2" / 250 lbs / 7'3" wingspan → ?
   - [ ] 7'3" / 255 lbs / 7'4" wingspan → ?

2. **Driving Dunk - Wingspan Variations (Single Height)**
   - Test 7'0" height with all valid wingspans (7'0" through 7'6")
   - Purpose: Determine if wingspan affects cap
   
3. **Standing Dunk - Same Pattern**

4. **Speed - Height Variations**
   - Likely follows height pattern (taller = slower)

5. **Strength - Weight Variations**
   - Likely follows weight pattern (heavier = stronger)

## Testing Protocol

### For Each Test:

1. **Create Character in Game**
   - Position: Center
   - Set exact height, weight, wingspan from test queue
   
2. **Record Attribute Cap**
   - Navigate to attribute screen
   - Record the MAX value (not current value)
   
3. **Document Result**
   - Add to `docs/center-findings.md`
   - Format: `6'8"H 220LBS 6'9"WS → XX cap`

4. **Add Test Case**
   - Update `pkg/attributes/center_test.go`
   - Add new test case with recorded cap

5. **Implement Logic (if pattern clear)**
   - Update function in `pkg/attributes/center.go`
   - Run `go test ./pkg/attributes/... -v`

## Pattern Recognition Tips

### Height-Based Patterns
- If values decrease/increase linearly with height
- Usually 3-4 point change per inch
- Example: 6'7"=99, 6'8"=96, 6'9"=93

### Wingspan-Based Patterns
- If values change with wingspan but not height
- Usually 1-2 point change per inch
- Common for dunking/blocking attributes

### Weight-Based Patterns
- If values change with weight but not height/wingspan
- Usually for Strength, possibly Speed
- May have threshold ranges (e.g., <240 vs 240+)

### Fixed Values
- Attribute is same for ALL builds
- Close Shot and Pass Accuracy are confirmed fixed

### Combo Patterns
- Some attributes may depend on multiple factors
- Example: Height + Wingspan together affect cap
- These are rarer but possible

## Example: Finding a Linear Pattern

**Driving Layup Testing Results:**
```
6'7"  → 99
6'8"  → 96
6'9"  → 93
6'10" → 90
```

**Pattern Recognition:**
- Decreases by 3 per inch
- Formula: `99 - ((heightInches - 79) * 3)`
- Can implement as calculation instead of lookup table

## After Finding Pattern

### 1. Document in center-findings.md
```markdown
### Driving Layup
- **Pattern**: Height-based linear decrease
- **Formula**: 99 - ((height - 79) * 3)
- **Range**: 99 (6'7") to 62 (7'4")
- **Test Cases**: [list all tested builds]
```

### 2. Implement in center.go
```go
func DrivingLayup(heightInches, weightLbs, wingspanInches int) int {
    // Linear decrease: 3 points per inch above 6'7"
    return 99 - ((heightInches - MustLengthToInches("6'7")) * 3)
}
```

### 3. Add Comprehensive Tests
```go
{
    name:  "all heights tested",
    tests: []struct{height int; want int}{
        {79, 99}, {80, 96}, {81, 93}, // ... etc
    },
}
```

## Quick Reference

**Start Game Testing**: Create Center → Set build → Check attribute cap  
**Record Result**: `docs/center-findings.md`  
**Add Test**: `pkg/attributes/center_test.go`  
**Implement**: `pkg/attributes/center.go`  
**Verify**: `go test ./pkg/attributes/... -v`

---

**Current Status**: Refactoring complete, tests passing, ready for systematic testing
