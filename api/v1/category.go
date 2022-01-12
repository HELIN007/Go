package v1

import (
	"GinProject/model"
	"GinProject/utils/errmsg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加分类
func AddCategory(c *gin.Context) {
	var cate model.Category
	err := c.ShouldBindJSON(&cate)
	if err != nil {
		fmt.Println("错误提示：", err)
	}
	//先判断分类是否存在
	code = model.CheckCategory(cate.Name)
	if code == errmsg.SUCCESS {
		model.CreateCategory(&cate)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   cate,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//查询分类列表
func GetCategory(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	//当为0时，抵消掉查询时limit的作用
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	cate, total := model.GetCategory(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   cate,
		"total":  total,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//编辑分类
func EditCategory(c *gin.Context) {
	//从上下文获取到id这个参数（转化为int）
	id, _ := strconv.Atoi(c.Param("id"))
	var cate model.Category
	err := c.ShouldBindJSON(&cate)
	if err != nil {
		fmt.Println("错误提示：", err)
	}
	code = model.CheckCategory(cate.Name)
	if code == errmsg.ErrorCategoryUsed {
		c.Abort()
	} else if code == errmsg.SUCCESS {
		model.EditCategory(id, &cate)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
