package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goClass/model/request"
	"net/http"
)

func Login(c *gin.Context) {
	var l request.Login
	_ = c.ShouldBindJSON(&l)
	fmt.Println("api", l)
	c.JSON(http.StatusOK, l)
}
