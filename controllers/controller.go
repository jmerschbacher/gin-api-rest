package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmerschbacher/gin-api-rest/database"
	"github.com/jmerschbacher/gin-api-rest/models"
	"net/http"
)

func ExibirTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func ExibirAlunoPorId(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"data": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func CriarAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusCreated, &aluno)
}

func ExcluirAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	if database.DB.First(&aluno, id); aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"data": "Aluno não existe"})
		return
	}

	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno excluído com sucesso"})
}

func EditarAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	if database.DB.First(&aluno, id); aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"data": "Aluno não existe"})
		return
	}

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}
