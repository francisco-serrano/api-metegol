package utils

import "github.com/jinzhu/gorm"

type Dependencies struct {
	Db *gorm.DB
}
