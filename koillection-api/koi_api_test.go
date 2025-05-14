package koiApi

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// objectTestConfig defines the configuration for testing an object type's lifecycle.
type objectTestConfig struct {
	name           string
	createFunc     func(ctx context.Context, client Client, obj interface{}) (interface{}, error)
	listFunc       func(ctx context.Context, client Client) ([]interface{}, error)
	deleteFunc     func(ctx context.Context, client Client, id ID) error
	getIDFunc      func(obj interface{}) ID
	sampleObj      interface{}
	requiresParent bool
	parentConfig   *objectTestConfig
}

// testObjectLifecycle tests the lifecycle: list, create, list (verify), delete, list (verify absence), delete non-existent.
func testObjectLifecycle(t *testing.T, client Client, cfg objectTestConfig) {
	ctx := context.Background()

	// Helper to list all objects, handling pagination
	listAll := func() ([]interface{}, error) {
		var allObjs []interface{}
		for page := 1; ; page++ {
			objs, err := cfg.listFunc(ctx, client)
			if err != nil {
				return nil, err
			}
			allObjs = append(allObjs, objs...)
			if len(objs) == 0 {
				break
			}
		}
		return allObjs, nil
	}

	// Create parent object if required (e.g., Collection for Item, Item for Datum)
	var parentID ID
	if cfg.requiresParent {
		t.Run("CreateParent", func(t *testing.T) {
			parentObj, err := cfg.parentConfig.createFunc(ctx, client, cfg.parentConfig.sampleObj)
			assert.NoError(t, err, "Failed to create parent object")
			assert.NotNil(t, parentObj, "Parent object is nil")
			parentID = cfg.parentConfig.getIDFunc(parentObj)
			assert.NotEmpty(t, parentID, "Parent object ID is empty")
		})
		defer func() {
			if parentID != "" {
				err := cfg.parentConfig.deleteFunc(ctx, client, parentID)
				assert.NoError(t, err, "Failed to clean up parent object")
			}
		}()
	}

	// Set parent reference if needed (e.g., Item.Collection, Datum.Item)
	if cfg.requiresParent {
		switch obj := cfg.sampleObj.(type) {
		case *Item:
			obj.Collection = new(string)
			*obj.Collection = "/api/collections/" + string(parentID)
		case *Datum:
			obj.Item = new(string)
			*obj.Item = "/api/items/" + string(parentID)
		}
	}

	// List initial objects
	t.Run("ListInitial", func(t *testing.T) {
		objs, err := listAll()
		assert.NoError(t, err, "Failed to list initial objects")
		t.Logf("Initial %s count: %d", cfg.name, len(objs))
	})

	// Create new object
	var createdObj interface{}
	var createdID ID
	t.Run("Create", func(t *testing.T) {
		var err error
		createdObj, err = cfg.createFunc(ctx, client, cfg.sampleObj)
		assert.NoError(t, err, "Failed to create object")
		assert.NotNil(t, createdObj, "Created object is nil")
		createdID = cfg.getIDFunc(createdObj)
		assert.NotEmpty(t, createdID, "Created object ID is empty")
	})

	// List to verify new object is present
	t.Run("ListAfterCreate", func(t *testing.T) {
		objs, err := listAll()
		assert.NoError(t, err, "Failed to list objects after creation")
		found := false
		for _, obj := range objs {
			if cfg.getIDFunc(obj) == createdID {
				found = true
				break
			}
		}
		assert.True(t, found, "Created object %s not found in list", createdID)
	})

	// Delete the created object
	t.Run("Delete", func(t *testing.T) {
		err := cfg.deleteFunc(ctx, client, createdID)
		assert.NoError(t, err, "Failed to delete object")
	})

	// List to verify object is gone
	t.Run("ListAfterDelete", func(t *testing.T) {
		objs, err := listAll()
		assert.NoError(t, err, "Failed to list objects after deletion")
		for _, obj := range objs {
			assert.NotEqual(t, createdID, cfg.getIDFunc(obj), "Deleted object %s still found in list", createdID)
		}
	})

	// Attempt to delete a non-existent object
	t.Run("DeleteNonExistent", func(t *testing.T) {
		nonExistentID := ID("non-existent-999")
		err := cfg.deleteFunc(ctx, client, nonExistentID)
		assert.Error(t, err, "Expected error when deleting non-existent object %s", nonExistentID)
	})
}

// TestClientLifecycle tests the lifecycle for multiple object types.
func TestClientLifecycle(t *testing.T) {
	// Load credentials
	creds, err := loadCredentials(t)
	if err != nil {
		t.Fatalf("Failed to load credentials: %v", err)
	}
	assert.NotEmpty(t, creds.URL, "Server URL is empty")
	assert.NotEmpty(t, creds.Username, "Username is empty")
	assert.NotEmpty(t, creds.Password, "Password is empty")

	// Initialize client (replace with your actual client implementation)
	client := NewHTTPClient(creds.URL, 30*time.Second)
	ctx := context.Background()

	// Authenticate
	token, err := client.CheckLogin(ctx, creds.Username, creds.Password)
	if err != nil {
		t.Fatalf("Failed to authenticate: %v", err)
	}
	assert.NotEmpty(t, token, "Authentication token is empty")
	// Assume client is configured to use token for subsequent requests

	// Define test configurations for each object type
	configs := []objectTestConfig{
		{
			name: "Album",
			createFunc: func(ctx context.Context, client Client, obj interface{}) (interface{}, error) {
				return client.CreateAlbum(ctx, obj.(*Album))
			},
			listFunc: func(ctx context.Context, client Client) ([]interface{}, error) {
				albums, err := client.ListAlbums(ctx, 1)
				if err != nil {
					return nil, err
				}
				objs := make([]interface{}, len(albums))
				for i, album := range albums {
					objs[i] = album
				}
				return objs, nil
			},
			deleteFunc: func(ctx context.Context, client Client, id ID) error {
				return client.DeleteAlbum(ctx, id)
			},
			getIDFunc: func(obj interface{}) ID {
				return obj.(*Album).ID
			},
			sampleObj: &Album{
				Title:      "Test Album",
				Visibility: VisibilityPublic,
				CreatedAt:  time.Now(),
			},
		},
		{
			name: "Collection",
			createFunc: func(ctx context.Context, client Client, obj interface{}) (interface{}, error) {
				return client.CreateCollection(ctx, obj.(*Collection))
			},
			listFunc: func(ctx context.Context, client Client) ([]interface{}, error) {
				collections, err := client.ListCollections(ctx, 1)
				if err != nil {
					return nil, err
				}
				objs := make([]interface{}, len(collections))
				for i, collection := range collections {
					objs[i] = collection
				}
				return objs, nil
			},
			deleteFunc: func(ctx context.Context, client Client, id ID) error {
				return client.DeleteCollection(ctx, id)
			},
			getIDFunc: func(obj interface{}) ID {
				return obj.(*Collection).ID
			},
			sampleObj: &Collection{
				Title:      "Test Collection",
				Visibility: VisibilityPublic,
				CreatedAt:  time.Now(),
			},
		},
		{
			name: "Item",
			createFunc: func(ctx context.Context, client Client, obj interface{}) (interface{}, error) {
				return client.CreateItem(ctx, obj.(*Item))
			},
			listFunc: func(ctx context.Context, client Client) ([]interface{}, error) {
				items, err := client.ListItems(ctx, 1)
				if err != nil {
					return nil, err
				}
				objs := make([]interface{}, len(items))
				for i, item := range items {
					objs[i] = item
				}
				return objs, nil
			},
			deleteFunc: func(ctx context.Context, client Client, id ID) error {
				return client.DeleteItem(ctx, id)
			},
			getIDFunc: func(obj interface{}) ID {
				return obj.(*Item).ID
			},
			sampleObj: &Item{
				Name:       "Test Item",
				Quantity:   1,
				Visibility: VisibilityPublic,
				CreatedAt:  time.Now(),
			},
			requiresParent: true,
			parentConfig: &objectTestConfig{
				name: "ParentCollection",
				createFunc: func(ctx context.Context, client Client, obj interface{}) (interface{}, error) {
					return client.CreateCollection(ctx, obj.(*Collection))
				},
				deleteFunc: func(ctx context.Context, client Client, id ID) error {
					return client.DeleteCollection(ctx, id)
				},
				getIDFunc: func(obj interface{}) ID {
					return obj.(*Collection).ID
				},
				sampleObj: &Collection{
					Title:      "Parent Collection",
					Visibility: VisibilityPublic,
					CreatedAt:  time.Now(),
				},
			},
		},
		{
			name: "Datum",
			createFunc: func(ctx context.Context, client Client, obj interface{}) (interface{}, error) {
				return client.CreateDatum(ctx, obj.(*Datum))
			},
			listFunc: func(ctx context.Context, client Client) ([]interface{}, error) {
				data, err := client.ListData(ctx, 1)
				if err != nil {
					return nil, err
				}
				objs := make([]interface{}, len(data))
				for i, datum := range data {
					objs[i] = datum
				}
				return objs, nil
			},
			deleteFunc: func(ctx context.Context, client Client, id ID) error {
				return client.DeleteDatum(ctx, id)
			},
			getIDFunc: func(obj interface{}) ID {
				return obj.(*Datum).ID
			},
			sampleObj: &Datum{
				Label:      "Test Datum",
				DatumType:  DatumTypeText,
				Value:      new(string),
				Visibility: VisibilityPublic,
				CreatedAt:  time.Now(),
			},
			requiresParent: true,
			parentConfig: &objectTestConfig{
				name: "ParentItem",
				createFunc: func(ctx context.Context, client Client, obj interface{}) (interface{}, error) {
					item := obj.(*struct {
						Collection *Collection
						Item       *Item
					})
					collection, err := client.CreateCollection(ctx, item.Collection)
					if err != nil {
						return nil, err
					}
					item.Item.Collection = new(string)
					*item.Item.Collection = "/api/collections/" + string(collection.ID)
					createdItem, err := client.CreateItem(ctx, item.Item)
					if err != nil {
						_ = client.DeleteCollection(ctx, collection.ID) // Cleanup
						return nil, err
					}
					return createdItem, nil
				},
				deleteFunc: func(ctx context.Context, client Client, id ID) error {
					return client.DeleteItem(ctx, id)
				},
				getIDFunc: func(obj interface{}) ID {
					return obj.(*Item).ID
				},
				sampleObj: &struct {
					Collection *Collection
					Item       *Item
				}{
					Collection: &Collection{
						Title:      "Parent Collection",
						Visibility: VisibilityPublic,
						CreatedAt:  time.Now(),
					},
					Item: &Item{
						Name:       "Parent Item",
						Quantity:   1,
						Visibility: VisibilityPublic,
						CreatedAt:  time.Now(),
					},
				},
			},
		},
		{
			name: "Tag",
			createFunc: func(ctx context.Context, client Client, obj interface{}) (interface{}, error) {
				return client.CreateTag(ctx, obj.(*Tag))
			},
			listFunc: func(ctx context.Context, client Client) ([]interface{}, error) {
				tags, err := client.ListTags(ctx, 1)
				if err != nil {
					return nil, err
				}
				objs := make([]interface{}, len(tags))
				for i, tag := range tags {
					objs[i] = tag
				}
				return objs, nil
			},
			deleteFunc: func(ctx context.Context, client Client, id ID) error {
				return client.DeleteTag(ctx, id)
			},
			getIDFunc: func(obj interface{}) ID {
				return obj.(*Tag).ID
			},
			sampleObj: &Tag{
				Label:      "Test Tag",
				Visibility: VisibilityPublic,
				CreatedAt:  time.Now(),
			},
		},
	}

	// Run tests for each object type
	for _, cfg := range configs {
		t.Run(cfg.name, func(t *testing.T) {
			testObjectLifecycle(t, client, cfg)
		})
	}
}
