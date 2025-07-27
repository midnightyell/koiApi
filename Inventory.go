package koiApi

import (
	"fmt"
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

// whichID
func (i *Inventory) whichID(inventoryID ...ID) ID {
	if len(inventoryID) > 0 {
		return inventoryID[0]
	}
	return i.ID
}

// Delete
func (i *Inventory) Delete(client Client, inventoryID ...ID) error {
	id := i.whichID(inventoryID...)
	return client.DeleteInventory(id)
}

// Get
func (i *Inventory) Get(client Client, inventoryID ...ID) (*Inventory, error) {
	id := i.whichID(inventoryID...)
	return client.GetInventory(id)
}

// IRI
func (i *Inventory) IRI() string {
	return fmt.Sprintf("/api/inventories/%s", i.ID)
}
