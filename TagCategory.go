package koiApi

import (
	"fmt"
	"time"
)

// TagCategoryInterface defines methods for interacting with TagCategory resources.
type TagCategoryInterface interface {
	Create(client Client) (*TagCategory, error)                      // HTTP POST /api/tag_categories
	Delete(client Client, tagCategoryID ...ID) error                 // HTTP DELETE /api/tag_categories/{id}
	Get(client Client, tagCategoryID ...ID) (*TagCategory, error)    // HTTP GET /api/tag_categories/{id}
	IRI() string                                                     // /api/tag_categories/{id}
	List(client Client) ([]*TagCategory, error)                      // HTTP GET /api/tag_categories
	ListTags(client Client, tagCategoryID ...ID) ([]*Tag, error)     // HTTP GET /api/tag_categories/{id}/tags
	Patch(client Client, tagCategoryID ...ID) (*TagCategory, error)  // HTTP PATCH /api/tag_categories/{id}
	Update(client Client, tagCategoryID ...ID) (*TagCategory, error) // HTTP PUT /api/tag_categories/{id}
	Summary() string
}

// TagCategory represents a tag category in Koillection, combining fields for JSON-LD and API interactions.
type TagCategory struct {
	Context     *Context   `json:"@context,omitempty" access:"rw"`    // JSON-LD only
	_ID         ID         `json:"@id,omitempty" access:"ro"`         // JSON-LD only
	Type        string     `json:"@type,omitempty" access:"rw"`       // JSON-LD only
	ID          ID         `json:"id,omitempty" access:"ro"`          // Identifier
	Label       string     `json:"label" access:"rw"`                 // Category label
	Description *string    `json:"description,omitempty" access:"rw"` // Category description
	Color       string     `json:"color" access:"rw"`                 // Color code
	Owner       *string    `json:"owner,omitempty" access:"ro"`       // Owner IRI
	CreatedAt   time.Time  `json:"createdAt" access:"ro"`             // Creation timestamp
	UpdatedAt   *time.Time `json:"updatedAt,omitempty" access:"ro"`   // Update timestamp

}

// whichID
func (tc *TagCategory) whichID(tagCategoryID ...ID) ID {
	if len(tagCategoryID) > 0 {
		return tagCategoryID[0]
	}
	return tc.ID
}

// Create
func (tc *TagCategory) Create(client Client) (*TagCategory, error) {
	return client.CreateTagCategory(tc)
}

// Delete
func (tc *TagCategory) Delete(client Client, tagCategoryID ...ID) error {
	id := tc.whichID(tagCategoryID...)
	return client.DeleteTagCategory(id)
}

// Get
func (tc *TagCategory) Get(client Client, tagCategoryID ...ID) (*TagCategory, error) {
	id := tc.whichID(tagCategoryID...)
	return client.GetTagCategory(id)
}

// IRI
func (tc *TagCategory) IRI() string {
	return fmt.Sprintf("/api/tag_categories/%s", tc.ID)
}

// List
func (tc *TagCategory) List(client Client) ([]*TagCategory, error) {
	return client.ListTagCategories()
}

// ListTags
func (tc *TagCategory) ListTags(client Client, tagCategoryID ...ID) ([]*Tag, error) {
	id := tc.whichID(tagCategoryID...)
	return client.ListTagCategoryTags(id)
}

// Patch
func (tc *TagCategory) Patch(client Client, tagCategoryID ...ID) (*TagCategory, error) {
	id := tc.whichID(tagCategoryID...)
	return client.PatchTagCategory(id, tc)
}

// Update
func (tc *TagCategory) Update(client Client, tagCategoryID ...ID) (*TagCategory, error) {
	id := tc.whichID(tagCategoryID...)
	return client.UpdateTagCategory(id, tc)
}
