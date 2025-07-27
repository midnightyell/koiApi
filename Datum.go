package koiApi

import (
	"fmt"
	"os"
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

// DatumInterface defines methods for interacting with Datum resources.
type DatumInterface interface {
	Create(client Client) (*Datum, error)                                            // HTTP POST /api/data
	Delete(client Client, datumID ...ID) error                                       // HTTP DELETE /api/data/{id}
	Get(client Client, datumID ...ID) (*Datum, error)                                // HTTP GET /api/data/{id}
	GetCollection(client Client, datumID ...ID) (*Collection, error)                 // HTTP GET /api/data/{id}/collection
	GetItem(client Client, datumID ...ID) (*Item, error)                             // HTTP GET /api/data/{id}/item
	IRI() string                                                                     // /api/data/{id}
	Patch(client Client, datumID ...ID) (*Datum, error)                              // HTTP PATCH /api/data/{id}
	Update(client Client, datumID ...ID) (*Datum, error)                             // HTTP PUT /api/data/{id}
	UploadFile(client Client, file []byte, datumID ...ID) (*Datum, error)            // HTTP POST /api/data/{id}/file
	UploadFileByFile(client Client, filename string, datumID ...ID) (*Datum, error)  // HTTP POST /api/data/{id}/file
	UploadImage(client Client, image []byte, datumID ...ID) (*Datum, error)          // HTTP POST /api/data/{id}/image
	UploadImageByFile(client Client, filename string, datumID ...ID) (*Datum, error) // HTTP POST /api/data/{id}/image
	UploadVideo(client Client, video []byte, datumID ...ID) (*Datum, error)          // HTTP POST /api/data/{id}/video
	UploadVideoByFile(client Client, filename string, datumID ...ID) (*Datum, error) // HTTP POST /api/data/{id}/video
	Summary() string
}

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

// whichID
func (d *Datum) whichID(datumID ...ID) ID {
	if len(datumID) > 0 {
		return datumID[0]
	}
	return d.ID
}

// Create
func (d *Datum) Create(client Client) (*Datum, error) {
	return client.CreateDatum(d)
}

// Delete
func (d *Datum) Delete(client Client, datumID ...ID) error {
	id := d.whichID(datumID...)
	return client.DeleteDatum(id)
}

// Get
func (d *Datum) Get(client Client, datumID ...ID) (*Datum, error) {
	id := d.whichID(datumID...)
	return client.GetDatum(id)
}

// GetCollection
func (d *Datum) GetCollection(client Client, datumID ...ID) (*Collection, error) {
	id := d.whichID(datumID...)
	return client.GetDatumCollection(id)
}

// GetItem
func (d *Datum) GetItem(client Client, datumID ...ID) (*Item, error) {
	id := d.whichID(datumID...)
	return client.GetDatumItem(id)
}

// IRI
func (d *Datum) IRI() string {
	return fmt.Sprintf("/api/data/%s", d.ID)
}

// Patch
func (d *Datum) Patch(client Client, datumID ...ID) (*Datum, error) {
	id := d.whichID(datumID...)
	return client.PatchDatum(id, d)
}

// Update
func (d *Datum) Update(client Client, datumID ...ID) (*Datum, error) {
	id := d.whichID(datumID...)
	return client.UpdateDatum(id, d)
}

// UploadFile
func (d *Datum) UploadFile(client Client, file []byte, datumID ...ID) (*Datum, error) {
	id := d.whichID(datumID...)
	return client.UploadDatumFile(id, file)
}

// UploadFileByFile
func (d *Datum) UploadFileByFile(client Client, filename string, datumID ...ID) (*Datum, error) {
	id := d.whichID(datumID...)
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return client.UploadDatumFile(id, file)
}

// UploadImage
func (d *Datum) UploadImage(client Client, image []byte, datumID ...ID) (*Datum, error) {
	id := d.whichID(datumID...)
	return client.UploadDatumImage(id, image)
}

// UploadImageByFile
func (d *Datum) UploadImageByFile(client Client, filename string, datumID ...ID) (*Datum, error) {
	id := d.whichID(datumID...)
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return client.UploadDatumImage(id, file)
}

// UploadVideo
func (d *Datum) UploadVideo(client Client, video []byte, datumID ...ID) (*Datum, error) {
	id := d.whichID(datumID...)
	return client.UploadDatumVideo(id, video)
}

// UploadVideoByFile
func (d *Datum) UploadVideoByFile(client Client, filename string, datumID ...ID) (*Datum, error) {
	id := d.whichID(datumID...)
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return client.UploadDatumVideo(id, file)
}
