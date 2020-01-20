package models

import "time"

// AuthToken holds the returned data from a call to the
// CreateTempAuthToken method
type AuthToken struct {
	ID             string `json:"id"`
	CreationDate   string `json:"creation_date"`
	ExpirationDate string `json:"expiration_date"`
}

// ProcessingResult holds the returned value from a call to
// the previously processed file endpoint
type ProcessingResult struct {
	ID            string            `json:"id"`
	Checksum      string            `json:"checksum"`
	ContentLength int64             `json:"content_length"`
	Findings      []string          `json:"findings"`
	CreationDate  time.Time         `json:"creation_date"`
	ContentType   string            `json:"content_type"`
	Metadata      map[string]string `json:"metadata"`
}

// PendingResult
type PendingResult struct {
	ID string `json:"id"`
}

type User struct {
	CreationDate  time.Time `json:"creation_date"`
	LastLoginDate time.Time `json:"last_login_date"`
}
type APIKey struct {
	Active                     bool      `json:"active"`
	CreationDate               time.Time `json:"creation_date"`
	LastSeenDate               time.Time `json:"last_seen_date"`
	DetectionCategoriesEnabled []string  `json:"detection_categories_enabled"`
	Tags                       [] string `json:"tags"`
}

type AccountInfo struct {
	Name             string            `json:"name"`
	Balance          int64             `json:"balance"`
	StartingBalance  int64             `json:"starting_balance"`
	BillingEmail     string            `json:"billing_email"`
	Subscription     string            `json:"subscription"`
	CreationDate     time.Time         `json:"creation_date"`
	ModificationDate time.Time         `json:"modification_date"`
	Users            map[string]User   `json:"users"`
	Keys             map[string]APIKey `json:"keys"`
}
