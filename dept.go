package main


import (
	"log"
	"fmt"
	"strings"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	lls "github.com/emirpasic/gods/stacks/linkedliststack"
)

// mptt with gorm

var db *gorm.DB

type Item struct {
	ID       int
	ParentID int
	Name     string
	Lft      int
	Rgt      int
}

func rebuildTree(parentID int, left int) (int, error) {
	var right int = left + 1
	items := []Item{}
	if err := db.Where("parent_id = ?", parentID).Find(&items).Error; err != nil {
		return 0, err
	}

	for _, item := range items {
		var err error
		right, err = rebuildTree(item.ID, right)
		if err != nil {
			return 0, err
		}
	}

	updates := map[string]interface{}{
		"lft": left,
		"rgt": right,
	}
	if err := db.Model(&Item{}).Where("id = ?", parentID).Updates(updates).Error; err != nil {
		return 0, err
	}

	return right + 1, nil
}

func displayTree(root string) error {
	rootItem := Item{}
	if err := db.Where("name = ?", root).First(&rootItem).Error; err != nil {
		return err
	}

	iteams := []Item{}
	rightStack := lls.New()
	if err := db.Where("lft BETWEEN ? AND ?", rootItem.Lft, rootItem.Rgt).Order("Lft ASC").Find(&iteams).Error; err != nil {
		return err
	}

	for _, item := range iteams {
		if rightStack.Size() > 0 {
			crt, _ := rightStack.Peek()
			for {
				if crt.(int) >= item.Lft {
					break
				}
				rightStack.Pop()
				crt, _ = rightStack.Peek()
			}
		}
		fmt.Println(strings.Repeat("  ", rightStack.Size()), item.Name)
		rightStack.Push(item.Rgt)
	}

	return nil
}

func create() {

}

func del() {

}

func getChildren() {

}

func getChildrenCount() {

}

func getDirectChildren() {

}

func getParents() {

}

func getParent() {

}

func main() {
	var dbErr error
	db, dbErr = gorm.Open("mysql", "vagrant:vagrant@(vagrant:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if dbErr != nil {
		log.Printf("err: %+v", dbErr)
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Item{})
	db.LogMode(true)

	if _, err := rebuildTree(1, 1); err != nil {
		log.Fatal(err)
	}

	if err := displayTree("1"); err != nil {
		log.Fatal(err)
	}

}