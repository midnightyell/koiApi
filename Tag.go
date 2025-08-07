package koiApi

import (
	"fmt"
	"time"
)

// Tag represents a tag in Koillection, combining fields for JSON-LD and API interactions.
type Tag struct {
	Context             *Context   `json:"@context,omitempty" access:"rw"`            // JSON-LD only
	_ID                 ID         `json:"@id,omitempty" access:"ro"`                 // JSON-LD only
	Type                string     `json:"@type,omitempty" access:"rw"`               // JSON-LD only
	ID                  ID         `json:"id,omitempty" access:"ro"`                  // Identifier
	Label               string     `json:"label" access:"rw"`                         // Tag label
	Description         *string    `json:"description,omitempty" access:"rw"`         // Tag description
	Image               *string    `json:"image,omitempty" access:"ro"`               // Image URL
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty" access:"ro"` // Small thumbnail URL
	Owner               *string    `json:"owner,omitempty" access:"ro"`               // Owner IRI
	Category            *string    `json:"category,omitempty" access:"rw"`            // Category IRI
	SeenCounter         int        `json:"seenCounter,omitempty" access:"ro"`         // View count
	Visibility          Visibility `json:"visibility,omitempty" access:"rw"`          // Visibility level
	CreatedAt           time.Time  `json:"createdAt" access:"ro"`                     // Creation timestamp
	UpdatedAt           *time.Time `json:"updatedAt,omitempty" access:"ro"`           // Update timestamp
	File                *string    `json:"file,omitempty" access:"wo"`                // Image file data

}

func (t *Tag) Summary() string {
	return fmt.Sprintf("%-40s %s", t.Label, t.ID)
}

// IRI
func (t *Tag) IRI() string {
	return fmt.Sprintf("/api/tags/%s", t.ID)
}

func (t *Tag) GetID() string {
	return string(t.ID)
}

func (t *Tag) Validate() error {
	var errs []string
	// label is required, type string; see components.schemas.Tag-tag.write.required
	if t.Label == "" {
		errs = append(errs, "tag label is required")
	}
	validateVisibility(t, &errs)
	return validationErrors(&errs)
}
