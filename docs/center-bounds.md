# Center Position - Physical Characteristic Bounds

This document tracks the valid ranges for weight and wingspan at each height for the Center position.

## Format

For each height:
- **Weight Range**: Min - Max (in lbs)
- **Wingspan Range**: Min - Max (in feet'inches")

---

## Confirmed Bounds

### 6'7" (Minimum Height)
- **Weight**: 215 - 270 lbs
- **Wingspan**: 6'7" - 7'1"
- **Status**: ✅ Confirmed

### 6'8"
- **Weight**: ??? - ??? lbs
- **Wingspan**: ???" - ???"
- **Status**: ⏳ Not yet tested

### 6'9"
- **Weight**: ??? - ??? lbs
- **Wingspan**: ???" - ???"
- **Status**: ⏳ Not yet tested

### 6'10"
- **Weight**: ??? - ??? lbs
- **Wingspan**: ???" - ???"
- **Status**: ⏳ Not yet tested

### 6'11"
- **Weight**: ??? - ??? lbs
- **Wingspan**: ???" - ???"
- **Status**: ⏳ Not yet tested

### 7'0"
- **Weight**: ??? - ??? lbs
- **Wingspan**: ???" - ???"
- **Status**: ⏳ Not yet tested

### 7'1"
- **Weight**: ??? - ??? lbs
- **Wingspan**: ???" - ???"
- **Status**: ⏳ Not yet tested

### 7'2"
- **Weight**: ??? - ??? lbs
- **Wingspan**: ???" - ???"
- **Status**: ⏳ Not yet tested

### 7'3"
- **Weight**: ??? - ??? lbs
- **Wingspan**: ???" - ???"
- **Status**: ⏳ Not yet tested

### 7'4" (Maximum Height)
- **Weight**: ??? - 290 lbs (min weight needs verification)
- **Wingspan**: ???" - 7'10" (min wingspan needs verification)
- **Status**: ⚠️ Partially confirmed (only maximums known)

---

## Testing Workflow

When testing a new height:

1. **Set height** in character builder
2. **Test minimum weight** - Reduce weight until builder won't allow lower
3. **Test maximum weight** - Increase weight until builder won't allow higher
4. **Test minimum wingspan** - Reduce wingspan until builder won't allow shorter
5. **Test maximum wingspan** - Increase wingspan until builder won't allow longer
6. **Document here** and update `pkg/attributes/bounds.go`

## Patterns to Watch For

As you collect data, look for:
- Does minimum weight stay constant (215) across all heights?
- Does wingspan range expand as height increases?
- Are there any height increments that have unusual bounds?
