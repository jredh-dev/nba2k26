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
- **Pattern**: Height-based (taller = lower cap)
- **Confirmed Values**:
  - 6'7"H → 99
  - 7'4"H → 62
- **Need to Test**: All intermediate heights (6'8" through 7'3")
- **Hypothesis**: Linear decrease, ~3-4 points per inch

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
