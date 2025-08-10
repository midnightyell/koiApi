package koiApi

import (
	"fmt"
	"time"
)

// Collection represents a collection in Koillection, combining fields for JSON-LD and API interactions.
type Collection struct {
	Context              Context    `json:"@context,omitempty" access:"rw"`             // JSON-LD only
	_ID                  ID         `json:"@id,omitempty" access:"ro"`                  // JSON-LD only
	Type                 string     `json:"@type,omitempty" access:"rw"`                // JSON-LD only
	ID                   ID         `json:"id,omitempty" access:"ro"`                   // Identifier
	Title                string     `json:"title" access:"rw"`                          // Collection title
	Parent               string     `json:"parent,omitempty" access:"rw"`               // Parent collection IRI
	Owner                string     `json:"owner,omitempty" access:"ro"`                // Owner IRI
	Color                string     `json:"color,omitempty" access:"ro"`                // Color code
	Image                string     `json:"image,omitempty" access:"ro"`                // Image URL
	SeenCounter          int        `json:"seenCounter,omitempty" access:"ro"`          // View count
	ItemsDefaultTemplate string     `json:"itemsDefaultTemplate,omitempty" access:"rw"` // Default template IRI
	Visibility           Visibility `json:"visibility,omitempty" access:"rw"`           // Visibility level
	ParentVisibility     string     `json:"parentVisibility,omitempty" access:"ro"`     // Parent visibility
	FinalVisibility      Visibility `json:"finalVisibility,omitempty" access:"ro"`      // Effective visibility
	ScrapedFromURL       string     `json:"scrapedFromUrl,omitempty" access:"ro"`       // Source URL
	CreatedAt            time.Time  `json:"createdAt" access:"ro"`                      // Creation timestamp
	UpdatedAt            time.Time  `json:"updatedAt,omitempty" access:"ro"`            // Update timestamp
	File                 string     `json:"file,omitempty" access:"wo"`                 // Image file data
	DeleteImage          bool       `json:"deleteImage,omitempty" access:"wo"`          // Flag to delete image
}

func (c *Collection) Summary() string {
	return fmt.Sprintf("%-40s %s", c.Title, c.ID)
}

// GetID
func (a *Collection) GetID() string {
	return string(a.ID)
}

// Validate
func (a *Collection) Validate() error {
	var errs []string
	// title is required, type string; see components.schemas.Collection-collection.write.required
	if a.Title == "" {
		errs = append(errs, "collection title is required")
	}
	// visibility enum ["public", "internal", "private"]; see components.schemas.Collection-collection.write.properties.visibility
	if a.Visibility != "" {
		switch a.Visibility {
		case VisibilityPublic, VisibilityInternal, VisibilityPrivate:
		default:
			errs = append(errs, fmt.Sprintf("invalid visibility: %s; must be public, internal, or private", a.Visibility))
		}
	}
	return validationErrors(&errs)
}

// IRI
func (a *Collection) IRI() string {
	return IRI(a)
}
