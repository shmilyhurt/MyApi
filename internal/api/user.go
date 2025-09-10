package api

import (
	"MyApi/internal/model"
	"MyApi/internal/service"
	"MyApi/pkg/database"
	"MyApi/pkg/response"
	"MyApi/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorWithCode(c, http.StatusBadRequest, "参数错误", nil)
		return
	}

	// 检查用户名是否存在
	var count int64
	database.DB.Model(&model.User{}).Where("username = ?", req.Username).Count(&count)
	if count > 0 {
		response.ErrorWithCode(c, http.StatusBadRequest, "用户名已存在", nil)
		return
	}

	// 生成 salt & hash
	salt, _ := utils.GenerateSalt()
	hashedPwd, _ := utils.HashPassword(req.Password, salt)

	user := model.User{
		Username: req.Username,
		Password: hashedPwd,
		Salt:     salt,
	}
	if err := database.DB.Create(&user).Error; err != nil {
		response.ErrorWithCode(c, http.StatusInternalServerError, "注册失败", nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorWithCode(c, http.StatusBadRequest, "参数错误", nil)
		return
	}

	var user model.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		response.ErrorWithCode(c, http.StatusUnauthorized, "用户名或密码错误", nil)
		return
	}

	// 校验密码
	if !utils.CheckPasswordHash(req.Password, user.Salt, user.Password) {
		response.ErrorWithCode(c, http.StatusUnauthorized, "用户名或密码错误", nil)
		return
	}

	// 生成 token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		response.ErrorWithCode(c, http.StatusInternalServerError, "生成 token 失败", nil)
		return
	}

	response.SuccessWithMsg(c, "登录成功", gin.H{
		"token": token,
	})
}

func CreateUser(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
		Pwd  string `json:"pwd"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorWithCode(c, http.StatusBadRequest, "参数错误", nil)
		return
	}

	if err := service.CreateUser(req.Name, req.Pwd); err != nil {
		response.ErrorWithCode(c, http.StatusInternalServerError, "创建用户失败", nil)
		return
	}

	response.SuccessWithMsg(c, "创建成功", nil)
}

func ListUsers(c *gin.Context) {
	users, err := service.ListUsers()
	if err != nil {
		response.ErrorWithCode(c, http.StatusInternalServerError, "查询失败", nil)
		return
	}
	response.SuccessWithMsg(c, "", users)
}
