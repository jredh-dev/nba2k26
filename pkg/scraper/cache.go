// SPDX-License-Identifier: AGPL-3.0
// Copyright (C) 2025 NBA 2K26 Attribute System - Cache

package scraper

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Cache handles persistent storage of API responses
type Cache struct {
	cacheDir string
	enabled  bool
}

// NewCache creates a new cache instance
// If cacheDir is empty, uses default: ./data/cache/
func NewCache(cacheDir string) *Cache {
	if cacheDir == "" {
		cacheDir = filepath.Join("data", "cache")
	}

	return &Cache{
		cacheDir: cacheDir,
		enabled:  true,
	}
}

// Disable turns off caching
func (c *Cache) Disable() {
	c.enabled = false
}

// Enable turns on caching
func (c *Cache) Enable() {
	c.enabled = true
}

// ensureCacheDir creates the cache directory if it doesn't exist
func (c *Cache) ensureCacheDir() error {
	if !c.enabled {
		return nil
	}
	return os.MkdirAll(c.cacheDir, 0755)
}

// buildCacheKey generates a unique key for a build configuration
func (c *Cache) buildCacheKey(position string, height, wingspan, weight int) string {
	// Create a unique string representation
	key := fmt.Sprintf("%s_%d_%d_%d", position, height, wingspan, weight)
	// Hash it for a clean filename
	hash := sha256.Sum256([]byte(key))
	return fmt.Sprintf("%x.json", hash[:8])
}

// Get retrieves cached data for a build configuration
func (c *Cache) Get(position string, height, wingspan, weight int) (*AttributeCaps, error) {
	if !c.enabled {
		return nil, fmt.Errorf("cache disabled")
	}

	filename := c.buildCacheKey(position, height, wingspan, weight)
	filepath := filepath.Join(c.cacheDir, filename)

	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err // Cache miss
	}

	var attrs AttributeCaps
	if err := json.Unmarshal(data, &attrs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cached data: %w", err)
	}

	return &attrs, nil
}

// Set stores data in the cache
func (c *Cache) Set(attrs *AttributeCaps) error {
	if !c.enabled {
		return nil
	}

	if err := c.ensureCacheDir(); err != nil {
		return fmt.Errorf("failed to create cache directory: %w", err)
	}

	filename := c.buildCacheKey(attrs.Position, attrs.Height, attrs.Wingspan, attrs.Weight)
	filepath := filepath.Join(c.cacheDir, filename)

	data, err := json.MarshalIndent(attrs, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	if err := os.WriteFile(filepath, data, 0644); err != nil {
		return fmt.Errorf("failed to write cache file: %w", err)
	}

	return nil
}

// Has checks if a cache entry exists
func (c *Cache) Has(position string, height, wingspan, weight int) bool {
	if !c.enabled {
		return false
	}

	filename := c.buildCacheKey(position, height, wingspan, weight)
	filepath := filepath.Join(c.cacheDir, filename)

	_, err := os.Stat(filepath)
	return err == nil
}

// Clear removes all cached entries
func (c *Cache) Clear() error {
	if !c.enabled {
		return nil
	}

	return os.RemoveAll(c.cacheDir)
}

// Size returns the number of cached entries
func (c *Cache) Size() (int, error) {
	if !c.enabled {
		return 0, nil
	}

	entries, err := os.ReadDir(c.cacheDir)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil
		}
		return 0, err
	}

	count := 0
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".json" {
			count++
		}
	}

	return count, nil
}
