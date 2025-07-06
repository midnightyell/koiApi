package koiApi

import (
	"context"
	"fmt"
	"time"
)

// LoanInterface defines methods for interacting with Loan resources.
type LoanInterface interface {
	Create(ctx context.Context, client Client) (*Loan, error)                // HTTP POST /api/loans
	Delete(ctx context.Context, client Client, loanID ...ID) error           // HTTP DELETE /api/loans/{id}
	Get(ctx context.Context, client Client, loanID ...ID) (*Loan, error)     // HTTP GET /api/loans/{id}
	GetItem(ctx context.Context, client Client, loanID ...ID) (*Item, error) // HTTP GET /api/loans/{id}/item
	IRI() string                                                             // /api/loans/{id}
	List(ctx context.Context, client Client) ([]*Loan, error)                // HTTP GET /api/loans
	Patch(ctx context.Context, client Client, loanID ...ID) (*Loan, error)   // HTTP PATCH /api/loans/{id}
	Update(ctx context.Context, client Client, loanID ...ID) (*Loan, error)  // HTTP PUT /api/loans/{id}
	Summary() string
}

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

// whichID
func (l *Loan) whichID(loanID ...ID) ID {
	if len(loanID) > 0 {
		return loanID[0]
	}
	return l.ID
}

// Create
func (l *Loan) Create(ctx context.Context, client Client) (*Loan, error) {
	return client.CreateLoan(ctx, l)
}

// Delete
func (l *Loan) Delete(ctx context.Context, client Client, loanID ...ID) error {
	id := l.whichID(loanID...)
	return client.DeleteLoan(ctx, id)
}

// Get
func (l *Loan) Get(ctx context.Context, client Client, loanID ...ID) (*Loan, error) {
	id := l.whichID(loanID...)
	return client.GetLoan(ctx, id)
}

// GetItem
func (l *Loan) GetItem(ctx context.Context, client Client, loanID ...ID) (*Item, error) {
	id := l.whichID(loanID...)
	return client.GetLoanItem(ctx, id)
}

// IRI
func (l *Loan) IRI() string {
	return fmt.Sprintf("/api/loans/%s", l.ID)
}

// List
func (l *Loan) List(ctx context.Context, client Client) ([]*Loan, error) {
	var allLoans []*Loan
	for page := 1; ; page++ {
		loans, err := client.ListLoans(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to list loans: %w", err)
		}
		if len(loans) == 0 {
			break
		}
		allLoans = append(allLoans, loans...)
	}
	return allLoans, nil
}

// Patch
func (l *Loan) Patch(ctx context.Context, client Client, loanID ...ID) (*Loan, error) {
	id := l.whichID(loanID...)
	return client.PatchLoan(ctx, id, l)
}

// Update
func (l *Loan) Update(ctx context.Context, client Client, loanID ...ID) (*Loan, error) {
	id := l.whichID(loanID...)
	return client.UpdateLoan(ctx, id, l)
}
