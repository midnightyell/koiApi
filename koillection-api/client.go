package koiApi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//File Uploads: Assumes a file field for multipart/form-data uploads. If the API expects a different field name (e.g., fileImage for Datum uploads), the uploadFile method can be updated.
// Pagination: Assumes the page query parameter is sufficient for pagination. If the API uses other parameters (e.g., per_page), the listResources method can be extended.

// Errors for common HTTP status codes.
var (
	ErrInvalidInput  = errors.New("invalid input")
	ErrNotFound      = errors.New("resource not found")
	ErrUnprocessable = errors.New("unprocessable entity")
	ErrUnauthorized  = errors.New("unauthorized")
)

// httpClient implements the Client interface using net/http.
type httpClient struct {
	baseURL    string
	httpClient *http.Client
	token      string
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

// NewHTTPClient creates a new HTTP client for the Koillection API.
func NewHTTPClient(baseURL string, timeout time.Duration) Client {

	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println("Error creating cookie jar:", err)
		return nil
	}
	return &httpClient{
		baseURL: strings.TrimSuffix(baseURL, "/"),
		httpClient: &http.Client{
			Jar:     jar,
			Timeout: timeout,
		},
	}
}

// doRequest sends an HTTP request and returns the response body.
func (c *httpClient) doRequest(ctx context.Context, method, path string, body io.Reader, isMultipart bool) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, body)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	//if c.token != "" && path != "/api/authentication_token" {
	//   Not needed because we use cookiejar
	//   req.Header.Set("Authorization", "Bearer "+c.token)
	//}

	if body != nil && !isMultipart {
		req.Header.Set("Content-Type", "application/ld+json")
	} else if isMultipart {
		req.Header.Set("Content-Type", "multipart/form-data")
	}
	if path == "/api/metrics" {
		req.Header.Set("Accept", "text/plain")
	} else {
		req.Header.Set("Accept", "application/ld+json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}

	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusNoContent:
		return resp, nil
	case http.StatusBadRequest:
		resp.Body.Close()
		return nil, ErrInvalidInput
	case http.StatusUnauthorized:
		resp.Body.Close()
		return nil, ErrUnauthorized
	case http.StatusNotFound:
		resp.Body.Close()
		return nil, ErrNotFound
	case http.StatusUnprocessableEntity:
		resp.Body.Close()
		return nil, ErrUnprocessable
	default:
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(body))
	}
}

// getResource retrieves a single resource and decodes it into the provided struct.
func (c *httpClient) getResource(ctx context.Context, path string, out interface{}) error {
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil, false)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(out)
}

// listResources retrieves a paginated list of resources and decodes the member array.
func (c *httpClient) listResources(ctx context.Context, path string, page int, out interface{}) error {
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return fmt.Errorf("parsing URL: %w", err)
	}
	q := u.Query()
	if page > 0 {
		q.Set("page", strconv.Itoa(page))
	}
	u.RawQuery = q.Encode()

	resp, err := c.doRequest(ctx, http.MethodGet, u.Path+"?"+u.RawQuery, nil, false)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Handle JSON-LD response with "member" array.
	headerContent := resp.Header.Get("Content-Type")
	if strings.Contains(headerContent, "application/ld+json") {
		var wrapper struct {
			Member json.RawMessage `json:"member"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
			return fmt.Errorf("decoding response: %w", err)
		}
		return json.Unmarshal(wrapper.Member, out)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bodyBytes, out)

}

// postResource creates a resource and decodes the response into the provided struct.
func (c *httpClient) postResource(ctx context.Context, path string, in, out interface{}) error {
	body, err := json.Marshal(in)
	if err != nil {
		return fmt.Errorf("encoding request body: %w", err)
	}

	resp, err := c.doRequest(ctx, http.MethodPost, path, bytes.NewReader(body), false)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if out != nil {
		return json.NewDecoder(resp.Body).Decode(out)
	}
	return nil
}

// putResource updates a resource and decodes the response into the provided struct.
func (c *httpClient) putResource(ctx context.Context, path string, in, out interface{}) error {
	body, err := json.Marshal(in)
	if err != nil {
		return fmt.Errorf("encoding request body: %w", err)
	}

	resp, err := c.doRequest(ctx, http.MethodPut, path, bytes.NewReader(body), false)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(out)
}

// patchResource partially updates a resource and decodes the response into the provided struct.
func (c *httpClient) patchResource(ctx context.Context, path string, in, out interface{}) error {
	body, err := json.Marshal(in)
	if err != nil {
		return fmt.Errorf("encoding request body: %w", err)
	}

	resp, err := c.doRequest(ctx, http.MethodPatch, path, bytes.NewReader(body), false)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(out)
}

// deleteResource deletes a resource.
func (c *httpClient) deleteResource(ctx context.Context, path string) error {
	resp, err := c.doRequest(ctx, http.MethodDelete, path, nil, false)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

// uploadFile uploads a file using multipart/form-data and decodes the response.
func (c *httpClient) uploadFile(ctx context.Context, path string, file []byte, out interface{}) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "file")
	if err != nil {
		return fmt.Errorf("creating form file: %w", err)
	}
	if _, err := part.Write(file); err != nil {
		return fmt.Errorf("writing file: %w", err)
	}
	if err := writer.Close(); err != nil {
		return fmt.Errorf("closing writer: %w", err)
	}

	resp, err := c.doRequest(ctx, http.MethodPost, path, body, true)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(out)
}

// CheckLogin authenticates a user and returns a JWT token.
func (c *httpClient) CheckLogin(ctx context.Context, username, password string) (string, error) {
	reqBody := map[string]string{
		"username": username,
		"password": password,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("encoding request body: %w", err)
	}

	resp, err := c.doRequest(ctx, http.MethodPost, "/api/authentication_token", bytes.NewReader(body), false)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Token string `json:"token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("decoding response: %w", err)
	}

	c.token = result.Token
	return result.Token, nil
}

// GetMetrics retrieves system or user-specific metrics.
func (c *httpClient) GetMetrics(ctx context.Context) (*Metrics, error) {
	var metrics Metrics
	if err := c.getResource(ctx, "/api/metrics", &metrics); err != nil {
		return nil, err
	}
	return &metrics, nil
}

// CreateAlbum creates a new album.
func (c *httpClient) CreateAlbum(ctx context.Context, album *Album) (*Album, error) {
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
func (c *httpClient) ListAlbums(ctx context.Context, page int) ([]*Album, error) {
	var albums []*Album
	if err := c.listResources(ctx, "/api/albums", page, &albums); err != nil {
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

// ListAlbumChildren retrieves child albums of an album.
func (c *httpClient) ListAlbumChildren(ctx context.Context, id ID, page int) ([]*Album, error) {
	var albums []*Album
	if err := c.listResources(ctx, "/api/albums/"+string(id)+"/children", page, &albums); err != nil {
		return nil, err
	}
	return albums, nil
}

// UploadAlbumImage uploads an image for an album.
func (c *httpClient) UploadAlbumImage(ctx context.Context, id ID, file []byte) (*Album, error) {
	var album Album
	if err := c.uploadFile(ctx, "/api/albums/"+string(id)+"/image", file, &album); err != nil {
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

// ListAlbumPhotos retrieves photos in an album.
func (c *httpClient) ListAlbumPhotos(ctx context.Context, id ID, page int) ([]*Photo, error) {
	var photos []*Photo
	if err := c.listResources(ctx, "/api/albums/"+string(id)+"/photos", page, &photos); err != nil {
		return nil, err
	}
	return photos, nil
}

// CreateChoiceList creates a new choice list.
func (c *httpClient) CreateChoiceList(ctx context.Context, choiceList *ChoiceList) (*ChoiceList, error) {
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
func (c *httpClient) ListChoiceLists(ctx context.Context, page int) ([]*ChoiceList, error) {
	var choiceLists []*ChoiceList
	if err := c.listResources(ctx, "/api/choice_lists", page, &choiceLists); err != nil {
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

// CreateCollection creates a new collection.
func (c *httpClient) CreateCollection(ctx context.Context, collection *Collection) (*Collection, error) {
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
func (c *httpClient) ListCollections(ctx context.Context, page int) ([]*Collection, error) {
	var collections []*Collection
	if err := c.listResources(ctx, "/api/collections", page, &collections); err != nil {
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

// ListCollectionChildren retrieves child collections of a collection.
func (c *httpClient) ListCollectionChildren(ctx context.Context, id ID, page int) ([]*Collection, error) {
	var collections []*Collection
	if err := c.listResources(ctx, "/api/collections/"+string(id)+"/children", page, &collections); err != nil {
		return nil, err
	}
	return collections, nil
}

// UploadCollectionImage uploads an image for a collection.
func (c *httpClient) UploadCollectionImage(ctx context.Context, id ID, file []byte) (*Collection, error) {
	var collection Collection
	if err := c.uploadFile(ctx, "/api/collections/"+string(id)+"/image", file, &collection); err != nil {
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

// ListCollectionItems retrieves items in a collection.
func (c *httpClient) ListCollectionItems(ctx context.Context, id ID, page int) ([]*Item, error) {
	var items []*Item
	if err := c.listResources(ctx, "/api/collections/"+string(id)+"/items", page, &items); err != nil {
		return nil, err
	}
	return items, nil
}

// ListCollectionData retrieves data fields in a collection.
func (c *httpClient) ListCollectionData(ctx context.Context, id ID, page int) ([]*Datum, error) {
	var data []*Datum
	if err := c.listResources(ctx, "/api/collections/"+string(id)+"/data", page, &data); err != nil {
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

// CreateDatum creates a new datum.
func (c *httpClient) CreateDatum(ctx context.Context, datum *Datum) (*Datum, error) {
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

// ListData retrieves a list of data fields.
func (c *httpClient) ListData(ctx context.Context, page int) ([]*Datum, error) {
	var data []*Datum
	if err := c.listResources(ctx, "/api/data", page, &data); err != nil {
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
	if err := c.uploadFile(ctx, "/api/data/"+string(id)+"/file", file, &datum); err != nil {
		return nil, err
	}
	return &datum, nil
}

// UploadDatumImage uploads an image for a datum.
func (c *httpClient) UploadDatumImage(ctx context.Context, id ID, image []byte) (*Datum, error) {
	var datum Datum
	if err := c.uploadFile(ctx, "/api/data/"+string(id)+"/image", image, &datum); err != nil {
		return nil, err
	}
	return &datum, nil
}

// UploadDatumVideo uploads a video for a datum.
func (c *httpClient) UploadDatumVideo(ctx context.Context, id ID, video []byte) (*Datum, error) {
	var datum Datum
	if err := c.uploadFile(ctx, "/api/data/"+string(id)+"/video", video, &datum); err != nil {
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

// CreateField creates a new field.
func (c *httpClient) CreateField(ctx context.Context, field *Field) (*Field, error) {
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
func (c *httpClient) ListFields(ctx context.Context, page int) ([]*Field, error) {
	var fields []*Field
	if err := c.listResources(ctx, "/api/fields", page, &fields); err != nil {
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

// ListTemplateFields retrieves fields associated with a template.
func (c *httpClient) ListTemplateFields(ctx context.Context, templateid ID, page int) ([]*Field, error) {
	var fields []*Field
	if err := c.listResources(ctx, "/api/templates/"+string(templateid)+"/fields", page, &fields); err != nil {
		return nil, err
	}
	return fields, nil
}

// ListInventories retrieves a list of inventories.
func (c *httpClient) ListInventories(ctx context.Context, page int) ([]*Inventory, error) {
	var inventories []*Inventory
	if err := c.listResources(ctx, "/api/inventories", page, &inventories); err != nil {
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
func (c *httpClient) ListItems(ctx context.Context, page int) ([]*Item, error) {
	var items []*Item
	if err := c.listResources(ctx, "/api/items", page, &items); err != nil {
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
	if err := c.uploadFile(ctx, "/api/items/"+string(id)+"/image", file, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// ListItemRelatedItems retrieves related items for an item.
func (c *httpClient) ListItemRelatedItems(ctx context.Context, id ID, page int) ([]*Item, error) {
	var items []*Item
	if err := c.listResources(ctx, "/api/items/"+string(id)+"/related_items", page, &items); err != nil {
		return nil, err
	}
	return items, nil
}

// ListItemLoans retrieves loans for an item.
func (c *httpClient) ListItemLoans(ctx context.Context, id ID, page int) ([]*Loan, error) {
	var loans []*Loan
	if err := c.listResources(ctx, "/api/items/"+string(id)+"/loans", page, &loans); err != nil {
		return nil, err
	}
	return loans, nil
}

// ListItemTags retrieves tags for an item.
func (c *httpClient) ListItemTags(ctx context.Context, id ID, page int) ([]*Tag, error) {
	var tags []*Tag
	if err := c.listResources(ctx, "/api/items/"+string(id)+"/tags", page, &tags); err != nil {
		return nil, err
	}
	return tags, nil
}

// ListItemData retrieves data fields for an item.
func (c *httpClient) ListItemData(ctx context.Context, id ID, page int) ([]*Datum, error) {
	var data []*Datum
	if err := c.listResources(ctx, "/api/items/"+string(id)+"/data", page, &data); err != nil {
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

// CreateLoan creates a new loan.
func (c *httpClient) CreateLoan(ctx context.Context, loan *Loan) (*Loan, error) {
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
func (c *httpClient) ListLoans(ctx context.Context, page int) ([]*Loan, error) {
	var loans []*Loan
	if err := c.listResources(ctx, "/api/loans", page, &loans); err != nil {
		return nil, err
	}
	return loans, nil
}

// UpdateLoan replaces an existing loan.
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
func (c *httpClient) ListLogs(ctx context.Context, page int) ([]*Log, error) {
	var logs []*Log
	if err := c.listResources(ctx, "/api/logs", page, &logs); err != nil {
		return nil, err
	}
	return logs, nil
}

// CreatePhoto creates a new photo.
func (c *httpClient) CreatePhoto(ctx context.Context, photo *Photo) (*Photo, error) {
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
func (c *httpClient) ListPhotos(ctx context.Context, page int) ([]*Photo, error) {
	var photos []*Photo
	if err := c.listResources(ctx, "/api/photos", page, &photos); err != nil {
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
	if err := c.uploadFile(ctx, "/api/photos/"+string(id)+"/image", file, &photo); err != nil {
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

// CreateTag creates a new tag.
func (c *httpClient) CreateTag(ctx context.Context, tag *Tag) (*Tag, error) {
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
func (c *httpClient) ListTags(ctx context.Context, page int) ([]*Tag, error) {
	var tags []*Tag
	if err := c.listResources(ctx, "/api/tags", page, &tags); err != nil {
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
	if err := c.uploadFile(ctx, "/api/tags/"+string(id)+"/image", file, &tag); err != nil {
		return nil, err
	}
	return &tag, nil
}

// ListTagItems retrieves items associated with a tag.
func (c *httpClient) ListTagItems(ctx context.Context, id ID, page int) ([]*Item, error) {
	var items []*Item
	if err := c.listResources(ctx, "/api/tags/"+string(id)+"/items", page, &items); err != nil {
		return nil, err
	}
	return items, nil
}

// GetTagsCategory retrieves the category associated with a tag.
func (c *httpClient) GetCategoryOfTag(ctx context.Context, id ID) (*TagCategory, error) {
	var category TagCategory
	if err := c.getResource(ctx, "/api/tags/"+string(id)+"/category", &category); err != nil {
		return nil, err
	}
	return &category, nil
}

// CreateTagCategory creates a new tag category.
func (c *httpClient) CreateTagCategory(ctx context.Context, category *TagCategory) (*TagCategory, error) {
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
func (c *httpClient) ListTagCategories(ctx context.Context, page int) ([]*TagCategory, error) {
	var categories []*TagCategory
	if err := c.listResources(ctx, "/api/tag_categories", page, &categories); err != nil {
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

// ListTagCategoryTags retrieves tags in a tag category.
func (c *httpClient) ListTagCategoryTags(ctx context.Context, id ID, page int) ([]*Tag, error) {
	var tags []*Tag
	if err := c.listResources(ctx, "/api/tag_categories/"+string(id)+"/tags", page, &tags); err != nil {
		return nil, err
	}
	return tags, nil
}

// CreateTemplate creates a new template.
func (c *httpClient) CreateTemplate(ctx context.Context, template *Template) (*Template, error) {
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
func (c *httpClient) ListTemplates(ctx context.Context, page int) ([]*Template, error) {
	var templates []*Template
	if err := c.listResources(ctx, "/api/templates", page, &templates); err != nil {
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
func (c *httpClient) ListUsers(ctx context.Context, page int) ([]*User, error) {
	var users []*User
	if err := c.listResources(ctx, "/api/users", page, &users); err != nil {
		return nil, err
	}
	return users, nil
}

// CreateWish creates a new wish.
func (c *httpClient) CreateWish(ctx context.Context, wish *Wish) (*Wish, error) {
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
func (c *httpClient) ListWishes(ctx context.Context, page int) ([]*Wish, error) {
	var wishes []*Wish
	if err := c.listResources(ctx, "/api/wishes", page, &wishes); err != nil {
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
	if err := c.uploadFile(ctx, "/api/wishes/"+string(id)+"/image", file, &wish); err != nil {
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

// CreateWishlist creates a new wishlist.
func (c *httpClient) CreateWishlist(ctx context.Context, wishlist *Wishlist) (*Wishlist, error) {
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
func (c *httpClient) ListWishlists(ctx context.Context, page int) ([]*Wishlist, error) {
	var wishlists []*Wishlist
	if err := c.listResources(ctx, "/api/wishlists", page, &wishlists); err != nil {
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

// ListWishlistWishes retrieves wishes in a wishlist.
func (c *httpClient) ListWishlistWishes(ctx context.Context, id ID, page int) ([]*Wish, error) {
	var wishes []*Wish
	if err := c.listResources(ctx, "/api/wishlists/"+string(id)+"/wishes", page, &wishes); err != nil {
		return nil, err
	}
	return wishes, nil
}

// ListWishlistChildren retrieves child wishlists of a wishlist.
func (c *httpClient) ListWishlistChildren(ctx context.Context, id ID, page int) ([]*Wishlist, error) {
	var wishlists []*Wishlist
	if err := c.listResources(ctx, "/api/wishlists/"+string(id)+"/children", page, &wishlists); err != nil {
		return nil, err
	}
	return wishlists, nil
}

// UploadWishlistImage uploads an image for a wishlist.
func (c *httpClient) UploadWishlistImage(ctx context.Context, id ID, file []byte) (*Wishlist, error) {
	var wishlist Wishlist
	if err := c.uploadFile(ctx, "/api/wishlists/"+string(id)+"/image", file, &wishlist); err != nil {
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
