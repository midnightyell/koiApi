package koiApi

import (
	"context"
	"fmt"
	"os"
	"time"
)

type DatumType string

const (
	DatumTypeText       DatumType = "text"
	DatumTypeTextarea   DatumType = "textarea"
	DatumTypeCountry    DatumType = "country"
	DatumTypeDate       DatumType = "date"
	DatumTypeRating     DatumType = "rating"
	DatumTypeNumber     DatumType = "number"
	DatumTypePrice      DatumType = "price"
	DatumTypeLink       DatumType = "link"
	DatumTypeList       DatumType = "list"
	DatumTypeChoiceList DatumType = "choice-list"
	DatumTypeCheckbox   DatumType = "checkbox"
	DatumTypeImage      DatumType = "image"
	DatumTypeFile       DatumType = "file"
	DatumTypeSign       DatumType = "sign"
	DatumTypeVideo      DatumType = "video"
	DatumTypeBlankLine  DatumType = "blank-line"
	DatumTypeSection    DatumType = "section"
)

func (dt DatumType) String() string {
	return string(dt)
}

// DatumInterface defines methods for interacting with Datum resources.
type DatumInterface interface {
	Create(ctx context.Context, client Client) (*Datum, error)                                            // HTTP POST /api/data
	Delete(ctx context.Context, client Client, datumID ...ID) error                                       // HTTP DELETE /api/data/{id}
	Get(ctx context.Context, client Client, datumID ...ID) (*Datum, error)                                // HTTP GET /api/data/{id}
	GetCollection(ctx context.Context, client Client, datumID ...ID) (*Collection, error)                 // HTTP GET /api/data/{id}/collection
	GetItem(ctx context.Context, client Client, datumID ...ID) (*Item, error)                             // HTTP GET /api/data/{id}/item
	IRI() string                                                                                          // /api/data/{id}
	List(ctx context.Context, client Client) ([]*Datum, error)                                            // HTTP GET /api/data
	Patch(ctx context.Context, client Client, datumID ...ID) (*Datum, error)                              // HTTP PATCH /api/data/{id}
	Update(ctx context.Context, client Client, datumID ...ID) (*Datum, error)                             // HTTP PUT /api/data/{id}
	UploadFile(ctx context.Context, client Client, file []byte, datumID ...ID) (*Datum, error)            // HTTP POST /api/data/{id}/file
	UploadFileByFile(ctx context.Context, client Client, filename string, datumID ...ID) (*Datum, error)  // HTTP POST /api/data/{id}/file
	UploadImage(ctx context.Context, client Client, image []byte, datumID ...ID) (*Datum, error)          // HTTP POST /api/data/{id}/image
	UploadImageByFile(ctx context.Context, client Client, filename string, datumID ...ID) (*Datum, error) // HTTP POST /api/data/{id}/image
	UploadVideo(ctx context.Context, client Client, video []byte, datumID ...ID) (*Datum, error)          // HTTP POST /api/data/{id}/video
	UploadVideoByFile(ctx context.Context, client Client, filename string, datumID ...ID) (*Datum, error) // HTTP POST /api/data/{id}/video
}

// Datum represents a custom data field in Koillection, combining fields for JSON-LD and API interactions.
type Datum struct {
	Context             *Context   `json:"@context,omitempty" access:"rw"`            // JSON-LD only
	_ID                 ID         `json:"@id,omitempty" access:"ro"`                 // JSON-LD only
	Type                string     `json:"@type,omitempty" access:"rw"`               // JSON-LD only
	ID                  ID         `json:"id,omitempty" access:"ro"`                  // Identifier
	Item                *string    `json:"item,omitempty" access:"rw"`                // Item IRI
	Collection          *string    `json:"collection,omitempty" access:"rw"`          // Collection IRI
	DatumType           DatumType  `json:"type" access:"rw"`                          // Custom data field type
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
func (d *Datum) Create(ctx context.Context, client Client) (*Datum, error) {
	return client.CreateDatum(ctx, d)
}

// Delete
func (d *Datum) Delete(ctx context.Context, client Client, datumID ...ID) error {
	id := d.whichID(datumID...)
	return client.DeleteDatum(ctx, id)
}

// Get
func (d *Datum) Get(ctx context.Context, client Client, datumID ...ID) (*Datum, error) {
	id := d.whichID(datumID...)
	return client.GetDatum(ctx, id)
}

// GetCollection
func (d *Datum) GetCollection(ctx context.Context, client Client, datumID ...ID) (*Collection, error) {
	id := d.whichID(datumID...)
	return client.GetDatumCollection(ctx, id)
}

// GetItem
func (d *Datum) GetItem(ctx context.Context, client Client, datumID ...ID) (*Item, error) {
	id := d.whichID(datumID...)
	return client.GetDatumItem(ctx, id)
}

// IRI
func (d *Datum) IRI() string {
	return fmt.Sprintf("/api/data/%s", d.ID)
}

// List
func (d *Datum) List(ctx context.Context, client Client) ([]*Datum, error) {
	var allData []*Datum
	for page := 1; ; page++ {
		data, err := client.ListData(ctx, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list data on page %d: %w", page, err)
		}
		if len(data) == 0 {
			break
		}
		allData = append(allData, data...)
	}
	return allData, nil
}

// Patch
func (d *Datum) Patch(ctx context.Context, client Client, datumID ...ID) (*Datum, error) {
	id := d.whichID(datumID...)
	return client.PatchDatum(ctx, id, d)
}

// Update
func (d *Datum) Update(ctx context.Context, client Client, datumID ...ID) (*Datum, error) {
	id := d.whichID(datumID...)
	return client.UpdateDatum(ctx, id, d)
}

// UploadFile
func (d *Datum) UploadFile(ctx context.Context, client Client, file []byte, datumID ...ID) (*Datum, error) {
	id := d.whichID(datumID...)
	return client.UploadDatumFile(ctx, id, file)
}

// UploadFileByFile
func (d *Datum) UploadFileByFile(ctx context.Context, client Client, filename string, datumID ...ID) (*Datum, error) {
	id := d.whichID(datumID...)
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return client.UploadDatumFile(ctx, id, file)
}

// UploadImage
func (d *Datum) UploadImage(ctx context.Context, client Client, image []byte, datumID ...ID) (*Datum, error) {
	id := d.whichID(datumID...)
	return client.UploadDatumImage(ctx, id, image)
}

// UploadImageByFile
func (d *Datum) UploadImageByFile(ctx context.Context, client Client, filename string, datumID ...ID) (*Datum, error) {
	id := d.whichID(datumID...)
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return client.UploadDatumImage(ctx, id, file)
}

// UploadVideo
func (d *Datum) UploadVideo(ctx context.Context, client Client, video []byte, datumID ...ID) (*Datum, error) {
	id := d.whichID(datumID...)
	return client.UploadDatumVideo(ctx, id, video)
}

// UploadVideoByFile
func (d *Datum) UploadVideoByFile(ctx context.Context, client Client, filename string, datumID ...ID) (*Datum, error) {
	id := d.whichID(datumID...)
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return client.UploadDatumVideo(ctx, id, file)
}
