#!/bin/bash
# Find the absolute best center build for badge availability

set -e

BADGE_CHECKER="./bin/badge-checker"

echo "Searching for optimal center build (badge count)..."
echo "This will test ~200 realistic builds"
echo "=================================================="
echo ""

# Create temp file
tmpfile=$(mktemp)

# Test realistic center builds
for h_in in 79 80 81 82 83 84 85 86 87 88; do
    h_ft=$((h_in / 12))
    h_inch=$((h_in % 12))
    
    # Test wingspan variations (+0 to +8 inches from height)
    for ws_delta in 0 2 4 6 8; do
        ws_in=$((h_in + ws_delta))
        ws_ft=$((ws_in / 12))
        ws_inch=$((ws_in % 12))
        
        # Test weight variations (220, 240, 260, 280)
        for wt in 220 240 260 280; do
            echo -n "Testing ${h_ft}'${h_inch}\" / ${ws_ft}'${ws_inch}\" / ${wt}lbs... "
            
            count=$($BADGE_CHECKER --height "$h_in" --wingspan "$ws_in" --weight "$wt" 2>/dev/null | grep -E "💎|🥇|🥈|🥉|🔶" | wc -l | tr -d ' ')
            
            echo "$count badges"
            echo "$count|$h_in|$ws_in|$wt" >> "$tmpfile"
        done
    done
done

echo ""
echo "=================================================="
echo "TOP 10 BUILDS BY BADGE COUNT"
echo "=================================================="
echo ""

sort -rn -t'|' -k1 "$tmpfile" | head -10 | while IFS='|' read -r count h ws wt; do
    h_ft=$((h / 12))
    h_in=$((h % 12))
    ws_ft=$((ws / 12))
    ws_in=$((ws % 12))
    echo "$count badges: ${h_ft}'${h_in}\" height, ${ws_ft}'${ws_in}\" wingspan, ${wt}lbs"
done

echo ""
echo "Showing detailed breakdown of #1 build..."
echo "=================================================="
echo ""

# Get the top build
top=$(sort -rn -t'|' -k1 "$tmpfile" | head -1)
IFS='|' read -r count h ws wt <<< "$top"

$BADGE_CHECKER --height "$h" --wingspan "$ws" --weight "$wt"

rm -f "$tmpfile"
