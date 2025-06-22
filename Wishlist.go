package koiApi

import (
	"context"
	"fmt"
	"os"
	"time"
)

// WishlistInterface defines methods for interacting with Wishlist resources.
type WishlistInterface interface {
	Create(ctx context.Context, client Client) (*Wishlist, error)                                               // HTTP POST /api/wishlists
	Delete(ctx context.Context, client Client, wishlistID ...ID) error                                          // HTTP DELETE /api/wishlists/{id}
	Get(ctx context.Context, client Client, wishlistID ...ID) (*Wishlist, error)                                // HTTP GET /api/wishlists/{id}
	GetParent(ctx context.Context, client Client, wishlistID ...ID) (*Wishlist, error)                          // HTTP GET /api/wishlists/{id}/parent
	IRI() string                                                                                                // /api/wishlists/{id}
	List(ctx context.Context, client Client) ([]*Wishlist, error)                                               // HTTP GET /api/wishlists
	ListChildren(ctx context.Context, client Client, wishlistID ...ID) ([]*Wishlist, error)                     // HTTP GET /api/wishlists/{id}/children
	ListWishes(ctx context.Context, client Client, wishlistID ...ID) ([]*Wish, error)                           // HTTP GET /api/wishlists/{id}/wishes
	Patch(ctx context.Context, client Client, wishlistID ...ID) (*Wishlist, error)                              // HTTP PATCH /api/wishlists/{id}
	Update(ctx context.Context, client Client, wishlistID ...ID) (*Wishlist, error)                             // HTTP PUT /api/wishlists/{id}
	UploadImage(ctx context.Context, client Client, file []byte, wishlistID ...ID) (*Wishlist, error)           // HTTP POST /api/wishlists/{id}/image
	UploadImageByFile(ctx context.Context, client Client, filename string, wishlistID ...ID) (*Wishlist, error) // HTTP POST /api/wishlists/{id}/image
}

// Wishlist represents a wishlist in Koillection, combining fields for JSON-LD and API interactions.
type Wishlist struct {
	Context          *Context   `json:"@context,omitempty" access:"rw"`         // JSON-LD only
	_ID              ID         `json:"@id,omitempty" access:"ro"`              // JSON-LD only
	Type             string     `json:"@type,omitempty" access:"rw"`            // JSON-LD only
	ID               ID         `json:"id,omitempty" access:"ro"`               // Identifier
	Name             string     `json:"name" access:"rw"`                       // Wishlist name
	Owner            *string    `json:"owner,omitempty" access:"ro"`            // Owner IRI
	Color            string     `json:"color" access:"ro"`                      // Color code
	Parent           *string    `json:"parent,omitempty" access:"rw"`           // Parent wishlist IRI
	Image            *string    `json:"image,omitempty" access:"ro"`            // Image URL
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
func (w *Wishlist) whichID(wishlistID ...ID) ID {
	if len(wishlistID) > 0 {
		return wishlistID[0]
	}
	return w.ID
}

// Create
func (w *Wishlist) Create(ctx context.Context, client Client) (*Wishlist, error) {
	return client.CreateWishlist(ctx, w)
}

// Delete
func (w *Wishlist) Delete(ctx context.Context, client Client, wishlistID ...ID) error {
	id := w.whichID(wishlistID...)
	return client.DeleteWishlist(ctx, id)
}

// Get
func (w *Wishlist) Get(ctx context.Context, client Client, wishlistID ...ID) (*Wishlist, error) {
	id := w.whichID(wishlistID...)
	return client.GetWishlist(ctx, id)
}

// GetParent
func (w *Wishlist) GetParent(ctx context.Context, client Client, wishlistID ...ID) (*Wishlist, error) {
	id := w.whichID(wishlistID...)
	return client.GetWishlistParent(ctx, id)
}

// IRI
func (w *Wishlist) IRI() string {
	return fmt.Sprintf("/api/wishlists/%s", w.ID)
}

// List
func (w *Wishlist) List(ctx context.Context, client Client) ([]*Wishlist, error) {
	var allWishlists []*Wishlist
	for page := 1; ; page++ {
		wishlists, err := client.ListWishlists(ctx, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list wishlists on page %d: %w", page, err)
		}
		if len(wishlists) == 0 {
			break
		}
		allWishlists = append(allWishlists, wishlists...)
	}
	return allWishlists, nil
}

// ListChildren
func (w *Wishlist) ListChildren(ctx context.Context, client Client, wishlistID ...ID) ([]*Wishlist, error) {
	id := w.whichID(wishlistID...)
	var allChildren []*Wishlist
	for page := 1; ; page++ {
		children, err := client.ListWishlistChildren(ctx, id, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list child wishlists for ID %s on page %d: %w", id, page, err)
		}
		if len(children) == 0 {
			break
		}
		allChildren = append(allChildren, children...)
	}
	return allChildren, nil
}

// ListWishes
func (w *Wishlist) ListWishes(ctx context.Context, client Client, wishlistID ...ID) ([]*Wish, error) {
	id := w.whichID(wishlistID...)
	var allWishes []*Wish
	for page := 1; ; page++ {
		wishes, err := client.ListWishlistWishes(ctx, id, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list wishes for ID %s on page %d: %w", id, page, err)
		}
		if len(wishes) == 0 {
			break
		}
		allWishes = append(allWishes, wishes...)
	}
	return allWishes, nil
}

// Patch
func (w *Wishlist) Patch(ctx context.Context, client Client, wishlistID ...ID) (*Wishlist, error) {
	id := w.whichID(wishlistID...)
	return client.PatchWishlist(ctx, id, w)
}

// Update
func (w *Wishlist) Update(ctx context.Context, client Client, wishlistID ...ID) (*Wishlist, error) {
	id := w.whichID(wishlistID...)
	return client.UpdateWishlist(ctx, id, w)
}

// UploadImage
func (w *Wishlist) UploadImage(ctx context.Context, client Client, file []byte, wishlistID ...ID) (*Wishlist, error) {
	id := w.whichID(wishlistID...)
	return client.UploadWishlistImage(ctx, id, file)
}

// UploadImageByFile
func (w *Wishlist) UploadImageByFile(ctx context.Context, client Client, filename string, wishlistID ...ID) (*Wishlist, error) {
	id := w.whichID(wishlistID...)
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return client.UploadWishlistImage(ctx, id, file)
}
