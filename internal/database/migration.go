package database

import (
	"github.com/charlallison/go-rest-api/internal/comment"
	"github.com/jinzhu/gorm"
)

// MigrateDb - Migrates the database and creates the comment table
func MigrateDb(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
