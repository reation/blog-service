package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
	"net/http"
)

type Tag struct {}

func NewTag() Tag{
	return Tag{}
}

func (t Tag) Get (c *gin.Context){}
func (t Tag) List (c *gin.Context){
	//c.JSON(200, gin.H{"detail" : global.DatabaseSetting.Password})


	returnInfo := errcode.ServerError
	if returnInfo.StatusCode() != http.StatusOK {
		global.Logger.Info("http response was error")
	}
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}
func (t Tag) Create (c *gin.Context){}
func (t Tag) Update (c *gin.Context){}
func (t Tag) Delete (c *gin.Context){}