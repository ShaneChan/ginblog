package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CreateCategory 新增分类
func CreateCategory(data *Category) int {
	// data.Password = ScryptPwd(data.Password)
	err := db.Create(&data).Error

	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}

// CheckCategory 查询分类是否存在
func CheckCategory(name string) int {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)

	if category.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}

	return errmsg.SUCCESS
}

// GetCategory 查询分类列表
func GetCategory(pageSize int, pageNum int) []Category {
	var categories []Category
	offset := (pageNum - 1) * pageSize

	if pageNum == -1 && pageSize == -1 {
		offset = -1
	}

	err := db.Limit(pageSize).Offset(offset).Find(&categories).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}

	return categories
}

// todo 查询分类下的所有文章

// EditCategory 编辑分类
func EditCategory(id int, data *Category) int {
	var maps = make(map[string]interface{})
	var category Category
	maps["name"] = data.Name

	err = db.Model(&category).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}

// DeleteCategory 删除分类
func DeleteCategory(id int) int {
	var category Category
	err := db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}
