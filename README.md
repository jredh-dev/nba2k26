# NBA 2K26 Character Attribute System

Reverse engineering the attribute calculation system for NBA 2K26 character creation.

## Overview

This project documents how physical characteristics (height, weight, wingspan) affect attribute caps and floors for different player positions. Data is collected by systematically testing character combinations in the game and recording the thresholds where attribute caps change.

## Data Structure

Character attributes are documented in YAML files organized by position:

```
data/
  center/
    close_shot.yaml
    driving_layup.yaml
    driving_dunk.yaml
    standing_dunk.yaml
    post_control.yaml
    mid_range_shot.yaml
    three_point_shot.yaml
    free_throw.yaml
    pass_accuracy.yaml
    ball_handle.yaml
    speed_with_ball.yaml
    interior_defense.yaml
    perimeter_defense.yaml
    steal.yaml
    block.yaml
    offensive_rebound.yaml
    defensive_rebound.yaml
    speed.yaml
    agility.yaml
    strength.yaml
    vertical.yaml
```

## YAML Format

Each attribute file documents:
- Base cap (maximum potential)
- Modifiers based on height, weight, and wingspan
- Thresholds where caps change

Example:
```yaml
position: Center
attribute: driving_dunk
modifiers:
  wingspan:
    6'7": +6   # At 6'7" wingspan, +6 adjustment
    6'8": +8   # At 6'8" wingspan, +8 adjustment
    6'9": +9   # At 6'9" wingspan, +9 adjustment
    6'10": +10 # At 6'10" wingspan, +10 adjustment
```

## Data Collection Process

1. Select position (e.g., Center)
2. Choose attribute to test (e.g., Driving Dunk)
3. Start at minimum height for position
4. Test all weight/wingspan combinations
5. Record cap changes and thresholds
6. Document in YAML file

## License

AGPL-3.0
