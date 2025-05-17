// uploadFile uploads a file using multipart/form-data and decodes the response.
func uploadFile(ctx context.Context, path string, file []byte) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "upload")
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

	return err
}







{
	"title": "Photo Title",
	"comment": "Photo Comment",
	"place": "Photo Place",
	"file": "Photo File",
	"visibility": "public",
	"album": ""
  }

  "data": "https://i.ebayimg.com/images/g/gpUAAOSwwhloIfpB/s-l1600.webp"






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
			if len(joinedValue) > maxFieldSize {
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
	} else if len(c.lastRequestBody) > maxFieldSize {
		fmt.Printf("  <%d bytes of data>\n", len(c.lastRequestBody))
	} else {
		var jsonData interface{}
		if err := json.Unmarshal(c.lastRequestBody, &jsonData); err == nil {
			// If body is valid JSON, pretty-print it.
			prettyJSON, err := json.MarshalIndent(jsonData, "  ", "  ")
			if err == nil {
				fmt.Printf("  %s\n", string(prettyJSON))
			} else {
				fmt.Printf("  %s\n", string(c.lastRequestBody))
			}
		} else {
			// If not JSON, print as plain text.
			fmt.Printf("  %s\n", string(c.lastRequestBody))
		}
	}

	// Print response headers.
	fmt.Println("Response Headers:")
	if c.lastResponse == nil || len(c.lastResponse.Header) == 0 {
		fmt.Println("  No headers")
	} else {
		for key, values := range c.lastResponse.Header {
			joinedValue := strings.Join(values, ", ")
			if len(joinedValue) > maxFieldSize {
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
	} else if len(c.rawError) > maxFieldSize {
		fmt.Printf("  <%d bytes of data>\n", len(c.rawError))
	} else {
		var jsonData interface{}
		if err := json.Unmarshal([]byte(c.rawError), &jsonData); err == nil {
			// If rawError is valid JSON, pretty-print it.
			prettyJSON, err := json.MarshalIndent(jsonData, "  ", "  ")
			if err == nil {
				fmt.Printf("  %s\n", string(prettyJSON))
			} else {
				fmt.Printf("  %s\n", c.rawError)
			}
		} else {
			// If not JSON, print as plain text.
			fmt.Printf("  %s\n", c.rawError)
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