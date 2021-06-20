package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goClass/model/request"
	"goClass/model/response"
	"goClass/utils"
)

func Login(c *gin.Context) {
	var l request.Login
	_ = c.ShouldBindJSON(&l)
	if err := utils.Verify(l, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("api", l)
	response.Ok(c)
}
