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

func (a *Field) Summary() string {
	return fmt.Sprintf("%-40s %s", a.Name, a.ID)
}

// GetID
func (a *Field) GetID() string {
	return string(a.ID)
}

// Validate
func (a *Field) Validate() error {
	var errs []string
	// name is required, type string; see components.schemas.Field-a.write.required
	if a.Name == "" {
		errs = append(errs, "field name is required")
	}
	// type is required, enum; see components.schemas.Field-a.write.required
	if a.FieldType == "" {
		errs = append(errs, "field type is required")
	} else {
		validTypes := []string{"text", "textarea", "country", "date", "rating", "number", "price", "link", "list", "choice-list", "checkbox", "image", "file", "sign", "video", "blank-line", "section"}
		valid := false
		for _, t := range validTypes {
			if string(a.FieldType) == t {
				valid = true
				break
			}
		}
		if !valid {
			errs = append(errs, fmt.Sprintf("invalid field type: %s; must be one of %v", a.FieldType, validTypes))
		}
	}
	// template is required, type string or null (IRI); see components.schemas.Field-a.write.required
	if a.Template == nil {
		errs = append(errs, "field template IRI is required")
	}
	validateVisibility(a, &errs)
	return validationErrors(&errs)
}

// IRI
func (a *Field) IRI() string {
	return IRI(a)
}
