package koiApi

import (
	"fmt"
	"strings"
)

// validateItem enforces schema restrictions for Item creation, collecting all validation errors.
func (c *httpClient) validateItem(item *Item) error {
	var errs []string

	// name is required, type string; see components.schemas.Item-item.write.required
	if item.Name == "" {
		errs = append(errs, "item name is required")
	}
	// collection is required, type string or null (IRI); see components.schemas.Item-item.write.required
	if item.Collection == nil || *item.Collection == "" {
		errs = append(errs, "item collection IRI is required")
	}
	// quantity minimum 1, type integer; see components.schemas.Item-item.write.properties.quantity
	if item.Quantity < 1 {
		item.Quantity = 1 // The API says it should use a default of 1, but errors out instead
		//errs = append(errs, "item quantity must be at least 1")
	}
	// visibility enum ["public", "internal", "private"]; see components.schemas.Item-item.write.properties.visibility
	if item.Visibility != "" {
		switch item.Visibility {
		case VisibilityPublic, VisibilityInternal, VisibilityPrivate:
			// Valid
		default:
			errs = append(errs, fmt.Sprintf("invalid visibility: %s; must be public, internal, or private", item.Visibility))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errs, "; "))
	}
	return nil
}

// validateAlbum enforces schema restrictions for Album creation, collecting all validation errors.
func (c *httpClient) validateAlbum(album *Album) error {
	var errs []string
	// title is required, type string; see components.schemas.Album-album.write.required
	if album.Title == "" {
		errs = append(errs, "album title is required")
	}
	// visibility enum ["public", "internal", "private"]; see components.schemas.Album-album.write.properties.visibility
	if album.Visibility != "" {
		switch album.Visibility {
		case VisibilityPublic, VisibilityInternal, VisibilityPrivate:
		default:
			errs = append(errs, fmt.Sprintf("invalid visibility: %s; must be public, internal, or private", album.Visibility))
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errs, "; "))
	}
	return nil
}

// validateChoiceList enforces schema restrictions for ChoiceList creation, collecting all validation errors.
func (c *httpClient) validateChoiceList(choiceList *ChoiceList) error {
	var errs []string
	// choices array of unique strings; see components.schemas.ChoiceList-choiceList.write.properties.choices
	if len(choiceList.Choices) > 0 {
		seen := make(map[string]struct{})
		for _, choice := range choiceList.Choices {
			if _, exists := seen[choice]; exists {
				errs = append(errs, fmt.Sprintf("duplicate choice found: %s", choice))
			}
			seen[choice] = struct{}{}
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errs, "; "))
	}
	return nil
}

// validateCollection enforces schema restrictions for Collection creation, collecting all validation errors.
func (c *httpClient) validateCollection(collection *Collection) error {
	var errs []string
	// title is required, type string; see components.schemas.Collection-collection.write.required
	if collection.Title == "" {
		errs = append(errs, "collection title is required")
	}
	// visibility enum ["public", "internal", "private"]; see components.schemas.Collection-collection.write.properties.visibility
	if collection.Visibility != "" {
		switch collection.Visibility {
		case VisibilityPublic, VisibilityInternal, VisibilityPrivate:
		default:
			errs = append(errs, fmt.Sprintf("invalid visibility: %s; must be public, internal, or private", collection.Visibility))
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errs, "; "))
	}
	return nil
}

// validateDatum enforces schema restrictions for Datum creation, collecting all validation errors.
func (c *httpClient) validateDatum(datum *Datum) error {
	var errs []string
	// type is required, enum; see components.schemas.Datum-datum.write.required
	if datum.DatumType == "" {
		errs = append(errs, "datum type is required")
	} else {
		validTypes := []string{"text", "textarea", "country", "date", "rating", "number", "price", "link", "list", "choice-list", "checkbox", "image", "file", "sign", "video", "blank-line", "section"}
		valid := false
		for _, t := range validTypes {
			if string(datum.DatumType) == t {
				valid = true
				break
			}
		}
		if !valid {
			errs = append(errs, fmt.Sprintf("invalid datum type: %s; must be one of %v", datum.DatumType, validTypes))
		}
	}
	// label is required, type string; see components.schemas.Datum-datum.write.required
	if datum.Label == "" {
		errs = append(errs, "datum label is required")
	}
	// visibility enum ["public", "internal", "private"]; see components.schemas.Datum-datum.write.properties.visibility
	if datum.Visibility != "" {
		switch datum.Visibility {
		case VisibilityPublic, VisibilityInternal, VisibilityPrivate:
		default:
			errs = append(errs, fmt.Sprintf("invalid visibility: %s; must be public, internal, or private", datum.Visibility))
		}
	}
	// currency follows https://schema.org/priceCurrency; see components.schemas.Datum-datum.write.properties.currency
	if datum.Currency != nil && *datum.Currency != "" {
		if !validateCurrency(*datum.Currency) {
			errs = append(errs, fmt.Sprintf("invalid currency code: %s", *datum.Currency))
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errs, "; "))
	}
	return nil
}

// validateField enforces schema restrictions for Field creation, collecting all validation errors.
func (c *httpClient) validateField(field *Field) error {
	var errs []string
	// name is required, type string; see components.schemas.Field-field.write.required
	if field.Name == "" {
		errs = append(errs, "field name is required")
	}
	// position is required, type integer; see components.schemas.Field-field.write.required
	//if field.Position == nil {
	//	errs = append(errs, "field position is required")
	//}
	// type is required, enum; see components.schemas.Field-field.write.required
	if field.FieldType == "" {
		errs = append(errs, "field type is required")
	} else {
		validTypes := []string{"text", "textarea", "country", "date", "rating", "number", "price", "link", "list", "choice-list", "checkbox", "image", "file", "sign", "video", "blank-line", "section"}
		valid := false
		for _, t := range validTypes {
			if string(field.FieldType) == t {
				valid = true
				break
			}
		}
		if !valid {
			errs = append(errs, fmt.Sprintf("invalid field type: %s; must be one of %v", field.FieldType, validTypes))
		}
	}
	// template is required, type string or null (IRI); see components.schemas.Field-field.write.required
	if field.Template == nil {
		errs = append(errs, "field template IRI is required")
	}
	// visibility enum ["public", "internal", "private"]; see components.schemas.Field-field.write.properties.visibility
	if field.Visibility != "" {
		switch field.Visibility {
		case VisibilityPublic, VisibilityInternal, VisibilityPrivate:
		default:
			errs = append(errs, fmt.Sprintf("invalid visibility: %s; must be public, internal, or private", field.Visibility))
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errs, "; "))
	}
	return nil
}

// validateLoan enforces schema restrictions for Loan creation, collecting all validation errors.
func (c *httpClient) validateLoan(loan *Loan) error {
	var errs []string
	// item is required, type string or null (IRI); see components.schemas.Loan-loan.write.required
	if loan.Item == nil {
		errs = append(errs, "loan item IRI is required")
	}
	// lentTo is required, type string; see components.schemas.Loan-loan.write.required
	if loan.LentTo == "" {
		errs = append(errs, "loan lentTo is required")
	}
	// lentAt is required, type string, format date-time; see components.schemas.Loan-loan.write.required
	if loan.LentAt.IsZero() {
		errs = append(errs, "loan lentAt is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errs, "; "))
	}
	return nil
}

// validatePhoto enforces schema restrictions for Photo creation, collecting all validation errors.
func (c *httpClient) validatePhoto(photo *Photo) error {
	var errs []string
	// title is required, type string; see components.schemas.Photo-photo.write.required
	if photo.Title == "" {
		errs = append(errs, "photo title is required")
	}
	// album is required, type string or null (IRI); see components.schemas.Photo-photo.write.required
	if photo.Album == nil {
		errs = append(errs, "photo album IRI is required")
	}
	// visibility enum ["public", "internal", "private"]; see components.schemas.Photo-photo.write.properties.visibility
	if photo.Visibility != "" {
		switch photo.Visibility {
		case VisibilityPublic, VisibilityInternal, VisibilityPrivate:
		default:
			errs = append(errs, fmt.Sprintf("invalid visibility: %s; must be public, internal, or private", photo.Visibility))
		}
	}
	// takenAt type string or null, format date-time; see components.schemas.Photo-photo.write.properties.takenAt
	if photo.TakenAt != nil && photo.TakenAt.IsZero() {
		errs = append(errs, "invalid takenAt: must be a valid date-time or null")
	}
	if len(errs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errs, "; "))
	}
	return nil
}

// validateTag enforces schema restrictions for Tag creation, collecting all validation errors.
func (c *httpClient) validateTag(tag *Tag) error {
	var errs []string
	// label is required, type string; see components.schemas.Tag-tag.write.required
	if tag.Label == "" {
		errs = append(errs, "tag label is required")
	}
	// visibility enum ["public", "internal", "private"]; see components.schemas.Tag-tag.write.properties.visibility
	if tag.Visibility != "" {
		switch tag.Visibility {
		case VisibilityPublic, VisibilityInternal, VisibilityPrivate:
		default:
			errs = append(errs, fmt.Sprintf("invalid visibility: %s; must be public, internal, or private", tag.Visibility))
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errs, "; "))
	}
	return nil
}

// validateTagCategory enforces schema restrictions for TagCategory creation, collecting all validation errors.
func (c *httpClient) validateTagCategory(category *TagCategory) error {
	var errs []string
	// label is required, type string; see components.schemas.TagCategory-tagCategory.write.required
	if category.Label == "" {
		errs = append(errs, "tag category label is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errs, "; "))
	}
	return nil
}

// validateTemplate enforces schema restrictions for Template creation, collecting all validation errors.
func (c *httpClient) validateTemplate(template *Template) error {
	var errs []string
	// name is required, type string; see components.schemas.Template-template.write.required
	if template.Name == "" {
		errs = append(errs, "template name is required")
	}
	if len(errs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errs, "; "))
	}
	return nil
}

// validateWish enforces schema restrictions for Wish creation, collecting all validation errors.
func (c *httpClient) validateWish(wish *Wish) error {
	var errs []string
	// name is required, type string; see components.schemas.Wish-wish.write.required
	if wish.Name == "" {
		errs = append(errs, "wish name is required")
	}
	// wishlist is required, type string or null (IRI); see components.schemas.Wish-wish.write.required
	if wish.Wishlist == nil {
		errs = append(errs, "wish wishlist IRI is required")
	}
	// visibility enum ["public", "internal", "private"]; see components.schemas.Wish-wish.write.properties.visibility
	if wish.Visibility != "" {
		switch wish.Visibility {
		case VisibilityPublic, VisibilityInternal, VisibilityPrivate:
		default:
			errs = append(errs, fmt.Sprintf("invalid visibility: %s; must be public, internal, or private", wish.Visibility))
		}
	}
	// currency follows https://schema.org/priceCurrency; see components.schemas.Wish-wish.write.properties.currency
	if wish.Currency != nil && *wish.Currency != "" {
		if !validateCurrency(*wish.Currency) {
			errs = append(errs, fmt.Sprintf("invalid currency code: %s", *wish.Currency))
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errs, "; "))
	}
	return nil
}

// validateWishlist enforces schema restrictions for Wishlist creation, collecting all validation errors.
func (c *httpClient) validateWishlist(wishlist *Wishlist) error {
	var errs []string
	// name is required, type string; see components.schemas.Wishlist-wishlist.write.required
	if wishlist.Name == "" {
		errs = append(errs, "wishlist name is required")
	}
	// visibility enum ["public", "internal", "private"]; see components.schemas.Wishlist-wishlist.write.properties.visibility
	if wishlist.Visibility != "" {
		switch wishlist.Visibility {
		case VisibilityPublic, VisibilityInternal, VisibilityPrivate:
		default:
			errs = append(errs, fmt.Sprintf("invalid visibility: %s; must be public, internal, or private", wishlist.Visibility))
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errs, "; "))
	}
	return nil
}
