package koiApi

import (
	"fmt"
	"strings"
)

// validateItem enforces schema restrictions for Item creation, collecting all validation errors.
func (c *koiClient) validateItem(item *Item) error {
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

// validateLoan enforces schema restrictions for Loan creation, collecting all validation errors.
func (c *koiClient) validateLoan(loan *Loan) error {
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
func (c *koiClient) validatePhoto(photo *Photo) error {
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
func (c *koiClient) validateTag(tag *Tag) error {
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
func (c *koiClient) validateTagCategory(category *TagCategory) error {
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
func (c *koiClient) validateTemplate(template *Template) error {
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
func (c *koiClient) validateWish(wish *Wish) error {
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
func (c *koiClient) validateWishlist(wishlist *Wishlist) error {
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
