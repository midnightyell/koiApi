package koiApi

import (
	"context"
	"fmt"
	"time"
)

// LogInterface defines methods for interacting with Log resources.
type LogInterface interface {
	Get(ctx context.Context, client Client, logID ...ID) (*Log, error) // HTTP GET /api/logs/{id}
	IRI() string                                                       // /api/logs/{id}
	List(ctx context.Context, client Client) ([]*Log, error)           // HTTP GET /api/logs
}

// Log represents an action or event in Koillection, combining fields for JSON-LD and API interactions.
type Log struct {
	Context       *Context  `json:"@context,omitempty" access:"rw"` // JSON-LD only
	_ID           ID        `json:"@id,omitempty" access:"ro"`      // JSON-LD only
	Type          string    `json:"@type,omitempty" access:"rw"`    // JSON-LD only
	ID            ID        `json:"id,omitempty" access:"ro"`       // Identifier
	LogType       *string   `json:"type,omitempty" access:"rw"`     // Log type
	LoggedAt      time.Time `json:"loggedAt" access:"rw"`           // Log timestamp
	ObjectID      string    `json:"objectId" access:"rw"`           // Object identifier
	ObjectLabel   string    `json:"objectLabel" access:"rw"`        // Object label
	ObjectClass   string    `json:"objectClass" access:"rw"`        // Object class
	ObjectDeleted bool      `json:"objectDeleted" access:"ro"`      // Deletion status
	Owner         *string   `json:"owner,omitempty" access:"ro"`    // Owner IRI
}

// whichID
func (l *Log) whichID(logID ...ID) ID {
	if len(logID) > 0 {
		return logID[0]
	}
	return l.ID
}

// Get
func (l *Log) Get(ctx context.Context, client Client, logID ...ID) (*Log, error) {
	id := l.whichID(logID...)
	return client.GetLog(ctx, id)
}

// IRI
func (l *Log) IRI() string {
	return fmt.Sprintf("/api/logs/%s", l.ID)
}

// List
func (l *Log) List(ctx context.Context, client Client) ([]*Log, error) {
	var allLogs []*Log
	for page := 1; ; page++ {
		logs, err := client.ListLogs(ctx, page)
		if err != nil {
			return nil, fmt.Errorf("failed to list logs on page %d: %w", page, err)
		}
		if len(logs) == 0 {
			break
		}
		allLogs = append(allLogs, logs...)
	}
	return allLogs, nil
}
