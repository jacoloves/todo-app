package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"todo-backend/prisma/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var client *db.PrismaClient

func init() {
	client = db.NewClient()
}

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	err := client.Connect()
	if err != nil {
		panic(fmt.Errorf("Could not connect to database: %v", err))
	}

	defer func() {
		err := client.Disconnect()
		if err != nil {
			panic(fmt.Errorf("Could not disconnect from database: %v", err))
		}
	}()

	r.GET("/tasks", getTasks)
	r.POST("/tasks", createTask)
	r.PUT("/tasks/:id/toggle", toggleTask)
	r.DELETE("/tasks/:id", deleteTask)

	r.Run(":8081")
}

func getTasks(c *gin.Context) {
	tasks, err := client.Task.FindMany().Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func createTask(c *gin.Context) {
	var task db.TaskModel
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdTask, err := client.Task.CreateOne(
		db.Task.Text.Set(task.Text),
		db.Task.Completed.Set(false),
	).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, createdTask)
}

func toggleTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	task, err := client.Task.FindUnique(
		db.Task.ID.Equals(id),
	).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	updatedTask, err := client.Task.FindUnique(
		db.Task.ID.Equals(id),
	).Update(
		db.Task.Completed.Set(!task.Completed),
	).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}

func deleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	_, err = client.Task.FindUnique(
		db.Task.ID.Equals(id),
	).Delete().Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
