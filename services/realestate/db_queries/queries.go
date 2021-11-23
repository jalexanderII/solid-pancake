package db_queries

import "gorm.io/gorm"

func NeighborhoodInQuery(neighborhoods []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("neighborhood IN (?)", neighborhoods)
	}
}

func PriceBetween(min int, max int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("rent BETWEEN ? AND ?", min, max)
	}
}

func BedsAtMost(beds int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("beds <= ?", beds)
	}
}

func BathsAtMost(baths int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("baths <= ?", baths)
	}
}
