package koiApi

import (
	"fmt"
	"os"
	"time"
)

// WishlistInterface defines methods for interacting with Wishlist resources.
type WishlistInterface interface {
	Create(client Client) (*Wishlist, error)                                               // HTTP POST /api/wishlists
	Delete(client Client, wishlistID ...ID) error                                          // HTTP DELETE /api/wishlists/{id}
	Get(client Client, wishlistID ...ID) (*Wishlist, error)                                // HTTP GET /api/wishlists/{id}
	GetParent(client Client, wishlistID ...ID) (*Wishlist, error)                          // HTTP GET /api/wishlists/{id}/parent
	IRI() string                                                                           // /api/wishlists/{id}
	List(client Client) ([]*Wishlist, error)                                               // HTTP GET /api/wishlists
	ListChildren(client Client, wishlistID ...ID) ([]*Wishlist, error)                     // HTTP GET /api/wishlists/{id}/children
	ListWishes(client Client, wishlistID ...ID) ([]*Wish, error)                           // HTTP GET /api/wishlists/{id}/wishes
	Patch(client Client, wishlistID ...ID) (*Wishlist, error)                              // HTTP PATCH /api/wishlists/{id}
	Update(client Client, wishlistID ...ID) (*Wishlist, error)                             // HTTP PUT /api/wishlists/{id}
	UploadImage(client Client, file []byte, wishlistID ...ID) (*Wishlist, error)           // HTTP POST /api/wishlists/{id}/image
	UploadImageByFile(client Client, filename string, wishlistID ...ID) (*Wishlist, error) // HTTP POST /api/wishlists/{id}/image
	Summary() string
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
func (w *Wishlist) Create(client Client) (*Wishlist, error) {
	return client.CreateWishlist(w)
}

// Delete
func (w *Wishlist) Delete(client Client, wishlistID ...ID) error {
	id := w.whichID(wishlistID...)
	return client.DeleteWishlist(id)
}

// Get
func (w *Wishlist) Get(client Client, wishlistID ...ID) (*Wishlist, error) {
	id := w.whichID(wishlistID...)
	return client.GetWishlist(id)
}

// GetParent
func (w *Wishlist) GetParent(client Client, wishlistID ...ID) (*Wishlist, error) {
	id := w.whichID(wishlistID...)
	return client.GetWishlistParent(id)
}

// IRI
func (w *Wishlist) IRI() string {
	return fmt.Sprintf("/api/wishlists/%s", w.ID)
}

// List
func (w *Wishlist) List(client Client) ([]*Wishlist, error) {
	return client.ListWishlists()
}

// ListChildren
func (w *Wishlist) ListChildren(client Client, wishlistID ...ID) ([]*Wishlist, error) {
	id := w.whichID(wishlistID...)
	return client.ListWishlistChildren(id)
}

// ListWishes
func (w *Wishlist) ListWishes(client Client, wishlistID ...ID) ([]*Wish, error) {
	id := w.whichID(wishlistID...)
	return client.ListWishlistWishes(id)
}

// Patch
func (w *Wishlist) Patch(client Client, wishlistID ...ID) (*Wishlist, error) {
	id := w.whichID(wishlistID...)
	return client.PatchWishlist(id, w)
}

// Update
func (w *Wishlist) Update(client Client, wishlistID ...ID) (*Wishlist, error) {
	id := w.whichID(wishlistID...)
	return client.UpdateWishlist(id, w)
}

// UploadImage
func (w *Wishlist) UploadImage(client Client, file []byte, wishlistID ...ID) (*Wishlist, error) {
	id := w.whichID(wishlistID...)
	return client.UploadWishlistImage(id, file)
}

// UploadImageByFile
func (w *Wishlist) UploadImageByFile(client Client, filename string, wishlistID ...ID) (*Wishlist, error) {
	id := w.whichID(wishlistID...)
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return client.UploadWishlistImage(id, file)
}
