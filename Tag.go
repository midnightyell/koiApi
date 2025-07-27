package koiApi

import (
	"fmt"
	"os"
	"time"
)

// TagInterface defines methods for interacting with Tag resources.
type TagInterface interface {
	Create(client Client) (*Tag, error)                                          // HTTP POST /api/tags
	Delete(client Client, tagID ...ID) error                                     // HTTP DELETE /api/tags/{id}
	Get(client Client, tagID ...ID) (*Tag, error)                                // HTTP GET /api/tags/{id}
	GetCategory(client Client, tagID ...ID) (*TagCategory, error)                // HTTP GET /api/tags/{id}/category
	IRI() string                                                                 // /api/tags/{id}
	List(client Client) ([]*Tag, error)                                          // HTTP GET /api/tags
	ListItems(client Client, tagID ...ID) ([]*Item, error)                       // HTTP GET /api/tags/{id}/items
	Patch(client Client, tagID ...ID) (*Tag, error)                              // HTTP PATCH /api/tags/{id}
	Update(client Client, tagID ...ID) (*Tag, error)                             // HTTP PUT /api/tags/{id}
	UploadImage(client Client, file []byte, tagID ...ID) (*Tag, error)           // HTTP POST /api/tags/{id}/image
	UploadImageByFile(client Client, filename string, tagID ...ID) (*Tag, error) // HTTP POST /api/tags/{id}/image
	Summary() string
}

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

// whichID
func (t *Tag) whichID(tagID ...ID) ID {
	if len(tagID) > 0 {
		return tagID[0]
	}
	return t.ID
}

// Create
func (t *Tag) Create(client Client) (*Tag, error) {
	return client.CreateTag(t)
}

// Delete
func (t *Tag) Delete(client Client, tagID ...ID) error {
	id := t.whichID(tagID...)
	return client.DeleteTag(id)
}

// Get
func (t *Tag) Get(client Client, tagID ...ID) (*Tag, error) {
	id := t.whichID(tagID...)
	return client.GetTag(id)
}

// GetCategory
func (t *Tag) GetCategory(client Client, tagID ...ID) (*TagCategory, error) {
	id := t.whichID(tagID...)
	return client.GetCategoryOfTag(id)
}

// IRI
func (t *Tag) IRI() string {
	return fmt.Sprintf("/api/tags/%s", t.ID)
}

// List
func (t *Tag) List(client Client) ([]*Tag, error) {
	return client.ListTags()
}

// ListItems
func (t *Tag) ListItems(client Client, tagID ...ID) ([]*Item, error) {
	id := t.whichID(tagID...)
	return client.ListTagItems(id)
}

// Patch
func (t *Tag) Patch(client Client, tagID ...ID) (*Tag, error) {
	id := t.whichID(tagID...)
	return client.PatchTag(id, t)
}

// Update
func (t *Tag) Update(client Client, tagID ...ID) (*Tag, error) {
	id := t.whichID(tagID...)
	return client.UpdateTag(id, t)
}

// UploadImage
func (t *Tag) UploadImage(client Client, file []byte, tagID ...ID) (*Tag, error) {
	id := t.whichID(tagID...)
	return client.UploadTagImage(id, file)
}

// UploadImageByFile
func (t *Tag) UploadImageByFile(client Client, filename string, tagID ...ID) (*Tag, error) {
	id := t.whichID(tagID...)
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return client.UploadTagImage(id, file)
}
