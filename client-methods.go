package koiApi

import (
	"fmt"
)

type Client interface {
	CheckLogin() (string, error)                                                // HTTP POST /api/authentication_token
	GetMetrics() (*Metrics, error)                                              // HTTP GET /api/metrics
	CreateAlbum(album *Album) (*Album, error)                                   // HTTP POST /api/albums
	GetAlbum(id ID) (*Album, error)                                             // HTTP GET /api/albums/{id}
	ListAlbums(queryParams ...string) ([]*Album, error)                         // HTTP GET /api/albums
	UpdateAlbum(id ID, album *Album) (*Album, error)                            // HTTP PUT /api/albums/{id}
	PatchAlbum(id ID, album *Album) (*Album, error)                             // HTTP PATCH /api/albums/{id}
	DeleteAlbum(id ID) error                                                    // HTTP DELETE /api/albums/{id}
	ListAlbumChildren(id ID, queryParams ...string) ([]*Album, error)           // HTTP GET /api/albums/{id}/children
	UploadAlbumImage(id ID, file []byte) (*Album, error)                        // HTTP POST /api/albums/{id}/image
	GetAlbumParent(id ID) (*Album, error)                                       // HTTP GET /api/albums/{id}/parent
	ListAlbumPhotos(id ID, queryParams ...string) ([]*Photo, error)             // HTTP GET /api/albums/{id}/photos
	CreateChoiceList(choiceList *ChoiceList) (*ChoiceList, error)               // HTTP POST /api/choice_lists
	GetChoiceList(id ID) (*ChoiceList, error)                                   // HTTP GET /api/choice_lists/{id}
	ListChoiceLists(queryParams ...string) ([]*ChoiceList, error)               // HTTP GET /api/choice_lists
	UpdateChoiceList(id ID, choiceList *ChoiceList) (*ChoiceList, error)        // HTTP PUT /api/choice_lists/{id}
	PatchChoiceList(id ID, choiceList *ChoiceList) (*ChoiceList, error)         // HTTP PATCH /api/choice_lists/{id}
	DeleteChoiceList(id ID) error                                               // HTTP DELETE /api/choice_lists/{id}
	CreateCollection(collection *Collection) (*Collection, error)               // HTTP POST /api/collections
	GetCollection(id ID) (*Collection, error)                                   // HTTP GET /api/collections/{id}
	ListCollections(queryParams ...string) ([]*Collection, error)               // HTTP GET /api/collections
	UpdateCollection(id ID, collection *Collection) (*Collection, error)        // HTTP PUT /api/collections/{id}
	PatchCollection(id ID, collection *Collection) (*Collection, error)         // HTTP PATCH /api/collections/{id}
	DeleteCollection(id ID) error                                               // HTTP DELETE /api/collections/{id}
	ListCollectionChildren(id ID, queryParams ...string) ([]*Collection, error) // HTTP GET /api/collections/{id}/children
	UploadCollectionImage(id ID, file []byte) (*Collection, error)              // HTTP POST /api/collections/{id}/image
	GetCollectionParent(id ID) (*Collection, error)                             // HTTP GET /api/collections/{id}/parent
	ListCollectionItems(id ID, queryParams ...string) ([]*Item, error)          // HTTP GET /api/collections/{id}/items
	ListCollectionData(id ID, queryParams ...string) ([]*Datum, error)          // HTTP GET /api/collections/{id}/data
	GetCollectionDefaultTemplate(id ID) (*Template, error)                      // HTTP GET /api/collections/{id}/items_default_template
	CreateDatum(datum *Datum) (*Datum, error)                                   // HTTP POST /api/data
	GetDatum(id ID) (*Datum, error)                                             // HTTP GET /api/data/{id}
	ListData(queryParams ...string) ([]*Datum, error)                           // HTTP GET /api/data
	UpdateDatum(id ID, datum *Datum) (*Datum, error)                            // HTTP PUT /api/data/{id}
	PatchDatum(id ID, datum *Datum) (*Datum, error)                             // HTTP PATCH /api/data/{id}
	DeleteDatum(id ID) error                                                    // HTTP DELETE /api/data/{id}
	UploadDatumFile(id ID, file []byte) (*Datum, error)                         // HTTP POST /api/data/{id}/file
	UploadDatumImage(id ID, image []byte) (*Datum, error)                       // HTTP POST /api/data/{id}/image
	UploadDatumVideo(id ID, video []byte) (*Datum, error)                       // HTTP POST /api/data/{id}/video
	GetDatumItem(id ID) (*Item, error)                                          // HTTP GET /api/data/{id}/item
	GetDatumCollection(id ID) (*Collection, error)                              // HTTP GET /api/data/{id}/collection
	CreateField(field *Field) (*Field, error)                                   // HTTP POST /api/fields
	GetField(id ID) (*Field, error)                                             // HTTP GET /api/fields/{id}
	ListFields(queryParams ...string) ([]*Field, error)                         // HTTP GET /api/fields
	UpdateField(id ID, field *Field) (*Field, error)                            // HTTP PUT /api/fields/{id}
	PatchField(id ID, field *Field) (*Field, error)                             // HTTP PATCH /api/fields/{id}
	DeleteField(id ID) error                                                    // HTTP DELETE /api/fields/{id}
	GetFieldTemplate(id ID) (*Template, error)                                  // HTTP GET /api/fields/{id}/template
	ListTemplateFields(templateid ID, queryParams ...string) ([]*Field, error)  // HTTP GET /api/templates/{id}/fields
	ListInventories(queryParams ...string) ([]*Inventory, error)                // HTTP GET /api/inventories
	GetInventory(id ID) (*Inventory, error)                                     // HTTP GET /api/inventories/{id}
	DeleteInventory(id ID) error                                                // HTTP DELETE /api/inventories/{id}
	CreateItem(item *Item) (*Item, error)                                       // HTTP POST /api/items
	GetItem(id ID) (*Item, error)                                               // HTTP GET /api/items/{id}
	ListItems(queryParams ...string) ([]*Item, error)                           // HTTP GET /api/items
	UpdateItem(id ID, item *Item) (*Item, error)                                // HTTP PUT /api/items/{id}
	PatchItem(id ID, item *Item) (*Item, error)                                 // HTTP PATCH /api/items/{id}
	DeleteItem(id ID) error                                                     // HTTP DELETE /api/items/{id}
	SearchItems(queryParams ...string) ([]*Item, error)                         // HTTP GET /search
	UploadItemImage(id ID, file []byte) (*Item, error)                          // HTTP POST /api/items/{id}/image
	ListItemRelatedItems(id ID, queryParams ...string) ([]*Item, error)         // HTTP GET /api/items/{id}/related_items
	ListItemLoans(id ID, queryParams ...string) ([]*Loan, error)                // HTTP GET /api/items/{id}/loans
	ListItemTags(id ID, queryParams ...string) ([]*Tag, error)                  // HTTP GET /api/items/{id}/tags
	ListItemData(id ID, queryParams ...string) ([]*Datum, error)                // HTTP GET /api/items/{id}/data
	GetItemCollection(id ID) (*Collection, error)                               // HTTP GET /api/items/{id}/collection
	CreateLoan(loan *Loan) (*Loan, error)                                       // HTTP POST /api/loans
	GetLoan(id ID) (*Loan, error)                                               // HTTP GET /api/loans/{id}
	ListLoans(queryParams ...string) ([]*Loan, error)                           // HTTP GET /api/loans
	UpdateLoan(id ID, loan *Loan) (*Loan, error)                                // HTTP PUT /api/loans/{id}
	PatchLoan(id ID, loan *Loan) (*Loan, error)                                 // HTTP PATCH /api/loans/{id}
	DeleteLoan(id ID) error                                                     // HTTP DELETE /api/loans/{id}
	GetLoanItem(id ID) (*Item, error)                                           // HTTP GET /api/loans/{id}/item
	GetLog(id ID) (*Log, error)                                                 // HTTP GET /api/logs/{id}
	ListLogs(queryParams ...string) ([]*Log, error)                             // HTTP GET /api/logs
	CreatePhoto(photo *Photo) (*Photo, error)                                   // HTTP POST /api/photos
	GetPhoto(id ID) (*Photo, error)                                             // HTTP GET /api/photos/{id}
	ListPhotos(queryParams ...string) ([]*Photo, error)                         // HTTP GET /api/photos
	UpdatePhoto(id ID, photo *Photo) (*Photo, error)                            // HTTP PUT /api/photos/{id}
	PatchPhoto(id ID, photo *Photo) (*Photo, error)                             // HTTP PATCH /api/photos/{id}
	DeletePhoto(id ID) error                                                    // HTTP DELETE /api/photos/{id}
	UploadPhotoImage(id ID, file []byte) (*Photo, error)                        // HTTP POST /api/photos/{id}/image
	GetPhotoAlbum(id ID) (*Album, error)                                        // HTTP GET /api/photos/{id}/album
	CreateTag(tag *Tag) (*Tag, error)                                           // HTTP POST /api/tags
	GetTag(id ID) (*Tag, error)                                                 // HTTP GET /api/tags/{id}
	ListTags(queryParams ...string) ([]*Tag, error)                             // HTTP GET /api/tags
	UpdateTag(id ID, tag *Tag) (*Tag, error)                                    // HTTP PUT /api/tags/{id}
	PatchTag(id ID, tag *Tag) (*Tag, error)                                     // HTTP PATCH /api/tags/{id}
	DeleteTag(id ID) error                                                      // HTTP DELETE /api/tags/{id}
	UploadTagImage(id ID, file []byte) (*Tag, error)                            // HTTP POST /api/tags/{id}/image
	ListTagItems(id ID, queryParams ...string) ([]*Item, error)                 // HTTP GET /api/tags/{id}/items
	GetCategoryOfTag(id ID) (*TagCategory, error)                               // HTTP GET /api/tags/{id}/category
	CreateTagCategory(category *TagCategory) (*TagCategory, error)              // HTTP POST /api/tag_categories
	GetTagCategory(id ID) (*TagCategory, error)                                 // HTTP GET /api/tag_categories/{id}
	ListTagCategories(queryParams ...string) ([]*TagCategory, error)            // HTTP GET /api/tag_categories
	UpdateTagCategory(id ID, category *TagCategory) (*TagCategory, error)       // HTTP PUT /api/tag_categories/{id}
	PatchTagCategory(id ID, category *TagCategory) (*TagCategory, error)        // HTTP PATCH /api/tag_categories/{id}
	DeleteTagCategory(id ID) error                                              // HTTP DELETE /api/tag_categories/{id}
	ListTagCategoryTags(id ID, queryParams ...string) ([]*Tag, error)           // HTTP GET /api/tag_categories/{id}/tags
	CreateTemplate(template *Template) (*Template, error)                       // HTTP POST /api/templates
	GetTemplate(id ID) (*Template, error)                                       // HTTP GET /api/templates/{id}
	ListTemplates(queryParams ...string) ([]*Template, error)                   // HTTP GET /api/templates
	UpdateTemplate(id ID, template *Template) (*Template, error)                // HTTP PUT /api/templates/{id}
	PatchTemplate(id ID, template *Template) (*Template, error)                 // HTTP PATCH /api/templates/{id}
	DeleteTemplate(id ID) error                                                 // HTTP DELETE /api/templates/{id}
	GetUser(id ID) (*User, error)                                               // HTTP GET /api/users/{id}
	ListUsers(queryParams ...string) ([]*User, error)                           // HTTP GET /api/users
	CreateWish(wish *Wish) (*Wish, error)                                       // HTTP POST /api/wishes
	GetWish(id ID) (*Wish, error)                                               // HTTP GET /api/wishes/{id}
	ListWishes(queryParams ...string) ([]*Wish, error)                          // HTTP GET /api/wishes
	UpdateWish(id ID, wish *Wish) (*Wish, error)                                // HTTP PUT /api/wishes/{id}
	PatchWish(id ID, wish *Wish) (*Wish, error)                                 // HTTP PATCH /api/wishes/{id}
	DeleteWish(id ID) error                                                     // HTTP DELETE /api/wishes/{id}
	UploadWishImage(id ID, file []byte) (*Wish, error)                          // HTTP POST /api/wishes/{id}/image
	GetWishWishlist(id ID) (*Wishlist, error)                                   // HTTP GET /api/wishes/{id}/wishlist
	CreateWishlist(wishlist *Wishlist) (*Wishlist, error)                       // HTTP POST /api/wishlists
	GetWishlist(id ID) (*Wishlist, error)                                       // HTTP GET /api/wishlists/{id}
	ListWishlists(queryParams ...string) ([]*Wishlist, error)                   // HTTP GET /api/wishlists
	UpdateWishlist(id ID, wishlist *Wishlist) (*Wishlist, error)                // HTTP PUT /api/wishlists/{id}
	PatchWishlist(id ID, wishlist *Wishlist) (*Wishlist, error)                 // HTTP PATCH /api/wishlists/{id}
	DeleteWishlist(id ID) error                                                 // HTTP DELETE /api/wishlists/{id}
	ListWishlistWishes(id ID, queryParams ...string) ([]*Wish, error)           // HTTP GET /api/wishlists/{id}/wishes
	ListWishlistChildren(id ID, queryParams ...string) ([]*Wishlist, error)     // HTTP GET /api/wishlists/{id}/children
	UploadWishlistImage(id ID, file []byte) (*Wishlist, error)                  // HTTP POST /api/wishlists/{id}/image
	GetWishlistParent(id ID) (*Wishlist, error)                                 // HTTP GET /api/wishlists/{id}/parent
	GetResponse() string
	PrintError()
	DeleteAllData() error
}

// GetMetrics retrieves system or user-specific metrics.
func (c *koiClient) GetMetrics() (*Metrics, error) {
	var metrics Metrics
	if err := c.getResource("/api/metrics", &metrics); err != nil {
		return nil, err
	}
	return &metrics, nil
}

// CreateAlbum creates a new album with schema validation.
func (c *koiClient) CreateAlbum(album *Album) (*Album, error) {
	if err := c.validateAlbum(album); err != nil {
		return nil, err
	}
	var result Album
	if err := c.postResource("/api/albums", album, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAlbum retrieves an album by its ID.
func (c *koiClient) GetAlbum(id ID) (*Album, error) {
	var album Album
	if err := c.getResource("/api/albums/"+string(id), &album); err != nil {
		return nil, err
	}
	return &album, nil
}

// ListAlbums retrieves a list of albums.
func (c *koiClient) ListAlbums(queryParams ...string) ([]*Album, error) {
	var albums []*Album
	if err := c.listResources("/api/albums", &albums, queryParams...); err != nil {
		return nil, err
	}
	return albums, nil
}

// UpdateAlbum replaces an existing album.
func (c *koiClient) UpdateAlbum(id ID, album *Album) (*Album, error) {
	var result Album
	if err := c.putResource("/api/albums/"+string(id), album, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchAlbum updates an existing album partially.
func (c *koiClient) PatchAlbum(id ID, album *Album) (*Album, error) {
	var result Album
	if err := c.patchResource("/api/albums/"+string(id), album, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteAlbum deletes an album by its ID.
func (c *koiClient) DeleteAlbum(id ID) error {
	return c.deleteResource("/api/albums/" + string(id))
}

// ListAlbumChildren retrieves the children of an album.
func (c *koiClient) ListAlbumChildren(id ID, queryParams ...string) ([]*Album, error) {
	var albums []*Album
	path := fmt.Sprintf("/api/albums/%s/children", id)
	if err := c.listResources(path, &albums, queryParams...); err != nil {
		return nil, err
	}
	return albums, nil
}

// UploadAlbumImage uploads an image for an album.
func (c *koiClient) UploadAlbumImage(id ID, file []byte) (*Album, error) {
	var album Album
	if err := c.uploadFile("/api/albums/"+string(id)+"/image", file, "file", &album); err != nil {
		return nil, err
	}
	return &album, nil
}

// GetAlbumParent retrieves the parent album of an album.
func (c *koiClient) GetAlbumParent(id ID) (*Album, error) {
	var album Album
	if err := c.getResource("/api/albums/"+string(id)+"/parent", &album); err != nil {
		return nil, err
	}
	return &album, nil
}

// ListAlbumPhotos retrieves the photos in an album.
func (c *koiClient) ListAlbumPhotos(id ID, queryParams ...string) ([]*Photo, error) {
	var photos []*Photo
	path := fmt.Sprintf("/api/albums/%s/photos", id)
	if err := c.listResources(path, &photos, queryParams...); err != nil {
		return nil, err
	}
	return photos, nil
}

// CreateChoiceList creates a new choice list with schema validation.
func (c *koiClient) CreateChoiceList(choiceList *ChoiceList) (*ChoiceList, error) {
	if err := c.validateChoiceList(choiceList); err != nil {
		return nil, err
	}
	var result ChoiceList
	if err := c.postResource("/api/choice_lists", choiceList, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetChoiceList retrieves a choice list by its ID.
func (c *koiClient) GetChoiceList(id ID) (*ChoiceList, error) {
	var choiceList ChoiceList
	if err := c.getResource("/api/choice_lists/"+string(id), &choiceList); err != nil {
		return nil, err
	}
	return &choiceList, nil
}

// ListChoiceLists retrieves a list of choice lists.
func (c *koiClient) ListChoiceLists(queryParams ...string) ([]*ChoiceList, error) {
	var choiceLists []*ChoiceList
	if err := c.listResources("/api/choice_lists", &choiceLists, queryParams...); err != nil {
		return nil, err
	}
	return choiceLists, nil
}

// UpdateChoiceList replaces an existing choice list.
func (c *koiClient) UpdateChoiceList(id ID, choiceList *ChoiceList) (*ChoiceList, error) {
	var result ChoiceList
	if err := c.putResource("/api/choice_lists/"+string(id), choiceList, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchChoiceList updates an existing choice list partially.
func (c *koiClient) PatchChoiceList(id ID, choiceList *ChoiceList) (*ChoiceList, error) {
	var result ChoiceList
	if err := c.patchResource("/api/choice_lists/"+string(id), choiceList, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteChoiceList deletes a choice list by its ID.
func (c *koiClient) DeleteChoiceList(id ID) error {
	return c.deleteResource("/api/choice_lists/" + string(id))
}

// CreateCollection creates a new collection with schema validation.
func (c *koiClient) CreateCollection(collection *Collection) (*Collection, error) {
	if err := c.validateCollection(collection); err != nil {
		return nil, err
	}
	var result Collection
	if err := c.postResource("/api/collections", collection, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCollection retrieves a collection by its ID.
func (c *koiClient) GetCollection(id ID) (*Collection, error) {
	var collection Collection
	if err := c.getResource("/api/collections/"+string(id), &collection); err != nil {
		return nil, err
	}
	return &collection, nil
}

// ListCollections retrieves a list of collections.
func (c *koiClient) ListCollections(queryParams ...string) ([]*Collection, error) {
	var collections []*Collection
	if err := c.listResources("/api/collections", &collections, queryParams...); err != nil {
		return nil, err
	}
	return collections, nil
}

// UpdateCollection replaces an existing collection.
func (c *koiClient) UpdateCollection(id ID, collection *Collection) (*Collection, error) {
	var result Collection
	if err := c.putResource("/api/collections/"+string(id), collection, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchCollection updates an existing collection partially.
func (c *koiClient) PatchCollection(id ID, collection *Collection) (*Collection, error) {
	var result Collection
	if err := c.patchResource("/api/collections/"+string(id), collection, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteCollection deletes a collection by its ID.
func (c *koiClient) DeleteCollection(id ID) error {
	return c.deleteResource("/api/collections/" + string(id))
}

// ListCollectionChildren retrieves the children of a collection.
func (c *koiClient) ListCollectionChildren(id ID, queryParams ...string) ([]*Collection, error) {
	var collections []*Collection
	path := fmt.Sprintf("/api/collections/%s/children", id)
	if err := c.listResources(path, &collections, queryParams...); err != nil {
		return nil, err
	}
	return collections, nil
}

// UploadCollectionImage uploads an image for a collection.
func (c *koiClient) UploadCollectionImage(id ID, file []byte) (*Collection, error) {
	var collection Collection
	if err := c.uploadFile("/api/collections/"+string(id)+"/image", file, "file", &collection); err != nil {
		return nil, err
	}
	return &collection, nil
}

// GetCollectionParent retrieves the parent collection of a collection.
func (c *koiClient) GetCollectionParent(id ID) (*Collection, error) {
	var collection Collection
	if err := c.getResource("/api/collections/"+string(id)+"/parent", &collection); err != nil {
		return nil, err
	}
	return &collection, nil
}

// ListCollectionItems retrieves the items in a collection.
func (c *koiClient) ListCollectionItems(id ID, queryParams ...string) ([]*Item, error) {
	var items []*Item
	path := fmt.Sprintf("/api/collections/%s/items", id)
	if err := c.listResources(path, &items, queryParams...); err != nil {
		return nil, err
	}
	return items, nil
}

// ListCollectionData retrieves the data in a collection.
func (c *koiClient) ListCollectionData(id ID, queryParams ...string) ([]*Datum, error) {
	var data []*Datum
	path := fmt.Sprintf("/api/collections/%s/data", id)
	if err := c.listResources(path, &data, queryParams...); err != nil {
		return nil, err
	}
	return data, nil
}

// GetCollectionDefaultTemplate retrieves the default template for items in a collection.
func (c *koiClient) GetCollectionDefaultTemplate(id ID) (*Template, error) {
	var template Template
	if err := c.getResource("/api/collections/"+string(id)+"/items_default_template", &template); err != nil {
		return nil, err
	}
	return &template, nil
}

// CreateDatum creates a new datum with schema validation.
func (c *koiClient) CreateDatum(datum *Datum) (*Datum, error) {
	if err := c.validateDatum(datum); err != nil {
		return nil, err
	}
	var result Datum
	if err := c.postResource("/api/data", datum, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDatum retrieves a datum by its ID.
func (c *koiClient) GetDatum(id ID) (*Datum, error) {
	var datum Datum
	if err := c.getResource("/api/data/"+string(id), &datum); err != nil {
		return nil, err
	}
	return &datum, nil
}

// ListData retrieves a list of data.
func (c *koiClient) ListData(queryParams ...string) ([]*Datum, error) {
	var data []*Datum
	if err := c.listResources("/api/data", &data, queryParams...); err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateDatum replaces an existing datum.
func (c *koiClient) UpdateDatum(id ID, datum *Datum) (*Datum, error) {
	var result Datum
	if err := c.putResource("/api/data/"+string(id), datum, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchDatum updates an existing datum partially.
func (c *koiClient) PatchDatum(id ID, datum *Datum) (*Datum, error) {
	var result Datum
	if err := c.patchResource("/api/data/"+string(id), datum, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteDatum deletes a datum by its ID.
func (c *koiClient) DeleteDatum(id ID) error {
	return c.deleteResource("/api/data/" + string(id))
}

// UploadDatumFile uploads a file for a datum.
func (c *koiClient) UploadDatumFile(id ID, file []byte) (*Datum, error) {
	var datum Datum
	if err := c.uploadFile("/api/data/"+string(id)+"/file", file, "fileFile", &datum); err != nil {
		return nil, err
	}
	return &datum, nil
}

// UploadDatumImage uploads an image for a datum.
func (c *koiClient) UploadDatumImage(id ID, image []byte) (*Datum, error) {
	var datum Datum
	if err := c.uploadFile("/api/data/"+string(id)+"/image", image, "fileImage", &datum); err != nil {
		return nil, err
	}
	return &datum, nil
}

// UploadDatumVideo uploads a video for a datum.
func (c *koiClient) UploadDatumVideo(id ID, video []byte) (*Datum, error) {
	var datum Datum
	if err := c.uploadFile("/api/data/"+string(id)+"/video", video, "fileVideo", &datum); err != nil {
		return nil, err
	}
	return &datum, nil
}

// GetDatumItem retrieves the item associated with a datum.
func (c *koiClient) GetDatumItem(id ID) (*Item, error) {
	var item Item
	if err := c.getResource("/api/data/"+string(id)+"/item", &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// GetDatumCollection retrieves the collection associated with a datum.
func (c *koiClient) GetDatumCollection(id ID) (*Collection, error) {
	var collection Collection
	if err := c.getResource("/api/data/"+string(id)+"/collection", &collection); err != nil {
		return nil, err
	}
	return &collection, nil
}

// CreateField creates a new field with schema validation.
func (c *koiClient) CreateField(field *Field) (*Field, error) {
	if err := c.validateField(field); err != nil {
		return nil, err
	}
	var result Field
	if err := c.postResource("/api/fields", field, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetField retrieves a field by its ID.
func (c *koiClient) GetField(id ID) (*Field, error) {
	var field Field
	if err := c.getResource("/api/fields/"+string(id), &field); err != nil {
		return nil, err
	}
	return &field, nil
}

// ListFields retrieves a list of fields.
func (c *koiClient) ListFields(queryParams ...string) ([]*Field, error) {
	var fields []*Field
	if err := c.listResources("/api/fields", &fields, queryParams...); err != nil {
		return nil, err
	}
	return fields, nil
}

// UpdateField replaces an existing field.
func (c *koiClient) UpdateField(id ID, field *Field) (*Field, error) {
	var result Field
	if err := c.putResource("/api/fields/"+string(id), field, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchField updates an existing field partially.
func (c *koiClient) PatchField(id ID, field *Field) (*Field, error) {
	var result Field
	if err := c.patchResource("/api/fields/"+string(id), field, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteField deletes a field by its ID.
func (c *koiClient) DeleteField(id ID) error {
	return c.deleteResource("/api/fields/" + string(id))
}

// GetFieldTemplate retrieves the template associated with a field.
func (c *koiClient) GetFieldTemplate(id ID) (*Template, error) {
	var template Template
	if err := c.getResource("/api/fields/"+string(id)+"/template", &template); err != nil {
		return nil, err
	}
	return &template, nil
}

// ListTemplateFields retrieves the fields of a template.
func (c *koiClient) ListTemplateFields(templateid ID, queryParams ...string) ([]*Field, error) {
	var fields []*Field
	path := fmt.Sprintf("/api/templates/%s/fields", templateid)
	if err := c.listResources(path, &fields, queryParams...); err != nil {
		return nil, err
	}
	return fields, nil
}

// ListInventories retrieves a list of inventories.
func (c *koiClient) ListInventories(queryParams ...string) ([]*Inventory, error) {
	var inventories []*Inventory
	if err := c.listResources("/api/inventories", &inventories, queryParams...); err != nil {
		return nil, err
	}
	return inventories, nil
}

// GetInventory retrieves an inventory by its ID.
func (c *koiClient) GetInventory(id ID) (*Inventory, error) {
	var inventory Inventory
	if err := c.getResource("/api/inventories/"+string(id), &inventory); err != nil {
		return nil, err
	}
	return &inventory, nil
}

// DeleteInventory deletes an inventory by its ID.
func (c *koiClient) DeleteInventory(id ID) error {
	return c.deleteResource("/api/inventories/" + string(id))
}

// CreateItem creates a new item.
func (c *koiClient) CreateItem(item *Item) (*Item, error) {
	if err := c.validateItem(item); err != nil {
		return nil, err
	}
	var result Item
	if err := c.postResource("/api/items", item, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetItem retrieves an item by its ID.
func (c *koiClient) GetItem(id ID) (*Item, error) {
	var item Item
	if err := c.getResource("/api/items/"+string(id), &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// ListItems retrieves a list of items.
func (c *koiClient) ListItems(queryParams ...string) ([]*Item, error) {
	var items []*Item
	if err := c.listResources("/api/items", &items, queryParams...); err != nil {
		return nil, err
	}
	return items, nil
}

// SearchItems retrieves a list of items that match a query
func (c *koiClient) SearchItems(queryParams ...string) ([]*Item, error) {
	var items []*Item
	if err := c.listResources("/search", &items, queryParams...); err != nil {
		return nil, err
	}
	return items, nil
}

// UpdateItem replaces an existing item.
func (c *koiClient) UpdateItem(id ID, item *Item) (*Item, error) {
	var result Item
	if err := c.putResource("/api/items/"+string(id), item, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchItem updates an existing item partially.
func (c *koiClient) PatchItem(id ID, item *Item) (*Item, error) {
	var result Item
	if err := c.patchResource("/api/items/"+string(id), item, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteItem deletes an item by its ID.
func (c *koiClient) DeleteItem(id ID) error {
	return c.deleteResource("/api/items/" + string(id))
}

// UploadItemImage uploads an image for an item.
func (c *koiClient) UploadItemImage(id ID, file []byte) (*Item, error) {
	var item Item
	if err := c.uploadFile("/api/items/"+string(id)+"/image", file, "file", &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// ListItemRelatedItems retrieves the related items of an item.
func (c *koiClient) ListItemRelatedItems(id ID, queryParams ...string) ([]*Item, error) {
	var items []*Item
	path := fmt.Sprintf("/api/items/%s/related_items", id)
	if err := c.listResources(path, &items, queryParams...); err != nil {
		return nil, err
	}
	return items, nil
}

// ListItemLoans retrieves the loans of an item.
func (c *koiClient) ListItemLoans(id ID, queryParams ...string) ([]*Loan, error) {
	var loans []*Loan
	path := fmt.Sprintf("/api/items/%s/loans", id)
	if err := c.listResources(path, &loans, queryParams...); err != nil {
		return nil, err
	}
	return loans, nil
}

// ListItemTags retrieves the tags of an item.
func (c *koiClient) ListItemTags(id ID, queryParams ...string) ([]*Tag, error) {
	var tags []*Tag
	path := fmt.Sprintf("/api/items/%s/tags", id)
	if err := c.listResources(path, &tags, queryParams...); err != nil {
		return nil, err
	}
	return tags, nil
}

// ListItemData retrieves the data of an item.
func (c *koiClient) ListItemData(id ID, queryParams ...string) ([]*Datum, error) {
	var data []*Datum
	path := fmt.Sprintf("/api/items/%s/data", id)
	if err := c.listResources(path, &data, queryParams...); err != nil {
		return nil, err
	}
	return data, nil
}

// GetItemCollection retrieves the collection associated with an item.
func (c *koiClient) GetItemCollection(id ID) (*Collection, error) {
	var collection Collection
	if err := c.getResource("/api/items/"+string(id)+"/collection", &collection); err != nil {
		return nil, err
	}
	return &collection, nil
}

// / CreateLoan creates a new loan with schema validation.
func (c *koiClient) CreateLoan(loan *Loan) (*Loan, error) {
	if err := c.validateLoan(loan); err != nil {
		return nil, err
	}
	var result Loan
	if err := c.postResource("/api/loans", loan, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetLoan retrieves a loan by its ID.
func (c *koiClient) GetLoan(id ID) (*Loan, error) {
	var loan Loan
	if err := c.getResource("/api/loans/"+string(id), &loan); err != nil {
		return nil, err
	}
	return &loan, nil
}

// ListLoans retrieves a list of loans.
func (c *koiClient) ListLoans(queryParams ...string) ([]*Loan, error) {
	var loans []*Loan
	if err := c.listResources("/api/loans", &loans, queryParams...); err != nil {
		return nil, err
	}
	return loans, nil
}

// UpdateLoan replaces(an existing loan.
func (c *koiClient) UpdateLoan(id ID, loan *Loan) (*Loan, error) {
	var result Loan
	if err := c.putResource("/api/loans/"+string(id), loan, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchLoan updates an existing loan partially.
func (c *koiClient) PatchLoan(id ID, loan *Loan) (*Loan, error) {
	var result Loan
	if err := c.patchResource("/api/loans/"+string(id), loan, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteLoan deletes a loan by its ID.
func (c *koiClient) DeleteLoan(id ID) error {
	return c.deleteResource("/api/loans/" + string(id))
}

// GetLoanItem retrieves the item associated with a loan.
func (c *koiClient) GetLoanItem(id ID) (*Item, error) {
	var item Item
	if err := c.getResource("/api/loans/"+string(id)+"/item", &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// GetLog retrieves a log by its ID.
func (c *koiClient) GetLog(id ID) (*Log, error) {
	var log Log
	if err := c.getResource("/api/logs/"+string(id), &log); err != nil {
		return nil, err
	}
	return &log, nil
}

// ListLogs retrieves a list of logs.
func (c *koiClient) ListLogs(queryParams ...string) ([]*Log, error) {
	var logs []*Log
	if err := c.listResources("/api/logs", &logs, queryParams...); err != nil {
		return nil, err
	}
	return logs, nil
}

// CreatePhoto creates a new photo with schema validation.
func (c *koiClient) CreatePhoto(photo *Photo) (*Photo, error) {
	if err := c.validatePhoto(photo); err != nil {
		return nil, err
	}
	var result Photo
	if err := c.postResource("/api/photos", photo, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPhoto retrieves a photo by its ID.
func (c *koiClient) GetPhoto(id ID) (*Photo, error) {
	var photo Photo
	if err := c.getResource("/api/photos/"+string(id), &photo); err != nil {
		return nil, err
	}
	return &photo, nil
}

// ListPhotos retrieves a list of photos.
func (c *koiClient) ListPhotos(queryParams ...string) ([]*Photo, error) {
	var photos []*Photo
	if err := c.listResources("/api/photos", &photos, queryParams...); err != nil {
		return nil, err
	}
	return photos, nil
}

// UpdatePhoto replaces an existing photo.
func (c *koiClient) UpdatePhoto(id ID, photo *Photo) (*Photo, error) {
	var result Photo
	if err := c.putResource("/api/photos/"+string(id), photo, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchPhoto updates an existing photo partially.
func (c *koiClient) PatchPhoto(id ID, photo *Photo) (*Photo, error) {
	var result Photo
	if err := c.patchResource("/api/photos/"+string(id), photo, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeletePhoto deletes a photo by its ID.
func (c *koiClient) DeletePhoto(id ID) error {
	return c.deleteResource("/api/photos/" + string(id))
}

// UploadPhotoImage uploads an image for a photo.
func (c *koiClient) UploadPhotoImage(id ID, file []byte) (*Photo, error) {
	var photo Photo
	if err := c.uploadFile("/api/photos/"+string(id)+"/image", file, "file", &photo); err != nil {
		return nil, err
	}
	return &photo, nil
}

// GetPhotoAlbum retrieves the album associated with a photo.
func (c *koiClient) GetPhotoAlbum(id ID) (*Album, error) {
	var album Album
	if err := c.getResource("/api/photos/"+string(id)+"/album", &album); err != nil {
		return nil, err
	}
	return &album, nil
}

// CreateTag creates a new tag with schema validation.
func (c *koiClient) CreateTag(tag *Tag) (*Tag, error) {
	if err := c.validateTag(tag); err != nil {
		return nil, err
	}
	var result Tag
	if err := c.postResource("/api/tags", tag, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTag retrieves a tag by its ID.
func (c *koiClient) GetTag(id ID) (*Tag, error) {
	var tag Tag
	if err := c.getResource("/api/tags/"+string(id), &tag); err != nil {
		return nil, err
	}
	return &tag, nil
}

// ListTags retrieves a list of tags.
func (c *koiClient) ListTags(queryParams ...string) ([]*Tag, error) {
	var tags []*Tag
	if err := c.listResources("/api/tags", &tags, queryParams...); err != nil {
		return nil, err
	}
	return tags, nil
}

// UpdateTag replaces an existing tag.
func (c *koiClient) UpdateTag(id ID, tag *Tag) (*Tag, error) {
	var result Tag
	if err := c.putResource("/api/tags/"+string(id), tag, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchTag updates an existing tag partially.
func (c *koiClient) PatchTag(id ID, tag *Tag) (*Tag, error) {
	var result Tag
	if err := c.patchResource("/api/tags/"+string(id), tag, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteTag deletes a tag by its ID.
func (c *koiClient) DeleteTag(id ID) error {
	return c.deleteResource("/api/tags/" + string(id))
}

// UploadTagImage uploads an image for a tag.
func (c *koiClient) UploadTagImage(id ID, file []byte) (*Tag, error) {
	var tag Tag
	if err := c.uploadFile("/api/tags/"+string(id)+"/image", file, "file", &tag); err != nil {
		return nil, err
	}
	return &tag, nil
}

// ListTagItems retrieves the items of a tag.
func (c *koiClient) ListTagItems(id ID, queryParams ...string) ([]*Item, error) {
	var items []*Item
	path := fmt.Sprintf("/api/tags/%s/items", id)
	if err := c.listResources(path, &items, queryParams...); err != nil {
		return nil, err
	}
	return items, nil
}

// GetCategoryOfTag retrieves the category associated with a tag.
func (c *koiClient) GetCategoryOfTag(id ID) (*TagCategory, error) {
	var category TagCategory
	if err := c.getResource("/api/tags/"+string(id)+"/category", &category); err != nil {
		return nil, err
	}
	return &category, nil
}

// CreateTagCategory creates a new tag category with schema validation.
func (c *koiClient) CreateTagCategory(category *TagCategory) (*TagCategory, error) {
	if err := c.validateTagCategory(category); err != nil {
		return nil, err
	}
	var result TagCategory
	if err := c.postResource("/api/tag_categories", category, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTagCategory retrieves a tag category by its ID.
func (c *koiClient) GetTagCategory(id ID) (*TagCategory, error) {
	var category TagCategory
	if err := c.getResource("/api/tag_categories/"+string(id), &category); err != nil {
		return nil, err
	}
	return &category, nil
}

// ListTagCategories retrieves a list of tag categories.
func (c *koiClient) ListTagCategories(queryParams ...string) ([]*TagCategory, error) {
	var categories []*TagCategory
	if err := c.listResources("/api/tag_categories", &categories, queryParams...); err != nil {
		return nil, err
	}
	return categories, nil
}

// UpdateTagCategory replaces an existing tag category.
func (c *koiClient) UpdateTagCategory(id ID, category *TagCategory) (*TagCategory, error) {
	var result TagCategory
	if err := c.putResource("/api/tag_categories/"+string(id), category, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchTagCategory updates an existing tag category partially.
func (c *koiClient) PatchTagCategory(id ID, category *TagCategory) (*TagCategory, error) {
	var result TagCategory
	if err := c.patchResource("/api/tag_categories/"+string(id), category, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteTagCategory deletes a tag category by its ID.
func (c *koiClient) DeleteTagCategory(id ID) error {
	return c.deleteResource("/api/tag_categories/" + string(id))
}

// ListTagCategoryTags retrieves the tags of a tag category.
func (c *koiClient) ListTagCategoryTags(id ID, queryParams ...string) ([]*Tag, error) {
	var tags []*Tag
	path := fmt.Sprintf("/api/tag_categories/%s/tags", id)
	if err := c.listResources(path, &tags, queryParams...); err != nil {
		return nil, err
	}
	return tags, nil
}

// CreateTemplate creates a new template with schema validation.
func (c *koiClient) CreateTemplate(template *Template) (*Template, error) {
	if err := c.validateTemplate(template); err != nil {
		return nil, err
	}
	var result Template
	if err := c.postResource("/api/templates", template, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTemplate retrieves a template by its ID.
func (c *koiClient) GetTemplate(id ID) (*Template, error) {
	var template Template
	if err := c.getResource("/api/templates/"+string(id), &template); err != nil {
		return nil, err
	}
	return &template, nil
}

// ListTemplates retrieves a list of templates.
func (c *koiClient) ListTemplates(queryParams ...string) ([]*Template, error) {
	var templates []*Template
	if err := c.listResources("/api/templates", &templates, queryParams...); err != nil {
		return nil, err
	}
	return templates, nil
}

// UpdateTemplate replaces an existing template.
func (c *koiClient) UpdateTemplate(id ID, template *Template) (*Template, error) {
	var result Template
	if err := c.putResource("/api/templates/"+string(id), template, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchTemplate updates an existing template partially.
func (c *koiClient) PatchTemplate(id ID, template *Template) (*Template, error) {
	var result Template
	if err := c.patchResource("/api/templates/"+string(id), template, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteTemplate deletes a template by its ID.
func (c *koiClient) DeleteTemplate(id ID) error {
	return c.deleteResource("/api/templates/" + string(id))
}

// GetUser retrieves a user by its ID.
func (c *koiClient) GetUser(id ID) (*User, error) {
	var user User
	if err := c.getResource("/api/users/"+string(id), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// ListUsers retrieves a list of users.
func (c *koiClient) ListUsers(queryParams ...string) ([]*User, error) {
	var users []*User
	if err := c.listResources("/api/users", &users, queryParams...); err != nil {
		return nil, err
	}
	return users, nil
}

// CreateWish creates a new wish with schema validation.
func (c *koiClient) CreateWish(wish *Wish) (*Wish, error) {
	if err := c.validateWish(wish); err != nil {
		return nil, err
	}
	var result Wish
	if err := c.postResource("/api/wishes", wish, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWish retrieves a wish by its ID.
func (c *koiClient) GetWish(id ID) (*Wish, error) {
	var wish Wish
	if err := c.getResource("/api/wishes/"+string(id), &wish); err != nil {
		return nil, err
	}
	return &wish, nil
}

// ListWishes retrieves a list of wishes.
func (c *koiClient) ListWishes(queryParams ...string) ([]*Wish, error) {
	var wishes []*Wish
	if err := c.listResources("/api/wishes", &wishes, queryParams...); err != nil {
		return nil, err
	}
	return wishes, nil
}

// UpdateWish replaces an existing wish.
func (c *koiClient) UpdateWish(id ID, wish *Wish) (*Wish, error) {
	var result Wish
	if err := c.putResource("/api/wishes/"+string(id), wish, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchWish updates an existing wish partially.
func (c *koiClient) PatchWish(id ID, wish *Wish) (*Wish, error) {
	var result Wish
	if err := c.patchResource("/api/wishes/"+string(id), wish, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteWish deletes a wish by its ID.
func (c *koiClient) DeleteWish(id ID) error {
	return c.deleteResource("/api/wishes/" + string(id))
}

// UploadWishImage uploads an image for a wish.
func (c *koiClient) UploadWishImage(id ID, file []byte) (*Wish, error) {
	var wish Wish
	if err := c.uploadFile("/api/wishes/"+string(id)+"/image", file, "file", &wish); err != nil {
		return nil, err
	}
	return &wish, nil
}

// GetWishWishlist retrieves the wishlist associated with a wish.
func (c *koiClient) GetWishWishlist(id ID) (*Wishlist, error) {
	var wishlist Wishlist
	if err := c.getResource("/api/wishes/"+string(id)+"/wishlist", &wishlist); err != nil {
		return nil, err
	}
	return &wishlist, nil
}

// CreateWishlist creates a new wishlist with schema validation.
func (c *koiClient) CreateWishlist(wishlist *Wishlist) (*Wishlist, error) {
	if err := c.validateWishlist(wishlist); err != nil {
		return nil, err
	}
	var result Wishlist
	if err := c.postResource("/api/wishlists", wishlist, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWishlist retrieves a wishlist by its ID.
func (c *koiClient) GetWishlist(id ID) (*Wishlist, error) {
	var wishlist Wishlist
	if err := c.getResource("/api/wishlists/"+string(id), &wishlist); err != nil {
		return nil, err
	}
	return &wishlist, nil
}

// ListWishlists retrieves a list of wishlists.
func (c *koiClient) ListWishlists(queryParams ...string) ([]*Wishlist, error) {
	var wishlists []*Wishlist
	if err := c.listResources("/api/wishlists", &wishlists, queryParams...); err != nil {
		return nil, err
	}
	return wishlists, nil
}

// UpdateWishlist replaces an existing wishlist.
func (c *koiClient) UpdateWishlist(id ID, wishlist *Wishlist) (*Wishlist, error) {
	var result Wishlist
	if err := c.putResource("/api/wishlists/"+string(id), wishlist, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchWishlist updates an existing wishlist partially.
func (c *koiClient) PatchWishlist(id ID, wishlist *Wishlist) (*Wishlist, error) {
	var result Wishlist
	if err := c.patchResource("/api/wishlists/"+string(id), wishlist, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteWishlist deletes a wishlist by its ID.
func (c *koiClient) DeleteWishlist(id ID) error {
	return c.deleteResource("/api/wishlists/" + string(id))
}

// ListWishlistWishes retrieves the wishes in a wishlist.
func (c *koiClient) ListWishlistWishes(id ID, queryParams ...string) ([]*Wish, error) {
	var wishes []*Wish
	path := fmt.Sprintf("/api/wishlists/%s/wishes", id)
	if err := c.listResources(path, &wishes, queryParams...); err != nil {
		return nil, err
	}
	return wishes, nil
}

// ListWishlistChildren retrieves the children of a wishlist.
func (c *koiClient) ListWishlistChildren(id ID, queryParams ...string) ([]*Wishlist, error) {
	var wishlists []*Wishlist
	path := fmt.Sprintf("/api/wishlists/%s/children", id)
	if err := c.listResources(path, &wishlists, queryParams...); err != nil {
		return nil, err
	}
	return wishlists, nil
}

// UploadWishlistImage uploads an image for a wishlist.
func (c *koiClient) UploadWishlistImage(id ID, file []byte) (*Wishlist, error) {
	var wishlist Wishlist
	if err := c.uploadFile("/api/wishlists/"+string(id)+"/image", file, "file", &wishlist); err != nil {
		return nil, err
	}
	return &wishlist, nil
}

// GetWishlistParent retrieves the parent wishlist of a wishlist.
func (c *koiClient) GetWishlistParent(id ID) (*Wishlist, error) {
	var wishlist Wishlist
	if err := c.getResource("/api/wishlists/"+string(id)+"/parent", &wishlist); err != nil {
		return nil, err
	}
	return &wishlist, nil
}
