package koiApi

import (
	"fmt"
	"time"
)

type PhotoImage *string

// Photo represents a photo in Koillection, combining fields for JSON-LD and API interactions.
type Photo struct {
	Context             *Context   `json:"@context,omitempty" access:"rw"`            // JSON-LD only
	_ID                 ID         `json:"@id,omitempty" access:"ro"`                 // JSON-LD only
	Type                string     `json:"@type,omitempty" access:"rw"`               // JSON-LD only
	ID                  ID         `json:"id,omitempty" access:"ro"`                  // Identifier
	Title               string     `json:"title" access:"rw"`                         // Photo title
	Comment             *string    `json:"comment,omitempty" access:"rw"`             // Photo comment
	Place               *string    `json:"place,omitempty" access:"rw"`               // Photo location
	Album               *string    `json:"album" access:"rw"`                         // Album IRI
	Owner               *string    `json:"owner,omitempty" access:"ro"`               // Owner IRI
	Image               *string    `json:"image,omitempty" access:"ro"`               // Image URL
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty" access:"ro"` // Small thumbnail URL
	TakenAt             *time.Time `json:"takenAt,omitempty" access:"ro"`             // Date taken
	Visibility          Visibility `json:"visibility,omitempty" access:"rw"`          // Visibility level
	ParentVisibility    *string    `json:"parentVisibility,omitempty" access:"ro"`    // Parent visibility
	FinalVisibility     Visibility `json:"finalVisibility,omitempty" access:"ro"`     // Effective visibility
	CreatedAt           time.Time  `json:"createdAt" access:"ro"`                     // Creation timestamp
	UpdatedAt           *time.Time `json:"updatedAt,omitempty" access:"ro"`           // Update timestamp
	File                *string    `json:"file,omitempty" access:"wo"`                // Image file data

}

// IRI
func (p *Photo) IRI() string {
	return fmt.Sprintf("/api/photos/%s", p.ID)
}

func (p *Photo) GetID() string {
	return string(p.ID)
}

func (p *Photo) Validate() error {
	return nil
}
