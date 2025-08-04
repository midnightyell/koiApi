package koiApi

import (
	"fmt"
	"time"
)

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
	Context             *Context   `json:"@context,omitempty" access:"rw"`            // JSON-LD only
	_ID                 ID         `json:"@id,omitempty" access:"ro"`                 // JSON-LD only
	Type                string     `json:"@type,omitempty" access:"rw"`               // JSON-LD only
	ID                  ID         `json:"id,omitempty" access:"ro"`                  // Identifier
	Item                *string    `json:"item,omitempty" access:"rw"`                // Item IRI
	Collection          *string    `json:"collection,omitempty" access:"rw"`          // Collection IRI
	DatumType           string     `json:"type" access:"rw"`                          // Custom data field type
	Label               string     `json:"label" access:"rw"`                         // Field label
	Value               *string    `json:"value,omitempty" access:"rw"`               // Field value
	Position            *int       `json:"position,omitempty" access:"rw"`            // Field position
	Currency            *string    `json:"currency,omitempty" access:"rw"`            // Currency code
	Image               *string    `json:"image,omitempty" access:"ro"`               // Image URL
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty" access:"ro"` // Small thumbnail URL
	ImageLargeThumbnail *string    `json:"imageLargeThumbnail,omitempty" access:"ro"` // Large thumbnail URL
	File                *string    `json:"file,omitempty" access:"ro"`                // File URL
	Video               *string    `json:"video,omitempty" access:"ro"`               // Video URL
	OriginalFilename    *string    `json:"originalFilename,omitempty" access:"ro"`    // Original file name
	ChoiceList          *string    `json:"choiceList,omitempty" access:"rw"`          // Choice list IRI
	Owner               *string    `json:"owner,omitempty" access:"ro"`               // Owner IRI
	Visibility          Visibility `json:"visibility,omitempty" access:"rw"`          // Visibility level
	ParentVisibility    *string    `json:"parentVisibility,omitempty" access:"ro"`    // Parent visibility
	FinalVisibility     Visibility `json:"finalVisibility,omitempty" access:"ro"`     // Effective visibility
	CreatedAt           time.Time  `json:"createdAt" access:"ro"`                     // Creation timestamp
	UpdatedAt           *time.Time `json:"updatedAt,omitempty" access:"ro"`           // Update timestamp
	FileImage           *string    `json:"fileImage,omitempty" access:"wo"`           // Image file data
	FileFile            *string    `json:"fileFile,omitempty" access:"wo"`            // File data
	FileVideo           *string    `json:"fileVideo,omitempty" access:"wo"`           // Video file data
}

// GetID
func (a *Datum) GetID() string {
	return string(a.ID)
}

// Validate
func (a *Datum) Validate() error {
	if a.Collection == nil && a.Item == nil {
		return fmt.Errorf("datum must belong to either a collection or an item")
	}
	if a.DatumType == "" {
		return fmt.Errorf("datum type cannot be empty")
	}
	if a.Label == "" {
		return fmt.Errorf("datum label cannot be empty")
	}
	return nil
}

// Create
func (a *Datum) Create() (*Datum, error) {
	return Create(a)
}

// Delete
func (a *Datum) Delete() error {
	return Delete(a)
}

// Get
func (a *Datum) Get() (*Datum, error) {
	res, err := Get(a)
	return res.(*Datum), err
}

// GetDefaultTemplate
func (a *Datum) GetDefaultTemplate() (*Template, error) {
	res, err := Get(a)
	return res.(*Template), err
}

// GetParent
func (a *Datum) GetParent() (*Datum, error) {
	res, err := Get(a)
	return res.(*Datum), err
}

// IRI
func (a *Datum) IRI() string {
	return IRI(a)
}

// List
func (a *Datum) List() ([]*Datum, error) {
	res, err := List(a)
	return res.([]*Datum), err
}

// Patch
func (a *Datum) Patch() (*Datum, error) {
	return Patch(a)
}

// Update
func (a *Datum) Update() (*Datum, error) {
	return Update(a)
}

// UploadImage
func (a *Datum) UploadImage(file []byte) (*Datum, error) {
	return Upload(a, file)
}

// UploadImageFromFile
func (a *Datum) UploadImageFromFile(filename string) (*Datum, error) {
	return UploadFromFile(a, filename)
}

func (a *Datum) UploadVideo(file []byte) (*Datum, error) {
	return Upload(a, file)
}

func (a *Datum) UploadVideoFromFile(filename string) (*Datum, error) {
	return UploadFromFile(a, filename)
}
