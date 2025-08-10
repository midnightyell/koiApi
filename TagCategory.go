package koiApi

import (
	"fmt"
	"time"
)

// TagCategory represents a tag category in Koillection, combining fields for JSON-LD and API interactions.
type TagCategory struct {
	Context     Context   `json:"@context,omitempty" access:"rw"`    // JSON-LD only
	_ID         ID        `json:"@id,omitempty" access:"ro"`         // JSON-LD only
	Type        string    `json:"@type,omitempty" access:"rw"`       // JSON-LD only
	ID          ID        `json:"id,omitempty" access:"ro"`          // Identifier
	Label       string    `json:"label" access:"rw"`                 // Category label
	Description string    `json:"description,omitempty" access:"rw"` // Category description
	Color       string    `json:"color" access:"rw"`                 // Color code
	Owner       string    `json:"owner,omitempty" access:"ro"`       // Owner IRI
	CreatedAt   time.Time `json:"createdAt" access:"ro"`             // Creation timestamp
	UpdatedAt   time.Time `json:"updatedAt,omitempty" access:"ro"`   // Update timestamp

}

func (tc *TagCategory) Summary() string {
	return fmt.Sprintf("%-40s %s", tc.Label, tc.ID)
}

// IRI
func (tc *TagCategory) IRI() string {
	return fmt.Sprintf("/api/tag_categories/%s", tc.ID)
}

func (tc *TagCategory) GetID() string {
	return string(tc.ID)
}

func (tc *TagCategory) Validate() error {
	var errs []string
	// label is required, type string; see components.schemas.TagCategory-tagCategory.write.required
	if tc.Label == "" {
		errs = append(errs, "tag category label is required")
	}
	validateVisibility(tc, &errs)
	return validationErrors(&errs)

}
