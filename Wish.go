package koiApi

import (
	"context"
	"fmt"
	"os"
	"time"
)

// WishInterface defines methods for interacting with Wish resources.
type WishInterface interface {
	Create(ctx context.Context, client Client) (*Wish, error)                                           // HTTP POST /api/wishes
	Delete(ctx context.Context, client Client, wishID ...ID) error                                      // HTTP DELETE /api/wishes/{id}
	Get(ctx context.Context, client Client, wishID ...ID) (*Wish, error)                                // HTTP GET /api/wishes/{id}
	GetWishlist(ctx context.Context, client Client, wishID ...ID) (*Wishlist, error)                    // HTTP GET /api/wishes/{id}/wishlist
	IRI() string                                                                                        // /api/wishes/{id}
	List(ctx context.Context, client Client) ([]*Wish, error)                                           // HTTP GET /api/wishes
	Patch(ctx context.Context, client Client, wishID ...ID) (*Wish, error)                              // HTTP PATCH /api/wishes/{id}
	Update(ctx context.Context, client Client, wishID ...ID) (*Wish, error)                             // HTTP PUT /api/wishes/{id}
	UploadImage(ctx context.Context, client Client, file []byte, wishID ...ID) (*Wish, error)           // HTTP POST /api/wishes/{id}/image
	UploadImageByFile(ctx context.Context, client Client, filename string, wishID ...ID) (*Wish, error) // HTTP POST /api/wishes/{id}/image
}

// Wish represents a wish in Koillection, combining fields for JSON-LD and API interactions.
type Wish struct {
	Context             *Context   `json:"@context,omitempty" access:"rw"`            // JSON-LD only
	_ID                 ID         `json:"@id,omitempty" access:"ro"`                 // JSON-LD only
	Type                string     `json:"@type,omitempty" access:"rw"`               // JSON-LD only
	ID                  ID         `json:"id,omitempty" access:"ro"`                  // Identifier
	Name                string     `json:"name" access:"rw"`                          // Wish name
	URL                 *string    `json:"url,omitempty" access:"rw"`                 // Wish URL
	Price               *string    `json:"price,omitempty" access:"rw"`               // Wish price
	Currency            *string    `json:"currency,omitempty" access:"rw"`            // Currency code
	Wishlist            *string    `json:"wishlist" access:"rw"`                      // Wishlist IRI
	Owner               *string    `json:"owner,omitempty" access:"ro"`               // Owner IRI
	Comment             *string    `json:"comment,omitempty" access:"rw"`             // Wish comment
	Image               *string    `json:"image,omitempty" access:"ro"`               // Image URL
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty" access:"ro"` // Small thumbnail URL
	Visibility          Visibility `json:"visibility,omitempty" access:"rw"`          // Visibility level
	ParentVisibility    *string    `json:"parentVisibility,omitempty" access:"ro"`    // Parent visibility
	FinalVisibility     Visibility `json:"finalVisibility,omitempty" access:"ro"`     // Effective visibility
	ScrapedFromURL      *string    `json:"scrapedFromUrl,omitempty" access:"ro"`      // Source URL
	CreatedAt           time.Time  `json:"createdAt" access:"ro"`                     // Creation timestamp
	UpdatedAt           *time.Time `json:"updatedAt,omitempty" access:"ro"`           // Update timestamp
	File                *string    `json:"file,omitempty" access:"wo"`                // Image file data
}

// whichID
func (w *Wish) whichID(wishID ...ID) ID {
	if len(wishID) > 0 {
		return wishID[0]
	}
	return w.ID
}

// Create
func (w *Wish) Create(ctx context.Context, client Client) (*Wish, error) {
	return client.CreateWish(ctx, w)
}

// Delete
func (w *Wish) Delete(ctx context.Context, client Client, wishID ...ID) error {
	id := w.whichID(wishID...)
	return client.DeleteWish(ctx, id)
}

// Get
func (w *Wish) Get(ctx context.Context, client Client, wishID ...ID) (*Wish, error) {
	id := w.whichID(wishID...)
	return client.GetWish(ctx, id)
}

// GetWishlist
func (w *Wish) GetWishlist(ctx context.Context, client Client, wishID ...ID) (*Wishlist, error) {
	id := w.whichID(wishID...)
	return client.GetWishWishlist(ctx, id)
}

// IRI
func (w *Wish) IRI() string {
	return fmt.Sprintf("/api/wishes/%s", w.ID)
}

// List
func (w *Wish) List(ctx context.Context, client Client) ([]*Wish, error) {
	return client.ListWishes(ctx)
}

// Patch
func (w *Wish) Patch(ctx context.Context, client Client, wishID ...ID) (*Wish, error) {
	id := w.whichID(wishID...)
	return client.PatchWish(ctx, id, w)
}

// Update
func (w *Wish) Update(ctx context.Context, client Client, wishID ...ID) (*Wish, error) {
	id := w.whichID(wishID...)
	return client.UpdateWish(ctx, id, w)
}

// UploadImage
func (w *Wish) UploadImage(ctx context.Context, client Client, file []byte, wishID ...ID) (*Wish, error) {
	id := w.whichID(wishID...)
	return client.UploadWishImage(ctx, id, file)
}

// UploadImageByFile
func (w *Wish) UploadImageByFile(ctx context.Context, client Client, filename string, wishID ...ID) (*Wish, error) {
	id := w.whichID(wishID...)
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return client.UploadWishImage(ctx, id, file)
}
