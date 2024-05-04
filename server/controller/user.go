package controller

import (
	"book-management-system-server/dto"
	"book-management-system-server/response"
	"book-management-system-server/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v5"
)

type UserCtrl struct{}

// 账号注册
func (u *UserCtrl) Register(c *gin.Context) {
	dto := dto.UserRegisterDto{}

	if err := c.MustBindWith(&dto, binding.JSON); err != nil {
		response.ValidateError(c, err, &dto)
		return
	}

	if err := service.Service.UserService.Register(&dto); err != nil {
		if err.Error() == "已存在该账号" {
			c.String(http.StatusBadRequest, "已存在该账号")
			return
		}
		fmt.Printf("注册账号失败: %v\n", err)
		c.String(http.StatusInternalServerError, "用户注册失败，系统异常")
		return
	}

	c.String(http.StatusOK, "用户注册成功")
}

// 登录
func (u UserCtrl) Login(c *gin.Context) {
	dto := dto.UserLoginDto{}

	if err := c.MustBindWith(&dto, binding.JSON); err != nil {
		response.ValidateError(c, err, &dto)
	}

	// 获取用户信息
	user, err := service.Service.UserService.GetUser(&dto)
	if err != nil {
		if err.Error() == "账号或密码错误" {
			c.String(http.StatusBadRequest, "账号或密码错误")
			return
		}
		fmt.Printf("获取用户信息失败，系统异常: %v\n", err)
		c.String(http.StatusInternalServerError, "获取用户信息失败，系统异常")
		return
	}

	// 生成jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      user.ID,
		"account": user.Account,
		"name":    user.Name,
	})
	tokenString, err := token.SignedString([]byte("test123"))
	if err != nil {
		fmt.Printf("jwt签名异常: %v\n", err)
		c.String(http.StatusInternalServerError, "jwt签名异常，系统异常")
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id":      user.ID,
		"account": user.Account,
		"name":    user.Name,
		"token":   tokenString,
	})
}
