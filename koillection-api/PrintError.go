package koiApi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// hasNonPrintable checks if a string contains non-printable characters (outside ASCII 32–126).
func hasNonPrintable(s string) bool {
	for _, r := range s {
		if r < 32 || r > 126 {
			return true
		}
	}
	return false
}

// sanitizeNonJSONBody checks a non-JSON body for non-printable characters and returns a printable representation.
func sanitizeNonJSONBody(body string) string {
	if hasNonPrintable(body) {
		return fmt.Sprintf("<%d bytes of binary data>", len(body))
	}
	return body
}

// replaceNonPrintableElements recursively inspects a JSON structure, replaces elements containing non-printable characters, and limits arrays to 3 elements.
func replaceNonPrintableElements(data interface{}) (interface{}, error) {
	switch v := data.(type) {
	case string:
		// Replace strings with non-printable characters.
		if hasNonPrintable(v) {
			return fmt.Sprintf("<%d bytes of binary data>", len(v)), nil
		}
		return v, nil
	case []interface{}:
		// Limit array to 3 elements.
		limit := 3
		if len(v) < limit {
			limit = len(v)
		}
		result := make([]interface{}, limit)
		for i := 0; i < limit; i++ {
			modified, err := replaceNonPrintableElements(v[i])
			if err != nil {
				return nil, err
			}
			result[i] = modified
		}
		// Check if any element is a placeholder, indicating binary data.
		for _, item := range result {
			if str, ok := item.(string); ok && strings.Contains(str, "bytes of binary data") {
				bytes, err := json.Marshal(v)
				if err != nil {
					return fmt.Sprintf("<unknown bytes of binary data>"), nil
				}
				return fmt.Sprintf("<%d bytes of binary data>", len(bytes)), nil
			}
		}
		// Add "plus N more" if array has more than 3 elements.
		if remaining := len(v) - limit; remaining > 0 {
			result = append(result, fmt.Sprintf("plus %d more", remaining))
		}
		return result, nil
	case map[string]interface{}:
		// Process map values recursively.
		result := make(map[string]interface{})
		for key, value := range v {
			modified, err := replaceNonPrintableElements(value)
			if err != nil {
				return nil, err
			}
			result[key] = modified
		}
		// Check if any value is a placeholder.
		for _, value := range result {
			if str, ok := value.(string); ok && strings.Contains(str, "bytes of binary data") {
				bytes, err := json.Marshal(v)
				if err != nil {
					return fmt.Sprintf("<unknown bytes of binary data>"), nil
				}
				return fmt.Sprintf("<%d bytes of binary data>", len(bytes)), nil
			}
		}
		return result, nil
	default:
		// Return non-string/array/map values unchanged.
		return v, nil
	}
}

// PrintError prints the request headers, request body, response headers, response body, and error struct or raw error text from the httpClient struct to stdout.
func (c *httpClient) PrintError(ctx context.Context) {
	const maxFieldSize = 1024 // Threshold for large fields in bytes.

	// Print request headers.
	fmt.Println("Request Headers:")
	if c.lastRequest == nil || len(c.lastRequest.Header) == 0 {
		fmt.Println("  No headers")
	} else {
		for key, values := range c.lastRequest.Header {
			joinedValue := strings.Join(values, ", ")
			if len(joinedValue) > maxFieldSize || hasNonPrintable(joinedValue) {
				fmt.Printf("  %s: <%d bytes of data>\n", key, len(joinedValue))
			} else {
				fmt.Printf("  %s: %s\n", key, joinedValue)
			}
		}
	}

	// Print request body.
	fmt.Println("Request Body:")
	if c.lastRequestBody == nil {
		fmt.Println("  No body")
	} else {
		var jsonData interface{}
		if err := json.Unmarshal(c.lastRequestBody, &jsonData); err == nil {
			// If body is valid JSON, replace elements with non-printable characters and limit arrays.
			modifiedData, _ := replaceNonPrintableElements(jsonData)
			prettyJSON, err := json.MarshalIndent(modifiedData, "  ", "  ")
			if err == nil {
				fmt.Printf("  %s\n", string(prettyJSON))
			} else {
				fmt.Printf("  %s\n", sanitizeNonJSONBody(string(c.lastRequestBody)))
			}
		} else {
			// If not JSON, sanitize for non-printable characters.
			fmt.Printf("  %s\n", sanitizeNonJSONBody(string(c.lastRequestBody)))
		}
	}

	// Print response headers.
	fmt.Println("Response Headers:")
	if c.lastResponse == nil || len(c.lastResponse.Header) == 0 {
		fmt.Println("  No headers")
	} else {
		for key, values := range c.lastResponse.Header {
			joinedValue := strings.Join(values, ", ")
			if len(joinedValue) > maxFieldSize || hasNonPrintable(joinedValue) {
				fmt.Printf("  %s: <%d bytes of data>\n", key, len(joinedValue))
			} else {
				fmt.Printf("  %s: %s\n", key, joinedValue)
			}
		}
	}

	// Print response body.
	fmt.Println("Response Body:")
	if c.rawError == "" {
		fmt.Println("  No body")
	} else {
		var jsonData interface{}
		if err := json.Unmarshal([]byte(c.rawError), &jsonData); err == nil {
			// If rawError is valid JSON, replace elements with non-printable characters and limit arrays.
			modifiedData, _ := replaceNonPrintableElements(jsonData)
			prettyJSON, err := json.MarshalIndent(modifiedData, "  ", "  ")
			if err == nil {
				fmt.Printf("  %s\n", string(prettyJSON))
			} else {
				fmt.Printf("  %s\n", sanitizeNonJSONBody(c.rawError))
			}
		} else {
			// If not JSON, sanitize for non-printable characters.
			fmt.Printf("  %s\n", sanitizeNonJSONBody(c.rawError))
		}
	}

	// Print error details.
	if c.lastResponse != nil && (c.lastResponse.StatusCode == http.StatusBadRequest || c.lastResponse.StatusCode == http.StatusUnprocessableEntity) && c.koiError == nil {
		// RawError was printed as response body for 400/422 if unmarshaling failed.
		return
	}
	if c.koiError != nil {
		fmt.Printf("Error Response (Status %d):\n", c.koiError.Status)
		fmt.Printf("  Title: %s\n", c.koiError.Title)
		fmt.Printf("  Detail: %s\n", c.koiError.Detail)
		fmt.Printf("  Description: %s\n", c.koiError.Description)
		fmt.Printf("  Context: %s\n", c.koiError.Context)
		fmt.Printf("  ID: %s\n", c.koiError.ID)
		fmt.Printf("  Type: %s\n", c.koiError.Type)
		fmt.Printf("  Instance: %s\n", c.koiError.Instance)
		if len(c.koiError.Violations) > 0 {
			fmt.Println("  Violations:")
			limit := 3
			if len(c.koiError.Violations) < limit {
				limit = len(c.koiError.Violations)
			}
			for i := 0; i < limit; i++ {
				v := c.koiError.Violations[i]
				fmt.Printf("    %d. Property: %s, Message: %s\n", i+1, v.PropertyPath, v.Message)
			}
			if remaining := len(c.koiError.Violations) - limit; remaining > 0 {
				fmt.Printf("    plus %d more\n", remaining)
			}
		}
		return
	}
	if c.rawError != "" {
		// RawError was printed as response body.
		return
	}
	fmt.Println("Error Response: No error data available")
}
