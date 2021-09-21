package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"Foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext;" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// CreateArticle 新增分类
func CreateArticle(data *Article) int {
	// data.Password = ScryptPwd(data.Password)
	err := db.Create(&data).Error

	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}

// GetArticle 查询文章列表
func GetArticle(pageSize int, pageNum int) ([]Article, int) {
	var articles []Article
	offset := (pageNum - 1) * pageSize

	err := db.Preload("Category").Limit(pageSize).Offset(offset).Find(&articles).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}

	return articles, errmsg.SUCCESS
}

// GetCategoryArticle 查询分类下的所有文章
func GetCategoryArticle(id int, pageSize int, pageNum int) ([]Article, int) {
	var categoryArticles []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&categoryArticles).Error

	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST
	}

	return categoryArticles, errmsg.SUCCESS
}

// todo 查询单个文章
func GetArticleInfo(id int) (Article, int) {
	var article Article
	err = db.Preload("Category").Where("id = ?", id).First(&article).Error

	if err != nil {
		return article, errmsg.ERROR_ARTICLE_NOT_EXIST
	}

	return article, errmsg.SUCCESS
}

// EditArticle 编辑文章
func EditArticle(id int, data *Article) int {
	var maps = make(map[string]interface{})
	var category Category
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Model(&category).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}

// DeleteArticle 删除文章
func DeleteArticle(id int) int {
	var article Article
	err := db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}
