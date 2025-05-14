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
    UploadAlbumImageFromFile(ctx context.Context, id ID, filename string) (*Album, error)
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
    UploadCollectionImageFromFile(ctx context.Context, id ID, filename string) (*Collection, error)
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
    UploadDatumFileFromFile(ctx context.Context, id ID, filename string) (*Datum, error)
    UploadDatumImage(ctx context.Context, id ID, image []byte) (*Datum, error)
    UploadDatumImageFromFile(ctx context.Context, id ID, filename string) (*Datum, error)
    UploadDatumVideo(ctx context.Context, id ID, video []byte) (*Datum, error)
    UploadDatumVideoFromFile(ctx context.Context, id ID, filename string) (*Datum, error)
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
    UploadItemImageFromFile(ctx context.Context, id ID, filename string) (*Item, error)
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
    UploadPhotoImageFromFile(ctx context.Context, id ID, filename string) (*Photo, error)
    GetPhotoAlbum(ctx context.Context, id ID) (*Album, error)
    CreateTag(ctx context.Context, tag *Tag) (*Tag, error)
    GetTag(ctx context.Context, id ID) (*Tag, error)
    ListTags(ctx context.Context, page int) ([]*Tag, error)
    UpdateTag(ctx context.Context, id ID, tag *Tag) (*Tag, error)
    PatchTag(ctx context.Context, id ID, tag *Tag) (*Tag, error)
    DeleteTag(ctx context.Context, id ID) error
    UploadTagImage(ctx context.Context, id ID, file []byte) (*Tag, error)
    UploadTagImageFromFile(ctx context.Context, id ID, filename string) (*Tag, error)
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
    UploadWishImageFromFile(ctx context.Context, id ID, filename string) (*Wish, error)
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
    UploadWishlistImageFromFile(ctx context.Context, id ID, filename string) (*Wishlist, error)
    GetWishlistParent(ctx context.Context, id ID) (*Wishlist, error)
}

import "os"

// UploadAlbumImageFromFile reads the file from filename and calls UploadAlbumImage.
func (c *YourRealClientImplementation) UploadAlbumImageFromFile(ctx context.Context, id ID, filename string) (*Album, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return c.UploadAlbumImage(ctx, id, file)
}

// UploadCollectionImageFromFile reads the file from filename and calls UploadCollectionImage.
func (c *YourRealClientImplementation) UploadCollectionImageFromFile(ctx context.Context, id ID, filename string) (*Collection, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return c.UploadCollectionImage(ctx, id, file)
}

// UploadDatumFileFromFile reads the file from filename and calls UploadDatumFile.
func (c *YourRealClientImplementation) UploadDatumFileFromFile(ctx context.Context, id ID, filename string) (*Datum, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return c.UploadDatumFile(ctx, id, file)
}

// UploadDatumImageFromFile reads the file from filename and calls UploadDatumImage.
func (c *YourRealClientImplementation) UploadDatumImageFromFile(ctx context.Context, id ID, filename string) (*Datum, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return c.UploadDatumImage(ctx, id, file)
}

// UploadDatumVideoFromFile reads the file from filename and calls UploadDatumVideo.
func (c *YourRealClientImplementation) UploadDatumVideoFromFile(ctx context.Context, id ID, filename string) (*Datum, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return c.UploadDatumVideo(ctx, id, file)
}

// UploadItemImageFromFile reads the file from filename and calls UploadItemImage.
func (c *YourRealClientImplementation) UploadItemImageFromFile(ctx context.Context, id ID, filename string) (*Item, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return c.UploadItemImage(ctx, id, file)
}

// UploadPhotoImageFromFile reads the file from filename and calls UploadPhotoImage.
func (c *YourRealClientImplementation) UploadPhotoImageFromFile(ctx context.Context, id ID, filename string) (*Photo, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return c.UploadPhotoImage(ctx, id, file)
}

// UploadTagImageFromFile reads the file from filename and calls UploadTagImage.
func (c *YourRealClientImplementation) UploadTagImageFromFile(ctx context.Context, id ID, filename string) (*Tag, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return c.UploadTagImage(ctx, id, file)
}

// UploadWishImageFromFile reads the file from filename and calls UploadWishImage.
func (c *YourRealClientImplementation) UploadWishImageFromFile(ctx context.Context, id ID, filename string) (*Wish, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return c.UploadWishImage(ctx, id, file)
}

// UploadWishlistImageFromFile reads the file from filename and calls UploadWishlistImage.
func (c *YourRealClientImplementation) UploadWishlistImageFromFile(ctx context.Context, id ID, filename string) (*Wishlist, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return c.UploadWishlistImage(ctx, id, file)
}