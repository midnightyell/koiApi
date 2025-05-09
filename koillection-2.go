package koillection

import (
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
	Visibility       string     `json:"visibility,omitempty"`       // Read and write
	ParentVisibility *string    `json:"parentVisibility,omitempty"` // Read-only
	FinalVisibility  string     `json:"finalVisibility,omitempty"`  // Read-only
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
	Visibility           string     `json:"visibility,omitempty"`           // Read and write
	ParentVisibility     *string    `json:"parentVisibility,omitempty"`     // Read-only
	FinalVisibility      string     `json:"finalVisibility,omitempty"`      // Read-only
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
	DatumType           string     `json:"type"`                          // Read and write
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
	Visibility          string     `json:"visibility,omitempty"`          // Read and write
	ParentVisibility    *string    `json:"parentVisibility,omitempty"`    // Read-only
	FinalVisibility     string     `json:"finalVisibility,omitempty"`     // Read-only
	CreatedAt           time.Time  `json:"createdAt"`                     // Read-only
	UpdatedAt           *time.Time `json:"updatedAt,omitempty"`           // Read-only
	FileImage           *string    `json:"fileImage,omitempty"`           // Write-only, binary data via multipart form
	FileFile            *string    `json:"fileFile,omitempty"`            // Write-only, binary data via multipart form
	FileVideo           *string    `json:"fileVideo,omitempty"`           // Write-only, binary data via multipart form
}

// Field represents a template field in Koillection, combining read and write fields (aligned with Field.jsonld-field.read and Field.jsonld-field.write).
type Field struct {
	Context    *Context `json:"@context,omitempty"`   // JSON-LD only
	ID         ID       `json:"@id,omitempty"`        // JSON-LD only (maps to "id" in JSON, read-only)
	Type       string   `json:"@type,omitempty"`      // JSON-LD only
	Name       string   `json:"name"`                 // Read and write
	Position   int      `json:"position"`             // Read and write
	FieldType  string   `json:"type"`                 // Read and write
	ChoiceList *IRI     `json:"choiceList,omitempty"` // Read and write
	Template   *IRI     `json:"template"`             // Read and write
	Visibility string   `json:"visibility,omitempty"` // Read and write
	Owner      *IRI     `json:"owner,omitempty"`      // Read-only
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
	Quantity            int        `json:"quantity"`                      // Read and write
	Collection          *IRI       `json:"collection"`                    // Read and write
	Owner               *IRI       `json:"owner,omitempty"`               // Read-only
	Image               *string    `json:"image,omitempty"`               // Read-only
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty"` // Read-only
	ImageLargeThumbnail *string    `json:"imageLargeThumbnail,omitempty"` // Read-only
	SeenCounter         int        `json:"seenCounter,omitempty"`         // Read-only
	Visibility          string     `json:"visibility,omitempty"`          // Read and write
	ParentVisibility    *string    `json:"parentVisibility,omitempty"`    // Read-only
	FinalVisibility     string     `json:"finalVisibility,omitempty"`     // Read-only
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
	Visibility          string     `json:"visibility,omitempty"`          // Read and write
	ParentVisibility    *string    `json:"parentVisibility,omitempty"`    // Read-only
	FinalVisibility     string     `json:"finalVisibility,omitempty"`     // Read-only
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
	Visibility          string     `json:"visibility,omitempty"`          // Read and write
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
	PlainPassword                *string    `json:"plainPassword,omitempty"`      // Read and write (assumed)
	Avatar                       *string    `json:"avatar,omitempty"`             // Read and write (assumed)
	Currency                     string     `json:"currency"`                     // Read and write (assumed)
	Locale                       string     `json:"locale"`                       // Read and write (assumed)
	Timezone                     string     `json:"timezone"`                     // Read and write (assumed)
	DateFormat                   string     `json:"dateFormat"`                   // Read and write (assumed)
	DiskSpaceAllowed             int        `json:"diskSpaceAllowed"`             // Read and write (assumed)
	Visibility                   string     `json:"visibility"`                   // Read and write (assumed)
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
	Visibility          string     `json:"visibility,omitempty"`          // Read and write
	ParentVisibility    *string    `json:"parentVisibility,omitempty"`    // Read-only
	FinalVisibility     string     `json:"finalVisibility,omitempty"`     // Read-only
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
	Visibility       string     `json:"visibility,omitempty"`       // Read and write
	ParentVisibility *string    `json:"parentVisibility,omitempty"` // Read-only
	FinalVisibility  string     `json:"finalVisibility,omitempty"`  // Read-only
	CreatedAt        time.Time  `json:"createdAt"`                  // Read-only
	UpdatedAt        *time.Time `json:"updatedAt,omitempty"`        // Read-only
	File             *string    `json:"file,omitempty"`             // Write-only, binary data via multipart form
	DeleteImage      *bool      `json:"deleteImage,omitempty"`      // Write-only
}
