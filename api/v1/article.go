package v1

import (
	"GinProject/model"
	"GinProject/utils/errmsg"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//添加文章
func AddArticle(c *gin.Context) {
	var article model.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		fmt.Println("错误提示：", err)
	}
	code = model.CreateArticle(&article)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   article,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//查询文章列表
func GetArticle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	//当为0时，抵消掉查询时limit的作用
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	article, code, total := model.GetArticle(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   article,
		"total":  total,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//
// GetArticleInfo godoc
// @Summary      查询单个文章
// @Description  查询单个文章接口
// @Tags         文章接口
// @Param        id   path      int            true  "文章编号"  default(1)
// @Success      200  {object}  ResponseUser   "查询文章成功"
// @Failure      400  {object}  ResponseError  "查询文章失败"
// @Router       /article/info/{id}  [get]
func GetArticleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArticleInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//查询分类下的所有文章
func GetCateArt(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("cid"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	//当为0时，抵消掉查询时limit的作用
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code, total := model.GetCateArt(cid, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"total":  total,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//编辑文章
func EditArticle(c *gin.Context) {
	//从上下文获取到id这个参数（转化为int）
	id, _ := strconv.Atoi(c.Param("id"))
	var article model.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		fmt.Println("错误提示：", err)
	}
	code = model.EditArticle(id, &article)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
