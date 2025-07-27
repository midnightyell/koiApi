package koiApi

import (
	"fmt"
)

// FieldType represents the type of a template field (same values as DatumType).
type FieldType string // Read and write

const (
	FieldTypeText       FieldType = "text"
	FieldTypeTextarea   FieldType = "textarea"
	FieldTypeCountry    FieldType = "country"
	FieldTypeDate       FieldType = "date"
	FieldTypeRating     FieldType = "rating"
	FieldTypeNumber     FieldType = "number"
	FieldTypePrice      FieldType = "price"
	FieldTypeLink       FieldType = "link"
	FieldTypeList       FieldType = "list"
	FieldTypeChoiceList FieldType = "choice-list"
	FieldTypeCheckbox   FieldType = "checkbox"
	FieldTypeImage      FieldType = "image"
	FieldTypeFile       FieldType = "file"
	FieldTypeSign       FieldType = "sign"
	FieldTypeVideo      FieldType = "video"
	FieldTypeBlankLine  FieldType = "blank-line"
	FieldTypeSection    FieldType = "section"
)

func (ft FieldType) String() string {
	return string(ft)
}

// FieldInterface defines methods for interacting with Field resources.
type FieldInterface interface {
	Create(client Client) (*Field, error)                        // HTTP POST /api/fields
	Delete(client Client, fieldID ...ID) error                   // HTTP DELETE /api/fields/{id}
	Get(client Client, fieldID ...ID) (*Field, error)            // HTTP GET /api/fields/{id}
	GetTemplate(client Client, fieldID ...ID) (*Template, error) // HTTP GET /api/fields/{id}/template
	IRI() string                                                 // /api/fields/{id}
	Patch(client Client, fieldID ...ID) (*Field, error)          // HTTP PATCH /api/fields/{id}
	Update(client Client, fieldID ...ID) (*Field, error)         // HTTP PUT /api/fields/{id}
	Summary() string
}

// Field represents a template field in Koillection, combining fields for JSON-LD and API interactions.
type Field struct {
	Context    *Context   `json:"@context,omitempty" access:"rw"`   // JSON-LD only
	_ID        ID         `json:"@id,omitempty" access:"ro"`        // JSON-LD only
	Type       string     `json:"@type,omitempty" access:"rw"`      // JSON-LD only
	ID         ID         `json:"id,omitempty" access:"ro"`         // Identifier
	Name       string     `json:"name" access:"rw"`                 // Field name
	Position   int        `json:"position" access:"rw"`             // Field position
	FieldType  FieldType  `json:"type" access:"rw"`                 // Field type
	ChoiceList *string    `json:"choiceList,omitempty" access:"rw"` // Choice list IRI
	Template   *string    `json:"template" access:"rw"`             // Template IRI
	Visibility Visibility `json:"visibility,omitempty" access:"rw"` // Visibility level
	Owner      *string    `json:"owner,omitempty" access:"ro"`      // Owner IRI

}

// whichID
func (f *Field) whichID(fieldID ...ID) ID {
	if len(fieldID) > 0 {
		return fieldID[0]
	}
	return f.ID
}

// Create
func (f *Field) Create(client Client) (*Field, error) {
	return client.CreateField(f)
}

// Delete
func (f *Field) Delete(client Client, fieldID ...ID) error {
	id := f.whichID(fieldID...)
	return client.DeleteField(id)
}

// Get
func (f *Field) Get(client Client, fieldID ...ID) (*Field, error) {
	id := f.whichID(fieldID...)
	return client.GetField(id)
}

// GetTemplate
func (f *Field) GetTemplate(client Client, fieldID ...ID) (*Template, error) {
	id := f.whichID(fieldID...)
	return client.GetFieldTemplate(id)
}

// IRI
func (f *Field) IRI() string {
	return fmt.Sprintf("/api/fields/%s", f.ID)
}

// Patch
func (f *Field) Patch(client Client, fieldID ...ID) (*Field, error) {
	id := f.whichID(fieldID...)
	return client.PatchField(id, f)
}

// Update
func (f *Field) Update(client Client, fieldID ...ID) (*Field, error) {
	id := f.whichID(fieldID...)
	return client.UpdateField(id, f)
}
