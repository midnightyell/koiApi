package koiApi

import (
	"fmt"
	"os"
	"time"
)

type PhotoImage *string

// PhotoInterface defines methods for interacting with Photo resources.
type PhotoInterface interface {
	Create(client Client) (*Photo, error)                                            // HTTP POST /api/photos
	Delete(client Client, photoID ...ID) error                                       // HTTP DELETE /api/photos/{id}
	Get(client Client, photoID ...ID) (*Photo, error)                                // HTTP GET /api/photos/{id}
	GetAlbum(client Client, photoID ...ID) (*Album, error)                           // HTTP GET /api/photos/{id}/album
	IRI() string                                                                     // /api/photos/{id}
	List(client Client) ([]*Photo, error)                                            // HTTP GET /api/photos
	Patch(client Client, photoID ...ID) (*Photo, error)                              // HTTP PATCH /api/photos/{id}
	Update(client Client, photoID ...ID) (*Photo, error)                             // HTTP PUT /api/photos/{id}
	UploadImage(client Client, file []byte, photoID ...ID) (*Photo, error)           // HTTP POST /api/photos/{id}/image
	UploadImageByFile(client Client, filename string, photoID ...ID) (*Photo, error) // HTTP POST /api/photos/{id}/image
	Summary() string
}

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

// whichID
func (p *Photo) whichID(photoID ...ID) ID {
	if len(photoID) > 0 {
		return photoID[0]
	}
	return p.ID
}

// Create
func (p *Photo) Create(client Client) (*Photo, error) {
	return client.CreatePhoto(p)
}

// Delete
func (p *Photo) Delete(client Client, photoID ...ID) error {
	id := p.whichID(photoID...)
	return client.DeletePhoto(id)
}

// Get
func (p *Photo) Get(client Client, photoID ...ID) (*Photo, error) {
	id := p.whichID(photoID...)
	return client.GetPhoto(id)
}

// GetAlbum
func (p *Photo) GetAlbum(client Client, photoID ...ID) (*Album, error) {
	id := p.whichID(photoID...)
	return client.GetPhotoAlbum(id)
}

// IRI
func (p *Photo) IRI() string {
	return fmt.Sprintf("/api/photos/%s", p.ID)
}

// Patch
func (p *Photo) Patch(client Client, photoID ...ID) (*Photo, error) {
	id := p.whichID(photoID...)
	return client.PatchPhoto(id, p)
}

// Update
func (p *Photo) Update(client Client, photoID ...ID) (*Photo, error) {
	id := p.whichID(photoID...)
	return client.UpdatePhoto(id, p)
}

// UploadImage
func (p *Photo) UploadImage(client Client, file []byte, photoID ...ID) (*Photo, error) {
	id := p.whichID(photoID...)
	return client.UploadPhotoImage(id, file)
}

// UploadImageByFile
func (p *Photo) UploadImageByFile(client Client, filename string, photoID ...ID) (*Photo, error) {
	id := p.whichID(photoID...)
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return client.UploadPhotoImage(id, file)
}
