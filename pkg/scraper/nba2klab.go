// SPDX-License-Identifier: AGPL-3.0
// Copyright (C) 2025 NBA 2K26 Attribute System

package scraper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// NBA2KLabAPIURL is the endpoint for fetching attribute caps
	NBA2KLabAPIURL = "https://www.nba2klab.com/.netlify/functions/char"

	// NBA2KLabAuthToken is the JWT token for API authentication
	// This is a public token used by the NBA2KLab website frontend
	NBA2KLabAuthToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6ImFkbWluIiwiYWRtaW4iOnRydWUsImp0aSI6ImQ2MTEwYzAxLWMwYjUtNDUzNy1iNDZhLTI0NTk5Mjc2YjY1NiIsImlhdCI6MTU5MjU2MDk2MCwiZXhwIjoxNTkyNTY0NjE5fQ.QgFSQtFaK_Ktauadttq1Is7f9w0SUtKcL8xCmkAvGLw"

	// GameYear is the NBA 2K game year (26 for NBA 2K26)
	GameYear = 26
)

// Client handles requests to the NBA2KLab API
type Client struct {
	httpClient *http.Client
	baseURL    string
	authToken  string
	year       int
}

// NewClient creates a new NBA2KLab API client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL:   NBA2KLabAPIURL,
		authToken: NBA2KLabAuthToken,
		year:      GameYear,
	}
}

// AttributeCaps represents the complete attribute data for a player build
type AttributeCaps struct {
	Position         string `json:"position"`
	Height           int    `json:"height"`
	Wingspan         int    `json:"wingspan"`
	Weight           int    `json:"weight"`
	CloseShot        int    `json:"close_shot"`
	DrivingLayup     int    `json:"driving_layup"`
	DrivingDunk      int    `json:"driving_dunk"`
	StandingDunk     int    `json:"standing_dunk"`
	PostControl      int    `json:"post_control"`
	MidRangeShot     int    `json:"mid_range_shot"`
	ThreePointShot   int    `json:"three_point_shot"`
	FreeThrow        int    `json:"free_throw"`
	PassAccuracy     int    `json:"pass_accuracy"`
	BallHandle       int    `json:"ball_handle"`
	SpeedWithBall    int    `json:"speed_with_ball"`
	InteriorDefense  int    `json:"interior_defense"`
	PerimeterDefense int    `json:"perimeter_defense"`
	Steal            int    `json:"steal"`
	Block            int    `json:"block"`
	OffensiveRebound int    `json:"offensive_rebound"`
	DefensiveRebound int    `json:"defensive_rebound"`
	Speed            int    `json:"speed"`
	Strength         int    `json:"strength"`
	Vertical         int    `json:"vertical"`
	Agility          int    `json:"agility"`
}

// apiRequest represents the request body for the NBA2KLab API
type apiRequest struct {
	Filters []filter `json:"filters"`
	Year    int      `json:"year"`
}

// filter represents a single filter parameter
type filter struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

// apiResponse represents the response from the NBA2KLab API
type apiResponse struct {
	Results []AttributeCaps `json:"results"`
}

// GetAttributeCaps fetches attribute caps for a specific build configuration
func (c *Client) GetAttributeCaps(position string, heightInches, wingspanInches, weight int) (*AttributeCaps, error) {
	// Construct request body
	reqBody := apiRequest{
		Filters: []filter{
			{Name: "position", Value: position},
			{Name: "height", Value: heightInches},
			{Name: "wingspan", Value: wingspanInches},
			{Name: "weight", Value: weight},
		},
		Year: c.year,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", c.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.authToken))

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	// Parse response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var apiResp apiResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// Validate response
	if len(apiResp.Results) == 0 {
		return nil, fmt.Errorf("no results returned for position=%s height=%d wingspan=%d weight=%d",
			position, heightInches, wingspanInches, weight)
	}

	return &apiResp.Results[0], nil
}

// ScrapeRange fetches attribute caps for a range of builds
// heightRange, wingspanRange: [min, max] inclusive in inches
// weightRange: [min, max, step] (e.g., [215, 270, 5] for 215, 220, 225, ..., 270)
func (c *Client) ScrapeRange(position string, heightRange, wingspanRange [2]int, weightRange [3]int) ([]*AttributeCaps, error) {
	var results []*AttributeCaps

	minHeight, maxHeight := heightRange[0], heightRange[1]
	minWingspan, maxWingspan := wingspanRange[0], wingspanRange[1]
	minWeight, maxWeight, stepWeight := weightRange[0], weightRange[1], weightRange[2]

	total := (maxHeight - minHeight + 1) * (maxWingspan - minWingspan + 1) * ((maxWeight-minWeight)/stepWeight + 1)
	current := 0

	for height := minHeight; height <= maxHeight; height++ {
		for wingspan := minWingspan; wingspan <= maxWingspan; wingspan++ {
			for weight := minWeight; weight <= maxWeight; weight += stepWeight {
				current++
				fmt.Printf("Scraping %d/%d: %s H=%d\" WS=%d\" W=%dlbs\n",
					current, total, position, height, wingspan, weight)

				caps, err := c.GetAttributeCaps(position, height, wingspan, weight)
				if err != nil {
					fmt.Printf("  ⚠️  Error: %v\n", err)
					continue // Skip errors and continue
				}

				results = append(results, caps)

				// Rate limiting: 100ms delay between requests
				time.Sleep(100 * time.Millisecond)
			}
		}
	}

	return results, nil
}

// ScrapeCentersByBounds scrapes all valid Center builds based on physical bounds
// bounds: map[heightInches]PhysicalBounds
func (c *Client) ScrapeCentersByBounds(bounds map[int]struct {
	MinWeight     int
	MaxWeight     int
	MinWingspan   int
	MaxWingspan   int
	DefaultWeight int
}) ([]*AttributeCaps, error) {
	var results []*AttributeCaps

	// Calculate total iterations for progress tracking
	total := 0
	for _, b := range bounds {
		heightCount := 1
		wingspanCount := b.MaxWingspan - b.MinWingspan + 1
		weightCount := (b.MaxWeight-b.MinWeight)/5 + 1 // Step by 5 lbs
		total += heightCount * wingspanCount * weightCount
	}

	current := 0

	// Iterate through each height
	for height, b := range bounds {
		for wingspan := b.MinWingspan; wingspan <= b.MaxWingspan; wingspan++ {
			for weight := b.MinWeight; weight <= b.MaxWeight; weight += 5 {
				current++
				fmt.Printf("Scraping %d/%d: Center H=%d\" WS=%d\" W=%dlbs\n",
					current, total, height, wingspan, weight)

				caps, err := c.GetAttributeCaps("Center", height, wingspan, weight)
				if err != nil {
					fmt.Printf("  ⚠️  Error: %v\n", err)
					continue
				}

				results = append(results, caps)

				// Rate limiting: 100ms delay between requests
				time.Sleep(100 * time.Millisecond)
			}
		}
	}

	return results, nil
}
