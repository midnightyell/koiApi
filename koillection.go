// package koillection
package main

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"
)

// IRI represents an IRI reference in the Koillection API.
type IRI string
type ID string
type Metrics map[string]string

// Album represents an album in Koillection.
type Album struct {
	ID               ID         `json:"id,omitempty,omitempty"`
	Title            string     `json:"title,omitempty"`
	Color            string     `json:"color,omitempty"`
	Image            *string    `json:"image,omitempty"`
	Owner            *IRI       `json:"owner,omitempty"`
	Parent           *IRI       `json:"parent,omitempty"`
	SeenCounter      int        `json:"seenCounter,omitempty"`
	Visibility       string     `json:"visibility,omitempty"`
	ParentVisibility *string    `json:"parentVisibility,omitempty"`
	FinalVisibility  string     `json:"finalVisibility,omitempty"`
	CreatedAt        time.Time  `json:"createdAt,omitempty"`
	UpdatedAt        *time.Time `json:"updatedAt,omitempty"`
}

// ChoiceList represents a predefined list of options in Koillection.
type ChoiceList struct {
	ID        ID         `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Choices   []string   `json:"choices,omitempty"`
	Owner     *IRI       `json:"owner,omitempty"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// Collection represents a collection in Koillection.
type Collection struct {
	ID                   ID         `json:"id,omitempty"`
	Title                string     `json:"title,omitempty"`
	Parent               *IRI       `json:"parent,omitempty"`
	Owner                *IRI       `json:"owner,omitempty"`
	Color                string     `json:"color,omitempty"`
	Image                *string    `json:"image,omitempty"`
	SeenCounter          int        `json:"seenCounter,omitempty"`
	ItemsDefaultTemplate *IRI       `json:"itemsDefaultTemplate,omitempty"`
	Visibility           string     `json:"visibility,omitempty"`
	ParentVisibility     *string    `json:"parentVisibility,omitempty"`
	FinalVisibility      string     `json:"finalVisibility,omitempty"`
	ScrapedFromURL       *string    `json:"scrapedFromUrl,omitempty"`
	CreatedAt            time.Time  `json:"createdAt,omitempty"`
	UpdatedAt            *time.Time `json:"updatedAt,omitempty"`
}

// Datum represents a custom data field in Koillection.
type Datum struct {
	ID                  ID         `json:"id,omitempty"`
	Item                *IRI       `json:"item,omitempty"`
	Collection          *IRI       `json:"collection,omitempty"`
	Type                string     `json:"type,omitempty"`
	Label               string     `json:"label,omitempty"`
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
	CreatedAt           time.Time  `json:"createdAt,omitempty"`
	UpdatedAt           *time.Time `json:"updatedAt,omitempty"`
}

// Field represents a template field in Koillection.
type Field struct {
	ID         ID     `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Position   int    `json:"position,omitempty"`
	Type       string `json:"type,omitempty"`
	ChoiceList *IRI   `json:"choiceList,omitempty"`
	Template   *IRI   `json:"template,omitempty"`
	Visibility string `json:"visibility,omitempty"`
	Owner      *IRI   `json:"owner,omitempty"`
}

// Inventory represents an inventory record in Koillection.
type Inventory struct {
	ID        ID         `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Content   []string   `json:"content,omitempty"`
	Owner     *IRI       `json:"owner,omitempty"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// Item represents an item within a collection.
type Item struct {
	ID                  ID         `json:"id,omitempty"`
	Name                string     `json:"name,omitempty"`
	Quantity            int        `json:"quantity,omitempty"`
	Collection          *IRI       `json:"collection,omitempty"`
	Owner               *IRI       `json:"owner,omitempty"`
	Image               *string    `json:"image,omitempty"`
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty"`
	ImageLargeThumbnail *string    `json:"imageLargeThumbnail,omitempty"`
	SeenCounter         int        `json:"seenCounter,omitempty"`
	Visibility          string     `json:"visibility,omitempty"`
	ParentVisibility    *string    `json:"parentVisibility,omitempty"`
	FinalVisibility     string     `json:"finalVisibility,omitempty"`
	ScrapedFromURL      *string    `json:"scrapedFromUrl,omitempty"`
	CreatedAt           time.Time  `json:"createdAt,omitempty"`
	UpdatedAt           *time.Time `json:"updatedAt,omitempty"`
}

// Loan represents a loan record in Koillection.
type Loan struct {
	ID         ID         `json:"id,omitempty"`
	Item       *IRI       `json:"item,omitempty"`
	LentTo     string     `json:"lentTo,omitempty"`
	LentAt     time.Time  `json:"lentAt,omitempty"`
	ReturnedAt *time.Time `json:"returnedAt,omitempty"`
	Owner      *IRI       `json:"owner,omitempty"`
}

// Log represents an action or event in Koillection.
type Log struct {
	ID            ID        `json:"id,omitempty"`
	Type          *string   `json:"type,omitempty"`
	LoggedAt      time.Time `json:"loggedAt,omitempty"`
	ObjectID      string    `json:"objectId,omitempty"`
	ObjectLabel   string    `json:"objectLabel,omitempty"`
	ObjectClass   string    `json:"objectClass,omitempty"`
	ObjectDeleted bool      `json:"objectDeleted,omitempty"`
	Owner         *IRI      `json:"owner,omitempty"`
}

// Photo represents a photo in Koillection.
type Photo struct {
	ID                  ID         `json:"id,omitempty"`
	Title               string     `json:"title,omitempty"`
	Comment             *string    `json:"comment,omitempty"`
	Place               *string    `json:"place,omitempty"`
	Album               *IRI       `json:"album,omitempty"`
	Owner               *IRI       `json:"owner,omitempty"`
	Image               *string    `json:"image,omitempty"`
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty"`
	TakenAt             *time.Time `json:"takenAt,omitempty"`
	Visibility          string     `json:"visibility,omitempty"`
	ParentVisibility    *string    `json:"parentVisibility,omitempty"`
	FinalVisibility     string     `json:"finalVisibility,omitempty"`
	CreatedAt           time.Time  `json:"createdAt,omitempty"`
	UpdatedAt           *time.Time `json:"updatedAt,omitempty"`
}

// Tag represents a tag in Koillection.
type Tag struct {
	ID                  ID         `json:"id,omitempty"`
	Label               string     `json:"label,omitempty"`
	Description         *string    `json:"description,omitempty"`
	Image               *string    `json:"image,omitempty"`
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty"`
	Owner               *IRI       `json:"owner,omitempty"`
	Category            *IRI       `json:"category,omitempty"`
	SeenCounter         int        `json:"seenCounter,omitempty"`
	Visibility          string     `json:"visibility,omitempty"`
	CreatedAt           time.Time  `json:"createdAt,omitempty"`
	UpdatedAt           *time.Time `json:"updatedAt,omitempty"`
}

// TagCategory represents a tag category in Koillection.
type TagCategory struct {
	ID          ID         `json:"id,omitempty"`
	Label       string     `json:"label,omitempty"`
	Description *string    `json:"description,omitempty"`
	Color       string     `json:"color,omitempty"`
	Owner       *IRI       `json:"owner,omitempty"`
	CreatedAt   time.Time  `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}

// Template represents a template in Koillection.
type Template struct {
	ID        ID         `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Owner     *IRI       `json:"owner,omitempty"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// User represents a user in Koillection.
type User struct {
	ID                           ID         `json:"id,omitempty"`
	Username                     string     `json:"username,omitempty"`
	Email                        string     `json:"email,omitempty"`
	Avatar                       *string    `json:"avatar,omitempty"`
	Currency                     string     `json:"currency,omitempty"`
	Locale                       string     `json:"locale,omitempty"`
	Timezone                     string     `json:"timezone,omitempty"`
	DateFormat                   string     `json:"dateFormat,omitempty"`
	DiskSpaceAllowed             int        `json:"diskSpaceAllowed,omitempty"`
	Visibility                   string     `json:"visibility,omitempty"`
	LastDateOfActivity           *time.Time `json:"lastDateOfActivity,omitempty"`
	WishlistsFeatureEnabled      bool       `json:"wishlistsFeatureEnabled,omitempty"`
	TagsFeatureEnabled           bool       `json:"tagsFeatureEnabled,omitempty"`
	SignsFeatureEnabled          bool       `json:"signsFeatureEnabled,omitempty"`
	AlbumsFeatureEnabled         bool       `json:"albumsFeatureEnabled,omitempty"`
	LoansFeatureEnabled          bool       `json:"loansFeatureEnabled,omitempty"`
	TemplatesFeatureEnabled      bool       `json:"templatesFeatureEnabled,omitempty"`
	HistoryFeatureEnabled        bool       `json:"historyFeatureEnabled,omitempty"`
	StatisticsFeatureEnabled     bool       `json:"statisticsFeatureEnabled,omitempty"`
	ScrapingFeatureEnabled       bool       `json:"scrapingFeatureEnabled,omitempty"`
	SearchInDataByDefaultEnabled bool       `json:"searchInDataByDefaultEnabled,omitempty"`
	DisplayItemsNameInGridView   bool       `json:"displayItemsNameInGridView,omitempty"`
	SearchResultsDisplayMode     string     `json:"searchResultsDisplayMode,omitempty"`
	CreatedAt                    time.Time  `json:"createdAt,omitempty"`
	UpdatedAt                    *time.Time `json:"updatedAt,omitempty"`
}

// Wish represents a wish in Koillection.
type Wish struct {
	ID                  ID         `json:"id,omitempty"`
	Name                string     `json:"name,omitempty"`
	URL                 *string    `json:"url,omitempty"`
	Price               *string    `json:"price,omitempty"`
	Currency            *string    `json:"currency,omitempty"`
	Wishlist            *IRI       `json:"wishlist,omitempty"`
	Owner               *IRI       `json:"owner,omitempty"`
	Comment             *string    `json:"comment,omitempty"`
	Image               *string    `json:"image,omitempty"`
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty"`
	Visibility          string     `json:"visibility,omitempty"`
	ParentVisibility    *string    `json:"parentVisibility,omitempty"`
	FinalVisibility     string     `json:"finalVisibility,omitempty"`
	ScrapedFromURL      *string    `json:"scrapedFromUrl,omitempty"`
	CreatedAt           time.Time  `json:"createdAt,omitempty"`
	UpdatedAt           *time.Time `json:"updatedAt,omitempty"`
}

// Wishlist represents a wishlist in Koillection.
type Wishlist struct {
	ID               ID         `json:"id,omitempty"`
	Name             string     `json:"name,omitempty"`
	Owner            *IRI       `json:"owner,omitempty"`
	Color            string     `json:"color,omitempty"`
	Parent           *IRI       `json:"parent,omitempty"`
	Image            *string    `json:"image,omitempty"`
	SeenCounter      int        `json:"seenCounter,omitempty"`
	Visibility       string     `json:"visibility,omitempty"`
	ParentVisibility *string    `json:"parentVisibility,omitempty"`
	FinalVisibility  string     `json:"finalVisibility,omitempty"`
	CreatedAt        time.Time  `json:"createdAt,omitempty"`
	UpdatedAt        *time.Time `json:"updatedAt,omitempty"`
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
