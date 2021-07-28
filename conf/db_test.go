package conf

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
)

type User struct {
	gorm.Model
	Name string
	Age int
}

/**
 * bash: go test gogodata/conf -v -run TestDB
 */
func TestDB(t *testing.T) {
	DB.Create(&User{Name: "test", Age: 123})
	DB.Delete(&User{}, 1)

	var users []User
	result := DB.Find(&users)
	fmt.Println(result.RowsAffected)
}