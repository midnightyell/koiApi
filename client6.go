package koillection

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Errors for common HTTP status codes and input validation.
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

// NewHTTPClient creates a new HTTP client for the Koillection API.
func NewHTTPClient(baseURL string, timeout time.Duration) Client {
	return &httpClient{
		baseURL: strings.TrimSuffix(baseURL, "/"),
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// getID extracts and validates the ID from an object pointer.
func (c *httpClient) getID(obj interface{}) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("%w: object is nil", ErrInvalidInput)
	}
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return "", fmt.Errorf("%w: object must be a non-nil pointer, got %T", ErrInvalidInput, obj)
	}
	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return "", fmt.Errorf("%w: object must be a struct, got %T", ErrInvalidInput, obj)
	}
	idField := v.FieldByName("ID")
	if !idField.IsValid() || idField.Kind() != reflect.String {
		return "", fmt.Errorf("%w: object %T has no valid ID field", ErrInvalidInput, obj)
	}
	id := idField.String()
	if id == "" {
		return "", fmt.Errorf("%w: object %T has empty ID", ErrInvalidInput, obj)
	}
	return id, nil
}

// doRequest sends an HTTP request and returns the response body.
func (c *httpClient) doRequest(ctx context.Context, method, path string, body io.Reader, isMultipart bool) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, body)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	if c.token != "" && path != "/api/authentication_token" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	if body != nil && !isMultipart {
		req.Header.Set("Content-Type", "application/json")
	} else if isMultipart {
		req.Header.Set("Content-Type", "multipart/form-data")
	}
	req.Header.Set("Accept", "application/json")

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
	var wrapper struct {
		Member json.RawMessage `json:"member"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}

	return json.Unmarshal(wrapper.Member, out)
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

// GetAlbum retrieves an album.
func (c *httpClient) GetAlbum(ctx context.Context, album *Album) (*Album, error) {
	id, err := c.getID(album)
	if err != nil {
		return nil, err
	}
	var result Album
	if err := c.getResource(ctx, "/api/albums/"+id, &result); err != nil {
		return nil, err
	}
	return &result, nil
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
func (c *httpClient) UpdateAlbum(ctx context.Context, album *Album) (*Album, error) {
	id, err := c.getID(album)
	if err != nil {
		return nil, err
	}
	var result Album
	if err := c.putResource(ctx, "/api/albums/"+id, album, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchAlbum updates an existing album partially.
func (c *httpClient) PatchAlbum(ctx context.Context, album *Album) (*Album, error) {
	id, err := c.getID(album)
	if err != nil {
		return nil, err
	}
	var result Album
	if err := c.patchResource(ctx, "/api/albums/"+id, album, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteAlbum deletes an album.
func (c *httpClient) DeleteAlbum(ctx context.Context, album *Album) error {
	id, err := c.getID(album)
	if err != nil {
		return err
	}
	return c.deleteResource(ctx, "/api/albums/"+id)
}

// ListAlbumChildren retrieves child albums of an album.
func (c *httpClient) ListAlbumChildren(ctx context.Context, album *Album, page int) ([]*Album, error) {
	id, err := c.getID(album)
	if err != nil {
		return nil, err
	}
	var albums []*Album
	if err := c.listResources(ctx, "/api/albums/"+id+"/children", page, &albums); err != nil {
		return nil, err
	}
	return albums, nil
}

// UploadAlbumImage uploads an image for an album.
func (c *httpClient) UploadAlbumImage(ctx context.Context, album *Album, file []byte) (*Album, error) {
	id, err := c.getID(album)
	if err != nil {
		return nil, err
	}
	var result Album
	if err := c.uploadFile(ctx, "/api/albums/"+id+"/image", file, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAlbumParent retrieves the parent album of an album.
func (c *httpClient) GetAlbumParent(ctx context.Context, album *Album) (*Album, error) {
	id, err := c.getID(album)
	if err != nil {
		return nil, err
	}
	var result Album
	if err := c.getResource(ctx, "/api/albums/"+id+"/parent", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ListAlbumPhotos retrieves photos in an album.
func (c *httpClient) ListAlbumPhotos(ctx context.Context, album *Album, page int) ([]*Photo, error) {
	id, err := c.getID(album)
	if err != nil {
		return nil, err
	}
	var photos []*Photo
	if err := c.listResources(ctx, "/api/albums/"+id+"/photos", page, &photos); err != nil {
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

// GetChoiceList retrieves a choice list.
func (c *httpClient) GetChoiceList(ctx context.Context, choiceList *ChoiceList) (*ChoiceList, error) {
	id, err := c.getID(choiceList)
	if err != nil {
		return nil, err
	}
	var result ChoiceList
	if err := c.getResource(ctx, "/api/choice_lists/"+id, &result); err != nil {
		return nil, err
	}
	return &result, nil
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
func (c *httpClient) UpdateChoiceList(ctx context.Context, choiceList *ChoiceList) (*ChoiceList, error) {
	id, err := c.getID(choiceList)
	if err != nil {
		return nil, err
	}
	var result ChoiceList
	if err := c.putResource(ctx, "/api/choice_lists/"+id, choiceList, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchChoiceList updates an existing choice list partially.
func (c *httpClient) PatchChoiceList(ctx context.Context, choiceList *ChoiceList) (*ChoiceList, error) {
	id, err := c.getID(choiceList)
	if err != nil {
		return nil, err
	}
	var result ChoiceList
	if err := c.patchResource(ctx, "/api/choice_lists/"+id, choiceList, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteChoiceList deletes a choice list.
func (c *httpClient) DeleteChoiceList(ctx context.Context, choiceList *ChoiceList) error {
	id, err := c.getID(choiceList)
	if err != nil {
		return err
	}
	return c.deleteResource(ctx, "/api/choice_lists/"+id)
}

// CreateCollection creates a new collection.
func (c *httpClient) CreateCollection(ctx context.Context, collection *Collection) (*Collection, error) {
	var result Collection
	if err := c.postResource(ctx, "/api/collections", collection, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCollection retrieves a collection.
func (c *httpClient) GetCollection(ctx context.Context, collection *Collection) (*Collection, error) {
	id, err := c.getID(collection)
	if err != nil {
		return nil, err
	}
	var result Collection
	if err := c.getResource(ctx, "/api/collections/"+id, &result); err != nil {
		return nil, err
	}
	return &result, nil
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
func (c *httpClient) UpdateCollection(ctx context.Context, collection *Collection) (*Collection, error) {
	id, err := c.getID(collection)
	if err != nil {
		return nil, err
	}
	var result Collection
	if err := c.putResource(ctx, "/api/collections/"+id, collection, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchCollection updates an existing collection partially.
func (c *httpClient) PatchCollection(ctx context.Context, collection *Collection) (*Collection, error) {
	id, err := c.getID(collection)
	if err != nil {
		return nil, err
	}
	var result Collection
	if err := c.patchResource(ctx, "/api/collections/"+id, collection, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteCollection deletes a collection.
func (c *httpClient) DeleteCollection(ctx context.Context, collection *Collection) error {
	id, err := c.getID(collection)
	if err != nil {
		return err
	}
	return c.deleteResource(ctx, "/api/collections/"+id)
}

// ListCollectionChildren retrieves child collections of a collection.
func (c *httpClient) ListCollectionChildren(ctx context.Context, collection *Collection, page int) ([]*Collection, error) {
	id, err := c.getID(collection)
	if err != nil {
		return nil, err
	}
	var collections []*Collection
	if err := c.listResources(ctx, "/api/collections/"+id+"/children", page, &collections); err != nil {
		return nil, err
	}
	return collections, nil
}

// UploadCollectionImage uploads an image for a collection.
func (c *httpClient) UploadCollectionImage(ctx context.Context, collection *Collection, file []byte) (*Collection, error) {
	id, err := c.getID(collection)
	if err != nil {
		return nil, err
	}
	var result Collection
	if err := c.uploadFile(ctx, "/api/collections/"+id+"/image", file, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCollectionParent retrieves the parent collection of a collection.
func (c *httpClient) GetCollectionParent(ctx context.Context, collection *Collection) (*Collection, error) {
	id, err := c.getID(collection)
	if err != nil {
		return nil, err
	}
	var result Collection
	if err := c.getResource(ctx, "/api/collections/"+id+"/parent", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ListCollectionItems retrieves items in a collection.
func (c *httpClient) ListCollectionItems(ctx context.Context, collection *Collection, page int) ([]*Item, error) {
	id, err := c.getID(collection)
	if err != nil {
		return nil, err
	}
	var items []*Item
	if err := c.listResources(ctx, "/api/collections/"+id+"/items", page, &items); err != nil {
		return nil, err
	}
	return items, nil
}

// ListCollectionData retrieves data fields in a collection.
func (c *httpClient) ListCollectionData(ctx context.Context, collection *Collection, page int) ([]*Datum, error) {
	id, err := c.getID(collection)
	if err != nil {
		return nil, err
	}
	var data []*Datum
	if err := c.listResources(ctx, "/api/collections/"+id+"/data", page, &data); err != nil {
		return nil, err
	}
	return data, nil
}

// GetCollectionDefaultTemplate retrieves the default template for items in a collection.
func (c *httpClient) GetCollectionDefaultTemplate(ctx context.Context, collection *Collection) (*Template, error) {
	id, err := c.getID(collection)
	if err != nil {
		return nil, err
	}
	var result Template
	if err := c.getResource(ctx, "/api/collections/"+id+"/items_default_template", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateDatum creates a new datum.
func (c *httpClient) CreateDatum(ctx context.Context, datum *Datum) (*Datum, error) {
	var result Datum
	if err := c.postResource(ctx, "/api/data", datum, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDatum retrieves a datum.
func (c *httpClient) GetDatum(ctx context.Context, datum *Datum) (*Datum, error) {
	id, err := c.getID(datum)
	if err != nil {
		return nil, err
	}
	var result Datum
	if err := c.getResource(ctx, "/api/data/"+id, &result); err != nil {
		return nil, err
	}
	return &result, nil
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
func (c *httpClient) UpdateDatum(ctx context.Context, datum *Datum) (*Datum, error) {
	id, err := c.getID(datum)
	if err != nil {
		return nil, err
	}
	var result Datum
	if err := c.putResource(ctx, "/api/data/"+id, datum, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchDatum updates an existing datum partially.
func (c *httpClient) PatchDatum(ctx context.Context, datum *Datum) (*Datum, error) {
	id, err := c.getID(datum)
	if err != nil {
		return nil, err
	}
	var result Datum
	if err := c.patchResource(ctx, "/api/data/"+id, datum, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteDatum deletes a datum.
func (c *httpClient) DeleteDatum(ctx context.Context, datum *Datum) error {
	id, err := c.getID(datum)
	if err != nil {
		return err
	}
	return c.deleteResource(ctx, "/api/data/"+id)
}

// UploadDatumFile uploads a file for a datum.
func (c *httpClient) UploadDatumFile(ctx context.Context, datum *Datum, file []byte) (*Datum, error) {
	id, err := c.getID(datum)
	if err != nil {
		return nil, err
	}
	var result Datum
	if err := c.uploadFile(ctx, "/api/data/"+id+"/file", file, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UploadDatumImage uploads an image for a datum.
func (c *httpClient) UploadDatumImage(ctx context.Context, datum *Datum, image []byte) (*Datum, error) {
	id, err := c.getID(datum)
	if err != nil {
		return nil, err
	}
	var result Datum
	if err := c.uploadFile(ctx, "/api/data/"+id+"/image", image, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UploadDatumVideo uploads a video for a datum.
func (c *httpClient) UploadDatumVideo(ctx context.Context, datum *Datum, video []byte) (*Datum, error) {
	id, err := c.getID(datum)
	if err != nil {
		return nil, err
	}
	var result Datum
	if err := c.uploadFile(ctx, "/api/data/"+id+"/video", video, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDatumItem retrieves the item associated with a datum.
func (c *httpClient) GetDatumItem(ctx context.Context, datum *Datum) (*Item, error) {
	id, err := c.getID(datum)
	if err != nil {
		return nil, err
	}
	var result Item
	if err := c.getResource(ctx, "/api/data/"+id+"/item", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDatumCollection retrieves the collection associated with a datum.
func (c *httpClient) GetDatumCollection(ctx context.Context, datum *Datum) (*Collection, error) {
	id, err := c.getID(datum)
	if err != nil {
		return nil, err
	}
	var result Collection
	if err := c.getResource(ctx, "/api/data/"+id+"/collection", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateField creates a new field.
func (c *httpClient) CreateField(ctx context.Context, field *Field) (*Field, error) {
	var result Field
	if err := c.postResource(ctx, "/api/fields", field, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetField retrieves a field.
func (c *httpClient) GetField(ctx context.Context, field *Field) (*Field, error) {
	id, err := c.getID(field)
	if err != nil {
		return nil, err
	}
	var result Field
	if err := c.getResource(ctx, "/api/fields/"+id, &result); err != nil {
		return nil, err
	}
	return &result, nil
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
func (c *httpClient) UpdateField(ctx context.Context, field *Field) (*Field, error) {
	id, err := c.getID(field)
	if err != nil {
		return nil, err
	}
	var result Field
	if err := c.putResource(ctx, "/api/fields/"+id, field, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchField updates an existing field partially.
func (c *httpClient) PatchField(ctx context.Context, field *Field) (*Field, error) {
	id, err := c.getID(field)
	if err != nil {
		return nil, err
	}
	var result Field
	if err := c.patchResource(ctx, "/api/fields/"+id, field, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteField deletes a field.
func (c *httpClient) DeleteField(ctx context.Context, field *Field) error {
	id, err := c.getID(field)
	if err != nil {
		return err
	}
	return c.deleteResource(ctx, "/api/fields/"+id)
}

// GetFieldTemplate retrieves the template associated with a field.
func (c *httpClient) GetFieldTemplate(ctx context.Context, field *Field) (*Template, error) {
	id, err := c.getID(field)
	if err != nil {
		return nil, err
	}
	var result Template
	if err := c.getResource(ctx, "/api/fields/"+id+"/template", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ListTemplateFields retrieves fields associated with a template.
func (c *httpClient) ListTemplateFields(ctx context.Context, template *Template, page int) ([]*Field, error) {
	id, err := c.getID(template)
	if err != nil {
		return nil, err
	}
	var fields []*Field
	if err := c.listResources(ctx, "/api/templates/"+id+"/fields", page, &fields); err != nil {
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

// GetInventory retrieves an inventory.
func (c *httpClient) GetInventory(ctx context.Context, inventory *Inventory) (*Inventory, error) {
	id, err := c.getID(inventory)
	if err != nil {
		return nil, err
	}
	var result Inventory
	if err := c.getResource(ctx, "/api/inventories/"+id, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteInventory deletes an inventory.
func (c *httpClient) DeleteInventory(ctx context.Context, inventory *Inventory) error {
	id, err := c.getID(inventory)
	if err != nil {
		return err
	}
	return c.deleteResource(ctx, "/api/inventories/"+id)
}

// CreateItem creates a new item.
func (c *httpClient) CreateItem(ctx context.Context, item *Item) (*Item, error) {
	var result Item
	if err := c.postResource(ctx, "/api/items", item, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetItem retrieves an item.
func (c *httpClient) GetItem(ctx context.Context, item *Item) (*Item, error) {
	id, err := c.getID(item)
	if err != nil {
		return nil, err
	}
	var result Item
	if err := c.getResource(ctx, "/api/items/"+id, &result); err != nil {
		return nil, err
	}
	return &result, nil
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
func (c *httpClient) UpdateItem(ctx context.Context, item *Item) (*Item, error) {
	id, err := c.getID(item)
	if err != nil {
		return nil, err
	}
	var result Item
	if err := c.putResource(ctx, "/api/items/"+id, item, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchItem updates an existing item partially.
func (c *httpClient) PatchItem(ctx context.Context, item *Item) (*Item, error) {
	id, err := c.getID(item)
	if err != nil {
		return nil, err
	}
	var result Item
	if err := c.patchResource(ctx, "/api/items/"+id, item, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteItem deletes an item.
func (c *httpClient) DeleteItem(ctx context.Context, item *Item) error {
	id, err := c.getID(item)
	if err != nil {
		return err
	}
	return c.deleteResource(ctx, "/api/items/"+id)
}

// UploadItemImage uploads an image for an item.
func (c *httpClient) UploadItemImage(ctx context.Context, item *Item, file []byte) (*Item, error) {
	id, err := c.getID(item)
	if err != nil {
		return nil, err
	}
	var result Item
	if err := c.uploadFile(ctx, "/api/items/"+id+"/image", file, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ListItemRelatedItems retrieves related items for an item.
func (c *httpClient) ListItemRelatedItems(ctx context.Context, item *Item, page int) ([]*Item, error) {
	id, err := c.getID(item)
	if err != nil {
		return nil, err
	}
	var items []*Item
	if err := c.listResources(ctx, "/api/items/"+id+"/related_items", page, &items); err != nil {
		return nil, err
	}
	return items, nil
}

// ListItemLoans retrieves loans for an item.
func (c *httpClient) ListItemLoans(ctx context.Context, item *Item, page int) ([]*Loan, error) {
	id, err := c.getID(item)
	if err != nil {
		return nil, err
	}
	var loans []*Loan
	if err := c.listResources(ctx, "/api/items/"+id+"/loans", page, &loans); err != nil {
		return nil, err
	}
	return loans, nil
}

// ListItemTags retrieves tags for an item.
func (c *httpClient) ListItemTags(ctx context.Context, item *Item, page int) ([]*Tag, error) {
	id, err := c.getID(item)
	if err != nil {
		return nil, err
	}
	var tags []*Tag
	if err := c.listResources(ctx, "/api/items/"+id+"/tags", page, &tags); err != nil {
		return nil, err
	}
	return tags, nil
}

// ListItemData retrieves data fields for an item.
func (c *httpClient) ListItemData(ctx context.Context, item *Item, page int) ([]*Datum, error) {
	id, err := c.getID(item)
	if err != nil {
		return nil, err
	}
	var data []*Datum
	if err := c.listResources(ctx, "/api/items/"+id+"/data", page, &data); err != nil {
		return nil, err
	}
	return data, nil
}

// GetItemCollection retrieves the collection associated with an item.
func (c *httpClient) GetItemCollection(ctx context.Context, item *Item) (*Collection, error) {
	id, err := c.getID(item)
	if err != nil {
		return nil, err
	}
	var result Collection
	if err := c.getResource(ctx, "/api/items/"+id+"/collection", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateLoan creates a new loan.
func (c *httpClient) CreateLoan(ctx context.Context, loan *Loan) (*Loan, error) {
	var result Loan
	if err := c.postResource(ctx, "/api/loans", loan, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetLoan retrieves a loan.
func (c *httpClient) GetLoan(ctx context.Context, loan *Loan) (*Loan, error) {
	id, err := c.getID(loan)
	if err != nil {
		return nil, err
	}
	var result Loan
	if err := c.getResource(ctx, "/api/loans/"+id, &result); err != nil {
		return nil, err
	}
	return &result, nil
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
func (c *httpClient) UpdateLoan(ctx context.Context, loan *Loan) (*Loan, error) {
	id, err := c.getID(loan)
	if err != nil {
		return nil, err
	}
	var result Loan
	if err := c.putResource(ctx, "/api/loans/"+id, loan, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchLoan updates an existing loan partially.
func (c *httpClient) PatchLoan(ctx context.Context, loan *Loan) (*Loan, error) {
	id, err := c.getID(loan)
	if err != nil {
		return nil, err
	}
	var result Loan
	if err := c.patchResource(ctx, "/api/loans/"+id, loan, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteLoan deletes a loan.
func (c *httpClient) DeleteLoan(ctx context.Context, loan *Loan) error {
	id, err := c.getID(loan)
	if err != nil {
		return err
	}
	return c.deleteResource(ctx, "/api/loans/"+id)
}

// GetLoanItem retrieves the item associated with a loan.
func (c *httpClient) GetLoanItem(ctx context.Context, loan *Loan) (*Item, error) {
	id, err := c.getID(loan)
	if err != nil {
		return nil, err
	}
	var result Item
	if err := c.getResource(ctx, "/api/loans/"+id+"/item", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetLog retrieves a log.
func (c *httpClient) GetLog(ctx context.Context, log *Log) (*Log, error) {
	id, err := c.getID(log)
	if err != nil {
		return nil, err
	}
	var result Log
	if err := c.getResource(ctx, "/api/logs/"+id, &result); err != nil {
		return nil, err
	}
	return &result, nil
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

// GetPhoto retrieves a photo.
func (c *httpClient) GetPhoto(ctx context.Context, photo *Photo) (*Photo, error) {
	id, err := c.getID(photo)
	if err != nil {
		return nil, err
	}
	var result Photo
	if err := c.getResource(ctx, "/api/photos/"+id, &result); err != nil {
		return nil, err
	}
	return &result, nil
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
func (c *httpClient) UpdatePhoto(ctx context.Context, photo *Photo) (*Photo, error) {
	id, err := c.getID(photo)
	if err != nil {
		return nil, err
	}
	var result Photo
	if err := c.putResource(ctx, "/api/photos/"+id, photo, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchPhoto updates an existing photo partially.
func (c *httpClient) PatchPhoto(ctx context.Context, photo *Photo) (*Photo, error) {
	id, err := c.getID(photo)
	if err != nil {
		return nil, err
	}
	var result Photo
	if err := c.patchResource(ctx, "/api/photos/"+id, photo, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeletePhoto deletes a photo.
func (c *httpClient) DeletePhoto(ctx context.Context, photo *Photo) error {
	id, err := c.getID(photo)
	if err != nil {
		return err
	}
	return c.deleteResource(ctx, "/api/photos/"+id)
}

// UploadPhotoImage uploads an image for a photo.
func (c *httpClient) UploadPhotoImage(ctx context.Context, photo *Photo, file []byte) (*Photo, error) {
	id, err := c.getID(photo)
	if err != nil {
		return nil, err
	}
	var result Photo
	if err := c.uploadFile(ctx, "/api/photos/"+id+"/image", file, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPhotoAlbum retrieves the album associated with a photo.
func (c *httpClient) GetPhotoAlbum(ctx context.Context, photo *Photo) (*Album, error) {
	id, err := c.getID(photo)
	if err != nil {
		return nil, err
	}
	var result Album
	if err := c.getResource(ctx, "/api/photos/"+id+"/album", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateTag creates a new tag.
func (c *httpClient) CreateTag(ctx context.Context, tag *Tag) (*Tag, error) {
	var result Tag
	if err := c.postResource(ctx, "/api/tags", tag, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTag retrieves a tag.
func (c *httpClient) GetTag(ctx context.Context, tag *Tag) (*Tag, error) {
	id, err := c.getID(tag)
	if err != nil {
		return nil, err
	}
	var result Tag
	if err := c.getResource(ctx, "/api/tags/"+id, &result); err != nil {
		return nil, err
	}
	return &result, nil
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
func (c *httpClient) UpdateTag(ctx context.Context, tag *Tag) (*Tag, error) {
	id, err := c.getID(tag)
	if err != nil {
		return nil, err
	}
	var result Tag
	if err := c.putResource(ctx, "/api/tags/"+id, tag, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchTag updates an existing tag partially.
func (c *httpClient) PatchTag(ctx context.Context, tag *Tag) (*Tag, error) {
	id, err := c.getID(tag)
	if err != nil {
		return nil, err
	}
	var result Tag
	if err := c.patchResource(ctx, "/api/tags/"+id, tag, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteTag deletes a tag.
func (c *httpClient) DeleteTag(ctx context.Context, tag *Tag) error {
	id, err := c.getID(tag)
	if err != nil {
		return err
	}
	return c.deleteResource(ctx, "/api/tags/"+id)
}

// UploadTagImage uploads an image for a tag.
func (c *httpClient) UploadTagImage(ctx context.Context, tag *Tag, file []byte) (*Tag, error) {
	id, err := c.getID(tag)
	if err != nil {
		return nil, err
	}
	var result Tag
	if err := c.uploadFile(ctx, "/api/tags/"+id+"/image", file, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ListTagItems retrieves items associated with a tag.
func (c *httpClient) ListTagItems(ctx context.Context, tag *Tag, page int) ([]*Item, error) {
	id, err := c.getID(tag)
	if err != nil {
		return nil, err
	}
	var items []*Item
	if err := c.listResources(ctx, "/api/tags/"+id+"/items", page, &items); err != nil {
		return nil, err
	}
	return items, nil
}

// GetCategoryOfTag retrieves the category associated with a tag.
func (c *httpClient) GetCategoryOfTag(ctx context.Context, tag *Tag) (*TagCategory, error) {
	id, err := c.getID(tag)
	if err != nil {
		return nil, err
	}
	var result TagCategory
	if err := c.getResource(ctx, "/api/tags/"+id+"/category", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateTagCategory creates a new tag category.
func (c *httpClient) CreateTagCategory(ctx context.Context, category *TagCategory) (*TagCategory, error) {
	var result TagCategory
	if err := c.postResource(ctx, "/api/tag_categories", category, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTagCategory retrieves a tag category.
func (c *httpClient) GetTagCategory(ctx context.Context, category *TagCategory) (*TagCategory, error) {
	id, err := c.getID(category)
	if err != nil {
		return nil, err
	}
	var result TagCategory
	if err := c.getResource(ctx, "/api/tag_categories/"+id, &result); err != nil {
		return nil, err
	}
	return &result, nil
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
func (c *httpClient) UpdateTagCategory(ctx context.Context, category *TagCategory) (*TagCategory, error) {
	id, err := c.getID(category)
	if err != nil {
		return nil, err
	}
	var result TagCategory
	if err := c.putResource(ctx, "/api/tag_categories/"+id, category, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchTagCategory updates an existing tag category partially.
func (c *httpClient) PatchTagCategory(ctx context.Context, category *TagCategory) (*TagCategory, error) {
	id, err := c.getID(category)
	if err != nil {
		return nil, err
	}
	var result TagCategory
	if err := c.patchResource(ctx, "/api/tag_categories/"+id, category, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteTagCategory deletes a tag category.
func (c *httpClient) DeleteTagCategory(ctx context.Context, category *TagCategory) error {
	id, err := c.getID(category)
	if err != nil {
		return err
	}
	return c.deleteResource(ctx, "/api/tag_categories/"+id)
}

// ListTagCategoryTags retrieves tags in a tag category.
func (c *httpClient) ListTagCategoryTags(ctx context.Context, category *TagCategory, page int) ([]*Tag, error) {
	id, err := c.getID(category)
	if err != nil {
		return nil, err
	}
	var tags []*Tag
	if err := c.listResources(ctx, "/api/tag_categories/"+id+"/tags", page, &tags); err != nil {
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

// GetTemplate retrieves a template.
func (c *httpClient) GetTemplate(ctx context.Context, template *Template) (*Template, error) {
	id, err := c.getID(template)
	if err != nil {
		return nil, err
	}
	var result Template
	if err := c.getResource(ctx, "/api/templates/"+id, &result); err != nil {
		return nil, err
	}
	return &result, nil
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
func (c *httpClient) UpdateTemplate(ctx context.Context, template *Template) (*Template, error) {
	id, err := c.getID(template)
	if err != nil {
		return nil, err
	}
	var result Template
	if err := c.putResource(ctx, "/api/templates/"+id, template, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchTemplate updates an existing template partially.
func (c *httpClient) PatchTemplate(ctx context.Context, template *Template) (*Template, error) {
	id, err := c.getID(template)
	if err != nil {
		return nil, err
	}
	var result Template
	if err := c.patchResource(ctx, "/api/templates/"+id, template, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteTemplate deletes a template.
func (c *httpClient) DeleteTemplate(ctx context.Context, template *Template) error {
	id, err := c.getID(template)
	if err != nil {
		return err
	}
	return c.deleteResource(ctx, "/api/templates/"+id)
}

// GetUser retrieves a user.
func (c *httpClient) GetUser(ctx context.Context, user *User) (*User, error) {
	id, err := c.getID(user)
	if err != nil {
		return nil, err
	}
	var result User
	if err := c.getResource(ctx, "/api/users/"+id, &result); err != nil {
		return nil, err
	}
	return &result, nil
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

// GetWish retrieves a wish.
func (c *httpClient) GetWish(ctx context.Context, wish *Wish) (*Wish, error) {
	id, err := c.getID(wish)
	if err != nil {
		return nil, err
	}
	var result Wish
	if err := c.getResource(ctx, "/api/wishes/"+id, &result); err != nil {
		return nil, err
	}
	return &result, nil
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
func (c *httpClient) UpdateWish(ctx context.Context, wish *Wish) (*Wish, error) {
	id, err := c.getID(wish)
	if err != nil {
		return nil, err
	}
	var result Wish
	if err := c.putResource(ctx, "/api/wishes/"+id, wish, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchWish updates an existing wish partially.
func (c *httpClient) PatchWish(ctx context.Context, wish *Wish) (*Wish, error) {
	id, err := c.getID(wish)
	if err != nil {
		return nil, err
	}
	var result Wish
	if err := c.patchResource(ctx, "/api/wishes/"+id, wish, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteWish deletes a wish.
func (c *httpClient) DeleteWish(ctx context.Context, wish *Wish) error {
	id, err := c.getID(wish)
	if err != nil {
		return err
	}
	return c.deleteResource(ctx, "/api/wishes/"+id)
}

// UploadWishImage uploads an image for a wish.
func (c *httpClient) UploadWishImage(ctx context.Context, wish *Wish, file []byte) (*Wish, error) {
	id, err := c.getID(wish)
	if err != nil {
		return nil, err
	}
	var result Wish
	if err := c.uploadFile(ctx, "/api/wishes/"+id+"/image", file, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWishWishlist retrieves the wishlist associated with a wish.
func (c *httpClient) GetWishWishlist(ctx context.Context, wish *Wish) (*Wishlist, error) {
	id, err := c.getID(wish)
	if err != nil {
		return nil, err
	}
	var result Wishlist
	if err := c.getResource(ctx, "/api/wishes/"+id+"/wishlist", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateWishlist creates a new wishlist.
func (c *httpClient) CreateWishlist(ctx context.Context, wishlist *Wishlist) (*Wishlist, error) {
	var result Wishlist
	if err := c.postResource(ctx, "/api/wishlists", wishlist, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWishlist retrieves a wishlist.
func (c *httpClient) GetWishlist(ctx context.Context, wishlist *Wishlist) (*Wishlist, error) {
	id, err := c.getID(wishlist)
	if err != nil {
		return nil, err
	}
	var result Wishlist
	if err := c.getResource(ctx, "/api/wishlists/"+id, &result); err != nil {
		return nil, err
	}
	return &result, nil
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
func (c *httpClient) UpdateWishlist(ctx context.Context, wishlist *Wishlist) (*Wishlist, error) {
	id, err := c.getID(wishlist)
	if err != nil {
		return nil, err
	}
	var result Wishlist
	if err := c.putResource(ctx, "/api/wishlists/"+id, wishlist, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PatchWishlist updates an existing wishlist partially.
func (c *httpClient) PatchWishlist(ctx context.Context, wishlist *Wishlist) (*Wishlist, error) {
	id, err := c.getID(wishlist)
	if err != nil {
		return nil, err
	}
	var result Wishlist
	if err := c.patchResource(ctx, "/api/wishlists/"+id, wishlist, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteWishlist deletes a wishlist.
func (c *httpClient) DeleteWishlist(ctx context.Context, wishlist *Wishlist) error {
	id, err := c.getID(wishlist)
	if err != nil {
		return err
	}
	return c.deleteResource(ctx, "/api/wishlists/"+id)
}

// ListWishlistWishes retrieves wishes in a wishlist.
func (c *httpClient) ListWishlistWishes(ctx context.Context, wishlist *Wishlist, page int) ([]*Wish, error) {
	id, err := c.getID(wishlist)
	if err != nil {
		return nil, err
	}
	var wishes []*Wish
	if err := c.listResources(ctx, "/api/wishlists/"+id+"/wishes", page, &wishes); err != nil {
		return nil, err
	}
	return wishes, nil
}

// ListWishlistChildren retrieves child wishlists of a wishlist.
func (c *httpClient) ListWishlistChildren(ctx context.Context, wishlist *Wishlist, page int) ([]*Wishlist, error) {
	id, err := c.getID(wishlist)
	if err != nil {
		return nil, err
	}
	var wishlists []*Wishlist
	if err := c.listResources(ctx, "/api/wishlists/"+id+"/children", page, &wishlists); err != nil {
		return nil, err
	}
	return wishlists, nil
}

// UploadWishlistImage uploads an image for a wishlist.
func (c *httpClient) UploadWishlistImage(ctx context.Context, wishlist *Wishlist, file []byte) (*Wishlist, error) {
	id, err := c.getID(wishlist)
	if err != nil {
		return nil, err
	}
	var result Wishlist
	if err := c.uploadFile(ctx, "/api/wishlists/"+id+"/image", file, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetWishlistParent retrieves the parent wishlist of a wishlist.
func (c *httpClient) GetWishlistParent(ctx context.Context, wishlist *Wishlist) (*Wishlist, error) {
	id, err := c.getID(wishlist)
	if err != nil {
		return nil, err
	}
	var result Wishlist
	if err := c.getResource(ctx, "/api/wishlists/"+id+"/parent", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DisplayMetricsTable displays the metrics as a text-based table.
func DisplayMetricsTable(w io.Writer, metrics *Metrics) {
	if metrics == nil || len(*metrics) == 0 {
		fmt.Fprintln(w, "No metrics available")
		return
	}

	// Find the longest key and value for column widths.
	maxKeyLen, maxValueLen := 4, 5 // Minimum lengths for "Key" and "Value" headers.
	for key, value := range *metrics {
		if len(key) > maxKeyLen {
			maxKeyLen = len(key)
		}
		if len(value) > maxValueLen {
			maxValueLen = len(value)
		}
	}

	// Create format strings for headers and rows.
	headerFormat := fmt.Sprintf("%%-%ds | %%-%ds\n", maxKeyLen, maxValueLen)
	rowFormat := fmt.Sprintf("%%-%ds | %%-%ds\n", maxKeyLen, maxValueLen)
	separator := strings.Repeat("-", maxKeyLen) + "-+-" + strings.Repeat("-", maxValueLen)

	// Print table.
	fmt.Fprintf(w, headerFormat, "Key", "Value")
	fmt.Fprintln(w, separator)
	for key, value := range *metrics {
		fmt.Fprintf(w, rowFormat, key, value)
	}
}
