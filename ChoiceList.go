package koiApi

import (
	"context"
	"fmt"
	"time"
)

// ChoiceListInterface defines methods for interacting with ChoiceList resources.
type ChoiceListInterface interface {
	Create(ctx context.Context, client Client) (*ChoiceList, error)                     // HTTP POST /api/choice_lists
	Delete(ctx context.Context, client Client, choiceListID ...ID) error                // HTTP DELETE /api/choice_lists/{id}
	Get(ctx context.Context, client Client, choiceListID ...ID) (*ChoiceList, error)    // HTTP GET /api/choice_lists/{id}
	IRI() string                                                                        // /api/choice_lists/{id}
	List(ctx context.Context, client Client) ([]*ChoiceList, error)                     // HTTP GET /api/choice_lists
	Patch(ctx context.Context, client Client, choiceListID ...ID) (*ChoiceList, error)  // HTTP PATCH /api/choice_lists/{id}
	Update(ctx context.Context, client Client, choiceListID ...ID) (*ChoiceList, error) // HTTP PUT /api/choice_lists/{id}
}

// ChoiceList represents a choice list in Koillection, combining fields for JSON-LD and API interactions.
type ChoiceList struct {
	Context   *Context   `json:"@context,omitempty" access:"rw"`  // JSON-LD only
	_ID       ID         `json:"@id,omitempty" access:"ro"`       // JSON-LD only
	Type      string     `json:"@type,omitempty" access:"rw"`     // JSON-LD only
	ID        ID         `json:"id,omitempty" access:"ro"`        // Identifier
	Name      string     `json:"name" access:"rw"`                // Choice list name
	Choices   []string   `json:"choices" access:"rw"`             // List of choices
	Owner     *string    `json:"owner,omitempty" access:"ro"`     // Owner IRI
	CreatedAt time.Time  `json:"createdAt" access:"ro"`           // Creation timestamp
	UpdatedAt *time.Time `json:"updatedAt,omitempty" access:"ro"` // Update timestamp
}

// whichID
func (cl *ChoiceList) whichID(choiceListID ...ID) ID {
	if len(choiceListID) > 0 {
		return choiceListID[0]
	}
	return cl.ID
}

// Create
func (cl *ChoiceList) Create(ctx context.Context, client Client) (*ChoiceList, error) {
	return client.CreateChoiceList(ctx, cl)
}

// Delete
func (cl *ChoiceList) Delete(ctx context.Context, client Client, choiceListID ...ID) error {
	id := cl.whichID(choiceListID...)
	return client.DeleteChoiceList(ctx, id)
}

// Get
func (cl *ChoiceList) Get(ctx context.Context, client Client, choiceListID ...ID) (*ChoiceList, error) {
	id := cl.whichID(choiceListID...)
	return client.GetChoiceList(ctx, id)
}

// IRI
func (cl *ChoiceList) IRI() string {
	return fmt.Sprintf("/api/choice_lists/%s", cl.ID)
}

// List
func (cl *ChoiceList) List(ctx context.Context, client Client) ([]*ChoiceList, error) {
	var allChoiceLists []*ChoiceList
	for page := 1; ; page++ {
		choiceLists, err := client.ListChoiceLists(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to list choice lists on page %d: %w", err)
		}
		if len(choiceLists) == 0 {
			break
		}
		allChoiceLists = append(allChoiceLists, choiceLists...)
	}
	return allChoiceLists, nil
}

// Patch
func (cl *ChoiceList) Patch(ctx context.Context, client Client, choiceListID ...ID) (*ChoiceList, error) {
	id := cl.whichID(choiceListID...)
	return client.PatchChoiceList(ctx, id, cl)
}

// Update
func (cl *ChoiceList) Update(ctx context.Context, client Client, choiceListID ...ID) (*ChoiceList, error) {
	id := cl.whichID(choiceListID...)
	return client.UpdateChoiceList(ctx, id, cl)
}
