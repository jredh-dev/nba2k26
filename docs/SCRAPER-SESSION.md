# NBA2KLab Data Scraper - Session Summary

## What We Built

### 1. Scraper Package (`pkg/scraper/`)

**File**: `pkg/scraper/nba2klab.go`

- **`Client`**: HTTP client for NBA2KLab API with authentication
- **`AttributeCaps`**: Complete struct with all 24 attributes
- **`GetAttributeCaps()`**: Fetch single build data
- **`ScrapeRange()`**: Bulk scrape with custom ranges
- **`ScrapeCentersByBounds()`**: Scrape all valid Centers

**Features**:
- Rate limiting (100ms delay between requests)
- Error handling for invalid builds
- Progress tracking
- JWT authentication handling

### 2. Command-Line Tool (`cmd/scraper/`)

**Usage**:
```bash
# Sample scrape (16 builds)
go run cmd/scraper/main.go --sample

# Full Center scrape (~1000 builds)
go run cmd/scraper/main.go --position Center

# Custom output path
go run cmd/scraper/main.go --output my_data.json
```

**Output**: JSON array in `data/<Position>_caps.json`

### 3. Test Coverage

**File**: `pkg/scraper/nba2klab_test.go`

- ✅ `TestGetAttributeCaps`: Validates known builds
- ✅ `TestGetAttributeCaps_InvalidBuild`: Error handling
- ✅ `TestScrapeRange_SmallSample`: Bulk scraping (skipped in short mode)

**Key Test**: Resolves the 7'4" data inconsistency (confirms DrivingDunk = 64 at 270 lbs)

## API Details

**Endpoint**: `https://www.nba2klab.com/.netlify/functions/char`

**Request Format**:
```json
{
  "filters": [
    {"name": "position", "value": "Center"},
    {"name": "height", "value": 79},
    {"name": "wingspan", "value": 82},
    {"name": "weight", "value": 243}
  ],
  "year": 26
}
```

**Response Format**:
```json
{
  "results": [{
    "position": "Center",
    "height": 79,
    "wingspan": 82,
    "weight": 243,
    "close_shot": 99,
    "driving_layup": 99,
    ...
  }]
}
```

## Data Validation Findings

### Resolved: DrivingDunk Data Inconsistency

**Issue**: Wingspan test data claimed 7'4"/7'4"/270lbs → 66 cap  
**API says**: 7'4"/7'4"/270lbs → **64 cap**

**Conclusion**: 
- ✅ Weight test data was correct
- ❌ Wingspan test data was wrong (likely collected at different weight)

See: `docs/DATA-INCONSISTENCY-ISSUE.md`

## Next Steps

1. **Wait for full scrape to complete** (~1000 builds, 100 seconds)
2. **Validate scraped data** against our manual tests
3. **Update attribute functions** to use scraped data
4. **Generate test fixtures** from scraped data
5. **Scrape other positions** (PG, SG, SF, PF)

## File Structure

```
nba2k26/
├── cmd/
│   └── scraper/
│       ├── main.go          # CLI tool
│       └── README.md        # Usage docs
├── pkg/
│   ├── scraper/
│   │   ├── nba2klab.go      # API client
│   │   └── nba2klab_test.go # Tests
│   └── attributes/
│       ├── bounds.go
│       ├── center.go
│       └── ...
├── data/
│   └── Center_caps.json     # Scraped data (gitignored)
├── .gitignore               # Excludes data/*.json
└── scrape.log               # Background scraper output
```

## Key Decisions

1. **Rate Limiting**: 100ms between requests to be respectful of API
2. **Error Handling**: Continue on errors, log warnings
3. **Data Format**: JSON for easy analysis and version control
4. **Gitignore**: Exclude scraped data (can be regenerated)
5. **Progress Tracking**: Show current/total for long scrapes

## Commands Reference

```bash
# Run tests
go test ./pkg/scraper -v

# Quick test (skip API calls)
go test ./pkg/scraper -v -short

# Sample scrape
go run cmd/scraper/main.go --sample

# Full scrape (background)
nohup go run cmd/scraper/main.go --position Center > scrape.log 2>&1 &

# Check progress
tail -f scrape.log

# Validate results
jq 'length' data/Center_caps.json
jq '[.[] | {h: .height, ws: .wingspan, w: .weight}] | unique | length' data/Center_caps.json
```

## Performance

**Estimated Times**:
- Single request: ~200-300ms (API + network)
- 16 builds (sample): ~2-3 seconds
- ~1000 builds (full Center): ~100-120 seconds

**Bottleneck**: API rate limiting (100ms delay)

## Attribution

Data source: NBA2KLab (https://www.nba2klab.com/myplayer-builder)  
API discovered via browser DevTools network inspection
