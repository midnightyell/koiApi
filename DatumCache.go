package koiApi

import (
	"fmt"
)

type DatumCache struct {
	cache map[string]map[string][]Item // map[Label]map[Value][]Item
}

// NewDatumCache initializes a new DatumCache.
func NewDatumCache() *DatumCache {
	return &DatumCache{
		cache: make(map[string]map[string][]Item),
	}
}

// AddDatum adds a datum and its associated item to the cache.
func (dc *DatumCache) AddDatum(datum Datum, item Item) {
	if _, exists := dc.cache[datum.Label]; !exists {
		dc.cache[datum.Label] = make(map[string][]Item)
	}
	dc.cache[datum.Label][datum.Value] = append(dc.cache[datum.Label][datum.Value], item)
}

func (c *koiClient) GetCollectionDataByTitle(title string) (*DatumCache, error) {
	// Step 1: List all collections
	collections, err := List(Collection{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch collections: %w", err)
	}

	// Step 2: Find the collection with the matching title
	var targetCollection *Collection
	for _, collection := range collections {
		if collection.Title == title {
			targetCollection = &collection
			break
		}
	}

	if targetCollection == nil {
		return nil, fmt.Errorf("collection with title '%s' not found", title)
	}

	// Step 3: List all items for the found collection
	items, err := ListItems(*targetCollection)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch items for collection '%s': %w", title, err)
	}

	// Step 4: Fetch all datum for each item and populate the cache
	cache := NewDatumCache()
	for _, item := range items {
		data, err := ListData(item)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch data for item %s in collection '%s': %w", item.ID, title, err)
		}

		for _, datum := range data {
			cache.AddDatum(*datum, *item)
		}
	}

	return cache, nil
}
