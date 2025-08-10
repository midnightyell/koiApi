package koiApi

import (
	"fmt"
	"time"
)

// User represents a user in Koillection, combining fields for JSON-LD and API interactions.
type User struct {
	Context                      Context    `json:"@context,omitempty" access:"rw"`           // JSON-LD only
	_ID                          ID         `json:"@id,omitempty" access:"ro"`                // JSON-LD only
	Type                         string     `json:"@type,omitempty" access:"rw"`              // JSON-LD only
	ID                           ID         `json:"id,omitempty" access:"ro"`                 // Identifier
	Username                     string     `json:"username" access:"rw"`                     // User name
	Email                        string     `json:"email" access:"rw"`                        // Email address
	PlainPassword                string     `json:"plainPassword,omitempty" access:"rw"`      // Password
	Avatar                       string     `json:"avatar,omitempty" access:"rw"`             // Avatar URL
	Currency                     string     `json:"currency" access:"rw"`                     // Currency preference
	Locale                       string     `json:"locale" access:"rw"`                       // Language preference
	Timezone                     string     `json:"timezone" access:"rw"`                     // Timezone preference
	DateFormat                   DateFormat `json:"dateFormat" access:"rw"`                   // Date format preference
	DiskSpaceAllowed             int        `json:"diskSpaceAllowed" access:"rw"`             // Storage limit
	Visibility                   Visibility `json:"visibility" access:"rw"`                   // Visibility level
	LastDateOfActivity           time.Time  `json:"lastDateOfActivity,omitempty" access:"ro"` // Last activity timestamp
	WishlistsFeatureEnabled      bool       `json:"wishlistsFeatureEnabled" access:"rw"`      // Wishlists feature toggle
	TagsFeatureEnabled           bool       `json:"tagsFeatureEnabled" access:"rw"`           // Tags feature toggle
	SignsFeatureEnabled          bool       `json:"signsFeatureEnabled" access:"rw"`          // Signs feature toggle
	AlbumsFeatureEnabled         bool       `json:"albumsFeatureEnabled" access:"rw"`         // Albums feature toggle
	LoansFeatureEnabled          bool       `json:"loansFeatureEnabled" access:"rw"`          // Loans feature toggle
	TemplatesFeatureEnabled      bool       `json:"templatesFeatureEnabled" access:"rw"`      // Templates feature toggle
	HistoryFeatureEnabled        bool       `json:"historyFeatureEnabled" access:"rw"`        // History feature toggle
	StatisticsFeatureEnabled     bool       `json:"statisticsFeatureEnabled" access:"rw"`     // Statistics feature toggle
	ScrapingFeatureEnabled       bool       `json:"scrapingFeatureEnabled" access:"rw"`       // Scraping feature toggle
	SearchInDataByDefaultEnabled bool       `json:"searchInDataByDefaultEnabled" access:"rw"` // Search data toggle
	DisplayItemsNameInGridView   bool       `json:"displayItemsNameInGridView" access:"rw"`   // Grid view name toggle
	SearchResultsDisplayMode     string     `json:"searchResultsDisplayMode" access:"rw"`     // Search display mode
	CreatedAt                    time.Time  `json:"createdAt" access:"ro"`                    // Creation timestamp
	UpdatedAt                    time.Time  `json:"updatedAt,omitempty" access:"ro"`          // Update timestamp

}

func (u *User) Summary() string {
	return fmt.Sprintf("%-40s %s", u.Username, u.ID)
}

// IRI
func (u *User) IRI() string {
	return fmt.Sprintf("/api/users/%s", u.ID)
}

func (u *User) GetID() string {
	return string(u.ID)
}

// Validate checks if the user has a valid username and email.
func (u *User) Validate() error {
	if u.Username == "" {
		return fmt.Errorf("username cannot be empty")
	}
	if u.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}
	return nil
}
