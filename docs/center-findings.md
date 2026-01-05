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
- **Pattern**: Height-based + weight affects at 7'4"
- **Confirmed Values**:
  - 6'7"H (any weight) → 99
  - 7'4"H + weight variations:
    - 230 lbs → 77
    - 232 lbs → 76
    - 236 lbs → 75
    - 240 lbs → 74
    - 244 lbs → 73
    - 249 lbs → 72
    - 252 lbs → 71
    - 257 lbs → 70
    - 261 lbs → 69
    - 265 lbs → 68
    - 269 lbs → 67
    - 273 lbs → 66
    - 277 lbs → 65
    - 281 lbs → 64
    - 287 lbs → 63
    - 290 lbs → 62
- **Need to Test**: All intermediate heights (6'8" through 7'3") to understand full pattern
- **Hypothesis**: Pattern appears formulaic (~4 lbs per cap point at 7'4"), but game's rounding method unclear. Hard-coding values until more data collected.

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
