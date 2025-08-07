package koiApi

import (
	"fmt"
	"time"
)

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

func (i *Item) Summary() string {
	return fmt.Sprintf("%8.8s   %s", i.ID[len(i.ID)-8:], i.Name)
}

// IRI
func (i *Item) IRI() string {
	return fmt.Sprintf("/api/items/%s", i.ID)
}

func (i *Item) GetID() string {
	return string(i.ID)
}

func (i *Item) Validate() error {
	var errs []string
	// name is required, type string; see components.schemas.Item-i.write.required
	if i.Name == "" {
		errs = append(errs, "item name is required")
	}
	// collection is required, type string or null (IRI); see components.schemas.Item-i.write.required
	if i.Collection == nil || *i.Collection == "" {
		errs = append(errs, "item collection IRI is required")
	}
	// quantity minimum 1, type integer; see components.schemas.Item-i.write.properties.quantity
	if i.Quantity < 1 {
		i.Quantity = 1 // The API says it should use a default of 1, but errors out instead
		//errs = append(errs, "item quantity must be at least 1")
	}
	validateVisibility(i, &errs)
	return validationErrors(&errs)
}
