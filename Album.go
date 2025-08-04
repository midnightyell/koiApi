package koiApi

import (
	"fmt"
	"time"
)

// Album represents an album in Koillection, combining fields for JSON-LD and API interactions.
type Album struct {
	Context          *Context   `json:"@context,omitempty" access:"rw"`         // JSON-LD only
	_ID              ID         `json:"@id,omitempty" access:"ro"`              // JSON-LD only
	Type             string     `json:"@type,omitempty" access:"rw"`            // JSON-LD only
	ID               ID         `json:"id,omitempty" access:"ro"`               // Identifier
	Title            string     `json:"title" access:"rw"`                      // Album title
	Color            string     `json:"color,omitempty" access:"ro"`            // Color code
	Image            *string    `json:"image,omitempty" access:"ro"`            // Image URL
	Owner            *string    `json:"owner,omitempty" access:"ro"`            // Owner IRI
	Parent           *string    `json:"parent,omitempty" access:"rw"`           // Parent album IRI
	SeenCounter      int        `json:"seenCounter,omitempty" access:"ro"`      // View count
	Visibility       Visibility `json:"visibility,omitempty" access:"rw"`       // Visibility level
	ParentVisibility *string    `json:"parentVisibility,omitempty" access:"ro"` // Parent visibility
	FinalVisibility  Visibility `json:"finalVisibility,omitempty" access:"ro"`  // Effective visibility
	CreatedAt        time.Time  `json:"createdAt" access:"ro"`                  // Creation timestamp
	UpdatedAt        *time.Time `json:"updatedAt,omitempty" access:"ro"`        // Update timestamp
	File             *string    `json:"file,omitempty" access:"wo"`             // Image file data
	DeleteImage      *bool      `json:"deleteImage,omitempty" access:"wo"`      // Flag to delete image
}

// GetID
func (a *Album) GetID() string {
	return string(a.ID)
}

// Create
func (a *Album) Create() (*Album, error) {
	return Create(a)
}

// Validate
func (a *Album) Validate() error {
	if a.Title == "" {
		return fmt.Errorf("album title cannot be empty")
	}
	return nil
}

// Delete
func (a *Album) Delete() error {
	return Delete(a)
}

// Get
func (a *Album) Get() (*Album, error) {
	res, err := Get(a)
	return res.(*Album), err
}

// GetParent
func (a *Album) GetParent() (*Album, error) {
	res, err := Get(a)
	return res.(*Album), err
}

// IRI
func (a *Album) IRI() string {
	return IRI(a)
}

// List
func (a *Album) List() ([]*Album, error) {
	res, err := List(a)
	return res.([]*Album), err
}

// ListChildren
func (a *Album) ListChildren() ([]*Album, error) {
	res, err := List(a)
	return res.([]*Album), err
}

// ListPhotos
func (a *Album) ListPhotos() ([]*Photo, error) {
	res, err := List(a)
	return res.([]*Photo), err
}

// Patch
func (a *Album) Patch() (*Album, error) {
	return Patch(a)
}

// Update
func (a *Album) Update() (*Album, error) {
	return Update(a)
}

// UploadImage
func (a *Album) UploadImage(file []byte) (*Album, error) {
	return Upload(a, file)
}

// UploadImageFromFile
func (a *Album) UploadImageFromFile(filename string) (*Album, error) {
	return UploadFromFile(a, filename)
}
