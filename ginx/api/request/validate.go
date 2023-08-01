package request

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/api/response"
	"github.com/thorraythorray/go-proj/pkg/validator"
)

func FieldsValidate(c *gin.Context, req interface{}) {
	var err error
	if c.Request.Method == "GET" || c.Request.Method == "HEAD" {
		err = c.ShouldBindQuery(req)
	} else {
		err = c.ShouldBindJSON(req)
	}
	if err != nil {
		response.RequestFailed(c, err.Error())
		return
	}
	if req != nil {
		errMsg := validator.ValidateWithSturct(req)
		if errMsg != "" {
			response.RequestFailed(c, errMsg)
			return
		}
	}
}