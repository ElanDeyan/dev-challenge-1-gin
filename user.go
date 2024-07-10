package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"name" xml:"name" yaml:"name"`
	Email    string `json:"email" xml:"email" yaml:"email"`
	Message  string `json:"message" xml:"message" yaml:"message"`
	Interest string `json:"interest" xml:"interest" yaml:"interest"`
}

func CreateUser(ctx *gin.Context) {
	var user User

	err := ctx.ShouldBind(&user)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H {"msg": err.Error()})
		return
	}

	successfulMessage := SuccessfulMessage{Message: fmt.Sprintf("Thank %s, for getting in touch and sharing your interests. We look forward to hearing from you soon.", user.Name)}

	ctx.IndentedJSON(http.StatusAccepted, successfulMessage)
}
