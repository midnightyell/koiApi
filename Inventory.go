package koiApi

import (
	"time"
)

// InventoryInterface defines methods for interacting with Inventory resources.
type InventoryInterface interface {
	Delete(client Client, inventoryID ...ID) error            // HTTP DELETE /api/inventories/{id}
	Get(client Client, inventoryID ...ID) (*Inventory, error) // HTTP GET /api/inventories/{id}
	IRI() string                                              // /api/inventories/{id}
	Summary() string
}

// Inventory represents an inventory record in Koillection, combining fields for JSON-LD and API interactions.
type Inventory struct {
	Context   *Context   `json:"@context,omitempty" access:"rw"`  // JSON-LD only
	_ID       ID         `json:"@id,omitempty" access:"ro"`       // JSON-LD only
	Type      string     `json:"@type,omitempty" access:"rw"`     // JSON-LD only
	ID        ID         `json:"id,omitempty" access:"ro"`        // Identifier
	Name      string     `json:"name" access:"rw"`                // Inventory name
	Content   []string   `json:"content" access:"rw"`             // Inventory content
	Owner     *string    `json:"owner,omitempty" access:"ro"`     // Owner IRI
	CreatedAt time.Time  `json:"createdAt" access:"ro"`           // Creation timestamp
	UpdatedAt *time.Time `json:"updatedAt,omitempty" access:"ro"` // Update timestamp

}

// GetID
func (a *Inventory) GetID() string {
	return string(a.ID)
}

// Create
// Not supported for Inventory?
func (a *Inventory) Create() (*Inventory, error) {
	return Create(a)
}

// Validate
func (a *Inventory) Validate() error {
	return nil
}

// Delete
func (a *Inventory) Delete() error {
	return Delete(a)
}

// Get
func (a *Inventory) Get() (*Inventory, error) {
	res, err := Get(a)
	return res.(*Inventory), err
}

// IRI
func (a *Inventory) IRI() string {
	return IRI(a)
}
