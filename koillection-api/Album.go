package koiApi

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// AlbumInterface defines methods for interacting with Album resources.
type AlbumInterface interface {
	Create(ctx context.Context, client Client) (*Album, error)
	Get(ctx context.Context, client Client, id ID) (*Album, error)
	List(ctx context.Context, client Client) ([]*Album, error)
	Update(ctx context.Context, client Client, id ID) (*Album, error)
	Patch(ctx context.Context, client Client, id ID) (*Album, error)
	Delete(ctx context.Context, client Client, id ID) error
	ListChildren(ctx context.Context, client Client, id ID) ([]*Album, error)
	UploadImage(ctx context.Context, client Client, id ID, file []byte) (*Album, error)
	GetParent(ctx context.Context, client Client, id ID) (*Album, error)
	ListPhotos(ctx context.Context, client Client, id ID) ([]*Photo, error)
	Validate(ctx context.Context, client Client) error
}

// Album represents an album in Koillection, combining read and write fields.
type Album struct {
	Context          *Context   `json:"@context,omitempty"`         // JSON-LD only
	ID_              ID         `json:"@id,omitempty"`              // JSON-LD only (maps to "@id" in JSON, read-only)
	ID               ID         `json:"id,omitempty"`               // JSON-LD only (maps to "id" in JSON, read-only)
	Type             string     `json:"@type,omitempty"`            // JSON-LD only
	Title            string     `json:"title"`                      // Read and write
	Color            string     `json:"color,omitempty"`            // Read-only
	Image            *string    `json:"image,omitempty"`            // Read-only
	Owner            *string    `json:"owner,omitempty"`            // Read-only, IRI
	Parent           *string    `json:"parent,omitempty"`           // Read and write, IRI
	SeenCounter      int        `json:"seenCounter,omitempty"`      // Read-only
	Visibility       Visibility `json:"visibility,omitempty"`       // Read and write
	ParentVisibility *string    `json:"parentVisibility,omitempty"` // Read-only
	FinalVisibility  Visibility `json:"finalVisibility,omitempty"`  // Read-only
	CreatedAt        time.Time  `json:"createdAt"`                  // Read-only
	UpdatedAt        *time.Time `json:"updatedAt,omitempty"`        // Read-only
	File             *string    `json:"file,omitempty"`             // Write-only, binary data via multipart form
	DeleteImage      *bool      `json:"deleteImage,omitempty"`      // Write-only
}

// Validate checks the Album's fields for validity, using ctx for cancellation and client for optional IRI validation.
func (a *Album) Validate(ctx context.Context, client Client) error {
	// Check for context cancellation
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("validation cancelled: %w", err)
	}

	// Required fields
	if a.Title == "" {
		return fmt.Errorf("title must not be empty")
	}

	// Visibility must be a valid value
	switch a.Visibility {
	case VisibilityPublic, VisibilityInternal, VisibilityPrivate, "":
		// Valid or unset (server may set default)
	default:
		return fmt.Errorf("invalid visibility value: %s", a.Visibility)
	}

	// Optional fields
	if a.Parent != nil {
		if *a.Parent == "" {
			return fmt.Errorf("parent IRI must not be empty if set")
		}
		if !strings.HasPrefix(*a.Parent, "/api/albums/") {
			return fmt.Errorf("parent IRI must start with /api/albums/: %s", *a.Parent)
		}
		// Optionally validate Parent exists if client is provided
		if client != nil {
			parts := strings.Split(*a.Parent, "/")
			if len(parts) < 4 {
				return fmt.Errorf("invalid parent IRI format: %s", *a.Parent)
			}
			parentID := ID(parts[3])
			_, err := client.GetAlbum(ctx, parentID)
			if err != nil {
				return fmt.Errorf("invalid parent album %s: %w", *a.Parent, err)
			}
		}
	}

	if a.File != nil && *a.File == "" {
		return fmt.Errorf("file must not be empty if set")
	}

	// Read-only fields for creation vs. update
	if a.ID == "" && a.ID_ == "" {
		// Creation: read-only fields should be empty
		if a.ID_ != "" {
			return fmt.Errorf("ID_ must be empty for creation")
		}
		if a.ID != "" {
			return fmt.Errorf("ID must be empty for creation")
		}
		if a.Type != "" && a.Type != "Album" {
			return fmt.Errorf("Type must be empty or 'Album' for creation: %s", a.Type)
		}
	} else {
		// Update: ID should be non-empty
		if a.ID == "" {
			return fmt.Errorf("ID must not be empty for update")
		}
	}

	return nil
}

// Create calls Client.CreateAlbum to create a new Album.
func (a *Album) Create(ctx context.Context, client Client) (*Album, error) {
	return client.CreateAlbum(ctx, a)
}

// Get retrieves an Album by ID using Client.GetAlbum.
func (a *Album) Get(ctx context.Context, client Client, id ID) (*Album, error) {
	return client.GetAlbum(ctx, id)
}

// List retrieves all Albums across all pages using Client.ListAlbums.
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

// Update updates an Album by ID using Client.UpdateAlbum.
func (a *Album) Update(ctx context.Context, client Client, id ID) (*Album, error) {
	return client.UpdateAlbum(ctx, id, a)
}

// Patch partially updates an Album by ID using Client.PatchAlbum.
func (a *Album) Patch(ctx context.Context, client Client, id ID) (*Album, error) {
	return client.PatchAlbum(ctx, id, a)
}

// Delete removes an Album by ID using Client.DeleteAlbum.
func (a *Album) Delete(ctx context.Context, client Client, id ID) error {
	return client.DeleteAlbum(ctx, id)
}

// ListChildren retrieves all child Albums for the given ID across all pages using Client.ListAlbumChildren.
func (a *Album) ListChildren(ctx context.Context, client Client, id ID) ([]*Album, error) {
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

// UploadImage uploads an image for an Album using Client.UploadAlbumImage.
func (a *Album) UploadImage(ctx context.Context, client Client, id ID, file []byte) (*Album, error) {
	return client.UploadAlbumImage(ctx, id, file)
}

// GetParent retrieves the parent Album using Client.GetAlbumParent.
func (a *Album) GetParent(ctx context.Context, client Client, id ID) (*Album, error) {
	return client.GetAlbumParent(ctx, id)
}

// ListPhotos retrieves all Photos for the given ID across all pages using Client.ListAlbumPhotos.
func (a *Album) ListPhotos(ctx context.Context, client Client, id ID) ([]*Photo, error) {
	var allPhotos []*Photo
	for page := 1; ; page++ {
		photos, err := client.ListAlbumPhotos(ctx, id, page)
		if err != nil {
			return fmt.Errorf("failed to list photos for ID %s on page %d: %w", id, page, err)
		}
		if len(photos) == 0 {
			break
		}
		allPhotos = append(allPhotos, photos...)
	}
	return allPhotos, nil
}
