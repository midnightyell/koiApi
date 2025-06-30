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

// Errors for common HTTP status codes.
var (
	ErrInvalidInput  = errors.New("invalid input")
	ErrNotFound      = errors.New("resource not found")
	ErrUnprocessable = errors.New("unprocessable entity")
	ErrUnauthorized  = errors.New("unauthorized")
)

// KoiError represents a union of the 400 and 422 error response structures.
type KoiError struct {
	Context     string      `json:"@context,omitempty"`
	ID          string      `json:"@id,omitempty"`
	Type        string      `json:"@type,omitempty"`
	Title       string      `json:"title,omitempty"`
	Detail      string      `json:"detail,omitempty"`
	Status      int         `json:"status,omitempty"`
	Instance    string      `json:"instance,omitempty"`
	Description string      `json:"description,omitempty"`
	Violations  []Violation `json:"violations,omitempty"`
}

// Violation represents a single violation in a 422 error response.
type Violation struct {
	PropertyPath string `json:"propertyPath"`
	Message      string `json:"message"`
}

// httpClient implements the Client interface using net/http.
type httpClient struct {
	baseURL         string
	httpClient      *http.Client
	token           string
	lastError       error
	lastRequest     *http.Request
	lastRequestBody []byte
	lastResponse    *http.Response
	koiError        *KoiError
	rawError        string
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
		lastRequest:     nil,
		lastRequestBody: nil,
		lastResponse:    nil,
		koiError:        nil,
		rawError:        "",
	}
}

// GetResponse retrieves the response from the httpClient struct.
func (c *httpClient) GetResponse(ctx context.Context) string {
	if c.koiError != nil {
		// Return the structured error if available.
		errBytes, err := json.MarshalIndent(c.koiError, "", "  ")
		if err != nil {
			return fmt.Sprintf("Error marshaling KoiError: %v\nRaw Error: %s", err, c.rawError)
		}
		return fmt.Sprintf("Response Status: %s\nBody: %s", c.lastResponse.Status, string(errBytes))
	}
	if c.rawError != "" {
		// Return the raw error text.
		return fmt.Sprintf("Response Status: %s\nBody: %s", c.lastResponse.Status, c.rawError)
	}
	if c.lastResponse == nil {
		return "No response"
	}
	body, err := io.ReadAll(c.lastResponse.Body)
	// Reset the response body so it can be read again if needed.
	c.lastResponse.Body = io.NopCloser(bytes.NewReader(body))
	if err != nil {
		return fmt.Sprintf("Error reading response body: %v\nRaw Error: %s", err, c.rawError)
	}
	return fmt.Sprintf("Response Status: %s\nBody: %s", c.lastResponse.Status, string(body))
}

// doRequest sends an HTTP request, stores it and the response in the httpClient struct, and returns the response.
func (c *httpClient) doRequest(ctx context.Context, method, path string, body io.Reader, multipartContentType string) (*http.Response, error) {
	var bodyBytes []byte
	if body != nil {
		var err error
		bodyBytes, err = io.ReadAll(body)
		if err != nil {
			c.lastRequest = nil
			c.lastRequestBody = nil
			c.lastResponse = nil
			c.koiError = nil
			c.rawError = fmt.Sprintf("Error reading request body: %v", err)
			c.lastError = err
			return nil, fmt.Errorf("reading request body: %w", err)
		}
	}

	// Reset the body for the request.
	var reqBody io.Reader
	if bodyBytes != nil {
		reqBody = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, reqBody)
	if err != nil {
		c.lastRequest = nil
		c.lastRequestBody = nil
		c.lastResponse = nil
		c.koiError = nil
		c.rawError = fmt.Sprintf("Error creating request: %v", err)
		c.lastError = err
		return nil, fmt.Errorf("creating request: %w", err)
	}

	if multipartContentType != "" {
		req.Header.Set("Content-Type", multipartContentType)
	} else if body != nil {
		req.Header.Set("Content-Type", "application/ld+json")
	}
	if path == "/api/metrics" {
		req.Header.Set("Accept", "text/plain")
	} else {
		req.Header.Set("Accept", "application/ld+json")
	}

	resp, err := c.httpClient.Do(req)
	c.lastError = err
	c.lastRequest = req
	c.lastRequestBody = bodyBytes
	c.lastResponse = resp
	c.koiError = nil
	c.rawError = ""

	if err != nil {
		c.rawError = fmt.Sprintf("Request failed: %v", err)
		return nil, fmt.Errorf("sending request: %w", err)
	}

	// Read the response body for all status codes.
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.rawError = fmt.Sprintf("Error reading response body: %v", err)
	} else {
		c.rawError = string(respBodyBytes)
	}
	// Reset the response body so callers can read it.
	resp.Body = io.NopCloser(bytes.NewReader(respBodyBytes))

	// Handle 400 and 422 errors by attempting to unmarshal into KoiError.
	if resp.StatusCode == http.StatusBadRequest || resp.StatusCode == http.StatusUnprocessableEntity {
		if err == nil {
			var koiErr KoiError
			if err := json.Unmarshal(respBodyBytes, &koiErr); err != nil {
				c.koiError = nil // Explicitly set to nil if unmarshaling fails.
			} else {
				c.koiError = &koiErr
			}
		}
	}

	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusNoContent:
		return resp, nil
	case http.StatusBadRequest:
		return resp, ErrInvalidInput
	case http.StatusUnauthorized:
		return resp, ErrUnauthorized
	case http.StatusNotFound:
		return resp, ErrNotFound
	case http.StatusUnprocessableEntity:
		return resp, ErrUnprocessable
	default:
		return resp, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
}

// getResource retrieves a single resource and decodes it into the provided struct.
func (c *httpClient) getResource(ctx context.Context, path string, out interface{}) error {
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil, "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(out)
}

// listResources retrieves a paginated list of resources and decodes the member array.
func (c *httpClient) listResources(ctx context.Context, path string, page int, out interface{}, queryParams ...string) error {
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return fmt.Errorf("parsing URL: %w", err)
	}
	q := u.Query()
	if page > 0 {
		q.Set("page", strconv.Itoa(page))
	}
	// Append query parameters
	for _, param := range queryParams {
		if param != "" {
			parts := strings.SplitN(param, "=", 2)
			if len(parts) == 2 {
				q.Set(parts[0], parts[1])
			}
		}
	}
	u.RawQuery = q.Encode()

	resp, err := c.doRequest(ctx, http.MethodGet, u.Path+"?"+u.RawQuery, nil, "")
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

	resp, err := c.doRequest(ctx, http.MethodPost, path, bytes.NewReader(body), "")
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

	resp, err := c.doRequest(ctx, http.MethodPut, path, bytes.NewReader(body), "")
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

	resp, err := c.doRequest(ctx, http.MethodPatch, path, bytes.NewReader(body), "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(out)
}

// deleteResource deletes a resource.
func (c *httpClient) deleteResource(ctx context.Context, path string) error {
	resp, err := c.doRequest(ctx, http.MethodDelete, path, nil, "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// uploadFile uploads a file using multipart/form-data and decodes the response.
func (c *httpClient) uploadFile(ctx context.Context, path string, file []byte, fieldName string, out interface{}) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fieldName, "upload")
	if err != nil {
		return fmt.Errorf("creating form file: %w", err)
	}
	if _, err = part.Write(file); err != nil {
		return fmt.Errorf("writing file: %w", err)
	}
	if err := writer.Close(); err != nil {
		return fmt.Errorf("closing writer: %w", err)
	}

	contentType := writer.FormDataContentType()
	resp, err := c.doRequest(ctx, http.MethodPost, path, body, contentType)
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

	resp, err := c.doRequest(ctx, http.MethodPost, "/api/authentication_token", bytes.NewReader(body), "")
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
