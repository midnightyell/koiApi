package koillection

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"
)

// ID represents a non-empty, non-nil identifier in the Koillection API.
type ID string

// NewID creates a new ID, ensuring it is non-empty.
func NewID(id string) (ID, error) {
	if id == "" {
		return "", errors.New("ID cannot be empty")
	}
	return ID(id), nil
}

// String returns the string representation of the ID.
func (id ID) String() string {
	return string(id)
}

// IRI represents an IRI reference in the Koillection API.
type IRI string

// Album represents an album in Koillection.
type Album struct {
	ID               ID         `json:"id"`
	Title            string     `json:"title"`
	Color            string     `json:"color"`
	Image            *string    `json:"image"`
	Owner            *IRI       `json:"owner"`
	Parent           *IRI       `json:"parent"`
	SeenCounter      int        `json:"seenCounter"`
	Visibility       string     `json:"visibility"`
	ParentVisibility *string    `json:"parentVisibility"`
	FinalVisibility  string     `json:"finalVisibility"`
	CreatedAt        time.Time  `json:"createdAt"`
	UpdatedAt        *time.Time `json:"updatedAt"`
}
func (obj *Album) id() (ID) {	return string(obj.ID) }
func (obj *Album) iri() (IRI) { path, _ := GetIRI(obj); return path }



// ChoiceList represents a predefined list of options in Koillection.
type ChoiceList struct {
	ID        ID         `json:"id"`
	Name      string     `json:"name"`
	Choices   []string   `json:"choices"`
	Owner     *IRI       `json:"owner"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
func (obj *ChoiceList) id() (ID) {	return string(obj.ID) }
func (obj *ChoiceList) iri() (IRI) { path, _ := GetIRI(obj); return path }



// Collection represents a collection in Koillection.
type Collection struct {
	ID                   ID         `json:"id"`
	Title                string     `json:"title"`
	Parent               *IRI       `json:"parent"`
	Owner                *IRI       `json:"owner"`
	Color                string     `json:"color"`
	Image                *string    `json:"image"`
	SeenCounter          int        `json:"seenCounter"`
	ItemsDefaultTemplate *IRI       `json:"itemsDefaultTemplate"`
	Visibility           string     `json:"visibility"`
	ParentVisibility     *string    `json:"parentVisibility"`
	FinalVisibility      string     `json:"finalVisibility"`
	ScrapedFromURL       *string    `json:"scrapedFromUrl"`
	CreatedAt            time.Time  `json:"createdAt"`
	UpdatedAt            *time.Time `json:"updatedAt"`
}
func (obj Collection) id() (ID) {	return string(obj.ID) }
func (obj Collection) iri() (IRI) { path, _ := GetIRI(obj); return path }


// Datum represents a custom data field in Koillection.
type Datum struct {
	ID                  ID         `json:"id"`
	Item                *IRI       `json:"item"`
	Collection          *IRI       `json:"collection"`
	Type                string     `json:"type"`
	Label               string     `json:"label"`
	Value               *string    `json:"value"`
	Position            *int       `json:"position"`
	Currency            *string    `json:"currency"`
	Image               *string    `json:"image"`
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail"`
	ImageLargeThumbnail *string    `json:"imageLargeThumbnail"`
	File                *string    `json:"file"`
	Video               *string    `json:"video"`
	OriginalFilename    *string    `json:"originalFilename"`
	ChoiceList          *IRI       `json:"choiceList"`
	Owner               *IRI       `json:"owner"`
	Visibility          string     `json:"visibility"`
	ParentVisibility    *string    `json:"parentVisibility"`
	FinalVisibility     string     `json:"finalVisibility"`
	CreatedAt           time.Time  `json:"createdAt"`
	UpdatedAt           *time.Time `json:"updatedAt"`
}
func (obj *Datum) id() (ID) {	return string(obj.ID) }
func (obj *Datum) iri() (IRI) { path, _ := GetIRI(obj); return path }


// Field represents a template field in Koillection.
type Field struct {
	ID         ID     `json:"id"`
	Name       string `json:"name"`
	Position   int    `json:"position"`
	Type       string `json:"type"`
	ChoiceList *IRI   `json:"choiceList"`
	Template   *IRI   `json:"template"`
	Visibility string `json:"visibility"`
	Owner      *IRI   `json:"owner"`
}
func (obj *) id() (ID) {	return string(obj.ID) }
func (obj *) iri() (IRI) { path, _ := GetIRI(obj); return path }


// Inventory represents an inventory record in Koillection.
type Inventory struct {
	ID        ID         `json:"id"`
	Name      string     `json:"name"`
	Content   []string   `json:"content"`
	Owner     *IRI       `json:"owner"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
func (obj *Inventory) id() (ID) {	return string(obj.ID) }
func (obj *Inventory) iri() (IRI) { path, _ := GetIRI(obj); return path }


// Item represents an item within a collection.
type Item struct {
	ID                  ID         `json:"id"`
	Name                string     `json:"name"`
	Quantity            int        `json:"quantity"`
	Collection          *IRI       `json:"collection"`
	Owner               *IRI       `json:"owner"`
	Image               *string    `json:"image"`
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail"`
	ImageLargeThumbnail *string    `json:"imageLargeThumbnail"`
	SeenCounter         int        `json:"seenCounter"`
	Visibility          string     `json:"visibility"`
	ParentVisibility    *string    `json:"parentVisibility"`
	FinalVisibility     string     `json:"finalVisibility"`
	ScrapedFromURL      *string    `json:"scrapedFromUrl"`
	CreatedAt           time.Time  `json:"createdAt"`
	UpdatedAt           *time.Time `json:"updatedAt"`
}
func (obj *Item) id() (ID) {	return string(obj.ID) }
func (obj *Item) iri() (IRI) { path, _ := GetIRI(obj); return path }


// Loan represents a loan record in Koillection.
type Loan struct {
	ID         ID         `json:"id"`
	Item       *IRI       `json:"item"`
	LentTo     string     `json:"lentTo"`
	LentAt     time.Time  `json:"lentAt"`
	ReturnedAt *time.Time `json:"returnedAt"`
	Owner      *IRI       `json:"owner"`
}
func (obj *Loan) id() (ID) {	return string(obj.ID) }
func (obj *Loan) iri() (IRI) { path, _ := GetIRI(obj); return path }


// Log represents an action or event in Koillection.
type Log struct {
	ID            ID        `json:"id"`
	Type          *string   `json:"type"`
	LoggedAt      time.Time `json:"loggedAt"`
	ObjectID      ID        `json:"objectId"`
	ObjectLabel   string    `json:"objectLabel"`
	ObjectClass   string    `json:"objectClass"`
	ObjectDeleted bool      `json:"objectDeleted"`
	Owner         *IRI      `json:"owner"`
}
func (obj *Log) id() (ID) {	return string(obj.ID) }
func (obj *Log) iri() (IRI) { path, _ := GetIRI(obj); return path }


// Photo represents a photo in Koillection.
type Photo struct {
	ID                  ID         `json:"id"`
	Title               string     `json:"title"`
	Comment             *string    `json:"comment"`
	Place               *string    `json:"place"`
	Album               *IRI       `json:"album"`
	Owner               *IRI       `json:"owner"`
	Image               *string    `json:"image"`
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail"`
	TakenAt             *time.Time `json:"takenAt"`
	Visibility          string     `json:"visibility"`
	ParentVisibility    *string    `json:"parentVisibility"`
	FinalVisibility     string     `json:"finalVisibility"`
	CreatedAt           time.Time  `json:"createdAt"`
	UpdatedAt           *time.Time `json:"updatedAt"`
}
func (obj *Photo) id() (ID) {	return string(obj.ID) }
func (obj *Photo) iri() (IRI) { path, _ := GetIRI(obj); return path }


// Tag represents a tag in Koillection.
type Tag struct {
	ID                  ID         `json:"id"`
	Label               string     `json:"label"`
	Description         *string    `json:"description"`
	Image               *string    `json:"image"`
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail"`
	Owner               *IRI       `json:"owner"`
	Category            *IRI       `json:"category"`
	SeenCounter         int        `json:"seenCounter"`
	Visibility          string     `json:"visibility"`
	CreatedAt           time.Time  `json:"createdAt"`
	UpdatedAt           *time.Time `json:"updatedAt"`
}
func (obj *Tag) id() (ID) {	return string(obj.ID) }
func (obj *Tag) iri() (IRI) { path, _ := GetIRI(obj); return path }


// TagCategory represents a tag category in Koillection.
type TagCategory struct {
	ID          ID         `json:"id"`
	Label       string     `json:"label"`
	Description *string    `json:"description"`
	Color       string     `json:"color"`
	Owner       *IRI       `json:"owner"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}
func (obj *TagCategory) id() (ID) {	return string(obj.ID) }
func (obj *TagCategory) iri() (IRI) { path, _ := GetIRI(obj); return path }


// Template represents a template in Koillection.
type Template struct {
	ID        ID         `json:"id"`
	Name      string     `json:"name"`
	Owner     *IRI       `json:"owner"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
func (obj *Template) id() (ID) {	return string(obj.ID) }
func (obj *Template) iri() (IRI) { path, _ := GetIRI(obj); return path }


// User represents a user in Koillection.
type User struct {
	ID                           ID         `json:"id"`
	Username                     string     `json:"username"`
	Email                        string     `json:"email"`
	Avatar                       *string    `json:"avatar"`
	Currency                     string     `json:"currency"`
	Locale                       string     `json:"locale"`
	Timezone                     string     `json:"timezone"`
	DateFormat                   string     `json:"dateFormat"`
	DiskSpaceAllowed             int        `json:"diskSpaceAllowed"`
	Visibility                   string     `json:"visibility"`
	LastDateOfActivity           *time.Time `json:"lastDateOfActivity"`
	WishlistsFeatureEnabled      bool       `json:"wishlistsFeatureEnabled"`
	TagsFeatureEnabled           bool       `json:"tagsFeatureEnabled"`
	SignsFeatureEnabled          bool       `json:"signsFeatureEnabled"`
	AlbumsFeatureEnabled         bool       `json:"albumsFeatureEnabled"`
	LoansFeatureEnabled          bool       `json:"loansFeatureEnabled"`
	TemplatesFeatureEnabled      bool       `json:"templatesFeatureEnabled"`
	HistoryFeatureEnabled        bool       `json:"historyFeatureEnabled"`
	StatisticsFeatureEnabled     bool       `json:"statisticsFeatureEnabled"`
	ScrapingFeatureEnabled       bool       `json:"scrapingFeatureEnabled"`
	SearchInDataByDefaultEnabled bool       `json:"searchInDataByDefaultEnabled"`
	DisplayItemsNameInGridView   bool       `json:"displayItemsNameInGridView"`
	SearchResultsDisplayMode     string     `json:"searchResultsDisplayMode"`
	CreatedAt                    time.Time  `json:"createdAt"`
	UpdatedAt                    *time.Time `json:"updatedAt"`
}
func (obj *User) id() (ID) {	return string(obj.ID) }
func (obj *User) iri() (IRI) { path, _ := GetIRI(obj); return path }


// Wish represents a wish in Koillection.
type Wish struct {
	ID                  ID         `json:"id"`
	Name                string     `json:"name"`
	URL                 *string    `json:"url"`
	Price               *string    `json:"price"`
	Currency            *string    `json:"currency"`
	Wishlist            *IRI       `json:"wishlist"`
	Owner               *IRI       `json:"owner"`
	Comment             *string    `json:"comment"`
	Image               *string    `json:"image"`
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail"`
	Visibility          string     `json:"visibility"`
	ParentVisibility    *string    `json:"parentVisibility"`
	FinalVisibility     string     `json:"finalVisibility"`
	ScrapedFromURL      *string    `json:"scrapedFromUrl"`
	CreatedAt           time.Time  `json:"createdAt"`
	UpdatedAt           *time.Time `json:"updatedAt"`
}
func (obj *Wish) id() (ID) {	return string(obj.ID) }
func (obj *Wish) iri() (IRI) { path, _ := GetIRI(obj); return path }


// Wishlist represents a wishlist in Koillection.
type Wishlist struct {
	ID               ID         `json:"id"`
	Name             string     `json:"name"`
	Owner            *IRI       `json:"owner"`
	Color            string     `json:"color"`
	Parent           *IRI       `json:"parent"`
	Image            *string    `json:"image"`
	SeenCounter      int        `json:"seenCounter"`
	Visibility       string     `json:"visibility"`
	ParentVisibility *string    `json:"parentVisibility"`
	FinalVisibility  string     `json:"finalVisibility"`
	CreatedAt        time.Time  `json:"createdAt"`
	UpdatedAt        *time.Time `json:"updatedAt"`
}
func (obj *Wishlist) id() (ID) {	return string(obj.ID) }
func (obj *Wishlist) iri() (IRI) { path, _ := GetIRI(obj); return path }


// Metrics represents system or user-specific metrics in Koillection.
type Metrics map[string]string

// Client defines the interface for interacting with the Koillection REST API.
type Client interface {
	CheckLogin(ctx context.Context, username, password string) (string, error)                           // HTTP POST /api/authentication_token
	GetMetrics(ctx context.Context) (*Metrics, error)                                                    // HTTP GET /api/metrics
	CreateAlbum(ctx context.Context, album *Album) (*Album, error)                                       // HTTP POST /api/albums
	GetAlbum(ctx context.Context, id ID) (*Album, error)                                                 // HTTP GET /api/albums/{id}
	ListAlbums(ctx context.Context, page int) ([]*Album, error)                                          // HTTP GET /api/albums
	UpdateAlbum(ctx context.Context, album *Album) (*Album, error)                                       // HTTP PUT /api/albums/{id}
	PatchAlbum(ctx context.Context, album *Album) (*Album, error)                                        // HTTP PATCH /api/albums/{id}
	DeleteAlbum(ctx context.Context, album *Album) error                                                 // HTTP DELETE /api/albums/{id}
	ListAlbumChildren(ctx context.Context, album *Album, page int) ([]*Album, error)                     // HTTP GET /api/albums/{id}/children
	UploadAlbumImage(ctx context.Context, album *Album, file []byte) (*Album, error)                     // HTTP POST /api/albums/{id}/image
	GetAlbumParent(ctx context.Context, album *Album) (*Album, error)                                    // HTTP GET /api/albums/{id}/parent
	ListAlbumPhotos(ctx context.Context, album *Album, page int) ([]*Photo, error)                       // HTTP GET /api/albums/{id}/photos
	CreateChoiceList(ctx context.Context, choiceList *ChoiceList) (*ChoiceList, error)                   // HTTP POST /api/choice_lists
	GetChoiceList(ctx context.Context, choiceList *ChoiceList) (*ChoiceList, error)                      // HTTP GET /api/choice_lists/{id}
	ListChoiceLists(ctx context.Context, page int) ([]*ChoiceList, error)                                // HTTP GET /api/choice_lists
	UpdateChoiceList(ctx context.Context, choiceList *ChoiceList) (*ChoiceList, error)                   // HTTP PUT /api/choice_lists/{id}
	PatchChoiceList(ctx context.Context, choiceList *ChoiceList) (*ChoiceList, error)                    // HTTP PATCH /api/choice_lists/{id}
	DeleteChoiceList(ctx context.Context, choiceList *ChoiceList) error                                  // HTTP DELETE /api/choice_lists/{id}
	CreateCollection(ctx context.Context, collection *Collection) (*Collection, error)                   // HTTP POST /api/collections
	GetCollection(ctx context.Context, collection *Collection) (*Collection, error)                      // HTTP GET /api/collections/{id}
	ListCollections(ctx context.Context, page int) ([]*Collection, error)                                // HTTP GET /api/collections
	UpdateCollection(ctx context.Context, collection *Collection) (*Collection, error)                   // HTTP PUT /api/collections/{id}
	PatchCollection(ctx context.Context, collection *Collection) (*Collection, error)                    // HTTP PATCH /api/collections/{id}
	DeleteCollection(ctx context.Context, collection *Collection) error                                  // HTTP DELETE /api/collections/{id}
	ListCollectionChildren(ctx context.Context, collection *Collection, page int) ([]*Collection, error) // HTTP GET /api/collections/{id}/children
	UploadCollectionImage(ctx context.Context, collection *Collection, file []byte) (*Collection, error) // HTTP POST /api/collections/{id}/image
	GetCollectionParent(ctx context.Context, collection *Collection) (*Collection, error)                // HTTP GET /api/collections/{id}/parent
	ListCollectionItems(ctx context.Context, collection *Collection, page int) ([]*Item, error)          // HTTP GET /api/collections/{id}/items
	ListCollectionData(ctx context.Context, collection *Collection, page int) ([]*Datum, error)          // HTTP GET /api/collections/{id}/data
	GetCollectionDefaultTemplate(ctx context.Context, collection *Collection) (*Template, error)         // HTTP GET /api/collections/{id}/items_default_template
	CreateDatum(ctx context.Context, datum *Datum) (*Datum, error)                                       // HTTP POST /api/data
	GetDatum(ctx context.Context, datum *Datum) (*Datum, error)                                          // HTTP GET /api/data/{id}
	ListData(ctx context.Context, page int) ([]*Datum, error)                                            // HTTP GET /api/data
	UpdateDatum(ctx context.Context, datum *Datum) (*Datum, error)                                       // HTTP PUT /api/data/{id}
	PatchDatum(ctx context.Context, datum *Datum) (*Datum, error)                                        // HTTP PATCH /api/data/{id}
	DeleteDatum(ctx context.Context, datum *Datum) error                                                 // HTTP DELETE /api/data/{id}
	UploadDatumFile(ctx context.Context, datum *Datum, file []byte) (*Datum, error)                      // HTTP POST /api/data/{id}/file
	UploadDatumImage(ctx context.Context, datum *Datum, image []byte) (*Datum, error)                    // HTTP POST /api/data/{id}/image
	UploadDatumVideo(ctx context.Context, datum *Datum, video []byte) (*Datum, error)                    // HTTP POST /api/data/{id}/video
	GetDatumItem(ctx context.Context, datum *Datum) (*Item, error)                                       // HTTP GET /api/data/{id}/item
	GetDatumCollection(ctx context.Context, datum *Datum) (*Collection, error)                           // HTTP GET /api/data/{id}/collection
	CreateField(ctx context.Context, field *Field) (*Field, error)                                       // HTTP POST /api/fields
	GetField(ctx context.Context, field *Field) (*Field, error)                                          // HTTP GET /api/fields/{id}
	ListFields(ctx context.Context, page int) ([]*Field, error)                                          // HTTP GET /api/fields
	UpdateField(ctx context.Context, field *Field) (*Field, error)                                       // HTTP PUT /api/fields/{id}
	PatchField(ctx context.Context, field *Field) (*Field, error)                                        // HTTP PATCH /api/fields/{id}
	DeleteField(ctx context.Context, field *Field) error                                                 // HTTP DELETE /api/fields/{id}
	GetFieldTemplate(ctx context.Context, field *Field) (*Template, error)                               // HTTP GET /api/fields/{id}/template
	ListTemplateFields(ctx context.Context, template *Template, page int) ([]*Field, error)              // HTTP GET /api/templates/{id}/fields
	ListInventories(ctx context.Context, page int) ([]*Inventory, error)                                 // HTTP GET /api/inventories
	GetInventory(ctx context.Context, inventory *Inventory) (*Inventory, error)                          // HTTP GET /api/inventories/{id}
	DeleteInventory(ctx context.Context, inventory *Inventory) error                                     // HTTP DELETE /api/inventories/{id}
	CreateItem(ctx context.Context, item *Item) (*Item, error)                                           // HTTP POST /api/items
	GetItem(ctx context.Context, item *Item) (*Item, error)                                              // HTTP GET /api/items/{id}
	ListItems(ctx context.Context, page int) ([]*Item, error)                                            // HTTP GET /api/items
	UpdateItem(ctx context.Context, item *Item) (*Item, error)                                           // HTTP PUT /api/items/{id}
	PatchItem(ctx context.Context, item *Item) (*Item, error)                                            // HTTP PATCH /api/items/{id}
	DeleteItem(ctx context.Context, item *Item) error                                                    // HTTP DELETE /api/items/{id}
	UploadItemImage(ctx context.Context, item *Item, file []byte) (*Item, error)                         // HTTP POST /api/items/{id}/image
	ListItemRelatedItems(ctx context.Context, item *Item, page int) ([]*Item, error)                     // HTTP GET /api/items/{id}/related_items
	ListItemLoans(ctx context.Context, item *Item, page int) ([]*Loan, error)                            // HTTP GET /api/items/{id}/loans
	ListItemTags(ctx context.Context, item *Item, page int) ([]*Tag, error)                              // HTTP GET /api/items/{id}/tags
	ListItemData(ctx context.Context, item *Item, page int) ([]*Datum, error)                            // HTTP GET /api/items/{id}/data
	GetItemCollection(ctx context.Context, item *Item) (*Collection, error)                              // HTTP GET /api/items/{id}/collection
	CreateLoan(ctx context.Context, loan *Loan) (*Loan, error)                                           // HTTP POST /api/loans
	GetLoan(ctx context.Context, loan *Loan) (*Loan, error)                                              // HTTP GET /api/loans/{id}
	ListLoans(ctx context.Context, page int) ([]*Loan, error)                                            // HTTP GET /api/loans
	UpdateLoan(ctx context.Context, loan *Loan) (*Loan, error)                                           // HTTP PUT /api/loans/{id}
	PatchLoan(ctx context.Context, loan *Loan) (*Loan, error)                                            // HTTP PATCH /api/loans/{id}
	DeleteLoan(ctx context.Context, loan *Loan) error                                                    // HTTP DELETE /api/loans/{id}
	GetLoanItem(ctx context.Context, loan *Loan) (*Item, error)                                          // HTTP GET /api/loans/{id}/item
	GetLog(ctx context.Context, log *Log) (*Log, error)                                                  // HTTP GET /api/logs/{id}
	ListLogs(ctx context.Context, page int) ([]*Log, error)                                              // HTTP GET /api/logs
	CreatePhoto(ctx context.Context, photo *Photo) (*Photo, error)                                       // HTTP POST /api/photos
	GetPhoto(ctx context.Context, photo *Photo) (*Photo, error)                                          // HTTP GET /api/photos/{id}
	ListPhotos(ctx context.Context, page int) ([]*Photo, error)                                          // HTTP GET /api/photos
	UpdatePhoto(ctx context.Context, photo *Photo) (*Photo, error)                                       // HTTP PUT /api/photos/{id}
	PatchPhoto(ctx context.Context, photo *Photo) (*Photo, error)                                        // HTTP PATCH /api/photos/{id}
	DeletePhoto(ctx context.Context, photo *Photo) error                                                 // HTTP DELETE /api/photos/{id}
	UploadPhotoImage(ctx context.Context, photo *Photo, file []byte) (*Photo, error)                     // HTTP POST /api/photos/{id}/image
	GetPhotoAlbum(ctx context.Context, photo *Photo) (*Album, error)                                     // HTTP GET /api/photos/{id}/album
	CreateTag(ctx context.Context, tag *Tag) (*Tag, error)                                               // HTTP POST /api/tags
	GetTag(ctx context.Context, tag *Tag) (*Tag, error)                                                  // HTTP GET /api/tags/{id}
	ListTags(ctx context.Context, page int) ([]*Tag, error)                                              // HTTP GET /api/tags
	UpdateTag(ctx context.Context, tag *Tag) (*Tag, error)                                               // HTTP PUT /api/tags/{id}
	PatchTag(ctx context.Context, tag *Tag) (*Tag, error)                                                // HTTP PATCH /api/tags/{id}
	DeleteTag(ctx context.Context, tag *Tag) error                                                       // HTTP DELETE /api/tags/{id}
	UploadTagImage(ctx context.Context, tag *Tag, file []byte) (*Tag, error)                             // HTTP POST /api/tags/{id}/image
	ListTagItems(ctx context.Context, tag *Tag, page int) ([]*Item, error)                               // HTTP GET /api/tags/{id}/items
	GetCategoryOfTag(ctx context.Context, tag *Tag) (*TagCategory, error)                                // HTTP GET /api/tags/{id}/category
	CreateTagCategory(ctx context.Context, category *TagCategory) (*TagCategory, error)                  // HTTP POST /api/tag_categories
	GetTagCategory(ctx context.Context, category *TagCategory) (*TagCategory, error)                     // HTTP GET /api/tag_categories/{id}
	ListTagCategories(ctx context.Context, page int) ([]*TagCategory, error)                             // HTTP GET /api/tag_categories
	UpdateTagCategory(ctx context.Context, category *TagCategory) (*TagCategory, error)                  // HTTP PUT /api/tag_categories/{id}
	PatchTagCategory(ctx context.Context, category *TagCategory) (*TagCategory, error)                   // HTTP PATCH /api/tag_categories/{id}
	DeleteTagCategory(ctx context.Context, category *TagCategory) error                                  // HTTP DELETE /api/tag_categories/{id}
	ListTagCategoryTags(ctx context.Context, category *TagCategory, page int) ([]*Tag, error)            // HTTP GET /api/tag_categories/{id}/tags
	CreateTemplate(ctx context.Context, template *Template) (*Template, error)                           // HTTP POST /api/templates
	GetTemplate(ctx context.Context, template *Template) (*Template, error)                              // HTTP GET /api/templates/{id}
	ListTemplates(ctx context.Context, page int) ([]*Template, error)                                    // HTTP GET /api/templates
	UpdateTemplate(ctx context.Context, template *Template) (*Template, error)                           // HTTP PUT /api/templates/{id}
	PatchTemplate(ctx context.Context, template *Template) (*Template, error)                            // HTTP PATCH /api/templates/{id}
	DeleteTemplate(ctx context.Context, template *Template) error                                        // HTTP DELETE /api/templates/{id}
	GetUser(ctx context.Context, user *User) (*User, error)                                              // HTTP GET /api/users/{id}
	ListUsers(ctx context.Context, page int) ([]*User, error)                                            // HTTP GET /api/users
	CreateWish(ctx context.Context, wish *Wish) (*Wish, error)                                           // HTTP POST /api/wishes
	GetWish(ctx context.Context, wish *Wish) (*Wish, error)                                              // HTTP GET /api/wishes/{id}
	ListWishes(ctx context.Context, page int) ([]*Wish, error)                                           // HTTP GET /api/wishes
	UpdateWish(ctx context.Context, wish *Wish) (*Wish, error)                                           // HTTP PUT /api/wishes/{id}
	PatchWish(ctx context.Context, wish *Wish) (*Wish, error)                                            // HTTP PATCH /api/wishes/{id}
	DeleteWish(ctx context.Context, wish *Wish) error                                                    // HTTP DELETE /api/wishes/{id}
	UploadWishImage(ctx context.Context, wish *Wish, file []byte) (*Wish, error)                         // HTTP POST /api/wishes/{id}/image
	GetWishWishlist(ctx context.Context, wish *Wish) (*Wishlist, error)                                  // HTTP GET /api/wishes/{id}/wishlist
	CreateWishlist(ctx context.Context, wishlist *Wishlist) (*Wishlist, error)                           // HTTP POST /api/wishlists
	GetWishlist(ctx context.Context, wishlist *Wishlist) (*Wishlist, error)                              // HTTP GET /api/wishlists/{id}
	ListWishlists(ctx context.Context, page int) ([]*Wishlist, error)                                    // HTTP GET /api/wishlists
	UpdateWishlist(ctx context.Context, wishlist *Wishlist) (*Wishlist, error)                           // HTTP PUT /api/wishlists/{id}
	PatchWishlist(ctx context.Context, wishlist *Wishlist) (*Wishlist, error)                            // HTTP PATCH /api/wishlists/{id}
	DeleteWishlist(ctx context.Context, wishlist *Wishlist) error                                        // HTTP DELETE /api/wishlists/{id}
	ListWishlistWishes(ctx context.Context, wishlist *Wishlist, page int) ([]*Wish, error)               // HTTP GET /api/wishlists/{id}/wishes
	ListWishlistChildren(ctx context.Context, wishlist *Wishlist, page int) ([]*Wishlist, error)         // HTTP GET /api/wishlists/{id}/children
	UploadWishlistImage(ctx context.Context, wishlist *Wishlist, file []byte) (*Wishlist, error)         // HTTP POST /api/wishlists/{id}/image
	GetWishlistParent(ctx context.Context, wishlist *Wishlist) (*Wishlist, error)                        // HTTP GET /api/wishlists/{id}/parent
}

// iriTable maps object types to their IRI path templates.
var iriTable = map[reflect.Type]string{
	reflect.TypeOf(&Album{}):       "/api/albums/%s",
	reflect.TypeOf(&ChoiceList{}):  "/api/choice_lists/%s",
	reflect.TypeOf(&Collection{}):  "/api/collections/%s",
	reflect.TypeOf(&Datum{}):       "/api/data/%s",
	reflect.TypeOf(&Field{}):       "/api/fields/%s",
	reflect.TypeOf(&Inventory{}):   "/api/inventories/%s",
	reflect.TypeOf(&Item{}):        "/api/items/%s",
	reflect.TypeOf(&Loan{}):        "/api/loans/%s",
	reflect.TypeOf(&Log{}):         "/api/logs/%s",
	reflect.TypeOf(&Photo{}):       "/api/photos/%s",
	reflect.TypeOf(&Tag{}):         "/api/tags/%s",
	reflect.TypeOf(&TagCategory{}): "/api/tag_categories/%s",
	reflect.TypeOf(&Template{}):    "/api/templates/%s",
	reflect.TypeOf(&User{}):        "/api/users/%s",
	reflect.TypeOf(&Wish{}):        "/api/wishes/%s",
	reflect.TypeOf(&Wishlist{}):    "/api/wishlists/%s",
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
	if !idField.IsValid() || idField.Type() != reflect.TypeOf(ID("")) {
		return "", fmt.Errorf("object %T has no valid ID field", obj)
	}
	id := idField.Interface().(ID)
	if id == "" {
		return "", fmt.Errorf("object %T has empty ID", obj)
	}

	// Construct the IRI using the path template and ID.
	return IRI(fmt.Sprintf(pathTemplate, id)), nil
}
