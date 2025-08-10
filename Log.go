package koiApi

import (
	"fmt"
	"time"
)

// Log represents an action or event in Koillection, combining fields for JSON-LD and API interactions.
type Log struct {
	Context       Context   `json:"@context,omitempty" access:"rw"` // JSON-LD only
	_ID           ID        `json:"@id,omitempty" access:"ro"`      // JSON-LD only
	Type          string    `json:"@type,omitempty" access:"rw"`    // JSON-LD only
	ID            ID        `json:"id,omitempty" access:"ro"`       // Identifier
	LogType       string    `json:"type,omitempty" access:"rw"`     // Log type
	LoggedAt      time.Time `json:"loggedAt" access:"rw"`           // Log timestamp
	ObjectID      string    `json:"objectId" access:"rw"`           // Object identifier
	ObjectLabel   string    `json:"objectLabel" access:"rw"`        // Object label
	ObjectClass   string    `json:"objectClass" access:"rw"`        // Object class
	ObjectDeleted bool      `json:"objectDeleted" access:"ro"`      // Deletion status
	Owner         string    `json:"owner,omitempty" access:"ro"`    // Owner IRI

}

func (l *Log) Summary() string {
	return fmt.Sprintf("%-40s %s", l.ObjectLabel, l.ID)
}

// IRI
func (l *Log) IRI() string {
	return fmt.Sprintf("/api/logs/%s", l.ID)
}

func (l *Log) GetID() string {
	return string(l.ID)
}

func (l *Log) Validate() error {
	return nil
}
