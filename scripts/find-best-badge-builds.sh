#!/bin/bash
# Find center builds with the most available badges

set -e

BADGE_CHECKER="./bin/badge-checker"

# Build the tool if not already built
if [ ! -f "$BADGE_CHECKER" ]; then
    echo "Building badge-checker..."
    go build -o bin/badge-checker ./cmd/badge-checker/
fi

echo "Testing various center builds to find optimal badge availability..."
echo "=================================================================="
echo ""

# Function to test a build and count badges
test_build() {
    local height=$1
    local wingspan=$2
    local weight=$3
    
    local count=$($BADGE_CHECKER --height "$height" --wingspan "$wingspan" --weight "$weight" 2>/dev/null | grep -c "💎\|🥇\|🥈\|🥉\|🔶" || echo "0")
    
    echo "H: $height  WS: $wingspan  W: ${weight}lbs  →  $count badges"
}

echo "=== Small Centers (6'6\" - 6'9\") ==="
test_build "6-6" "6-10" "220"
test_build "6-6" "7-0" "220"
test_build "6-9" "7-0" "230"
test_build "6-9" "7-2" "240"
echo ""

echo "=== Medium Centers (6'10\" - 7'0\") ==="
test_build "6-10" "7-0" "240"
test_build "6-10" "7-2" "250"
test_build "7-0" "7-2" "250"
test_build "7-0" "7-3" "260"
test_build "7-0" "7-4" "270"
echo ""

echo "=== Tall Centers (7'1\" - 7'4\") ==="
test_build "7-1" "7-3" "260"
test_build "7-1" "7-5" "270"
test_build "7-2" "7-4" "270"
test_build "7-2" "7-6" "280"
test_build "7-3" "7-5" "280"
test_build "7-3" "7-7" "290"
test_build "7-4" "7-6" "290"
test_build "7-4" "7-8" "300"
echo ""

echo "=== Testing wingspan variations for 7'0\" centers ==="
for ws in 6-10 7-0 7-2 7-4 7-6 7-8; do
    test_build "7-0" "$ws" "260"
done
echo ""

echo "=== Testing weight variations for 7'0\" / 7'3\" wingspan ==="
for wt in 220 230 240 250 260 270 280 290 300; do
    test_build "7-0" "7-3" "$wt"
done
echo ""

echo "Finding top 5 builds with most badges..."
echo "========================================="

# Create temporary file for results
tmpfile=$(mktemp)

# Test a wider range and store results
for h in 78 79 80 81 82 83 84 85 86 87 88; do
    for ws in 82 84 86 88 90 92 94; do
        for wt in 220 240 260 280 300; do
            count=$($BADGE_CHECKER --height "$h" --wingspan "$ws" --weight "$wt" 2>/dev/null | grep -c "💎\|🥇\|🥈\|🥉\|🔶" || echo "0")
            echo "$count|$h|$ws|$wt" >> "$tmpfile"
        done
    done
done

# Sort and display top 5
echo ""
echo "Top 5 Builds:"
sort -rn -t'|' -k1 "$tmpfile" | head -5 | while IFS='|' read -r count h ws wt; do
    h_ft=$((h / 12))
    h_in=$((h % 12))
    ws_ft=$((ws / 12))
    ws_in=$((ws % 12))
    echo "  $count badges: ${h_ft}'${h_in}\" height, ${ws_ft}'${ws_in}\" wingspan, ${wt}lbs"
done

rm -f "$tmpfile"

echo ""
echo "Note: Badge availability is limited by attribute implementations (3/21 complete)"
echo "Full badge availability requires implementing remaining attributes."
