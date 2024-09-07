package controller

import (
	"fmt"
	"net/http"
	"sam/my-erp/config"
	"sam/my-erp/model"
	repo "sam/my-erp/repository"

	"github.com/gin-gonic/gin"
)

type ContactController struct {
}

func GetContactController() *ContactController {
	return &ContactController{}
}

func (contactController *ContactController) All(c *gin.Context) {

	fmt.Println(config.GetDb())

	contactRepo := repo.GetContactRepo(config.GetDb())

	contacts, err := contactRepo.All()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success get data",
		"data":    contacts,
	})
}

func (contactController *ContactController) Create(c *gin.Context) {
	var contact model.Contact
	contactRepo := repo.GetContactRepo(config.GetDb())

	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": err,
		})
	}

	if err := contactRepo.Create(&contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data Created",
		"data":    []*model.Contact{&contact},
	})

}
