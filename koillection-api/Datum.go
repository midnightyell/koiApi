package koiApi

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/currency"
)

// DatumInterface defines methods for interacting with Datum resources.
type DatumInterface interface {
	Create(ctx context.Context, client Client) (*Datum, error)
	Get(ctx context.Context, client Client, id ID) (*Datum, error)
	List(ctx context.Context, client Client) ([]*Datum, error)
	Update(ctx context.Context, client Client, id ID) (*Datum, error)
	Delete(ctx context.Context, client Client, id ID) error
	ListDatumItems(ctx context.Context, client Client, id ID) ([]*Item, error)
	Validate(ctx context.Context, client Client) error
}

// Datum represents a datum in Koillection, combining read and write fields.
type Datum struct {
	Context             *Context   `json:"@context,omitempty" access:"rw"`            // JSON-LD only
	ID_                 ID         `json:"@id,omitempty" access:"ro"`                 // JSON-LD only (maps to "@id" in JSON, read-only)
	ID                  ID         `json:"id,omitempty" access:"ro"`                  // Maps to "id" in JSON, read-only
	Type                string     `json:"@type,omitempty" access:"rw"`               // JSON-LD only
	DatumType           DatumType  `json:"datumType" access:"rw"`                     // Read and write
	Label               string     `json:"label" access:"rw"`                         // Read and write
	Value               *string    `json:"value,omitempty" access:"rw"`               // Read and write
	Currency            *string    `json:"currency,omitempty" access:"rw"`            // Read and write
	Item                *string    `json:"item,omitempty" access:"rw"`                // Read and write, IRI
	Position            *int       `json:"position,omitempty" access:"rw"`            // Read and write
	Visibility          Visibility `json:"visibility,omitempty" access:"rw"`          // Read and write
	CreatedAt           time.Time  `json:"createdAt" access:"ro"`                     // Read-only
	UpdatedAt           *time.Time `json:"updatedAt,omitempty" access:"ro"`           // Read-only
	Image               *string    `json:"image,omitempty" access:"ro"`               // Read-only
	ImageSmallThumbnail *string    `json:"imageSmallThumbnail,omitempty" access:"ro"` // Read-only
}

// Validate checks the Datum's fields for validity, using ctx for cancellation and client for optional IRI validation.
func (d *Datum) Validate(ctx context.Context, client Client) error {
	// Check for context cancellation
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("validation cancelled: %w", err)
	}

	// Required fields
	if d.DatumType == "" {
		return fmt.Errorf("datumType must not be empty")
	}
	// Assume DatumType is an enum-like type; no specific validation unless enum values are defined
	if d.Label == "" {
		return fmt.Errorf("label must not be empty")
	}

	// Optional fields
	if d.Value != nil {
		if *d.Value == "" {
			return fmt.Errorf("value must not be empty if set")
		}
		if d.DatumType == DatumTypePrice {
			if _, err := strconv.ParseFloat(*d.Value, 64); err != nil {
				return fmt.Errorf("value must be a valid float for price type: %s", *d.Value)
			}
		}
	}

	if d.Currency != nil {
		if *d.Currency == "" {
			return fmt.Errorf("currency must not be empty if set")
		}
		if _, err := currency.ParseISO(*d.Currency); err != nil {
			return fmt.Errorf("currency must be a valid ISO 4217 code: %s", *d.Currency)
		}
	}

	if d.Item != nil {
		if *d.Item == "" {
			return fmt.Errorf("item IRI must not be empty if set")
		}
		if !strings.HasPrefix(*d.Item, "/api/items/") {
			return fmt.Errorf("item IRI must start with /api/items/: %s", *d.Item)
		}
		// Optionally validate Item exists if client is provided
		if client != nil {
			parts := strings.Split(*d.Item, "/")
			if len(parts) < 4 {
				return fmt.Errorf("invalid item IRI format: %s", *d.Item)
			}
			itemID := ID(parts[3])
			_, err := client.GetItem(ctx, itemID)
			if err != nil {
				return fmt.Errorf("invalid item %s: %w", *d.Item, err)
			}
		}
	}

	// Visibility must be a valid value
	switch d.Visibility {
	case VisibilityPublic, VisibilityInternal, VisibilityPrivate, "":
		// Valid or unset (server may set default)
	default:
		return fmt.Errorf("invalid visibility value: %s", d.Visibility)
	}

	// Read-only fields for creation vs. update
	if d.ID == "" && d.ID_ == "" {
		// Creation: read-only fields should be empty
		if d.ID_ != "" {
			return fmt.Errorf("ID_ must be empty for creation")
		}
		if d.ID != "" {
			return fmt.Errorf("ID must be empty for creation")
		}
		if d.Type != "" && d.Type != "Datum" {
			return fmt.Errorf("Type must be empty or 'Datum' for creation: %s", d.Type)
		}
	} else {
		// Update: ID should be non-empty
		if d.ID == "" {
			return fmt.Errorf("ID must not be empty for update")
		}
	}

	return nil
}

// Create calls Client.CreateDatum to create a new Datum.
func (d *Datum) Create(ctx context.Context, client Client) (*Datum, error) {
	return client.CreateDatum(ctx, d)
}

// Get retrieves a Datum by ID using Client.GetDatum.
func (d *Datum) Get(ctx context.Context, client Client, id ID) (*Datum, error) {
	return client.GetDatum(ctx, id)
}

// List retrieves all Data across all pages using Client.ListData.
func (d *Datum) List(ctx context.Context, client Client) ([]*Datum, error) {
	var allData []*Datum
	for page := 1; ; page++ {
		data, err := client.ListData(ctx, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list data on page %d: %w", page, err)
		}
		if len(data) == 0 {
			break
		}
		allData = append(allData, data...)
	}
	return allData, nil
}

// Update updates a Datum by ID using Client.UpdateDatum.
func (d *Datum) Update(ctx context.Context, client Client, id ID) (*Datum, error) {
	return client.UpdateDatum(ctx, id, d)
}

// Delete removes a Datum by ID using Client.DeleteDatum.
func (d *Datum) Delete(ctx context.Context, client Client, id ID) error {
	return client.DeleteDatum(ctx, id)
}

// ListDatumItems retrieves all Items associated with the Datum ID across all pages using Client.ListDatumItems.
func (d *Datum) ListDatumItems(ctx context.Context, client Client, id ID) ([]*Item, error) {
	var allItems []*Item
	for page := 1; ; page++ {
		items, err := client.ListDatumItems(ctx, id, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list items for Datum ID %s on page %d: %w", id, page, err)
		}
		if len(items) == 0 {
			break
		}
		allItems = append(allItems, items...)
	}
	return allItems, nil
}
