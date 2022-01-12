package model

import (
	"GinProject/utils/errmsg"

	"gorm.io/gorm"
)

// Article Cid为1时，显示第一个category
type Article struct {
	//必须确保传入的cid值，在category表中是存在的
	//Category是父类结构体
	Category Category `gorm:"foreignKey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// CreateArticle 新建文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetArticleInfo 查询单个文章
func GetArticleInfo(id int) (Article, int) {
	var article Article
	err = db.Preload("Category").Where("id = ?", id).Find(&article).Error
	if err != nil {
		return article, errmsg.ErrorArticleExist
	}
	return article, errmsg.SUCCESS
}

// GetArticle 查询文章列表，分页显示
func GetArticle(pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	//预加载结构体里面的其他结构体
	err := db.Preload("Category").Model(&articleList).Count(&total).Limit(pageSize).Offset(offset).Find(&articleList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}

// EditArticle 编辑文章
func EditArticle(id int, data *Article) int {
	//根据 `struct` 更新属性，只会更新非零值的字段
	//根据 `map` 更新属性
	var article Article
	maps := map[string]interface{}{
		"title":   data.Title,
		"cid":     data.Cid,
		"desc":    data.Desc,
		"content": data.Content,
		"img":     data.Img,
	}
	//maps["name"] = data.Username
	err = db.Model(&article).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCateArt 查询分类下的所有文章
func GetCateArt(cid int, pageSize int, pageNum int) ([]Article, int, int64) {
	var cateArt []Article
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	err = db.Preload("Category").Where("cid = ?", cid).Model(&cateArt).Count(&total).Limit(pageSize).Offset(offset).Find(&cateArt).Error
	if err != nil {
		return nil, errmsg.ErrorCategoryNotExist, 0
	}
	return cateArt, errmsg.SUCCESS, total
}

// DeleteArticle 删除文章
func DeleteArticle(id int) int {
	var article Article
	err = db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
