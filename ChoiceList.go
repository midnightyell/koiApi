package koiApi

import (
	"time"
)

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

// IRI
func (a *ChoiceList) IRI() string {
	return IRI(a)
}

// GetID
func (a *ChoiceList) GetID() string {
	return string(a.ID)
}

// Validate
func (a *ChoiceList) Validate() error {
	return nil
}

// Create
func (a *ChoiceList) Create() (*ChoiceList, error) {
	return Create(a)
}

// Delete
func (a *ChoiceList) Delete() error {
	return Delete(a)
}

// Get
func (a *ChoiceList) Get() (*ChoiceList, error) {
	res, err := Get(a)
	return res.(*ChoiceList), err
}

// List
func (a *ChoiceList) List() ([]*ChoiceList, error) {
	res, err := List(a)
	return res.([]*ChoiceList), err
}

// Patch
func (a *ChoiceList) Patch() (*ChoiceList, error) {
	return Patch(a)
}

// Update
func (a *ChoiceList) Update() (*ChoiceList, error) {
	return Update(a)
}
