package koiApi

import (
	"context"
	"fmt"
	"os"
	"time"
)

// ItemInterface defines methods for interacting with Item resources.
type ItemInterface interface {
	Create(ctx context.Context, client Client) (*Item, error)                                           // HTTP POST /api/items
	Delete(ctx context.Context, client Client, itemID ...ID) error                                      // HTTP DELETE /api/items/{id}
	Get(ctx context.Context, client Client, itemID ...ID) (*Item, error)                                // HTTP GET /api/items/{id}
	GetCollection(ctx context.Context, client Client, itemID ...ID) (*Collection, error)                // HTTP GET /api/items/{id}/collection
	IRI() string                                                                                        // /api/items/{id}
	List(ctx context.Context, client Client) ([]*Item, error)                                           // HTTP GET /api/items
	ListData(ctx context.Context, client Client, itemID ...ID) ([]*Datum, error)                        // HTTP GET /api/items/{id}/data
	ListLoans(ctx context.Context, client Client, itemID ...ID) ([]*Loan, error)                        // HTTP GET /api/items/{id}/loans
	ListRelatedItems(ctx context.Context, client Client, itemID ...ID) ([]*Item, error)                 // HTTP GET /api/items/{id}/related_items
	ListTags(ctx context.Context, client Client, itemID ...ID) ([]*Tag, error)                          // HTTP GET /api/items/{id}/tags
	Patch(ctx context.Context, client Client, itemID ...ID) (*Item, error)                              // HTTP PATCH /api/items/{id}
	Update(ctx context.Context, client Client, itemID ...ID) (*Item, error)                             // HTTP PUT /api/items/{id}
	UploadImage(ctx context.Context, client Client, file []byte, itemID ...ID) (*Item, error)           // HTTP POST /api/items/{id}/image
	UploadImageByFile(ctx context.Context, client Client, filename string, itemID ...ID) (*Item, error) // HTTP POST /api/items/{id}/image
}

// Item represents an item within a collection, combining fields for JSON-LD and API interactions.
type Item struct {
	Context             *Context   `json:"@context,omitempty" access:"rw"`            // JSON-LD only
	_ID                 ID         `json:"@id,omitempty" access:"ro"`                 // JSON-LD only
	Type                string     `json:"@type,omitempty" access:"rw"`               // JSON-LD only
	ID                  ID         `json:"id,omitempty" access:"ro"`                  // Identifier
	Name                string     `json:"name" access:"rw"`                          // Item name
	Quantity            int        `json:"quantity" access:"rw"`                      // Item quantity
	Collection          *string    `json:"collection" access:"rw"`                    // Collection IRI
	Owner               *string    `json:"owner,omitempty" access:"ro"`               // Owner IRI
	Image               *string    `json:"image,omitempty" access:"ro"`               // Image URL
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty" access:"ro"` // Small thumbnail URL
	ImageLargeThumbnail *string    `json:"imageLargeThumbnail,omitempty" access:"ro"` // Large thumbnail URL
	SeenCounter         int        `json:"seenCounter,omitempty" access:"ro"`         // View count
	Visibility          Visibility `json:"visibility,omitempty" access:"rw"`          // Visibility level
	ParentVisibility    *string    `json:"parentVisibility,omitempty" access:"ro"`    // Parent visibility
	FinalVisibility     Visibility `json:"finalVisibility,omitempty" access:"ro"`     // Effective visibility
	ScrapedFromURL      *string    `json:"scrapedFromUrl,omitempty" access:"ro"`      // Source URL
	CreatedAt           time.Time  `json:"createdAt" access:"ro"`                     // Creation timestamp
	UpdatedAt           *time.Time `json:"updatedAt,omitempty" access:"ro"`           // Update timestamp
	Tags                []string   `json:"tags,omitempty" access:"wo"`                // Tag IRIs
	RelatedItems        []string   `json:"relatedItems,omitempty" access:"wo"`        // Related item IRIs
	File                *string    `json:"file,omitempty" access:"wo"`                // Image file data
}

// whichID
func (i *Item) whichID(itemID ...ID) ID {
	if len(itemID) > 0 {
		return itemID[0]
	}
	return i.ID
}

// Create
func (i *Item) Create(ctx context.Context, client Client) (*Item, error) {
	return client.CreateItem(ctx, i)
}

// Delete
func (i *Item) Delete(ctx context.Context, client Client, itemID ...ID) error {
	id := i.whichID(itemID...)
	return client.DeleteItem(ctx, id)
}

// Get
func (i *Item) Get(ctx context.Context, client Client, itemID ...ID) (*Item, error) {
	id := i.whichID(itemID...)
	return client.GetItem(ctx, id)
}

// GetCollection
func (i *Item) GetCollection(ctx context.Context, client Client, itemID ...ID) (*Collection, error) {
	id := i.whichID(itemID...)
	return client.GetItemCollection(ctx, id)
}

// IRI
func (i *Item) IRI() string {
	return fmt.Sprintf("/api/items/%s", i.ID)
}

// List
func (i *Item) List(ctx context.Context, client Client) ([]*Item, error) {
	var allItems []*Item
	for page := 1; ; page++ {
		items, err := client.ListItems(ctx, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list items on page %d: %w", page, err)
		}
		if len(items) == 0 {
			break
		}
		allItems = append(allItems, items...)
	}
	return allItems, nil
}

// ListData
func (i *Item) ListData(ctx context.Context, client Client, itemID ...ID) ([]*Datum, error) {
	id := i.whichID(itemID...)
	var allData []*Datum
	for page := 1; ; page++ {
		data, err := client.ListItemData(ctx, id, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list data for ID %s on page %d: %w", id, page, err)
		}
		if len(data) == 0 {
			break
		}
		allData = append(allData, data...)
	}
	return allData, nil
}

// ListLoans
func (i *Item) ListLoans(ctx context.Context, client Client, itemID ...ID) ([]*Loan, error) {
	id := i.whichID(itemID...)
	var allLoans []*Loan
	for page := 1; ; page++ {
		loans, err := client.ListItemLoans(ctx, id, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list loans for ID %s on page %d: %w", id, page, err)
		}
		if len(loans) == 0 {
			break
		}
		allLoans = append(allLoans, loans...)
	}
	return allLoans, nil
}

// ListRelatedItems
func (i *Item) ListRelatedItems(ctx context.Context, client Client, itemID ...ID) ([]*Item, error) {
	id := i.whichID(itemID...)
	var allRelatedItems []*Item
	for page := 1; ; page++ {
		relatedItems, err := client.ListItemRelatedItems(ctx, id, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list related items for ID %s on page %d: %w", id, page, err)
		}
		if len(relatedItems) == 0 {
			break
		}
		allRelatedItems = append(allRelatedItems, relatedItems...)
	}
	return allRelatedItems, nil
}

// ListTags
func (i *Item) ListTags(ctx context.Context, client Client, itemID ...ID) ([]*Tag, error) {
	id := i.whichID(itemID...)
	var allTags []*Tag
	for page := 1; ; page++ {
		tags, err := client.ListItemTags(ctx, id, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list tags for ID %s on page %d: %w", id, page, err)
		}
		if len(tags) == 0 {
			break
		}
		allTags = append(allTags, tags...)
	}
	return allTags, nil
}

// Patch
func (i *Item) Patch(ctx context.Context, client Client, itemID ...ID) (*Item, error) {
	id := i.whichID(itemID...)
	return client.PatchItem(ctx, id, i)
}

// Update
func (i *Item) Update(ctx context.Context, client Client, itemID ...ID) (*Item, error) {
	id := i.whichID(itemID...)
	return client.UpdateItem(ctx, id, i)
}

// UploadImage
func (i *Item) UploadImage(ctx context.Context, client Client, file []byte, itemID ...ID) (*Item, error) {
	id := i.whichID(itemID...)
	return client.UploadItemImage(ctx, id, file)
}

// UploadImageByFile
func (i *Item) UploadImageByFile(ctx context.Context, client Client, filename string, itemID ...ID) (*Item, error) {
	id := i.whichID(itemID...)
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return client.UploadItemImage(ctx, id, file)
}
