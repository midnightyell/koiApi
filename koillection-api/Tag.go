package koiApi

import (
	"context"
	"fmt"
	"os"
	"time"
)

// TagInterface defines methods for interacting with Tag resources.
type TagInterface interface {
	Create(ctx context.Context, client Client) (*Tag, error)                                          // HTTP POST /api/tags
	Delete(ctx context.Context, client Client, tagID ...ID) error                                     // HTTP DELETE /api/tags/{id}
	Get(ctx context.Context, client Client, tagID ...ID) (*Tag, error)                                // HTTP GET /api/tags/{id}
	GetCategory(ctx context.Context, client Client, tagID ...ID) (*TagCategory, error)                // HTTP GET /api/tags/{id}/category
	IRI() string                                                                                      // /api/tags/{id}
	List(ctx context.Context, client Client) ([]*Tag, error)                                          // HTTP GET /api/tags
	ListItems(ctx context.Context, client Client, tagID ...ID) ([]*Item, error)                       // HTTP GET /api/tags/{id}/items
	Patch(ctx context.Context, client Client, tagID ...ID) (*Tag, error)                              // HTTP PATCH /api/tags/{id}
	Update(ctx context.Context, client Client, tagID ...ID) (*Tag, error)                             // HTTP PUT /api/tags/{id}
	UploadImage(ctx context.Context, client Client, file []byte, tagID ...ID) (*Tag, error)           // HTTP POST /api/tags/{id}/image
	UploadImageByFile(ctx context.Context, client Client, filename string, tagID ...ID) (*Tag, error) // HTTP POST /api/tags/{id}/image
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
func (t *Tag) Create(ctx context.Context, client Client) (*Tag, error) {
	return client.CreateTag(ctx, t)
}

// Delete
func (t *Tag) Delete(ctx context.Context, client Client, tagID ...ID) error {
	id := t.whichID(tagID...)
	return client.DeleteTag(ctx, id)
}

// Get
func (t *Tag) Get(ctx context.Context, client Client, tagID ...ID) (*Tag, error) {
	id := t.whichID(tagID...)
	return client.GetTag(ctx, id)
}

// GetCategory
func (t *Tag) GetCategory(ctx context.Context, client Client, tagID ...ID) (*TagCategory, error) {
	id := t.whichID(tagID...)
	return client.GetCategoryOfTag(ctx, id)
}

// IRI
func (t *Tag) IRI() string {
	return fmt.Sprintf("/api/tags/%s", t.ID)
}

// List
func (t *Tag) List(ctx context.Context, client Client) ([]*Tag, error) {
	var allTags []*Tag
	for page := 1; ; page++ {
		tags, err := client.ListTags(ctx, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list tags on page %d: %w", page, err)
		}
		if len(tags) == 0 {
			break
		}
		allTags = append(allTags, tags...)
	}
	return allTags, nil
}

// ListItems
func (t *Tag) ListItems(ctx context.Context, client Client, tagID ...ID) ([]*Item, error) {
	id := t.whichID(tagID...)
	var allItems []*Item
	for page := 1; ; page++ {
		items, err := client.ListTagItems(ctx, id, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list items for ID %s on page %d: %w", id, page, err)
		}
		if len(items) == 0 {
			break
		}
		allItems = append(allItems, items...)
	}
	return allItems, nil
}

// Patch
func (t *Tag) Patch(ctx context.Context, client Client, tagID ...ID) (*Tag, error) {
	id := t.whichID(tagID...)
	return client.PatchTag(ctx, id, t)
}

// Update
func (t *Tag) Update(ctx context.Context, client Client, tagID ...ID) (*Tag, error) {
	id := t.whichID(tagID...)
	return client.UpdateTag(ctx, id, t)
}

// UploadImage
func (t *Tag) UploadImage(ctx context.Context, client Client, file []byte, tagID ...ID) (*Tag, error) {
	id := t.whichID(tagID...)
	return client.UploadTagImage(ctx, id, file)
}

// UploadImageByFile
func (t *Tag) UploadImageByFile(ctx context.Context, client Client, filename string, tagID ...ID) (*Tag, error) {
	id := t.whichID(tagID...)
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return client.UploadTagImage(ctx, id, file)
}
