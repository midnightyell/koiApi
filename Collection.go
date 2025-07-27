package koiApi

import (
	"fmt"
	"os"
	"time"
)

// CollectionInterface defines methods for interacting with Collection resources.
type CollectionInterface interface {
	Create(client Client) (*Collection, error)                                                 // HTTP POST /api/collections
	Delete(client Client, collectionID ...ID) error                                            // HTTP DELETE /api/collections/{id}
	Get(client Client, collectionID ...ID) (*Collection, error)                                // HTTP GET /api/collections/{id}
	GetDefaultTemplate(client Client, collectionID ...ID) (*Template, error)                   // HTTP GET /api/collections/{id}/items_default_template
	GetParent(client Client, collectionID ...ID) (*Collection, error)                          // HTTP GET /api/collections/{id}/parent
	IRI() string                                                                               // /api/collections/{id}
	ListChildren(client Client, collectionID ...ID) ([]*Collection, error)                     // HTTP GET /api/collections/{id}/children
	ListCollectionData(client Client, collectionID ...ID) ([]*Datum, error)                    // HTTP GET /api/collections/{id}/data
	ListCollectionItems(client Client, collectionID ...ID) ([]*Item, error)                    // HTTP GET /api/collections/{id}/items
	Patch(client Client, collectionID ...ID) (*Collection, error)                              // HTTP PATCH /api/collections/{id}
	Update(client Client, collectionID ...ID) (*Collection, error)                             // HTTP PUT /api/collections/{id}
	UploadImage(client Client, file []byte, collectionID ...ID) (*Collection, error)           // HTTP POST /api/collections/{id}/image
	UploadImageByFile(client Client, filename string, collectionID ...ID) (*Collection, error) // HTTP POST /api/collections/{id}/image
	Summary() string
	Exists()
}

// Collection represents a collection in Koillection, combining fields for JSON-LD and API interactions.
type Collection struct {
	Context              *Context   `json:"@context,omitempty" access:"rw"`             // JSON-LD only
	_ID                  ID         `json:"@id,omitempty" access:"ro"`                  // JSON-LD only
	Type                 string     `json:"@type,omitempty" access:"rw"`                // JSON-LD only
	ID                   ID         `json:"id,omitempty" access:"ro"`                   // Identifier
	Title                string     `json:"title" access:"rw"`                          // Collection title
	Parent               *string    `json:"parent,omitempty" access:"rw"`               // Parent collection IRI
	Owner                *string    `json:"owner,omitempty" access:"ro"`                // Owner IRI
	Color                string     `json:"color,omitempty" access:"ro"`                // Color code
	Image                *string    `json:"image,omitempty" access:"ro"`                // Image URL
	SeenCounter          int        `json:"seenCounter,omitempty" access:"ro"`          // View count
	ItemsDefaultTemplate *string    `json:"itemsDefaultTemplate,omitempty" access:"rw"` // Default template IRI
	Visibility           Visibility `json:"visibility,omitempty" access:"rw"`           // Visibility level
	ParentVisibility     *string    `json:"parentVisibility,omitempty" access:"ro"`     // Parent visibility
	FinalVisibility      Visibility `json:"finalVisibility,omitempty" access:"ro"`      // Effective visibility
	ScrapedFromURL       *string    `json:"scrapedFromUrl,omitempty" access:"ro"`       // Source URL
	CreatedAt            time.Time  `json:"createdAt" access:"ro"`                      // Creation timestamp
	UpdatedAt            *time.Time `json:"updatedAt,omitempty" access:"ro"`            // Update timestamp
	File                 *string    `json:"file,omitempty" access:"wo"`                 // Image file data
	DeleteImage          *bool      `json:"deleteImage,omitempty" access:"wo"`          // Flag to delete image
}

// whichID
func (c *Collection) whichID(collectionID ...ID) ID {
	if len(collectionID) > 0 {
		return collectionID[0]
	}
	return c.ID
}

// Create
func (c *Collection) Create(client Client) (*Collection, error) {
	return client.CreateCollection(c)
}

// Delete
func (c *Collection) Delete(client Client, collectionID ...ID) error {
	id := c.whichID(collectionID...)
	return client.DeleteCollection(id)
}

// Get
func (c *Collection) Get(client Client, collectionID ...ID) (*Collection, error) {
	id := c.whichID(collectionID...)
	return client.GetCollection(id)
}

// GetDefaultTemplate
func (c *Collection) GetDefaultTemplate(client Client, collectionID ...ID) (*Template, error) {
	id := c.whichID(collectionID...)
	return client.GetCollectionDefaultTemplate(id)
}

// GetParent
func (c *Collection) GetParent(client Client, collectionID ...ID) (*Collection, error) {
	id := c.whichID(collectionID...)
	return client.GetCollectionParent(id)
}

// IRI
func (c *Collection) IRI() string {
	return fmt.Sprintf("/api/collections/%s", c.ID)
}

// Patch
func (c *Collection) Patch(client Client, collectionID ...ID) (*Collection, error) {
	id := c.whichID(collectionID...)
	return client.PatchCollection(id, c)
}

// Update
func (c *Collection) Update(client Client, collectionID ...ID) (*Collection, error) {
	id := c.whichID(collectionID...)
	return client.UpdateCollection(id, c)
}

// UploadImage
func (c *Collection) UploadImage(client Client, file []byte, collectionID ...ID) (*Collection, error) {
	id := c.whichID(collectionID...)
	return client.UploadCollectionImage(id, file)
}

// UploadImageByFile
func (c *Collection) UploadImageByFile(client Client, filename string, collectionID ...ID) (*Collection, error) {
	id := c.whichID(collectionID...)
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return c.UploadImage(client, file, id)
}
