package koiApi

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAlbumLifecycle(t *testing.T) {
	client, := setupClient(t)
	defer cleanupCollections(t, client, )

	// Create Album
	album := &Album{
		Title:      "Test Album " + time.Now().Format("15:04:05"),
		Visibility: VisibilityPublic,
	}
	createdAlbum, err := client.CreateAlbum(album)
	assert.NoError(t, err, "Failed to create album")
	assert.NotEmpty(t, createdAlbum.ID, "Album ID is empty")
	assert.Equal(t, album.Title, createdAlbum.Title, "Album title mismatch")

	// Upload Image
	_, err = createdAlbum.UploadImageByFile(client, "./testdata/picture001.jpg")
	assert.NoError(t, err, "Failed to upload album image")

	// Get Album
	fetchedAlbum, err := client.GetAlbum(createdAlbum.ID)
	assert.NoError(t, err, "Failed to fetch album")
	assert.Equal(t, createdAlbum.ID, fetchedAlbum.ID, "Fetched album ID mismatch")

	// Update Album
	createdAlbum.Title = "Updated Album"
	updatedAlbum, err := client.UpdateAlbum(createdAlbum.ID, createdAlbum)
	assert.NoError(t, err, "Failed to update album")
	assert.Equal(t, "Updated Album", updatedAlbum.Title, "Updated album title mismatch")

	// List Albums
	albums, err := client.ListAlbums(1)
	assert.NoError(t, err, "Failed to list albums")
	assert.True(t, len(albums) > 0, "No albums listed")

	// Delete Album
	err = client.DeleteAlbum(createdAlbum.ID)
	assert.NoError(t, err, "Failed to delete album")
}

func TestChoiceListLifecycle(t *testing.T) {
	client, := setupClient(t)

	// Create ChoiceList
	choiceList := &ChoiceList{
		Name:    "Test ChoiceList " + time.Now().Format("15:04:05"),
		Choices: []string{"Option1", "Option2"},
	}
	createdChoiceList, err := client.CreateChoiceList(choiceList)
	assert.NoError(t, err, "Failed to create choice list")
	assert.NotEmpty(t, createdChoiceList.ID, "ChoiceList ID is empty")
	assert.Equal(t, choiceList.Name, createdChoiceList.Name, "ChoiceList name mismatch")

	// Get ChoiceList
	fetchedChoiceList, err := client.GetChoiceList(createdChoiceList.ID)
	assert.NoError(t, err, "Failed to fetch choice list")
	assert.Equal(t, createdChoiceList.ID, fetchedChoiceList.ID, "Fetched choice list ID mismatch")

	// Update ChoiceList
	createdChoiceList.Choices = append(createdChoiceList.Choices, "Option3")
	updatedChoiceList, err := client.UpdateChoiceList(createdChoiceList.ID, createdChoiceList)
	assert.NoError(t, err, "Failed to update choice list")
	assert.Equal(t, 3, len(updatedChoiceList.Choices), "Updated choice list choices count mismatch")

	// Delete ChoiceList
	err = client.DeleteChoiceList(createdChoiceList.ID)
	assert.NoError(t, err, "Failed to delete choice list")
}

func TestDatumLifecycle(t *testing.T) {
	client, := setupClient(t)
	defer cleanupCollections(t, client, )

	// Create Collection
	collection := &Collection{
		Title:      "Test Collection " + time.Now().Format("15:04:05"),
		Visibility: VisibilityPublic,
	}
	createdCollection, err := client.CreateCollection(collection)
	assert.NoError(t, err, "Failed to create collection")

	// Create Item
	item := &Item{
		Name:       "Test Item",
		Collection: strPtr(createdCollection.IRI()),
		Visibility: VisibilityPublic,
	}
	createdItem, err := client.CreateItem(item)
	assert.NoError(t, err, "Failed to create item")

	// Create Datum
	datum := &Datum{
		Item:       strPtr(createdItem.IRI()),
		DatumType:  DatumTypeText,
		Label:      "Test Field",
		Value:      strPtr("Test Value"),
		Visibility: VisibilityPublic,
	}
	createdDatum, err := client.CreateDatum(datum)
	assert.NoError(t, err, "Failed to create datum")
	assert.NotEmpty(t, createdDatum.ID, "Datum ID is empty")
	assert.Equal(t, "Test Value", *createdDatum.Value, "Datum value mismatch")

	// Upload Image
	_, err = createdDatum.UploadImageByFile(client, "./testdata/picture001.jpg")
	assert.NoError(t, err, "Failed to upload datum image")

	// Get Datum
	fetchedDatum, err := client.GetDatum(createdDatum.ID)
	assert.NoError(t, err, "Failed to fetch datum")
	assert.Equal(t, createdDatum.ID, fetchedDatum.ID, "Fetched datum ID mismatch")

	// Update Datum
	createdDatum.Value = strPtr("Updated Value")
	updatedDatum, err := client.UpdateDatum(createdDatum.ID, createdDatum)
	assert.NoError(t, err, "Failed to update datum")
	assert.Equal(t, "Updated Value", *updatedDatum.Value, "Updated datum value mismatch")

	// Delete Datum
	err = client.DeleteDatum(createdDatum.ID)
	assert.NoError(t, err, "Failed to delete datum")
}

func TestPhotoLifecycle(t *testing.T) {
	client, := setupClient(t)
	defer cleanupCollections(t, client, )

	// Create Album
	album := &Album{
		Title:      "Test Photo Album " + time.Now().Format("15:04:05"),
		Visibility: VisibilityPublic,
	}
	createdAlbum, err := client.CreateAlbum(album)
	assert.NoError(t, err, "Failed to create album")

	// Create Photo
	photo := &Photo{
		Title:      "Test Photo",
		Album:      strPtr(createdAlbum.IRI()),
		Visibility: VisibilityPublic,
	}
	createdPhoto, err := client.CreatePhoto(photo)
	assert.NoError(t, err, "Failed to create photo")
	assert.NotEmpty(t, createdPhoto.ID, "Photo ID is empty")
	assert.Equal(t, photo.Title, createdPhoto.Title, "Photo title mismatch")

	// Upload Image
	_, err = createdPhoto.UploadImageByFile(client, "./testdata/picture001.jpg")
	assert.NoError(t, err, "Failed to upload photo image")

	// Get Photo
	fetchedPhoto, err := client.GetPhoto(createdPhoto.ID)
	assert.NoError(t, err, "Failed to fetch photo")
	assert.Equal(t, createdPhoto.ID, fetchedPhoto.ID, "Fetched photo ID mismatch")

	// Update Photo
	createdPhoto.Title = "Updated Photo"
	updatedPhoto, err := client.UpdatePhoto(createdPhoto.ID, createdPhoto)
	assert.NoError(t, err, "Failed to update photo")
	assert.Equal(t, "Updated Photo", updatedPhoto.Title, "Updated photo title mismatch")

	// List Photos in Album
	photos, err := client.ListAlbumPhotos(createdAlbum.ID, 1)
	assert.NoError(t, err, "Failed to list album photos")
	assert.True(t, len(photos) > 0, "No photos listed in album")

	// Delete Photo
	err = client.DeletePhoto(createdPhoto.ID)
	assert.NoError(t, err, "Failed to delete photo")
}
func TestWishlistLifecycle(t *testing.T) {
	client, := setupClient(t)
	defer cleanupCollections(t, client, )

	// Create Wishlist
	wishlist := &Wishlist{
		Name:       "Test Wishlist " + time.Now().Format("15:04:05"),
		Visibility: VisibilityPublic,
	}
	createdWishlist, err := client.CreateWishlist(wishlist)
	assert.NoError(t, err, "Failed to create wishlist")
	assert.NotEmpty(t, createdWishlist.ID, "Wishlist ID is empty")
	assert.Equal(t, wishlist.Name, createdWishlist.Name, "Wishlist name mismatch")

	// Upload Image
	_, err = createdWishlist.UploadImageByFile(client, "./testdata/picture001.jpg")
	assert.NoError(t, err, "Failed to upload wishlist image")

	// Get Wishlist
	fetchedWishlist, err := client.GetWishlist(createdWishlist.ID)
	assert.NoError(t, err, "Failed to fetch wishlist")
	assert.Equal(t, createdWishlist.ID, fetchedWishlist.ID, "Fetched wishlist ID mismatch")

	// Update Wishlist
	createdWishlist.Name = "Updated Wishlist"
	updatedWishlist, err := client.UpdateWishlist(createdWishlist.ID, createdWishlist)
	assert.NoError(t, err, "Failed to update wishlist")
	assert.Equal(t, "Updated Wishlist", updatedWishlist.Name, "Updated wishlist name mismatch")

	// Delete Wishlist
	err = client.DeleteWishlist(createdWishlist.ID)
	assert.NoError(t, err, "Failed to delete wishlist")
}
package koiApi

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestItemWithChoiceList(t *testing.T) {
	client, := setupClient(t)
	defer cleanupCollections(t, client, )

	// Create first ChoiceList
	choiceList1 := &ChoiceList{
		Name:    "Colors " + time.Now().Format("15:04:05"),
		Choices: []string{"Red", "Blue", "Green"},
	}
	createdChoiceList1, err := client.CreateChoiceList(choiceList1)
	assert.NoError(t, err, "Failed to create first choice list")
	assert.NotEmpty(t, createdChoiceList1.ID, "First ChoiceList ID is empty")
	assert.Equal(t, choiceList1.Name, createdChoiceList1.Name, "First ChoiceList name mismatch")
	assert.Equal(t, []string{"Red", "Blue", "Green"}, createdChoiceList1.Choices, "First ChoiceList choices mismatch")

	// Create second ChoiceList
	choiceList2 := &ChoiceList{
		Name:    "Sizes " + time.Now().Format("15:04:05"),
		Choices: []string{"Small", "Medium", "Large"},
	}
	createdChoiceList2, err := client.CreateChoiceList(choiceList2)
	assert.NoError(t, err, "Failed to create second choice list")
	assert.NotEmpty(t, createdChoiceList2.ID, "Second ChoiceList ID is empty")
	assert.Equal(t, choiceList2.Name, createdChoiceList2.Name, "Second ChoiceList name mismatch")
	assert.Equal(t, []string{"Small", "Medium", "Large"}, createdChoiceList2.Choices, "Second ChoiceList choices mismatch")

	// Create Collection
	collection := &Collection{
		Title:      "Test Collection " + time.Now().Format("15:04:05"),
		Visibility: VisibilityPublic,
	}
	createdCollection, err := client.CreateCollection(collection)
	assert.NoError(t, err, "Failed to create collection")
	assert.NotEmpty(t, createdCollection.ID, "Collection ID is empty")

	// Create Item
	item := &Item{
		Name:       "Test Item",
		Collection: strPtr(createdCollection.IRI()),
		Visibility: VisibilityPublic,
	}
	createdItem, err := client.CreateItem(item)
	assert.NoError(t, err, "Failed to create item")
	assert.NotEmpty(t, createdItem.ID, "Item ID is empty")

	// Create Datum referencing the first ChoiceList
	datum := &Datum{
		Item:        strPtr(createdItem.IRI()),
		DatumType:   DatumTypeChoiceList,
		Label:       "Color Selection",
		Value:       strPtr("Blue"), // Must be one of the choices in ChoiceList1
		ChoiceList:  strPtr(createdChoiceList1.IRI()),
		Visibility:  VisibilityPublic,
	}
	createdDatum, err := client.CreateDatum(datum)
	assert.NoError(t, err, "Failed to create datum")
	assert.NotEmpty(t, createdDatum.ID, "Datum ID is empty")
	assert.Equal(t, DatumTypeChoiceList, createdDatum.DatumType, "Datum type mismatch")
	assert.Equal(t, "Blue", *createdDatum.Value, "Datum value mismatch")
	assert.Equal(t, createdChoiceList1.IRI(), *createdDatum.ChoiceList, "Datum ChoiceList IRI mismatch")

	// Verify ChoiceList is populated correctly
	fetchedChoiceList, err := client.GetChoiceList(createdChoiceList1.ID)
	assert.NoError(t, err, "Failed to fetch choice list")
	assert.Contains(t, fetchedChoiceList.Choices, *createdDatum.Value, "ChoiceList does not contain selected value")

	// Fetch Item and verify Datum
	fetchedItem, err := client.GetItem(createdItem.ID)
	assert.NoError(t, err, "Failed to fetch item")
	data, err := client.ListItemData(createdItem.ID, 1)
	assert.NoError(t, err, "Failed to list item data")
	assert.Len(t, data, 1, "Expected one datum for item")
	assert.Equal(t, createdDatum.ID, data[0].ID, "Datum ID mismatch in item data")

	// Cleanup (datum, item, and choice lists are deleted via collection cleanup)
	err = client.DeleteChoiceList(createdChoiceList1.ID)
	assert.NoError(t, err, "Failed to delete first choice list")
	err = client.DeleteChoiceList(createdChoiceList2.ID)
	assert.NoError(t, err, "Failed to delete second choice list")
}

// strPtr returns a pointer to a string.
func strPtr(s string) *string {
	return &s
}
// setupClient initializes the client with credentials and authenticates.
func setupClient(t *testing.T) (Client, context.Context) {
	creds, err := loadCredentials(t)
	assert.NoError(t, err, "Failed to load credentials")
	client := NewHTTPClient(creds.URL, 30*time.Second)
	:= context.Background()
	token, err := client.CheckLogin(creds.Username, creds.Password)
	assert.NoError(t, err, "Failed to authenticate")
	assert.NotEmpty(t, token, "Authentication token is empty")
	return client, 
}

System: token, "Authentication token is empty")
	return client, 
}

// strPtr returns a pointer to a string.
func strPtr(s string) *string {
	return &s
}