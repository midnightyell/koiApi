package koiApi

import (
	"fmt"
	"time"
)

// Template represents a template in Koillection, combining fields for JSON-LD and API interactions.
type Template struct {
	Context   *Context   `json:"@context,omitempty" access:"rw"`  // JSON-LD only
	_ID       ID         `json:"@id,omitempty" access:"ro"`       // JSON-LD only
	Type      string     `json:"@type,omitempty" access:"rw"`     // JSON-LD only
	ID        ID         `json:"id,omitempty" access:"ro"`        // Identifier
	Name      string     `json:"name" access:"rw"`                // Template name
	Owner     *string    `json:"owner,omitempty" access:"ro"`     // Owner IRI
	CreatedAt time.Time  `json:"createdAt" access:"ro"`           // Creation timestamp
	UpdatedAt *time.Time `json:"updatedAt,omitempty" access:"ro"` // Update timestamp

}

func (t *Template) Summary() string {
	return fmt.Sprintf("%-40s %s", t.Name, t.ID)
}

// IRI
func (t *Template) IRI() string {
	return fmt.Sprintf("/api/templates/%s", t.ID)
}

func (t *Template) GetID() string {
	return string(t.ID)
}

func (t *Template) Validate() error {
	var errs []string
	// name is required, type string; see components.schemas.Template-template.write.required
	if t.Name == "" {
		errs = append(errs, "template name is required")
	}
	validateVisibility(t, &errs)
	return validationErrors(&errs)
}
