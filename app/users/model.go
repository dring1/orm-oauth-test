package users

import "time"

// User -
type Model struct {
	// ID    uuid.UUID `gorm:"primary_key;type:uuid"`
	Email        string `gorm:"type:varchar(100);primary_key;unique"`
	LastLoggedIn time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
