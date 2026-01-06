# NBA2KLab Scraper

Command-line tool to scrape attribute caps from NBA2KLab's MyPlayer Builder API.

## Installation

```bash
go build -o bin/scraper cmd/scraper/main.go
```

## Usage

### Sample Scrape (Testing)

Scrape a small subset of builds to verify API connectivity:

```bash
go run cmd/scraper/main.go --sample
```

This scrapes 6'7" Centers with limited wingspan/weight combinations.

### Full Center Scrape

Scrape all valid Center builds (all heights, wingspans, and weights):

```bash
go run cmd/scraper/main.go --position Center
```

**Warning**: This makes ~2,000 API calls and takes approximately 3-4 minutes with rate limiting.

### Custom Output

```bash
go run cmd/scraper/main.go --position Center --output my_data.json
```

## Output Format

Results are saved as JSON array of `AttributeCaps` objects:

```json
[
  {
    "position": "Center",
    "height": 79,
    "wingspan": 82,
    "weight": 243,
    "close_shot": 99,
    "driving_layup": 99,
    "driving_dunk": 90,
    ...
  }
]
```

## Rate Limiting

The scraper includes a 100ms delay between requests to avoid overwhelming the API.

## Data Source

Data is fetched from: `https://www.nba2klab.com/.netlify/functions/char`

This is the same API used by NBA2KLab's web-based MyPlayer Builder tool.
