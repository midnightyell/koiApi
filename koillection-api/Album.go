package koiApi

import (
	"context"
	"fmt"
	"os"
	"time"
)

// AlbumInterface defines methods for interacting with Album resources.
type AlbumInterface interface {
	Create(ctx context.Context, client Client) (*Album, error)                                            // HTTP POST /api/albums
	Delete(ctx context.Context, client Client, albumID ...ID) error                                       // HTTP DELETE /api/albums/{id}
	Get(ctx context.Context, client Client, albumID ...ID) (*Album, error)                                // HTTP GET /api/albums/{id}
	GetParent(ctx context.Context, client Client, albumID ...ID) (*Album, error)                          // HTTP GET /api/albums/{id}/parent
	IRI() string                                                                                          // /api/albums/{id}
	List(ctx context.Context, client Client) ([]*Album, error)                                            // HTTP GET /api/albums
	ListChildren(ctx context.Context, client Client, albumID ...ID) ([]*Album, error)                     // HTTP GET /api/albums/{id}/children
	ListPhotos(ctx context.Context, client Client, albumID ...ID) ([]*Photo, error)                       // HTTP GET /api/albums/{id}/photos
	Patch(ctx context.Context, client Client, albumID ...ID) (*Album, error)                              // HTTP PATCH /api/albums/{id}
	Update(ctx context.Context, client Client, albumID ...ID) (*Album, error)                             // HTTP PUT /api/albums/{id}
	UploadImage(ctx context.Context, client Client, file []byte, albumID ...ID) (*Album, error)           // HTTP POST /api/albums/{id}/image
	UploadImageByFile(ctx context.Context, client Client, filename string, albumID ...ID) (*Album, error) // HTTP POST /api/albums/{id}/image
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

// Create
func (a *Album) Create(ctx context.Context, client Client) (*Album, error) {
	return client.CreateAlbum(ctx, a)
}

// Delete
func (a *Album) Delete(ctx context.Context, client Client, albumID ...ID) error {
	id := a.whichID(albumID...)
	return client.DeleteAlbum(ctx, id)
}

// Get
func (a *Album) Get(ctx context.Context, client Client, albumID ...ID) (*Album, error) {
	id := a.whichID(albumID...)
	return client.GetAlbum(ctx, id)
}

// GetParent
func (a *Album) GetParent(ctx context.Context, client Client, albumID ...ID) (*Album, error) {
	id := a.whichID(albumID...)
	return client.GetAlbumParent(ctx, id)
}

// IRI
func (a *Album) IRI() string {
	return fmt.Sprintf("/api/albums/%s", a.ID)
}

// List
func (a *Album) List(ctx context.Context, client Client) ([]*Album, error) {
	var allAlbums []*Album
	for page := 1; ; page++ {
		albums, err := client.ListAlbums(ctx, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list albums on page %d: %w", page, err)
		}
		if len(albums) == 0 {
			break
		}
		allAlbums = append(allAlbums, albums...)
	}
	return allAlbums, nil
}

// ListChildren
func (a *Album) ListChildren(ctx context.Context, client Client, albumID ...ID) ([]*Album, error) {
	id := a.whichID(albumID...)
	var allChildren []*Album
	for page := 1; ; page++ {
		children, err := client.ListAlbumChildren(ctx, id, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list child albums for ID %s on page %d: %w", id, page, err)
		}
		if len(children) == 0 {
			break
		}
		allChildren = append(allChildren, children...)
	}
	return allChildren, nil
}

// ListPhotos
func (a *Album) ListPhotos(ctx context.Context, client Client, albumID ...ID) ([]*Photo, error) {
	id := a.whichID(albumID...)
	var allPhotos []*Photo
	for page := 1; ; page++ {
		photos, err := client.ListAlbumPhotos(ctx, id, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list photos for ID %s on page %d: %w", id, page, err)
		}
		if len(photos) == 0 {
			break
		}
		allPhotos = append(allPhotos, photos...)
	}
	return allPhotos, nil
}

// Patch
func (a *Album) Patch(ctx context.Context, client Client, albumID ...ID) (*Album, error) {
	id := a.whichID(albumID...)
	return client.PatchAlbum(ctx, id, a)
}

// Update
func (a *Album) Update(ctx context.Context, client Client, albumID ...ID) (*Album, error) {
	id := a.whichID(albumID...)
	return client.UpdateAlbum(ctx, id, a)
}

// UploadImage
func (a *Album) UploadImage(ctx context.Context, client Client, file []byte, albumID ...ID) (*Album, error) {
	id := a.whichID(albumID...)
	return client.UploadAlbumImage(ctx, id, file)
}

// UploadImageByFile
func (a *Album) UploadImageByFile(ctx context.Context, client Client, filename string, albumID ...ID) (*Album, error) {
	id := a.whichID(albumID...)
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return client.UploadAlbumImage(ctx, id, file)
}
