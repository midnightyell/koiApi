package koiApi

import (
	"fmt"
	"time"
)

// Wish represents a wish in Koillection, combining fields for JSON-LD and API interactions.
type Wish struct {
	Context             *Context   `json:"@context,omitempty" access:"rw"`            // JSON-LD only
	_ID                 ID         `json:"@id,omitempty" access:"ro"`                 // JSON-LD only
	Type                string     `json:"@type,omitempty" access:"rw"`               // JSON-LD only
	ID                  ID         `json:"id,omitempty" access:"ro"`                  // Identifier
	Name                string     `json:"name" access:"rw"`                          // Wish name
	URL                 *string    `json:"url,omitempty" access:"rw"`                 // Wish URL
	Price               *string    `json:"price,omitempty" access:"rw"`               // Wish price
	Currency            *string    `json:"currency,omitempty" access:"rw"`            // Currency code
	Wishlist            *string    `json:"wishlist" access:"rw"`                      // Wishlist IRI
	Owner               *string    `json:"owner,omitempty" access:"ro"`               // Owner IRI
	Comment             *string    `json:"comment,omitempty" access:"rw"`             // Wish comment
	Image               *string    `json:"image,omitempty" access:"ro"`               // Image URL
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty" access:"ro"` // Small thumbnail URL
	Visibility          Visibility `json:"visibility,omitempty" access:"rw"`          // Visibility level
	ParentVisibility    *string    `json:"parentVisibility,omitempty" access:"ro"`    // Parent visibility
	FinalVisibility     Visibility `json:"finalVisibility,omitempty" access:"ro"`     // Effective visibility
	ScrapedFromURL      *string    `json:"scrapedFromUrl,omitempty" access:"ro"`      // Source URL
	CreatedAt           time.Time  `json:"createdAt" access:"ro"`                     // Creation timestamp
	UpdatedAt           *time.Time `json:"updatedAt,omitempty" access:"ro"`           // Update timestamp
	File                *string    `json:"file,omitempty" access:"wo"`                // Image file data

}

func (w *Wish) Summary() string {
	return fmt.Sprintf("%-40s %s", w.Name, w.ID)
}

func (w *Wish) IRI() string {
	return fmt.Sprintf("/api/wishes/%s", w.ID)
}

func (w *Wish) GetID() string {
	return string(w.ID)
}

func (w *Wish) Validate() error {
	var errs []string
	// name is required, type string; see components.schemas.Wish-w.write.required
	if w.Name == "" {
		errs = append(errs, "wish name is required")
	}
	// wishlist is required, type string or null (IRI); see components.schemas.Wish-w.write.required
	if w.Wishlist == nil {
		errs = append(errs, "wish wishlist IRI is required")
	}
	// currency follows https://schema.org/priceCurrency; see components.schemas.Wish-w.write.properties.currency
	if w.Currency != nil && *w.Currency != "" {
		if !validateCurrency(*w.Currency) {
			errs = append(errs, fmt.Sprintf("invalid currency code: %s", *w.Currency))
		}
	}
	validateVisibility(w, &errs)
	return validationErrors(&errs)
}
