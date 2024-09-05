package controller

import (
	"net/http"
	"sam/my-erp/config"
	repo "sam/my-erp/repository"

	"github.com/gin-gonic/gin"
)

func GetContacts(c *gin.Context) {
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

func CreateContact(c *gin.Context) {

}
