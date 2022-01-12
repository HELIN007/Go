package model

import (
	"GinProject/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"ID"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ErrorCategoryUsed
	}
	return errmsg.SUCCESS
}

//新建分类
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询分类列表，分页显示
func GetCategory(pageSize int, pageNum int) ([]Category, int64) {
	var cate []Category
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	err := db.Model(&cate).Count(&total).Limit(pageSize).Offset(offset).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

//编辑分类
func EditCategory(id int, data *Category) int {
	//根据 `struct` 更新属性，只会更新非零值的字段
	//根据 `map` 更新属性
	var cate Category
	maps := map[string]interface{}{
		"name": data.Name,
	}
	//maps["name"] = data.Username
	err = db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//TODO 查询分类下的所以文章

//删除分类
func DeleteCategory(id int) int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
