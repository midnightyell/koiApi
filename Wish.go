package koiApi

import (
	"fmt"
	"os"
	"time"
)

// WishInterface defines methods for interacting with Wish resources.
type WishInterface interface {
	Create(client Client) (*Wish, error)                                           // HTTP POST /api/wishes
	Delete(client Client, wishID ...ID) error                                      // HTTP DELETE /api/wishes/{id}
	Get(client Client, wishID ...ID) (*Wish, error)                                // HTTP GET /api/wishes/{id}
	GetWishlist(client Client, wishID ...ID) (*Wishlist, error)                    // HTTP GET /api/wishes/{id}/wishlist
	IRI() string                                                                   // /api/wishes/{id}
	List(client Client) ([]*Wish, error)                                           // HTTP GET /api/wishes
	Patch(client Client, wishID ...ID) (*Wish, error)                              // HTTP PATCH /api/wishes/{id}
	Update(client Client, wishID ...ID) (*Wish, error)                             // HTTP PUT /api/wishes/{id}
	UploadImage(client Client, file []byte, wishID ...ID) (*Wish, error)           // HTTP POST /api/wishes/{id}/image
	UploadImageByFile(client Client, filename string, wishID ...ID) (*Wish, error) // HTTP POST /api/wishes/{id}/image
	Summary() string
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
func (w *Wish) Create(client Client) (*Wish, error) {
	return client.CreateWish(w)
}

// Delete
func (w *Wish) Delete(client Client, wishID ...ID) error {
	id := w.whichID(wishID...)
	return client.DeleteWish(id)
}

// Get
func (w *Wish) Get(client Client, wishID ...ID) (*Wish, error) {
	id := w.whichID(wishID...)
	return client.GetWish(id)
}

// GetWishlist
func (w *Wish) GetWishlist(client Client, wishID ...ID) (*Wishlist, error) {
	id := w.whichID(wishID...)
	return client.GetWishWishlist(id)
}

// IRI
func (w *Wish) IRI() string {
	return fmt.Sprintf("/api/wishes/%s", w.ID)
}

// List
func (w *Wish) List(client Client) ([]*Wish, error) {
	return client.ListWishes()
}

// Patch
func (w *Wish) Patch(client Client, wishID ...ID) (*Wish, error) {
	id := w.whichID(wishID...)
	return client.PatchWish(id, w)
}

// Update
func (w *Wish) Update(client Client, wishID ...ID) (*Wish, error) {
	id := w.whichID(wishID...)
	return client.UpdateWish(id, w)
}

// UploadImage
func (w *Wish) UploadImage(client Client, file []byte, wishID ...ID) (*Wish, error) {
	id := w.whichID(wishID...)
	return client.UploadWishImage(id, file)
}

// UploadImageByFile
func (w *Wish) UploadImageByFile(client Client, filename string, wishID ...ID) (*Wish, error) {
	id := w.whichID(wishID...)
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return client.UploadWishImage(id, file)
}
