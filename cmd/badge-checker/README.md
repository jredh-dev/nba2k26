# Badge Checker CLI

Query badge availability for NBA 2K26 character builds based on physical measurements.

## Features

- Check all available badges for a build
- Filter badges by category (Finishing, Shooting, Playmaking, Defense, Rebounding, Physicals, All-Around)
- Query specific badges
- Show calculated attribute caps
- Filter by minimum tier (Bronze, Silver, Gold, Hall of Fame, Legendary)

## Usage

```bash
# Build the tool
go build -o bin/badge-checker ./cmd/badge-checker/

# Check all badges for a 7'0" Center with 7'3" wingspan, 260 lbs
./bin/badge-checker --height 7-0 --wingspan 7-3 --weight 260

# Check only finishing badges
./bin/badge-checker --height 84 --wingspan 87 --weight 260 --category Finishing

# Check specific badge
./bin/badge-checker --height 7-0 --wingspan 7-3 --weight 260 --badge Posterizer

# Show calculated attribute values
./bin/badge-checker --height 7-0 --wingspan 7-3 --weight 260 --show-attributes

# Show all badges including unavailable (None tier)
./bin/badge-checker --height 7-0 --wingspan 7-3 --weight 260 --all

# Filter by minimum tier (only show Gold and above)
./bin/badge-checker --height 7-0 --wingspan 7-3 --weight 260 --min-tier Gold
```

## Input Formats

**Height/Wingspan:**
- Feet-inches: `7-0`, `6-6`, `7-3`
- Total inches: `84`, `78`, `87`

**Categories:**
- Finishing (Inside Scoring)
- Shooting (Outside Scoring)
- Playmaking
- Defense
- Rebounding
- Physicals (General Offense)
- AllAround

**Tiers:**
- Bronze
- Silver
- Gold
- HoF (Hall of Fame)
- Legendary

## Output Example

```
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Build: Center
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Height:   84" (7'0")
Wingspan: 87" (7'3")
Weight:   260 lbs

Available Badges (7):

Finishing (4):
  💎 Float Game (Legendary)
  💎 Paint Prodigy (Legendary)
  💎 Versatile Visionary (Legendary)
  🥈 Aerial Wizard (Gold)

Playmaking (3):
  💎 Bail Out (Legendary)
  💎 Break Starter (Legendary)
  💎 Dimer (Legendary)
```

## Current Limitations

**Note:** Badge availability is calculated based on attribute caps. Currently, only 3/21 attributes have been implemented:
- Close Shot
- Driving Layup
- Driving Dunk
- Pass Accuracy

As more attributes are implemented, more badges will become available. Badges requiring unimplemented attributes (e.g., Posterizer requiring Vertical) will show as unavailable until those attributes are added.

## Data Source

Badge requirements are sourced from NBA2KLab's badge requirements page:
https://www.nba2klab.com/badge-requirements

The data includes:
- 64 unique badges across 7 categories
- Attribute requirements for each tier (Bronze → Legendary)
- Height restrictions for position-specific badges
- Primary vs Secondary badge types

## License

AGPL-3.0
