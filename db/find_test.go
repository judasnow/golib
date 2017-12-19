package db

import (
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/juju/errors"
)

func init() {
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

func TestFind(t *testing.T) {
	ProductFind("code", "L1212", false)
}
