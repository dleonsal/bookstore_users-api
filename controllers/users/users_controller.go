package users

import (
	"github.com/dleonsal/bookstore_users-api/domain/users"
	"github.com/dleonsal/bookstore_users-api/services"
	"github.com/dleonsal/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context){
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		error := errors.NewBadRequestError("User id should be number")
		c.JSON(error.Status, error)
		return
	}

	result, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)

}

func CreateUser(c *gin.Context)  {
	var user users.User

	if error := c.ShouldBindJSON (&user); error != nil{
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, savErr := services.CreateUser(user)
	if savErr != nil {
		c.JSON(savErr.Status, savErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

/*func SearchUser(c *gin.Context)  {
	c.String(http.StatusNotImplemented, "Implement me!")
}*/