package koiApi

import (
	"context"
	"fmt"
	"time"
)

// TemplateInterface defines methods for interacting with Template resources.
type TemplateInterface interface {
	Create(ctx context.Context, client Client) (*Template, error)                   // HTTP POST /api/templates
	Delete(ctx context.Context, client Client, templateID ...ID) error              // HTTP DELETE /api/templates/{id}
	Get(ctx context.Context, client Client, templateID ...ID) (*Template, error)    // HTTP GET /api/templates/{id}
	IRI() string                                                                    // /api/templates/{id}
	List(ctx context.Context, client Client) ([]*Template, error)                   // HTTP GET /api/templates
	Patch(ctx context.Context, client Client, templateID ...ID) (*Template, error)  // HTTP PATCH /api/templates/{id}
	Update(ctx context.Context, client Client, templateID ...ID) (*Template, error) // HTTP PUT /api/templates/{id}
}

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

// whichID
func (t *Template) whichID(templateID ...ID) ID {
	if len(templateID) > 0 {
		return templateID[0]
	}
	return t.ID
}

// Create
func (t *Template) Create(ctx context.Context, client Client) (*Template, error) {
	return client.CreateTemplate(ctx, t)
}

// Delete
func (t *Template) Delete(ctx context.Context, client Client, templateID ...ID) error {
	id := t.whichID(templateID...)
	return client.DeleteTemplate(ctx, id)
}

// Get
func (t *Template) Get(ctx context.Context, client Client, templateID ...ID) (*Template, error) {
	id := t.whichID(templateID...)
	return client.GetTemplate(ctx, id)
}

// IRI
func (t *Template) IRI() string {
	return fmt.Sprintf("/api/templates/%s", t.ID)
}

// List
func (t *Template) List(ctx context.Context, client Client) ([]*Template, error) {
	return client.ListTemplates(ctx)
}

// Patch
func (t *Template) Patch(ctx context.Context, client Client, templateID ...ID) (*Template, error) {
	id := t.whichID(templateID...)
	return client.PatchTemplate(ctx, id, t)
}

// Update
func (t *Template) Update(ctx context.Context, client Client, templateID ...ID) (*Template, error) {
	id := t.whichID(templateID...)
	return client.UpdateTemplate(ctx, id, t)
}
