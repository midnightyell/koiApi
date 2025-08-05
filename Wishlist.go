package koiApi

import (
	"fmt"
	"os"
	"time"
)

// Wishlist represents a wishlist in Koillection, combining fields for JSON-LD and API interactions.
type Wishlist struct {
	Context          *Context   `json:"@context,omitempty" access:"rw"`         // JSON-LD only
	_ID              ID         `json:"@id,omitempty" access:"ro"`              // JSON-LD only
	Type             string     `json:"@type,omitempty" access:"rw"`            // JSON-LD only
	ID               ID         `json:"id,omitempty" access:"ro"`               // Identifier
	Name             string     `json:"name" access:"rw"`                       // Wishlist name
	Owner            *string    `json:"owner,omitempty" access:"ro"`            // Owner IRI
	Color            string     `json:"color" access:"ro"`                      // Color code
	Parent           *string    `json:"parent,omitempty" access:"rw"`           // Parent wishlist IRI
	Image            *string    `json:"image,omitempty" access:"ro"`            // Image URL
	SeenCounter      int        `json:"seenCounter,omitempty" access:"ro"`      // View count
	Visibility       Visibility `json:"visibility,omitempty" access:"rw"`       // Visibility level
	ParentVisibility *string    `json:"parentVisibility,omitempty" access:"ro"` // Parent visibility
	FinalVisibility  Visibility `json:"finalVisibility,omitempty" access:"ro"`  // Effective visibility
	CreatedAt        time.Time  `json:"createdAt" access:"ro"`                  // Creation timestamp
	UpdatedAt        *time.Time `json:"updatedAt,omitempty" access:"ro"`        // Update timestamp
	File             *string    `json:"file,omitempty" access:"wo"`             // Image file data
	DeleteImage      *bool      `json:"deleteImage,omitempty" access:"wo"`      // Flag to delete image

}

func (w *Wishlist) IRI() string {
	return IRI(w)
}

func (w *Wishlist) GetID() string {
	return string(w.ID)
}

func (w *Wishlist) Validate() error {
	if w.Name == "" {
		return fmt.Errorf("wishlist name cannot be empty")
	}
	if w.Owner == nil || *w.Owner == "" {
		return fmt.Errorf("owner cannot be empty")
	}
	if w.Image != nil && *w.Image != "" {
		if _, err := os.Stat(*w.Image); os.IsNotExist(err) {
			return fmt.Errorf("image file does not exist: %s", *w.Image)
		}
	}
	return nil
}
