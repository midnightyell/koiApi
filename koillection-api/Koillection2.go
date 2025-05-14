package koiApi

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/currency"
)

// ID represents a unique identifier for resources (JSON-LD @id or JSON id).
type ID string // read-only (maps to @id or id)

// Metrics represents a map of metrics data.
type Metrics map[string]string // read-only

// Context represents the JSON-LD @context field.
type Context string // JSON-LD only

// Visibility represents the visibility level of a resource.
type Visibility string // read and write

const (
	VisibilityPublic   Visibility = "public" // Default for most resources
	VisibilityInternal Visibility = "internal"
	VisibilityPrivate  Visibility = "private" // Default for User
)

func (v Visibility) String() string {
	return string(v)
}

// DatumType represents the type of a custom data field.
type DatumType string // read and write

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
type FieldType string // read and write

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
type DateFormat string // read and write (assumed)

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
	CheckLogin(ctx context.Context, username, password string) (string, error)
	GetMetrics(ctx context.Context) (*Metrics, error)
	CreateAlbum(ctx context.Context, album *Album) (*Album, error)
	GetAlbum(ctx context.Context, id ID) (*Album, error)
	ListAlbums(ctx context.Context, page int) ([]*Album, error)
	UpdateAlbum(ctx context.Context, id ID, album *Album) (*Album, error)
	PatchAlbum(ctx context.Context, id ID, album *Album) (*Album, error)
	DeleteAlbum(ctx context.Context, id ID) error
	ListAlbumChildren(ctx context.Context, id ID, page int) ([]*Album, error)
	UploadAlbumImage(ctx context.Context, id ID, file []byte) (*Album, error)
	GetAlbumParent(ctx context.Context, id ID) (*Album, error)
	ListAlbumPhotos(ctx context.Context, id ID, page int) ([]*Photo, error)
	CreateChoiceList(ctx context.Context, choiceList *ChoiceList) (*ChoiceList, error)
	GetChoiceList(ctx context.Context, id ID) (*ChoiceList, error)
	ListChoiceLists(ctx context.Context, page int) ([]*ChoiceList, error)
	UpdateChoiceList(ctx context.Context, id ID, choiceList *ChoiceList) (*ChoiceList, error)
	PatchChoiceList(ctx context.Context, id ID, choiceList *ChoiceList) (*ChoiceList, error)
	DeleteChoiceList(ctx context.Context, id ID) error
	CreateCollection(ctx context.Context, collection *Collection) (*Collection, error)
	GetCollection(ctx context.Context, id ID) (*Collection, error)
	ListCollections(ctx context.Context, page int) ([]*Collection, error)
	UpdateCollection(ctx context.Context, id ID, collection *Collection) (*Collection, error)
	PatchCollection(ctx context.Context, id ID, collection *Collection) (*Collection, error)
	DeleteCollection(ctx context.Context, id ID) error
	ListCollectionChildren(ctx context.Context, id ID, page int) ([]*Collection, error)
	UploadCollectionImage(ctx context.Context, id ID, file []byte) (*Collection, error)
	GetCollectionParent(ctx context.Context, id ID) (*Collection, error)
	ListCollectionItems(ctx context.Context, id ID, page int) ([]*Item, error)
	ListCollectionData(ctx context.Context, id ID, page int) ([]*Datum, error)
	GetCollectionDefaultTemplate(ctx context.Context, id ID) (*Template, error)
	CreateDatum(ctx context.Context, datum *Datum) (*Datum, error)
	GetDatum(ctx context.Context, id ID) (*Datum, error)
	ListData(ctx context.Context, page int) ([]*Datum, error)
	UpdateDatum(ctx context.Context, id ID, datum *Datum) (*Datum, error)
	PatchDatum(ctx context.Context, id ID, datum *Datum) (*Datum, error)
	DeleteDatum(ctx context.Context, id ID) error
	UploadDatumFile(ctx context.Context, id ID, file []byte) (*Datum, error)
	UploadDatumImage(ctx context.Context, id ID, image []byte) (*Datum, error)
	UploadDatumVideo(ctx context.Context, id ID, video []byte) (*Datum, error)
	GetDatumItem(ctx context.Context, id ID) (*Item, error)
	GetDatumCollection(ctx context.Context, id ID) (*Collection, error)
	CreateField(ctx context.Context, field *Field) (*Field, error)
	GetField(ctx context.Context, id ID) (*Field, error)
	ListFields(ctx context.Context, page int) ([]*Field, error)
	UpdateField(ctx context.Context, id ID, field *Field) (*Field, error)
	PatchField(ctx context.Context, id ID, field *Field) (*Field, error)
	DeleteField(ctx context.Context, id ID) error
	GetFieldTemplate(ctx context.Context, id ID) (*Template, error)
	ListTemplateFields(ctx context.Context, templateid ID, page int) ([]*Field, error)
	ListInventories(ctx context.Context, page int) ([]*Inventory, error)
	GetInventory(ctx context.Context, id ID) (*Inventory, error)
	DeleteInventory(ctx context.Context, id ID) error
	CreateItem(ctx context.Context, item *Item) (*Item, error)
	GetItem(ctx context.Context, id ID) (*Item, error)
	ListItems(ctx context.Context, page int) ([]*Item, error)
	UpdateItem(ctx context.Context, id ID, item *Item) (*Item, error)
	PatchItem(ctx context.Context, id ID, item *Item) (*Item, error)
	DeleteItem(ctx context.Context, id ID) error
	UploadItemImage(ctx context.Context, id ID, file []byte) (*Item, error)
	ListItemRelatedItems(ctx context.Context, id ID, page int) ([]*Item, error)
	ListItemLoans(ctx context.Context, id ID, page int) ([]*Loan, error)
	ListItemTags(ctx context.Context, id ID, page int) ([]*Tag, error)
	ListItemData(ctx context.Context, id ID, page int) ([]*Datum, error)
	GetItemCollection(ctx context.Context, id ID) (*Collection, error)
	CreateLoan(ctx context.Context, loan *Loan) (*Loan, error)
	GetLoan(ctx context.Context, id ID) (*Loan, error)
	ListLoans(ctx context.Context, page int) ([]*Loan, error)
	UpdateLoan(ctx context.Context, id ID, loan *Loan) (*Loan, error)
	PatchLoan(ctx context.Context, id ID, loan *Loan) (*Loan, error)
	DeleteLoan(ctx context.Context, id ID) error
	GetLoanItem(ctx context.Context, id ID) (*Item, error)
	GetLog(ctx context.Context, id ID) (*Log, error)
	ListLogs(ctx context.Context, page int) ([]*Log, error)
	CreatePhoto(ctx context.Context, photo *Photo) (*Photo, error)
	GetPhoto(ctx context.Context, id ID) (*Photo, error)
	ListPhotos(ctx context.Context, page int) ([]*Photo, error)
	UpdatePhoto(ctx context.Context, id ID, photo *Photo) (*Photo, error)
	PatchPhoto(ctx context.Context, id ID, photo *Photo) (*Photo, error)
	DeletePhoto(ctx context.Context, id ID) error
	UploadPhotoImage(ctx context.Context, id ID, file []byte) (*Photo, error)
	GetPhotoAlbum(ctx context.Context, id ID) (*Album, error)
	CreateTag(ctx context.Context, tag *Tag) (*Tag, error)
	GetTag(ctx context.Context, id ID) (*Tag, error)
	ListTags(ctx context.Context, page int) ([]*Tag, error)
	UpdateTag(ctx context.Context, id ID, tag *Tag) (*Tag, error)
	PatchTag(ctx context.Context, id ID, tag *Tag) (*Tag, error)
	DeleteTag(ctx context.Context, id ID) error
	UploadTagImage(ctx context.Context, id ID, file []byte) (*Tag, error)
	ListTagItems(ctx context.Context, id ID, page int) ([]*Item, error)
	GetCategoryOfTag(ctx context.Context, id ID) (*TagCategory, error)
	CreateTagCategory(ctx context.Context, category *TagCategory) (*TagCategory, error)
	GetTagCategory(ctx context.Context, id ID) (*TagCategory, error)
	ListTagCategories(ctx context.Context, page int) ([]*TagCategory, error)
	UpdateTagCategory(ctx context.Context, id ID, category *TagCategory) (*TagCategory, error)
	PatchTagCategory(ctx context.Context, id ID, category *TagCategory) (*TagCategory, error)
	DeleteTagCategory(ctx context.Context, id ID) error
	ListTagCategoryTags(ctx context.Context, id ID, page int) ([]*Tag, error)
	CreateTemplate(ctx context.Context, template *Template) (*Template, error)
	GetTemplate(ctx context.Context, id ID) (*Template, error)
	ListTemplates(ctx context.Context, page int) ([]*Template, error)
	UpdateTemplate(ctx context.Context, id ID, template *Template) (*Template, error)
	PatchTemplate(ctx context.Context, id ID, template *Template) (*Template, error)
	DeleteTemplate(ctx context.Context, id ID) error
	GetUser(ctx context.Context, id ID) (*User, error)
	ListUsers(ctx context.Context, page int) ([]*User, error)
	CreateWish(ctx context.Context, wish *Wish) (*Wish, error)
	GetWish(ctx context.Context, id ID) (*Wish, error)
	ListWishes(ctx context.Context, page int) ([]*Wish, error)
	UpdateWish(ctx context.Context, id ID, wish *Wish) (*Wish, error)
	PatchWish(ctx context.Context, id ID, wish *Wish) (*Wish, error)
	DeleteWish(ctx context.Context, id ID) error
	UploadWishImage(ctx context.Context, id ID, file []byte) (*Wish, error)
	GetWishWishlist(ctx context.Context, id ID) (*Wishlist, error)
	CreateWishlist(ctx context.Context, wishlist *Wishlist) (*Wishlist, error)
	GetWishlist(ctx context.Context, id ID) (*Wishlist, error)
	ListWishlists(ctx context.Context, page int) ([]*Wishlist, error)
	UpdateWishlist(ctx context.Context, id ID, wishlist *Wishlist) (*Wishlist, error)
	PatchWishlist(ctx context.Context, id ID, wishlist *Wishlist) (*Wishlist, error)
	DeleteWishlist(ctx context.Context, id ID) error
	ListWishlistWishes(ctx context.Context, id ID, page int) ([]*Wish, error)
	ListWishlistChildren(ctx context.Context, id ID, page int) ([]*Wishlist, error)
	UploadWishlistImage(ctx context.Context, id ID, file []byte) (*Wishlist, error)
	GetWishlistParent(ctx context.Context, id ID) (*Wishlist, error)
}

// AlbumInterface defines methods for interacting with Album resources.
type AlbumInterface interface {
	Create(ctx context.Context, client Client) (*Album, error)
	Get(ctx context.Context, client Client, id ID) (*Album, error)
	List(ctx context.Context, client Client) ([]*Album, error)
	Update(ctx context.Context, client Client, id ID) (*Album, error)
	Patch(ctx context.Context, client Client, id ID) (*Album, error)
	Delete(ctx context.Context, client Client, id ID) error
	ListChildren(ctx context.Context, client Client, id ID) ([]*Album, error)
	UploadImage(ctx context.Context, client Client, id ID, file []byte) (*Album, error)
	GetParent(ctx context.Context, client Client, id ID) (*Album, error)
	ListPhotos(ctx context.Context, client Client, id ID) ([]*Photo, error)
	Validate(ctx context.Context, client Client) error
	IRI() string
}

// Album represents an album in Koillection, combining read and write fields.
type Album struct {
	Context          *Context   `json:"@context,omitempty" access:"rw"`         // JSON-LD only
	_ID              ID         `json:"@id,omitempty" access:"ro"`              // JSON-LD, read-only
	ID               ID         `json:"id,omitempty" access:"ro"`               // read-only
	Type             string     `json:"@type,omitempty" access:"rw"`            // JSON-LD only
	Title            string     `json:"title" access:"rw"`                      // read and write
	Color            string     `json:"color,omitempty" access:"ro"`            // read-only
	Image            *string    `json:"image,omitempty" access:"ro"`            // read-only
	Owner            *string    `json:"owner,omitempty" access:"ro"`            // read-only, IRI
	Parent           *string    `json:"parent,omitempty" access:"rw"`           // read and write, IRI
	SeenCounter      int        `json:"seenCounter,omitempty" access:"ro"`      // read-only
	Visibility       Visibility `json:"visibility,omitempty" access:"rw"`       // read and write
	ParentVisibility *string    `json:"parentVisibility,omitempty" access:"ro"` // read-only
	FinalVisibility  Visibility `json:"finalVisibility,omitempty" access:"ro"`  // read-only
	CreatedAt        time.Time  `json:"createdAt" access:"ro"`                  // read-only
	UpdatedAt        *time.Time `json:"updatedAt,omitempty" access:"ro"`        // read-only
	File             *string    `json:"file,omitempty" access:"wo"`             // write-only, binary data via multipart form
	DeleteImage      *bool      `json:"deleteImage,omitempty" access:"wo"`      // write-only
}

// Validate checks the Album's fields for validity, using ctx for cancellation and client for optional IRI validation.
func (a *Album) Validate(ctx context.Context, client Client) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("validation cancelled: %w", err)
	}

	if a.Title == "" {
		return fmt.Errorf("title must not be empty")
	}

	switch a.Visibility {
	case VisibilityPublic, VisibilityInternal, VisibilityPrivate, "":
	default:
		return fmt.Errorf("invalid visibility value: %s", a.Visibility)
	}

	if a.Parent != nil {
		if *a.Parent == "" {
			return fmt.Errorf("parent IRI must not be empty if set")
		}
		if !strings.HasPrefix(*a.Parent, "/api/albums/") {
			return fmt.Errorf("parent IRI must start with /api/albums/: %s", *a.Parent)
		}
		if client != nil {
			parts := strings.Split(*a.Parent, "/")
			if len(parts) < 4 {
				return fmt.Errorf("invalid parent IRI format: %s", *a.Parent)
			}
			parentID := ID(parts[3])
			_, err := client.GetAlbum(ctx, parentID)
			if err != nil {
				return fmt.Errorf("invalid parent album %s: %w", *a.Parent, err)
			}
		}
	}

	if a.File != nil && *a.File == "" {
		return fmt.Errorf("file must not be empty if set")
	}

	if a.ID == "" && a._ID == "" {
		if a._ID != "" {
			return fmt.Errorf("_ID must be empty for creation")
		}
		if a.ID != "" {
			return fmt.Errorf("ID must be empty for creation")
		}
		if a.Type != "" && a.Type != "Album" {
			return fmt.Errorf("Type must be empty or 'Album' for creation: %s", a.Type)
		}
	} else {
		if a.ID == "" {
			return fmt.Errorf("ID must not be empty for update")
		}
	}

	return nil
}

// IRI returns the JSON-LD IRI for the Album.
func (a *Album) IRI() string {
	if a.ID != "" {
		return fmt.Sprintf("/api/albums/%s", a.ID)
	}
	if a._ID != "" {
		return fmt.Sprintf("/api/albums/%s", a._ID)
	}
	return ""
}

// Create calls Client.CreateAlbum to create a new Album.
func (a *Album) Create(ctx context.Context, client Client) (*Album, error) {
	return client.CreateAlbum(ctx, a)
}

// Get retrieves an Album by ID using Client.GetAlbum.
func (a *Album) Get(ctx context.Context, client Client, id ID) (*Album, error) {
	return client.GetAlbum(ctx, id)
}

// List retrieves all Albums across all pages using Client.ListAlbums.
func (a *Album) List(ctx context.Context, client Client) ([]*Album, error) {
	var allAlbums []*Album
	for page := 1; ; page++ {
		albums, err := client.ListAlbums(ctx, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list albums on page %d: %w", page, err)
		}
		if len(albums) == 0 {
			break
		}
		allAlbums = append(allAlbums, albums...)
	}
	return allAlbums, nil
}

// Update updates an Album by ID using Client.UpdateAlbum.
func (a *Album) Update(ctx context.Context, client Client, id ID) (*Album, error) {
	return client.UpdateAlbum(ctx, id, a)
}

// Patch partially updates an Album by ID using Client.PatchAlbum.
func (a *Album) Patch(ctx context.Context, client Client, id ID) (*Album, error) {
	return client.PatchAlbum(ctx, id, a)
}

// Delete removes an Album by ID using Client.DeleteAlbum.
func (a *Album) Delete(ctx context.Context, client Client, id ID) error {
	return client.DeleteAlbum(ctx, id)
}

// ListChildren retrieves all child Albums for the given ID across all pages using Client.ListAlbumChildren.
func (a *Album) ListChildren(ctx context.Context, client Client, id ID) ([]*Album, error) {
	var allChildren []*Album
	for page := 1; ; page++ {
		children, err := client.ListAlbumChildren(ctx, id, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list child albums for ID %s on page %d: %w", id, page, err)
		}
		if len(children) == 0 {
			break
		}
		allChildren = append(allChildren, children...)
	}
	return allChildren, nil
}

// UploadImage uploads an image for an Album using Client.UploadAlbumImage.
func (a *Album) UploadImage(ctx context.Context, client Client, id ID, file []byte) (*Album, error) {
	return client.UploadAlbumImage(ctx, id, file)
}

// GetParent retrieves the parent Album using Client.GetAlbumParent.
func (a *Album) GetParent(ctx context.Context, client Client, id ID) (*Album, error) {
	return client.GetAlbumParent(ctx, id)
}

// ListPhotos retrieves all Photos for the given ID across all pages using Client.ListAlbumPhotos.
func (a *Album) ListPhotos(ctx context.Context, client Client, id ID) ([]*Photo, error) {
	var allPhotos []*Photo
	for page := 1; ; page++ {
		photos, err := client.ListAlbumPhotos(ctx, id, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list photos for ID %s on page %d: %w", id, page, err)
		}
		if len(photos) == 0 {
			break
		}
		allPhotos = append(allPhotos, photos...)
	}
	return allPhotos, nil
}

// ChoiceListInterface defines methods for interacting with ChoiceList resources.
type ChoiceListInterface interface {
	Create(ctx context.Context, client Client) (*ChoiceList, error)
	Get(ctx context.Context, client Client, id ID) (*ChoiceList, error)
	List(ctx context.Context, client Client) ([]*ChoiceList, error)
	Update(ctx context.Context, client Client, id ID) (*ChoiceList, error)
	Patch(ctx context.Context, client Client, id ID) (*ChoiceList, error)
	Delete(ctx context.Context, client Client, id ID) error
	Validate(ctx context.Context, client Client) error
	IRI() string
}

// ChoiceList represents a choice list in Koillection, combining read and write fields.
type ChoiceList struct {
	Context   *Context   `json:"@context,omitempty" access:"rw"`  // JSON-LD only
	_ID       ID         `json:"@id,omitempty" access:"ro"`       // JSON-LD, read-only
	ID        ID         `json:"id,omitempty" access:"ro"`        // read-only
	Type      string     `json:"@type,omitempty" access:"rw"`     // JSON-LD only
	Name      string     `json:"name" access:"rw"`                // read and write
	Choices   []string   `json:"choices" access:"rw"`             // read and write
	Owner     *string    `json:"owner,omitempty" access:"ro"`     // read-only, IRI
	CreatedAt time.Time  `json:"createdAt" access:"ro"`           // read-only
	UpdatedAt *time.Time `json:"updatedAt,omitempty" access:"ro"` // read-only
}

// Validate checks the ChoiceList's fields for validity, using ctx for cancellation.
func (cl *ChoiceList) Validate(ctx context.Context, client Client) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("validation cancelled: %w", err)
	}

	if cl.Name == "" {
		return fmt.Errorf("name must not be empty")
	}

	if len(cl.Choices) == 0 {
		return fmt.Errorf("choices must contain at least one item")
	}

	if cl.ID == "" && cl._ID == "" {
		if cl._ID != "" {
			return fmt.Errorf("_ID must be empty for creation")
		}
		if cl.ID != "" {
			return fmt.Errorf("ID must be empty for creation")
		}
		if cl.Type != "" && cl.Type != "ChoiceList" {
			return fmt.Errorf("Type must be empty or 'ChoiceList' for creation: %s", cl.Type)
		}
	} else {
		if cl.ID == "" {
			return fmt.Errorf("ID must not be empty for update")
		}
	}

	return nil
}

// IRI returns the JSON-LD IRI for the ChoiceList.
func (cl *ChoiceList) IRI() string {
	if cl.ID != "" {
		return fmt.Sprintf("/api/choice_lists/%s", cl.ID)
	}
	if cl._ID != "" {
		return fmt.Sprintf("/api/choice_lists/%s", cl._ID)
	}
	return ""
}

// Create calls Client.CreateChoiceList to create a new ChoiceList.
func (cl *ChoiceList) Create(ctx context.Context, client Client) (*ChoiceList, error) {
	return client.CreateChoiceList(ctx, cl)
}

// Get retrieves a ChoiceList by ID using Client.GetChoiceList.
func (cl *ChoiceList) Get(ctx context.Context, client Client, id ID) (*ChoiceList, error) {
	return client.GetChoiceList(ctx, id)
}

// List retrieves all ChoiceLists across all pages using Client.ListChoiceLists.
func (cl *ChoiceList) List(ctx context.Context, client Client) ([]*ChoiceList, error) {
	var allChoiceLists []*ChoiceList
	for page := 1; ; page++ {
		choiceLists, err := client.ListChoiceLists(ctx, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list choice lists on page %d: %w", page, err)
		}
		if len(choiceLists) == 0 {
			break
		}
		allChoiceLists = append(allChoiceLists, choiceLists...)
	}
	return allChoiceLists, nil
}

// Update updates a ChoiceList by ID using Client.UpdateChoiceList.
func (cl *ChoiceList) Update(ctx context.Context, client Client, id ID) (*ChoiceList, error) {
	return client.UpdateChoiceList(ctx, id, cl)
}

// Patch partially updates a ChoiceList by ID using Client.PatchChoiceList.
func (cl *ChoiceList) Patch(ctx context.Context, client Client, id ID) (*ChoiceList, error) {
	return client.PatchChoiceList(ctx, id, cl)
}

// Delete removes a ChoiceList by ID using Client.DeleteChoiceList.
func (cl *ChoiceList) Delete(ctx context.Context, client Client, id ID) error {
	return client.DeleteChoiceList(ctx, id)
}

// CollectionInterface defines methods for interacting with Collection resources.
type CollectionInterface interface {
	Create(ctx context.Context, client Client) (*Collection, error)
	Get(ctx context.Context, client Client, id ID) (*Collection, error)
	List(ctx context.Context, client Client) ([]*Collection, error)
	Update(ctx context.Context, client Client, id ID) (*Collection, error)
	Patch(ctx context.Context, client Client, id ID) (*Collection, error)
	Delete(ctx context.Context, client Client, id ID) error
	ListChildren(ctx context.Context, client Client, id ID) ([]*Collection, error)
	UploadImage(ctx context.Context, client Client, id ID, file []byte) (*Collection, error)
	GetParent(ctx context.Context, client Client, id ID) (*Collection, error)
	ListCollectionItems(ctx context.Context, client Client, id ID) ([]*Item, error)
	ListCollectionData(ctx context.Context, client Client, id ID) ([]*Datum, error)
	GetDefaultTemplate(ctx context.Context, client Client, id ID) (*Template, error)
	Validate(ctx context.Context, client Client) error
	IRI() string
}

// Collection represents a collection in Koillection, combining read and write fields.
type Collection struct {
	Context              *Context   `json:"@context,omitempty" access:"rw"`             // JSON-LD only
	_ID                  ID         `json:"@id,omitempty" access:"ro"`                  // JSON-LD, read-only
	ID                   ID         `json:"id,omitempty" access:"ro"`                   // read-only
	Type                 string     `json:"@type,omitempty" access:"rw"`                // JSON-LD only
	Title                string     `json:"title" access:"rw"`                          // read and write
	Parent               *string    `json:"parent,omitempty" access:"rw"`               // read and write, IRI
	Owner                *string    `json:"owner,omitempty" access:"ro"`                // read-only, IRI
	Color                string     `json:"color,omitempty" access:"ro"`                // read-only
	Image                *string    `json:"image,omitempty" access:"ro"`                // read-only
	SeenCounter          int        `json:"seenCounter,omitempty" access:"ro"`          // read-only
	ItemsDefaultTemplate *string    `json:"itemsDefaultTemplate,omitempty" access:"rw"` // read and write, IRI
	Visibility           Visibility `json:"visibility,omitempty" access:"rw"`           // read and write
	ParentVisibility     *string    `json:"parentVisibility,omitempty" access:"ro"`     // read-only
	FinalVisibility      Visibility `json:"finalVisibility,omitempty" access:"ro"`      // read-only
	ScrapedFromURL       *string    `json:"scrapedFromUrl,omitempty" access:"ro"`       // read-only
	CreatedAt            time.Time  `json:"createdAt" access:"ro"`                      // read-only
	UpdatedAt            *time.Time `json:"updatedAt,omitempty" access:"ro"`            // read-only
	File                 *string    `json:"file,omitempty" access:"wo"`                 // write-only, binary data via multipart form
	DeleteImage          *bool      `json:"deleteImage,omitempty" access:"wo"`          // write-only
}

// Validate checks the Collection's fields for validity, using ctx for cancellation and client for optional IRI validation.
func (c *Collection) Validate(ctx context.Context, client Client) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("validation cancelled: %w", err)
	}

	if c.Title == "" {
		return fmt.Errorf("title must not be empty")
	}

	switch c.Visibility {
	case VisibilityPublic, VisibilityInternal, VisibilityPrivate, "":
	default:
		return fmt.Errorf("invalid visibility value: %s", c.Visibility)
	}

	if c.Parent != nil {
		if *c.Parent == "" {
			return fmt.Errorf("parent IRI must not be empty if set")
		}
		if !strings.HasPrefix(*c.Parent, "/api/collections/") {
			return fmt.Errorf("parent IRI must start with /api/collections/: %s", *c.Parent)
		}
		if client != nil {
			parts := strings.Split(*c.Parent, "/")
			if len(parts) < 4 {
				return fmt.Errorf("invalid parent IRI format: %s", *c.Parent)
			}
			parentID := ID(parts[3])
			_, err := client.GetCollection(ctx, parentID)
			if err != nil {
				return fmt.Errorf("invalid parent collection %s: %w", *c.Parent, err)
			}
		}
	}

	if c.ItemsDefaultTemplate != nil {
		if *c.ItemsDefaultTemplate == "" {
			return fmt.Errorf("itemsDefaultTemplate IRI must not be empty if set")
		}
		if !strings.HasPrefix(*c.ItemsDefaultTemplate, "/api/templates/") {
			return fmt.Errorf("itemsDefaultTemplate IRI must start with /api/templates/: %s", *c.ItemsDefaultTemplate)
		}
		if client != nil {
			parts := strings.Split(*c.ItemsDefaultTemplate, "/")
			if len(parts) < 4 {
				return fmt.Errorf("invalid itemsDefaultTemplate IRI format: %s", *c.ItemsDefaultTemplate)
			}
			templateID := ID(parts[3])
			_, err := client.GetTemplate(ctx, templateID)
			if err != nil {
				return fmt.Errorf("invalid itemsDefaultTemplate %s: %w", *c.ItemsDefaultTemplate, err)
			}
		}
	}

	if c.File != nil && *c.File == "" {
		return fmt.Errorf("file must not be empty if set")
	}

	if c.ID == "" && c._ID == "" {
		if c._ID != "" {
			return fmt.Errorf("_ID must be empty for creation")
		}
		if c.ID != "" {
			return fmt.Errorf("ID must be empty for creation")
		}
		if c.Type != "" && c.Type != "Collection" {
			return fmt.Errorf("Type must be empty or 'Collection' for creation: %s", c.Type)
		}
	} else {
		if c.ID == "" {
			return fmt.Errorf("ID must not be empty for update")
		}
	}

	return nil
}

// IRI returns the JSON-LD IRI for the Collection.
func (c *Collection) IRI() string {
	if c.ID != "" {
		return fmt.Sprintf("/api/collections/%s", c.ID)
	}
	if c._ID != "" {
		return fmt.Sprintf("/api/collections/%s", c._ID)
	}
	return ""
}

// Create calls Client.CreateCollection to create a new Collection.
func (c *Collection) Create(ctx context.Context, client Client) (*Collection, error) {
	return client.CreateCollection(ctx, c)
}

// Get retrieves a Collection by ID using Client.GetCollection.
func (c *Collection) Get(ctx context.Context, client Client, id ID) (*Collection, error) {
	return client.GetCollection(ctx, id)
}

// List retrieves all Collections across all pages using Client.ListCollections.
func (c *Collection) List(ctx context.Context, client Client) ([]*Collection, error) {
	var allCollections []*Collection
	for page := 1; ; page++ {
		collections, err := client.ListCollections(ctx, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list collections on page %d: %w", page, err)
		}
		if len(collections) == 0 {
			break
		}
		allCollections = append(allCollections, collections...)
	}
	return allCollections, nil
}

// Update updates a Collection by ID using Client.UpdateCollection.
func (c *Collection) Update(ctx context.Context, client Client, id ID) (*Collection, error) {
	return client.UpdateCollection(ctx, id, c)
}

// Patch partially updates a Collection by ID using Client.PatchCollection.
func (c *Collection) Patch(ctx context.Context, client Client, id ID) (*Collection, error) {
	return client.PatchCollection(ctx, id, c)
}

// Delete removes a Collection by ID using Client.DeleteCollection.
func (c *Collection) Delete(ctx context.Context, client Client, id ID) error {
	return client.DeleteCollection(ctx, id)
}

// ListChildren retrieves all child Collections for the given ID across all pages using Client.ListCollectionChildren.
func (c *Collection) ListChildren(ctx context.Context, client Client, id ID) ([]*Collection, error) {
	var allChildren []*Collection
	for page := 1; ; page++ {
		children, err := client.ListCollectionChildren(ctx, id, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list child collections for ID %s on page %d: %w", id, page, err)
		}
		if len(children) == 0 {
			break
		}
		allChildren = append(allChildren, children...)
	}
	return allChildren, nil
}

// UploadImage uploads an image for a Collection using Client.UploadCollectionImage.
func (c *Collection) UploadImage(ctx context.Context, client Client, id ID, file []byte) (*Collection, error) {
	return client.UploadCollectionImage(ctx, id, file)
}

// GetParent retrieves the parent Collection using Client.GetCollectionParent.
func (c *Collection) GetParent(ctx context.Context, client Client, id ID) (*Collection, error) {
	return client.GetCollectionParent(ctx, id)
}

// ListCollectionItems retrieves all Items for the given ID across all pages using Client.ListCollectionItems.
func (c *Collection) ListCollectionItems(ctx context.Context, client Client, id ID) ([]*Item, error) {
	var allItems []*Item
	for page := 1; ; page++ {
		items, err := client.ListCollectionItems(ctx, id, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list items for Collection ID %s on page %d: %w", id, page, err)
		}
		if len(items) == 0 {
			break
		}
		allItems = append(allItems, items...)
	}
	return allItems, nil
}

// ListCollectionData retrieves all Data for the given ID across all pages using Client.ListCollectionData.
func (c *Collection) ListCollectionData(ctx context.Context, client Client, id ID) ([]*Datum, error) {
	var allData []*Datum
	for page := 1; ; page++ {
		data, err := client.ListCollectionData(ctx, id, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list data for Collection ID %s on page %d: %w", id, page, err)
		}
		if len(data) == 0 {
			break
		}
		allData = append(allData, data...)
	}
	return allData, nil
}

// GetDefaultTemplate retrieves the default Template for the Collection using Client.GetCollectionDefaultTemplate.
func (c *Collection) GetDefaultTemplate(ctx context.Context, client Client, id ID) (*Template, error) {
	return client.GetCollectionDefaultTemplate(ctx, id)
}

// DatumInterface defines methods for interacting with Datum resources.
type DatumInterface interface {
	Create(ctx context.Context, client Client) (*Datum, error)
	Get(ctx context.Context, client Client, id ID) (*Datum, error)
	List(ctx context.Context, client Client) ([]*Datum, error)
	Update(ctx context.Context, client Client, id ID) (*Datum, error)
	Patch(ctx context.Context, client Client, id ID) (*Datum, error)
	Delete(ctx context.Context, client Client, id ID) error
	UploadFile(ctx context.Context, client Client, id ID, file []byte) (*Datum, error)
	UploadImage(ctx context.Context, client Client, id ID, image []byte) (*Datum, error)
	UploadVideo(ctx context.Context, client Client, id ID, video []byte) (*Datum, error)
	GetItem(ctx context.Context, client Client, id ID) (*Item, error)
	GetCollection(ctx context.Context, client Client, id ID) (*Collection, error)
	Validate(ctx context.Context, client Client) error
	IRI() string
}

// Datum represents a custom data field in Koillection, combining read and write fields.
type Datum struct {
	Context             *Context   `json:"@context,omitempty" access:"rw"`            // JSON-LD only
	_ID                 ID         `json:"@id,omitempty" access:"ro"`                 // JSON-LD, read-only
	ID                  ID         `json:"id,omitempty" access:"ro"`                  // read-only
	Type                string     `json:"@type,omitempty" access:"rw"`               // JSON-LD only
	Item                *string    `json:"item,omitempty" access:"rw"`                // read and write, IRI
	Collection          *string    `json:"collection,omitempty" access:"rw"`          // read and write, IRI
	DatumType           DatumType  `json:"type" access:"rw"`                          // read and write
	Label               string     `json:"label" access:"rw"`                         // read and write
	Value               *string    `json:"value,omitempty" access:"rw"`               // read and write
	Position            *int       `json:"position,omitempty" access:"rw"`            // read and write
	Currency            *string    `json:"currency,omitempty" access:"rw"`            // read and write
	Image               *string    `json:"image,omitempty" access:"ro"`               // read-only
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty" access:"ro"` // read-only
	ImageLargeThumbnail *string    `json:"imageLargeThumbnail,omitempty" access:"ro"` // read-only
	File                *string    `json:"file,omitempty" access:"ro"`                // read-only
	Video               *string    `json:"video,omitempty" access:"ro"`               // read-only
	OriginalFilename    *string    `json:"originalFilename,omitempty" access:"ro"`    // read-only
	ChoiceList          *string    `json:"choiceList,omitempty" access:"rw"`          // read and write, IRI
	Owner               *string    `json:"owner,omitempty" access:"ro"`               // read-only, IRI
	Visibility          Visibility `json:"visibility,omitempty" access:"rw"`          // read and write
	ParentVisibility    *string    `json:"parentVisibility,omitempty" access:"ro"`    // read-only
	FinalVisibility     Visibility `json:"finalVisibility,omitempty" access:"ro"`     // read-only
	CreatedAt           time.Time  `json:"createdAt" access:"ro"`                     // read-only
	UpdatedAt           *time.Time `json:"updatedAt,omitempty" access:"ro"`           // read-only
	FileImage           *string    `json:"fileImage,omitempty" access:"wo"`           // write-only, binary data via multipart form
	FileFile            *string    `json:"fileFile,omitempty" access:"wo"`            // write-only, binary data via multipart form
	FileVideo           *string    `json:"fileVideo,omitempty" access:"wo"`           // write-only, binary data via multipart form
}

// Validate checks the Datum's fields for validity, using ctx for cancellation and client for optional IRI validation.
func (d *Datum) Validate(ctx context.Context, client Client) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("validation cancelled: %w", err)
	}

	if d.DatumType == "" {
		return fmt.Errorf("datumType must not be empty")
	}
	if d.Label == "" {
		return fmt.Errorf("label must not be empty")
	}

	if d.Value != nil {
		if *d.Value == "" {
			return fmt.Errorf("value must not be empty if set")
		}
		if d.DatumType == DatumTypePrice {
			if _, err := strconv.ParseFloat(*d.Value, 64); err != nil {
				return fmt.Errorf("value must be a valid float for price type: %s", *d.Value)
			}
		}
	}

	if d.Currency != nil {
		if *d.Currency == "" {
			return fmt.Errorf("currency must not be empty if set")
		}
		if _, err := currency.ParseISO(*d.Currency); err != nil {
			return fmt.Errorf("currency must be a valid ISO 4217 code: %s", *d.Currency)
		}
	}

	if d.Item != nil {
		if *d.Item == "" {
			return fmt.Errorf("item IRI must not be empty if set")
		}
		if !strings.HasPrefix(*d.Item, "/api/items/") {
			return fmt.Errorf("item IRI must start with /api/items/: %s", *d.Item)
		}
		if client != nil {
			parts := strings.Split(*d.Item, "/")
			if len(parts) < 4 {
				return fmt.Errorf("invalid item IRI format: %s", *d.Item)
			}
			itemID := ID(parts[3])
			_, err := client.GetItem(ctx, itemID)
			if err != nil {
				return fmt.Errorf("invalid item %s: %w", *d.Item, err)
			}
		}
	}

	if d.Collection != nil {
		if *d.Collection == "" {
			return fmt.Errorf("collection IRI must not be empty if set")
		}
		if !strings.HasPrefix(*d.Collection, "/api/collections/") {
			return fmt.Errorf("collection IRI must start with /api/collections/: %s", *d.Collection)
		}
		if client != nil {
			parts := strings.Split(*d.Collection, "/")
			if len(parts) < 4 {
				return fmt.Errorf("invalid collection IRI format: %s", *d.Collection)
			}
			collectionID := ID(parts[3])
			_, err := client.GetCollection(ctx, collectionID)
			if err != nil {
				return fmt.Errorf("invalid collection %s: %w", *d.Collection, err)
			}
		}
	}

	if d.ChoiceList != nil {
		if *d.ChoiceList == "" {
			return fmt.Errorf("choiceList IRI must not be empty if set")
		}
		if !strings.HasPrefix(*d.ChoiceList, "/api/choice_lists/") {
			return fmt.Errorf("choiceList IRI must start with /api/choice_lists/: %s", *d.ChoiceList)
		}
		if client != nil {
			parts := strings.Split(*d.ChoiceList, "/")
			if len(parts) < 4 {
				return fmt.Errorf("invalid choiceList IRI format: %s", *d.ChoiceList)
			}
			choiceListID := ID(parts[3])
			_, err := client.GetChoiceList(ctx, choiceListID)
			if err != nil {
				return fmt.Errorf("invalid choiceList %s: %w", *d.ChoiceList, err)
			}
		}
	}

	switch d.Visibility {
	case VisibilityPublic, VisibilityInternal, VisibilityPrivate, "":
	default:
		return fmt.Errorf("invalid visibility value: %s", d.Visibility)
	}

	if d.FileImage != nil && *d.FileImage == "" {
		return fmt.Errorf("fileImage must not be empty if set")
	}
	if d.FileFile != nil && *d.FileFile == "" {
		return fmt.Errorf("fileFile must not be empty if set")
	}
	if d.FileVideo != nil && *d.FileVideo == "" {
		return fmt.Errorf("fileVideo must not be empty if set")
	}

	if d.ID == "" && d._ID == "" {
		if d._ID != "" {
			return fmt.Errorf("_ID must be empty for creation")
		}
		if d.ID != "" {
			return fmt.Errorf("ID must be empty for creation")
		}
		if d.Type != "" && d.Type != "Datum" {
			return fmt.Errorf("Type must be empty or 'Datum' for creation: %s", d.Type)
		}
	} else {
		if d.ID == "" {
			return fmt.Errorf("ID must not be empty for update")
		}
	}

	return nil
}

// IRI returns the JSON-LD IRI for the Datum.
func (d *Datum) IRI() string {
	if d.ID != "" {
		return fmt.Sprintf("/api/data/%s", d.ID)
	}
	if d._ID != "" {
		return fmt.Sprintf("/api/data/%s", d._ID)
	}
	return ""
}

// Create calls Client.CreateDatum to create a new Datum.
func (d *Datum) Create(ctx context.Context, client Client) (*Datum, error) {
	return client.CreateDatum(ctx, d)
}

// Get retrieves a Datum by ID using Client.GetDatum.
func (d *Datum) Get(ctx context.Context, client Client, id ID) (*Datum, error) {
	return client.GetDatum(ctx, id)
}

// List retrieves all Data across all pages using Client.ListData.
func (d *Datum) List(ctx context.Context, client Client) ([]*Datum, error) {
	var allData []*Datum
	for page := 1; ; page++ {
		data, err := client.ListData(ctx, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list data on page %d: %w", page, err)
		}
		if len(data) == 0 {
			break
		}
		allData = append(allData, data...)
	}
	return allData, nil
}

// Update updates a Datum by ID using Client.UpdateDatum.
func (d *Datum) Update(ctx context.Context, client Client, id ID) (*Datum, error) {
	return client.UpdateDatum(ctx, id, d)
}

// Patch partially updates a Datum by ID using Client.PatchDatum.
func (d *Datum) Patch(ctx context.Context, client Client, id ID) (*Datum, error) {
	return client.PatchDatum(ctx, id, d)
}

// Delete removes a Datum by ID using Client.DeleteDatum.
func (d *Datum) Delete(ctx context.Context, client Client, id ID) error {
	return client.DeleteDatum(ctx, id)
}

// UploadFile uploads a file for a Datum using Client.UploadDatumFile.
func (d *Datum) UploadFile(ctx context.Context, client Client, id ID, file []byte) (*Datum, error) {
	return client.UploadDatumFile(ctx, id, file)
}

// UploadImage uploads an image for a Datum using Client.UploadDatumImage.
func (d *Datum) UploadImage(ctx context.Context, client Client, id ID, image []byte) (*Datum, error) {
	return client.UploadDatumImage(ctx, id, image)
}

// UploadVideo uploads a video for a Datum using Client.UploadDatumVideo.
func (d *Datum) UploadVideo(ctx context.Context, client Client, id ID, video []byte) (*Datum, error) {
	return client.UploadDatumVideo(ctx, id, video)
}

// GetItem retrieves the Item associated with the Datum using Client.GetDatumItem.
func (d *Datum) GetItem(ctx context.Context, client Client, id ID) (*Item, error) {
	return client.GetDatumItem(ctx, id)
}

// GetCollection retrieves the Collection associated with the Datum using Client.GetDatumCollection.
func (d *Datum) GetCollection(ctx context.Context, client Client, id ID) (*Collection, error) {
	return client.GetDatumCollection(ctx, id)
}

// FieldInterface defines methods for interacting with Field resources.
type FieldInterface interface {
	Create(ctx context.Context, client Client) (*Field, error)
	Get(ctx context.Context, client Client, id ID) (*Field, error)
	List(ctx context.Context, client Client) ([]*Field, error)
	Update(ctx context.Context, client Client, id ID) (*Field, error)
	Patch(ctx context.Context, client Client, id ID) (*Field, error)
	Delete(ctx context.Context, client Client, id ID) error
	GetTemplate(ctx context.Context, client Client, id ID) (*Template, error)
	Validate(ctx context.Context, client Client) error
	IRI() string
}

// Field represents a template field in Koillection, combining read and write fields.
type Field struct {
	Context    *Context   `json:"@context,omitempty" access:"rw"`   // JSON-LD only
	_ID        ID         `json:"@id,omitempty" access:"ro"`        // JSON-LD, read-only
	ID         ID         `json:"id,omitempty" access:"ro"`         // read-only
	Type       string     `json:"@type,omitempty" access:"rw"`      // JSON-LD only
	Name       string     `json:"name" access:"rw"`                 // read and write
	Position   int        `json:"position" access:"rw"`             // read and write
	FieldType  FieldType  `json:"type" access:"rw"`                 // read and write
	ChoiceList *string    `json:"choiceList,omitempty" access:"rw"` // read and write, IRI
	Template   *string    `json:"template" access:"rw"`             // read and write, IRI
	Visibility Visibility `json:"visibility,omitempty" access:"rw"` // read and write
	Owner      *string    `json:"owner,omitempty" access:"ro"`      // read-only, IRI
}

// Validate checks the Field's fields for validity, using ctx for cancellation and client for optional IRI validation.
func (f *Field) Validate(ctx context.Context, client Client) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("validation cancelled: %w", err)
	}

	if f.Name == "" {
		return fmt.Errorf("name must not be empty")
	}

	if f.FieldType == "" {
		return fmt.Errorf("fieldType must not be empty")
	}

	if f.ChoiceList != nil {
		if *f.ChoiceList == "" {
			return fmt.Errorf("choiceList IRI must not be empty if set")
		}
		if !strings.HasPrefix(*f.ChoiceList, "/api/choice_lists/") {
			return fmt.Errorf("choiceList IRI must start with /api/choice_lists/: %s", *f.ChoiceList)
		}
		if client != nil {
			parts := strings.Split(*f.ChoiceList, "/")
			if len(parts) < 4 {
				return fmt.Errorf("invalid choiceList IRI format: %s", *f.ChoiceList)
			}
			choiceListID := ID(parts[3])
			_, err := client.GetChoiceList(ctx, choiceListID)
			if err != nil {
				return fmt.Errorf("invalid choiceList %s: %w", *f.ChoiceList, err)
			}
		}
	}

	if f.Template != nil {
		if *f.Template == "" {
			return fmt.Errorf("template IRI must not be empty if set")
		}
		if !strings.HasPrefix(*f.Template, "/api/templates/") {
			return fmt.Errorf("template IRI must start with /api/templates/: %s", *f.Template)
		}
		if client != nil {
			parts := strings.Split(*f.Template, "/")
			if len(parts) < 4 {
				return fmt.Errorf("invalid template IRI format: %s", *f.Template)
			}
			templateID := ID(parts[3])
			_, err := client.GetTemplate(ctx, templateID)
			if err != nil {
				return fmt.Errorf("invalid template %s: %w", *f.Template, err)
			}
		}
	}

	switch f.Visibility {
	case VisibilityPublic, VisibilityInternal, VisibilityPrivate, "":
	default:
		return fmt.Errorf("invalid visibility value: %s", f.Visibility)
	}

	if f.ID == "" && f._ID == "" {
		if f._ID != "" {
			return fmt.Errorf("_ID must be empty for creation")
		}
		if f.ID != "" {
			return fmt.Errorf("ID must be empty for creation")
		}
		if f.Type != "" && f.Type != "Field" {
			return fmt.Errorf("Type must be empty or 'Field' for creation: %s", f.Type)
		}
	} else {
		if f.ID == "" {
			return fmt.Errorf("ID must not be empty for update")
		}
	}

	return nil
}

// IRI returns the JSON-LD IRI for the Field.
func (f *Field) IRI() string {
	if f.ID != "" {
		return fmt.Sprintf("/api/fields/%s", f.ID)
	}
	if f._ID != "" {
		return fmt.Sprintf("/api/fields/%s", f._ID)
	}
	return ""
}

// Create calls Client.CreateField to create a new Field.
func (f *Field) Create(ctx context.Context, client Client) (*Field, error) {
	return client.CreateField(ctx, f)
}

// Get retrieves a Field by ID using Client.GetField.
func (f *Field) Get(ctx context.Context, client Client, id ID) (*Field, error) {
	return client.GetField(ctx, id)
}

// List retrieves all Fields across all pages using Client.ListFields.
func (f *Field) List(ctx context.Context, client Client) ([]*Field, error) {
	var allFields []*Field
	for page := 1; ; page++ {
		fields, err := client.ListFields(ctx, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list fields on page %d: %w", page, err)
		}
		if len(fields) == 0 {
			break
		}
		allFields = append(allFields, fields...)
	}
	return allFields, nil
}

// Update updates a Field by ID using Client.UpdateField.
func (f *Field) Update(ctx context.Context, client Client, id ID) (*Field, error) {
	return client.UpdateField(ctx, id, f)
}

// Patch partially updates a Field by ID using Client.PatchField.
func (f *Field) Patch(ctx context.Context, client Client, id ID) (*Field, error) {
	return client.PatchField(ctx, id, f)
}

// Delete removes a Field by ID using Client.DeleteField.
func (f *Field) Delete(ctx context.Context, client Client, id ID) error {
	return client.DeleteField(ctx, id)
}

// GetTemplate retrieves the Template associated with the Field using Client.GetFieldTemplate.
func (f *Field) GetTemplate(ctx context.Context, client Client, id ID) (*Template, error) {
	return client.GetFieldTemplate(ctx, id)
}

// InventoryInterface defines methods for interacting with Inventory resources.
type InventoryInterface interface {
	Get(ctx context.Context, client Client, id ID) (*Inventory, error)
	List(ctx context.Context, client Client) ([]*Inventory, error)
	Delete(ctx context.Context, client Client, id ID) error
	Validate(ctx context.Context, client Client) error
	IRI() string
}

// Inventory represents an inventory record in Koillection, combining read and write fields.
type Inventory struct {
	Context   *Context   `json:"@context,omitempty" access:"rw"`  // JSON-LD only
	_ID       ID         `json:"@id,omitempty" access:"ro"`       // JSON-LD, read-only
	ID        ID         `json:"id,omitempty" access:"ro"`        // read-only
	Type      string     `json:"@type,omitempty" access:"rw"`     // JSON-LD only
	Name      string     `json:"name" access:"rw"`                // read and write (assumed)
	Content   []string   `json:"content" access:"rw"`             // read and write (assumed)
	Owner     *string    `json:"owner,omitempty" access:"ro"`     // read-only, IRI
	CreatedAt time.Time  `json:"createdAt" access:"ro"`           // read-only
	UpdatedAt *time.Time `json:"updatedAt,omitempty" access:"ro"` // read-only
}

// Validate checks the Inventory's fields for validity, using ctx for cancellation.
func (i *Inventory) Validate(ctx context.Context, client Client) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("validation cancelled: %w", err)
	}

	if i.Name == "" {
		return fmt.Errorf("name must not be empty")
	}

	if len(i.Content) == 0 {
		return fmt.Errorf("content must contain at least one item")
	}

	if i.ID == "" && i._ID == "" {
		if i._ID != "" {
			return fmt.Errorf("_ID must be empty for creation")
		}
		if i.ID != "" {
			return fmt.Errorf("ID must be empty for creation")
		}
		if i.Type != "" && i.Type != "Inventory" {
			return fmt.Errorf("Type must be empty or 'Inventory' for creation: %s", i.Type)
		}
	} else {
		if i.ID == "" {
			return fmt.Errorf("ID must not be empty for update")
		}
	}

	return nil
}

// IRI returns the JSON-LD IRI for the Inventory.
func (i *Inventory) IRI() string {
	if i.ID != "" {
		return fmt.Sprintf("/api/inventories/%s", i.ID)
	}
	if i._ID != "" {
		return fmt.Sprintf("/api/inventories/%s", i._ID)
	}
	return ""
}

// Get retrieves an Inventory by ID using Client.GetInventory.
func (i *Inventory) Get(ctx context.Context, client Client, id ID) (*Inventory, error) {
	return client.GetInventory(ctx, id)
}

// List retrieves all Inventories across all pages using Client.ListInventories.
func (i *Inventory) List(ctx context.Context, client Client) ([]*Inventory, error) {
	var allInventories []*Inventory
	for page := 1; ; page++ {
		inventories, err := client.ListInventories(ctx, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list inventories on page %d: %w", page, err)
		}
		if len(inventories) == 0 {
			break
		}
		allInventories = append(allInventories, inventories...)
	}
	return allInventories, nil
}

// Delete removes an Inventory by ID using Client.DeleteInventory.
func (i *Inventory) Delete(ctx context.Context, client Client, id ID) error {
	return client.DeleteInventory(ctx, id)
}

// ItemInterface defines methods for interacting with Item resources.
type ItemInterface interface {
	Create(ctx context.Context, client Client) (*Item, error)
	Get(ctx context.Context, client Client, id ID) (*Item, error)
	List(ctx context.Context, client Client) ([]*Item, error)
	Update(ctx context.Context, client Client, id ID) (*Item, error)
	Patch(ctx context.Context, client Client, id ID) (*Item, error)
	Delete(ctx context.Context, client Client, id ID) error
	UploadImage(ctx context.Context, client Client, id ID, file []byte) (*Item, error)
	ListRelatedItems(ctx context.Context, client Client, id ID) ([]*Item, error)
	ListLoans(ctx context.Context, client Client, id ID) ([]*Loan, error)
	ListTags(ctx context.Context, client Client, id ID) ([]*Tag, error)
	ListItemData(ctx context.Context, client Client, id ID) ([]*Datum, error)
	GetCollection(ctx context.Context, client Client, id ID) (*Collection, error)
	Validate(ctx context.Context, client Client) error
	IRI() string
}

// Item represents an item within a collection, combining read and write fields.
type Item struct {
	Context             *Context   `json:"@context,omitempty" access:"rw"`            // JSON-LD only
	_ID                 ID         `json:"@id,omitempty" access:"ro"`                 // JSON-LD, read-only
	ID                  ID         `json:"id,omitempty" access:"ro"`                  // read-only
	Type                string     `json:"@type,omitempty" access:"rw"`               // JSON-LD only
	Name                string     `json:"name" access:"rw"`                          // read and write
	Quantity            int        `json:"quantity" access:"rw"`                      // read and write, must be >= 1
	Collection          *string    `json:"collection" access:"rw"`                    // read and write, IRI
	Owner               *string    `json:"owner,omitempty" access:"ro"`               // read-only, IRI
	Image               *string    `json:"image,omitempty" access:"ro"`               // read-only
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty" access:"ro"` // read-only
	ImageLargeThumbnail *string    `json:"imageLargeThumbnail,omitempty" access:"ro"` // read-only
	SeenCounter         int        `json:"seenCounter,omitempty" access:"ro"`         // read-only
	Visibility          Visibility `json:"visibility,omitempty" access:"rw"`          // read and write
	ParentVisibility    *string    `json:"parentVisibility,omitempty" access:"ro"`    // read-only
	FinalVisibility     Visibility `json:"finalVisibility,omitempty" access:"ro"`     // read-only
	ScrapedFromURL      *string    `json:"scrapedFromUrl,omitempty" access:"ro"`      // read-only
	CreatedAt           time.Time  `json:"createdAt" access:"ro"`                     // read-only
	UpdatedAt           *time.Time `json:"updatedAt,omitempty" access:"ro"`           // read-only
	Tags                []string   `json:"tags,omitempty" access:"wo"`                // write-only, IRI
	RelatedItems        []string   `json:"relatedItems,omitempty" access:"wo"`        // write-only, IRI
	File                *string    `json:"file,omitempty" access:"wo"`                // write-only, binary data via multipart form
}

// Validate checks the Item's fields for validity, using ctx for cancellation and client for optional IRI validation.
func (i *Item) Validate(ctx context.Context, client Client) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("validation cancelled: %w", err)
	}

	if i.Name == "" {
		return fmt.Errorf("name must not be empty")
	}

	if !ValidateQuantity(i.Quantity) {
		return fmt.Errorf("quantity must be at least 1: %d", i.Quantity)
	}

	if i.Collection != nil {
		if *i.Collection == "" {
			return fmt.Errorf("collection IRI must not be empty if set")
		}
		if !strings.HasPrefix(*i.Collection, "/api/collections/") {
			return fmt.Errorf("collection IRI must start with /api/collections/: %s", *i.Collection)
		}
		if client != nil {
			parts := strings.Split(*i.Collection, "/")
			if len(parts) < 4 {
				return fmt.Errorf("invalid collection IRI format: %s", *i.Collection)
			}
			collectionID := ID(parts[3])
			_, err := client.GetCollection(ctx, collectionID)
			if err != nil {
				return fmt.Errorf("invalid collection %s: %w", *i.Collection, err)
			}
		}
	}

	switch i.Visibility {
	case VisibilityPublic, VisibilityInternal, VisibilityPrivate, "":
	default:
		return fmt.Errorf("invalid visibility value: %s", i.Visibility)
	}

	if i.File != nil && *i.File == "" {
		return fmt.Errorf("file must not be empty if set")
	}

	for _, tag := range i.Tags {
		if tag == "" {
			return fmt.Errorf("tag IRI must not be empty")
		}
		if !strings.HasPrefix(tag, "/api/tags/") {
			return fmt.Errorf("tag IRI must start with /api/tags/: %s", tag)
		}
		if client != nil {
			parts := strings.Split(tag, "/")
			if len(parts) < 4 {
				return fmt.Errorf("invalid tag IRI format: %s", tag)
			}
			tagID := ID(parts[3])
			_, err := client.GetTag(ctx, tagID)
			if err != nil {
				return fmt.Errorf("invalid tag %s: %w", tag, err)
			}
		}
	}

	for _, relatedItem := range i.RelatedItems {
		if relatedItem == "" {
			return fmt.Errorf("relatedItem IRI must not be empty")
		}
		if !strings.HasPrefix(relatedItem, "/api/items/") {
			return fmt.Errorf("relatedItem IRI must start with /api/items/: %s", relatedItem)
		}
		if client != nil {
			parts := strings.Split(relatedItem, "/")
			if len(parts) < 4 {
				return fmt.Errorf("invalid relatedItem IRI format: %s", relatedItem)
			}
			itemID := ID(parts[3])
			_, err := client.GetItem(ctx, itemID)
			if err != nil {
				return fmt.Errorf("invalid related item %s: %w", relatedItem, err)
			}
		}
	}

	if i.ID == "" && i._ID == "" {
		if i._ID != "" {
			return fmt.Errorf("_ID must be empty for creation")
		}
		if i.ID != "" {
			return fmt.Errorf("ID must be empty for creation")
		}
		if i.Type != "" && i.Type != "Item" {
			return fmt.Errorf("Type must be empty or 'Item' for creation: %s", i.Type)
		}
	} else {
		if i.ID == "" {
			return fmt.Errorf("ID must not be empty for update")
		}
	}

	return nil
}

// IRI returns the JSON-LD IRI for the Item.
func (i *Item) IRI() string {
	if i.ID != "" {
		return fmt.Sprintf("/api/items/%s", i.ID)
	}
	if i._ID != "" {
		return fmt.Sprintf("/api/items/%s", i._ID)
	}
	return ""
}

// Create calls Client.CreateItem to create a new Item.
func (i *Item) Create(ctx context.Context, client Client) (*Item, error) {
	return client.CreateItem(ctx, i)
}

// Get retrieves an Item by ID using Client.GetItem.
func (i *Item) Get(ctx context.Context, client Client, id ID) (*Item, error) {
	return client.GetItem(ctx, id)
}

// List retrieves all Items across all pages using Client.ListItems.
func (i *Item) List(ctx context.Context, client Client) ([]*Item, error) {
	var allItems []*Item
	for page := 1; ; page++ {
		items, err := client.ListItems(ctx, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list items on page %d: %w", page, err)
		}
		if len(items) == 0 {
			break
		}
		allItems = append(allItems, items...)
	}
	return allItems, nil
}

// Update updates an Item by ID using Client.UpdateItem.
func (i *Item) Update(ctx context.Context, client Client, id ID) (*Item, error) {
	return client.UpdateItem(ctx, id, i)
}

// Patch partially updates an Item by ID using Client.PatchItem.
func (i *Item) Patch(ctx context.Context, client Client, id ID) (*Item, error) {
	return client.PatchItem(ctx, id, i)
}

// Delete removes an Item by ID using Client.DeleteItem.
func (i *Item) Delete(ctx context.Context, client Client, id ID) error {
	return client.DeleteItem(ctx, id)
}

// UploadImage uploads an image for an Item using Client.UploadItemImage.
func (i *Item) UploadImage(ctx context.Context, client Client, id ID, file []byte) (*Item, error) {
	return client.UploadItemImage(ctx, id, file)
}

// ListRelatedItems retrieves all related Items for the given ID across all pages using Client.ListItemRelatedItems.
func (i *Item) ListRelatedItems(ctx context.Context, client Client, id ID) ([]*Item, error) {
	var allRelatedItems []*Item
	for page := 1; ; page++ {
		relatedItems, err := client.ListItemRelatedItems(ctx, id, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list related items for ID %s on page %d: %w", id, page, err)
		}
		if len(relatedItems) == 0 {
			break
		}
		allRelatedItems = append(allRelatedItems, relatedItems...)
	}
	return allRelatedItems, nil
}

// ListLoans retrieves all Loans for the given ID across all pages using Client.ListItemLoans.
func (i *Item) ListLoans(ctx context.Context, client Client, id ID) ([]*Loan, error) {
	var allLoans []*Loan
	for page := 1; ; page++ {
		loans, err := client.ListItemLoans(ctx, id, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list loans for ID %s on page %d: %w", id, page, err)
		}
		if len(loans) == 0 {
			break
		}
		allLoans = append(allLoans, loans...)
	}
	return allLoans, nil
}

// ListTags retrieves all Tags for the given ID across all pages using Client.ListItem
