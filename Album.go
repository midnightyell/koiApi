package koiApi

import (
	"fmt"
	"time"
)

// AlbumInterface defines methods for interacting with Album resources.
type AlbumInterface interface {
	//Create(client Client) (*Album, error)                                            // HTTP POST /api/albums
	Delete(client Client, albumID ...ID) error                                       // HTTP DELETE /api/albums/{id}
	Get(client Client, albumID ...ID) (*Album, error)                                // HTTP GET /api/albums/{id}
	GetParent(client Client, albumID ...ID) (*Album, error)                          // HTTP GET /api/albums/{id}/parent
	IRI() string                                                                     // /api/albums/{id}
	List(client Client) ([]*Album, error)                                            // HTTP GET /api/albums
	ListChildren(client Client, albumID ...ID) ([]*Album, error)                     // HTTP GET /api/albums/{id}/children
	ListPhotos(client Client, albumID ...ID) ([]*Photo, error)                       // HTTP GET /api/albums/{id}/photos
	Patch(client Client, albumID ...ID) (*Album, error)                              // HTTP PATCH /api/albums/{id}
	Update(client Client, albumID ...ID) (*Album, error)                             // HTTP PUT /api/albums/{id}
	UploadImage(client Client, file []byte, albumID ...ID) (*Album, error)           // HTTP POST /api/albums/{id}/image
	UploadImageByFile(client Client, filename string, albumID ...ID) (*Album, error) // HTTP POST /api/albums/{id}/image
	Summary() string
}

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

// whichID
func (a *Album) whichID(albumID ...ID) ID {
	if len(albumID) > 0 {
		return albumID[0]
	}
	return a.ID
}

func (a *Album) GetID() string {
	return string(a.ID)
}

// Create
func (a *Album) Create() (*Album, error) {
	return Create(a)
}

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
	return Get(a)
}

// GetParent
func (a *Album) GetParent() (*Album, error) {
	return GetParent(a)
}

// IRI
func (a *Album) IRI() string {
	return fmt.Sprintf("/api/albums/%s", a.ID)
}

// List
func (a *Album) List() ([]*Album, error) {
	return List(a)
}

// ListChildren
func (a *Album) ListChildren() ([]*Album, error) {
	return ListChildren(a)
}

// ListPhotos
func (a *Album) ListPhotos() ([]*Photo, error) {
	return ListPhotos(a)
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
	return UploadImage(a, file)
}

// UploadImageFromFile
func (a *Album) UploadImageFromFile(filename string) (*Album, error) {
	return UploadImageFromFile(a, filename)
}
