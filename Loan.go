package koiApi

import (
	"fmt"
	"time"
)

// Loan represents a loan record in Koillection, combining fields for JSON-LD and API interactions.
type Loan struct {
	Context    *Context   `json:"@context,omitempty" access:"rw"`   // JSON-LD only
	_ID        ID         `json:"@id,omitempty" access:"ro"`        // JSON-LD only
	Type       string     `json:"@type,omitempty" access:"rw"`      // JSON-LD only
	ID         ID         `json:"id,omitempty" access:"ro"`         // Identifier
	Item       *string    `json:"item" access:"rw"`                 // Item IRI
	LentTo     string     `json:"lentTo" access:"rw"`               // Borrower name
	LentAt     time.Time  `json:"lentAt" access:"rw"`               // Loan start date
	ReturnedAt *time.Time `json:"returnedAt,omitempty" access:"rw"` // Loan return date
	Owner      *string    `json:"owner,omitempty" access:"ro"`      // Owner IRI

}

func (a *Loan) Summary() string {
	return fmt.Sprintf("%-40s %s", a.LentTo, a.ID)
}

// IRI
func (l *Loan) IRI() string {
	return fmt.Sprintf("/api/loans/%s", l.ID)
}

func (l *Loan) GetID() string {
	return string(l.ID)
}

func (l *Loan) Validate() error {
	var errs []string
	// item is required, type string or null (IRI); see components.schemas.Loan-loan.write.required
	if l.Item == nil {
		errs = append(errs, "loan item IRI is required")
	}
	// lentTo is required, type string; see components.schemas.Loan-loan.write.required
	if l.LentTo == "" {
		errs = append(errs, "loan lentTo is required")
	}
	// lentAt is required, type string, format date-time; see components.schemas.Loan-loan.write.required
	if l.LentAt.IsZero() {
		errs = append(errs, "loan lentAt is required")
	}
	return validationErrors(&errs)
}
