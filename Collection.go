package koiApi

import (
	"time"
)

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

// GetID
func (a *Collection) GetID() string {
	return string(a.ID)
}

// Validate
func (a *Collection) Validate() error {
	return nil
}

// Create
func (a *Collection) Create() (*Collection, error) {
	return Create(a)
}

// Delete
func (a *Collection) Delete() error {
	return Delete(a)
}

// Get
func (a *Collection) Get() (*Collection, error) {
	res, err := Get(a)
	return res.(*Collection), err
}

// GetDefaultTemplate
func (a *Collection) GetDefaultTemplate() (*Template, error) {
	res, err := Get(a)
	return res.(*Template), err
}

// GetParent
func (a *Collection) GetParent() (*Collection, error) {
	res, err := Get(a)
	return res.(*Collection), err
}

// IRI
func (a *Collection) IRI() string {
	return IRI(a)
}

// List
func (a *Collection) List() ([]*Collection, error) {
	res, err := List(a)
	return res.([]*Collection), err
}

// ListChildren
func (a *Collection) ListChildren() ([]*Collection, error) {
	res, err := List(a)
	return res.([]*Collection), err
}

// ListItems
func (a *Collection) ListItems() ([]*Item, error) {
	res, err := List(a)
	return res.([]*Item), err
}

// ListData
func (a *Collection) ListData() ([]*Datum, error) {
	res, err := List(a)
	return res.([]*Datum), err
}

// Patch
func (a *Collection) Patch() (*Collection, error) {
	return Patch(a)
}

// Update
func (a *Collection) Update() (*Collection, error) {
	return Update(a)
}

// UploadImage
func (a *Collection) UploadImage(file []byte) (*Collection, error) {
	return Upload(a, file)
}

// UploadImageFromFile
func (a *Collection) UploadFromFile(filename string) (*Collection, error) {
	return UploadFromFile(a, filename)
}
