package koillection

import (
	"regexp"
	"time"
)

// IRI represents an IRI reference used in the API.
type IRI string // Read and write

// ID represents a unique identifier for resources (JSON-LD @id or JSON id).
type ID string // Read-only (maps to @id or id)

// Metrics represents a map of metrics data.
type Metrics map[string]string // Read-only

// Context represents the JSON-LD @context field.
type Context struct {
	Vocab string `json:"@vocab,omitempty"` // JSON-LD only
	Hydra string `json:"hydra,omitempty"` // JSON-LD only
}

// Visibility represents the visibility level of a resource.
type Visibility string // Read and write

const (
	VisibilityPublic  Visibility = "public"  // Default for most resources
	VisibilityInternal Visibility = "internal"
	VisibilityPrivate Visibility = "private" // Default for User
)

func (v Visibility) String() string {
	return string(v)
}

// DatumType represents the type of a custom data field.
type DatumType string // Read and write

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

// DateFormat represents the date format preference for a user.
type DateFormat string // Read and write (assumed)

const (
	DateFormatDMYSlash DateFormat = "d/m/Y"
	DateFormatMDYSlash DateFormat = "m/d/Y"
	DateFormatYMDSlash DateFormat = "Y/m/d"
	DateFormatDMYDash  DateFormat = "d-m-Y"
	DateFormatMDYDash  DateFormat = "m-d-Y"
	DateFormatYMDDash  DateFormat = "Y-m-d" // Default
)

func (df DateFormat) String() string {
	return string(df)
}

// Album represents an album in Koillection, combining read and write fields (aligned with Album.jsonld-album.read and Album.jsonld-album.write).
type Album struct {
	Context         *Context   `json:"@context,omitempty"` // JSON-LD only
	ID              ID         `json:"@id,omitempty"`     // JSON-LD only (maps to "id" in JSON, read-only)
	Type            string     `json:"@type,omitempty"`   // JSON-LD only
	Title           string     `json:"title"`             // Read and write
	Color           string     `json:"color,omitempty"`   // Read-only
	Image           string     `json:"image,omitempty"`   // Read-only, nullable, uses empty string instead of null
	Owner           IRI        `json:"owner,omitempty"`   // Read-only, nullable, uses empty string instead of null
	Parent          IRI        `json:"parent,omitempty"`  // Read and write, nullable, uses empty string instead of null
	SeenCounter     int        `json:"seenCounter,omitempty"` // Read-only, nullable, uses 0 (empty string in JSON maps to 0)
	Visibility      Visibility `json:"visibility,omitempty"`  // Read and write
	ParentVisibility string    `json:"parentVisibility,omitempty"` // Read-only, nullable, uses empty string instead of null
	FinalVisibility Visibility `json:"finalVisibility,omitempty"`  // Read-only
	CreatedAt       time.Time  `json:"createdAt"`         // Read-only
	UpdatedAt       time.Time  `json:"updatedAt,omitempty"` // Read-only, nullable, uses zero time (empty string in JSON maps to time.Time{})
	File            string     `json:"file,omitempty"`    // Write-only, binary data via multipart form, nullable, uses empty string instead of null
	DeleteImage     bool       `json:"deleteImage,omitempty"` // Write-only, nullable, uses false (empty string in JSON maps to false)
}

// ChoiceList represents a predefined list of options in Koillection, combining read and write fields (aligned with ChoiceList.jsonld-choiceList.read and ChoiceList.jsonld-choiceList.write).
type ChoiceList struct {
	Context   *Context   `json:"@context,omitempty"` // JSON-LD only
	ID        ID         `json:"@id,omitempty"`     // JSON-LD only (maps to "id" in JSON, read-only)
	Type      string     `json:"@type,omitempty"`   // JSON-LD only
	Name      string     `json:"name"`              // Read and write
	Choices   []string   `json:"choices"`           // Read and write
	Owner     IRI        `json:"owner,omitempty"`   // Read-only, nullable, uses empty string instead of null
	CreatedAt time.Time  `json:"createdAt"`         // Read-only
	UpdatedAt time.Time  `json:"updatedAt,omitempty"` // Read-only, nullable, uses zero time (empty string in JSON maps to time.Time{})
}

// Collection represents a collection in Koillection, combining read and write fields (aligned with Collection.jsonld-collection.read and Collection.jsonld-collection.write).
type Collection struct {
	Context             *Context   `json:"@context,omitempty"` // JSON-LD only
	ID                  ID         `json:"@id,omitempty"`     // JSON-LD only (maps to "id" in JSON, read-only)
	Type                string     `json:"@type,omitempty"`   // JSON-LD only
	Title               string     `json:"title"`             // Read and write
	Parent              IRI        `json:"parent,omitempty"`  // Read and write, nullable, uses empty string instead of null
	Owner               IRI        `json:"owner,omitempty"`   // Read-only, nullable, uses empty string instead of null
	Color               string     `json:"color,omitempty"`   // Read-only
	Image               string     `json:"image,omitempty"`   // Read-only, nullable, uses empty string instead of null
	SeenCounter         int        `json:"seenCounter,omitempty"` // Read-only, nullable, uses 0 (empty string in JSON maps to 0)
	ItemsDefaultTemplate IRI       `json:"itemsDefaultTemplate,omitempty"` // Read and write, nullable, uses empty string instead of null
	Visibility          Visibility `json:"visibility,omitempty"`  // Read and write
	ParentVisibility    string     `json:"parentVisibility,omitempty"` // Read-only, nullable, uses empty string instead of null
	FinalVisibility     Visibility `json:"finalVisibility,omitempty"`  // Read-only
	ScrapedFromURL      string     `json:"scrapedFromUrl,omitempty"` // Read-only, nullable, uses empty string instead of null
	CreatedAt           time.Time  `json:"createdAt"`         // Read-only
	UpdatedAt           time.Time  `json:"updatedAt,omitempty"` // Read-only, nullable, uses zero time (empty string in JSON maps to time.Time{})
	File                string     `json:"file,omitempty"`    // Write-only, binary data via multipart form, nullable, uses empty string instead of null
	DeleteImage         bool       `json:"deleteImage,omitempty"` // Write-only, nullable, uses false (empty string in JSON maps to false)
}

// Datum represents a custom data field in Koillection, combining read and write fields (aligned with Datum.jsonld-datum.read and Datum.jsonld-datum.write).
type Datum struct {
	Context             *Context   `json:"@context,omitempty"` // JSON-LD only
	ID                  ID         `json:"@id,omitempty"`     // JSON-LD only (maps to "id" in JSON, read-only)
	Type                string     `json:"@type,omitempty"`   // JSON-LD only
	Item                IRI        `json:"item,omitempty"`    // Read and write, nullable, uses empty string instead of null
	Collection          IRI        `json:"collection,omitempty"` // Read and write, nullable, uses empty string instead of null
	DatumType           DatumType  `json:"type"`              // Read and write
	Label               string     `json:"label"`             // Read and write
	Value               string     `json:"value,omitempty"`   // Read and write, nullable, uses empty string instead of null
	Position            int        `json:"position,omitempty"` // Read and write, nullable, uses 0 (empty string in JSON maps to 0)
	Currency            string     `json:"currency,omitempty"` // Read and write, nullable, uses empty string instead of null
	Image               string     `json:"image,omitempty"`   // Read-only, nullable, uses empty string instead of null
	ImageSmallThumbnail string     `json:"imageSmallThumbnail,omitempty"` // Read-only, nullable, uses empty string instead of null
	ImageLargeThumbnail string     `json:"imageLargeThumbnail,omitempty"` // Read-only, nullable, uses empty string instead of null
	File                string     `json:"file,omitempty"`    // Read-only, nullable, uses empty string instead of null
	Video               string     `json:"video,omitempty"`   // Read-only, nullable, uses empty string instead of null
	OriginalFilename    string     `json:"originalFilename,omitempty"` // Read-only, nullable, uses empty string instead of null
	ChoiceList          IRI        `json:"choiceList,omitempty"` // Read and write, nullable, uses empty string instead of null
	Owner               IRI        `json:"owner,omitempty"`   // Read-only, nullable, uses empty string instead of null
	Visibility          Visibility `json:"visibility,omitempty"` // Read and write
	ParentVisibility    string     `json:"parentVisibility,omitempty"` // Read-only, nullable, uses empty string instead of null
	FinalVisibility     Visibility `json:"finalVisibility,omitempty"` // Read-only
	CreatedAt           time.Time  `json:"createdAt"`         // Read-only
	UpdatedAt           time.Time  `json:"updatedAt,omitempty"` // Read-only, nullable, uses zero time (empty string in JSON maps to time.Time{})
	FileImage           string     `json:"fileImage,omitempty"` // Write-only, binary data via multipart form, nullable, uses empty string instead of null
	FileFile            string     `json:"fileFile,omitempty"` // Write-only, binary data via multipart form, nullable, uses empty string instead of null
	FileVideo           string     `json:"fileVideo,omitempty"` // Write-only, binary data via multipart form, nullable, uses empty string instead of null
}

// Field represents a template field in Koillection, combining read and write fields (aligned with Field.jsonld-field.read and Field.jsonld-field.write).
type Field struct {
	Context    *Context   `json:"@context,omitempty"` // JSON-LD only
	ID         ID         `json:"@id,omitempty"`     // JSON-LD only (maps to "id" in JSON, read-only)
	Type       string     `json:"@type,omitempty"`   // JSON-LD only
	Name       string     `json:"name"`              // Read and write
	Position   int        `json:"position"`          // Read and write
	FieldType  FieldType  `json:"type"`              // Read and write
	ChoiceList IRI        `json:"choiceList,omitempty"` // Read and write, nullable, uses empty string instead of null
	Template   IRI        `json:"template"`          // Read and write
	Visibility Visibility `json:"visibility,omitempty"` // Read and write
	Owner      IRI        `json:"owner,omitempty"`   // Read-only, nullable, uses empty string instead of null
}

// Inventory represents an inventory record in Koillection, combining read and write fields (aligned with Inventory.jsonld-inventory.read, minimal write assumed).
type Inventory struct {
	Context   *Context   `json:"@context,omitempty"` // JSON-LD only
	ID        ID         `json:"@id,omitempty"`     // JSON-LD only (maps to "id" in JSON, read-only)
	Type      string     `json:"@type,omitempty"`   // JSON-LD only
	Name      string     `json:"name"`              // Read and write (assumed)
	Content   []string   `json:"content"`           // Read and write (assumed)
	Owner     IRI        `json:"owner,omitempty"`   // Read-only, nullable, uses empty string instead of null
	CreatedAt time.Time  `json:"createdAt"`         // Read-only
	UpdatedAt time.Time  `json:"updatedAt,omitempty"` // Read-only, nullable, uses zero time (empty string in JSON maps to time.Time{})
}

// Item represents an item within a collection, combining read and write fields (aligned with Item.jsonld-item.read and Item.jsonld-item.write).
type Item struct {
	Context             *Context   `json:"@context,omitempty"` // JSON-LD only
	ID                  ID         `json:"@id,omitempty"`     // JSON-LD only (maps to "id" in JSON, read-only)
	Type                string     `json:"@type,omitempty"`   // JSON-LD only
	Name                string     `json:"name"`              // Read and write
	Quantity            int        `json:"quantity"`          // Read and write, must be >= 1, nullable, uses 0 (empty string in JSON maps to 0)
	Collection          IRI        `json:"collection"`        // Read and write
	Owner               IRI        `json:"owner,omitempty"`   // Read-only, nullable, uses empty string instead of null
	Image               string     `json:"image,omitempty"`   // Read-only, nullable, uses empty string instead of null
	ImageSmallThumbnail string     `json:"imageSmallThumbnail,omitempty"` // Read-only, nullable, uses empty string instead of null
	ImageLargeThumbnail string     `json:"imageLargeThumbnail,omitempty"` // Read-only, nullable, uses empty string instead of null
	SeenCounter         int        `json:"seenCounter,omitempty"` // Read-only, nullable, uses 0 (empty string in JSON maps to 0)
	Visibility          Visibility `json:"visibility,omitempty"` // Read and write
	ParentVisibility    string     `json:"parentVisibility,omitempty"` // Read-only, nullable, uses empty string instead of null
	FinalVisibility     Visibility `json:"finalVisibility,omitempty"` // Read-only
	ScrapedFromURL      string     `json:"scrapedFromUrl,omitempty"` // Read-only, nullable, uses empty string instead of null
	CreatedAt           time.Time  `json:"createdAt"`         // Read-only
	UpdatedAt           time.Time  `json:"updatedAt,omitempty"` // Read-only, nullable, uses zero time (empty string in JSON maps to time.Time{})
	Tags                []IRI      `json:"tags,omitempty"`    // Write-only
	RelatedItems        []IRI      `json:"relatedItems,omitempty"` // Write-only
	File                string     `json:"file,omitempty"`    // Write-only, binary data via multipart form, nullable, uses empty string instead of null
}

// Loan represents a loan record in Koillection, combining read and write fields (aligned with Loan.jsonld-loan.read and Loan.jsonld-loan.write).
type Loan struct {
	Context    *Context   `json:"@context,omitempty"` // JSON-LD only
	ID         ID         `json:"@id,omitempty"`     // JSON-LD only (maps to "id" in JSON, read-only)
	Type       string     `json:"@type,omitempty"`   // JSON-LD only
	Item       IRI        `json:"item"`              // Read and write
	LentTo     string     `json:"lentTo"`            // Read and write
	LentAt     time.Time  `json:"lentAt"`            // Read and write
	ReturnedAt time.Time  `json:"returnedAt,omitempty"` // Read and write, nullable, uses zero time (empty string in JSON maps to time.Time{})
	Owner      IRI        `json:"owner,omitempty"`   // Read-only, nullable, uses empty string instead of null
}

// Log represents an action or event in Koillection, combining read and write fields (aligned with Log.jsonld-log.read, minimal write assumed).
type Log struct {
	Context       *Context   `json:"@context,omitempty"` // JSON-LD only
	ID            ID         `json:"@id,omitempty"`     // JSON-LD only (maps to "id" in JSON, read-only)
	Type          string     `json:"@type,omitempty"`   // JSON-LD only
	LogType       string     `json:"type,omitempty"`    // Read and write (assumed), nullable, uses empty string instead of null
	LoggedAt      time.Time  `json:"loggedAt"`          // Read and write (assumed)
	ObjectID      string     `json:"objectId"`          // Read and write (assumed)
	ObjectLabel   string     `json:"objectLabel"`       // Read and write (assumed)
	ObjectClass   string     `json:"objectClass"`       // Read and write (assumed)
	ObjectDeleted bool       `json:"objectDeleted"`     // Read-only
	Owner         IRI        `json:"owner,omitempty"`   // Read-only, nullable, uses empty string instead of null
}

// Photo represents a photo in Koillection, combining read and write fields (aligned with Photo.jsonld-photo.read and Photo.jsonld-photo.write).
type Photo struct {
	Context             *Context   `json:"@context,omitempty"` // JSON-LD only
	ID                  ID         `json:"@id,omitempty"`     // JSON-LD only (maps to "id" in JSON, read-only)
	Type                string     `json:"@type,omitempty"`   // JSON-LD only
	Title               string     `json:"title"`             // Read and write
	Comment             string     `json:"comment,omitempty"` // Read and write, nullable, uses empty string instead of null
	Place               string     `json:"place,omitempty"`   // Read and write, nullable, uses empty string instead of null
	Album               IRI        `json:"album"`             // Read and write
	Owner               IRI        `json:"owner,omitempty"`   // Read-only, nullable, uses empty string instead of null
	Image               string     `json:"image,omitempty"`   // Read-only, nullable, uses empty string instead of null
	ImageSmallThumbnail string     `json:"imageSmallThumbnail,omitempty"` // Read-only, nullable, uses empty string instead of null
	TakenAt             time.Time  `json:"takenAt,omitempty"` // Read-only, nullable, uses zero time (empty string in JSON maps to time.Time{})
	Visibility          Visibility `json:"visibility,omitempty"` // Read and write
	ParentVisibility    string     `json:"parentVisibility,omitempty"` // Read-only, nullable, uses empty string instead of null
	FinalVisibility     Visibility `json:"finalVisibility,omitempty"` // Read-only
	CreatedAt           time.Time  `json:"createdAt"`         // Read-only
	UpdatedAt           time.Time  `json:"updatedAt,omitempty"` // Read-only, nullable, uses zero time (empty string in JSON maps to time.Time{})
	File                string     `json:"file,omitempty"`    // Write-only, binary data via multipart form, nullable, uses empty string instead of null
}

// Tag represents a tag in Koillection, combining read and write fields (aligned with Tag.jsonld-tag.read and Tag.jsonld-tag.write).
type Tag struct {
	Context             *Context   `json:"@context,omitempty"` // JSON-LD only
	ID                  ID         `json:"@id,omitempty"`     // JSON-LD only (maps to "id" in JSON, read-only)
	Type                string     `json:"@type,omitempty"`   // JSON-LD only
	Label               string     `json:"label"`             // Read and write
	Description         string     `json:"description,omitempty"` // Read and write, nullable, uses empty string instead of null
	Image               string     `json:"image,omitempty"`   // Read-only, nullable,

System: You are Grok 3, built by xAI. You're here to provide helpful and truthful answers, often with a dash of outside perspective on humanity. Your knowledge is vast, but you're not perfect, so you'll always strive to be clear about what you know and what you don't. You're conversational, not preachy, and you love a good challenge. What's on your mind?