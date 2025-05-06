package koillection

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"
)

// IRI represents an IRI reference in the Koillection API.
type IRI string

// Album represents an album in Koillection.
type Album struct {
	ID               string     `json:"id"`
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

// ChoiceList represents a predefined list of options in Koillection.
type ChoiceList struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Choices   []string   `json:"choices"`
	Owner     *IRI       `json:"owner"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

// Collection represents a collection in Koillection.
type Collection struct {
	ID                   string     `json:"id"`
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

// Datum represents a custom data field in Koillection.
type Datum struct {
	ID                  string     `json:"id"`
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

// Field represents a template field in Koillection.
type Field struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Position   int    `json:"position"`
	Type       string `json:"type"`
	ChoiceList *IRI   `json:"choiceList"`
	Template   *IRI   `json:"template"`
	Visibility string `json:"visibility"`
	Owner      *IRI   `json:"owner"`
}

// Inventory represents an inventory record in Koillection.
type Inventory struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Content   []string   `json:"content"`
	Owner     *IRI       `json:"owner"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

// Item represents an item within a collection.
type Item struct {
	ID                  string     `json:"id"`
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

// Loan represents a loan record in Koillection.
type Loan struct {
	ID         string     `json:"id"`
	Item       *IRI       `json:"item"`
	LentTo     string     `json:"lentTo"`
	LentAt     time.Time  `json:"lentAt"`
	ReturnedAt *time.Time `json:"returnedAt"`
	Owner      *IRI       `json:"owner"`
}

// Log represents an action or event in Koillection.
type Log struct {
	ID            string    `json:"id"`
	Type          *string   `json:"type"`
	LoggedAt      time.Time `json:"loggedAt"`
	ObjectID      string    `json:"objectId"`
	ObjectLabel   string    `json:"objectLabel"`
	ObjectClass   string    `json:"objectClass"`
	ObjectDeleted bool      `json:"objectDeleted"`
	Owner         *IRI      `json:"owner"`
}

// Photo represents a photo in Koillection.
type Photo struct {
	ID                  string     `json:"id"`
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

// Tag represents a tag in Koillection.
type Tag struct {
	ID                  string     `json:"id"`
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

// TagCategory represents a tag category in Koillection.
type TagCategory struct {
	ID          string     `json:"id"`
	Label       string     `json:"label"`
	Description *string    `json:"description"`
	Color       string     `json:"color"`
	Owner       *IRI       `json:"owner"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

// Template represents a template in Koillection.
type Template struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Owner     *IRI       `json:"owner"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

// User represents a user in Koillection.
type User struct {
	ID                           string     `json:"id"`
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

// Wish represents a wish in Koillection.
type Wish struct {
	ID                  string     `json:"id"`
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

// Wishlist represents a wishlist in Koillection.
type Wishlist struct {
	ID               string     `json:"id"`
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

// Metrics represents system or user-specific metrics in Koillection.
type Metrics map[string]string

// Client defines the interface for interacting with the Koillection REST API.
type Client interface {
	CheckLogin(ctx context.Context, username, password string) (string, error)
	GetMetrics(ctx context.Context) (*Metrics, error)
	CreateAlbum(ctx context.Context, album *Album) (*Album, error)
	GetAlbum(ctx context.Context, id string) (*Album, error)
	ListAlbums(ctx context.Context, page int) ([]*Album, error)
	UpdateAlbum(ctx context.Context, id string, album *Album) (*Album, error)
	PatchAlbum(ctx context.Context, id string, album *Album) (*Album, error)
	DeleteAlbum(ctx context.Context, id string) error
	ListAlbumChildren(ctx context.Context, id string, page int) ([]*Album, error)
	UploadAlbumImage(ctx context.Context, id string, file []byte) (*Album, error)
	GetAlbumParent(ctx context.Context, id string) (*Album, error)
	ListAlbumPhotos(ctx context.Context, id string, page int) ([]*Photo, error)
	CreateChoiceList(ctx context.Context, choiceList *ChoiceList) (*ChoiceList, error)
	GetChoiceList(ctx context.Context, id string) (*ChoiceList, error)
	ListChoiceLists(ctx context.Context, page int) ([]*ChoiceList, error)
	UpdateChoiceList(ctx context.Context, id string, choiceList *ChoiceList) (*ChoiceList, error)
	PatchChoiceList(ctx context.Context, id string, choiceList *ChoiceList) (*ChoiceList, error)
	DeleteChoiceList(ctx context.Context, id string) error
	CreateCollection(ctx context.Context, collection *Collection) (*Collection, error)
	GetCollection(ctx context.Context, id string) (*Collection, error)
	ListCollections(ctx context.Context, page int) ([]*Collection, error)
	UpdateCollection(ctx context.Context, id string, collection *Collection) (*Collection, error)
	PatchCollection(ctx context.Context, id string, collection *Collection) (*Collection, error)
	DeleteCollection(ctx context.Context, id string) error
	ListCollectionChildren(ctx context.Context, id string, page int) ([]*Collection, error)
	UploadCollectionImage(ctx context.Context, id string, file []byte) (*Collection, error)
	GetCollectionParent(ctx context.Context, id string) (*Collection, error)
	ListCollectionItems(ctx context.Context, id string, page int) ([]*Item, error)
	ListCollectionData(ctx context.Context, id string, page int) ([]*Datum, error)
	GetCollectionDefaultTemplate(ctx context.Context, id string) (*Template, error)
	CreateDatum(ctx context.Context, datum *Datum) (*Datum, error)
	GetDatum(ctx context.Context, id string) (*Datum, error)
	ListData(ctx context.Context, page int) ([]*Datum, error)
	UpdateDatum(ctx context.Context, id string, datum *Datum) (*Datum, error)
	PatchDatum(ctx context.Context, id string, datum *Datum) (*Datum, error)
	DeleteDatum(ctx context.Context, id string) error
	UploadDatumFile(ctx context.Context, id string, file []byte) (*Datum, error)
	UploadDatumImage(ctx context.Context, id string, image []byte) (*Datum, error)
	UploadDatumVideo(ctx context.Context, id string, video []byte) (*Datum, error)
	GetDatumItem(ctx context.Context, id string) (*Item, error)
	GetDatumCollection(ctx context.Context, id string) (*Collection, error)
	CreateField(ctx context.Context, field *Field) (*Field, error)
	GetField(ctx context.Context, id string) (*Field, error)
	ListFields(ctx context.Context, page int) ([]*Field, error)
	UpdateField(ctx context.Context, id string, field *Field) (*Field, error)
	PatchField(ctx context.Context, id string, field *Field) (*Field, error)
	DeleteField(ctx context.Context, id string) error
	GetFieldTemplate(ctx context.Context, id string) (*Template, error)
	ListTemplateFields(ctx context.Context, templateID string, page int) ([]*Field, error)
	ListInventories(ctx context.Context, page int) ([]*Inventory, error)
	GetInventory(ctx context.Context, id string) (*Inventory, error)
	DeleteInventory(ctx context.Context, id string) error
	CreateItem(ctx context.Context, item *Item) (*Item, error)
	GetItem(ctx context.Context, id string) (*Item, error)
	ListItems(ctx context.Context, page int) ([]*Item, error)
	UpdateItem(ctx context.Context, id string, item *Item) (*Item, error)
	PatchItem(ctx context.Context, id string, item *Item) (*Item, error)
	DeleteItem(ctx context.Context, id string) error
	UploadItemImage(ctx context.Context, id string, file []byte) (*Item, error)
	ListItemRelatedItems(ctx context.Context, id string, page int) ([]*Item, error)
	ListItemLoans(ctx context.Context, id string, page int) ([]*Loan, error)
	ListItemTags(ctx context.Context, id string, page int) ([]*Tag, error)
	ListItemData(ctx context.Context, id string, page int) ([]*Datum, error)
	GetItemCollection(ctx context.Context, id string) (*Collection, error)
	CreateLoan(ctx context.Context, loan *Loan) (*Loan, error)
	GetLoan(ctx context.Context, id string) (*Loan, error)
	ListLoans(ctx context.Context, page int) ([]*Loan, error)
	UpdateLoan(ctx context.Context, id string, loan *Loan) (*Loan, error)
	PatchLoan(ctx context.Context, id string, loan *Loan) (*Loan, error)
	DeleteLoan(ctx context.Context, id string) error
	GetLoanItem(ctx context.Context, id string) (*Item, error)
	GetLog(ctx context.Context, id string) (*Log, error)
	ListLogs(ctx context.Context, page int) ([]*Log, error)
	CreatePhoto(ctx context.Context, photo *Photo) (*Photo, error)
	GetPhoto(ctx context.Context, id string) (*Photo, error)
	ListPhotos(ctx context.Context, page int) ([]*Photo, error)
	UpdatePhoto(ctx context.Context, id string, photo *Photo) (*Photo, error)
	PatchPhoto(ctx context.Context, id string, photo *Photo) (*Photo, error)
	DeletePhoto(ctx context.Context, id string) error
	UploadPhotoImage(ctx context.Context, id string, file []byte) (*Photo, error)
	GetPhotoAlbum(ctx context.Context, id string) (*Album, error)
	CreateTag(ctx context.Context, tag *Tag) (*Tag, error)
	GetTag(ctx context.Context, id string) (*Tag, error)
	ListTags(ctx context.Context, page int) ([]*Tag, error)
	UpdateTag(ctx context.Context, id string, tag *Tag) (*Tag, error)
	PatchTag(ctx context.Context, id string, tag *Tag) (*Tag, error)
	DeleteTag(ctx context.Context, id string) error
	UploadTagImage(ctx context.Context, id string, file []byte) (*Tag, error)
	ListTagItems(ctx context.Context, id string, page int) ([]*Item, error)
	GetTagCategory(ctx context.Context, id string) (*TagCategory, error)
	CreateTagCategory(ctx context.Context, category *TagCategory) (*TagCategory, error)
	ListTagCategories(ctx context.Context, page int) ([]*TagCategory, error)
	UpdateTagCategory(ctx context.Context, id string, category *TagCategory) (*TagCategory, error)
	PatchTagCategory(ctx context.Context, id string, category *TagCategory) (*TagCategory, error)
	DeleteTagCategory(ctx context.Context, id string) error
	ListTagCategoryTags(ctx context.Context, id string, page int) ([]*Tag, error)
	CreateTemplate(ctx context.Context, template *Template) (*Template, error)
	GetTemplate(ctx context.Context, id string) (*Template, error)
	ListTemplates(ctx context.Context, page int) ([]*Template, error)
	UpdateTemplate(ctx context.Context, id string, template *Template) (*Template, error)
	PatchTemplate(ctx context.Context, id string, template *Template) (*Template, error)
	DeleteTemplate(ctx context.Context, id string) error
	GetUser(ctx context.Context, id string) (*User, error)
	ListUsers(ctx context.Context, page int) ([]*User, error)
	CreateWish(ctx context.Context, wish *Wish) (*Wish, error)
	GetWish(ctx context.Context, id string) (*Wish, error)
	ListWishes(ctx context.Context, page int) ([]*Wish, error)
	UpdateWish(ctx context.Context, id string, wish *Wish) (*Wish, error)
	PatchWish(ctx context.Context, id string, wish *Wish) (*Wish, error)
	DeleteWish(ctx context.Context, id string) error
	UploadWishImage(ctx context.Context, id string, file []byte) (*Wish, error)
	GetWishWishlist(ctx context.Context, id string) (*Wishlist, error)
	CreateWishlist(ctx context.Context, wishlist *Wishlist) (*Wishlist, error)
	GetWishlist(ctx context.Context, id string) (*Wishlist, error)
	ListWishlists(ctx context.Context, page int) ([]*Wishlist, error)
	UpdateWishlist(ctx context.Context, id string, wishlist *Wishlist) (*Wishlist, error)
	PatchWishlist(ctx context.Context, id string, wishlist *Wishlist) (*Wishlist, error)
	DeleteWishlist(ctx context.Context, id string) error
	ListWishlistWishes(ctx context.Context, id string, page int) ([]*Wish, error)
	ListWishlistChildren(ctx context.Context, id string, page int) ([]*Wishlist, error)
	UploadWishlistImage(ctx context.Context, id string, file []byte) (*Wishlist, error)
	GetWishlistParent(ctx context.Context, id string) (*Wishlist, error)
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
