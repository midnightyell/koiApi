package koillection

import (
	"context"
	"time"
)

// Album represents an album in Koillection.
type Album struct {
	ID               string     `json:"id"`
	Title            string     `json:"title"`
	Color            string     `json:"color"`
	Image            *string    `json:"image"`
	Owner            *string    `json:"owner"`
	Parent           *string    `json:"parent"`
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
	Owner     *string    `json:"owner"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

// Collection represents a collection in Koillection.
type Collection struct {
	ID                   string     `json:"id"`
	Title                string     `json:"title"`
	Parent               *string    `json:"parent"`
	Owner                *string    `json:"owner"`
	Color                string     `json:"color"`
	Image                *string    `json:"image"`
	SeenCounter          int        `json:"seenCounter"`
	ItemsDefaultTemplate *string    `json:"itemsDefaultTemplate"`
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
	Item                *string    `json:"item"`
	Collection          *string    `json:"collection"`
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
	ChoiceList          *string    `json:"choiceList"`
	Owner               *string    `json:"owner"`
	Visibility          string     `json:"visibility"`
	ParentVisibility    *string    `json:"parentVisibility"`
	FinalVisibility     string     `json:"finalVisibility"`
	CreatedAt           time.Time  `json:"createdAt"`
	UpdatedAt           *time.Time `json:"updatedAt"`
}

// Field represents a template field in Koillection.
type Field struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Position   int     `json:"position"`
	Type       string  `json:"type"`
	ChoiceList *string `json:"choiceList"`
	Template   *string `json:"template"`
	Visibility string  `json:"visibility"`
	Owner      *string `json:"owner"`
}

// Inventory represents an inventory record in Koillection.
type Inventory struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Content   []string   `json:"content"`
	Owner     *string    `json:"owner"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

// Item represents an item within a collection.
type Item struct {
	ID                  string     `json:"id"`
	Name                string     `json:"name"`
	Quantity            int        `json:"quantity"`
	Collection          *string    `json:"collection"`
	Owner               *string    `json:"owner"`
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
	Item       *string    `json:"item"`
	LentTo     string     `json:"lentTo"`
	LentAt     time.Time  `json:"lentAt"`
	ReturnedAt *time.Time `json:"returnedAt"`
	Owner      *string    `json:"owner"`
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
	Owner         *string   `json:"owner"`
}

// Photo represents a photo in Koillection.
type Photo struct {
	ID                  string     `json:"id"`
	Title               string     `json:"title"`
	Comment             *string    `json:"comment"`
	Place               *string    `json:"place"`
	Album               *string    `json:"album"`
	Owner               *string    `json:"owner"`
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
	Owner               *string    `json:"owner"`
	Category            *string    `json:"category"`
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
	Owner       *string    `json:"owner"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

// Template represents a template in Koillection.
type Template struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Owner     *string    `json:"owner"`
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
	Wishlist            *string    `json:"wishlist"`
	Owner               *string    `json:"owner"`
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
	Owner            *string    `json:"owner"`
	Color            string     `json:"color"`
	Parent           *string    `json:"parent"`
	Image            *string    `json:"image"`
	SeenCounter      int        `json:"seenCounter"`
	Visibility       string     `json:"visibility"`
	ParentVisibility *string    `json:"parentVisibility"`
	FinalVisibility  string     `json:"finalVisibility"`
	CreatedAt        time.Time  `json:"createdAt"`
	UpdatedAt        *time.Time `json:"updatedAt"`
}

/*
// Metrics represents system or user-specific metrics in Koillection.
type Metrics struct {
	CollectionsCount int `json:"collections_count"`
	ItemsCount       int `json:"items_count"`
	WishlistsCount   int `json:"wishlists_count"`
	UsersCount       int `json:"users_count"`
	RequestsCount    int `json:"requests_count"`
}
*/

// Client defines the interface for interacting with the Koillection REST API.
type Client interface {
	// CheckLogin authenticates a user and returns a JWT token.
	CheckLogin(ctx context.Context, username, password string) (string, error)

	// GetMetrics retrieves system or user-specific metrics.
	GetMetrics(ctx context.Context) (*Metrics, error)

	// CreateAlbum creates a new album.
	CreateAlbum(ctx context.Context, album *Album) (*Album, error)

	// GetAlbum retrieves an album by its ID.
	GetAlbum(ctx context.Context, id string) (*Album, error)

	// ListAlbums retrieves a list of albums.
	ListAlbums(ctx context.Context, page int) ([]*Album, error)

	// UpdateAlbum replaces an existing album.
	UpdateAlbum(ctx context.Context, id string, album *Album) (*Album, error)

	// PatchAlbum updates an existing album partially.
	PatchAlbum(ctx context.Context, id string, album *Album) (*Album, error)

	// DeleteAlbum deletes an album by its ID.
	DeleteAlbum(ctx context.Context, id string) error

	// ListAlbumChildren retrieves child albums of an album.
	ListAlbumChildren(ctx context.Context, id string, page int) ([]*Album, error)

	// UploadAlbumImage uploads an image for an album.
	UploadAlbumImage(ctx context.Context, id string, file []byte) (*Album, error)

	// GetAlbumParent retrieves the parent album of an album.
	GetAlbumParent(ctx context.Context, id string) (*Album, error)

	// ListAlbumPhotos retrieves photos in an album.
	ListAlbumPhotos(ctx context.Context, id string, page int) ([]*Photo, error)

	// CreateChoiceList creates a new choice list.
	CreateChoiceList(ctx context.Context, choiceList *ChoiceList) (*ChoiceList, error)

	// GetChoiceList retrieves a choice list by its ID.
	GetChoiceList(ctx context.Context, id string) (*ChoiceList, error)

	// ListChoiceLists retrieves a list of choice lists.
	ListChoiceLists(ctx context.Context, page int) ([]*ChoiceList, error)

	// UpdateChoiceList replaces an existing choice list.
	UpdateChoiceList(ctx context.Context, id string, choiceList *ChoiceList) (*ChoiceList, error)

	// PatchChoiceList updates an existing choice list partially.
	PatchChoiceList(ctx context.Context, id string, choiceList *ChoiceList) (*ChoiceList, error)

	// DeleteChoiceList deletes a choice list by its ID.
	DeleteChoiceList(ctx context.Context, id string) error

	// CreateCollection creates a new collection.
	CreateCollection(ctx context.Context, collection *Collection) (*Collection, error)

	// GetCollection retrieves a collection by its ID.
	GetCollection(ctx context.Context, id string) (*Collection, error)

	// ListCollections retrieves a list of collections.
	ListCollections(ctx context.Context, page int) ([]*Collection, error)

	// UpdateCollection replaces an existing collection.
	UpdateCollection(ctx context.Context, id string, collection *Collection) (*Collection, error)

	// PatchCollection updates an existing collection partially.
	PatchCollection(ctx context.Context, id string, collection *Collection) (*Collection, error)

	// DeleteCollection deletes a collection by its ID.
	DeleteCollection(ctx context.Context, id string) error

	// ListCollectionChildren retrieves child collections of a collection.
	ListCollectionChildren(ctx context.Context, id string, page int) ([]*Collection, error)

	// UploadCollectionImage uploads an image for a collection.
	UploadCollectionImage(ctx context.Context, id string, file []byte) (*Collection, error)

	// GetCollectionParent retrieves the parent collection of a collection.
	GetCollectionParent(ctx context.Context, id string) (*Collection, error)

	// ListCollectionItems retrieves items in a collection.
	ListCollectionItems(ctx context.Context, id string, page int) ([]*Item, error)

	// ListCollectionData retrieves data fields in a collection.
	ListCollectionData(ctx context.Context, id string, page int) ([]*Datum, error)

	// GetCollectionDefaultTemplate retrieves the default template for items in a collection.
	GetCollectionDefaultTemplate(ctx context.Context, id string) (*Template, error)

	// CreateDatum creates a new datum.
	CreateDatum(ctx context.Context, datum *Datum) (*Datum, error)

	// GetDatum retrieves a datum by its ID.
	GetDatum(ctx context.Context, id string) (*Datum, error)

	// ListData retrieves a list of data fields.
	ListData(ctx context.Context, page int) ([]*Datum, error)

	// UpdateDatum replaces an existing datum.
	UpdateDatum(ctx context.Context, id string, datum *Datum) (*Datum, error)

	// PatchDatum updates an existing datum partially.
	PatchDatum(ctx context.Context, id string, datum *Datum) (*Datum, error)

	// DeleteDatum deletes a datum by its ID.
	DeleteDatum(ctx context.Context, id string) error

	// UploadDatumFile uploads a file for a datum.
	UploadDatumFile(ctx context.Context, id string, file []byte) (*Datum, error)

	// UploadDatumImage uploads an image for a datum.
	UploadDatumImage(ctx context.Context, id string, image []byte) (*Datum, error)

	// UploadDatumVideo uploads a video for a datum.
	UploadDatumVideo(ctx context.Context, id string, video []byte) (*Datum, error)

	// GetDatumItem retrieves the item associated with a datum.
	GetDatumItem(ctx context.Context, id string) (*Item, error)

	// GetDatumCollection retrieves the collection associated with a datum.
	GetDatumCollection(ctx context.Context, id string) (*Collection, error)

	// CreateField creates a new field.
	CreateField(ctx context.Context, field *Field) (*Field, error)

	// GetField retrieves a field by its ID.
	GetField(ctx context.Context, id string) (*Field, error)

	// ListFields retrieves a list of fields.
	ListFields(ctx context.Context, page int) ([]*Field, error)

	// UpdateField replaces an existing field.
	UpdateField(ctx context.Context, id string, field *Field) (*Field, error)

	// PatchField updates an existing field partially.
	PatchField(ctx context.Context, id string, field *Field) (*Field, error)

	// DeleteField deletes a field by its ID.
	DeleteField(ctx context.Context, id string) error

	// GetFieldTemplate retrieves the template associated with a field.
	GetFieldTemplate(ctx context.Context, id string) (*Template, error)

	// ListTemplateFields retrieves fields associated with a template.
	ListTemplateFields(ctx context.Context, templateID string, page int) ([]*Field, error)

	// ListInventories retrieves a list of inventories.
	ListInventories(ctx context.Context, page int) ([]*Inventory, error)

	// GetInventory retrieves an inventory by its ID.
	GetInventory(ctx context.Context, id string) (*Inventory, error)

	// DeleteInventory deletes an inventory by its ID.
	DeleteInventory(ctx context.Context, id string) error

	// CreateItem creates a new item.
	CreateItem(ctx context.Context, item *Item) (*Item, error)

	// GetItem retrieves an item by its ID.
	GetItem(ctx context.Context, id string) (*Item, error)

	// ListItems retrieves a list of items.
	ListItems(ctx context.Context, page int) ([]*Item, error)

	// UpdateItem replaces an existing item.
	UpdateItem(ctx context.Context, id string, item *Item) (*Item, error)

	// PatchItem updates an existing item partially.
	PatchItem(ctx context.Context, id string, item *Item) (*Item, error)

	// DeleteItem deletes an item by its ID.
	DeleteItem(ctx context.Context, id string) error

	// UploadItemImage uploads an image for an item.
	UploadItemImage(ctx context.Context, id string, file []byte) (*Item, error)

	// ListItemRelatedItems retrieves related items for an item.
	ListItemRelatedItems(ctx context.Context, id string, page int) ([]*Item, error)

	// ListItemLoans retrieves loans for an item.
	ListItemLoans(ctx context.Context, id string, page int) ([]*Loan, error)

	// ListItemTags retrieves tags for an item.
	ListItemTags(ctx context.Context, id string, page int) ([]*Tag, error)

	// ListItemData retrieves data fields for an item.
	ListItemData(ctx context.Context, id string, page int) ([]*Datum, error)

	// GetItemCollection retrieves the collection associated with an item.
	GetItemCollection(ctx context.Context, id string) (*Collection, error)

	// CreateLoan creates a new loan.
	CreateLoan(ctx context.Context, loan *Loan) (*Loan, error)

	// GetLoan retrieves a loan by its ID.
	GetLoan(ctx context.Context, id string) (*Loan, error)

	// ListLoans retrieves a list of loans.
	ListLoans(ctx context.Context, page int) ([]*Loan, error)

	// UpdateLoan replaces an existing loan.
	UpdateLoan(ctx context.Context, id string, loan *Loan) (*Loan, error)

	// PatchLoan updates an existing loan partially.
	PatchLoan(ctx context.Context, id string, loan *Loan) (*Loan, error)

	// DeleteLoan deletes a loan by its ID.
	DeleteLoan(ctx context.Context, id string) error

	// GetLoanItem retrieves the item associated with a loan.
	GetLoanItem(ctx context.Context, id string) (*Item, error)

	// GetLog retrieves a log by its ID.
	GetLog(ctx context.Context, id string) (*Log, error)

	// ListLogs retrieves a list of logs.
	ListLogs(ctx context.Context, page int) ([]*Log, error)

	// CreatePhoto creates a new photo.
	CreatePhoto(ctx context.Context, photo *Photo) (*Photo, error)

	// GetPhoto retrieves a photo by its ID.
	GetPhoto(ctx context.Context, id string) (*Photo, error)

	// ListPhotos retrieves a list of photos.
	ListPhotos(ctx context.Context, page int) ([]*Photo, error)

	// UpdatePhoto replaces an existing photo.
	UpdatePhoto(ctx context.Context, id string, photo *Photo) (*Photo, error)

	// PatchPhoto updates an existing photo partially.
	PatchPhoto(ctx context.Context, id string, photo *Photo) (*Photo, error)

	// DeletePhoto deletes a photo by its ID.
	DeletePhoto(ctx context.Context, id string) error

	// UploadPhotoImage uploads an image for a photo.
	UploadPhotoImage(ctx context.Context, id string, file []byte) (*Photo, error)

	// GetPhotoAlbum retrieves the album associated with a photo.
	GetPhotoAlbum(ctx context.Context, id string) (*Album, error)

	// CreateTag creates a new tag.
	CreateTag(ctx context.Context, tag *Tag) (*Tag, error)

	// GetTag retrieves a tag by its ID.
	GetTag(ctx context.Context, id string) (*Tag, error)

	// ListTags retrieves a list of tags.
	ListTags(ctx context.Context, page int) ([]*Tag, error)

	// UpdateTag replaces an existing tag.
	UpdateTag(ctx context.Context, id string, tag *Tag) (*Tag, error)

	// PatchTag updates an existing tag partially.
	PatchTag(ctx context.Context, id string, tag *Tag) (*Tag, error)

	// DeleteTag deletes a tag by its ID.
	DeleteTag(ctx context.Context, id string) error

	// UploadTagImage uploads an image for a tag.
	UploadTagImage(ctx context.Context, id string, file []byte) (*Tag, error)

	// ListTagItems retrieves items associated with a tag.
	ListTagItems(ctx context.Context, id string, page int) ([]*Item, error)

	// GetTagCategory retrieves the category associated with a tag.
	GetTagCategory(ctx context.Context, id string) (*TagCategory, error)

	// CreateTagCategory creates a new tag category.
	CreateTagCategory(ctx context.Context, category *TagCategory) (*TagCategory, error)

	// GetTagCategory retrieves a tag category by its ID.
	GetTagCategory(ctx context.Context, id string) (*TagCategory, error)

	// ListTagCategories retrieves a list of tag categories.
	ListTagCategories(ctx context.Context, page int) ([]*TagCategory, error)

	// UpdateTagCategory replaces an existing tag category.
	UpdateTagCategory(ctx context.Context, id string, category *TagCategory) (*TagCategory, error)

	// PatchTagCategory updates an existing tag category partially.
	PatchTagCategory(ctx context.Context, id string, category *TagCategory) (*TagCategory, error)

	// DeleteTagCategory deletes a tag category by its ID.
	DeleteTagCategory(ctx context.Context, id string) error

	// ListTagCategoryTags retrieves tags in a tag category.
	ListTagCategoryTags(ctx context.Context, id string, page int) ([]*Tag, error)

	// CreateTemplate creates a new template.
	CreateTemplate(ctx context.Context, template *Template) (*Template, error)

	// GetTemplate retrieves a template by its ID.
	GetTemplate(ctx context.Context, id string) (*Template, error)

	// ListTemplates retrieves a list of templates.
	ListTemplates(ctx context.Context, page int) ([]*Template, error)

	// UpdateTemplate replaces an existing template.
	UpdateTemplate(ctx context.Context, id string, template *Template) (*Template, error)

	// PatchTemplate updates an existing template partially.
	PatchTemplate(ctx context.Context, id string, template *Template) (*Template, error)

	// DeleteTemplate deletes a template by its ID.
	DeleteTemplate(ctx context.Context, id string) error

	// GetUser retrieves a user by its ID.
	GetUser(ctx context.Context, id string) (*User, error)

	// ListUsers retrieves a list of users.
	ListUsers(ctx context.Context, page int) ([]*User, error)

	// CreateWish creates a new wish.
	CreateWish(ctx context.Context, wish *Wish) (*Wish, error)

	// GetWish retrieves a wish by its ID.
	GetWish(ctx context.Context, id string) (*Wish, error)

	// ListWishes retrieves a list of wishes.
	ListWishes(ctx context.Context, page int) ([]*Wish, error)

	// UpdateWish replaces an existing wish.
	UpdateWish(ctx context.Context, id string, wish *Wish) (*Wish, error)

	// PatchWish updates an existing wish partially.
	PatchWish(ctx context.Context, id string, wish *Wish) (*Wish, error)

	// DeleteWish deletes a wish by its ID.
	DeleteWish(ctx context.Context, id string) error

	// UploadWishImage uploads an image for a wish.
	UploadWishImage(ctx context.Context, id string, file []byte) (*Wish, error)

	// GetWishWishlist retrieves the wishlist associated with a wish.
	GetWishWishlist(ctx context.Context, id string) (*Wishlist, error)

	// CreateWishlist creates a new wishlist.
	CreateWishlist(ctx context.Context, wishlist *Wishlist) (*Wishlist, error)

	// GetWishlist retrieves a wishlist by its ID.
	GetWishlist(ctx context.Context, id string) (*Wishlist, error)

	// ListWishlists retrieves a list of wishlists.
	ListWishlists(ctx context.Context, page int) ([]*Wishlist, error)

	// UpdateWishlist replaces an existing wishlist.
	UpdateWishlist(ctx context.Context, id string, wishlist *Wishlist) (*Wishlist, error)

	// PatchWishlist updates an existing wishlist partially.
	PatchWishlist(ctx context.Context, id string, wishlist *Wishlist) (*Wishlist, error)

	// DeleteWishlist deletes a wishlist by its ID.
	DeleteWishlist(ctx context.Context, id string) error

	// ListWishlistWishes retrieves wishes in a wishlist.
	ListWishlistWishes(ctx context.Context, id string, page int) ([]*Wish, error)

	// ListWishlistChildren retrieves child wishlists of a wishlist.
	ListWishlistChildren(ctx context.Context, id string, page int) ([]*Wishlist, error)

	// UploadWishlistImage uploads an image for a wishlist.
	UploadWishlistImage(ctx context.Context, id string, file []byte) (*Wishlist, error)

	// GetWishlistParent retrieves the parent wishlist of a wishlist.
	GetWishlistParent(ctx context.Context, id string) (*Wishlist, error)
}
