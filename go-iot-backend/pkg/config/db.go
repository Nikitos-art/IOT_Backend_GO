// used for tests that need a db connection
package config

import "github.com/jinzhu/gorm"

func SetDB(database *gorm.DB) {
	db = database
}
