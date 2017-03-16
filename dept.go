package main


import (
	"log"
	"fmt"
	"strings"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)


type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	return  s[:l-1], s[l-1]
}

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
	if err := db.Model(&Item{}).Where("parent_id = ?", parentID).Updates(updates).Error; err != nil {
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
	rightStack := make(stack, 0)
	if err := db.Where("lft BETWEEN ? AND ?", rootItem.Lft, rootItem.Rgt).Order("Lft ASC").Find(&iteams).Error; err != nil {
		return err
	}

	for _, item := range iteams {
		if len(rightStack) > 0 {
			var crt int = rightStack[len(rightStack) - 1]
			for {
				if crt >= item.Lft {
					break
				}
				rightStack.Pop()
				crt = rightStack[len(rightStack) - 1]
			}
		}
		fmt.Println(strings.Repeat(" ", len(rightStack)), item.Name)
		rightStack.Push(item.Rgt)
	}

	return nil
}

func create() {

}

func create2() {

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

	//if err := displayTree("2"); err != nil {
	//	log.Fatal(err)
	//}

	if _, err := rebuildTree(1, 1); err != nil {
		log.Fatal(err)
	}
}