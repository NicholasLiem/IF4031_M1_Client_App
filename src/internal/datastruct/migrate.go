package datastruct

import (
	"fmt"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, models ...interface{}) {
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			fmt.Println(err)
			return
		}
	}
}
