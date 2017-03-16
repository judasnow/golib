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

func create(name string, parentID int) error {
	item := Item{
		Name: name,
		ParentID: parentID,
	}
	if err := db.Create(&item).Error; err != nil {
		return err
	}
	if _, err := rebuildTree(1, 1); err != nil {
		return err
	} else {
		return nil
	}
}

func del(id int) {
	// 删除本部门，同时需要删除所有的子部门
	// rebuild
}

// 仅仅需要查询指定节点左右标记之间的元素即可
func getChildren(id int) ([]Item, error) {
	item := Item{}
	if err := db.Where("id = ?", id).Find(&item).Error; err != nil {
		return []Item{}, err
	}

	items := []Item{}
	if err := db.Where("lft BETWEEN ? AND ?", item.Lft, item.Rgt).Find(&items).Error; err != nil {
		return []Item{}, err
	}

	return items, nil
}

func getChildrenCount(id int) {

}

func getDirectChildren(id int) {

}

func getParents(id int) {

}

func getParent(id int) {

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

	//if _, err := rebuildTree(1, 1); err != nil {
	//	log.Fatal(err)
	//}

	if err := displayTree("1"); err != nil {
		log.Fatal(err)
	}

	//if err := create("2-1", 3); err != nil {
	//	log.Fatal(err)
	//}

}