# NBA 2K26 Character Attribute System

Reverse engineering the attribute calculation system for NBA 2K26 character creation.

## Overview

This project maps and models how physical characteristics (height, weight, wingspan) affect attribute caps and floors for different player positions. The goal is to reverse engineer the underlying formulas used by the game.

## Data Structure

Character attributes are stored in YAML files organized by position:

```
data/
  center/
    close_shot.yaml
    driving_layup.yaml
    driving_dunk.yaml
    ...
```

Each attribute file defines how height, weight, and wingspan modifiers affect the attribute's cap.

## License

AGPL-3.0
