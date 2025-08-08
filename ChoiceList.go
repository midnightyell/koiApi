package koiApi

import (
	"fmt"
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

func (a *ChoiceList) Summary() string {
	return fmt.Sprintf("%-40s %s", a.Name, a.ID)
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
	var errs []string
	// choices array of unique strings; see components.schemas.ChoiceList-choiceList.write.properties.choices
	if len(a.Choices) > 0 {
		seen := make(map[string]struct{})
		for _, choice := range a.Choices {
			if _, exists := seen[choice]; exists {
				errs = append(errs, fmt.Sprintf("duplicate choice found: %s", choice))
			}
			seen[choice] = struct{}{}
		}
	}
	return validationErrors(&errs)
}
