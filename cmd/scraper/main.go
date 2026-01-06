// SPDX-License-Identifier: AGPL-3.0
// Copyright (C) 2025 NBA 2K26 Attribute System

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jredh-dev/nba2k26/pkg/scraper"
)

func main() {
	var (
		position   = flag.String("position", "Center", "Position to scrape (Center, Point Guard, etc.)")
		outputFile = flag.String("output", "", "Output JSON file (default: data/<position>_caps.json)")
		sample     = flag.Bool("sample", false, "Run small sample scrape for testing")
	)
	flag.Parse()

	if *outputFile == "" {
		// Default output path
		*outputFile = filepath.Join("data", fmt.Sprintf("%s_caps.json", *position))
	}

	client := scraper.NewClient()

	var results []*scraper.AttributeCaps
	var err error

	if *sample {
		fmt.Println("Running sample scrape (6'7\" Center, limited range)...")
		results, err = client.ScrapeRange(
			*position,
			[2]int{79, 79},      // Height: 6'7" only
			[2]int{79, 82},      // Wingspan: 6'7" to 6'10"
			[3]int{215, 230, 5}, // Weight: 215-230 (step 5)
		)
	} else if *position == "Center" {
		fmt.Println("Scraping all valid Center builds...")
		results, err = scrapeCenters(client)
	} else {
		fmt.Printf("Position %s not yet implemented. Use --position Center\n", *position)
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error scraping: %v\n", err)
		os.Exit(1)
	}

	// Save results to JSON
	if err := saveJSON(*outputFile, results); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving results: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n✅ Successfully scraped %d builds\n", len(results))
	fmt.Printf("📁 Saved to: %s\n", *outputFile)
}

// scrapeCenters scrapes all valid Center builds using physical bounds
func scrapeCenters(client *scraper.Client) ([]*scraper.AttributeCaps, error) {
	// Define Center bounds (heights 79-88 inches / 6'7" to 7'4")
	bounds := map[int]struct {
		MinWeight     int
		MaxWeight     int
		MinWingspan   int
		MaxWingspan   int
		DefaultWeight int
	}{
		79: {MinWeight: 215, MaxWeight: 270, MinWingspan: 79, MaxWingspan: 85, DefaultWeight: 243}, // 6'7"
		80: {MinWeight: 215, MaxWeight: 275, MinWingspan: 80, MaxWingspan: 86, DefaultWeight: 245}, // 6'8"
		81: {MinWeight: 215, MaxWeight: 285, MinWingspan: 81, MaxWingspan: 87, DefaultWeight: 250}, // 6'9"
		82: {MinWeight: 215, MaxWeight: 285, MinWingspan: 82, MaxWingspan: 88, DefaultWeight: 250}, // 6'10"
		83: {MinWeight: 215, MaxWeight: 290, MinWingspan: 83, MaxWingspan: 89, DefaultWeight: 253}, // 6'11"
		84: {MinWeight: 215, MaxWeight: 290, MinWingspan: 84, MaxWingspan: 90, DefaultWeight: 253}, // 7'0"
		85: {MinWeight: 220, MaxWeight: 290, MinWingspan: 85, MaxWingspan: 91, DefaultWeight: 255}, // 7'1"
		86: {MinWeight: 220, MaxWeight: 290, MinWingspan: 86, MaxWingspan: 92, DefaultWeight: 255}, // 7'2"
		87: {MinWeight: 230, MaxWeight: 290, MinWingspan: 87, MaxWingspan: 93, DefaultWeight: 260}, // 7'3"
		88: {MinWeight: 230, MaxWeight: 290, MinWingspan: 88, MaxWingspan: 94, DefaultWeight: 260}, // 7'4"
	}

	return client.ScrapeCentersByBounds(bounds)
}

// saveJSON writes results to a JSON file
func saveJSON(filename string, data interface{}) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Marshal JSON with indentation
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Write to file
	if err := os.WriteFile(filename, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
