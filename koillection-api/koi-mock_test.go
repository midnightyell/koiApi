package koiApi

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockClient is a mock implementation of the Client interface for testing.
type MockClient struct {
	mock.Mock
}

// Implement all Client interface methods (only used methods are shown for brevity).
func (m *MockClient) CreateCollection(ctx context.Context, collection *Collection) (*Collection, error) {
	args := m.Called(ctx, collection)
	return args.Get(0).(*Collection), args.Error(1)
}

func (m *MockClient) CreateItem(ctx context.Context, item *Item) (*Item, error) {
	args := m.Called(ctx, item)
	return args.Get(0).(*Item), args.Error(1)
}

func (m *MockClient) CreateDatum(ctx context.Context, datum *Datum) (*Datum, error) {
	args := m.Called(ctx, datum)
	return args.Get(0).(*Datum), args.Error(1)
}

func (m *MockClient) GetItem(ctx context.Context, id ID) (*Item, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*Item), args.Error(1)
}

func (m *MockClient) DeleteItem(ctx context.Context, id ID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockClient) DeleteCollection(ctx context.Context, id ID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCollectionAndItemLifecycle(t *testing.T) {
	ctx := context.Background()
	client := new(MockClient)

	// Create collection
	collection := &Collection{
		Title:      "TEST",
		Visibility: VisibilityPublic,
	}
	createdCollection := &Collection{
		ID:         "collection-1",
		Title:      "TEST",
		Visibility: VisibilityPublic,
		CreatedAt:  time.Now(),
	}
	client.On("CreateCollection", ctx, collection).Return(createdCollection, nil)

	// Create item
	collectionID := "/api/collections/collection-1"
	item := &Item{
		Name:       "Test Item",
		Quantity:   1,
		Collection: &collectionID,
		Visibility: VisibilityPublic,
	}
	createdItem := &Item{
		ID:         "item-1",
		Name:       "Test Item",
		Quantity:   1,
		Collection: &collectionID,
		Visibility: VisibilityPublic,
		CreatedAt:  time.Now(),
	}
	client.On("CreateItem", ctx, item).Return(createdItem, nil)

	// Create one datum for each DatumType
	datumTypes := []DatumType{
		DatumTypeText, DatumTypeTextarea, DatumTypeCountry, DatumTypeDate,
		DatumTypeRating, DatumTypeNumber, DatumTypePrice, DatumTypeLink,
		DatumTypeList, DatumTypeChoiceList, DatumTypeCheckbox, DatumTypeImage,
		DatumTypeFile, DatumTypeSign, DatumTypeVideo, DatumTypeBlankLine, DatumTypeSection,
	}
	createdData := make([]*Datum, len(datumTypes))
	itemID := "/api/items/item-1"
	for i, dt := range datumTypes {
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
			value = "/path/to/file" // Placeholder, real test would upload file
		}
		datum := &Datum{
			Item:       &itemID,
			DatumType:  dt,
			Label:      string(dt) + " Field",
			Value:      &value,
			Visibility: VisibilityPublic,
		}
		createdDatum := &Datum{
			ID:         ID("datum-" + string(dt)),
			Item:       &itemID,
			DatumType:  dt,
			Label:      string(dt) + " Field",
			Value:      &value,
			Visibility: VisibilityPublic,
			CreatedAt:  time.Now(),
		}
		client.On("CreateDatum", ctx, datum).Return(createdDatum, nil)
		createdData[i] = createdDatum
	}

	// Fetch item
	fetchedItem := &Item{
		ID:         "item-1",
		Name:       "Test Item",
		Quantity:   1,
		Collection: &collectionID,
		Visibility: VisibilityPublic,
		CreatedAt:  createdItem.CreatedAt,
	}
	client.On("GetItem", ctx, ID("item-1")).Return(fetchedItem, nil)

	// Delete item and collection
	client.On("DeleteItem", ctx, ID("item-1")).Return(nil)
	client.On("DeleteCollection", ctx, ID("collection-1")).Return(nil)

	// Run test
	t.Run("CreateAndDeleteCollectionAndItem", func(t *testing.T) {
		// Create collection
		resultCollection, err := client.CreateCollection(ctx, collection)
		assert.NoError(t, err)
		assert.Equal(t, createdCollection.ID, resultCollection.ID)
		assert.Equal(t, "TEST", resultCollection.Title)

		// Create item
		resultItem, err := client.CreateItem(ctx, item)
		assert.NoError(t, err)
		assert.Equal(t, createdItem.ID, resultItem.ID)
		assert.Equal(t, "Test Item", resultItem.Name)

		// Create data for each DatumType
		for i, dt := range datumTypes {
			value := *createdData[i].Value
			datum := &Datum{
				Item:       &itemID,
				DatumType:  dt,
				Label:      string(dt) + " Field",
				Value:      &value,
				Visibility: VisibilityPublic,
			}
			resultDatum, err := client.CreateDatum(ctx, datum)
			assert.NoError(t, err)
			assert.Equal(t, createdData[i].ID, resultDatum.ID)
			assert.Equal(t, dt, resultDatum.DatumType)
			assert.Equal(t, value, *resultDatum.Value)
		}

		// Fetch item
		fetched, err := client.GetItem(ctx, ID("item-1"))
		assert.NoError(t, err)
		assert.Equal(t, createdItem.ID, fetched.ID)
		assert.Equal(t, createdItem.Name, fetched.Name)

		// Delete item
		err = client.DeleteItem(ctx, ID("item-1"))
		assert.NoError(t, err)

		// Delete collection
		err = client.DeleteCollection(ctx, ID("collection-1"))
		assert.NoError(t, err)
	})

	// Assert all mocks were called
	client.AssertExpectations(t)
}
