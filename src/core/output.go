package core

import (
	"GoMVC/config"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Output(gin *gin.Context, code int, result interface{}, strings map[string]string) {
	response := Response{
		Code:    code,
		Message: config.GetExceptionMessage(code),
		Data:    result,
	}

	gin.JSON(200, response)
}