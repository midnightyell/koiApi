package koiApi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Credentials holds the username and password from creds.yaml.
type Credentials struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// loadCredentials reads username and password from creds.yaml.
func loadCredentials(t *testing.T) (Credentials, error) {
	data, err := os.ReadFile("creds.json")
	if err != nil {
		return Credentials{}, err
	}
	var creds Credentials
	if err := json.Unmarshal(data, &creds); err != nil {
		return Credentials{}, err
	}
	return creds, nil
}

func TestCollectionAndItemLifecycleWithRealServer(t *testing.T) {
	// Load credentials
	creds, err := loadCredentials(t)
	assert.NoError(t, err, "Failed to load creds.yaml")
	assert.NotEmpty(t, creds.URL, "Username is empty")
	assert.NotEmpty(t, creds.Username, "Username is empty")
	assert.NotEmpty(t, creds.Password, "Password is empty")

	// Initialize client (assumes a real client implementation)
	client := NewKoiClient(creds.URL, 30*time.Second)

	:= context.Background()

	// Authenticate
	token, err := client.CheckLogin(creds.Username, creds.Password)
	assert.NoError(t, err, "Failed to authenticate")
	assert.NotEmpty(t, token, "Authentication token is empty")
	// Assume client is configured to use token for subsequent requests

	RemoveAllCollections(t, client, )

	// Create collection
	var testName = "TEST" + time.Now().Format("15:04:05")
	collection := &Collection{
		Title:      testName,
		Visibility: VisibilityPublic,
	}
	resultCollection, err := client.CreateCollection(collection)
	assert.NoError(t, err, "Failed to create collection")
	assert.NotEmpty(t, resultCollection.ID, "Collection ID is empty")
	assert.Equal(t, testName, resultCollection.Title, "Collection title mismatch")

	// Create IRI from collection ID
	collectionIRI := resultCollection.IRI()
	assert.NotEmpty(t, collectionIRI, "Collection IRI is empty")
	assert.Equal(t, fmt.Sprintf("/api/collections/%s", resultCollection.ID), collectionIRI, "Collection IRI mismatch")
	resultCollection.UploadImageByFile(client, "./picture002.jpg")

	// Create item
	item := &Item{
		Name:       "Test Item",
		Quantity:   1,
		Collection: &collectionIRI,
		Visibility: VisibilityPublic,
	}
	resultItem, err := client.CreateItem(item)
	assert.NoError(t, err, "Failed to create item")
	assert.NotEmpty(t, resultItem.ID, "Item ID is empty")
	assert.Equal(t, "Test Item", resultItem.Name, "Item name mismatch")

	// Create one datum for each DatumType
	datumTypes := []DatumType{
		DatumTypeText, DatumTypeTextarea, DatumTypeCountry, DatumTypeDate,
		DatumTypeRating, DatumTypeNumber, DatumTypePrice, DatumTypeLink,
		DatumTypeList,
		DatumTypeChoiceList,
		DatumTypeCheckbox, DatumTypeImage,
		DatumTypeFile, DatumTypeSign, DatumTypeVideo,
		DatumTypeBlankLine, DatumTypeSection,
	}
	itemIRI := resultItem.IRI()
	resultItem, err = resultItem.UploadImageByFile(client, "./picture001.jpg")
	if err != nil {
		client.PrintError()
		t.Fatalf("Failed to upload image: %v\n", err)
	}
	resultItem, err = resultItem.UploadImageByFile(client, "./picture002.jpg")
	if err != nil {
		t.Fatalf("Failed to upload image: %v\n", err)
	}

	assert.NotEmpty(t, itemIRI, "Item IRI is empty")
	assert.Equal(t, fmt.Sprintf("/api/items/%s", resultItem.ID), itemIRI, "Item IRI mismatch")

	// Create a datum for the it
	name := "Name of the thing"

	datum := &Datum{
		Item:       &itemIRI,
		DatumType:  DatumTypeText,
		Label:      "Name",
		Value:      &name,
		Visibility: VisibilityPublic,
	}

	datum.Print("%s:\n", *datum.Item)
	fmt.Printf("\n")

	resultDatum, err := client.CreateDatum(datum)
	assert.NoError(t, err, "Failed to create datum")
	assert.NotEmpty(t, resultDatum.ID, "Datum ID is empty")
	assert.Equal(t, DatumTypeText, resultDatum.DatumType, "Datum type mismatch")
	assert.Equal(t, name, *resultDatum.Value, "Datum value mismatch")
	assert.Equal(t, VisibilityPublic, resultDatum.Visibility, "Datum visibility mismatch")
	assert.Equal(t, itemIRI, *resultDatum.Item, "Datum item mismatch")

	for _, dt := range datumTypes {
		value := "test-value"
		currency := ""
		if dt == DatumTypeCheckbox {
			value = "1"
		} else if dt == DatumTypeDate {
			value = "2023-01-01"
		} else if dt == DatumTypeNumber {
			value = "42"
		} else if dt == DatumTypeRating {
			value = "5"
		} else if dt == DatumTypePrice {
			value = "99.99"
			currency = "EUR"
		} else if dt == DatumTypeCountry {
			value = "US"
		} else if dt == DatumTypeLink {
			value = "https://example.com"
		} else if dt == DatumTypeChoiceList {
			continue
		} else if dt == DatumTypeList {
			value = `["List item 1", "List item 2"]`
		} else if dt == DatumTypeSign || dt == DatumTypeImage || dt == DatumTypeFile || dt == DatumTypeVideo {
			continue
		} else if dt == DatumTypeBlankLine {
			value = "This is a blank line"
		} else if dt == DatumTypeSection {
			value = "Section title"
		}
		var strP *string = &value
		if *strP == "" {
			strP = nil
		}
		datum = &Datum{
			Item:       &itemIRI,
			DatumType:  dt,
			Label:      string(dt) + " Field",
			Value:      strP,
			Visibility: VisibilityPublic,
			Currency:   &currency,
		}

		datum.Print("%s:\n", *datum.Item)
		fmt.Printf("\n")

		err = validateDatumValue(dt, value)
		if err != nil {
			t.Fatalf("Invalid datum value '%s' for type %s: %v", value, dt, err)
		}

		resultDatum, err := client.CreateDatum(datum)
		assert.NoError(t, err, "Failed to create datum for type %s", dt)
		assert.NotEmpty(t, resultDatum.ID, "Datum ID is empty for type %s", dt)
		//assert.Equal(t, dt, resultDatum.DatumType, "Datum type mismatch for %s", dt)
		//assert.Equal(t, value, *resultDatum.Value, "Datum value mismatch for %s", dt)

		// Upload file for image, file, or video types
		//if datum.Value != nil {
		if dt == DatumTypeImage {
			_, err = resultDatum.UploadImageByFile(client, "./picture002.jpg")
			assert.NoError(t, err, "Failed to upload image for datum %s", dt)
		} else if dt == DatumTypeFile {
			_, err = resultDatum.UploadFileByFile(client, "./picture002.jpg")
			assert.NoError(t, err, "Failed to upload file for datum %s", dt)
		} else if dt == DatumTypeVideo {
			_, err = resultDatum.UploadVideoByFile(client, "./picture002.jpg")
			assert.NoError(t, err, "Failed to upload video for datum %s", dt)
		}
		//}
	}

	// Fetch item
	fetchedItem, err := client.GetItem(resultItem.ID)
	assert.NoError(t, err, "Failed to fetch item")
	assert.Equal(t, resultItem.ID, fetchedItem.ID, "Fetched item ID mismatch")
	assert.Equal(t, resultItem.Name, fetchedItem.Name, "Fetched item name mismatch")

	// Delete item
	//err = client.DeleteItem(resultItem.ID)
	//assert.NoError(t, err, "Failed to delete item")

	// Delete collection
	//err = client.DeleteCollection(resultCollection.ID)
	//assert.NoError(t, err, "Failed to delete collection")
}

// validateDatumValue validates the value for a given DatumType.
func validateDatumValue(dt DatumType, value string) error {
	switch dt {
	case DatumTypeCountry:
		// Simplified list of ISO 3166-1 alpha-2 codes
		validCountries := map[string]bool{
			"US": true, "FR": true, "JP": true, "GB": true, "CA": true,
			// Add more codes as needed
		}
		if !validCountries[value] {
			return fmt.Errorf("invalid country code: %s, must be a 2-letter ISO 3166-1 alpha-2 code", value)
		}
	case DatumTypeRating:
		num, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("rating must be an integer, got %s", value)
		}
		if num < 0 || num > 10 {
			return fmt.Errorf("rating must be between 0 and 10, got %d", num)
		}
	case DatumTypeCheckbox:
		if value != "0" && value != "1" {
			return fmt.Errorf("checkbox must be '0' or '1', got %s", value)
		}
	case DatumTypeLink:
		parsedURL, err := url.Parse(value)
		if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
			return fmt.Errorf("link must be a valid URL, got %s", value)
		}
	case DatumTypeChoiceList:
		var list []string
		if err := json.Unmarshal([]byte(value), &list); err != nil {
			return fmt.Errorf("choice-list must be a JSON string array, e.g., '[\"Value 1\", \"Value 2\"]', got %s", value)
		}
	case DatumTypeText, DatumTypeTextarea, DatumTypeList, DatumTypeSign, DatumTypeBlankLine, DatumTypeSection:
		if value == "" {
			return fmt.Errorf("%s value cannot be empty", dt)
		}
	case DatumTypeDate:
		if _, err := time.Parse("2006-01-02", value); err != nil {
			return fmt.Errorf("date must be in YYYY-MM-DD format, got %s", value)
		}
	case DatumTypeNumber:
		if _, err := strconv.Atoi(value); err != nil {
			return fmt.Errorf("number must be an integer, got %s", value)
		}
		/* 	case DatumTypePrice:
		if _, err := strconv.ParseFloat(value, 64); err != nil {
			return fmt.Errorf("price must be a number, got %s", value)
		}
		*/
	}
	return nil
}

func RemoveAllCollections(t *testing.T, client Client) {
	collections, err := client.ListCollections(1)
	if err != nil {
		t.Fatalf("Failed to list collections: %v", err)
	}
	for _, collection := range collections {
		err = client.DeleteCollection(collection.ID)
		if err != nil {
			t.Fatalf("Failed to delete collection %s: %v", collection.ID, err)
		}
	}
}
