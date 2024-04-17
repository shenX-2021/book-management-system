package controller

import (
	"book-management-system-server/dto"
	"book-management-system-server/response"
	"book-management-system-server/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookCtrl struct{}

// 新建新图书
func (ctrl *BookCtrl) Create(c *gin.Context) {
	var dto dto.BookDataDto
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		response.ValidateError(c, err, &dto)
		return
	}
	err = service.Service.BookService.Create(&dto)
	if err != nil {
		if err.Error() == "已存在相同ISBN的书籍" {
			c.String(http.StatusBadRequest, "已存在相同ISBN的书籍")
			return
		}
		fmt.Printf("新增书籍失败: %v", err)
		c.String(http.StatusInternalServerError, "新增书籍失败，系统异常")
		return
	}
	c.String(http.StatusOK, "create success")
}

// 获取图书列表
func (ctrl *BookCtrl) List(c *gin.Context) {

	name := c.Query("name")

	list, err := service.Service.BookService.List(name)
	if err != nil {
		fmt.Printf("获取书籍列表失败: %v", err)
		c.String(http.StatusInternalServerError, "获取书籍列表失败，系统异常")
	}
	c.JSON(http.StatusOK, list)
}

func (ctrl *BookCtrl) Edit(c *gin.Context) {
	var dto *dto.EdtiBookDataDto
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 0, 0)
	if err != nil {
		response.BadRequest(c, "错误的id值")
		return
	}

	err = c.ShouldBindJSON(&dto)
	if err != nil {
		response.ValidateError(c, err, &dto)
		return
	}

	err = service.Service.BookService.Edit(id, dto)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	c.String(http.StatusOK, "编辑图书成功")
}

// 删除图书
func (ctrl *BookCtrl) Delete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 0, 0)
	if err != nil {
		response.BadRequest(c, "错误的id值")
		return
	}

	err = service.Service.BookService.Delete(id)
	if err != nil {
		fmt.Printf("删除id为%d的书籍失败: %v", id, err)
		c.String(http.StatusInternalServerError, "删除书籍失败，系统异常")
	}
	c.String(http.StatusOK, "删除成功")
}
