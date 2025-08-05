package koiApi

import (
	"fmt"
	"time"
)

// ItemInterface defines methods for interacting with Item resources.
type ItemInterface interface {
	Create(client Client) (*Item, error)                                           // HTTP POST /api/items
	Delete(client Client, itemID ...ID) error                                      // HTTP DELETE /api/items/{id}
	Get(client Client, itemID ...ID) (*Item, error)                                // HTTP GET /api/items/{id}
	GetCollection(client Client, itemID ...ID) (*Collection, error)                // HTTP GET /api/items/{id}/collection
	IRI() string                                                                   // /api/items/{id}
	Patch(client Client, itemID ...ID) (*Item, error)                              // HTTP PATCH /api/items/{id}
	Update(client Client, itemID ...ID) (*Item, error)                             // HTTP PUT /api/items/{id}
	UploadImage(client Client, file []byte, itemID ...ID) (*Item, error)           // HTTP POST /api/items/{id}/image
	UploadImageByFile(client Client, filename string, itemID ...ID) (*Item, error) // HTTP POST /api/items/{id}/image
	Summary() string
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
func (i *Item) Create(client Client) (*Item, error) {
	return client.CreateItem(i)
}

// Delete
func (i *Item) Delete(client Client, itemID ...ID) error {
	id := i.whichID(itemID...)
	return client.DeleteItem(id)
}

// Get
func (i *Item) Get(client Client, itemID ...ID) (*Item, error) {
	id := i.whichID(itemID...)
	return client.GetItem(id)
}

// GetCollection
func (i *Item) GetCollection(client Client, itemID ...ID) (*Collection, error) {
	id := i.whichID(itemID...)
	return client.GetItemCollection(id)
}

// IRI
func (i *Item) IRI() string {
	return fmt.Sprintf("/api/items/%s", i.ID)
}

// List
func (i *Item) List(client Client) ([]*Item, error) {
	var allItems []*Item
	for page := 1; ; page++ {
		items, err := client.ListItems()
		if err != nil {
			return nil, fmt.Errorf("failed to list items on page %d: %w", err)
		}
		if len(items) == 0 {
			break
		}
		allItems = append(allItems, items...)
	}
	return allItems, nil
}

// ListData
func (i *Item) ListData(client Client, itemID ...ID) ([]*Datum, error) {
	id := i.whichID(itemID...)
	var allData []*Datum
	for page := 1; ; page++ {
		data, err := client.ListItemData(id)
		if err != nil {
			return nil, fmt.Errorf("failed to list data for ID %s: %w", id, err)
		}
		if len(data) == 0 {
			break
		}
		allData = append(allData, data...)
	}
	return allData, nil
}

// ListLoans
func (i *Item) ListLoans(client Client, itemID ...ID) ([]*Loan, error) {
	id := i.whichID(itemID...)
	var allLoans []*Loan
	for page := 1; ; page++ {
		loans, err := client.ListItemLoans(id)
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
func (a *Item) ListLoans() ([]*Loan, error) {
	return nil, fmt.Errorf("ListLoans not implemented for Album")
}

// Patch
func (a *Item) Patch() (*Item, error) {
	return Patch(a)
}

// Update
func (a *Item) Update() (*Item, error) {
	return Update(a)
}

// UploadImage
func (a *Item) UploadImage(file []byte) (*Item, error) {
	return Upload(a, file)
}

// UploadImageFromFile
func (a *Item) UploadImageFromFile(filename string) (*Item, error) {
	return UploadFromFile(a, filename)
}
