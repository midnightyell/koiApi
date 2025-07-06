package koiApi

import (
	"context"
	"fmt"
)

type Client interface {
	CheckLogin(ctx context.Context, username, password string) (string, error)                       // HTTP POST /api/authentication_token
	GetMetrics(ctx context.Context) (*Metrics, error)                                                // HTTP GET /api/metrics
	CreateAlbum(ctx context.Context, album *Album) (*Album, error)                                   // HTTP POST /api/albums
	GetAlbum(ctx context.Context, id ID) (*Album, error)                                             // HTTP GET /api/albums/{id}
	ListAlbums(ctx context.Context, queryParams ...string) ([]*Album, error)                         // HTTP GET /api/albums
	UpdateAlbum(ctx context.Context, id ID, album *Album) (*Album, error)                            // HTTP PUT /api/albums/{id}
	PatchAlbum(ctx context.Context, id ID, album *Album) (*Album, error)                             // HTTP PATCH /api/albums/{id}
	DeleteAlbum(ctx context.Context, id ID) error                                                    // HTTP DELETE /api/albums/{id}
	ListAlbumChildren(ctx context.Context, id ID, queryParams ...string) ([]*Album, error)           // HTTP GET /api/albums/{id}/children
	UploadAlbumImage(ctx context.Context, id ID, file []byte) (*Album, error)                        // HTTP POST /api/albums/{id}/image
	GetAlbumParent(ctx context.Context, id ID) (*Album, error)                                       // HTTP GET /api/albums/{id}/parent
	ListAlbumPhotos(ctx context.Context, id ID, queryParams ...string) ([]*Photo, error)             // HTTP GET /api/albums/{id}/photos
	CreateChoiceList(ctx context.Context, choiceList *ChoiceList) (*ChoiceList, error)               // HTTP POST /api/choice_lists
	GetChoiceList(ctx context.Context, id ID) (*ChoiceList, error)                                   // HTTP GET /api/choice_lists/{id}
	ListChoiceLists(ctx context.Context, queryParams ...string) ([]*ChoiceList, error)               // HTTP GET /api/choice_lists
	UpdateChoiceList(ctx context.Context, id ID, choiceList *ChoiceList) (*ChoiceList, error)        // HTTP PUT /api/choice_lists/{id}
	PatchChoiceList(ctx context.Context, id ID, choiceList *ChoiceList) (*ChoiceList, error)         // HTTP PATCH /api/choice_lists/{id}
	DeleteChoiceList(ctx context.Context, id ID) error                                               // HTTP DELETE /api/choice_lists/{id}
	CreateCollection(ctx context.Context, collection *Collection) (*Collection, error)               // HTTP POST /api/collections
	GetCollection(ctx context.Context, id ID) (*Collection, error)                                   // HTTP GET /api/collections/{id}
	ListCollections(ctx context.Context, queryParams ...string) ([]*Collection, error)               // HTTP GET /api/collections
	UpdateCollection(ctx context.Context, id ID, collection *Collection) (*Collection, error)        // HTTP PUT /api/collections/{id}
	PatchCollection(ctx context.Context, id ID, collection *Collection) (*Collection, error)         // HTTP PATCH /api/collections/{id}
	DeleteCollection(ctx context.Context, id ID) error                                               // HTTP DELETE /api/collections/{id}
	ListCollectionChildren(ctx context.Context, id ID, queryParams ...string) ([]*Collection, error) // HTTP GET /api/collections/{id}/children
	UploadCollectionImage(ctx context.Context, id ID, file []byte) (*Collection, error)              // HTTP POST /api/collections/{id}/image
	GetCollectionParent(ctx context.Context, id ID) (*Collection, error)                             // HTTP GET /api/collections/{id}/parent
	ListCollectionItems(ctx context.Context, id ID, queryParams ...string) ([]*Item, error)          // HTTP GET /api/collections/{id}/items
	ListCollectionData(ctx context.Context, id ID, queryParams ...string) ([]*Datum, error)          // HTTP GET /api/collections/{id}/data
	GetCollectionDefaultTemplate(ctx context.Context, id ID) (*Template, error)                      // HTTP GET /api/collections/{id}/items_default_template
	CreateDatum(ctx context.Context, datum *Datum) (*Datum, error)                                   // HTTP POST /api/data
	GetDatum(ctx context.Context, id ID) (*Datum, error)                                             // HTTP GET /api/data/{id}
	ListData(ctx context.Context, queryParams ...string) ([]*Datum, error)                           // HTTP GET /api/data
	UpdateDatum(ctx context.Context, id ID, datum *Datum) (*Datum, error)                            // HTTP PUT /api/data/{id}
	PatchDatum(ctx context.Context, id ID, datum *Datum) (*Datum, error)                             // HTTP PATCH /api/data/{id}
	DeleteDatum(ctx context.Context, id ID) error                                                    // HTTP DELETE /api/data/{id}
	UploadDatumFile(ctx context.Context, id ID, file []byte) (*Datum, error)                         // HTTP POST /api/data/{id}/file
	UploadDatumImage(ctx context.Context, id ID, image []byte) (*Datum, error)                       // HTTP POST /api/data/{id}/image
	UploadDatumVideo(ctx context.Context, id ID, video []byte) (*Datum, error)                       // HTTP POST /api/data/{id}/video
	GetDatumItem(ctx context.Context, id ID) (*Item, error)                                          // HTTP GET /api/data/{id}/item
	GetDatumCollection(ctx context.Context, id ID) (*Collection, error)                              // HTTP GET /api/data/{id}/collection
	CreateField(ctx context.Context, field *Field) (*Field, error)                                   // HTTP POST /api/fields
	GetField(ctx context.Context, id ID) (*Field, error)                                             // HTTP GET /api/fields/{id}
	ListFields(ctx context.Context, queryParams ...string) ([]*Field, error)                         // HTTP GET /api/fields
	UpdateField(ctx context.Context, id ID, field *Field) (*Field, error)                            // HTTP PUT /api/fields/{id}
	PatchField(ctx context.Context, id ID, field *Field) (*Field, error)                             // HTTP PATCH /api/fields/{id}
	DeleteField(ctx context.Context, id ID) error                                                    // HTTP DELETE /api/fields/{id}
	GetFieldTemplate(ctx context.Context, id ID) (*Template, error)                                  // HTTP GET /api/fields/{id}/template
	ListTemplateFields(ctx context.Context, templateid ID, queryParams ...string) ([]*Field, error)  // HTTP GET /api/templates/{id}/fields
	ListInventories(ctx context.Context, queryParams ...string) ([]*Inventory, error)                // HTTP GET /api/inventories
	GetInventory(ctx context.Context, id ID) (*Inventory, error)                                     // HTTP GET /api/inventories/{id}
	DeleteInventory(ctx context.Context, id ID) error                                                // HTTP DELETE /api/inventories/{id}
	CreateItem(ctx context.Context, item *Item) (*Item, error)                                       // HTTP POST /api/items
	GetItem(ctx context.Context, id ID) (*Item, error)                                               // HTTP GET /api/items/{id}
	ListItems(ctx context.Context, queryParams ...string) ([]*Item, error)                           // HTTP GET /api/items
	UpdateItem(ctx context.Context, id ID, item *Item) (*Item, error)                                // HTTP PUT /api/items/{id}
	PatchItem(ctx context.Context, id ID, item *Item) (*Item, error)                                 // HTTP PATCH /api/items/{id}
	DeleteItem(ctx context.Context, id ID) error                                                     // HTTP DELETE /api/items/{id}
	UploadItemImage(ctx context.Context, id ID, file []byte) (*Item, error)                          // HTTP POST /api/items/{id}/image
	ListItemRelatedItems(ctx context.Context, id ID, queryParams ...string) ([]*Item, error)         // HTTP GET /api/items/{id}/related_items
	ListItemLoans(ctx context.Context, id ID, queryParams ...string) ([]*Loan, error)                // HTTP GET /api/items/{id}/loans
	ListItemTags(ctx context.Context, id ID, queryParams ...string) ([]*Tag, error)                  // HTTP GET /api/items/{id}/tags
	ListItemData(ctx context.Context, id ID, queryParams ...string) ([]*Datum, error)                // HTTP GET /api/items/{id}/data
	GetItemCollection(ctx context.Context, id ID) (*Collection, error)                               // HTTP GET /api/items/{id}/collection
	CreateLoan(ctx context.Context, loan *Loan) (*Loan, error)                                       // HTTP POST /api/loans
	GetLoan(ctx context.Context, id ID) (*Loan, error)                                               // HTTP GET /api/loans/{id}
	ListLoans(ctx context.Context, queryParams ...string) ([]*Loan, error)                           // HTTP GET /api/loans
	UpdateLoan(ctx context.Context, id ID, loan *Loan) (*Loan, error)                                // HTTP PUT /api/loans/{id}
	PatchLoan(ctx context.Context, id ID, loan *Loan) (*Loan, error)                                 // HTTP PATCH /api/loans/{id}
	DeleteLoan(ctx context.Context, id ID) error                                                     // HTTP DELETE /api/loans/{id}
	GetLoanItem(ctx context.Context, id ID) (*Item, error)                                           // HTTP GET /api/loans/{id}/item
	GetLog(ctx context.Context, id ID) (*Log, error)                                                 // HTTP GET /api/logs/{id}
	ListLogs(ctx context.Context, queryParams ...string) ([]*Log, error)                             // HTTP GET /api/logs
	CreatePhoto(ctx context.Context, photo *Photo) (*Photo, error)                                   // HTTP POST /api/photos
	GetPhoto(ctx context.Context, id ID) (*Photo, error)                                             // HTTP GET /api/photos/{id}
	ListPhotos(ctx context.Context, queryParams ...string) ([]*Photo, error)                         // HTTP GET /api/photos
	UpdatePhoto(ctx context.Context, id ID, photo *Photo) (*Photo, error)                            // HTTP PUT /api/photos/{id}
	PatchPhoto(ctx context.Context, id ID, photo *Photo) (*Photo, error)                             // HTTP PATCH /api/photos/{id}
	DeletePhoto(ctx context.Context, id ID) error                                                    // HTTP DELETE /api/photos/{id}
	UploadPhotoImage(ctx context.Context, id ID, file []byte) (*Photo, error)                        // HTTP POST /api/photos/{id}/image
	GetPhotoAlbum(ctx context.Context, id ID) (*Album, error)                                        // HTTP GET /api/photos/{id}/album
	CreateTag(ctx context.Context, tag *Tag) (*Tag, error)                                           // HTTP POST /api/tags
	GetTag(ctx context.Context, id ID) (*Tag, error)                                                 // HTTP GET /api/tags/{id}
	ListTags(ctx context.Context, queryParams ...string) ([]*Tag, error)                             // HTTP GET /api/tags
	UpdateTag(ctx context.Context, id ID, tag *Tag) (*Tag, error)                                    // HTTP PUT /api/tags/{id}
	PatchTag(ctx context.Context, id ID, tag *Tag) (*Tag, error)                                     // HTTP PATCH /api/tags/{id}
	DeleteTag(ctx context.Context, id ID) error                                                      // HTTP DELETE /api/tags/{id}
	UploadTagImage(ctx context.Context, id ID, file []byte) (*Tag, error)                            // HTTP POST /api/tags/{id}/image
	ListTagItems(ctx context.Context, id ID, queryParams ...string) ([]*Item, error)                 // HTTP GET /api/tags/{id}/items
	GetCategoryOfTag(ctx context.Context, id ID) (*TagCategory, error)                               // HTTP GET /api/tags/{id}/category
	CreateTagCategory(ctx context.Context, category *TagCategory) (*TagCategory, error)              // HTTP POST /api/tag_categories
	GetTagCategory(ctx context.Context, id ID) (*TagCategory, error)                                 // HTTP GET /api/tag_categories/{id}
	ListTagCategories(ctx context.Context, queryParams ...string) ([]*TagCategory, error)            // HTTP GET /api/tag_categories
	UpdateTagCategory(ctx context.Context, id ID, category *TagCategory) (*TagCategory, error)       // HTTP PUT /api/tag_categories/{id}
	PatchTagCategory(ctx context.Context, id ID, category *TagCategory) (*TagCategory, error)        // HTTP PATCH /api/tag_categories/{id}
	DeleteTagCategory(ctx context.Context, id ID) error                                              // HTTP DELETE /api/tag_categories/{id}
	ListTagCategoryTags(ctx context.Context, id ID, queryParams ...string) ([]*Tag, error)           // HTTP GET /api/tag_categories/{id}/tags
	CreateTemplate(ctx context.Context, template *Template) (*Template, error)                       // HTTP POST /api/templates
	GetTemplate(ctx context.Context, id ID) (*Template, error)                                       // HTTP GET /api/templates/{id}
	ListTemplates(ctx context.Context, queryParams ...string) ([]*Template, error)                   // HTTP GET /api/templates
	UpdateTemplate(ctx context.Context, id ID, template *Template) (*Template, error)                // HTTP PUT /api/templates/{id}
	PatchTemplate(ctx context.Context, id ID, template *Template) (*Template, error)                 // HTTP PATCH /api/templates/{id}
	DeleteTemplate(ctx context.Context, id ID) error                                                 // HTTP DELETE /api/templates/{id}
	GetUser(ctx context.Context, id ID) (*User, error)                                               // HTTP GET /api/users/{id}
	ListUsers(ctx context.Context, queryParams ...string) ([]*User, error)                           // HTTP GET /api/users
	CreateWish(ctx context.Context, wish *Wish) (*Wish, error)                                       // HTTP POST /api/wishes
	GetWish(ctx context.Context, id ID) (*Wish, error)                                               // HTTP GET /api/wishes/{id}
	ListWishes(ctx context.Context, queryParams ...string) ([]*Wish, error)                          // HTTP GET /api/wishes
	UpdateWish(ctx context.Context, id ID, wish *Wish) (*Wish, error)                                // HTTP PUT /api/wishes/{id}
	PatchWish(ctx context.Context, id ID, wish *Wish) (*Wish, error)                                 // HTTP PATCH /api/wishes/{id}
	DeleteWish(ctx context.Context, id ID) error                                                     // HTTP DELETE /api/wishes/{id}
	UploadWishImage(ctx context.Context, id ID, file []byte) (*Wish, error)                          // HTTP POST /api/wishes/{id}/image
	GetWishWishlist(ctx context.Context, id ID) (*Wishlist, error)                                   // HTTP GET /api/wishes/{id}/wishlist
	CreateWishlist(ctx context.Context, wishlist *Wishlist) (*Wishlist, error)                       // HTTP POST /api/wishlists
	GetWishlist(ctx context.Context, id ID) (*Wishlist, error)                                       // HTTP GET /api/wishlists/{id}
	ListWishlists(ctx context.Context, queryParams ...string) ([]*Wishlist, error)                   // HTTP GET /api/wishlists
	UpdateWishlist(ctx context.Context, id ID, wishlist *Wishlist) (*Wishlist, error)                // HTTP PUT /api/wishlists/{id}
	PatchWishlist(ctx context.Context, id ID, wishlist *Wishlist) (*Wishlist, error)                 // HTTP PATCH /api/wishlists/{id}
	DeleteWishlist(ctx context.Context, id ID) error                                                 // HTTP DELETE /api/wishlists/{id}
	ListWishlistWishes(ctx context.Context, id ID, queryParams ...string) ([]*Wish, error)           // HTTP GET /api/wishlists/{id}/wishes
	ListWishlistChildren(ctx context.Context, id ID, queryParams ...string) ([]*Wishlist, error)     // HTTP GET /api/wishlists/{id}/children
	UploadWishlistImage(ctx context.Context, id ID, file []byte) (*Wishlist, error)                  // HTTP POST /api/wishlists/{id}/image
	GetWishlistParent(ctx context.Context, id ID) (*Wishlist, error)                                 // HTTP GET /api/wishlists/{id}/parent
	GetResponse(ctx context.Context) string
	PrintError(ctx context.Context)
	DeleteAllData(ctx context.Context) error
}

// GetMetrics retrieves system or user-specific metrics.
func (c *httpClient) GetMetrics(ctx context.Context) (*Metrics, error) {
	var metrics Metrics
	if err := c.getResource(ctx, "/api/metrics", &metrics); err != nil {
		return nil, err
	}
	return &metrics, nil
}

// CreateAlbum creates a new album with schema validation.
func (c *httpClient) CreateAlbum(ctx context.Context, album *Album) (*Album, error) {
	if err := c.validateAlbum(album); err != nil {
		return nil, err
	}
	var result Album
	if err := c.postResource(ctx, "/api/albums", album, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAlbum retrieves an album by its ID.
func (c *httpClient) GetAlbum(ctx context.Context, id ID) (*Album, error) {
	var album Album
	if err := c.getResource(ctx, "/api/albums/"+string(id), &album); err != nil {
		return nil, err
	}
	return &album, nil
}

// ListAlbums retrieves a list of albums.
func (c *httpClient) ListAlbums(ctx context.Context, queryParams ...string) ([]*Album, error) {
	var albums []*Album
	if err := c.listResources(ctx, "/api/albums", &albums, queryParams...); err != nil {
		return nil, err
	}
	return albums, nil
}

// UpdateAlbum replaces an existing album.
func (c *httpClient) UpdateAlbum(ctx context.Context, id ID, album *Album) (*Album, error) {
	var result Album
	if err := c.putResource(ctx, "/api/albums/"+string(id), album, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchAlbum updates an existing album partially.
func (c *httpClient) PatchAlbum(ctx context.Context, id ID, album *Album) (*Album, error) {
	var result Album
	if err := c.patchResource(ctx, "/api/albums/"+string(id), album, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteAlbum deletes an album by its ID.
func (c *httpClient) DeleteAlbum(ctx context.Context, id ID) error {
	return c.deleteResource(ctx, "/api/albums/"+string(id))
}

// ListAlbumChildren retrieves the children of an album.
func (c *httpClient) ListAlbumChildren(ctx context.Context, id ID, queryParams ...string) ([]*Album, error) {
	var albums []*Album
	path := fmt.Sprintf("/api/albums/%s/children", id)
	if err := c.listResources(ctx, path, &albums, queryParams...); err != nil {
		return nil, err
	}
	return albums, nil
}

// UploadAlbumImage uploads an image for an album.
func (c *httpClient) UploadAlbumImage(ctx context.Context, id ID, file []byte) (*Album, error) {
	var album Album
	if err := c.uploadFile(ctx, "/api/albums/"+string(id)+"/image", file, "file", &album); err != nil {
		return nil, err
	}
	return &album, nil
}

// GetAlbumParent retrieves the parent album of an album.
func (c *httpClient) GetAlbumParent(ctx context.Context, id ID) (*Album, error) {
	var album Album
	if err := c.getResource(ctx, "/api/albums/"+string(id)+"/parent", &album); err != nil {
		return nil, err
	}
	return &album, nil
}

// ListAlbumPhotos retrieves the photos in an album.
func (c *httpClient) ListAlbumPhotos(ctx context.Context, id ID, queryParams ...string) ([]*Photo, error) {
	var photos []*Photo
	path := fmt.Sprintf("/api/albums/%s/photos", id)
	if err := c.listResources(ctx, path, &photos, queryParams...); err != nil {
		return nil, err
	}
	return photos, nil
}

// CreateChoiceList creates a new choice list with schema validation.
func (c *httpClient) CreateChoiceList(ctx context.Context, choiceList *ChoiceList) (*ChoiceList, error) {
	if err := c.validateChoiceList(choiceList); err != nil {
		return nil, err
	}
	var result ChoiceList
	if err := c.postResource(ctx, "/api/choice_lists", choiceList, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetChoiceList retrieves a choice list by its ID.
func (c *httpClient) GetChoiceList(ctx context.Context, id ID) (*ChoiceList, error) {
	var choiceList ChoiceList
	if err := c.getResource(ctx, "/api/choice_lists/"+string(id), &choiceList); err != nil {
		return nil, err
	}
	return &choiceList, nil
}

// ListChoiceLists retrieves a list of choice lists.
func (c *httpClient) ListChoiceLists(ctx context.Context, queryParams ...string) ([]*ChoiceList, error) {
	var choiceLists []*ChoiceList
	if err := c.listResources(ctx, "/api/choice_lists", &choiceLists, queryParams...); err != nil {
		return nil, err
	}
	return choiceLists, nil
}

// UpdateChoiceList replaces an existing choice list.
func (c *httpClient) UpdateChoiceList(ctx context.Context, id ID, choiceList *ChoiceList) (*ChoiceList, error) {
	var result ChoiceList
	if err := c.putResource(ctx, "/api/choice_lists/"+string(id), choiceList, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchChoiceList updates an existing choice list partially.
func (c *httpClient) PatchChoiceList(ctx context.Context, id ID, choiceList *ChoiceList) (*ChoiceList, error) {
	var result ChoiceList
	if err := c.patchResource(ctx, "/api/choice_lists/"+string(id), choiceList, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteChoiceList deletes a choice list by its ID.
func (c *httpClient) DeleteChoiceList(ctx context.Context, id ID) error {
	return c.deleteResource(ctx, "/api/choice_lists/"+string(id))
}

// CreateCollection creates a new collection with schema validation.
func (c *httpClient) CreateCollection(ctx context.Context, collection *Collection) (*Collection, error) {
	if err := c.validateCollection(collection); err != nil {
		return nil, err
	}
	var result Collection
	if err := c.postResource(ctx, "/api/collections", collection, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCollection retrieves a collection by its ID.
func (c *httpClient) GetCollection(ctx context.Context, id ID) (*Collection, error) {
	var collection Collection
	if err := c.getResource(ctx, "/api/collections/"+string(id), &collection); err != nil {
		return nil, err
	}
	return &collection, nil
}

// ListCollections retrieves a list of collections.
func (c *httpClient) ListCollections(ctx context.Context, queryParams ...string) ([]*Collection, error) {
	var collections []*Collection
	if err := c.listResources(ctx, "/api/collections", &collections, queryParams...); err != nil {
		return nil, err
	}
	return collections, nil
}

// UpdateCollection replaces an existing collection.
func (c *httpClient) UpdateCollection(ctx context.Context, id ID, collection *Collection) (*Collection, error) {
	var result Collection
	if err := c.putResource(ctx, "/api/collections/"+string(id), collection, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchCollection updates an existing collection partially.
func (c *httpClient) PatchCollection(ctx context.Context, id ID, collection *Collection) (*Collection, error) {
	var result Collection
	if err := c.patchResource(ctx, "/api/collections/"+string(id), collection, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteCollection deletes a collection by its ID.
func (c *httpClient) DeleteCollection(ctx context.Context, id ID) error {
	return c.deleteResource(ctx, "/api/collections/"+string(id))
}

// ListCollectionChildren retrieves the children of a collection.
func (c *httpClient) ListCollectionChildren(ctx context.Context, id ID, queryParams ...string) ([]*Collection, error) {
	var collections []*Collection
	path := fmt.Sprintf("/api/collections/%s/children", id)
	if err := c.listResources(ctx, path, &collections, queryParams...); err != nil {
		return nil, err
	}
	return collections, nil
}

// UploadCollectionImage uploads an image for a collection.
func (c *httpClient) UploadCollectionImage(ctx context.Context, id ID, file []byte) (*Collection, error) {
	var collection Collection
	if err := c.uploadFile(ctx, "/api/collections/"+string(id)+"/image", file, "file", &collection); err != nil {
		return nil, err
	}
	return &collection, nil
}

// GetCollectionParent retrieves the parent collection of a collection.
func (c *httpClient) GetCollectionParent(ctx context.Context, id ID) (*Collection, error) {
	var collection Collection
	if err := c.getResource(ctx, "/api/collections/"+string(id)+"/parent", &collection); err != nil {
		return nil, err
	}
	return &collection, nil
}

// ListCollectionItems retrieves the items in a collection.
func (c *httpClient) ListCollectionItems(ctx context.Context, id ID, queryParams ...string) ([]*Item, error) {
	var items []*Item
	path := fmt.Sprintf("/api/collections/%s/items", id)
	if err := c.listResources(ctx, path, &items, queryParams...); err != nil {
		return nil, err
	}
	return items, nil
}

// ListCollectionData retrieves the data in a collection.
func (c *httpClient) ListCollectionData(ctx context.Context, id ID, queryParams ...string) ([]*Datum, error) {
	var data []*Datum
	path := fmt.Sprintf("/api/collections/%s/data", id)
	if err := c.listResources(ctx, path, &data, queryParams...); err != nil {
		return nil, err
	}
	return data, nil
}

// GetCollectionDefaultTemplate retrieves the default template for items in a collection.
func (c *httpClient) GetCollectionDefaultTemplate(ctx context.Context, id ID) (*Template, error) {
	var template Template
	if err := c.getResource(ctx, "/api/collections/"+string(id)+"/items_default_template", &template); err != nil {
		return nil, err
	}
	return &template, nil
}

// CreateDatum creates a new datum with schema validation.
func (c *httpClient) CreateDatum(ctx context.Context, datum *Datum) (*Datum, error) {
	if err := c.validateDatum(datum); err != nil {
		return nil, err
	}
	var result Datum
	if err := c.postResource(ctx, "/api/data", datum, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDatum retrieves a datum by its ID.
func (c *httpClient) GetDatum(ctx context.Context, id ID) (*Datum, error) {
	var datum Datum
	if err := c.getResource(ctx, "/api/data/"+string(id), &datum); err != nil {
		return nil, err
	}
	return &datum, nil
}

// ListData retrieves a list of data.
func (c *httpClient) ListData(ctx context.Context, queryParams ...string) ([]*Datum, error) {
	var data []*Datum
	if err := c.listResources(ctx, "/api/data", &data, queryParams...); err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateDatum replaces an existing datum.
func (c *httpClient) UpdateDatum(ctx context.Context, id ID, datum *Datum) (*Datum, error) {
	var result Datum
	if err := c.putResource(ctx, "/api/data/"+string(id), datum, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchDatum updates an existing datum partially.
func (c *httpClient) PatchDatum(ctx context.Context, id ID, datum *Datum) (*Datum, error) {
	var result Datum
	if err := c.patchResource(ctx, "/api/data/"+string(id), datum, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteDatum deletes a datum by its ID.
func (c *httpClient) DeleteDatum(ctx context.Context, id ID) error {
	return c.deleteResource(ctx, "/api/data/"+string(id))
}

// UploadDatumFile uploads a file for a datum.
func (c *httpClient) UploadDatumFile(ctx context.Context, id ID, file []byte) (*Datum, error) {
	var datum Datum
	if err := c.uploadFile(ctx, "/api/data/"+string(id)+"/file", file, "fileFile", &datum); err != nil {
		return nil, err
	}
	return &datum, nil
}

// UploadDatumImage uploads an image for a datum.
func (c *httpClient) UploadDatumImage(ctx context.Context, id ID, image []byte) (*Datum, error) {
	var datum Datum
	if err := c.uploadFile(ctx, "/api/data/"+string(id)+"/image", image, "fileImage", &datum); err != nil {
		return nil, err
	}
	return &datum, nil
}

// UploadDatumVideo uploads a video for a datum.
func (c *httpClient) UploadDatumVideo(ctx context.Context, id ID, video []byte) (*Datum, error) {
	var datum Datum
	if err := c.uploadFile(ctx, "/api/data/"+string(id)+"/video", video, "fileVideo", &datum); err != nil {
		return nil, err
	}
	return &datum, nil
}

// GetDatumItem retrieves the item associated with a datum.
func (c *httpClient) GetDatumItem(ctx context.Context, id ID) (*Item, error) {
	var item Item
	if err := c.getResource(ctx, "/api/data/"+string(id)+"/item", &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// GetDatumCollection retrieves the collection associated with a datum.
func (c *httpClient) GetDatumCollection(ctx context.Context, id ID) (*Collection, error) {
	var collection Collection
	if err := c.getResource(ctx, "/api/data/"+string(id)+"/collection", &collection); err != nil {
		return nil, err
	}
	return &collection, nil
}

// CreateField creates a new field with schema validation.
func (c *httpClient) CreateField(ctx context.Context, field *Field) (*Field, error) {
	if err := c.validateField(field); err != nil {
		return nil, err
	}
	var result Field
	if err := c.postResource(ctx, "/api/fields", field, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetField retrieves a field by its ID.
func (c *httpClient) GetField(ctx context.Context, id ID) (*Field, error) {
	var field Field
	if err := c.getResource(ctx, "/api/fields/"+string(id), &field); err != nil {
		return nil, err
	}
	return &field, nil
}

// ListFields retrieves a list of fields.
func (c *httpClient) ListFields(ctx context.Context, queryParams ...string) ([]*Field, error) {
	var fields []*Field
	if err := c.listResources(ctx, "/api/fields", &fields, queryParams...); err != nil {
		return nil, err
	}
	return fields, nil
}

// UpdateField replaces an existing field.
func (c *httpClient) UpdateField(ctx context.Context, id ID, field *Field) (*Field, error) {
	var result Field
	if err := c.putResource(ctx, "/api/fields/"+string(id), field, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchField updates an existing field partially.
func (c *httpClient) PatchField(ctx context.Context, id ID, field *Field) (*Field, error) {
	var result Field
	if err := c.patchResource(ctx, "/api/fields/"+string(id), field, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteField deletes a field by its ID.
func (c *httpClient) DeleteField(ctx context.Context, id ID) error {
	return c.deleteResource(ctx, "/api/fields/"+string(id))
}

// GetFieldTemplate retrieves the template associated with a field.
func (c *httpClient) GetFieldTemplate(ctx context.Context, id ID) (*Template, error) {
	var template Template
	if err := c.getResource(ctx, "/api/fields/"+string(id)+"/template", &template); err != nil {
		return nil, err
	}
	return &template, nil
}

// ListTemplateFields retrieves the fields of a template.
func (c *httpClient) ListTemplateFields(ctx context.Context, templateid ID, queryParams ...string) ([]*Field, error) {
	var fields []*Field
	path := fmt.Sprintf("/api/templates/%s/fields", templateid)
	if err := c.listResources(ctx, path, &fields, queryParams...); err != nil {
		return nil, err
	}
	return fields, nil
}

// ListInventories retrieves a list of inventories.
func (c *httpClient) ListInventories(ctx context.Context, queryParams ...string) ([]*Inventory, error) {
	var inventories []*Inventory
	if err := c.listResources(ctx, "/api/inventories", &inventories, queryParams...); err != nil {
		return nil, err
	}
	return inventories, nil
}

// GetInventory retrieves an inventory by its ID.
func (c *httpClient) GetInventory(ctx context.Context, id ID) (*Inventory, error) {
	var inventory Inventory
	if err := c.getResource(ctx, "/api/inventories/"+string(id), &inventory); err != nil {
		return nil, err
	}
	return &inventory, nil
}

// DeleteInventory deletes an inventory by its ID.
func (c *httpClient) DeleteInventory(ctx context.Context, id ID) error {
	return c.deleteResource(ctx, "/api/inventories/"+string(id))
}

// CreateItem creates a new item.
func (c *httpClient) CreateItem(ctx context.Context, item *Item) (*Item, error) {
	if err := c.validateItem(item); err != nil {
		return nil, err
	}
	var result Item
	if err := c.postResource(ctx, "/api/items", item, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetItem retrieves an item by its ID.
func (c *httpClient) GetItem(ctx context.Context, id ID) (*Item, error) {
	var item Item
	if err := c.getResource(ctx, "/api/items/"+string(id), &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// ListItems retrieves a list of items.
func (c *httpClient) ListItems(ctx context.Context, queryParams ...string) ([]*Item, error) {
	var items []*Item
	if err := c.listResources(ctx, "/api/items", &items, queryParams...); err != nil {
		return nil, err
	}
	return items, nil
}

// UpdateItem replaces an existing item.
func (c *httpClient) UpdateItem(ctx context.Context, id ID, item *Item) (*Item, error) {
	var result Item
	if err := c.putResource(ctx, "/api/items/"+string(id), item, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchItem updates an existing item partially.
func (c *httpClient) PatchItem(ctx context.Context, id ID, item *Item) (*Item, error) {
	var result Item
	if err := c.patchResource(ctx, "/api/items/"+string(id), item, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteItem deletes an item by its ID.
func (c *httpClient) DeleteItem(ctx context.Context, id ID) error {
	return c.deleteResource(ctx, "/api/items/"+string(id))
}

// UploadItemImage uploads an image for an item.
func (c *httpClient) UploadItemImage(ctx context.Context, id ID, file []byte) (*Item, error) {
	var item Item
	if err := c.uploadFile(ctx, "/api/items/"+string(id)+"/image", file, "file", &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// ListItemRelatedItems retrieves the related items of an item.
func (c *httpClient) ListItemRelatedItems(ctx context.Context, id ID, queryParams ...string) ([]*Item, error) {
	var items []*Item
	path := fmt.Sprintf("/api/items/%s/related_items", id)
	if err := c.listResources(ctx, path, &items, queryParams...); err != nil {
		return nil, err
	}
	return items, nil
}

// ListItemLoans retrieves the loans of an item.
func (c *httpClient) ListItemLoans(ctx context.Context, id ID, queryParams ...string) ([]*Loan, error) {
	var loans []*Loan
	path := fmt.Sprintf("/api/items/%s/loans", id)
	if err := c.listResources(ctx, path, &loans, queryParams...); err != nil {
		return nil, err
	}
	return loans, nil
}

// ListItemTags retrieves the tags of an item.
func (c *httpClient) ListItemTags(ctx context.Context, id ID, queryParams ...string) ([]*Tag, error) {
	var tags []*Tag
	path := fmt.Sprintf("/api/items/%s/tags", id)
	if err := c.listResources(ctx, path, &tags, queryParams...); err != nil {
		return nil, err
	}
	return tags, nil
}

// ListItemData retrieves the data of an item.
func (c *httpClient) ListItemData(ctx context.Context, id ID, queryParams ...string) ([]*Datum, error) {
	var data []*Datum
	path := fmt.Sprintf("/api/items/%s/data", id)
	if err := c.listResources(ctx, path, &data, queryParams...); err != nil {
		return nil, err
	}
	return data, nil
}

// GetItemCollection retrieves the collection associated with an item.
func (c *httpClient) GetItemCollection(ctx context.Context, id ID) (*Collection, error) {
	var collection Collection
	if err := c.getResource(ctx, "/api/items/"+string(id)+"/collection", &collection); err != nil {
		return nil, err
	}
	return &collection, nil
}

// / CreateLoan creates a new loan with schema validation.
func (c *httpClient) CreateLoan(ctx context.Context, loan *Loan) (*Loan, error) {
	if err := c.validateLoan(loan); err != nil {
		return nil, err
	}
	var result Loan
	if err := c.postResource(ctx, "/api/loans", loan, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetLoan retrieves a loan by its ID.
func (c *httpClient) GetLoan(ctx context.Context, id ID) (*Loan, error) {
	var loan Loan
	if err := c.getResource(ctx, "/api/loans/"+string(id), &loan); err != nil {
		return nil, err
	}
	return &loan, nil
}

// ListLoans retrieves a list of loans.
func (c *httpClient) ListLoans(ctx context.Context, queryParams ...string) ([]*Loan, error) {
	var loans []*Loan
	if err := c.listResources(ctx, "/api/loans", &loans, queryParams...); err != nil {
		return nil, err
	}
	return loans, nil
}

// UpdateLoan replaces(an existing loan.
func (c *httpClient) UpdateLoan(ctx context.Context, id ID, loan *Loan) (*Loan, error) {
	var result Loan
	if err := c.putResource(ctx, "/api/loans/"+string(id), loan, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchLoan updates an existing loan partially.
func (c *httpClient) PatchLoan(ctx context.Context, id ID, loan *Loan) (*Loan, error) {
	var result Loan
	if err := c.patchResource(ctx, "/api/loans/"+string(id), loan, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteLoan deletes a loan by its ID.
func (c *httpClient) DeleteLoan(ctx context.Context, id ID) error {
	return c.deleteResource(ctx, "/api/loans/"+string(id))
}

// GetLoanItem retrieves the item associated with a loan.
func (c *httpClient) GetLoanItem(ctx context.Context, id ID) (*Item, error) {
	var item Item
	if err := c.getResource(ctx, "/api/loans/"+string(id)+"/item", &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// GetLog retrieves a log by its ID.
func (c *httpClient) GetLog(ctx context.Context, id ID) (*Log, error) {
	var log Log
	if err := c.getResource(ctx, "/api/logs/"+string(id), &log); err != nil {
		return nil, err
	}
	return &log, nil
}

// ListLogs retrieves a list of logs.
func (c *httpClient) ListLogs(ctx context.Context, queryParams ...string) ([]*Log, error) {
	var logs []*Log
	if err := c.listResources(ctx, "/api/logs", &logs, queryParams...); err != nil {
		return nil, err
	}
	return logs, nil
}

// CreatePhoto creates a new photo with schema validation.
func (c *httpClient) CreatePhoto(ctx context.Context, photo *Photo) (*Photo, error) {
	if err := c.validatePhoto(photo); err != nil {
		return nil, err
	}
	var result Photo
	if err := c.postResource(ctx, "/api/photos", photo, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPhoto retrieves a photo by its ID.
func (c *httpClient) GetPhoto(ctx context.Context, id ID) (*Photo, error) {
	var photo Photo
	if err := c.getResource(ctx, "/api/photos/"+string(id), &photo); err != nil {
		return nil, err
	}
	return &photo, nil
}

// ListPhotos retrieves a list of photos.
func (c *httpClient) ListPhotos(ctx context.Context, queryParams ...string) ([]*Photo, error) {
	var photos []*Photo
	if err := c.listResources(ctx, "/api/photos", &photos, queryParams...); err != nil {
		return nil, err
	}
	return photos, nil
}

// UpdatePhoto replaces an existing photo.
func (c *httpClient) UpdatePhoto(ctx context.Context, id ID, photo *Photo) (*Photo, error) {
	var result Photo
	if err := c.putResource(ctx, "/api/photos/"+string(id), photo, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchPhoto updates an existing photo partially.
func (c *httpClient) PatchPhoto(ctx context.Context, id ID, photo *Photo) (*Photo, error) {
	var result Photo
	if err := c.patchResource(ctx, "/api/photos/"+string(id), photo, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeletePhoto deletes a photo by its ID.
func (c *httpClient) DeletePhoto(ctx context.Context, id ID) error {
	return c.deleteResource(ctx, "/api/photos/"+string(id))
}

// UploadPhotoImage uploads an image for a photo.
func (c *httpClient) UploadPhotoImage(ctx context.Context, id ID, file []byte) (*Photo, error) {
	var photo Photo
	if err := c.uploadFile(ctx, "/api/photos/"+string(id)+"/image", file, "file", &photo); err != nil {
		return nil, err
	}
	return &photo, nil
}

// GetPhotoAlbum retrieves the album associated with a photo.
func (c *httpClient) GetPhotoAlbum(ctx context.Context, id ID) (*Album, error) {
	var album Album
	if err := c.getResource(ctx, "/api/photos/"+string(id)+"/album", &album); err != nil {
		return nil, err
	}
	return &album, nil
}

// CreateTag creates a new tag with schema validation.
func (c *httpClient) CreateTag(ctx context.Context, tag *Tag) (*Tag, error) {
	if err := c.validateTag(tag); err != nil {
		return nil, err
	}
	var result Tag
	if err := c.postResource(ctx, "/api/tags", tag, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTag retrieves a tag by its ID.
func (c *httpClient) GetTag(ctx context.Context, id ID) (*Tag, error) {
	var tag Tag
	if err := c.getResource(ctx, "/api/tags/"+string(id), &tag); err != nil {
		return nil, err
	}
	return &tag, nil
}

// ListTags retrieves a list of tags.
func (c *httpClient) ListTags(ctx context.Context, queryParams ...string) ([]*Tag, error) {
	var tags []*Tag
	if err := c.listResources(ctx, "/api/tags", &tags, queryParams...); err != nil {
		return nil, err
	}
	return tags, nil
}

// UpdateTag replaces an existing tag.
func (c *httpClient) UpdateTag(ctx context.Context, id ID, tag *Tag) (*Tag, error) {
	var result Tag
	if err := c.putResource(ctx, "/api/tags/"+string(id), tag, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchTag updates an existing tag partially.
func (c *httpClient) PatchTag(ctx context.Context, id ID, tag *Tag) (*Tag, error) {
	var result Tag
	if err := c.patchResource(ctx, "/api/tags/"+string(id), tag, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteTag deletes a tag by its ID.
func (c *httpClient) DeleteTag(ctx context.Context, id ID) error {
	return c.deleteResource(ctx, "/api/tags/"+string(id))
}

// UploadTagImage uploads an image for a tag.
func (c *httpClient) UploadTagImage(ctx context.Context, id ID, file []byte) (*Tag, error) {
	var tag Tag
	if err := c.uploadFile(ctx, "/api/tags/"+string(id)+"/image", file, "file", &tag); err != nil {
		return nil, err
	}
	return &tag, nil
}

// ListTagItems retrieves the items of a tag.
func (c *httpClient) ListTagItems(ctx context.Context, id ID, queryParams ...string) ([]*Item, error) {
	var items []*Item
	path := fmt.Sprintf("/api/tags/%s/items", id)
	if err := c.listResources(ctx, path, &items, queryParams...); err != nil {
		return nil, err
	}
	return items, nil
}

// GetCategoryOfTag retrieves the category associated with a tag.
func (c *httpClient) GetCategoryOfTag(ctx context.Context, id ID) (*TagCategory, error) {
	var category TagCategory
	if err := c.getResource(ctx, "/api/tags/"+string(id)+"/category", &category); err != nil {
		return nil, err
	}
	return &category, nil
}

// CreateTagCategory creates a new tag category with schema validation.
func (c *httpClient) CreateTagCategory(ctx context.Context, category *TagCategory) (*TagCategory, error) {
	if err := c.validateTagCategory(category); err != nil {
		return nil, err
	}
	var result TagCategory
	if err := c.postResource(ctx, "/api/tag_categories", category, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTagCategory retrieves a tag category by its ID.
func (c *httpClient) GetTagCategory(ctx context.Context, id ID) (*TagCategory, error) {
	var category TagCategory
	if err := c.getResource(ctx, "/api/tag_categories/"+string(id), &category); err != nil {
		return nil, err
	}
	return &category, nil
}

// ListTagCategories retrieves a list of tag categories.
func (c *httpClient) ListTagCategories(ctx context.Context, queryParams ...string) ([]*TagCategory, error) {
	var categories []*TagCategory
	if err := c.listResources(ctx, "/api/tag_categories", &categories, queryParams...); err != nil {
		return nil, err
	}
	return categories, nil
}

// UpdateTagCategory replaces an existing tag category.
func (c *httpClient) UpdateTagCategory(ctx context.Context, id ID, category *TagCategory) (*TagCategory, error) {
	var result TagCategory
	if err := c.putResource(ctx, "/api/tag_categories/"+string(id), category, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchTagCategory updates an existing tag category partially.
func (c *httpClient) PatchTagCategory(ctx context.Context, id ID, category *TagCategory) (*TagCategory, error) {
	var result TagCategory
	if err := c.patchResource(ctx, "/api/tag_categories/"+string(id), category, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteTagCategory deletes a tag category by its ID.
func (c *httpClient) DeleteTagCategory(ctx context.Context, id ID) error {
	return c.deleteResource(ctx, "/api/tag_categories/"+string(id))
}

// ListTagCategoryTags retrieves the tags of a tag category.
func (c *httpClient) ListTagCategoryTags(ctx context.Context, id ID, queryParams ...string) ([]*Tag, error) {
	var tags []*Tag
	path := fmt.Sprintf("/api/tag_categories/%s/tags", id)
	if err := c.listResources(ctx, path, &tags, queryParams...); err != nil {
		return nil, err
	}
	return tags, nil
}

// CreateTemplate creates a new template with schema validation.
func (c *httpClient) CreateTemplate(ctx context.Context, template *Template) (*Template, error) {
	if err := c.validateTemplate(template); err != nil {
		return nil, err
	}
	var result Template
	if err := c.postResource(ctx, "/api/templates", template, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTemplate retrieves a template by its ID.
func (c *httpClient) GetTemplate(ctx context.Context, id ID) (*Template, error) {
	var template Template
	if err := c.getResource(ctx, "/api/templates/"+string(id), &template); err != nil {
		return nil, err
	}
	return &template, nil
}

// ListTemplates retrieves a list of templates.
func (c *httpClient) ListTemplates(ctx context.Context, queryParams ...string) ([]*Template, error) {
	var templates []*Template
	if err := c.listResources(ctx, "/api/templates", &templates, queryParams...); err != nil {
		return nil, err
	}
	return templates, nil
}

// UpdateTemplate replaces an existing template.
func (c *httpClient) UpdateTemplate(ctx context.Context, id ID, template *Template) (*Template, error) {
	var result Template
	if err := c.putResource(ctx, "/api/templates/"+string(id), template, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchTemplate updates an existing template partially.
func (c *httpClient) PatchTemplate(ctx context.Context, id ID, template *Template) (*Template, error) {
	var result Template
	if err := c.patchResource(ctx, "/api/templates/"+string(id), template, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteTemplate deletes a template by its ID.
func (c *httpClient) DeleteTemplate(ctx context.Context, id ID) error {
	return c.deleteResource(ctx, "/api/templates/"+string(id))
}

// GetUser retrieves a user by its ID.
func (c *httpClient) GetUser(ctx context.Context, id ID) (*User, error) {
	var user User
	if err := c.getResource(ctx, "/api/users/"+string(id), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// ListUsers retrieves a list of users.
func (c *httpClient) ListUsers(ctx context.Context, queryParams ...string) ([]*User, error) {
	var users []*User
	if err := c.listResources(ctx, "/api/users", &users, queryParams...); err != nil {
		return nil, err
	}
	return users, nil
}

// CreateWish creates a new wish with schema validation.
func (c *httpClient) CreateWish(ctx context.Context, wish *Wish) (*Wish, error) {
	if err := c.validateWish(wish); err != nil {
		return nil, err
	}
	var result Wish
	if err := c.postResource(ctx, "/api/wishes", wish, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWish retrieves a wish by its ID.
func (c *httpClient) GetWish(ctx context.Context, id ID) (*Wish, error) {
	var wish Wish
	if err := c.getResource(ctx, "/api/wishes/"+string(id), &wish); err != nil {
		return nil, err
	}
	return &wish, nil
}

// ListWishes retrieves a list of wishes.
func (c *httpClient) ListWishes(ctx context.Context, queryParams ...string) ([]*Wish, error) {
	var wishes []*Wish
	if err := c.listResources(ctx, "/api/wishes", &wishes, queryParams...); err != nil {
		return nil, err
	}
	return wishes, nil
}

// UpdateWish replaces an existing wish.
func (c *httpClient) UpdateWish(ctx context.Context, id ID, wish *Wish) (*Wish, error) {
	var result Wish
	if err := c.putResource(ctx, "/api/wishes/"+string(id), wish, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchWish updates an existing wish partially.
func (c *httpClient) PatchWish(ctx context.Context, id ID, wish *Wish) (*Wish, error) {
	var result Wish
	if err := c.patchResource(ctx, "/api/wishes/"+string(id), wish, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteWish deletes a wish by its ID.
func (c *httpClient) DeleteWish(ctx context.Context, id ID) error {
	return c.deleteResource(ctx, "/api/wishes/"+string(id))
}

// UploadWishImage uploads an image for a wish.
func (c *httpClient) UploadWishImage(ctx context.Context, id ID, file []byte) (*Wish, error) {
	var wish Wish
	if err := c.uploadFile(ctx, "/api/wishes/"+string(id)+"/image", file, "file", &wish); err != nil {
		return nil, err
	}
	return &wish, nil
}

// GetWishWishlist retrieves the wishlist associated with a wish.
func (c *httpClient) GetWishWishlist(ctx context.Context, id ID) (*Wishlist, error) {
	var wishlist Wishlist
	if err := c.getResource(ctx, "/api/wishes/"+string(id)+"/wishlist", &wishlist); err != nil {
		return nil, err
	}
	return &wishlist, nil
}

// CreateWishlist creates a new wishlist with schema validation.
func (c *httpClient) CreateWishlist(ctx context.Context, wishlist *Wishlist) (*Wishlist, error) {
	if err := c.validateWishlist(wishlist); err != nil {
		return nil, err
	}
	var result Wishlist
	if err := c.postResource(ctx, "/api/wishlists", wishlist, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWishlist retrieves a wishlist by its ID.
func (c *httpClient) GetWishlist(ctx context.Context, id ID) (*Wishlist, error) {
	var wishlist Wishlist
	if err := c.getResource(ctx, "/api/wishlists/"+string(id), &wishlist); err != nil {
		return nil, err
	}
	return &wishlist, nil
}

// ListWishlists retrieves a list of wishlists.
func (c *httpClient) ListWishlists(ctx context.Context, queryParams ...string) ([]*Wishlist, error) {
	var wishlists []*Wishlist
	if err := c.listResources(ctx, "/api/wishlists", &wishlists, queryParams...); err != nil {
		return nil, err
	}
	return wishlists, nil
}

// UpdateWishlist replaces an existing wishlist.
func (c *httpClient) UpdateWishlist(ctx context.Context, id ID, wishlist *Wishlist) (*Wishlist, error) {
	var result Wishlist
	if err := c.putResource(ctx, "/api/wishlists/"+string(id), wishlist, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchWishlist updates an existing wishlist partially.
func (c *httpClient) PatchWishlist(ctx context.Context, id ID, wishlist *Wishlist) (*Wishlist, error) {
	var result Wishlist
	if err := c.patchResource(ctx, "/api/wishlists/"+string(id), wishlist, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteWishlist deletes a wishlist by its ID.
func (c *httpClient) DeleteWishlist(ctx context.Context, id ID) error {
	return c.deleteResource(ctx, "/api/wishlists/"+string(id))
}

// ListWishlistWishes retrieves the wishes in a wishlist.
func (c *httpClient) ListWishlistWishes(ctx context.Context, id ID, queryParams ...string) ([]*Wish, error) {
	var wishes []*Wish
	path := fmt.Sprintf("/api/wishlists/%s/wishes", id)
	if err := c.listResources(ctx, path, &wishes, queryParams...); err != nil {
		return nil, err
	}
	return wishes, nil
}

// ListWishlistChildren retrieves the children of a wishlist.
func (c *httpClient) ListWishlistChildren(ctx context.Context, id ID, queryParams ...string) ([]*Wishlist, error) {
	var wishlists []*Wishlist
	path := fmt.Sprintf("/api/wishlists/%s/children", id)
	if err := c.listResources(ctx, path, &wishlists, queryParams...); err != nil {
		return nil, err
	}
	return wishlists, nil
}

// UploadWishlistImage uploads an image for a wishlist.
func (c *httpClient) UploadWishlistImage(ctx context.Context, id ID, file []byte) (*Wishlist, error) {
	var wishlist Wishlist
	if err := c.uploadFile(ctx, "/api/wishlists/"+string(id)+"/image", file, "file", &wishlist); err != nil {
		return nil, err
	}
	return &wishlist, nil
}

// GetWishlistParent retrieves the parent wishlist of a wishlist.
func (c *httpClient) GetWishlistParent(ctx context.Context, id ID) (*Wishlist, error) {
	var wishlist Wishlist
	if err := c.getResource(ctx, "/api/wishlists/"+string(id)+"/parent", &wishlist); err != nil {
		return nil, err
	}
	return &wishlist, nil
}
