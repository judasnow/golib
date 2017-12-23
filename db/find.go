package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/juju/errors"
)


func Find(db *gorm.DB, key string, value interface{}, allowNull bool, m interface{}) (interface{}, error) {
	q := db.Where(fmt.Sprintf("%s = ?", key), value).First(&m)
	if q.RecordNotFound() {
		if allowNull {
			return m, nil
		} else {
			return m, errors.New("not null")
		}
	} else {
		queryErr := q.Error
		if queryErr != nil {
			return m, errors.Annotate(queryErr, "查询失败")
		} else {
			return m, nil
		}
	}
}
