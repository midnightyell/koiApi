package koiApi

import (
	"context"
	"fmt"
)

// DeleteAllData deletes all accessible data from the Koillection database.
// It returns an error if any deletions fail, but continues processing to maximize cleanup.
// The error contains a list of individual failures.
func (c *httpClient) DeleteAllData(ctx context.Context) error {
	var errs []error

	// Helper function to append errors
	addError := func(err error, resource, id string) {
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to delete %s %s: %w", resource, id, err))
		}
	}

	// Delete resources in dependency order: children first, then parents

	// 1. Delete Photos (depend on Albums)
	photos, err := c.ListPhotos(ctx)
	addError(err, "photos list", "")
	for _, photo := range photos {
		err := c.DeletePhoto(ctx, photo.ID)
		addError(err, "photo", string(photo.ID))
	}

	// 2. Delete Wishes (depend on Wishlists)
	wishes, err := c.ListWishes(ctx)
	addError(err, "wishes list", "")
	for _, wish := range wishes {
		err := c.DeleteWish(ctx, wish.ID)
		addError(err, "wish", string(wish.ID))
	}

	// 3. Delete Items (depend on Collections, have Loans, Tags, Data)
	items, err := c.ListItems(ctx)
	addError(err, "items list", "")
	for _, item := range items {
		// Delete associated Loans
		loans, err := c.ListItemLoans(ctx, item.ID)
		addError(err, "loans list for item", string(item.ID))
		for _, loan := range loans {
			err := c.DeleteLoan(ctx, loan.ID)
			addError(err, "loan", string(loan.ID))
		}

		// Delete associated Data
		data, err := c.ListItemData(ctx, item.ID)
		addError(err, "data list for item", string(item.ID))
		for _, datum := range data {
			err := c.DeleteDatum(ctx, datum.ID)
			addError(err, "datum", string(datum.ID))
		}

		// Note: Tags are not deleted here as they may be shared; item-tag relations are cleared by deleting the item
		err = c.DeleteItem(ctx, item.ID)
		addError(err, "item", string(item.ID))
	}

	// 4. Delete Albums (have children and photos)
	albums, err := c.ListAlbums(ctx)
	addError(err, "albums list", "")
	for _, album := range albums {
		// Delete child albums recursively
		children, err := c.ListAlbumChildren(ctx, album.ID)
		addError(err, "album children list", string(album.ID))
		for _, child := range children {
			err := c.DeleteAlbum(ctx, child.ID)
			addError(err, "child album", string(child.ID))
		}
		err = c.DeleteAlbum(ctx, album.ID)
		addError(err, "album", string(album.ID))
	}

	// 5. Delete Wishlists (have children and wishes)
	wishlists, err := c.ListWishlists(ctx)
	addError(err, "wishlists list", "")
	for _, wishlist := range wishlists {
		// Delete child wishlists recursively
		children, err := c.ListWishlistChildren(ctx, wishlist.ID)
		addError(err, "wishlist children list", string(wishlist.ID))
		for _, child := range children {
			err := c.DeleteWishlist(ctx, child.ID)
			addError(err, "child wishlist", string(child.ID))
		}
		err = c.DeleteWishlist(ctx, wishlist.ID)
		addError(err, "wishlist", string(wishlist.ID))
	}

	// 6. Delete Collections (have children, items, data)
	collections, err := c.ListCollections(ctx)
	addError(err, "collections list", "")
	for _, collection := range collections {
		// Delete child collections recursively
		children, err := c.ListCollectionChildren(ctx, collection.ID)
		addError(err, "collection children list", string(collection.ID))
		for _, child := range children {
			err := c.DeleteCollection(ctx, child.ID)
			addError(err, "child collection", string(child.ID))
		}
		// Delete associated Data
		data, err := c.ListCollectionData(ctx, collection.ID)
		addError(err, "data list for collection", string(collection.ID))
		for _, datum := range data {
			err := c.DeleteDatum(ctx, datum.ID)
			addError(err, "datum", string(datum.ID))
		}
		err = c.DeleteCollection(ctx, collection.ID)
		addError(err, "collection", string(collection.ID))
	}

	// 7. Delete Templates (have Fields)
	templates, err := c.ListTemplates(ctx)
	addError(err, "templates list", "")
	for _, template := range templates {
		// Delete associated Fields
		fields, err := c.ListTemplateFields(ctx, template.ID)
		addError(err, "fields list for template", string(template.ID))
		for _, field := range fields {
			err := c.DeleteField(ctx, field.ID)
			addError(err, "field", string(field.ID))
		}
		err = c.DeleteTemplate(ctx, template.ID)
		addError(err, "template", string(template.ID))
	}

	// 8. Delete Tags (depend on Tag Categories)
	tags, err := c.ListTags(ctx)
	addError(err, "tags list", "")
	for _, tag := range tags {
		err := c.DeleteTag(ctx, tag.ID)
		addError(err, "tag", string(tag.ID))
	}

	// 9. Delete Tag Categories
	tagCategories, err := c.ListTagCategories(ctx)
	addError(err, "tag categories list", "")
	for _, category := range tagCategories {
		err := c.DeleteTagCategory(ctx, category.ID)
		addError(err, "tag category", string(category.ID))
	}

	// 10. Delete Choice Lists
	choiceLists, err := c.ListChoiceLists(ctx)
	addError(err, "choice lists list", "")
	for _, choiceList := range choiceLists {
		err := c.DeleteChoiceList(ctx, choiceList.ID)
		addError(err, "choice list", string(choiceList.ID))
	}

	// 11. Delete Inventories
	inventories, err := c.ListInventories(ctx)
	addError(err, "inventories list", "")
	for _, inventory := range inventories {
		err := c.DeleteInventory(ctx, inventory.ID)
		addError(err, "inventory", string(inventory.ID))
	}

	/* // 12. Delete Logs (if accessible)
	logs, err := c.ListLogs(ctx, 1)
	addError(err, "logs list", "")
	for _, log := range logs {
		err := c.DeleteLog(ctx, log.ID)
		addError(err, "log", string(log.ID))
	}
	*/
	// Note: Users are not deleted as the API only provides GET operations for users

	// Combine errors if any
	if len(errs) > 0 {
		return fmt.Errorf("encountered %d errors during deletion: %v", len(errs), errs)
	}

	return nil
}
