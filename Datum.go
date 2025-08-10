package koiApi

import (
	"fmt"
	"time"
)

// TODO refactor into enum of some sort
const (
	DatumTypeText       string = "text"
	DatumTypeTextarea   string = "textarea"
	DatumTypeCountry    string = "country"
	DatumTypeDate       string = "date"
	DatumTypeRating     string = "rating"
	DatumTypeNumber     string = "number"
	DatumTypePrice      string = "price"
	DatumTypeLink       string = "link"
	DatumTypeList       string = "list"
	DatumTypeChoiceList string = "choice-list"
	DatumTypeCheckbox   string = "checkbox"
	DatumTypeImage      string = "image"
	DatumTypeFile       string = "file"
	DatumTypeSign       string = "sign"
	DatumTypeVideo      string = "video"
	DatumTypeBlankLine  string = "blank-line"
	DatumTypeSection    string = "section"
)

// Datum represents a custom data field in Koillection, combining fields for JSON-LD and API interactions.
type Datum struct {
	Context             Context    `json:"@context,omitempty" access:"rw"`            // JSON-LD only
	_ID                 ID         `json:"@id,omitempty" access:"ro"`                 // JSON-LD only
	Type                string     `json:"@type,omitempty" access:"rw"`               // JSON-LD only
	ID                  ID         `json:"id,omitempty" access:"ro"`                  // Identifier
	Item                string     `json:"item,omitempty" access:"rw"`                // Item IRI
	Collection          string     `json:"collection,omitempty" access:"rw"`          // Collection IRI
	DatumType           string     `json:"type" access:"rw"`                          // Custom data field type
	Label               string     `json:"label" access:"rw"`                         // Field label
	Value               string     `json:"value,omitempty" access:"rw"`               // Field value
	Position            int        `json:"position,omitempty" access:"rw"`            // Field position
	Currency            string     `json:"currency,omitempty" access:"rw"`            // Currency code
	Image               string     `json:"image,omitempty" access:"ro"`               // Image URL
	ImageSmallThumbnail string     `json:"imageSmallThumbnail,omitempty" access:"ro"` // Small thumbnail URL
	ImageLargeThumbnail string     `json:"imageLargeThumbnail,omitempty" access:"ro"` // Large thumbnail URL
	File                string     `json:"file,omitempty" access:"ro"`                // File URL
	Video               string     `json:"video,omitempty" access:"ro"`               // Video URL
	OriginalFilename    string     `json:"originalFilename,omitempty" access:"ro"`    // Original file name
	ChoiceList          string     `json:"choiceList,omitempty" access:"rw"`          // Choice list IRI
	Owner               string     `json:"owner,omitempty" access:"ro"`               // Owner IRI
	Visibility          Visibility `json:"visibility,omitempty" access:"rw"`          // Visibility level
	ParentVisibility    string     `json:"parentVisibility,omitempty" access:"ro"`    // Parent visibility
	FinalVisibility     Visibility `json:"finalVisibility,omitempty" access:"ro"`     // Effective visibility
	CreatedAt           time.Time  `json:"createdAt" access:"ro"`                     // Creation timestamp
	UpdatedAt           time.Time  `json:"updatedAt,omitempty" access:"ro"`           // Update timestamp
	FileImage           string     `json:"fileImage,omitempty" access:"wo"`           // Image file data
	FileFile            string     `json:"fileFile,omitempty" access:"wo"`            // File data
	FileVideo           string     `json:"fileVideo,omitempty" access:"wo"`           // Video file data
}

func (a *Datum) Summary() string {
	return fmt.Sprintf("%-25s: %-20s %s", a.Label, a.Value, a.Item)
}

// GetID
func (a *Datum) GetID() string {
	return string(a.ID)
}

// Validate
func (a *Datum) Validate() error {
	var errs []string
	// type is required, enum; see components.schemas.Datum-a.write.required
	if a.DatumType == "" {
		errs = append(errs, "datum type is required")
	} else {
		validTypes := []string{"text", "textarea", "country", "date", "rating", "number", "price", "link", "list", "choice-list", "checkbox", "image", "file", "sign", "video", "blank-line", "section"}
		valid := false
		for _, t := range validTypes {
			if string(a.DatumType) == t {
				valid = true
				break
			}
		}
		if !valid {
			errs = append(errs, fmt.Sprintf("invalid datum type: %s; must be one of %v", a.DatumType, validTypes))
		}
	}
	// label is required, type string; see components.schemas.Datum-a.write.required
	if a.Label == "" {
		errs = append(errs, "datum label is required")
	}
	// visibility enum ["public", "internal", "private"]; see components.schemas.Datum-a.write.properties.visibility
	if a.Visibility != "" {
		switch a.Visibility {
		case VisibilityPublic, VisibilityInternal, VisibilityPrivate:
		default:
			errs = append(errs, fmt.Sprintf("invalid visibility: %s; must be public, internal, or private", a.Visibility))
		}
	}
	// currency follows https://schema.org/priceCurrency; see components.schemas.Datum-a.write.properties.currency
	if a.Currency != "" {
		if !validateCurrency(a.Currency) {
			errs = append(errs, fmt.Sprintf("invalid currency code: %s", a.Currency))
		}
	}
	return validationErrors(&errs)
}

// IRI
func (a *Datum) IRI() string {
	return IRI(a)
}
