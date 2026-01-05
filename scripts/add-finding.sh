#!/bin/bash
# Helper script to quickly add test results to findings

# Usage: ./add-finding.sh "6'8\" 220 6'9\" DrivingLayup 96"

set -e

if [ $# -ne 5 ]; then
    echo "Usage: $0 HEIGHT WEIGHT WINGSPAN ATTRIBUTE VALUE"
    echo "Example: $0 \"6'8\\\"\" 220 \"6'9\\\"\" DrivingLayup 96"
    exit 1
fi

HEIGHT=$1
WEIGHT=$2
WINGSPAN=$3
ATTRIBUTE=$4
VALUE=$5

FINDINGS_FILE="docs/center-findings.md"

echo "Adding finding: ${HEIGHT}H ${WEIGHT}LBS ${WINGSPAN}WS → ${ATTRIBUTE} = ${VALUE}"

# Add to findings file (manual for now - could automate later)
echo ""
echo "Please manually add to $FINDINGS_FILE:"
echo "  - ${HEIGHT}H ${WEIGHT}LBS ${WINGSPAN}WS → ${VALUE}"
