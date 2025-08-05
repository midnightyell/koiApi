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

// GetID
func (a *Field) GetID() string {
	return string(a.ID)
}

// Validate
func (a *Field) Validate() error {
	if a.Name == "" {
		return fmt.Errorf("Name cannot be empty")
	}
	if a.FieldType == "" {
		return fmt.Errorf("FieldType cannot be empty")
	}
	return nil
}

// IRI
func (a *Field) IRI() string {
	return IRI(a)
}
