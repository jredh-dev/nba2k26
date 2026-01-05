# NBA 2K26 Character Attribute System

Reverse engineering the attribute calculation system for NBA 2K26 character creation.

## Overview

This project reverse engineers how physical characteristics (height, weight, wingspan) affect attribute caps for different player positions. Instead of storing raw data, we encode the game's logic as Go functions that can calculate attribute caps for any combination of physical characteristics.

## Project Structure

```
nba2k26/
├── docs/
│   ├── center-findings.md       # Testing notes and discovered patterns
│   ├── center-bounds.md         # Physical characteristic bounds
│   ├── TESTING-WORKFLOW.md      # Systematic testing guide
│   └── QUICK-TESTING-REF.md     # Quick reference for in-game testing
├── pkg/
│   └── attributes/
│       ├── center.go            # Center attribute calculators
│       ├── center_test.go       # Tests validating formulas
│       ├── bounds.go            # Physical characteristic bounds
│       └── conversion.go        # Height/weight conversion utilities
├── scripts/
│   └── add-finding.sh           # Helper script for adding test results
└── data/                        # (deprecated) Old YAML structure
```

## Approach

1. **Test in-game** - Create builds and record attribute caps
2. **Document patterns** - Add findings to `docs/center-findings.md`
3. **Implement functions** - Encode logic in `pkg/attributes/center.go`
4. **Write tests** - Validate formulas in `pkg/attributes/center_test.go`
5. **Run tests** - Verify: `go test ./pkg/attributes/... -v`

See `docs/TESTING-WORKFLOW.md` for detailed testing process.

## Example

```go
// CloseShot is always 99 regardless of physical characteristics
func CloseShot(heightInches, weightLbs, wingspanInches int) int {
    return 99
}

// DrivingLayup decreases with height (taller = lower cap)
func DrivingLayup(heightInches, weightLbs, wingspanInches int) int {
    switch heightInches {
    case MustLengthToInches("6'7"):
        return 99
    case MustLengthToInches("7'4"):
        return 62
    default:
        return 0 // Not yet tested
    }
}

// Test validates this works for all combinations
func TestCloseShot(t *testing.T) {
    assert.Equal(t, 99, CloseShot(MustLengthToInches("6'7"), 215, MustLengthToInches("6'7")))
    assert.Equal(t, 99, CloseShot(MustLengthToInches("7'4"), 290, MustLengthToInches("7'10")))
}
```

## Running Tests

```bash
go test ./pkg/attributes/... -v
```

## Current Status

**Center Position - Confirmed Attributes:**
- ✅ Close Shot - Always 99
- ✅ Pass Accuracy - Always 99
- 🔄 Driving Layup - Height-based (testing in progress)

**Stubbed (Need Testing):**
- ⏳ Driving Dunk, Standing Dunk, Post Control
- ⏳ Mid-Range Shot, Three-Point Shot, Free Throw
- ⏳ Ball Handle, Speed With Ball
- ⏳ Interior Defense, Perimeter Defense, Steal, Block
- ⏳ Offensive Rebound, Defensive Rebound
- ⏳ Speed, Agility, Strength, Vertical

**Physical Bounds Documented:**
- ✅ All 10 center heights (6'7" → 7'4") with valid weight/wingspan ranges

See `docs/center-findings.md` for detailed testing notes.

## Quick Start for Testing

1. Check `docs/QUICK-TESTING-REF.md` for next builds to test
2. Create build in NBA 2K26 and record attribute cap
3. Report findings (format: `6'8"H 220LBS 6'9"WS → DrivingLayup = 96`)
4. AI agent will update code/tests with your findings

## License

AGPL-3.0
