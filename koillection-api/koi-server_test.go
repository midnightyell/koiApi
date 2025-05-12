package koiApi

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

// Credentials holds the username and password from creds.yaml.
type Credentials struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// loadCredentials reads username and password from creds.yaml.
func loadCredentials(t *testing.T) (Credentials, error) {
	data, err := ioutil.ReadFile("creds.yaml")
	if err != nil {
		return Credentials{}, err
	}
	var creds Credentials
	if err := yaml.Unmarshal(data, &creds); err != nil {
		return Credentials{}, err
	}
	return creds, nil
}

func TestCollectionAndItemLifecycleWithRealServer(t *testing.T) {
	// Load credentials
	creds, err := loadCredentials(t)
	assert.NoError(t, err, "Failed to load creds.yaml")
	assert.NotEmpty(t, creds.Username, "Username is empty")
	assert.NotEmpty(t, creds.Password, "Password is empty")

	// Initialize client (assumes a real client implementation)
	// Replace with your actual client constructor, e.g., NewClient("http://localhost:80")
	client := &YourRealClientImplementation{BaseURL: "http://localhost:80"} // TODO: Replace with actual client

	ctx := context.Background()

	// Authenticate
	token, err := client.CheckLogin(ctx, creds.Username, creds.Password)
	assert.NoError(t, err, "Failed to authenticate")
	assert.NotEmpty(t, token, "Authentication token is empty")
	// Assume client is configured to use token for subsequent requests

	// Create collection
	collection := &Collection{
		Title:      "TEST",
		Visibility: VisibilityPublic,
	}
	resultCollection, err := client.CreateCollection(ctx, collection)
	assert.NoError(t, err, "Failed to create collection")
	assert.NotEmpty(t, resultCollection.ID, "Collection ID is empty")
	assert.Equal(t, "TEST", resultCollection.Title, "Collection title mismatch")

	// Create item
	collectionIRI := IRI("/api/collections/" + string(resultCollection.ID))
	item := &Item{
		Name:       "Test Item",
		Quantity:   1,
		Collection: &collectionIRI,
		Visibility: VisibilityPublic,
	}
	resultItem, err := client.CreateItem(ctx, item)
	assert.NoError(t, err, "Failed to create item")
	assert.NotEmpty(t, resultItem.ID, "Item ID is empty")
	assert.Equal(t, "Test Item", resultItem.Name, "Item name mismatch")

	// Create one datum for each DatumType
	datumTypes := []DatumType{
		DatumTypeText, DatumTypeTextarea, DatumTypeCountry, DatumTypeDate,
		DatumTypeRating, DatumTypeNumber, DatumTypePrice, DatumTypeLink,
		DatumTypeList, DatumTypeChoiceList, DatumTypeCheckbox, DatumTypeImage,
		DatumTypeFile, DatumTypeSign, DatumTypeVideo, DatumTypeBlankLine, DatumTypeSection,
	}
	itemIRI := IRI("/api/items/" + string(resultItem.ID))
	for _, dt := range datumTypes {
		value := "test-value"
		if dt == DatumTypeCheckbox {
			value = "true"
		} else if dt == DatumTypeDate {
			value = "2023-01-01"
		} else if dt == DatumTypeNumber || dt == DatumTypeRating {
			value = "42"
		} else if dt == DatumTypePrice {
			value = "99.99"
		} else if dt == DatumTypeImage || dt == DatumTypeFile || dt == DatumTypeVideo {
			value = "/path/to/file" // Placeholder for file-based types
		}
		datum := &Datum{
			Item:       &itemIRI,
			DatumType:  dt,
			Label:      string(dt) + " Field",
			Value:      &value,
			Visibility: VisibilityPublic,
		}
		resultDatum, err := client.CreateDatum(ctx, datum)
		assert.NoError(t, err, "Failed to create datum for type %s", dt)
		assert.NotEmpty(t, resultDatum.ID, "Datum ID is empty for type %s", dt)
		assert.Equal(t, dt, resultDatum.DatumType, "Datum type mismatch for %s", dt)
		assert.Equal(t, value, *resultDatum.Value, "Datum value mismatch for %s", dt)

		// Upload file for image, file, or video types
		fileData := []byte("placeholder file content") // Replace with actual file data
		if dt == DatumTypeImage {
			_, err = client.UploadDatumImage(ctx, resultDatum.ID, fileData)
			assert.NoError(t, err, "Failed to upload image for datum %s", dt)
		} else if dt == DatumTypeFile {
			_, err = client.UploadDatumFile(ctx, resultDatum.ID, fileData)
			assert.NoError(t, err, "Failed to upload file for datum %s", dt)
		} else if dt == DatumTypeVideo {
			_, err = client.UploadDatumVideo(ctx, resultDatum.ID, fileData)
			assert.NoError(t, err, "Failed to upload video for datum %s", dt)
		}
	}

	// Fetch item
	fetchedItem, err := client.GetItem(ctx, resultItem.ID)
	assert.NoError(t, err, "Failed to fetch item")
	assert.Equal(t, resultItem.ID, fetchedItem.ID, "Fetched item ID mismatch")
	assert.Equal(t, resultItem.Name, fetchedItem.Name, "Fetched item name mismatch")

	// Delete item
	err = client.DeleteItem(ctx, resultItem.ID)
	assert.NoError(t, err, "Failed to delete item")

	// Delete collection
	err = client.DeleteCollection(ctx, resultCollection.ID)
	assert.NoError(t, err, "Failed to delete collection")
}
