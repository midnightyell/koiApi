package koiApi

import (
	"fmt"
	"os"
	"time"
)

// AlbumInterface defines methods for interacting with Album resources.
type AlbumInterface interface {
	Create(client Client) (*Album, error)                                            // HTTP POST /api/albums
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

func (obj *Album) Create() (interface{}, error) {
	if err := obj.Validate(); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}
	koiOp := KoiPathForOp()

	pc, _, _, _ := runtime.Caller(0)
    return runtime.FuncForPC(pc).Name()

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

// 
func (a *Album) KoiPath() string {


// Create
func (a *Album) Create(client Client) (*Album, error) {
	return client.CreateAlbum(a)
}

// Delete
func (a *Album) Delete(client Client, albumID ...ID) error {
	id := a.whichID(albumID...)
	return client.DeleteAlbum(id)
}

// Get
func (a *Album) Get(client Client, albumID ...ID) (*Album, error) {
	id := a.whichID(albumID...)
	return client.GetAlbum(id)
}

// GetParent
func (a *Album) GetParent(client Client, albumID ...ID) (*Album, error) {
	id := a.whichID(albumID...)
	return client.GetAlbumParent(id)
}

// IRI
func (a *Album) IRI() string {
	return fmt.Sprintf("/api/albums/%s", a.ID)
}

// List
func (a *Album) List(client Client) ([]*Album, error) {
	var allAlbums []*Album
	for page := 1; ; page++ {
		albums, err := client.ListAlbums()
		if err != nil {
			return nil, fmt.Errorf("failed to list albums on page %d: %w", err)
		}
		if len(albums) == 0 {
			break
		}
		allAlbums = append(allAlbums, albums...)
	}
	return allAlbums, nil
}

// ListChildren
func (a *Album) ListChildren(client Client, albumID ...ID) ([]*Album, error) {
	id := a.whichID(albumID...)
	var allChildren []*Album
	for page := 1; ; page++ {
		children, err := client.ListAlbumChildren(id)
		if err != nil {
			return nil, fmt.Errorf("failed to list child albums for ID %s on page %d: %w", id, err)
		}
		if len(children) == 0 {
			break
		}
		allChildren = append(allChildren, children...)
	}
	return allChildren, nil
}

// ListPhotos
func (a *Album) ListPhotos(client Client, albumID ...ID) ([]*Photo, error) {
	id := a.whichID(albumID...)
	var allPhotos []*Photo
	for page := 1; ; page++ {
		photos, err := client.ListAlbumPhotos(id)
		if err != nil {
			return nil, fmt.Errorf("failed to list photos for ID %s on page %d: %w", id, err)
		}
		if len(photos) == 0 {
			break
		}
		allPhotos = append(allPhotos, photos...)
	}
	return allPhotos, nil
}

// Patch
func (a *Album) Patch(client Client, albumID ...ID) (*Album, error) {
	id := a.whichID(albumID...)
	return client.PatchAlbum(id, a)
}

// Update
func (a *Album) Update(client Client, albumID ...ID) (*Album, error) {
	id := a.whichID(albumID...)
	return client.UpdateAlbum(id, a)
}

// UploadImage
func (a *Album) UploadImage(client Client, file []byte, albumID ...ID) (*Album, error) {
	id := a.whichID(albumID...)
	return client.UploadAlbumImage(id, file)
}

// UploadImageByFile
func (a *Album) UploadImageByFile(client Client, filename string, albumID ...ID) (*Album, error) {
	id := a.whichID(albumID...)
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return client.UploadAlbumImage(id, file)
}
