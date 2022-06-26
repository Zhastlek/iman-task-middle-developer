package handlers

import (
	"iman-task/api/internal/service"

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
	router.POST("create-posts", h.Create)
	router.GET("/post", h.GetOne)
	router.GET("/posts", h.GetSome)
	router.DELETE("/delete-post", h.DeleteOne)
	router.PUT("/update-post", h.UpdateOne)
}

func (h *handler) Create(c *gin.Context) {

}

func (h *handler) GetOne(c *gin.Context) {
	// var p *models.Post
	// if err := c.BindJSON(&p); err != nil {
	// 	c.AbortWithStatusJSON(400, gin.H{
	// 		"status": failed,
	// 		"error":  err.Error(),
	// 	})
	// 	return
	// }
	// // log.Println("handler--get by name---->", p, p.SearchName)
	// post, err := h.getter.GetOneByName(p.SearchName)
	// if err != nil {
	// 	log.Printf("Error get one by name method handler: %v\n", err)
	// 	c.AbortWithStatusJSON(400, gin.H{
	// 		"status": failed,
	// 		"error":  inValidName,
	// 	})
	// 	return
	// }
	// // log.Println("get by name handler:----->", product)
	// c.JSON(http.StatusOK, gin.H{
	// 	"posts": post,
	// })
}

func (h *handler) GetSome(c *gin.Context) {

}

func (h *handler) UpdateOne(c *gin.Context) {

}

func (h *handler) DeleteOne(c *gin.Context) {

}
