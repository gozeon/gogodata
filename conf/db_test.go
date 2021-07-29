package conf

import (
	"fmt"
	"gogodata/model"
	"testing"
)

/**
 * bash: go test gogodata/conf -v -run TestDB
 */
func TestDB(t *testing.T) {
	DB.Create(&model.User{Name: "test", Age: 123})
	DB.Delete(&model.User{}, 1)

	var users []model.User
	result := DB.Find(&users)
	fmt.Println(result.RowsAffected)

	DB.Create(&model.Group{Name: "asd", CreateUser: "admin", DataSources: []model.DataSource{
		{Name: "xxx", Data: "xxx"},
		{Name: "aa", Data: "aa"},
	}})
}

/**
 * bash: go test gogodata/conf -v -run TestData
 */
func TestData(t *testing.T) {
	var dataSource model.DataSource
	dataSource.Name = "xxxx"
	dataSource.CreateUser = "admin"
	dataSource.GroupID = 12222

	result := DB.Create(&dataSource)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
	}
	fmt.Println(result.RowsAffected)
}
