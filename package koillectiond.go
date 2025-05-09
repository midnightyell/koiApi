package koillection

import (
	"time"
)

// IRI represents an IRI reference used in the API.
type IRI string

// ID represents a unique identifier for resources.
type ID string

// Metrics represents a map of metrics data.
type Metrics map[string]string

// Context represents the JSON-LD @context field.
type Context struct {
	Vocab string `json:"@vocab,omitempty"`
	Hydra string `json:"hydra,omitempty"`
}

// Album represents an album in Koillection (aligned with Album.jsonld-album.read).
type Album struct {
	Context         *Context   `json:"@context,omitempty"`
	ID              ID         `json:"@id,omitempty"`
	Type            string     `json:"@type,omitempty"`
	AlbumID         ID         `json:"id,omitempty"`
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
}

// ChoiceList represents a predefined list of options in Koillection (aligned with ChoiceList.jsonld-choiceList.read).
type ChoiceList struct {
	Context   *Context   `json:"@context,omitempty"`
	ID        ID         `json:"@id,omitempty"`
	Type      string     `json:"@type,omitempty"`
	ChoiceListID ID    `json:"id,omitempty"`
	Name      string     `json:"name"`
	Choices   []string   `json:"choices"`
	Owner     *IRI       `json:"owner,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// Collection represents a collection in Koillection (aligned with Collection.jsonld-collection.read).
type Collection struct {
	Context             *Context   `json:"@context,omitempty"`
	ID                  ID         `json:"@id,omitempty"`
	Type                string     `json:"@type,omitempty"`
	CollectionID        ID         `json:"id,omitempty"`
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
}

// Datum represents a custom data field in Koillection (aligned with Datum.jsonld-datum.read).
type Datum struct {
	Context             *Context   `json:"@context,omitempty"`
	ID                  ID         `json:"@id,omitempty"`
	Type                string     `json:"@type,omitempty"`
	DatumID             ID         `json:"id,omitempty"`
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
}

// Field represents a template field in Koillection (aligned with Field.jsonld-field.read).
type Field struct {
	Context    *Context   `json:"@context,omitempty"`
	ID         ID         `json:"@id,omitempty"`
	Type       string     `json:"@type,omitempty"`
	FieldID    ID         `json:"id,omitempty"`
	Name       string     `json:"name"`
	Position   int        `json:"position"`
	FieldType  string     `json:"type"`
	ChoiceList *IRI       `json:"choiceList,omitempty"`
	Template   *IRI       `json:"template"`
	Visibility string     `json:"visibility,omitempty"`
	Owner      *IRI       `json:"owner,omitempty"`
}

// Inventory represents an inventory record in Koillection (aligned with Inventory.jsonld-inventory.read).
type Inventory struct {
	Context      *Context   `json:"@context,omitempty"`
	ID           ID         `json:"@id,omitempty"`
	Type         string     `json:"@type,omitempty"`
	InventoryID  ID         `json:"id,omitempty"`
	Name         string     `json:"name"`
	Content      []string   `json:"content"`
	Owner        *IRI       `json:"owner,omitempty"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt,omitempty"`
}

// Item represents an item within a collection (aligned with Item.jsonld-item.read).
type Item struct {
	Context             *Context   `json:"@context,omitempty"`
	ID                  ID         `json:"@id,omitempty"`
	Type                string     `json:"@type,omitempty"`
	ItemID              ID         `json:"id,omitempty"`
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
}

// Loan represents a loan record in Koillection (aligned with Loan.jsonld-loan.read).
type Loan struct {
	Context    *Context   `json:"@context,omitempty"`
	ID         ID         `json:"@id,omitempty"`
	Type       string     `json:"@type,omitempty"`
	LoanID     ID         `json:"id,omitempty"`
	Item       *IRI       `json:"item"`
	LentTo     string     `json:"lentTo"`
	LentAt     time.Time  `json:"lentAt"`
	ReturnedAt *time.Time `json:"returnedAt,omitempty"`
	Owner      *IRI       `json:"owner,omitempty"`
}

// Log represents an action or event in Koillection (aligned with Log.jsonld-log.read).
type Log struct {
	Context       *Context   `json:"@context,omitempty"`
	ID            ID         `json:"@id,omitempty"`
	Type          string     `json:"@type,omitempty"`
	LogID         ID         `json:"id,omitempty"`
	LogType       *string    `json:"type,omitempty"`
	LoggedAt      time.Time  `json:"loggedAt"`
	ObjectID      string     `json:"objectId"`
	ObjectLabel   string     `json:"objectLabel"`
	ObjectClass   string     `json:"objectClass"`
	ObjectDeleted bool       `json:"objectDeleted"`
	Owner         *IRI       `json:"owner,omitempty"`
}

// Photo represents a photo in Koillection (aligned with Photo.jsonld-photo.read).
type Photo struct {
	Context             *Context   `json:"@context,omitempty"`
	ID                  ID         `json:"@id,omitempty"`
	Type                string     `json:"@type,omitempty"`
	PhotoID             ID         `json:"id,omitempty"`
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
}

// Tag represents a tag in Koillection (aligned with Tag.jsonld-tag.read).
type Tag struct {
	Context             *Context   `json:"@context,omitempty"`
	ID                  ID         `json:"@id,omitempty"`
	Type                string     `json:"@type,omitempty"`
	TagID               ID         `json:"id,omitempty"`
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
}

// TagCategory represents a tag category in Koillection (aligned with TagCategory.jsonld-tagCategory.read).
type TagCategory struct {
	Context     *Context   `json:"@context,omitempty"`
	ID          ID         `json:"@id,omitempty"`
	Type        string     `json:"@type,omitempty"`
	TagCategoryID ID      `json:"id,omitempty"`
	Label       string     `json:"label"`
	Description *string    `json:"description,omitempty"`
	Color       string     `json:"color"`
	Owner       *IRI       `json:"owner,omitempty"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}

// Template represents a template in Koillection (aligned with Template.jsonld-template.read).
type Template struct {
	Context     *Context   `json:"@context,omitempty"`
	ID          ID         `json:"@id,omitempty"`
	Type        string     `json:"@type,omitempty"`
	TemplateID  ID         `json:"id,omitempty"`
	Name        string     `json:"name"`
	Owner       *IRI       `json:"owner,omitempty"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}

// User represents a user in Koillection (aligned with User.jsonld-user.read).
type User struct {
	Context                     *Context   `json:"@context,omitempty"`
	ID                          ID         `json:"@id,omitempty"`
	Type                        string     `json:"@type,omitempty"`
	UserID                      ID         `json:"id,omitempty"`
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

// Wish represents a wish in Koillection (aligned with Wish.jsonld-wish.read).
type Wish struct {
	Context             *Context   `json:"@context,omitempty"`
	ID                  ID         `json:"@id,omitempty"`
	Type                string     `json:"@type,omitempty"`
	WishID              ID         `json:"id,omitempty"`
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
}

// Wishlist represents a wishlist in Koillection (aligned with Wishlist.jsonld-wishlist.read).
type Wishlist struct {
	Context         *Context   `json:"@context,omitempty"`
	ID              ID         `json:"@id,omitempty"`
	Type            string     `json:"@type,omitempty"`
	WishlistID      ID         `json:"id,omitempty"`
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
}