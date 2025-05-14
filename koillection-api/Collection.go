package koiApi

import (
	"context"
	"fmt"
	"time"
)

// CollectionInterface defines methods for interacting with Collection resources.
type CollectionInterface interface {
	Create(ctx context.Context, client Client) (*Collection, error)
	Get(ctx context.Context, client Client, id ID) (*Collection, error)
	List(ctx context.Context, client Client) ([]*Collection, error)
	Update(ctx context.Context, client Client, id ID) (*Collection, error)
	Delete(ctx context.Context, client Client, id ID) error
	ListCollectionItems(ctx context.Context, client Client, id ID) ([]*Item, error)
	Validate(ctx context.Context, client Client) error
}

// Collection represents a collection in Koillection, combining read and write fields.
type Collection struct {
	Context         *Context   `json:"@context,omitempty"`        // JSON-LD only
	ID_             ID         `json:"@id,omitempty"`             // JSON-LD only (maps to "@id" in JSON, read-only)
	ID              ID         `json:"id,omitempty"`              // JSON-LD only (maps to "id" in JSON, read-only)
	Type            string     `json:"@type,omitempty"`           // JSON-LD only
	Title           string     `json:"title"`                     // Read and write
	Visibility      Visibility `json:"visibility,omitempty"`      // Read and write
	Owner           *string    `json:"owner,omitempty"`           // Read-only, IRI
	SeenCounter     int        `json:"seenCounter,omitempty"`     // Read-only
	ChildrenCounter int        `json:"childrenCounter,omitempty"` // Read-only
	ItemsCounter    int        `json:"itemsCounter,omitempty"`    // Read-only
	CreatedAt       time.Time  `json:"createdAt"`                 // Read-only
	UpdatedAt       *time.Time `json:"updatedAt,omitempty"`       // Read-only
}

// Validate checks the Collection's fields for validity, using ctx for cancellation.
func (c *Collection) Validate(ctx context.Context, client Client) error {
	// Check for context cancellation
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("validation cancelled: %w", err)
	}

	// Required fields
	if c.Title == "" {
		return fmt.Errorf("title must not be empty")
	}

	// Visibility must be a valid value
	switch c.Visibility {
	case VisibilityPublic, VisibilityInternal, VisibilityPrivate, "":
		// Valid or unset (server may set default)
	default:
		return fmt.Errorf("invalid visibility value: %s", c.Visibility)
	}

	// Read-only fields for creation vs. update
	if c.ID == "" && c.ID_ == "" {
		// Creation: read-only fields should be empty
		if c.ID_ != "" {
			return fmt.Errorf("ID_ must be empty for creation")
		}
		if c.ID != "" {
			return fmt.Errorf("ID must be empty for creation")
		}
		if c.Type != "" && c.Type != "Collection" {
			return fmt.Errorf("Type must be empty or 'Collection' for creation: %s", c.Type)
		}
	} else {
		// Update: ID should be non-empty
		if c.ID == "" {
			return fmt.Errorf("ID must not be empty for update")
		}
	}

	return nil
}

// Create calls Client.CreateCollection to create a new Collection.
func (c *Collection) Create(ctx context.Context, client Client) (*Collection, error) {
	return client.CreateCollection(ctx, c)
}

// Get retrieves a Collection by ID using Client.GetCollection.
func (c *Collection) Get(ctx context.Context, client Client, id ID) (*Collection, error) {
	return client.GetCollection(ctx, id)
}

// List retrieves all Collections across all pages using Client.ListCollections.
func (c *Collection) List(ctx context.Context, client Client) ([]*Collection, error) {
	var allCollections []*Collection
	for page := 1; ; page++ {
		collections, err := client.ListCollections(ctx, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list collections on page %d: %w", page, err)
		}
		if len(collections) == 0 {
			break
		}
		allCollections = append(allCollections, collections...)
	}
	return allCollections, nil
}

// Update updates a Collection by ID using Client.UpdateCollection.
func (c *Collection) Update(ctx context.Context, client Client, id ID) (*Collection, error) {
	return client.UpdateCollection(ctx, id, c)
}

// Delete removes a Collection by ID using Client.DeleteCollection.
func (c *Collection) Delete(ctx context.Context, client Client, id ID) error {
	return client.DeleteCollection(ctx, id)
}

// ListCollectionItems retrieves all Items associated with the Collection ID across all pages using Client.ListCollectionItems.
func (c *Collection) ListCollectionItems(ctx context.Context, client Client, id ID) ([]*Item, error) {
	var allItems []*Item
	for page := 1; ; page++ {
		items, err := client.ListCollectionItems(ctx, id, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list items for Collection ID %s on page %d: %w", id, page, err)
		}
		if len(items) == 0 {
			break
		}
		allItems = append(allItems, items...)
	}
	return allItems, nil
}
