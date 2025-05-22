package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义用户结构体
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// 模拟的用户数据存储
var users = map[string]User{
	"1": {"1", "Alice", "alice@example.com"},
	"2": {"2", "Bob", "bob@example.com"},
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// 获取所有用户
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	// 根据ID获取单个用户
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		if user, ok := users[id]; ok {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		}
	})

	// 添加新用户
	r.POST("/users", func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		users[newUser.ID] = newUser
		c.JSON(http.StatusCreated, newUser)
	})

	// 更新用户信息
	r.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		if _, ok := users[id]; !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		var updatedUser User
		if err := c.ShouldBindJSON(&updatedUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		users[id] = updatedUser
		c.JSON(http.StatusOK, updatedUser)
	})

	// 删除用户
	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		if _, ok := users[id]; ok {
			delete(users, id)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		}
	})

	// 启动服务，默认在0.0.0.0:8080启动服务
	r.Run()
}