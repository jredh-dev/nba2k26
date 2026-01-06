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

### Driving Dunk ✅ (WINGSPAN DATA COMPLETE - Weight Effects TODO)
- **Pattern**: Height + Wingspan both affect caps (Weight ALSO matters!)
- **Wingspan Impact**: Confirmed - larger wingspan = higher cap
- **Height Impact**: Taller heights have lower base caps
- **Weight Impact**: Discovered but not yet implemented (need modifier system)
- **Confirmed Wingspan Data (at 270 lbs baseline)**:
  - 6'7"H: 95-99 cap range (6'7"WS → 7'1"WS)
  - 6'8"H: 94-99 cap range (6'8"WS → 7'2"WS)
  - 6'9"H: 92-99 cap range (6'9"WS → 7'3"WS)
  - 6'10"H: 90-96 cap range (6'10"WS → 7'4"WS)
  - 6'11"H: 86-92 cap range (6'11"WS → 7'5"WS)
  - 7'0"H: 83-89 cap range (7'0"WS → 7'6"WS)
  - 7'1"H: 77-82 cap range (7'1"WS → 7'7"WS)
  - 7'2"H: 72-77 cap range (7'2"WS → 7'8"WS)
  - 7'3"H: 68-72 cap range (7'3"WS → 7'9"WS)
  - 7'4"H: 66-70 cap range (7'4"WS → 7'10"WS)
- **Key Insights**:
  - All 10 heights tested with full wingspan ranges
  - Wingspan increases cap by ~1 point per inch (not perfectly linear)
  - Height decreases base cap (taller = lower base)
  - Weight also affects cap (needs modifier system to implement)
  - Current implementation: baseline weight (270 lbs) for all heights

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
