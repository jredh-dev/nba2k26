# Center Position - Attribute Findings

## Testing Notes

This document tracks observed patterns when testing Center attributes in NBA 2K26.

---

## Confirmed Patterns

### Close Shot
- **Always 99** regardless of height, weight, or wingspan
- Tested: 6'7"H 6'7"WS 215LBS → 99
- Tested: (add more as you test)

### Pass Accuracy
- **Always 99** regardless of height, weight, or wingspan
- Tested: (add your test cases)

---

## Attributes To Test

### Finishing
- [ ] Driving Layup
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
