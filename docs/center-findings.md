# Center Position - Attribute Findings

## Testing Notes

This document tracks observed patterns when testing Center attributes in NBA 2K26.

---

## Confirmed Patterns

### Close Shot ✅
- **Pattern**: Fixed value (no modifiers)
- **Cap**: 99
- **Test Cases**:
  - 6'7"H 215LBS 6'7"WS → 99
  - 7'4"H 290LBS 7'10"WS → 99
  - 7'0"H 250LBS 7'4"WS → 99

### Pass Accuracy ✅
- **Pattern**: Fixed value (no modifiers)
- **Cap**: 99
- **Test Cases**:
  - 6'7"H 215LBS 6'7"WS → 99
  - 7'4"H 290LBS 7'10"WS → 99
  - 7'0"H 250LBS 7'4"WS → 99

### Driving Layup 🔄 (In Progress)
- **Pattern**: Height-based + weight affects at tall heights
- **Confirmed Values**:
  - 6'7"H (any weight) → 99
  - 7'2"H + weight variations:
    - 223 lbs → 84
    - 244 lbs → 80
    - 269 lbs → 75
    - 290 lbs → 71
    - (Full range: 71-84, **13 point spread**)
  - 7'3"H + weight variations:
    - 230 lbs → 80
    - 250 lbs → 75
    - 270 lbs → 70
    - 290 lbs → 64
    - (Full range: 64-80, **16 point spread**)
  - 7'4"H + weight variations:
    - 230 lbs → 77
    - 250 lbs → 75
    - 270 lbs → 67
    - 290 lbs → 62
    - (Full range: 62-77, **15 point spread**)
- **Need to Test**: Heights 6'8" through 7'1" to understand when weight starts mattering
- **Pattern Analysis**: 
  - At 7'2": ~5 lbs per cap point (67 lbs = 13 points)
  - At 7'3": ~4 lbs per cap point (60 lbs = 16 points)
  - At 7'4": ~4 lbs per cap point (60 lbs = 15 points)
  - **Observation**: Wider intervals at 7'2", pattern NOT perfectly consistent across heights

---

## Attributes To Test

### Finishing
- [x] Close Shot (✅ Always 99)
- [x] Pass Accuracy (✅ Always 99)
- [🔄] Driving Layup (Height-based, testing in progress)
- [ ] Driving Dunk
- [ ] Standing Dunk
- [ ] Post Control

### Shooting
- [ ] Mid-Range Shot
- [ ] Three-Point Shot
- [ ] Free Throw

### Playmaking
- [ ] Ball Handle
- [ ] Speed With Ball

### Defense/Rebounding
- [ ] Interior Defense
- [ ] Perimeter Defense
- [ ] Steal
- [ ] Block
- [ ] Offensive Rebound
- [ ] Defensive Rebound

### Athleticism
- [ ] Speed
- [ ] Agility
- [ ] Strength
- [ ] Vertical

---

## Pattern Template

When you find a pattern, document it like this:

### [Attribute Name]
- **Pattern**: Description (e.g., "Wingspan only", "Height + Weight combo", "Always fixed")
- **Base Cap**: X (if applicable)
- **Modifiers**:
  - Height: describe effect
  - Weight: describe effect  
  - Wingspan: describe effect
- **Test Cases**:
  - 6'7"H 215LBS 6'7"WS → result
  - 6'8"H 220LBS 6'9"WS → result
