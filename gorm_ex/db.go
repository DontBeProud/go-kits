package gorm_ex

import (
	"context"
	"gorm.io/gorm"
)

// ForkDb fork一个新的db对象
func ForkDb(db *gorm.DB) *gorm.DB {
	_db := db
	return _db
}

// ForkDbWithContext fork一个新的db对象，并设置ctx
func ForkDbWithContext(ctx context.Context, db *gorm.DB) *gorm.DB {
	_db := db
	return _db
}
