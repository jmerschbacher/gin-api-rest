package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmerschbacher/gin-api-rest/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/aluno", controllers.ExibirTodosAlunos)
	r.GET("/aluno/:id", controllers.ExibirAlunoPorId)
	r.POST("/aluno", controllers.CriarAluno)
	r.DELETE("/aluno/:id", controllers.ExcluirAluno)
	r.PATCH("/aluno/:id", controllers.EditarAluno)
	err := r.Run()
	if err != nil {
		panic(err.Error())
	}
}
