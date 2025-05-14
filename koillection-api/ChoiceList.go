package koiApi

import (
	"context"
	"fmt"
	"time"
)

// ChoiceListInterface defines methods for interacting with ChoiceList resources.
type ChoiceListInterface interface {
	Create(ctx context.Context, client Client) (*ChoiceList, error)
	Get(ctx context.Context, client Client, id ID) (*ChoiceList, error)
	List(ctx context.Context, client Client) ([]*ChoiceList, error)
	Update(ctx context.Context, client Client, id ID) (*ChoiceList, error)
	Delete(ctx context.Context, client Client, id ID) error
	ListChoiceListItems(ctx context.Context, client Client, id ID) ([]*Item, error)
	Validate(ctx context.Context, client Client) error
}

// ChoiceList represents a choice list in Koillection, combining read and write fields.
type ChoiceList struct {
	Context    *Context   `json:"@context,omitempty"`   // JSON-LD only
	ID_        ID         `json:"@id,omitempty"`        // JSON-LD only (maps to "@id" in JSON, read-only)
	ID         ID         `json:"id,omitempty"`         // JSON-LD only (maps to "id" in JSON, read-only)
	Type       string     `json:"@type,omitempty"`      // JSON-LD only
	Label      string     `json:"label"`                // Read and write
	Choices    []string   `json:"choices"`              // Read and write
	Visibility Visibility `json:"visibility,omitempty"` // Read and write
	CreatedAt  time.Time  `json:"createdAt"`            // Read-only
	UpdatedAt  *time.Time `json:"updatedAt,omitempty"`  // Read-only
}

// Validate checks the ChoiceList's fields for validity, using ctx for cancellation.
func (cl *ChoiceList) Validate(ctx context.Context, client Client) error {
	// Check for context cancellation
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("validation cancelled: %w", err)
	}

	// Required fields
	if cl.Label == "" {
		return fmt.Errorf("label must not be empty")
	}

	if len(cl.Choices) == 0 {
		return fmt.Errorf("choices must contain at least one item")
	}

	// Visibility must be a valid value
	switch cl.Visibility {
	case VisibilityPublic, VisibilityInternal, VisibilityPrivate, "":
		// Valid or unset (server may set default)
	default:
		return fmt.Errorf("invalid visibility value: %s", cl.Visibility)
	}

	// Read-only fields for creation vs. update
	if cl.ID == "" && cl.ID_ == "" {
		// Creation: read-only fields should be empty
		if cl.ID_ != "" {
			return fmt.Errorf("ID_ must be empty for creation")
		}
		if cl.ID != "" {
			return fmt.Errorf("ID must be empty for creation")
		}
		if cl.Type != "" && cl.Type != "ChoiceList" {
			return fmt.Errorf("Type must be empty or 'ChoiceList' for creation: %s", cl.Type)
		}
	} else {
		// Update: ID should be non-empty
		if cl.ID == "" {
			return fmt.Errorf("ID must not be empty for update")
		}
	}

	return nil
}

// Create calls Client.CreateChoiceList to create a new ChoiceList.
func (cl *ChoiceList) Create(ctx context.Context, client Client) (*ChoiceList, error) {
	return client.CreateChoiceList(ctx, cl)
}

// Get retrieves a ChoiceList by ID using Client.GetChoiceList.
func (cl *ChoiceList) Get(ctx context.Context, client Client, id ID) (*ChoiceList, error) {
	return client.GetChoiceList(ctx, id)
}

// List retrieves all ChoiceLists across all pages using Client.ListChoiceLists.
func (cl *ChoiceList) List(ctx context.Context, client Client) ([]*ChoiceList, error) {
	var allChoiceLists []*ChoiceList
	for page := 1; ; page++ {
		choiceLists, err := client.ListChoiceLists(ctx, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list choice lists on page %d: %w", page, err)
		}
		if len(choiceLists) == 0 {
			break
		}
		allChoiceLists = append(allChoiceLists, choiceLists...)
	}
	return allChoiceLists, nil
}

// Update updates a ChoiceList by ID using Client.UpdateChoiceList.
func (cl *ChoiceList) Update(ctx context.Context, client Client, id ID) (*ChoiceList, error) {
	return client.UpdateChoiceList(ctx, id, cl)
}

// Delete removes a ChoiceList by ID using Client.DeleteChoiceList.
func (cl *ChoiceList) Delete(ctx context.Context, client Client, id ID) error {
	return client.DeleteChoiceList(ctx, id)
}

// ListChoiceListItems retrieves all Items associated with the ChoiceList ID across all pages using Client.ListChoiceListItems.
func (cl *ChoiceList) ListChoiceListItems(ctx context.Context, client Client, id ID) ([]*Item, error) {
	var allItems []*Item
	for page := 1; ; page++ {
		items, err := client.ListChoiceListItems(ctx, id, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list items for ChoiceList ID %s on page %d: %w", id, page, err)
		}
		if len(items) == 0 {
			break
		}
		allItems = append(allItems, items...)
	}
	return allItems, nil
}
