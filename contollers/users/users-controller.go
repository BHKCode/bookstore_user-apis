package users

import (
	"github.com/bhawana/bookstore_user-apis/domain/users"
	"github.com/bhawana/bookstore_user-apis/services"
	"github.com/bhawana/bookstore_user-apis/utils/errors"
	"github.com/gin-gonic/gin"
	"strconv"

	//"io/ioutil"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	/*	byte,err:=ioutil.ReadAll(c.Request.Body)
		if err!=nil{
			//TODO--Handle error
			return
		}
		if err:=json.Unmarshal(byte,&user); err!=nil{
			//TODO--Handle json error
			fmt.Println(err)
			return
		}*/

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		//TODO return bad request to caller
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO--Handle create user error
		c.JSON(saveErr.Status,saveErr)
		//c.JSON(http, saveErr)
		return
	}
	//fmt.Println(user)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId,userError := strconv.ParseInt(c.Param())
	c.String(http.StatusNotImplemented, "Implement Me!")
}
