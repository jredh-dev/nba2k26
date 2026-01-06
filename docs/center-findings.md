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

### Driving Layup ✅ (COMPLETE!)
- **Pattern**: Height is primary factor, weight affects caps at 6'11" and taller
- **Wingspan Impact**: None - wingspan does NOT affect this attribute
- **Confirmed Values**:
  - 6'7"H → 99 (weight independent)
  - 6'8"H → 99 (weight independent)
  - 6'9"H → 98 (weight independent)
  - 6'10"H → 96 (weight independent)
  - 6'11"H → 92-94 (weight dependent, 3 point range)
  - 7'0"H → 89-93 (weight dependent, 5 point range)
  - 7'1"H → 77-86 (weight dependent, 10 point range)
  - 7'2"H → 71-84 (weight dependent, 13 point range)
  - 7'3"H → 64-80 (weight dependent, 16 point range)
  - 7'4"H → 62-77 (weight dependent, 15 point range)
- **Key Insights**: 
  - Weight doesn't matter until 6'11"
  - Weight impact increases dramatically with height
  - Taller players have much wider cap ranges based on weight

### Driving Dunk 🔄 (PARTIAL - Testing in Progress)
- **Pattern**: Wingspan affects caps (CONFIRMED)
- **Height Impact**: Unknown (need to test more heights)
- **Weight Impact**: Appears not to matter at 6'7" (need confirmation at other heights)
- **Confirmed Values at 6'7" height**:
  - 6'7"H 270LBS 6'7"WS → 95
  - 6'7"H 270LBS 6'8"WS → 97
  - 6'7"H 270LBS 6'9"WS → 98
  - 6'7"H 270LBS 7'1"WS → 99
- **Key Insights**:
  - Wingspan DOES affect Driving Dunk (unlike Driving Layup)
  - At 6'7" height: +2 cap per inch of wingspan increase (approximately)
  - Need to test: Other heights, weight variations, full wingspan range

---

## Attributes To Test

### Finishing
- [x] Close Shot (✅ Always 99)
- [x] Pass Accuracy (✅ Always 99)
- [x] Driving Layup (✅ COMPLETE - All 10 heights)
- [🔄] Driving Dunk (Testing in progress - wingspan DOES affect caps!)
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
