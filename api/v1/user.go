package v1

import (
	"GinProject/middleware"
	"GinProject/model"
	"GinProject/utils/errmsg"
	"GinProject/utils/validatoor"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

type ResponseUser struct {
	Status int    `json:"status" example:"200"`
	Msg    string `json:"msg" example:"OK"`
}

type ResponseError struct {
	Status int    `json:"status" example:"500"`
	Msg    string `json:"msg" example:"Error"`
}

type UserInfo struct {
	Username string `json:"username" example:"lin"`     // 用户名
	Password string `json:"password" example:"1233456"` // 密码
	Role     int    `json:"role" example:"2"`           // 权限码
}

// AddUser godoc
// @Summary      新增用户
// @Description  新增用户接口
// @Tags         用户接口
// @Param        userinfo  body      UserInfo       true  "用户信息"
// @Success      200       {object}  ResponseUser   "新增用户成功"
// @Failure      400       {object}  ResponseError  "新增用户失败"
// @Router       /user/add  [post]
func AddUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	//数据验证
	msg, code := validatoor.Validate(&user)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    msg,
		})
		return
	}
	if err != nil {
		fmt.Println("错误提示：", err)
	}
	//先判断用户是否存在
	code = model.CheckUser(user.Username)
	middleware.Infof("新增用户%v:%v，状态是%v", user.Username, user.Password, errmsg.GetErrMsg(code))
	if code == errmsg.SUCCESS {
		model.CreateUser(&user)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	//当为0时，抵消掉查询时limit的作用
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	users, total := model.GetUsers(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   users,
		"total":  total,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//编辑用户
func EditUser(c *gin.Context) {
	//从上下文获取到id这个参数（转化为int）
	id, _ := strconv.Atoi(c.Param("id"))
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("错误提示：", err)
	}
	//code = model.CheckUser(user.Username)
	//if code == errmsg.ErrorUsernameUsed {
	//	c.Abort()
	//} else if code == errmsg.SUCCESS {
	//	model.EditUser(id, &user)
	//}
	model.EditUser(id, &user)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
