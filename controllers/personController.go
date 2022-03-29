package controllers

import (
	"net/http"
	"strconv"

	"example.com/sarang-apis/models"
	"example.com/sarang-apis/services"
	"github.com/gin-gonic/gin"
)

type PersonController struct {
	PersonService services.PersonService
}

func New(personservice services.PersonService) PersonController {
	return PersonController{
		PersonService: personservice,
	}
}

func (pc *PersonController) CreatePerson(ctx *gin.Context) {
	var person models.Person
	if err := ctx.ShouldBindJSON(&person); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	err := pc.PersonService.CreatePerson(&person)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (pc *PersonController) GetPerson(ctx *gin.Context) {
	idPerson, err := strconv.Atoi(ctx.Param("id"))
	person, err := pc.PersonService.GetPerson(&idPerson)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, person)
}

func (pc *PersonController) GetAll(ctx *gin.Context) {
	persons, err := pc.PersonService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, persons)
}

func (pc *PersonController) UpdatePerson(ctx *gin.Context) {
	var person models.Person
	if err := ctx.ShouldBindJSON(&person); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	err := pc.PersonService.UpdatePerson(&person)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (pc *PersonController) DeletePerson(ctx *gin.Context) {
	idPerson, erro := strconv.Atoi(ctx.Param("id"))
	if erro != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": erro.Error()})
		return
	}
	err := pc.PersonService.DeletePerson(&idPerson)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (pc *PersonController) RegisterPersonRouters(rg *gin.RouterGroup) {
	personroute := rg.Group("/person")
	personroute.POST("/create", pc.CreatePerson)
	personroute.GET("/get/:id", pc.GetPerson)
	personroute.GET("/getall", pc.GetAll)
	personroute.PATCH("/update", pc.UpdatePerson)
	personroute.DELETE("/delete/:id", pc.DeletePerson)
}
