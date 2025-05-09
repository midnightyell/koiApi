package koillection

import (
	"time"
)

// IRI represents an IRI reference used in the API.
type IRI string

// ID represents a unique identifier for resources (JSON-LD @id or JSON id).
type ID string

// Metrics represents a map of metrics data.
type Metrics map[string]string

// Context represents the JSON-LD @context field.
type Context struct {
	Vocab string `json:"@vocab,omitempty"`
	Hydra string `json:"hydra,omitempty"`
}

// Album represents an album in Koillection, combining read and write fields (aligned with Album.jsonld-album.read and Album.jsonld-album.write).
type Album struct {
	Context         *Context   `json:"@context,omitempty"` // JSON-LD only
	ID              ID         `json:"@id,omitempty"`     // JSON-LD @id, maps to "id" in JSON
	Type            string     `json:"@type,omitempty"`   // JSON-LD only
	Title           string     `json:"title"`
	Color           string     `json:"color,omitempty"`
	Image           *string    `json:"image,omitempty"`
	Owner           *IRI       `json:"owner,omitempty"`
	Parent          *IRI       `json:"parent,omitempty"`
	SeenCounter     int        `json:"seenCounter,omitempty"`
	Visibility      string     `json:"visibility,omitempty"`
	ParentVisibility *string   `json:"parentVisibility,omitempty"`
	FinalVisibility string     `json:"finalVisibility,omitempty"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       *time.Time `json:"updatedAt,omitempty"`
	File            *string    `json:"file,omitempty"`        // Write-only, binary data via multipart form
	DeleteImage     *bool      `json:"deleteImage,omitempty"` // Write-only
}

// ChoiceList represents a predefined list of options in Koillection, combining read and write fields (aligned with ChoiceList.jsonld-choiceList.read and ChoiceList.jsonld-choiceList.write).
type ChoiceList struct {
	Context   *Context   `json:"@context,omitempty"` // JSON-LD only
	ID        ID         `json:"@id,omitempty"`     // JSON-LD @id, maps to "id" in JSON
	Type      string     `json:"@type,omitempty"`   // JSON-LD only
	Name      string     `json:"name"`
	Choices   []string   `json:"choices"`
	Owner     *IRI       `json:"owner,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// Collection represents a collection in Koillection, combining read and write fields (aligned with Collection.jsonld-collection.read and Collection.jsonld-collection.write).
type Collection struct {
	Context             *Context   `json:"@context,omitempty"` // JSON-LD only
	ID                  ID         `json:"@id,omitempty"`     // JSON-LD @id, maps to "id" in JSON
	Type                string     `json:"@type,omitempty"`   // JSON-LD only
	Title               string     `json:"title"`
	Parent              *IRI       `json:"parent,omitempty"`
	Owner               *IRI       `json:"owner,omitempty"`
	Color               string     `json:"color,omitempty"`
	Image               *string    `json:"image,omitempty"`
	SeenCounter         int        `json:"seenCounter,omitempty"`
	ItemsDefaultTemplate *IRI      `json:"itemsDefaultTemplate,omitempty"`
	Visibility          string     `json:"visibility,omitempty"`
	ParentVisibility    *string    `json:"parentVisibility,omitempty"`
	FinalVisibility     string     `json:"finalVisibility,omitempty"`
	ScrapedFromURL      *string    `json:"scrapedFromUrl,omitempty"`
	CreatedAt           time.Time  `json:"createdAt"`
	UpdatedAt           *time.Time `json:"updatedAt,omitempty"`
	File                *string    `json:"file,omitempty"`        // Write-only, binary data via multipart form
	DeleteImage         *bool      `json:"deleteImage,omitempty"` // Write-only
}

// Datum represents a custom data field in Koillection, combining read and write fields (aligned with Datum.jsonld-datum.read and Datum.jsonld-datum.write).
type Datum struct {
	Context             *Context   `json:"@context,omitempty"` // JSON-LD only
	ID                  ID         `json:"@id,omitempty"`     // JSON-LD @id, maps to "id" in JSON
	Type                string     `json:"@type,omitempty"`   // JSON-LD only
	Item                *IRI       `json:"item,omitempty"`
	Collection          *IRI       `json:"collection,omitempty"`
	DatumType           string     `json:"type"`
	Label               string     `json:"label"`
	Value               *string    `json:"value,omitempty"`
	Position            *int       `json:"position,omitempty"`
	Currency            *string    `json:"currency,omitempty"`
	Image               *string    `json:"image,omitempty"`
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty"`
	ImageLargeThumbnail *string    `json:"imageLargeThumbnail,omitempty"`
	File                *string    `json:"file,omitempty"`
	Video               *string    `json:"video,omitempty"`
	OriginalFilename    *string    `json:"originalFilename,omitempty"`
	ChoiceList          *IRI       `json:"choiceList,omitempty"`
	Owner               *IRI       `json:"owner,omitempty"`
	Visibility          string     `json:"visibility,omitempty"`
	ParentVisibility    *string    `json:"parentVisibility,omitempty"`
	FinalVisibility     string     `json:"finalVisibility,omitempty"`
	CreatedAt           time.Time  `json:"createdAt"`
	UpdatedAt           *time.Time `json:"updatedAt,omitempty"`
	FileImage           *string    `json:"fileImage,omitempty"` // Write-only, binary data via multipart form
	FileFile            *string    `json:"fileFile,omitempty"`  // Write-only, binary data via multipart form
	FileVideo           *string    `json:"fileVideo,omitempty"` // Write-only, binary data via multipart form
}

// Field represents a template field in Koillection, combining read and write fields (aligned with Field.jsonld-field.read and Field.jsonld-field.write).
type Field struct {
	Context    *Context   `json:"@context,omitempty"` // JSON-LD only
	ID         ID         `json:"@id,omitempty"`     // JSON-LD @id, maps to "id" in JSON
	Type       string     `json:"@type,omitempty"`   // JSON-LD only
	Name       string     `json:"name"`
	Position   int        `json:"position"`
	FieldType  string     `json:"type"`
	ChoiceList *IRI       `json:"choiceList,omitempty"`
	Template   *IRI       `json:"template"`
	Visibility string     `json:"visibility,omitempty"`
	Owner      *IRI       `json:"owner,omitempty"`
}

// Inventory represents an inventory record in Koillection, combining read and write fields (aligned with Inventory.jsonld-inventory.read, minimal write assumed).
type Inventory struct {
	Context   *Context   `json:"@context,omitempty"` // JSON-LD only
	ID        ID         `json:"@id,omitempty"`     // JSON-LD @id, maps to "id" in JSON
	Type      string     `json:"@type,omitempty"`   // JSON-LD only
	Name      string     `json:"name"`
	Content   []string   `json:"content"`
	Owner     *IRI       `json:"owner,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// Item represents an item within a collection, combining read and write fields (aligned with Item.jsonld-item.read and Item.jsonld-item.write).
type Item struct {
	Context             *Context   `json:"@context,omitempty"` // JSON-LD only
	ID                  ID         `json:"@id,omitempty"`     // JSON-LD @id, maps to "id" in JSON
	Type                string     `json:"@type,omitempty"`   // JSON-LD only
	Name                string     `json:"name"`
	Quantity            int        `json:"quantity"`
	Collection          *IRI       `json:"collection"`
	Owner               *IRI       `json:"owner,omitempty"`
	Image               *string    `json:"image,omitempty"`
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty"`
	ImageLargeThumbnail *string    `json:"imageLargeThumbnail,omitempty"`
	SeenCounter         int        `json:"seenCounter,omitempty"`
	Visibility          string     `json:"visibility,omitempty"`
	ParentVisibility    *string    `json:"parentVisibility,omitempty"`
	FinalVisibility     string     `json:"finalVisibility,omitempty"`
	ScrapedFromURL      *string    `json:"scrapedFromUrl,omitempty"`
	CreatedAt           time.Time  `json:"createdAt"`
	UpdatedAt           *time.Time `json:"updatedAt,omitempty"`
	Tags                []IRI      `json:"tags,omitempty"`        // Write-only
	RelatedItems        []IRI      `json:"relatedItems,omitempty"` // Write-only
	File                *string    `json:"file,omitempty"`        // Write-only, binary data via multipart form
}

// Loan represents a loan record in Koillection, combining read and write fields (aligned with Loan.jsonld-loan.read and Loan.jsonld-loan.write).
type Loan struct {
	Context    *Context   `json:"@context,omitempty"` // JSON-LD only
	ID         ID         `json:"@id,omitempty"`     // JSON-LD @id, maps to "id" in JSON
	Type       string     `json:"@type,omitempty"`   // JSON-LD only
	Item       *IRI       `json:"item"`
	LentTo     string     `json:"lentTo"`
	LentAt     time.Time  `json:"lentAt"`
	ReturnedAt *time.Time `json:"returnedAt,omitempty"`
	Owner      *IRI       `json:"owner,omitempty"`
}

// Log represents an action or event in Koillection, combining read and write fields (aligned with Log.jsonld-log.read, minimal write assumed).
type Log struct {
	Context       *Context   `json:"@context,omitempty"` // JSON-LD only
	ID            ID         `json:"@id,omitempty"`     // JSON-LD @id, maps to "id" in JSON
	Type          string     `json:"@type,omitempty"`   // JSON-LD only
	LogType       *string    `json:"type,omitempty"`
	LoggedAt      time.Time  `json:"loggedAt"`
	ObjectID      string     `json:"objectId"`
	ObjectLabel   string     `json:"objectLabel"`
	ObjectClass   string     `json:"objectClass"`
	ObjectDeleted bool       `json:"objectDeleted"`
	Owner         *IRI       `json:"owner,omitempty"`
}

// Photo represents a photo in Koillection, combining read and write fields (aligned with Photo.jsonld-photo.read and Photo.jsonld-photo.write).
type Photo struct {
	Context             *Context   `json:"@context,omitempty"` // JSON-LD only
	ID                  ID         `json:"@id,omitempty"`     // JSON-LD @id, maps to "id" in JSON
	Type                string     `json:"@type,omitempty"`   // JSON-LD only
	Title               string     `json:"title"`
	Comment             *string    `json:"comment,omitempty"`
	Place               *string    `json:"place,omitempty"`
	Album               *IRI       `json:"album"`
	Owner               *IRI       `json:"owner,omitempty"`
	Image               *string    `json:"image,omitempty"`
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty"`
	TakenAt             *time.Time `json:"takenAt,omitempty"`
	Visibility          string     `json:"visibility,omitempty"`
	ParentVisibility    *string    `json:"parentVisibility,omitempty"`
	FinalVisibility     string     `json:"finalVisibility,omitempty"`
	CreatedAt           time.Time  `json:"createdAt"`
	UpdatedAt           *time.Time `json:"updatedAt,omitempty"`
	File                *string    `json:"file,omitempty"` // Write-only, binary data via multipart form
}

// Tag represents a tag in Koillection, combining read and write fields (aligned with Tag.jsonld-tag.read and Tag.jsonld-tag.write).
type Tag struct {
	Context             *Context   `json:"@context,omitempty"` // JSON-LD only
	ID                  ID         `json:"@id,omitempty"`     // JSON-LD @id, maps to "id" in JSON
	Type                string     `json:"@type,omitempty"`   // JSON-LD only
	Label               string     `json:"label"`
	Description         *string    `json:"description,omitempty"`
	Image               *string    `json:"image,omitempty"`
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty"`
	Owner               *IRI       `json:"owner,omitempty"`
	Category            *IRI       `json:"category,omitempty"`
	SeenCounter         int        `json:"seenCounter,omitempty"`
	Visibility          string     `json:"visibility,omitempty"`
	CreatedAt           time.Time  `json:"createdAt"`
	UpdatedAt           *time.Time `json:"updatedAt,omitempty"`
	File                *string    `json:"file,omitempty"` // Write-only, binary data via multipart form
}

// TagCategory represents a tag category in Koillection, combining read and write fields (aligned with TagCategory.jsonld-tagCategory.read and TagCategory.jsonld-tagCategory.write).
type TagCategory struct {
	Context     *Context   `json:"@context,omitempty"` // JSON-LD only
	ID          ID         `json:"@id,omitempty"`     // JSON-LD @id, maps to "id" in JSON
	Type        string     `json:"@type,omitempty"`   // JSON-LD only
	Label       string     `json:"label"`
	Description *string    `json:"description,omitempty"`
	Color       string     `json:"color"`
	Owner       *IRI       `json:"owner,omitempty"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}

// Template represents a template in Koillection, combining read and write fields (aligned with Template.jsonld-template.read and Template.jsonld-template.write).
type Template struct {
	Context   *Context   `json:"@context,omitempty"` // JSON-LD only
	ID        ID         `json:"@id,omitempty"`     // JSON-LD @id, maps to "id" in JSON
	Type      string     `json:"@type,omitempty"`   // JSON-LD only
	Name      string     `json:"name"`
	Owner     *IRI       `json:"owner,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// User represents a user in Koillection, combining read and write fields (aligned with User.jsonld-user.read, minimal write assumed).
type User struct {
	Context                     *Context   `json:"@context,omitempty"` // JSON-LD only
	ID                          ID         `json:"@id,omitempty"`     // JSON-LD @id, maps to "id" in JSON
	Type                        string     `json:"@type,omitempty"`   // JSON-LD only
	Username                    string     `json:"username"`
	Email                       string     `json:"email"`
	PlainPassword               *string    `json:"plainPassword,omitempty"`
	Avatar                      *string    `json:"avatar,omitempty"`
	Currency                    string     `json:"currency"`
	Locale                      string     `json:"locale"`
	Timezone                    string     `json:"timezone"`
	DateFormat                  string     `json:"dateFormat"`
	DiskSpaceAllowed            int        `json:"diskSpaceAllowed"`
	Visibility                  string     `json:"visibility"`
	LastDateOfActivity          *time.Time `json:"lastDateOfActivity,omitempty"`
	WishlistsFeatureEnabled     bool       `json:"wishlistsFeatureEnabled"`
	TagsFeatureEnabled          bool       `json:"tagsFeatureEnabled"`
	SignsFeatureEnabled         bool       `json:"signsFeatureEnabled"`
	AlbumsFeatureEnabled        bool       `json:"albumsFeatureEnabled"`
	LoansFeatureEnabled         bool       `json:"loansFeatureEnabled"`
	TemplatesFeatureEnabled     bool       `json:"templatesFeatureEnabled"`
	HistoryFeatureEnabled       bool       `json:"historyFeatureEnabled"`
	StatisticsFeatureEnabled    bool       `json:"statisticsFeatureEnabled"`
	ScrapingFeatureEnabled      bool       `json:"scrapingFeatureEnabled"`
	SearchInDataByDefaultEnabled bool       `json:"searchInDataByDefaultEnabled"`
	DisplayItemsNameInGridView  bool       `json:"displayItemsNameInGridView"`
	SearchResultsDisplayMode    string     `json:"searchResultsDisplayMode"`
	CreatedAt                   time.Time  `json:"createdAt"`
	UpdatedAt                   *time.Time `json:"updatedAt,omitempty"`
}

// Wish represents a wish in Koillection, combining read and write fields (aligned with Wish.jsonld-wish.read and Wish.jsonld-wish.write).
type Wish struct {
	Context             *Context   `json:"@context,omitempty"` // JSON-LD only
	ID                  ID         `json:"@id,omitempty"`     // JSON-LD @id, maps to "id" in JSON
	Type                string     `json:"@type,omitempty"`   // JSON-LD only
	Name                string     `json:"name"`
	URL                 *string    `json:"url,omitempty"`
	Price               *string    `json:"price,omitempty"`
	Currency            *string    `json:"currency,omitempty"`
	Wishlist            *IRI       `json:"wishlist"`
	Owner               *IRI       `json:"owner,omitempty"`
	Comment             *string    `json:"comment,omitempty"`
	Image               *string    `json:"image,omitempty"`
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty"`
	Visibility          string     `json:"visibility,omitempty"`
	ParentVisibility    *string    `json:"parentVisibility,omitempty"`
	FinalVisibility     string     `json:"finalVisibility,omitempty"`
	ScrapedFromURL      *string    `json:"scrapedFromUrl,omitempty"`
	CreatedAt           time.Time  `json:"createdAt"`
	UpdatedAt           *time.Time `json:"updatedAt,omitempty"`
	File                *string    `json:"file,omitempty"` // Write-only, binary data via multipart form
}

// Wishlist represents a wishlist in Koillection, combining read and write fields (aligned with Wishlist.jsonld-wishlist.read and Wishlist.jsonld-wishlist.write).
type Wishlist struct {
	Context         *Context   `json:"@context,omitempty"` // JSON-LD only
	ID              ID         `json:"@id,omitempty"`     // JSON-LD @id, maps to "id" in JSON
	Type            string     `json:"@type,omitempty"`   // JSON-LD only
	Name            string     `json:"name"`
	Owner           *IRI       `json:"owner,omitempty"`
	Color           string     `json:"color"`
	Parent          *IRI       `json:"parent,omitempty"`
	Image           *string    `json:"image,omitempty"`
	SeenCounter     int        `json:"seenCounter,omitempty"`
	Visibility      string     `json:"visibility,omitempty"`
	ParentVisibility *string   `json:"parentVisibility,omitempty"`
	FinalVisibility string     `json:"finalVisibility,omitempty"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       *time.Time `json:"updatedAt,omitempty"`
	File            *string    `json:"file,omitempty"`        // Write-only, binary data via multipart form
	DeleteImage     *bool      `json:"deleteImage,omitempty"` // Write-only
}˜