package datastruct

import "gorm.io/gorm"

func Migrate(db *gorm.DB, models ...interface{}) {
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return
		}
	}
}
