package handlers

import (
	"iman-task/api/internal/models"
	"iman-task/api/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	creater service.ServiceCreater
	getter  service.ServiceGetter
	editer  service.ServiceEditor
}

type Register interface {
	Register(router *gin.Engine)
}

func NewHandler(creater service.ServiceCreater, getter service.ServiceGetter, editer service.ServiceEditor) Register {
	return &handler{
		creater: creater,
		getter:  getter,
		editer:  editer,
	}
}

var (
	failed      = "failure"
	success     = "success"
	inValidId   = "invalid post id"
	inValidName = "invalid post name"
)

func (h *handler) Register(router *gin.Engine) {
	router.POST("/create-posts", h.Create)
	router.GET("/post", h.GetOne)
	router.GET("/posts", h.GetSome)
	router.DELETE("/delete-post", h.DeleteOne)
	router.PUT("/update-post", h.UpdateOne)
}

func (h *handler) Create(c *gin.Context) {
	var u *models.URL
	if err := c.BindJSON(&u); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  err.Error(),
		})
		return
	}
	status, err := h.creater.Create(u.URL)
	if err != nil {
		log.Printf("error answer handler create method:--> %v\n", err)
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}

func (h *handler) GetOne(c *gin.Context) {
	var p *models.Post
	if err := c.BindJSON(&p); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  err.Error(),
		})
		return
	}
	post, err := h.getter.GetOne(p)
	if err != nil {
		log.Printf("Error get one method handler: %v\n", err)
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  inValidName,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": post,
	})
}

func (h *handler) GetSome(c *gin.Context) {
	var ids *models.PostsID
	if err := c.BindJSON(&ids); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  err.Error(),
		})
		return
	}
	log.Println(ids.Ids)
	posts, err := h.getter.GetSome(ids.Ids)
	if err != nil {
		log.Printf("Error get some method handler: %v\n", err)
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func (h *handler) UpdateOne(c *gin.Context) {
	var p *models.Post
	if err := c.BindJSON(&p); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  err.Error(),
		})
		return
	}
	status, err := h.editer.UpdateOne(p)
	if err != nil {
		log.Printf("Error update one method handler: %v\n", err)
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  inValidName,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}

func (h *handler) DeleteOne(c *gin.Context) {
	var p *models.Post
	if err := c.BindJSON(&p); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  err.Error(),
		})
		return
	}
	status, err := h.editer.DeleteOne(p.Id)
	if err != nil {
		log.Printf("Error delete one method handler: %v\n", err)
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  inValidName,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}
