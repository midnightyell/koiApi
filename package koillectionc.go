package koillection

import (
	"time"
)

// IRI represents an IRI reference used in the API.
type IRI string

// ID represents a unique identifier for resources (JSON-LD @id).
type ID string

// Metrics represents a map of metrics data.
type Metrics map[string]string

// Context represents the JSON-LD @context field.
type Context struct {
	Vocab string `json:"@vocab,omitempty"`
	Hydra string `json:"hydra,omitempty"`
}

// AlbumRead represents an album in Koillection for read operations (aligned with Album.jsonld-album.read).
type AlbumRead struct {
	Context         *Context   `json:"@context,omitempty"`
	ID              ID         `json:"@id,omitempty"`
	Type            string     `json:"@type,omitempty"`
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

// AlbumWrite represents an album in Koillection for write operations (aligned with Album.jsonld-album.write).
type AlbumWrite struct {
	Title       string  `json:"title"`
	File        *string `json:"file,omitempty"`        // Binary data, typically handled via multipart form
	DeleteImage *bool   `json:"deleteImage,omitempty"`
	Parent      *IRI    `json:"parent,omitempty"`
	Visibility  string  `json:"visibility,omitempty"`
}

// ChoiceListRead represents a predefined list of options in Koillection for read operations (aligned with ChoiceList.jsonld-choiceList.read).
type ChoiceListRead struct {
	Context   *Context   `json:"@context,omitempty"`
	ID        ID         `json:"@id,omitempty"`
	Type      string     `json:"@type,omitempty"`
	Name      string     `json:"name"`
	Choices   []string   `json:"choices"`
	Owner     *IRI       `json:"owner,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// ChoiceListWrite represents a predefined list of options in Koillection for write operations (aligned with ChoiceList.jsonld-choiceList.write).
type ChoiceListWrite struct {
	Name    string   `json:"name"`
	Choices []string `json:"choices"`
}

// CollectionRead represents a collection in Koillection for read operations (aligned with Collection.jsonld-collection.read).
type CollectionRead struct {
	Context             *Context   `json:"@context,omitempty"`
	ID                  ID         `json:"@id,omitempty"`
	Type                string     `json:"@type,omitempty"`
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

// CollectionWrite represents a collection in Koillection for write operations (aligned with Collection.jsonld-collection.write).
type CollectionWrite struct {
	Title               string  `json:"title"`
	Parent              *IRI    `json:"parent,omitempty"`
	File                *string `json:"file,omitempty"`        // Binary data, typically handled via multipart form
	DeleteImage         *bool   `json:"deleteImage,omitempty"`
	ItemsDefaultTemplate *IRI   `json:"itemsDefaultTemplate,omitempty"`
	Visibility          string  `json:"visibility,omitempty"`
}

// DatumRead represents a custom data field in Koillection for read operations (aligned with Datum.jsonld-datum.read).
type DatumRead struct {
	Context             *Context   `json:"@context,omitempty"`
	ID                  ID         `json:"@id,omitempty"`
	Type                string     `json:"@type,omitempty"`
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

// DatumWrite represents a custom data field in Koillection for write operations (aligned with Datum.jsonld-datum.write).
type DatumWrite struct {
	Item        *IRI    `json:"item,omitempty"`
	Collection  *IRI    `json:"collection,omitempty"`
	Type        string  `json:"type"`
	Label       string  `json:"label"`
	Value       *string `json:"value,omitempty"`
	Position    *int    `json:"position,omitempty"`
	Currency    *string `json:"currency,omitempty"`
	FileImage   *string `json:"fileImage,omitempty"` // Binary data, typically handled via multipart form
	FileFile    *string `json:"fileFile,omitempty"`  // Binary data, typically handled via multipart form
	FileVideo   *string `json:"fileVideo,omitempty"` // Binary data, typically handled via multipart form
	ChoiceList  *IRI    `json:"choiceList,omitempty"`
	Visibility  string  `json:"visibility,omitempty"`
}

// FieldRead represents a template field in Koillection for read operations (aligned with Field.jsonld-field.read).
type FieldRead struct {
	Context    *Context   `json:"@context,omitempty"`
	ID         ID         `json:"@id,omitempty"`
	Type       string     `json:"@type,omitempty"`
	Name       string     `json:"name"`
	Position   int        `json:"position"`
	FieldType  string     `json:"type"`
	ChoiceList *IRI       `json:"choiceList,omitempty"`
	Template   *IRI       `json:"template"`
	Visibility string     `json:"visibility,omitempty"`
	Owner      *IRI       `json:"owner,omitempty"`
}

// FieldWrite represents a template field in Koillection for write operations (aligned with Field.jsonld-field.write).
type FieldWrite struct {
	Name       string `json:"name"`
	Position   int    `json:"position"`
	Type       string `json:"type"`
	ChoiceList *IRI   `json:"choiceList,omitempty"`
	Template   *IRI   `json:"template"`
	Visibility string `json:"visibility,omitempty"`
}

// InventoryRead represents an inventory record in Koillection for read operations (aligned with Inventory.jsonld-inventory.read).
type InventoryRead struct {
	Context   *Context   `json:"@context,omitempty"`
	ID        ID         `json:"@id,omitempty"`
	Type      string     `json:"@type,omitempty"`
	Name      string     `json:"name"`
	Content   []string   `json:"content"`
	Owner     *IRI       `json:"owner,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// InventoryWrite represents an inventory record in Koillection for write operations (no write schema provided, assuming similar to read).
type InventoryWrite struct {
	Name    string   `json:"name"`
	Content []string `json:"content"`
}

// ItemRead represents an item within a collection for read operations (aligned with Item.jsonld-item.read).
type ItemRead struct {
	Context             *Context   `json:"@context,omitempty"`
	ID                  ID         `json:"@id,omitempty"`
	Type                string     `json:"@type,omitempty"`
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

// ItemWrite represents an item within a collection for write operations (aligned with Item.jsonld-item.write).
type ItemWrite struct {
	Name        string   `json:"name"`
	Quantity    int      `json:"quantity"`
	Collection  *IRI     `json:"collection"`
	Tags        []IRI    `json:"tags,omitempty"`
	RelatedItems []IRI   `json:"relatedItems,omitempty"`
	File        *string  `json:"file,omitempty"` // Binary data, typically handled via multipart form
	Visibility  string   `json:"visibility,omitempty"`
}

// LoanRead represents a loan record in Koillection for read operations (aligned with Loan.jsonld-loan.read).
type LoanRead struct {
	Context    *Context   `json:"@context,omitempty"`
	ID         ID         `json:"@id,omitempty"`
	Type       string     `json:"@type,omitempty"`
	Item       *IRI       `json:"item"`
	LentTo     string     `json:"lentTo"`
	LentAt     time.Time  `json:"lentAt"`
	ReturnedAt *time.Time `json:"returnedAt,omitempty"`
	Owner      *IRI       `json:"owner,omitempty"`
}

// LoanWrite represents a loan record in Koillection for write operations (aligned with Loan.jsonld-loan.write).
type LoanWrite struct {
	Item       *IRI       `json:"item"`
	LentTo     string     `json:"lentTo"`
	LentAt     time.Time  `json:"lentAt"`
	ReturnedAt *time.Time `json:"returnedAt,omitempty"`
}

// LogRead represents an action or event in Koillection for read operations (aligned with Log.jsonld-log.read).
type LogRead struct {
	Context       *Context   `json:"@context,omitempty"`
	ID            ID         `json:"@id,omitempty"`
	Type          string     `json:"@type,omitempty"`
	LogType       *string    `json:"type,omitempty"`
	LoggedAt      time.Time  `json:"loggedAt"`
	ObjectID      string     `json:"objectId"`
	ObjectLabel   string     `json:"objectLabel"`
	ObjectClass   string     `json:"objectClass"`
	ObjectDeleted bool       `json:"objectDeleted"`
	Owner         *IRI       `json:"owner,omitempty"`
}

// LogWrite represents an action or event in Koillection for write operations (no write schema provided, assuming minimal).
type LogWrite struct {
	Type        *string   `json:"type,omitempty"`
	LoggedAt    time.Time `json:"loggedAt"`
	ObjectID    string    `json:"objectId"`
	ObjectLabel string    `json:"objectLabel"`
	ObjectClass string    `json:"objectClass"`
}

// PhotoRead represents a photo in Koillection for read operations (aligned with Photo.jsonld-photo.read).
type PhotoRead struct {
	Context             *Context   `json:"@context,omitempty"`
	ID                  ID         `json:"@id,omitempty"`
	Type                string     `json:"@type,omitempty"`
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

// PhotoWrite represents a photo in Koillection for write operations (aligned with Photo.jsonld-photo.write).
type PhotoWrite struct {
	Title      string  `json:"title"`
	Comment    *string `json:"comment,omitempty"`
	Place      *string `json:"place,omitempty"`
	Album      *IRI    `json:"album"`
	File       *string `json:"file,omitempty"` // Binary data, typically handled via multipart form
	Visibility string  `json:"visibility,omitempty"`
}

// TagRead represents a tag in Koillection for read operations (aligned with Tag.jsonld-tag.read).
type TagRead struct {
	Context             *Context   `json:"@context,omitempty"`
	ID                  ID         `json:"@id,omitempty"`
	Type                string     `json:"@type,omitempty"`
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

// TagWrite represents a tag in Koillection for write operations (aligned with Tag.jsonld-tag.write).
type TagWrite struct {
	Label       string  `json:"label"`
	Description *string `json:"description,omitempty"`
	File        *string `json:"file,omitempty"` // Binary data, typically handled via multipart form
	Category    *IRI    `json:"category,omitempty"`
	Visibility  string  `json:"visibility,omitempty"`
}

// TagCategoryRead represents a tag category in Koillection for read operations (aligned with TagCategory.jsonld-tagCategory.read).
type TagCategoryRead struct {
	Context     *Context   `json:"@context,omitempty"`
	ID          ID         `json:"@id,omitempty"`
	Type        string     `json:"@type,omitempty"`
	Label       string     `json:"label"`
	Description *string    `json:"description,omitempty"`
	Color       string     `json:"color"`
	Owner       *IRI       `json:"owner,omitempty"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}

// TagCategoryWrite represents a tag category in Koillection for write operations (aligned with TagCategory.jsonld-tagCategory.write).
type TagCategoryWrite struct {
	Label       string  `json:"label"`
	Description *string `json:"description,omitempty"`
	Color       string  `json:"color"`
}

// TemplateRead represents a template in Koillection for read operations (aligned with Template.jsonld-template.read).
type TemplateRead struct {
	Context   *Context   `json:"@context,omitempty"`
	ID        ID         `json:"@id,omitempty"`
	Type      string     `json:"@type,omitempty"`
	Name      string     `json:"name"`
	Owner     *IRI       `json:"owner,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// TemplateWrite represents a template in Koillection for write operations (aligned with Template.jsonld-template.write).
type TemplateWrite struct {
	Name string `json:"name"`
}

// UserRead represents a user in Koillection for read operations (aligned with User.jsonld-user.read).
type UserRead struct {
	Context                     *Context   `json:"@context,omitempty"`
	ID                          ID         `json:"@id,omitempty"`
	Type                        string     `json:"@type,omitempty"`
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

// UserWrite represents a user in Koillection for write operations (no write schema provided, assuming minimal).
type UserWrite struct {
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
}

// WishRead represents a wish in Koillection for read operations (aligned with Wish.jsonld-wish.read).
type WishRead struct {
	Context             *Context   `json:"@context,omitempty"`
	ID                  ID         `json:"@id,omitempty"`
	Type                string     `json:"@type,omitempty"`
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

// WishWrite represents a wish in Koillection for write operations (aligned with Wish.jsonld-wish.write).
type WishWrite struct {
	Name       string  `json:"name"`
	URL        *string `json:"url,omitempty"`
	Price      *string `json:"price,omitempty"`
	Currency   *string `json:"currency,omitempty"`
	Wishlist   *IRI    `json:"wishlist"`
	Comment    *string `json:"comment,omitempty"`
	File       *string `json:"file,omitempty"` // Binary data, typically handled via multipart form
	Visibility string  `json:"visibility,omitempty"`
}

// WishlistRead represents a wishlist in Koillection for read operations (aligned with Wishlist.jsonld-wishlist.read).
type WishlistRead struct {
	Context         *Context   `json:"@context,omitempty"`
	ID              ID         `json:"@id,omitempty"`
	Type            string     `json:"@type,omitempty"`
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

// WishlistWrite represents a wishlist in Koillection for write operations (aligned with Wishlist.jsonld-wishlist.write).
type WishlistWrite struct {
	Name        string  `json:"name"`
	Parent      *IRI    `json:"parent,omitempty"`
	File        *string `json:"file,omitempty"`        // Binary data, typically handled via multipart form
	DeleteImage *bool   `json:"deleteImage,omitempty"`
	Visibility  string  `json:"visibility,omitempty"`
}