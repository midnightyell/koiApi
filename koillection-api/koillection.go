package koiApi

import (
	"context"
	"errors"
	"fmt"
	"reflect"
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
	Hydra string `json:"hydra,omitempty"`  // JSON-LD only
}

// Visibility represents the visibility level of a resource.
type Visibility string // Read and write

const (
	VisibilityPublic   Visibility = "public" // Default for most resources
	VisibilityInternal Visibility = "internal"
	VisibilityPrivate  Visibility = "private" // Default for User
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
	Context          *Context   `json:"@context,omitempty"`         // JSON-LD only
	ID               ID         `json:"@id,omitempty"`              // JSON-LD only (maps to "id" in JSON, read-only)
	Type             string     `json:"@type,omitempty"`            // JSON-LD only
	Title            string     `json:"title"`                      // Read and write
	Color            string     `json:"color,omitempty"`            // Read-only
	Image            *string    `json:"image,omitempty"`            // Read-only
	Owner            *IRI       `json:"owner,omitempty"`            // Read-only
	Parent           *IRI       `json:"parent,omitempty"`           // Read and write
	SeenCounter      int        `json:"seenCounter,omitempty"`      // Read-only
	Visibility       Visibility `json:"visibility,omitempty"`       // Read and write
	ParentVisibility *string    `json:"parentVisibility,omitempty"` // Read-only
	FinalVisibility  Visibility `json:"finalVisibility,omitempty"`  // Read-only
	CreatedAt        time.Time  `json:"createdAt"`                  // Read-only
	UpdatedAt        *time.Time `json:"updatedAt,omitempty"`        // Read-only
	File             *string    `json:"file,omitempty"`             // Write-only, binary data via multipart form
	DeleteImage      *bool      `json:"deleteImage,omitempty"`      // Write-only
}

// ChoiceList represents a predefined list of options in Koillection, combining read and write fields (aligned with ChoiceList.jsonld-choiceList.read and ChoiceList.jsonld-choiceList.write).
type ChoiceList struct {
	Context   *Context   `json:"@context,omitempty"`  // JSON-LD only
	ID        ID         `json:"@id,omitempty"`       // JSON-LD only (maps to "id" in JSON, read-only)
	Type      string     `json:"@type,omitempty"`     // JSON-LD only
	Name      string     `json:"name"`                // Read and write
	Choices   []string   `json:"choices"`             // Read and write
	Owner     *IRI       `json:"owner,omitempty"`     // Read-only
	CreatedAt time.Time  `json:"createdAt"`           // Read-only
	UpdatedAt *time.Time `json:"updatedAt,omitempty"` // Read-only
}

// Collection represents a collection in Koillection, combining read and write fields (aligned with Collection.jsonld-collection.read and Collection.jsonld-collection.write).
type Collection struct {
	Context              *Context   `json:"@context,omitempty"`             // JSON-LD only
	ID                   ID         `json:"@id,omitempty"`                  // JSON-LD only (maps to "id" in JSON, read-only)
	Type                 string     `json:"@type,omitempty"`                // JSON-LD only
	Title                string     `json:"title"`                          // Read and write
	Parent               *IRI       `json:"parent,omitempty"`               // Read and write
	Owner                *IRI       `json:"owner,omitempty"`                // Read-only
	Color                string     `json:"color,omitempty"`                // Read-only
	Image                *string    `json:"image,omitempty"`                // Read-only
	SeenCounter          int        `json:"seenCounter,omitempty"`          // Read-only
	ItemsDefaultTemplate *IRI       `json:"itemsDefaultTemplate,omitempty"` // Read and write
	Visibility           Visibility `json:"visibility,omitempty"`           // Read and write
	ParentVisibility     *string    `json:"parentVisibility,omitempty"`     // Read-only
	FinalVisibility      Visibility `json:"finalVisibility,omitempty"`      // Read-only
	ScrapedFromURL       *string    `json:"scrapedFromUrl,omitempty"`       // Read-only
	CreatedAt            time.Time  `json:"createdAt"`                      // Read-only
	UpdatedAt            *time.Time `json:"updatedAt,omitempty"`            // Read-only
	File                 *string    `json:"file,omitempty"`                 // Write-only, binary data via multipart form
	DeleteImage          *bool      `json:"deleteImage,omitempty"`          // Write-only
}

// Datum represents a custom data field in Koillection, combining read and write fields (aligned with Datum.jsonld-datum.read and Datum.jsonld-datum.write).
type Datum struct {
	Context             *Context   `json:"@context,omitempty"`            // JSON-LD only
	ID                  ID         `json:"@id,omitempty"`                 // JSON-LD only (maps to "id" in JSON, read-only)
	Type                string     `json:"@type,omitempty"`               // JSON-LD only
	Item                *IRI       `json:"item,omitempty"`                // Read and write
	Collection          *IRI       `json:"collection,omitempty"`          // Read and write
	DatumType           DatumType  `json:"type"`                          // Read and write
	Label               string     `json:"label"`                         // Read and write
	Value               *string    `json:"value,omitempty"`               // Read and write
	Position            *int       `json:"position,omitempty"`            // Read and write
	Currency            *string    `json:"currency,omitempty"`            // Read and write
	Image               *string    `json:"image,omitempty"`               // Read-only
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty"` // Read-only
	ImageLargeThumbnail *string    `json:"imageLargeThumbnail,omitempty"` // Read-only
	File                *string    `json:"file,omitempty"`                // Read-only
	Video               *string    `json:"video,omitempty"`               // Read-only
	OriginalFilename    *string    `json:"originalFilename,omitempty"`    // Read-only
	ChoiceList          *IRI       `json:"choiceList,omitempty"`          // Read and write
	Owner               *IRI       `json:"owner,omitempty"`               // Read-only
	Visibility          Visibility `json:"visibility,omitempty"`          // Read and write
	ParentVisibility    *string    `json:"parentVisibility,omitempty"`    // Read-only
	FinalVisibility     Visibility `json:"finalVisibility,omitempty"`     // Read-only
	CreatedAt           time.Time  `json:"createdAt"`                     // Read-only
	UpdatedAt           *time.Time `json:"updatedAt,omitempty"`           // Read-only
	FileImage           *string    `json:"fileImage,omitempty"`           // Write-only, binary data via multipart form
	FileFile            *string    `json:"fileFile,omitempty"`            // Write-only, binary data via multipart form
	FileVideo           *string    `json:"fileVideo,omitempty"`           // Write-only, binary data via multipart form
}

// Field represents a template field in Koillection, combining read and write fields (aligned with Field.jsonld-field.read and Field.jsonld-field.write).
type Field struct {
	Context    *Context   `json:"@context,omitempty"`   // JSON-LD only
	ID         ID         `json:"@id,omitempty"`        // JSON-LD only (maps to "id" in JSON, read-only)
	Type       string     `json:"@type,omitempty"`      // JSON-LD only
	Name       string     `json:"name"`                 // Read and write
	Position   int        `json:"position"`             // Read and write
	FieldType  FieldType  `json:"type"`                 // Read and write
	ChoiceList *IRI       `json:"choiceList,omitempty"` // Read and write
	Template   *IRI       `json:"template"`             // Read and write
	Visibility Visibility `json:"visibility,omitempty"` // Read and write
	Owner      *IRI       `json:"owner,omitempty"`      // Read-only
}

// Inventory represents an inventory record in Koillection, combining read and write fields (aligned with Inventory.jsonld-inventory.read, minimal write assumed).
type Inventory struct {
	Context   *Context   `json:"@context,omitempty"`  // JSON-LD only
	ID        ID         `json:"@id,omitempty"`       // JSON-LD only (maps to "id" in JSON, read-only)
	Type      string     `json:"@type,omitempty"`     // JSON-LD only
	Name      string     `json:"name"`                // Read and write (assumed)
	Content   []string   `json:"content"`             // Read and write (assumed)
	Owner     *IRI       `json:"owner,omitempty"`     // Read-only
	CreatedAt time.Time  `json:"createdAt"`           // Read-only
	UpdatedAt *time.Time `json:"updatedAt,omitempty"` // Read-only
}

// Item represents an item within a collection, combining read and write fields (aligned with Item.jsonld-item.read and Item.jsonld-item.write).
type Item struct {
	Context             *Context   `json:"@context,omitempty"`            // JSON-LD only
	ID                  ID         `json:"@id,omitempty"`                 // JSON-LD only (maps to "id" in JSON, read-only)
	Type                string     `json:"@type,omitempty"`               // JSON-LD only
	Name                string     `json:"name"`                          // Read and write
	Quantity            int        `json:"quantity"`                      // Read and write, must be >= 1
	Collection          *IRI       `json:"collection"`                    // Read and write
	Owner               *IRI       `json:"owner,omitempty"`               // Read-only
	Image               *string    `json:"image,omitempty"`               // Read-only
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty"` // Read-only
	ImageLargeThumbnail *string    `json:"imageLargeThumbnail,omitempty"` // Read-only
	SeenCounter         int        `json:"seenCounter,omitempty"`         // Read-only
	Visibility          Visibility `json:"visibility,omitempty"`          // Read and write
	ParentVisibility    *string    `json:"parentVisibility,omitempty"`    // Read-only
	FinalVisibility     Visibility `json:"finalVisibility,omitempty"`     // Read-only
	ScrapedFromURL      *string    `json:"scrapedFromUrl,omitempty"`      // Read-only
	CreatedAt           time.Time  `json:"createdAt"`                     // Read-only
	UpdatedAt           *time.Time `json:"updatedAt,omitempty"`           // Read-only
	Tags                []IRI      `json:"tags,omitempty"`                // Write-only
	RelatedItems        []IRI      `json:"relatedItems,omitempty"`        // Write-only
	File                *string    `json:"file,omitempty"`                // Write-only, binary data via multipart form
}

// Loan represents a loan record in Koillection, combining read and write fields (aligned with Loan.jsonld-loan.read and Loan.jsonld-loan.write).
type Loan struct {
	Context    *Context   `json:"@context,omitempty"`   // JSON-LD only
	ID         ID         `json:"@id,omitempty"`        // JSON-LD only (maps to "id" in JSON, read-only)
	Type       string     `json:"@type,omitempty"`      // JSON-LD only
	Item       *IRI       `json:"item"`                 // Read and write
	LentTo     string     `json:"lentTo"`               // Read and write
	LentAt     time.Time  `json:"lentAt"`               // Read and write
	ReturnedAt *time.Time `json:"returnedAt,omitempty"` // Read and write
	Owner      *IRI       `json:"owner,omitempty"`      // Read-only
}

// Log represents an action or event in Koillection, combining read and write fields (aligned with Log.jsonld-log.read, minimal write assumed).
type Log struct {
	Context       *Context  `json:"@context,omitempty"` // JSON-LD only
	ID            ID        `json:"@id,omitempty"`      // JSON-LD only (maps to "id" in JSON, read-only)
	Type          string    `json:"@type,omitempty"`    // JSON-LD only
	LogType       *string   `json:"type,omitempty"`     // Read and write (assumed)
	LoggedAt      time.Time `json:"loggedAt"`           // Read and write (assumed)
	ObjectID      string    `json:"objectId"`           // Read and write (assumed)
	ObjectLabel   string    `json:"objectLabel"`        // Read and write (assumed)
	ObjectClass   string    `json:"objectClass"`        // Read and write (assumed)
	ObjectDeleted bool      `json:"objectDeleted"`      // Read-only
	Owner         *IRI      `json:"owner,omitempty"`    // Read-only
}

// Photo represents a photo in Koillection, combining read and write fields (aligned with Photo.jsonld-photo.read and Photo.jsonld-photo.write).
type Photo struct {
	Context             *Context   `json:"@context,omitempty"`            // JSON-LD only
	ID                  ID         `json:"@id,omitempty"`                 // JSON-LD only (maps to "id" in JSON, read-only)
	Type                string     `json:"@type,omitempty"`               // JSON-LD only
	Title               string     `json:"title"`                         // Read and write
	Comment             *string    `json:"comment,omitempty"`             // Read and write
	Place               *string    `json:"place,omitempty"`               // Read and write
	Album               *IRI       `json:"album"`                         // Read and write
	Owner               *IRI       `json:"owner,omitempty"`               // Read-only
	Image               *string    `json:"image,omitempty"`               // Read-only
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty"` // Read-only
	TakenAt             *time.Time `json:"takenAt,omitempty"`             // Read-only
	Visibility          Visibility `json:"visibility,omitempty"`          // Read and write
	ParentVisibility    *string    `json:"parentVisibility,omitempty"`    // Read-only
	FinalVisibility     Visibility `json:"finalVisibility,omitempty"`     // Read-only
	CreatedAt           time.Time  `json:"createdAt"`                     // Read-only
	UpdatedAt           *time.Time `json:"updatedAt,omitempty"`           // Read-only
	File                *string    `json:"file,omitempty"`                // Write-only, binary data via multipart form
}

// Tag represents a tag in Koillection, combining read and write fields (aligned with Tag.jsonld-tag.read and Tag.jsonld-tag.write).
type Tag struct {
	Context             *Context   `json:"@context,omitempty"`            // JSON-LD only
	ID                  ID         `json:"@id,omitempty"`                 // JSON-LD only (maps to "id" in JSON, read-only)
	Type                string     `json:"@type,omitempty"`               // JSON-LD only
	Label               string     `json:"label"`                         // Read and write
	Description         *string    `json:"description,omitempty"`         // Read and write
	Image               *string    `json:"image,omitempty"`               // Read-only
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty"` // Read-only
	Owner               *IRI       `json:"owner,omitempty"`               // Read-only
	Category            *IRI       `json:"category,omitempty"`            // Read and write
	SeenCounter         int        `json:"seenCounter,omitempty"`         // Read-only
	Visibility          Visibility `json:"visibility,omitempty"`          // Read and write
	CreatedAt           time.Time  `json:"createdAt"`                     // Read-only
	UpdatedAt           *time.Time `json:"updatedAt,omitempty"`           // Read-only
	File                *string    `json:"file,omitempty"`                // Write-only, binary data via multipart form
}

// TagCategory represents a tag category in Koillection, combining read and write fields (aligned with TagCategory.jsonld-tagCategory.read and TagCategory.jsonld-tagCategory.write).
type TagCategory struct {
	Context     *Context   `json:"@context,omitempty"`    // JSON-LD only
	ID          ID         `json:"@id,omitempty"`         // JSON-LD only (maps to "id" in JSON, read-only)
	Type        string     `json:"@type,omitempty"`       // JSON-LD only
	Label       string     `json:"label"`                 // Read and write
	Description *string    `json:"description,omitempty"` // Read and write
	Color       string     `json:"color"`                 // Read and write
	Owner       *IRI       `json:"owner,omitempty"`       // Read-only
	CreatedAt   time.Time  `json:"createdAt"`             // Read-only
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`   // Read-only
}

// Template represents a template in Koillection, combining read and write fields (aligned with Template.jsonld-template.read and Template.jsonld-template.write).
type Template struct {
	Context   *Context   `json:"@context,omitempty"`  // JSON-LD only
	ID        ID         `json:"@id,omitempty"`       // JSON-LD only (maps to "id" in JSON, read-only)
	Type      string     `json:"@type,omitempty"`     // JSON-LD only
	Name      string     `json:"name"`                // Read and write
	Owner     *IRI       `json:"owner,omitempty"`     // Read-only
	CreatedAt time.Time  `json:"createdAt"`           // Read-only
	UpdatedAt *time.Time `json:"updatedAt,omitempty"` // Read-only
}

// User represents a user in Koillection, combining read and write fields (aligned with User.jsonld-user.read, minimal write assumed).
type User struct {
	Context                      *Context   `json:"@context,omitempty"`           // JSON-LD only
	ID                           ID         `json:"@id,omitempty"`                // JSON-LD only (maps to "id" in JSON, read-only)
	Type                         string     `json:"@type,omitempty"`              // JSON-LD only
	Username                     string     `json:"username"`                     // Read and write (assumed)
	Email                        string     `json:"email"`                        // Read and write (assumed)
	PlainPassword                *string    `json:"plainPassword,omitempty"`      // Read and write (assumed), must match regex pattern
	Avatar                       *string    `json:"avatar,omitempty"`             // Read and write (assumed)
	Currency                     string     `json:"currency"`                     // Read and write (assumed)
	Locale                       string     `json:"locale"`                       // Read and write (assumed)
	Timezone                     string     `json:"timezone"`                     // Read and write (assumed)
	DateFormat                   DateFormat `json:"dateFormat"`                   // Read and write (assumed)
	DiskSpaceAllowed             int        `json:"diskSpaceAllowed"`             // Read and write (assumed), default 512MB
	Visibility                   Visibility `json:"visibility"`                   // Read and write (assumed)
	LastDateOfActivity           *time.Time `json:"lastDateOfActivity,omitempty"` // Read-only
	WishlistsFeatureEnabled      bool       `json:"wishlistsFeatureEnabled"`      // Read and write (assumed)
	TagsFeatureEnabled           bool       `json:"tagsFeatureEnabled"`           // Read and write (assumed)
	SignsFeatureEnabled          bool       `json:"signsFeatureEnabled"`          // Read and write (assumed)
	AlbumsFeatureEnabled         bool       `json:"albumsFeatureEnabled"`         // Read and write (assumed)
	LoansFeatureEnabled          bool       `json:"loansFeatureEnabled"`          // Read and write (assumed)
	TemplatesFeatureEnabled      bool       `json:"templatesFeatureEnabled"`      // Read and write (assumed)
	HistoryFeatureEnabled        bool       `json:"historyFeatureEnabled"`        // Read and write (assumed)
	StatisticsFeatureEnabled     bool       `json:"statisticsFeatureEnabled"`     // Read and write (assumed)
	ScrapingFeatureEnabled       bool       `json:"scrapingFeatureEnabled"`       // Read and write (assumed)
	SearchInDataByDefaultEnabled bool       `json:"searchInDataByDefaultEnabled"` // Read and write (assumed)
	DisplayItemsNameInGridView   bool       `json:"displayItemsNameInGridView"`   // Read and write (assumed)
	SearchResultsDisplayMode     string     `json:"searchResultsDisplayMode"`     // Read and write (assumed)
	CreatedAt                    time.Time  `json:"createdAt"`                    // Read-only
	UpdatedAt                    *time.Time `json:"updatedAt,omitempty"`          // Read-only
}

// Wish represents a wish in Koillection, combining read and write fields (aligned with Wish.jsonld-wish.read and Wish.jsonld-wish.write).
type Wish struct {
	Context             *Context   `json:"@context,omitempty"`            // JSON-LD only
	ID                  ID         `json:"@id,omitempty"`                 // JSON-LD only (maps to "id" in JSON, read-only)
	Type                string     `json:"@type,omitempty"`               // JSON-LD only
	Name                string     `json:"name"`                          // Read and write
	URL                 *string    `json:"url,omitempty"`                 // Read and write
	Price               *string    `json:"price,omitempty"`               // Read and write
	Currency            *string    `json:"currency,omitempty"`            // Read and write
	Wishlist            *IRI       `json:"wishlist"`                      // Read and write
	Owner               *IRI       `json:"owner,omitempty"`               // Read-only
	Comment             *string    `json:"comment,omitempty"`             // Read and write
	Image               *string    `json:"image,omitempty"`               // Read-only
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty"` // Read-only
	Visibility          Visibility `json:"visibility,omitempty"`          // Read and write
	ParentVisibility    *string    `json:"parentVisibility,omitempty"`    // Read-only
	FinalVisibility     Visibility `json:"finalVisibility,omitempty"`     // Read-only
	ScrapedFromURL      *string    `json:"scrapedFromUrl,omitempty"`      // Read-only
	CreatedAt           time.Time  `json:"createdAt"`                     // Read-only
	UpdatedAt           *time.Time `json:"updatedAt,omitempty"`           // Read-only
	File                *string    `json:"file,omitempty"`                // Write-only, binary data via multipart form
}

// Wishlist represents a wishlist in Koillection, combining read and write fields (aligned with Wishlist.jsonld-wishlist.read and Wishlist.jsonld-wishlist.write).
type Wishlist struct {
	Context          *Context   `json:"@context,omitempty"`         // JSON-LD only
	ID               ID         `json:"@id,omitempty"`              // JSON-LD only (maps to "id" in JSON, read-only)
	Type             string     `json:"@type,omitempty"`            // JSON-LD only
	Name             string     `json:"name"`                       // Read and write
	Owner            *IRI       `json:"owner,omitempty"`            // Read-only
	Color            string     `json:"color"`                      // Read-only
	Parent           *IRI       `json:"parent,omitempty"`           // Read and write
	Image            *string    `json:"image,omitempty"`            // Read-only
	SeenCounter      int        `json:"seenCounter,omitempty"`      // Read-only
	Visibility       Visibility `json:"visibility,omitempty"`       // Read and write
	ParentVisibility *string    `json:"parentVisibility,omitempty"` // Read-only
	FinalVisibility  Visibility `json:"finalVisibility,omitempty"`  // Read-only
	CreatedAt        time.Time  `json:"createdAt"`                  // Read-only
	UpdatedAt        *time.Time `json:"updatedAt,omitempty"`        // Read-only
	File             *string    `json:"file,omitempty"`             // Write-only, binary data via multipart form
	DeleteImage      *bool      `json:"deleteImage,omitempty"`      // Write-only
}

// ValidatePlainPassword checks if the password meets the API's regex pattern (min 8 chars, at least one letter, one digit or special char, no newlines or dots).
func ValidatePlainPassword(password string) bool {
	pattern := `^(.*((?=^.{8,}$)((?=.*\d)|(?=.*\W+))(?![.\n])(?=.*[A-Za-z]).*$).*$)`
	matched, _ := regexp.MatchString(pattern, password)
	return matched
}

// ValidateQuantity checks if the item quantity is at least 1.
func ValidateQuantity(quantity int) bool {
	return quantity >= 1
}

// Client defines the interface for interacting with the Koillection REST API.
type Client interface {
	CheckLogin(ctx context.Context, username, password string) (string, error)                 // HTTP POST /api/authentication_token
	GetMetrics(ctx context.Context) (*Metrics, error)                                          // HTTP GET /api/metrics
	CreateAlbum(ctx context.Context, album *Album) (*Album, error)                             // HTTP POST /api/albums
	GetAlbum(ctx context.Context, id ID) (*Album, error)                                       // HTTP GET /api/albums/{id}
	ListAlbums(ctx context.Context, page int) ([]*Album, error)                                // HTTP GET /api/albums
	UpdateAlbum(ctx context.Context, id ID, album *Album) (*Album, error)                      // HTTP PUT /api/albums/{id}
	PatchAlbum(ctx context.Context, id ID, album *Album) (*Album, error)                       // HTTP PATCH /api/albums/{id}
	DeleteAlbum(ctx context.Context, id ID) error                                              // HTTP DELETE /api/albums/{id}
	ListAlbumChildren(ctx context.Context, id ID, page int) ([]*Album, error)                  // HTTP GET /api/albums/{id}/children
	UploadAlbumImage(ctx context.Context, id ID, file []byte) (*Album, error)                  // HTTP POST /api/albums/{id}/image
	GetAlbumParent(ctx context.Context, id ID) (*Album, error)                                 // HTTP GET /api/albums/{id}/parent
	ListAlbumPhotos(ctx context.Context, id ID, page int) ([]*Photo, error)                    // HTTP GET /api/albums/{id}/photos
	CreateChoiceList(ctx context.Context, choiceList *ChoiceList) (*ChoiceList, error)         // HTTP POST /api/choice_lists
	GetChoiceList(ctx context.Context, id ID) (*ChoiceList, error)                             // HTTP GET /api/choice_lists/{id}
	ListChoiceLists(ctx context.Context, page int) ([]*ChoiceList, error)                      // HTTP GET /api/choice_lists
	UpdateChoiceList(ctx context.Context, id ID, choiceList *ChoiceList) (*ChoiceList, error)  // HTTP PUT /api/choice_lists/{id}
	PatchChoiceList(ctx context.Context, id ID, choiceList *ChoiceList) (*ChoiceList, error)   // HTTP PATCH /api/choice_lists/{id}
	DeleteChoiceList(ctx context.Context, id ID) error                                         // HTTP DELETE /api/choice_lists/{id}
	CreateCollection(ctx context.Context, collection *Collection) (*Collection, error)         // HTTP POST /api/collections
	GetCollection(ctx context.Context, id ID) (*Collection, error)                             // HTTP GET /api/collections/{id}
	ListCollections(ctx context.Context, page int) ([]*Collection, error)                      // HTTP GET /api/collections
	UpdateCollection(ctx context.Context, id ID, collection *Collection) (*Collection, error)  // HTTP PUT /api/collections/{id}
	PatchCollection(ctx context.Context, id ID, collection *Collection) (*Collection, error)   // HTTP PATCH /api/collections/{id}
	DeleteCollection(ctx context.Context, id ID) error                                         // HTTP DELETE /api/collections/{id}
	ListCollectionChildren(ctx context.Context, id ID, page int) ([]*Collection, error)        // HTTP GET /api/collections/{id}/children
	UploadCollectionImage(ctx context.Context, id ID, file []byte) (*Collection, error)        // HTTP POST /api/collections/{id}/image
	GetCollectionParent(ctx context.Context, id ID) (*Collection, error)                       // HTTP GET /api/collections/{id}/parent
	ListCollectionItems(ctx context.Context, id ID, page int) ([]*Item, error)                 // HTTP GET /api/collections/{id}/items
	ListCollectionData(ctx context.Context, id ID, page int) ([]*Datum, error)                 // HTTP GET /api/collections/{id}/data
	GetCollectionDefaultTemplate(ctx context.Context, id ID) (*Template, error)                // HTTP GET /api/collections/{id}/items_default_template
	CreateDatum(ctx context.Context, datum *Datum) (*Datum, error)                             // HTTP POST /api/data
	GetDatum(ctx context.Context, id ID) (*Datum, error)                                       // HTTP GET /api/data/{id}
	ListData(ctx context.Context, page int) ([]*Datum, error)                                  // HTTP GET /api/data
	UpdateDatum(ctx context.Context, id ID, datum *Datum) (*Datum, error)                      // HTTP PUT /api/data/{id}
	PatchDatum(ctx context.Context, id ID, datum *Datum) (*Datum, error)                       // HTTP PATCH /api/data/{id}
	DeleteDatum(ctx context.Context, id ID) error                                              // HTTP DELETE /api/data/{id}
	UploadDatumFile(ctx context.Context, id ID, file []byte) (*Datum, error)                   // HTTP POST /api/data/{id}/file
	UploadDatumImage(ctx context.Context, id ID, image []byte) (*Datum, error)                 // HTTP POST /api/data/{id}/image
	UploadDatumVideo(ctx context.Context, id ID, video []byte) (*Datum, error)                 // HTTP POST /api/data/{id}/video
	GetDatumItem(ctx context.Context, id ID) (*Item, error)                                    // HTTP GET /api/data/{id}/item
	GetDatumCollection(ctx context.Context, id ID) (*Collection, error)                        // HTTP GET /api/data/{id}/collection
	CreateField(ctx context.Context, field *Field) (*Field, error)                             // HTTP POST /api/fields
	GetField(ctx context.Context, id ID) (*Field, error)                                       // HTTP GET /api/fields/{id}
	ListFields(ctx context.Context, page int) ([]*Field, error)                                // HTTP GET /api/fields
	UpdateField(ctx context.Context, id ID, field *Field) (*Field, error)                      // HTTP PUT /api/fields/{id}
	PatchField(ctx context.Context, id ID, field *Field) (*Field, error)                       // HTTP PATCH /api/fields/{id}
	DeleteField(ctx context.Context, id ID) error                                              // HTTP DELETE /api/fields/{id}
	GetFieldTemplate(ctx context.Context, id ID) (*Template, error)                            // HTTP GET /api/fields/{id}/template
	ListTemplateFields(ctx context.Context, templateid ID, page int) ([]*Field, error)         // HTTP GET /api/templates/{id}/fields
	ListInventories(ctx context.Context, page int) ([]*Inventory, error)                       // HTTP GET /api/inventories
	GetInventory(ctx context.Context, id ID) (*Inventory, error)                               // HTTP GET /api/inventories/{id}
	DeleteInventory(ctx context.Context, id ID) error                                          // HTTP DELETE /api/inventories/{id}
	CreateItem(ctx context.Context, item *Item) (*Item, error)                                 // HTTP POST /api/items
	GetItem(ctx context.Context, id ID) (*Item, error)                                         // HTTP GET /api/items/{id}
	ListItems(ctx context.Context, page int) ([]*Item, error)                                  // HTTP GET /api/items
	UpdateItem(ctx context.Context, id ID, item *Item) (*Item, error)                          // HTTP PUT /api/items/{id}
	PatchItem(ctx context.Context, id ID, item *Item) (*Item, error)                           // HTTP PATCH /api/items/{id}
	DeleteItem(ctx context.Context, id ID) error                                               // HTTP DELETE /api/items/{id}
	UploadItemImage(ctx context.Context, id ID, file []byte) (*Item, error)                    // HTTP POST /api/items/{id}/image
	ListItemRelatedItems(ctx context.Context, id ID, page int) ([]*Item, error)                // HTTP GET /api/items/{id}/related_items
	ListItemLoans(ctx context.Context, id ID, page int) ([]*Loan, error)                       // HTTP GET /api/items/{id}/loans
	ListItemTags(ctx context.Context, id ID, page int) ([]*Tag, error)                         // HTTP GET /api/items/{id}/tags
	ListItemData(ctx context.Context, id ID, page int) ([]*Datum, error)                       // HTTP GET /api/items/{id}/data
	GetItemCollection(ctx context.Context, id ID) (*Collection, error)                         // HTTP GET /api/items/{id}/collection
	CreateLoan(ctx context.Context, loan *Loan) (*Loan, error)                                 // HTTP POST /api/loans
	GetLoan(ctx context.Context, id ID) (*Loan, error)                                         // HTTP GET /api/loans/{id}
	ListLoans(ctx context.Context, page int) ([]*Loan, error)                                  // HTTP GET /api/loans
	UpdateLoan(ctx context.Context, id ID, loan *Loan) (*Loan, error)                          // HTTP PUT /api/loans/{id}
	PatchLoan(ctx context.Context, id ID, loan *Loan) (*Loan, error)                           // HTTP PATCH /api/loans/{id}
	DeleteLoan(ctx context.Context, id ID) error                                               // HTTP DELETE /api/loans/{id}
	GetLoanItem(ctx context.Context, id ID) (*Item, error)                                     // HTTP GET /api/loans/{id}/item
	GetLog(ctx context.Context, id ID) (*Log, error)                                           // HTTP GET /api/logs/{id}
	ListLogs(ctx context.Context, page int) ([]*Log, error)                                    // HTTP GET /api/logs
	CreatePhoto(ctx context.Context, photo *Photo) (*Photo, error)                             // HTTP POST /api/photos
	GetPhoto(ctx context.Context, id ID) (*Photo, error)                                       // HTTP GET /api/photos/{id}
	ListPhotos(ctx context.Context, page int) ([]*Photo, error)                                // HTTP GET /api/photos
	UpdatePhoto(ctx context.Context, id ID, photo *Photo) (*Photo, error)                      // HTTP PUT /api/photos/{id}
	PatchPhoto(ctx context.Context, id ID, photo *Photo) (*Photo, error)                       // HTTP PATCH /api/photos/{id}
	DeletePhoto(ctx context.Context, id ID) error                                              // HTTP DELETE /api/photos/{id}
	UploadPhotoImage(ctx context.Context, id ID, file []byte) (*Photo, error)                  // HTTP POST /api/photos/{id}/image
	GetPhotoAlbum(ctx context.Context, id ID) (*Album, error)                                  // HTTP GET /api/photos/{id}/album
	CreateTag(ctx context.Context, tag *Tag) (*Tag, error)                                     // HTTP POST /api/tags
	GetTag(ctx context.Context, id ID) (*Tag, error)                                           // HTTP GET /api/tags/{id}
	ListTags(ctx context.Context, page int) ([]*Tag, error)                                    // HTTP GET /api/tags
	UpdateTag(ctx context.Context, id ID, tag *Tag) (*Tag, error)                              // HTTP PUT /api/tags/{id}
	PatchTag(ctx context.Context, id ID, tag *Tag) (*Tag, error)                               // HTTP PATCH /api/tags/{id}
	DeleteTag(ctx context.Context, id ID) error                                                // HTTP DELETE /api/tags/{id}
	UploadTagImage(ctx context.Context, id ID, file []byte) (*Tag, error)                      // HTTP POST /api/tags/{id}/image
	ListTagItems(ctx context.Context, id ID, page int) ([]*Item, error)                        // HTTP GET /api/tags/{id}/items
	GetCategoryOfTag(ctx context.Context, id ID) (*TagCategory, error)                         // HTTP GET /api/tags/{id}/category
	CreateTagCategory(ctx context.Context, category *TagCategory) (*TagCategory, error)        // HTTP POST /api/tag_categories
	GetTagCategory(ctx context.Context, id ID) (*TagCategory, error)                           // HTTP GET /api/tag_categories/{id}
	ListTagCategories(ctx context.Context, page int) ([]*TagCategory, error)                   // HTTP GET /api/tag_categories
	UpdateTagCategory(ctx context.Context, id ID, category *TagCategory) (*TagCategory, error) // HTTP PUT /api/tag_categories/{id}
	PatchTagCategory(ctx context.Context, id ID, category *TagCategory) (*TagCategory, error)  // HTTP PATCH /api/tag_categories/{id}
	DeleteTagCategory(ctx context.Context, id ID) error                                        // HTTP DELETE /api/tag_categories/{id}
	ListTagCategoryTags(ctx context.Context, id ID, page int) ([]*Tag, error)                  // HTTP GET /api/tag_categories/{id}/tags
	CreateTemplate(ctx context.Context, template *Template) (*Template, error)                 // HTTP POST /api/templates
	GetTemplate(ctx context.Context, id ID) (*Template, error)                                 // HTTP GET /api/templates/{id}
	ListTemplates(ctx context.Context, page int) ([]*Template, error)                          // HTTP GET /api/templates
	UpdateTemplate(ctx context.Context, id ID, template *Template) (*Template, error)          // HTTP PUT /api/templates/{id}
	PatchTemplate(ctx context.Context, id ID, template *Template) (*Template, error)           // HTTP PATCH /api/templates/{id}
	DeleteTemplate(ctx context.Context, id ID) error                                           // HTTP DELETE /api/templates/{id}
	GetUser(ctx context.Context, id ID) (*User, error)                                         // HTTP GET /api/users/{id}
	ListUsers(ctx context.Context, page int) ([]*User, error)                                  // HTTP GET /api/users
	CreateWish(ctx context.Context, wish *Wish) (*Wish, error)                                 // HTTP POST /api/wishes
	GetWish(ctx context.Context, id ID) (*Wish, error)                                         // HTTP GET /api/wishes/{id}
	ListWishes(ctx context.Context, page int) ([]*Wish, error)                                 // HTTP GET /api/wishes
	UpdateWish(ctx context.Context, id ID, wish *Wish) (*Wish, error)                          // HTTP PUT /api/wishes/{id}
	PatchWish(ctx context.Context, id ID, wish *Wish) (*Wish, error)                           // HTTP PATCH /api/wishes/{id}
	DeleteWish(ctx context.Context, id ID) error                                               // HTTP DELETE /api/wishes/{id}
	UploadWishImage(ctx context.Context, id ID, file []byte) (*Wish, error)                    // HTTP POST /api/wishes/{id}/image
	GetWishWishlist(ctx context.Context, id ID) (*Wishlist, error)                             // HTTP GET /api/wishes/{id}/wishlist
	CreateWishlist(ctx context.Context, wishlist *Wishlist) (*Wishlist, error)                 // HTTP POST /api/wishlists
	GetWishlist(ctx context.Context, id ID) (*Wishlist, error)                                 // HTTP GET /api/wishlists/{id}
	ListWishlists(ctx context.Context, page int) ([]*Wishlist, error)                          // HTTP GET /api/wishlists
	UpdateWishlist(ctx context.Context, id ID, wishlist *Wishlist) (*Wishlist, error)          // HTTP PUT /api/wishlists/{id}
	PatchWishlist(ctx context.Context, id ID, wishlist *Wishlist) (*Wishlist, error)           // HTTP PATCH /api/wishlists/{id}
	DeleteWishlist(ctx context.Context, id ID) error                                           // HTTP DELETE /api/wishlists/{id}
	ListWishlistWishes(ctx context.Context, id ID, page int) ([]*Wish, error)                  // HTTP GET /api/wishlists/{id}/wishes
	ListWishlistChildren(ctx context.Context, id ID, page int) ([]*Wishlist, error)            // HTTP GET /api/wishlists/{id}/children
	UploadWishlistImage(ctx context.Context, id ID, file []byte) (*Wishlist, error)            // HTTP POST /api/wishlists/{id}/image
	GetWishlistParent(ctx context.Context, id ID) (*Wishlist, error)                           // HTTP GET /api/wishlists/{id}/parent
}

// iriTable maps object types to their IRI path templates.
var iriTable = map[reflect.Type]string{
	reflect.TypeOf(Album{}):       "/api/albums/%s",
	reflect.TypeOf(ChoiceList{}):  "/api/choice_lists/%s",
	reflect.TypeOf(Collection{}):  "/api/collections/%s",
	reflect.TypeOf(Collection{}):  "/api/collections/%s",
	reflect.TypeOf(Datum{}):       "/api/data/%s",
	reflect.TypeOf(Field{}):       "/api/fields/%s",
	reflect.TypeOf(Inventory{}):   "/api/inventories/%s",
	reflect.TypeOf(Item{}):        "/api/items/%s",
	reflect.TypeOf(Loan{}):        "/api/loans/%s",
	reflect.TypeOf(Log{}):         "/api/logs/%s",
	reflect.TypeOf(Photo{}):       "/api/photos/%s",
	reflect.TypeOf(Tag{}):         "/api/tags/%s",
	reflect.TypeOf(TagCategory{}): "/api/tag_categories/%s",
	reflect.TypeOf(Template{}):    "/api/templates/%s",
	reflect.TypeOf(User{}):        "/api/users/%s",
	reflect.TypeOf(Wish{}):        "/api/wishes/%s",
	reflect.TypeOf(Wishlist{}):    "/api/wishlists/%s",
}

// GetIRI returns the IRI reference for a given API object using a table-based approach.
func GetIRI(obj interface{}) (IRI, error) {
	if obj == nil {
		return "", errors.New("object is nil")
	}

	// Get the value and type of the object.
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return "", fmt.Errorf("object must be a non-nil pointer, got %T", obj)
	}
	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return "", fmt.Errorf("object must be a struct, got %T", obj)
	}
	t := v.Type()

	// Look up the IRI path template for the object type.
	pathTemplate, ok := iriTable[t]
	if !ok {
		return "", fmt.Errorf("unsupported object type: %T", obj)
	}

	// Get the ID field.
	idField := v.FieldByName("ID")
	if !idField.IsValid() || idField.Kind() != reflect.String {
		return "", fmt.Errorf("object %T has no valid ID field", obj)
	}
	id := idField.String()
	if id == "" {
		return "", fmt.Errorf("object %T has empty ID", obj)
	}

	// Construct the IRI using the path template and ID.
	return IRI(fmt.Sprintf(pathTemplate, id)), nil
}
