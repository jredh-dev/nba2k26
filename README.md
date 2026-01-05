# NBA 2K26 Character Attribute System

Reverse engineering the attribute calculation system for NBA 2K26 character creation.

## Overview

This project reverse engineers how physical characteristics (height, weight, wingspan) affect attribute caps for different player positions. Instead of storing raw data, we encode the game's logic as Go functions that can calculate attribute caps for any combination of physical characteristics.

## Project Structure

```
nba2k26/
├── docs/
│   └── center-findings.md    # Testing notes and discovered patterns
├── pkg/
│   └── attributes/
│       ├── center.go         # Center attribute calculators
│       └── center_test.go    # Tests validating formulas
└── data/                     # (deprecated) Old YAML structure
```

## Approach

1. **Test in-game** - Document patterns in `docs/center-findings.md`
2. **Encode as Go functions** - Implement calculators in `pkg/attributes/center.go`
3. **Validate with tests** - Write tests proving your formulas match the game

## Example

```go
// CloseShot is always 99 regardless of physical characteristics
func CloseShot(height, weight, wingspan string) int {
    return 99
}

// Test validates this works for all combinations
func TestCloseShot(t *testing.T) {
    assert.Equal(t, 99, CloseShot("6'7\"", "215", "6'7\""))
    assert.Equal(t, 99, CloseShot("7'4\"", "290", "7'10\""))
}
```

## Running Tests

```bash
go test ./pkg/attributes/... -v
```

## Current Status

**Confirmed Attributes:**
- ✅ Close Shot - Always 99
- ✅ Pass Accuracy - Always 99

**Stubbed (Need Testing):**
- ⏳ Driving Layup, Driving Dunk, Standing Dunk, Post Control
- ⏳ Mid-Range Shot, Three-Point Shot, Free Throw
- ⏳ Ball Handle, Speed With Ball
- ⏳ Interior Defense, Perimeter Defense, Steal, Block
- ⏳ Offensive Rebound, Defensive Rebound
- ⏳ Speed, Agility, Strength, Vertical

See `docs/center-findings.md` for testing notes.

## License

AGPL-3.0
