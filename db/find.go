package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"errors"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func ProductFind(key string, value interface{}, allowNull bool) (Product, error) {
	product := Product{}
	i, findErr := Find(key, value, allowNull, product)
	if p, ok := i.(Product); !ok {
		return Product{}, errors.New("返回类型错误")
	} else {
		return p, findErr
	}
}

func Find(key string, value interface{}, allowNull bool, m interface{}) (interface{}, error) {
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
			return m, queryErr
		} else {
			return m, nil
		}
	}
}

func Main() {
	var err error
	db, err = gorm.Open("mysql", "vagrant:vagrant@(vagrant:3306)/golib?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	// Migrate the schema
	db.AutoMigrate(&Product{})

	p, err := ProductFind("code", "L1212", true)
	fmt.Println(p)
}
